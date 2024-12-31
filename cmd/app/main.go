package app

import (
	"context"
	"log"

	"github.com/char5742/ecsite-sample/pkg/config"
)

func main() {
	cfg := config.LoadConfig()
	log.Printf("Starting %s on port %s\n", cfg.AppName, cfg.Port)
	if err := run(context.Background()); err != nil {
		log.Printf("Error running app: %v\n", err)
	}
}
