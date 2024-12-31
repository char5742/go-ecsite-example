package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/char5742/ecsite-sample/internal/app"
	"github.com/char5742/ecsite-sample/pkg/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("Error running app: %v\n", err)
	}
}

func run(ctx context.Context) error {
	cfg := config.LoadConfig()
	l, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("Listening on %s\n", url)
	mux := app.NewMux()
	s := app.NewServer(l, mux)
	return s.Run(ctx)
}
