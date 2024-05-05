package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRsaGeneratorNewKey(t *testing.T) {
	generator := NewRsaGenerator(2048)

	keyBytes, err := generator.NewKey()
	require.NoError(t, err)
	require.NotEmpty(t, keyBytes)
	require.Contains(t, string(keyBytes), "PRIVATE KEY")
}

func BenchmarkRsaGenerator(t *testing.B) {
	generator := NewRsaGenerator(2048)
	for i := 0; i < t.N; i++ {
		generator.NewKey()
	}
}
