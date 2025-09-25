package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"boot.rohi.tv/internal/request"
	"boot.rohi.tv/internal/response"
	"boot.rohi.tv/internal/server"
)

const port = 42069

func main() {
server, err := server.Serve(port, func(w io.Writer, req *request.Request) *server.HandlerError {
			if req.RequestLine.RequestTarget == "/yourproblem" {
				return &server.HandlerError{
					StatusCode: response.StatusBadRequest,
					Message:    "Your problem is not my proble\n",
				}
			} else if req.RequestLine.RequestTarget == "/myproble" {
				return &server.HandlerError{
					StatusCode: response.StatusInternalServerError,
					Message:    "Woopsie, my bad\n",
				}
			} else {
				w.Write([]byte("All good, frfr\n"))
			}
			return nil
		})

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	defer server.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}
