package utils

func IntToTwobytes(len int)[]byte{
	size := make([]byte, 2)
	size[0] = (byte)(len % 256)
	size[1] = (byte)(len / 256)

	return size
}

func TwoBytesToInt(size []byte)int{
	num := 0

	num += (int)(size[0])
	num += (int)(size[1]) * 256

	return num
}
