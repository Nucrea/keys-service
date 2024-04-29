package app

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type IRsaGenerator interface {
	NewKey() ([]byte, error)
}

func NewRsaGenerator(bits int) IRsaGenerator {
	return &RsaGenerator{bits}
}

type RsaGenerator struct {
	bits int
}

func (r *RsaGenerator) NewKey() ([]byte, error) {
	key, err := rsa.GenerateKey(rand.Reader, r.bits)
	if err != nil {
		return nil, err
	}

	keyBytes, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyBytes,
	}), nil
}
