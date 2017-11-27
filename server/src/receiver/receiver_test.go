package receiver

import "testing"

func TestReceiver_twoBytesToInt(t *testing.T) {
	size := make([]byte, 2)
	size[0] = 15
	size[1] = 2

	num := twoBytesToInt(size)

	if num != 527 {
		t.Error("Expected 527, got ", num)
	}
}
