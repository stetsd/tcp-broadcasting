package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/google/uuid"
)

type Server struct {
	id   uuid.UUID
	conn net.Listener
}

func NewServer() *Server {
	ID := uuid.New()

	slog.Info("room created", "id", ID)
	return &Server{id: ID}
}

func (s *Server) Run() error {
	conn, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		return fmt.Errorf("server initialization error: %w", err)
	}

	s.conn = conn

	return nil
}
