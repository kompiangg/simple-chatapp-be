package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
)

func ProvideServer(r *mux.Router) {
	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatal("ERROR error while starting the server", err)
		}
	}()
	log.Println("INFO starting the server on port 8080")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("INFO shutting down the server")
}