package main

import (
	"crypto/ecdsa"
	"fmt"
	"os"
)

func deriveSharedKey(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	x, _ := publicKey.Curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	return x.Bytes(), nil
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
