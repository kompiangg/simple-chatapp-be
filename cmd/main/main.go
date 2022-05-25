package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/kompiangg/chatapp/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := websocket.Upgrader(w, r)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%+V\n", err)
	}

	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Meow")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	setupRoutes()

	go func ()  {
		if err := http.ListenAndServe(":8080", nil) ; err != nil {
			log.Fatal("ERROR error while starting the server", err)
		}	
	}()
	log.Println("INFO starting the server on port 8080")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("INFO shutting down the server")
}