package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessMessages_WithRequiredLenghtData_ReturnMessage(t *testing.T) {
	type test struct {
		msgs   [][]string
		answer string
		err    error
	}

	testcases := []test{
		{
			msgs: [][]string{
				{"0", "1", "", ""},
				{"", "", "2", ""},
				{"", "", "", "3"},
			},
			answer: "0 1 2 3 ",
			err:    nil,
		},
		{
			msgs: [][]string{
				{"", "", "", ""},
				{"", "1", "2", "3"},
				{"0", "1", "", "3"},
			},
			answer: "0 1 2 3 ",
			err:    nil,
		},
	}

	for _, val := range testcases {

		got, err := ProcessMessages(val.msgs)

		if val.err != nil {
			assert.Error(t, err, val.err)
		}
		assert.Equal(t, val.answer, got)
	}
}

func TestProcessMessages_InvalidAmountOfArrays_ReturnError(t *testing.T) {
	msgData := [][]string{
		{"0", "1", "2", "3"},
	}
	got, err := ProcessMessages(msgData)

	assert.Error(t, err, "Expected array lenght to be > 1 is required")
	assert.Equal(t, "", got)
}

func TestProcessMessages_EmtyArrays_ReturnError(t *testing.T) {
	msgData := [][]string{
		{"", "", "", ""},
		{"", "", "", ""},
		{"", "", "", ""},
	}
	got, err := ProcessMessages(msgData)

	assert.Error(t, err, "Unable to determine the message")
	assert.Equal(t, "", got)
}
