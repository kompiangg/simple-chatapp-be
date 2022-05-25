package main

import (
	"github.com/gorilla/mux"
	"github.com/kompiangg/chatapp/pkg/controller"
	"github.com/kompiangg/chatapp/pkg/server"
	"github.com/kompiangg/chatapp/pkg/websocket"
)

func providePool() *websocket.Pool {
	pool := websocket.NewPool()
	go pool.Start()

	return pool
}

func main() {
	r := mux.NewRouter()
	pool := providePool()

	controller.InitRoutes(pool, r)
	server.ProvideServer(r)
}