package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kompiangg/chatapp/internal/global"
	"github.com/kompiangg/chatapp/internal/ping/service"
)

type pingController struct {
	r *mux.Router
	p service.PingServiceAPI
}

func (p pingController) Ping(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode("Message: meow")
}

func (p *pingController) InitializeController() {
	subRouter := p.r.PathPrefix(global.API_PATH_PING).Subrouter()
	subRouter.HandleFunc("", p.Ping).Methods(http.MethodGet, http.MethodOptions)
}

func ProvidePingController(r *mux.Router, p service.PingServiceAPI) *pingController {
	return &pingController{
		r: r,
		p: p,
	}
}