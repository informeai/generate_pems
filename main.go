package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func main() {
	// Gerar a chave privada
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	// Serializar a chave privada para o formato PEM
	privateKeyFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer privateKeyFile.Close()
	privByt, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	privateKeyPEM := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privByt,
	}

	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		panic(err)
	}

	// Gerar a chave pública
	publicKey := &privateKey.PublicKey

	// Serializar a chave pública para o formato PEM
	publicKeyFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer publicKeyFile.Close()
	publicByt, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	publicKeyPEM := &pem.Block{
		Type:  "EC PUBLIC KEY",
		Bytes: publicByt,
	}

	if err := pem.Encode(publicKeyFile, publicKeyPEM); err != nil {
		panic(err)
	}
}
