package controller

import "net/http"

type Controller interface {
	BatchPostSatellitesMessage(w http.ResponseWriter, r *http.Request)
	PostSatelliteMessage(w http.ResponseWriter, r *http.Request)
	GetTransmition(w http.ResponseWriter, r *http.Request)
	ExecuteOrder(w http.ResponseWriter, r *http.Request)
}
