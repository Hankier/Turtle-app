package main

import (
	"net"
	"bufio"
	"fmt"
	"os"
	"github.com/Nik-U/pbc"
	"time"
	"strconv"
	"io"
	"io/ioutil"
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

var debug = false

func (c *Crypto) init(){
	f, _ := os.Open("ss512.param")
	defer f.Close()
	params, _ := pbc.NewParams(f)
	c.pairing = pbc.NewPairing(params)
	if debug {
		fmt.Println("-----PAIRING-----")
		fmt.Print(params.String())
		fmt.Println("-----END_PAIRING-----")
		fmt.Println()
	}
	c.gen = c.pairing.NewUncheckedElement(pbc.G1)
	c.otherPublicKey = c.pairing.NewUncheckedElement(pbc.G1)
	c.otherSessionGenPublicKey = c.pairing.NewUncheckedElement(pbc.G1)
	c.otherSessionMultiplier = c.pairing.NewUncheckedElement(pbc.Zr)
	c.otherSessionPublicKey = c.pairing.NewUncheckedElement(pbc.G1)
	c.secretKey = c.pairing.NewUncheckedElement(pbc.Zr)
	c.publicKey = c.pairing.NewUncheckedElement(pbc.G1)
	c.sessionGenSecretKey = c.pairing.NewUncheckedElement(pbc.Zr)
	c.sessionGenPublicKey = c.pairing.NewUncheckedElement(pbc.G1)
	c.sessionMultiplier = c.pairing.NewUncheckedElement(pbc.Zr)
	c.sessionGen = c.pairing.NewUncheckedElement(pbc.G2)
	c.sessionPublicKey = c.pairing.NewUncheckedElement(pbc.G2)
}

func (c *Crypto) initGenerator(){
	c.gen.Rand()
	if debug{
		fmt.Println("-----GEN-----")
		fmt.Println("gen: ", c.gen.Bytes())
		fmt.Println("-----END_GEN-----")
		fmt.Println()
	}
}

func (c *Crypto) setGenerator(bytes []byte){
	c.gen.SetBytes(bytes)
	if debug {
		fmt.Println("-----SET_GEN-----")
		fmt.Println("gen: ", c.gen.Bytes())
		fmt.Println("-----END_SET_GEN-----")
		fmt.Println()
	}
}

func (c *Crypto) setOtherPublicKey(bytes []byte){
	c.otherPublicKey.SetBytes(bytes)
	if debug {
		fmt.Println("-----OTHER_KEYS-----")
		fmt.Println("public: " + c.otherPublicKey.String())
		fmt.Println("-----END_OTHERS_KEYS-----")
		fmt.Println()
	}
}

func (c *Crypto) setOtherSessionGenPublicKey(bytes []byte){
	c.otherSessionGenPublicKey.SetBytes(bytes)
	if debug {
		fmt.Println("-----OTHER_SESSION_GEN_KEYS-----")
		fmt.Println("public: " + c.otherSessionGenPublicKey.String())
		fmt.Println("-----END_OTHER_SESSION_GEN_KEYS-----")
		fmt.Println()
	}
}

func (c *Crypto) setOtherSessionMultiplier(bytes []byte){
	c.otherSessionMultiplier.SetBytes(bytes)
	if debug{
		fmt.Println("-----OTHER_SESSION_MULTIPLIER-----")
		fmt.Println("multiplier: " + c.otherSessionMultiplier.String())
		fmt.Println("-----END_OTHER_SESSION_MULTIPLIER-----")
		fmt.Println()
	}
}

func (c *Crypto) setOtherSessionPublicKey(bytes []byte){
	c.otherSessionPublicKey.SetBytes(bytes)
	if debug{
		fmt.Println("-----OTHER_SESSION_PUBLIC_KEY-----")
		fmt.Print("session public key: ")
		fmt.Println(c.otherSessionPublicKey)
		fmt.Println("-----END_OTHER_SESSION_PUBLIC_KEY-----")
		fmt.Println()
	}
}

func (c *Crypto) generateKeys(){
	c.secretKey.Rand()
	c.publicKey.PowZn(c.gen, c.secretKey)
	if debug{
		fmt.Println("-----KEYS-----")
		fmt.Println("public: " + c.publicKey.String())
		fmt.Println("secret: " + c.secretKey.String())
		fmt.Println("-----END_KEYS-----")
		fmt.Println()
	}
}

func (c *Crypto) generateSessionGenKeys(){
	c.sessionGenSecretKey.Rand()
	c.sessionGenPublicKey.PowZn(c.gen, c.sessionGenSecretKey)
	if debug{
		fmt.Println("-----SESSION_GEN_KEYS-----")
		fmt.Println("public: " + c.sessionGenPublicKey.String())
		fmt.Println("secret: " + c.sessionGenSecretKey.String())
		fmt.Println("-----END_SESSION_GEN_KEYS-----")
		fmt.Println()
	}
}

func (c *Crypto) generateSessionMultiplier(){
	c.sessionMultiplier.Rand()
	if debug{
		fmt.Println("-----SESSION_MULTIPLIER-----")
		fmt.Println("multiplier: " + c.sessionMultiplier.String())
		fmt.Println("-----END_SESSION_MULTIPLIER-----")
		fmt.Println()
	}
}

func (c *Crypto) generateSessionGen(invert bool){
	bytes1 := c.sessionGenPublicKey.Bytes()
	bytes2 := c.otherSessionGenPublicKey.Bytes()
	var hash []byte
	if invert {
		hash = append(bytes1, bytes2...)
	} else {
		hash = append(bytes2, bytes1...)
	}
	c.sessionGen.SetFromHash(hash)
	if debug{
		fmt.Println("-----SESSION_GEN-----")
		fmt.Println("gen: ", c.sessionGen.Bytes())
		fmt.Println("-----END_SESSION_GEN-----")
		fmt.Println()
	}
}

func (c *Crypto) generateSessionPublicKey(){
	exp := c.pairing.NewUncheckedElement(pbc.Zr).Set0().ThenAdd(c.secretKey).ThenMulZn(c.otherSessionMultiplier).ThenAdd(c.sessionGenSecretKey)
	c.sessionPublicKey.PowZn(c.sessionGen, exp)
	if debug{
		fmt.Println("-----SESSION_PUBLIC_KEY-----")
		fmt.Print("exp: ")
		fmt.Println(exp)
		fmt.Print("session public key: ")
		fmt.Println(c.sessionPublicKey)
		fmt.Println("-----END_SESSION_PUBLIC_KEY-----")
		fmt.Println()
	}
}

func (c *Crypto) checkKeys() bool{
	pairing1 := c.pairing.NewUncheckedElement(pbc.GT)
	pairing1.Pair(c.otherSessionPublicKey, c.gen)

	temp1 := c.pairing.NewUncheckedElement(pbc.G1).Set(c.otherPublicKey).ThenPowZn(c.sessionMultiplier).ThenMul(c.otherSessionGenPublicKey)
	pairing2 := c.pairing.NewUncheckedElement(pbc.GT)
	pairing2.Pair(c.sessionGen, temp1)

	if debug{
		fmt.Print("pairing1: ")
		fmt.Println(pairing1)
		fmt.Print("pairing2: ")
		fmt.Println(pairing2)
	}

	return pairing1.Equals(pairing2)
}

func (c *Crypto) generateCommonKey(){
	exp := c.pairing.NewUncheckedElement(pbc.Zr)
	exp.Set0().ThenAdd(c.secretKey).ThenMulZn(c.otherSessionMultiplier).ThenAdd(c.sessionGenSecretKey)
	c.sessionCommonKey = c.pairing.NewUncheckedElement(pbc.G2)
	c.sessionCommonKey.PowZn(c.otherSessionPublicKey, exp)
	if debug{
		fmt.Println("-----SESSION_COMMON_KEY-----")
		fmt.Print("exp: ")
		fmt.Println(exp)
		fmt.Print("session common key: ")
		fmt.Println(c.sessionCommonKey)
		fmt.Println("-----END_SESSION_COMMON_KEY-----")
	}
}

func listen(crypto *Crypto, conn net.Conn){
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for{
		cmd, data, err := readCmd(reader)
		if err != nil { fmt.Println(err); break}

		//fmt.Println("Cmd recv:" + cmd)

		switch cmd {
		case "part1":
			crypto.setOtherPublicKey(data[0])
			crypto.setOtherSessionGenPublicKey(data[1])
			fmt.Println("-> A, X\n")

			crypto.generateSessionGenKeys()
			crypto.generateSessionMultiplier()
			fmt.Println("gen y, Y=g^y, cb\n")

			writeCmd(writer, "part2",
				crypto.publicKey.Bytes(),
				crypto.sessionGenPublicKey.Bytes(),
				crypto.sessionMultiplier.Bytes())
			fmt.Println("B, Y, cb ->\n")
			break
		case "part2":
			crypto.setOtherPublicKey(data[0])
			crypto.setOtherSessionGenPublicKey(data[1])
			crypto.setOtherSessionMultiplier(data[2])
			fmt.Println("-> B, Y, cb\n")

			crypto.generateSessionGen(false)
			crypto.generateSessionPublicKey()
			crypto.generateSessionMultiplier()
			fmt.Println("gen ĝ=𝓗(X,Y), Sa=ĝ^(x+a*cb), ca\n")

			writeCmd(writer, "part3",
				crypto.sessionMultiplier.Bytes(),
				crypto.sessionPublicKey.Bytes())
			fmt.Println("Sa, ca ->\n")
			break
		case "part3":
			crypto.setOtherSessionMultiplier(data[0])
			crypto.setOtherSessionPublicKey(data[1])
			fmt.Println("-> Sa, ca\n")

			crypto.generateSessionGen(true)
			crypto.generateSessionPublicKey()
			fmt.Println("gen ĝ=𝓗(X,Y), Sb=ĝ^(y+b*ca)\n")

			fmt.Print("ê(Sa,g) = ê(ĝ, X*A^cb) ")
			if crypto.checkKeys(){
				fmt.Println("passed\n")
			} else {
				fmt.Println("failed\n")
			}

			crypto.generateCommonKey()
			fmt.Println("gen K=Sa^(y+b*ca)\n")

			writeCmd(writer, "part4", crypto.sessionPublicKey.Bytes())
			fmt.Println("Sb ->\n")

			fmt.Println("K=", crypto.sessionCommonKey)
			break;
		case "part4":
			crypto.setOtherSessionPublicKey(data[0])
			fmt.Println("-> Sb\n")

			fmt.Print("ê(Sb,g) = ê(ĝ, Y*B^ca) ")
			if crypto.checkKeys(){
				fmt.Println("passed\n")
			} else {
				fmt.Println("failed\n")
			}

			crypto.generateCommonKey()
			fmt.Println("gen K=Sb^(x+a*cb)\n")

			fmt.Println("K=", crypto.sessionCommonKey)

			break
		}
	}
	conn.Close()
	os.Exit(0)
}

func writeCmd(writer *bufio.Writer, cmd string, elements... []byte){
	writer.Write([]byte(cmd + "\n"))
	writer.Write([]byte(strconv.Itoa(len(elements)) + "\n"))
	for _, element := range elements{
		writer.Write([]byte(strconv.Itoa(len(element)) + "\n"))
		writer.Write(element)
	}
	writer.Flush()
}

func readCmd(reader *bufio.Reader) (string, [][]byte, error){
	cmdB, err := reader.ReadBytes('\n')
	if err != nil{return "", nil, err}
	cmd := string(cmdB[0:len(cmdB) -1])
	//fmt.Println("Got cmd:", cmd)

	sizeB, err := reader.ReadBytes('\n')
	if err != nil{return "", nil, err}
	size, _ := strconv.Atoi(string(sizeB[0:len(sizeB) - 1]))
	//fmt.Println("Size:", size)

	data := make([][]byte, size)

	for i := 0; i < size; i++{
		elSizeB, err := reader.ReadBytes('\n')
		if err != nil{return "", nil, err}
		elSize, _ := strconv.Atoi(string(elSizeB[0:len(elSizeB) - 1]))
		//fmt.Println("El size:", elSize)

		elData := make([]byte, elSize)
		_, err = io.ReadFull(reader, elData)
		if err != nil{return "", nil, err}
		//fmt.Println("El data:", elData)

		data[i] = elData
	}

	return cmd, data, nil
}

func main(){
	crypto := new(Crypto)
	crypto.init()
	loadGen(crypto)
	crypto.generateKeys()
	fmt.Println("public: " + crypto.publicKey.String())
	fmt.Println("secret: " + crypto.secretKey.String())
	fmt.Println()

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	for err != nil{
		fmt.Println("Error connecting " + err.Error())
		fmt.Println("Waiting 5s")
		time.Sleep(5 * time.Second)
	}

	go listen(crypto, conn)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)

	for{
		//fmt.Print("Cmd:")
		text, _ := reader.ReadString('\n')

		if text == "start\n"{
			fmt.Println("gen x, X=g^x\n")
			crypto.generateSessionGenKeys()

			fmt.Println("A, X ->\n")
			writeCmd(writer, "part1", crypto.publicKey.Bytes(), crypto.sessionGenPublicKey.Bytes())
		}
	}
}

func loadGen(c *Crypto){
	var bytes = make([]byte, 128)
	bytes, _ = ioutil.ReadFile("gen.bin")
	c.setGenerator(bytes)
}

func saveGen(c *Crypto){
	ioutil.WriteFile("gen.bin", c.gen.Bytes(), 0664);
}