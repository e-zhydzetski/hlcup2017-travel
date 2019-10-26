package xhttp

import (
	"context"
	"github.com/e-zhydzetski/hlcup2017-travel/pkg/x/xerror"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"golang.org/x/sync/errgroup"
)

const ErrGracefulShutdown = xerror.Error("graceful shutdown")

func StartServer(ctx context.Context, addr string, handler http.Handler) error {
	g, ctx := errgroup.WithContext(ctx)
	{
		server := &http.Server{
			Addr:    addr,
			Handler: chi.ServerBaseContext(ctx, handler),
		}
		g.Go(func() error {
			<-ctx.Done()
			return server.Shutdown(context.Background()) // wait forever for live connections, maybe add timeout
		})
		g.Go(server.ListenAndServe)
	}
	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case /*sig :=*/ <-c:
			return ErrGracefulShutdown
		}
	})
	return g.Wait()
}
