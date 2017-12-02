package sender

import "testing"

func TestSenderImpl_addSizeToMessage(t *testing.T){
	bytes := make([]byte, 5)
	bytes[0] = 123
	bytes[1] = 2
	bytes[2] = 32
	bytes[3] = 5
	bytes[4] = 3

	bytes = addSizeToMessage(bytes)

	if bytes[0] != 5 || bytes[1] != 0{
		t.Error("Expected 5 0, got ", bytes[0], " ", bytes[1])
	}
}

func TestSenderImpl_messagesToSingle(t *testing.T) {
	bytes := make([][]byte, 2)
	bytes[0] = make([]byte, 1)
	bytes[0][0] = 15
	bytes[1] = make([]byte, 1)
	bytes[1][0] = 14

	result := messagesToSingle(bytes)

	if result[0] != 15 || result[1] != 14{
		t.Error("Expected 15 14, got ", result[0], " ", result[1])
	}
}
