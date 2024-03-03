package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func shutdown(ctx context.Context) <-chan struct{} {
	errCh := make(chan struct{}, 1)

	ctx, stop := signal.NotifyContext(ctx,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		defer func() {
			stop()
			close(errCh)
		}()

		<-ctx.Done()
	}()

	return errCh
}
