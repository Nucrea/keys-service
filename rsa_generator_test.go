package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRsaGenerator(t *testing.T) {
	generator := RSAGenerator{}

	keyBytes, err := generator.GenerateRSAKey()
	require.NoError(t, err)
	require.NotEmpty(t, keyBytes)
}
