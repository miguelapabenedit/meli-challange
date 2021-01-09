package service

/*Service provide a variaty of functions to process message and locate transmiters
 */
type Service interface {
	GetLocation(distances ...float32) (x, y float32)
	GetMessage(messages ...[]string) (msg string)
}
