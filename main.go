package main

import (
	"crypto/ed25519"
	"fmt"
	"log"
	"time"

	"github.com/o1egl/paseto"
)

func main() {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Println("cannot create keys")
	}

	jsonToken := paseto.JSONToken{
		Expiration: time.Now().Add(24 * time.Hour),
	}

	jsonToken.Set("data", "this is a signed message")
	footer := "some footer"

	// Sign data
	token, err := paseto.NewV2().Sign(privateKey, jsonToken, footer)
	if err != nil {
		log.Println("verify error")
	}
	fmt.Println(token)

	// Verify data
	var newJsonToken paseto.JSONToken
	var newFooter string
	err = paseto.NewV2().Verify(token, publicKey, &newJsonToken, &newFooter)
	if err != nil {
		log.Println("verify error")
	}
}
