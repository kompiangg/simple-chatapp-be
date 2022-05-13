package main

import (
	"net/http"
)

// func initGlobalRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	return r
// }

func main() {
	// r := initGlobalRouter()
	// controller.InitRoutes(r)
	http.ListenAndServe(":8080", nil)
}