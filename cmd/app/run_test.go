package app

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/char5742/ecsite-sample/pkg/config"
	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	in := "message"
	cfg := config.LoadConfig()
	addr := "http://localhost:" + cfg.Port
	rsp, err := http.Get(addr + in)
	if err != nil {
		t.Errorf("Error making GET request: %v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}
	want := "Hello, " + in + "!"
	if string(got) != want {
		t.Errorf("want %q, got %q", want, got)
	}

	cancel()

	if err := eg.Wait(); err != nil {
		t.Errorf("Error running app: %v", err)
	}
}
