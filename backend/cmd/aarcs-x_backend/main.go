package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Akshansh-29072005/AARCS-X/backend/internals/config"
)

func main(){
	//load config
	cfg := config.MustLoad()
	//database setup
	//router setup
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the AARCS-X backend"))
	})
	//server setup
	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	
	slog.Info("Server started", slog.String("Address",cfg.Addr))

	//created a channel for the server to run
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	go func ()  {
		err := server.ListenAndServe()
		if err !=nil{
			log.Fatal("Failed to start the server")
		}
	} ()

	<-done

	slog.Info("Shutting down the server!")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil{
		slog.Error("Failed to shutdown the server", slog.String("error", err.Error()))
	}

	//Above lines can also be written as below also!
	// if err := server.Shutdown(ctx); err != nil{
	// 	slog.Error("Failed to shutdown the server", slog.String("error", err.Error()))
	// }

	slog.Info("Server shutdown sucessfully!")
}