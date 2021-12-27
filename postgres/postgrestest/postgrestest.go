package postgrestest

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/Saser/pdp/postgres"
)

func FromImage(t testing.TB, ctx context.Context, path string) *postgres.Database {
	t.Helper()
	runfile, err := bazel.Runfile(path)
	if err != nil {
		t.Fatalf("error finding runfile %q. Make sure it is added as a data dependency of the test.", path)
	}
	data, err := os.ReadFile(runfile)
	if err != nil {
		t.Fatalf("error reading image archive for %q: %v", path, err)
	}
	buf := bytes.NewReader(data)
	dc, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("error creating Docker client: %v", err)
	}
	t.Cleanup(func() {
		if err := dc.Close(); err != nil {
			t.Errorf("error closing Docker client: %v", err)
		}
	})
	res, err := dc.ImageLoad(ctx, buf, true /*quiet*/)
	if err != nil {
		t.Fatalf("error loading Docker image: %v", err)
	}
	t.Cleanup(func() {
		if err := dc.Close(); err != nil {
			t.Errorf("error closing ImageLoadResponse: %v", err)
		}
	})
	if _, err := io.ReadAll(res.Body); err != nil {
		t.Errorf("error reading ImageLoadResponse body: %v", err)
	}
	// Construct the image name by prepending "bazel" and replacing "/image.tar"
	// with ":image". For example:
	//     postgres/postgrestest/testschema_image.tar
	// becomes
	//     bazel/postgres/postgrestest:testschema_image
	idx := strings.LastIndex(path, "/")
	if idx == -1 {
		panic(errors.New("did not find / in image archive path -- this is a bug in postgrestest"))
	}
	tag := strings.TrimSuffix(path[idx+1:], ".tar") // everything after the last slash except ".tar"
	imageName := "bazel/" + path[:idx] + ":" + tag
	cfg := &container.Config{
		Image: imageName,
	}
	containerName := strings.TrimSuffix(strings.ReplaceAll(path, "/", "_"), ".tar")
	cb, err := dc.ContainerCreate(ctx, cfg, nil, nil, nil, containerName)
	if err != nil {
		t.Fatalf("error creating container: %v", err)
	}
	t.Cleanup(func() {
		if err := dc.ContainerRemove(ctx, cb.ID, types.ContainerRemoveOptions{
			RemoveVolumes: true,
		}); err != nil {
			t.Errorf("error removing container: %v", err)
		}
	})
	if err := dc.ContainerStart(ctx, cb.ID, types.ContainerStartOptions{}); err != nil {
		t.Fatalf("error starting container: %v", err)
	}
	t.Cleanup(func() {
		if err := dc.ContainerStop(ctx, cb.ID, nil); err != nil {
			t.Fatalf("error stopping container: %v", err)
		}
	})
	// TODO: need to do something useful here to check that the container
	// started correctly and can begin serving connections. Failure scenarios:
	// * schema is wrong
	// * environment variables are not set or have invalid values
	return nil
}
