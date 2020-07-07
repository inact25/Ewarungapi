package configs

import (
	"github.com/gorilla/mux"
	"github.com/inact25/E-WarungApi/utils"
	"log"
	"net/http"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func RunServer(router *mux.Router) {
	//routerPort ambil dari env, 6969 default port
	port := utils.GetCustomConf("RouterPort", "6969")
	log.Printf("Server is now listening at %v.....\n", port)
	err := http.ListenAndServe("localhost: "+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
