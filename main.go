package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"server/server"
	"syscall"
)

func main() {

	ctx := context.Background()
	serverDoneChan := make(chan os.Signal, 1)

	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":8085")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server started on port 8085")
	<-serverDoneChan

	srv.Shutdown(ctx)
	log.Println("Server stopped")
}
