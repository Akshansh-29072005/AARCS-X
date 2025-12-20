package main

import (
	"fmt"
	"log"
	"net/http"

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
	fmt.Printf("Server Started %s",cfg.HTTPServer.Addr)
	err := server.ListenAndServe()
	if err !=nil{
		log.Fatal("Failed to start the server")
	}
}