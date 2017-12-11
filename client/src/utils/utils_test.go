package utils

import "testing"

func TestUtils_IntTo2bytes(t *testing.T){
	num := 5
	bytes := IntToTwobytes(num)

	if bytes[0] != 5 || bytes[1] != 0{
		t.Error("Expected 5 0, got ", bytes[0], " ", bytes[1])
	}

	num = 1000
	bytes = IntToTwobytes(num)

	if bytes[0] != 232 || bytes[1] != 3{
		t.Error("Expected 232 3, got", bytes[0], bytes[1])
	}
}

func TestUtils_TwoBytesToInt(t *testing.T) {
	size := make([]byte, 2)
	size[0] = 15
	size[1] = 2

	num := TwoBytesToInt(size)

	if num != 527 {
		t.Error("Expected 527, got ", num)
	}
}
