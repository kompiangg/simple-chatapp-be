package controller

import (
	"github.com/gorilla/mux"
	"github.com/kompiangg/chatapp/pkg/websocket"

	PingControllerPkg "github.com/kompiangg/chatapp/internal/ping/controller"
	PingApiPkg "github.com/kompiangg/chatapp/internal/ping/service"

	webSocketControllerPkg "github.com/kompiangg/chatapp/internal/websocket/controller"
	WebSocketApiImpl "github.com/kompiangg/chatapp/internal/websocket/service/impl"
)

func InitRoutes(pool *websocket.Pool, r *mux.Router) {
	routes := r.NewRoute().Subrouter()

	pingServiceImpl := PingApiPkg.NewPingService()
	pingControllerImpl := PingControllerPkg.ProvidePingController(routes, pingServiceImpl)
	pingControllerImpl.InitializeController()

	webSocketService := WebSocketApiImpl.ProvideWebSocketService(pool)
	webSocketController := webSocketControllerPkg.ProvideWebSocketController(routes, webSocketService)
	webSocketController.InitializeController()
}