package postgrestest

import (
	"context"
	"os"
	"testing"

	"github.com/bazelbuild/rules_go/go/tools/bazel"

	"github.com/Saser/pdp/postgres"
)

func FromImage(t testing.TB, ctx context.Context, path string) *postgres.Database {
	t.Helper()
	runfile, err := bazel.Runfile(path)
	if err != nil {
		t.Fatalf("error finding runfile %q. Make sure it is added as a data dependency of the test.", path)
	}
	f, err := os.Open(runfile)
	if err != nil {
		t.Fatalf("error opening image archive for %q: %v", path, err)
	}
	t.Cleanup(func() {
		if err := f.Close(); err != nil {
			t.Errorf("error closing image archive file handle: %v", err)
		}
	})
	t.Fatal("postgrestest: from image: not implemented")
	return nil
}
