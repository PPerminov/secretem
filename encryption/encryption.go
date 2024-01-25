package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func deriveSharedKey(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	x, _ := publicKey.Curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	return x.Bytes(), nil
}

func loadECCPrivateKey(privateKeyFile string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func loadECCPublicKey(publicKeyFile string) (*ecdsa.PublicKey, error) {
	keyBytes, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
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

func main() {
	// Load private key
	privateKeyFile := "path/to/ecc_private_key.pem"
	privateKey, err := loadECCPrivateKey(privateKeyFile)
	if err != nil {
		fmt.Println("Error loading ECC private key:", err)
		os.Exit(1)
	}

	// Load public key
	publicKeyFile := "path/to/ecc_public_key.pem"
	publicKey, err := loadECCPublicKey(publicKeyFile)
	if err != nil {
		fmt.Println("Error loading ECC public key:", err)
		os.Exit(1)
	}

	// Derive shared key
	sharedKey, err := deriveSharedKey(privateKey, publicKey)
	if err != nil {
		fmt.Println("Error deriving shared key:", err)
		os.Exit(1)
	}

	fmt.Printf("Derived shared key: %x\n", sharedKey)
}
