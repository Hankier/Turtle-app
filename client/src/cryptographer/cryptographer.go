package cryptographer

type Cryptographer interface{
	encrypt([]byte)[]byte
	decrypt([]byte)[]byte
}
