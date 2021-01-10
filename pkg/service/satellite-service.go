package service

import (
	"errors"
	"log"

	"github.com/miguelapabenedit/meli-challange/internal/algorithm"
	"github.com/miguelapabenedit/meli-challange/pkg/entity"
)

type service struct{}

func NewSatelliteService() Service {
	return &service{}
}

var (
	satelliteKenobi = entity.Satellite{
		Name:     "kenobi",
		Message:  nil,
		Distance: -1,
		Position: entity.Position{X: -500, Y: -200}}
	satelliteSkywalker = entity.Satellite{
		Name:     "skywalker",
		Message:  nil,
		Distance: -1,
		Position: entity.Position{X: 100, Y: -100}}
	satelliteSato = entity.Satellite{
		Name:     "sato",
		Message:  nil,
		Distance: -1,
		Position: entity.Position{X: 500, Y: 100}}
)

func (*service) GetLocation(distances ...float32) (x, y float32) {
	posX, posY, err := algorithm.Trilateration(
		float64(satelliteKenobi.Position.X),
		float64(satelliteKenobi.Position.Y),
		float64(distances[0]),
		float64(satelliteSkywalker.Position.X),
		float64(satelliteSkywalker.Position.Y),
		float64(distances[1]),
		float64(satelliteSato.Position.X),
		float64(satelliteSato.Position.Y),
		float64(distances[2]))

	if err != nil {
		log.Println(err.Error())
	}

	x = float32(posX)
	y = float32(posY)
	return
}

func (*service) GetMessage(messages ...[]string) (msg string) {
	msg, err := algorithm.ProcessMessages(messages)

	if err != nil {
		log.Println(err.Error())
	}

	return
}

func (*service) GetTransmitionUpdateBatch(satellites []entity.Satellite) (*entity.Position, string, error) {
	for _, v := range satellites {
		err := updateSatelliteTransmition(&v)

		if err != nil {
			log.Println(err.Error())
			return nil, "", err
		}
	}

	position, err := getLocation()

	if err != nil {
		log.Println(err.Error())
		return nil, "", err
	}

	message, err := getMessage(satelliteKenobi.Message, satelliteSkywalker.Message, satelliteSato.Message)

	if err != nil {
		log.Println(err.Error())
		return nil, "", err
	}

	return position, message, nil
}

func (*service) PutSatelliteTransmition(satellite *entity.Satellite) error {
	err := updateSatelliteTransmition(satellite)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (*service) GetTransmition() (*entity.Position, string, error) {
	position, err := getLocation()

	if err != nil {
		log.Println(err.Error())
		return nil, "", err
	}

	message, err := getMessage(satelliteKenobi.Message, satelliteSkywalker.Message, satelliteSato.Message)

	if err != nil {
		log.Println(err.Error())
		return nil, "", err
	}

	return position, message, nil
}

func getLocation() (*entity.Position, error) {
	posX, posY, err := algorithm.Trilateration(
		float64(satelliteKenobi.Position.X),
		float64(satelliteKenobi.Position.Y),
		float64(satelliteKenobi.Distance),
		float64(satelliteSkywalker.Position.X),
		float64(satelliteSkywalker.Position.Y),
		float64(satelliteSkywalker.Distance),
		float64(satelliteSato.Position.X),
		float64(satelliteSato.Position.Y),
		float64(satelliteSato.Distance))

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	position := entity.Position{
		X: float32(posX),
		Y: float32(posY),
	}

	return &position, nil
}

func getMessage(msgs ...[]string) (string, error) {
	msg, err := algorithm.ProcessMessages(msgs)

	if err != nil {
		return "", err
	}

	return msg, nil
}

func updateSatelliteTransmition(transmition *entity.Satellite) error {
	err := validateTransmitionData(transmition.Distance, transmition.Message)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	switch transmition.Name {
	case satelliteKenobi.Name:
		satelliteKenobi.Distance = transmition.Distance
		satelliteKenobi.Message = transmition.Message

	case satelliteSkywalker.Name:
		satelliteSkywalker.Distance = transmition.Distance
		satelliteSkywalker.Message = transmition.Message

	case satelliteSato.Name:
		satelliteSato.Distance = transmition.Distance
		satelliteSato.Message = transmition.Message
	default:
		log.Printf("The satellite name: %s, is not valid", transmition.Name)
	}

	return nil
}

func validateTransmitionData(distance float32, msg []string) error {
	errs := ""
	if distance < 0 {
		errs += "Distance cant be negative\n\r"
	}
	if msg == nil || len(msg) == 0 {
		errs += "Msg can't be empty\n\r"
	}

	if errs != "" {
		return errors.New(errs)
	}

	return nil
}

func (*service) DeleteTransmitionData() {
	satelliteKenobi.Distance = -1
	satelliteKenobi.Message = nil
	satelliteSato.Distance = -1
	satelliteSato.Message = nil
	satelliteSkywalker.Distance = -1
	satelliteSkywalker.Message = nil
}
