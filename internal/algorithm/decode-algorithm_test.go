package algorithm

// func TestProcessMessages_WithRequiredLenghtData_ReturnMessage(t *testing.T) {
// 	type test struct {
// 		dataMsg1 []string
// 		dataMsg2 []string
// 		dataMsg3 []string
// 		answer   string
// 		err      error
// 	}

// 	testcases := []test{
// 		{
// 			dataMsg1: []string{"0", "1", "", ""},
// 			dataMsg2: []string{"", "", "2", ""},
// 			dataMsg3: []string{"", "1", "", "3"},
// 			answer:   "0 1 2 3 ",
// 			err:      nil,
// 		},
// 		{
// 			dataMsg1: []string{"", "", "", ""},
// 			dataMsg2: []string{"", "1", "2", "3"},
// 			dataMsg3: []string{"0", "1", "", "3"},
// 			answer:   "0 1 2 3 ",
// 			err:      nil,
// 		},
// 	}

// 	for _, val := range testcases {

// 		got, err := ProcessMessages(val.dataMsg1, val.dataMsg2, val.dataMsg3)

// 		if val.err != nil {
// 			assert.Error(t, err, val.err)
// 		}
// 		assert.Equal(t, val.answer, got)
// 	}
// }

// func TestProcessMessages_InvalidAmountOfArrays_ReturnError(t *testing.T) {
// 	msgData := []string{"", "", ""}
// 	got, err := ProcessMessages(msgData)

// 	assert.Error(t, err, "Expected array lenght to be > 1 is required")
// 	assert.Equal(t, "", got)
// }

// func TestProcessMessages_EmtyArrays_ReturnError(t *testing.T) {
// 	msgData1 := []string{"", "", ""}
// 	msgData2 := []string{"", "", ""}
// 	msgData3 := []string{"", "", ""}
// 	got, err := ProcessMessages(msgData1, msgData2, msgData3)

// 	assert.Error(t, err, "Unable to determine the message")
// 	assert.Equal(t, "", got)
// }
