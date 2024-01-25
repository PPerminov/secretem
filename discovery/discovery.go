package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func encryptWithECCPublicKey(publicKey *ecdsa.PublicKey, plaintext []byte) ([]byte, error) {
	ciphertext, x, y, err := ecdsa.Encrypt(rand.Reader, publicKey, plaintext)
	if err != nil {
		return nil, err
	}
	// You might want to include x and y in the final result for decryption
	result := append(ciphertext, x.Bytes()...)
	result = append(result, y.Bytes()...)
	return result, nil
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
	// Load public key
	publicKeyFile := "path/to/ecc_public_key.pem"
	publicKey, err := loadECCPublicKey(publicKeyFile)
	if err != nil {
		fmt.Println("Error loading ECC public key:", err)
		os.Exit(1)
	}

	// String to encrypt
	plaintext := []byte("Hello, World!")

	// Encrypt the string with the ECC public key
	ciphertext, err := encryptWithECCPublicKey(publicKey, plaintext)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		os.Exit(1)
	}

	fmt.Printf("Encrypted text: %x\n", ciphertext)
}
