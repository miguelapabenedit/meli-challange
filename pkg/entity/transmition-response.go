package entity

/*TransmitionResponse represents a message and a distance of a incoming ship communication
 */
type TransmitionResponse struct {
	Position Position `json:'position'`
	Message  string   `json:'message'`
}
