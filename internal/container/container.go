package container

import (
	"context"

	"tcp-broadcasting/internal/server"
)

type Container struct {
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Run(ctx context.Context) error {
	// TODO решение временное пока не будет ручки создания комнаты
	srv := server.NewServer()

	err := srv.Run()
	if err != nil {
		return err
	}

	return nil
}
