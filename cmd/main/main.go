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

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("Websocket EndPoint Hitted")
	ws, err := websocket.Upgrader(w, r)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &websocket.Client{
		Conn: ws,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes(pool *websocket.Pool) {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Meow")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	pool := websocket.NewPool()
	go pool.Start()

	setupRoutes(pool)

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