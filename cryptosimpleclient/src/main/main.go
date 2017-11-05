package main

import (
	"net"
	"bufio"
	"fmt"
	"os"
	"github.com/Nik-U/pbc"
	_ "time"
)

type Crypto struct{
	pairing                  *pbc.Pairing
	gen                      *pbc.Element //g
	sessionGen               *pbc.Element //g dash
	secretKey                *pbc.Element //a, b
	publicKey                *pbc.Element //A / B
	otherPublicKey           *pbc.Element //A / B
	sessionGenSecretKey      *pbc.Element //x, y
	sessionGenPublicKey      *pbc.Element //X / Y
	otherSessionGenPublicKey *pbc.Element //X / Y
	sessionMultiplier        *pbc.Element //ca / cb
	otherSessionMultiplier   *pbc.Element //ca / cb
	sessionPublicKey         *pbc.Element //Sa / Sb
	otherSessionPublicKey    *pbc.Element //Sa / Sb
	sessionCommonKey         *pbc.Element //K
}

func (c *Crypto) initPairing(){
	params := pbc.GenerateA(160, 512)
	c.pairing = pbc.NewPairing(params)
	fmt.Println("-----PAIRING-----")
	fmt.Print(params.String())
	fmt.Println("-----END_PAIRING-----")
	fmt.Println()
}

func (c *Crypto) initGenerator(){
	c.gen = c.pairing.NewG1().Rand()
}

func (c *Crypto) setPairing(pairing *pbc.Pairing){
	c.pairing = pairing
}

func (c *Crypto) setGenerator(gen *pbc.Element){
	c.gen = gen
}

func (c *Crypto) generateKeys(){
	c.secretKey = c.pairing.NewZr()
	c.secretKey.Rand()
	c.publicKey = c.pairing.NewG1().PowZn(c.gen, c.secretKey)
	fmt.Println("-----KEYS-----")
	fmt.Println("public: " + c.publicKey.String())
	fmt.Println("secret: " + c.secretKey.String())
	fmt.Println("-----END_KEYS-----")
	fmt.Println()
}

func (c *Crypto) generateSessionGenKeys(){
	c.sessionGenSecretKey = c.pairing.NewZr()
	c.sessionGenSecretKey.Rand()
	c.sessionGenPublicKey = c.pairing.NewG1().PowZn(c.gen, c.sessionGenSecretKey)
	fmt.Println("-----SESSION_GEN_KEYS-----")
	fmt.Println("public: " + c.sessionGenPublicKey.String())
	fmt.Println("secret: " + c.sessionGenSecretKey.String())
	fmt.Println("-----END_SESSION_GEN_KEYS-----")
	fmt.Println()
}

func (c *Crypto) generateSessionMultiplier(){
	c.sessionMultiplier = c.pairing.NewZr()
	c.sessionMultiplier.Rand()
	fmt.Println("-----SESSION_MULTIPLIER-----")
	fmt.Println("multiplier: " + c.sessionMultiplier.String())
	fmt.Println("-----END_SESSION_MULTIPLIER-----")
	fmt.Println()
}

func (c *Crypto) generateSessionGen(invert bool){
	c.sessionGen = c.pairing.NewG2()
	bytes1 := c.sessionGenPublicKey.Bytes()
	bytes2 := c.otherSessionGenPublicKey.Bytes()
	var hash []byte
	if invert {
		hash = append(bytes1, bytes2...)
	} else {
		hash = append(bytes2, bytes1...)
	}
	c.sessionGen.SetFromHash(hash)
	fmt.Println("-----SESSION_GEN-----")
	fmt.Print("public key 1: ")
	fmt.Println(bytes1)
	fmt.Print("public key 2: ")
	fmt.Println(bytes2)
	fmt.Print("hash: ")
	fmt.Println(hash)
	fmt.Println("-----END_SESSION_GEN-----")
	fmt.Println()
}

