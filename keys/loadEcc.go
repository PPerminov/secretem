package keys

import (
	"chat/logAndErrors"
	"chat/utils"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

func loadECCPrivateKey(privateKeyFile string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := utils.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}
	blockPriv, _ := pem.Decode(keyBytes)
	if blockPriv == nil {
		return nil, logAndErrors.Wrapper(err, "failed to parse PEM block containing the public key")
	}

	a, _ := x509.ParsePKCS8PrivateKey(blockPriv.Bytes)
	if err != nil {
		return nil, err
	}

	return a.(*ecdsa.PrivateKey), nil
}

func loadECCPublicKey(publicKeyFile string) (*ecdsa.PublicKey, error) {
	keyBytes, err := utils.ReadFile(publicKeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, logAndErrors.Wrapper(err, "failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, logAndErrors.Wrapper(err, "")
	}

	return pub.(*ecdsa.PublicKey), nil
}
