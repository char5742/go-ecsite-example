package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/char5742/ecsite-sample/pkg/config"
	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	cfg := config.GetConfig()
	s := &http.Server{
		Addr: cfg.Port,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error starting server: %v\n", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %+v\n",
			err)
	}
	return eg.Wait()

}
