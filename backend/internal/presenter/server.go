package presenter

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

type HTTPServer struct {
	port   int
	router chi.Router
}

func NewHTTPServer(port int) *HTTPServer {
	r := NewRouter()
	return &HTTPServer{
		port:   port,
		router: r,
	}
}

func (s *HTTPServer) Serve(ctx context.Context) {
	addr := fmt.Sprintf(":%d", s.port)
	go func() {
		slog.Info("Server is listening", slog.String("addr", addr))
		if err := http.ListenAndServe(addr, s.router); err != http.ErrServerClosed {
			slog.Error(err.Error())
			panic("http server is stopped unexpectedly")
		}

		slog.Info("Server is stopped")
	}()

	waitFinishSignal()
	slog.Info("Server is shutting down...")
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	if err := s.shutdown(ctx); err != nil {
		slog.Error(err.Error())
		panic("http server is stopped unexpectedly")
	}
}

// TODO: implement graceful shutdown
func (s *HTTPServer) shutdown(_ context.Context) error {
	return nil
}

func waitFinishSignal() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, os.Interrupt)
	<-sigChan
}