func (c *Crypto) generateSessionPublicKey(){
	c.sessionPublicKey = c.pairing.NewG2()
	exp := c.pairing.NewZr().Set0().ThenAdd(c.secretKey).ThenMulZn(c.otherSessionMultiplier).ThenAdd(c.sessionGenSecretKey)
	c.sessionPublicKey.PowZn(c.sessionGen, exp)
	fmt.Println("-----SESSION_PUBLIC_KEY-----")
	fmt.Print("exp: ")
	fmt.Println(exp)
	fmt.Print("session public key: ")
	fmt.Println(c.sessionPublicKey)
	fmt.Println("-----END_SESSION_PUBLIC_KEY-----")
	fmt.Println()
}

func (c *Crypto) checkKeys(){
	pairing1 := c.pairing.NewGT()
	pairing1.Pair(c.otherSessionPublicKey, c.gen)

	temp1 := c.pairing.NewG1().Set(c.otherPublicKey).ThenPowZn(c.sessionMultiplier).ThenMul(c.otherSessionGenPublicKey)
	pairing2 := c.pairing.NewGT()
	pairing2.Pair(c.sessionGen, temp1)

	fmt.Print("pairing1: ")
	fmt.Println(pairing1)
	fmt.Print("pairing2: ")
	fmt.Println(pairing2)

	if pairing1.Equals(pairing2){
		fmt.Println("Pairing check PASSED!!!")
	} else {
		fmt.Println("Pairing check FAILED!!!")
	}
	fmt.Println()
}

func (c *Crypto) generateCommonKey(){
	exp := c.pairing.NewZr()
	exp.Set0().ThenAdd(c.secretKey).ThenMulZn(c.otherSessionMultiplier).ThenAdd(c.sessionGenSecretKey)
	c.sessionCommonKey = c.pairing.NewG2()
	c.sessionCommonKey.PowZn(c.otherSessionPublicKey, exp)
	fmt.Println("-----SESSION_COMMON_KEY-----")
	fmt.Print("exp: ")
	fmt.Println(exp)
	fmt.Print("session common key: ")
	fmt.Println(c.sessionCommonKey)
	fmt.Println("-----END_SESSION_COMMON_KEY-----")
}

func listen(conn net.Conn){
	reader := bufio.NewReader(conn)
	for{
		msg, err := reader.ReadString('\n')
			if err != nil{
				fmt.Println(err)
				break
			}
		fmt.Println("Received:" + msg)
	}
	conn.Close()
	os.Exit(0)
}

func main(){
	crypto := new(Crypto)
	crypto.initPairing()
	crypto.initGenerator()
	crypto.generateKeys()


	crypto2 := new(Crypto)
	crypto2.setPairing(crypto.pairing)
	crypto2.setGenerator(crypto.gen)
	crypto2.generateKeys()

	//debug
	crypto.generateSessionGenKeys()
	crypto2.otherPublicKey = crypto.publicKey
	crypto2.otherSessionGenPublicKey = crypto.sessionGenPublicKey

	crypto2.generateSessionGenKeys()
	crypto2.generateSessionMultiplier()
	crypto.otherPublicKey = crypto2.publicKey
	crypto.otherSessionGenPublicKey = crypto2.sessionGenPublicKey
	crypto.otherSessionMultiplier = crypto2.sessionMultiplier

	crypto.generateSessionGen(false)
	crypto.generateSessionPublicKey()
	crypto.generateSessionMultiplier()
	crypto2.otherSessionMultiplier = crypto.sessionMultiplier
	crypto2.otherSessionPublicKey = crypto.sessionPublicKey

	crypto2.generateSessionGen(true)
	crypto2.generateSessionPublicKey()
	crypto2.checkKeys()

	crypto.otherSessionPublicKey = crypto2.sessionPublicKey
	crypto.checkKeys()

	crypto.generateCommonKey()
	crypto2.generateCommonKey()

	/*
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	for err != nil{
		fmt.Println("Error connecting " + err.Error())
		fmt.Println("Waiting 5s")
		time.Sleep(5 * time.Second)
	}

	go listen(conn)

	reader := bufio.NewReader(os.Stdin)

	for{
		fmt.Print("Text to send:")
		text, _ := reader.ReadString('\n')



		fmt.Fprintf(conn, text)
	}
	*/
}
