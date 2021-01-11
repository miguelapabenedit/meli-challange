package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/miguelapabenedit/meli-challange/docs"
	_ "github.com/miguelapabenedit/meli-challange/docs"
	"github.com/miguelapabenedit/meli-challange/pkg/controller"
	"github.com/miguelapabenedit/meli-challange/pkg/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	satelliteController        = controller.NewSatelliteController(service.NewSatelliteService())
	port                string = os.Getenv("PORT")
	host                string = os.Getenv("HOST")
)

// @title Meli Challange Satellite API
// @version 1.0
// @description This api serves the Rebel Alliance by providing apis to recieve and decode position and messages of alliance transmitions
// @contact.name API Support
// @contact.email miguell.beneditt@gmail.com
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/topsecret/", satelliteController.BatchPostSatellitesMessage).Methods(http.MethodPost)
	r.HandleFunc("/topsecret_split/{satellite}", satelliteController.PostSatelliteMessage).Methods(http.MethodPost)
	r.HandleFunc("/topsecret_split/", satelliteController.GetTransmition).Methods(http.MethodGet)
	r.HandleFunc("/topsecret/order/66", satelliteController.ExecuteOrder).Methods(http.MethodDelete)

	docs.SwaggerInfo.Host = host + ":" + port
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Printf("Running up on host:%s port: %s", host, port)
	log.Fatalln(http.ListenAndServe(":"+port, r))
}
