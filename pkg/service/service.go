package service

import "github.com/miguelapabenedit/meli-challange/pkg/entity"

/*Service provide a variaty of functions to process message and locate transmiters
 */
type Service interface {
	GetLocation(distances ...float32) (x, y float32)
	GetMessage(messages ...[]string) (msg string)
	GetTransmitionUpdateBatch(satellites []entity.Satellite) (*entity.Position, string, error)
	PutSatelliteTransmition(sattellite *entity.Satellite) error
	GetTransmition() (*entity.Position, string, error)
	DeleteTransmitionData()
}
