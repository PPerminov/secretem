package keys

import (
	"chat/logAndErrors"
	"chat/utils"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func loadRSAPrivateKey(privateKeyFile string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := utils.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, logAndErrors.Wrapper(err, "failed to parse PEM block containing the private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func loadRSAPublicKey(publicKeyFile string) (*ecdsa.PublicKey, error) {
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
		return nil, err
	}

	publicKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to parse public key")
	}

	return publicKey, nil
}
