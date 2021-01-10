package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguelapabenedit/meli-challange/pkg/controller"
	"github.com/miguelapabenedit/meli-challange/pkg/service"
)

var (
	satelliteController = controller.NewSatelliteController(service.NewSatelliteService())
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/topsecret/", satelliteController.BatchPostSatellitesMessage).Methods(http.MethodPost)
	r.HandleFunc("/topsecret_split/{satellite}", satelliteController.PostSatelliteMessage).Methods(http.MethodPost)
	r.HandleFunc("/topsecret_split/", satelliteController.GetTransmition).Methods(http.MethodGet)
	r.HandleFunc("/topsecret/order/66", satelliteController.ExecuteOrder).Methods(http.MethodDelete)
	log.Println("Running up")
	log.Fatalln(http.ListenAndServe(":5000", r))
}
