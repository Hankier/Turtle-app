package main

import (
	"fmt"
	"io/ioutil"
)

func test1(){
	crypto := new(Crypto)
	crypto.init()
	crypto.initGenerator()
	crypto.generateKeys()
	crypto.generateSessionGenKeys()

	crypto2 := new(Crypto)
	crypto2.init()
	crypto2.setGenerator(crypto.gen.Bytes())
	crypto2.generateKeys()
	crypto2.setOtherPublicKey(crypto.publicKey.Bytes())
	crypto2.setOtherSessionGenPublicKey(crypto.sessionGenPublicKey.Bytes())
	crypto2.generateSessionGenKeys()
	crypto2.generateSessionMultiplier()


	crypto.setOtherPublicKey(crypto2.publicKey.Bytes())
	crypto.setOtherSessionGenPublicKey(crypto2.sessionGenPublicKey.Bytes())
	crypto.setOtherSessionMultiplier(crypto2.sessionMultiplier.Bytes())
	crypto.generateSessionGen(false)
	crypto.generateSessionPublicKey()
	crypto.generateSessionMultiplier()


	crypto2.setOtherSessionMultiplier(crypto.sessionMultiplier.Bytes())
	crypto2.setOtherSessionPublicKey(crypto.sessionPublicKey.Bytes())
	crypto2.generateSessionGen(true)
	crypto2.generateSessionPublicKey()
	crypto2.checkKeys()

	crypto.setOtherSessionPublicKey(crypto2.sessionPublicKey.Bytes())
	crypto.checkKeys()

	crypto.generateCommonKey()
	crypto2.generateCommonKey()

	saveCryptoToFile(crypto, "crypto")
	saveCryptoToFile(crypto2, "crypto2")
}

func saveCryptoToFile(c *Crypto, name string){
	ioutil.WriteFile(name + "_gen", c.gen.Bytes(), 0664);
	ioutil.WriteFile(name + "_secretKey", c.secretKey.Bytes(), 0664);
	ioutil.WriteFile(name + "_publicKey", c.publicKey.Bytes(), 0664);
	ioutil.WriteFile(name + "_sessionGen", c.sessionGen.Bytes(), 0664);
	ioutil.WriteFile(name + "_sessionGenSecretKey", c.sessionGenSecretKey.Bytes(), 0664);
	ioutil.WriteFile(name + "_sessionGenPublicKey", c.sessionGenPublicKey.Bytes(), 0664);
	ioutil.WriteFile(name + "_sessionMultiplier", c.sessionMultiplier.Bytes(), 0664);
	ioutil.WriteFile(name + "_sessionPublicKey", c.sessionPublicKey.Bytes(), 0664);
}

func loadCryptoFromFile(c *Crypto, name string){
	var bytes = make([]byte, 1024)
	bytes, err := ioutil.ReadFile(name + "_gen")
	fmt.Println(err)
	c.gen.SetBytes(bytes[:])
	fmt.Println("-----GEN-----")
	fmt.Println("gen: ", c.gen.Bytes())
	fmt.Println("-----END_GEN-----")
	fmt.Println()

	bytes, _ = ioutil.ReadFile(name + "_secretKey")
	c.secretKey.SetBytes(bytes[:])
	bytes, _ = ioutil.ReadFile(name + "_publicKey")
	c.publicKey.SetBytes(bytes[:])
	fmt.Println("-----KEYS-----")
	fmt.Println("public: " + c.publicKey.String())
	fmt.Println("secret: " + c.secretKey.String())
	fmt.Println("-----END_KEYS-----")
	fmt.Println()

	bytes, _ = ioutil.ReadFile(name + "_sessionGen")
	c.sessionGen.SetBytes(bytes[:])
	fmt.Println("-----SESSION_GEN-----")
	fmt.Print("session gen: ")
	fmt.Println(c.sessionGen.Bytes())
	fmt.Println("-----END_SESSION_GEN-----")
	fmt.Println()

	bytes, _ = ioutil.ReadFile(name + "_sessionGenSecretKey")
	c.sessionGenSecretKey.SetBytes(bytes[:])
	bytes, _ = ioutil.ReadFile(name + "_sessionGenPublicKey")
	c.sessionGenPublicKey.SetBytes(bytes[:])
	fmt.Println("-----SESSION_GEN_KEYS-----")
	fmt.Println("public: " + c.sessionGenPublicKey.String())
	fmt.Println("secret: " + c.sessionGenSecretKey.String())
	fmt.Println("-----END_SESSION_GEN_KEYS-----")
	fmt.Println()

	bytes, _ = ioutil.ReadFile(name + "_sessionMultiplier")
	c.sessionMultiplier.SetBytes(bytes[:])
	fmt.Println("-----SESSION_MULTIPLIER-----")
	fmt.Println("multiplier: " + c.sessionMultiplier.String())
	fmt.Println("-----END_SESSION_MULTIPLIER-----")
	fmt.Println()

	bytes, _ = ioutil.ReadFile(name + "_sessionPublicKey")
	c.sessionPublicKey.SetBytes(bytes[:])
	fmt.Println("-----SESSION_PUBLIC_KEY-----")
	fmt.Print("session public key: ")
	fmt.Println(c.sessionPublicKey)
	fmt.Println("-----END_SESSION_PUBLIC_KEY-----")
	fmt.Println()
}

