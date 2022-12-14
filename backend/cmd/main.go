package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/adykaaa/online-notes/api/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	httpServer := http.NewServer(r, ":80")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Printf("Server run interrupted by signal %s", s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		fmt.Errorf("app - Run - httpServer.Shutdown: %w", err)
	}
}
