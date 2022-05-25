package controller

import (
	"github.com/gorilla/mux"
	"github.com/kompiangg/chatapp/pkg/websocket"

	PingControllerPkg "github.com/kompiangg/chatapp/internal/ping/controller"

	webSocketControllerPkg "github.com/kompiangg/chatapp/internal/websocket/controller"
	WebSocketApiImpl "github.com/kompiangg/chatapp/internal/websocket/service/impl"
)

func InitRoutes(pool *websocket.Pool, r *mux.Router) {
	routes := r.NewRoute().Subrouter()

	pingControllerImpl := PingControllerPkg.ProvidePingController(routes)
	pingControllerImpl.InitializeController()

	webSocketService := WebSocketApiImpl.ProvideWebSocketService(pool)
	webSocketController := webSocketControllerPkg.ProvideWebSocketController(routes, webSocketService)
	webSocketController.InitializeController()
}