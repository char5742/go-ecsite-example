package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {

	return &Server{
		srv: &http.Server{
			Handler: mux},
		l: l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	eg, ctx := errgroup.WithContext(ctx)
	defer stop()
	eg.Go(func() error {
		// サーバーを開始
		// http.ErrServerClosedはサーバーが正常にシャットダウンされたことを示すエラーなので、無視する
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			log.Printf("Error starting server: %v\n", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %+v\n",
			err)
	}
	// グレースフルシャットダウンのための待機
	return eg.Wait()

}
