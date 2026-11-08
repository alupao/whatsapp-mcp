package main

import (
	"context"
	"path/filepath"
	"testing"

	"go.mau.fi/whatsmeow/store/sqlstore"
)

func TestUpgradedSQLiteStore(t *testing.T) {
	ctx := context.Background()
	dsn := "file:" + filepath.Join(t.TempDir(), "session.db") + "?_foreign_keys=on"
	for i := 0; i < 2; i++ {
		container, err := sqlstore.New(ctx, "sqlite3", dsn, nil)
		if err != nil {
			t.Fatalf("open %d: %v", i, err)
		}
		device, err := container.GetFirstDevice(ctx)
		closeErr := container.Close()
		if err != nil {
			t.Fatalf("get fresh device %d: %v", i, err)
		}
		if device == nil || device.ID != nil {
			t.Fatalf("expected an unpaired device on open %d", i)
		}
		if closeErr != nil {
			t.Fatalf("close %d: %v", i, closeErr)
		}
	}
}
