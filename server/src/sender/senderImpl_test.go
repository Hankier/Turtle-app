package sender

import "testing"

func TestSenderImpl_intTo2bytes(t *testing.T){
	num := 5
	bytes := intTo2bytes(num)

	if bytes[0] != 5 || bytes[1] != 0{
		t.Error("Expected 5 0, got ", bytes[0], " ", bytes[1])
	}

	num = 1000
	bytes = intTo2bytes(num)

	if bytes[0] != 232 || bytes[1] != 3{
		t.Error("Expected 232 3, got", bytes[0], bytes[1])
	}
}

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
