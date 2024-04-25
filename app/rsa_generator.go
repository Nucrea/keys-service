package app

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

type RSAGenerator struct{}

func (r RSAGenerator) GenerateRSAKey() ([]byte, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	keyBytes, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}

	return keyBytes, nil
}
