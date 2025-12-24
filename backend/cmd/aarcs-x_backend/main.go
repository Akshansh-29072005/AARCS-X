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
	"github.com/Akshansh-29072005/AARCS-X/backend/internals/http/handlers/student"
	"github.com/Akshansh-29072005/AARCS-X/backend/internals/storage/sqlite"
)

func main(){
	//load config
	cfg := config.MustLoad()
	//database setup

	storage, err := sqlite.New(cfg)
	if err != nil{
		log.Fatal(err)
	}

	slog.Info("storage intialized", slog.String("env", cfg.Env), slog.String("version","1.0.0"))

	//router setup
	//This route is used for making new students.
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	//This route is used for extracting the student's info using the Id of the student.
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
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

	err = server.Shutdown(ctx)

	if err != nil{
		slog.Error("Failed to shutdown the server", slog.String("error", err.Error()))
	}

	//Above lines can also be written as below also!
	// if err := server.Shutdown(ctx); err != nil{
	// 	slog.Error("Failed to shutdown the server", slog.String("error", err.Error()))
	// }

	slog.Info("Server shutdown sucessfully!")
}