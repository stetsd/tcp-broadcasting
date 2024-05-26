package main

import (
	"context"
	"log"
	"log/slog"

	"tcp-broadcasting/internal/container"
)

func main() {
	slog.Info("Hello mazafaka")

	cnt := container.NewContainer()

	err := cnt.Run(context.Background())
	if err != nil {
		log.Fatalf("container running error: %s", err)
	}
}
