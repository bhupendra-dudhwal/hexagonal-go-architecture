package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"project_structure/internal/builder"
	"syscall"
	"time"
)

func main() {

	app := builder.NewAppBuilder().
		SetConfig().
		SetLogger().
		SetTimezone().
		SetMiddleware().
		SetResponse().
		SetUserService().
		SetAuthService().
		SetHealthService().
		SetHandlers().
		SetServer().
		Build()

	go func() {
		app.Log.Info().Str("port", ":8081").Msg("starting server")
		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Log.Fatal().Err(err).Msg("server crashed")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	app.Log.Info().Msg("shutdown signal received")

	// Siganl received not wait to finish ongoing task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		app.Log.Error().Err(err).Msg("error during shutdown")
	} else {
		app.Log.Info().Msg("server shutdown gracefully")
	}
}
