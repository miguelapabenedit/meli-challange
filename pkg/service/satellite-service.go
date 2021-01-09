package service

import (
	"fmt"
	"log"

	"github.com/miguelapabenedit/meli-challange/internal/algorithm"
	"github.com/miguelapabenedit/meli-challange/pkg/entity"
)

type service struct{}

var satellites []entity.Satellite = []entity.Satellite{
	{
		Name:            "Kenobi",
		PositionX:       -500,
		PositionY:       -200,
		LastTransmition: entity.Transmition{Distance: 0, Message: nil},
	},
	{
		Name:            "Skywalker",
		PositionX:       100,
		PositionY:       -100,
		LastTransmition: entity.Transmition{Distance: 0, Message: nil},
	},
	{
		Name:            "Sato",
		PositionX:       500,
		PositionY:       100,
		LastTransmition: entity.Transmition{Distance: 0, Message: nil},
	}}

func NewSatelliteService() service {
	fmt.Println(satellites)
	return service{}
}

func (*service) GetLocation(distances ...float32) (x, y float32) {
	posX, posY, err := algorithm.Trilateration(-500, -200, float64(distances[0]), 100, -100, float64(distances[1]), 500, 100, float64(distances[2]))

	if err != nil {
		log.Println(err.Error())
	}

	x = float32(posX)
	y = float32(posY)
	return
}

func (*service) GetMessage(messages ...[]string) (msg string) {
	return "sda"
}
