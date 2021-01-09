package entity

/*Satellite holds the name and position of the objet with the last Transmition detected
 */
type Satellite struct {
	Name            string      `json:'name'`
	PositionX       float32     `json:'positionX'`
	PositionY       float32     `json:'positionY'`
	LastTransmition Transmition `json:'lastTransmition'`
}

/*Transmition representes a communication with a distance and message from a transmitter object/ship
 */
type Transmition struct {
	Distance float32  `json:'distance'`
	Message  []string `json:'message'`
}
