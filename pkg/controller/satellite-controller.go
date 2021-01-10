package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/miguelapabenedit/meli-challange/pkg/entity"
	"github.com/miguelapabenedit/meli-challange/pkg/service"
)

type controller struct {
}

var serv service.Service

func NewSatelliteController(satelliteService service.Service) Controller {
	serv = satelliteService
	return &controller{}
}

func (controller) BatchPostSatellitesMessage(w http.ResponseWriter, r *http.Request) {
	var satellites []entity.Satellite
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bodyBytes, &satellites)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	position, msg, err := serv.GetTransmitionUpdateBatch(satellites)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	transmition := entity.TransmitionResponse{Position: *position, Message: msg}

	transmitionJSON, err := json.Marshal(transmition)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(transmitionJSON)
	return
}

func (controller) PostSatelliteMessage(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "/")
	satelliteName := urlPathSegment[len(urlPathSegment)-1]

	if satelliteName == "" {
		log.Println("The satellite name can't be empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newSatelliteTransmition entity.Satellite
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bodyBytes, &newSatelliteTransmition)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newSatelliteTransmition.Name = satelliteName
	err = serv.PutSatelliteTransmition(&newSatelliteTransmition)
	return
}

func (controller) GetTransmition(w http.ResponseWriter, r *http.Request) {
	position, msg, err := serv.GetTransmition()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errJSON, _ := json.Marshal("there is not enough information")
		w.Header().Set("Content-Type", "application/json")
		w.Write(errJSON)
		return
	}

	transmition := entity.TransmitionResponse{Position: *position, Message: msg}

	transmitionJSON, err := json.Marshal(transmition)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(transmitionJSON)
	return
}

func (controller) ExecuteOrder(w http.ResponseWriter, r *http.Request) {
	serv.DeleteTransmitionData()
	w.WriteHeader(http.StatusOK)
	errJSON, _ := json.Marshal("Yes my Lord")
	w.Header().Set("Content-Type", "application/json")
	w.Write(errJSON)
	return
}
