package keys

import (
	"chat/utils"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
)

func loadRSAPrivateKey(privateKeyFile string) (*rsa.PrivateKey, error) {
	keyBytes, err := utils.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	return privateKey, nil
}

func loadRSAPublicKey(publicKeyFile string) (*rsa.PublicKey, error) {
	keyBytes, err := utils.ReadFile(publicKeyFile)
	if err != nil {
		return nil, err
	}
	// Decode the PEM data
	block, _ := pem.Decode(keyBytes)
	// Parse the RSA public key
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	return publicKey, nil
}
