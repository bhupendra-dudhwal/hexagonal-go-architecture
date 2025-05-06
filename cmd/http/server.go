package main

import (
	"context"
	"fmt"
	"log"
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
		SetServer().
		Build()

	fmt.Println(time.Now().In(app.Timezone))

	go func() {
		fmt.Println("server running on :8080")
		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("error: %w", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	fmt.Println("cancel signal received")

	// Siganl received not wait to finish ongoing task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		fmt.Printf("error: while trying to shoutdown server")
	}
}
