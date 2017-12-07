package receiverKeyHandler

import "crypto/rsa"
import "golang.org/x/crypto/openpgp/elgamal"

type ReceiverKeyHandlerImpl struct{
	publicKeyRSA   *rsa.PublicKey
	publicKeyElGamal *elgamal.PublicKey
}
