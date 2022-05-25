package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kompiangg/chatapp/internal/global"
	serviceAPI "github.com/kompiangg/chatapp/internal/websocket/service/api"
	"github.com/kompiangg/chatapp/pkg/websocket"
)

// log.Println("Websocket EndPoint Hitted")
// 	ws, err := websocket.Upgrader(w, r)
// 	if err != nil {
// 		log.Println(err)
// 		fmt.Fprintf(w, "%+V\n", err)
// 	}

type webSocketController struct {
	r *mux.Router
	s serviceAPI.WebSocketServiceAPI
}

func (w *webSocketController) createClient(rw http.ResponseWriter, r *http.Request) {
	log.Println("Websocket EndPoint Hitted")
	ws, err := websocket.Upgrader(rw, r)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(rw, "%+V\n", err)
	}

	w.s.CreateClient(ws)
}

func (w *webSocketController) InitializeController() {
	subRouter := w.r.PathPrefix(global.API_PATH_WEBSOCKET).Subrouter()
	subRouter.HandleFunc("", w.createClient)
}

func ProvideWebSocketController(r *mux.Router, s serviceAPI.WebSocketServiceAPI) *webSocketController {
	return &webSocketController{
		r: r,
		s: s,
	}
}