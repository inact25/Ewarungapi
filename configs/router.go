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

	routerHost := utils.GetCustomConf("RouterHost", "yourRouterHost")
	routerPort := utils.GetCustomConf("RouterPort", "8080")

	log.Printf("Server is now listening at %v.....\n", routerPort)
	err := http.ListenAndServe(routerHost+": "+routerPort, router)
	if err != nil {
		log.Fatal(err)
	}
}
