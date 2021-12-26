package postgrestest

import (
	"context"
	"testing"
)

func TestFromImage(t *testing.T) {
	ctx := context.Background()
	db := FromImage(t, ctx, "postgres/postgrestest/testschema_image.tar")
	if _, err := db.Get(ctx); err != nil {
		t.Fatal(err)
	}
}
