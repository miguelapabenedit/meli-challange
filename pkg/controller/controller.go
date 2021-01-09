package controller

import "net/http"

type controller interface {
	BatchPostSatellitesMessage(m http.ResponseWriter, r *http.Response)
	PostSatelliteMessage(m http.ResponseWriter, r *http.Response)
}
