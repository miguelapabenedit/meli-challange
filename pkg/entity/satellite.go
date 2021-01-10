package entity

/*Satellite holds the name and position of the objet with the last Transmition detected
 */
type Satellite struct {
	Name     string   `json:'name'`
	Message  []string `json:'message'`
	Distance float32  `json:'distance'`
	Position Position `json:'position'`
}
