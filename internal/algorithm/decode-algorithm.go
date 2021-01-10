package algorithm

import "errors"

/*ProcessMessages decode incomplete messages base on several collections of the same message
[Required]  array lenght> 1
*/
func ProcessMessages(msg [][]string) (string, error) {
	if len(msg) < 2 {
		return "", errors.New("Expected array lenght to be > 1 is required")
	}

	numberOfArrays := len(msg)
	maxLenght := getMaxArrayLengh(msg)
	revertMessage := make([]string, maxLenght)
	var currentArrayPosition int

	for currentArrayPosition < numberOfArrays {
		fixedIndex := 0
		currentMsg := msg[currentArrayPosition]

		for i := len(currentMsg) - 1; i >= 0; i-- {
			if currentMsg[i] != "" {
				revertMessage[fixedIndex] = currentMsg[i]
			}
			fixedIndex++
		}
		currentArrayPosition++
	}

	var decodeMessage string

	for i := len(revertMessage) - 1; i >= 0; i-- {
		if revertMessage[i] != "" {
			decodeMessage += revertMessage[i] + " "
		}
	}

	if decodeMessage == "" {
		return decodeMessage, errors.New("Unable to determine the message")
	}

	return decodeMessage, nil
}

/*getMaxArrayLenght given a collection of arrays it retuns the biggest lenght detected
 */
func getMaxArrayLengh(arrays [][]string) (maxLenght int) {
	for _, val := range arrays {
		if len(val) > maxLenght {
			maxLenght = len(val)
		}
	}
	return
}
