package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	conf, err := NewConfig("./config_test.yaml")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, uint(16), conf.GetThreadsCount())
	assert.Equal(t, uint(5000), conf.GetStackSize())
	assert.Equal(t, uint16(8080), conf.GetPort())
}
