package sender

import "testing"

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
