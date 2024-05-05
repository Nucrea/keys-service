package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Run("err file not exist", func(t *testing.T) {
		_, err := NewConfig("./not_existing_config.yaml")
		require.Error(t, err)
	})

	t.Run("err validating config", func(t *testing.T) {
		_, err := NewConfig("./config_bad_test.yaml")
		require.Error(t, err)
	})

	t.Run("all ok", func(t *testing.T) {
		conf, err := NewConfig("./config_test.yaml")
		require.NoError(t, err)
		assert.Equal(t, uint(16), conf.GetThreadsCount())
		assert.Equal(t, uint(5000), conf.GetStackSize())
		assert.Equal(t, uint16(8080), conf.GetPort())
	})
}
