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

// BatchPostSatellitesMessage godoc
// @Summary Process a Batch of incoming satellites transmitions update
// @ID post-batch
// @Description returns the position and message of a given transmitter
// @Tags branch
// @Accept json
// @Produce  json
// @Param branch body []entity.Satellite true "Satellite"
// @Success 200 {object} entity.TransmitionResponse
// @Failure 400,404 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server error"
// @Router /topsecret/ [post]
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

// PostSatelliteMessage godoc
// @Summary Updates a Satellite transmition data
// @ID post-satellite-message
// @Description updates a Satellite transmition data base on the
// @Tags branch
// @Accept json
// @Produce  json
// @Param satellite path string true "Satellite Name"
// @Param branch body []entity.Satellite true "Satellite"
// @Success 200
// @Failure 400,404 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server error"
// @Router /topsecret_split/{satellite} [post]
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

// GetTransmition godoc
// @Summary Try to retrieve the position and message of the transmition
// @ID get-transmition
// @Description Retrieves the information of the satellites base on tu current data
// @Tags branch
// @Accept json
// @Produce  json
// @Success 200 {object} entity.TransmitionResponse
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server error"
// @Router /topsecret_split/ [get]
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

// ExecuteOrder godoc
// @Summary Executes the holy command to erase existing transmition
// @ID execute-order
// @Description deletes all the distances and messages from the satellites, for a fresh start.
// @Tags branch
// @Accept json
// @Produce  json
// @Success 200
// @Router /topsecret/order/66 [delete]
func (controller) ExecuteOrder(w http.ResponseWriter, r *http.Request) {
	serv.DeleteTransmitionData()
	w.WriteHeader(http.StatusOK)
	errJSON, _ := json.Marshal("Yes my Lord")
	w.Header().Set("Content-Type", "application/json")
	w.Write(errJSON)
	return
}
