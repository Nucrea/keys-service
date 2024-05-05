package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		cmdArgs := []string{}
		args, err := NewArgs(cmdArgs)
		assert.Error(t, err)
		assert.Nil(t, args)
	})

	t.Run("no error", func(t *testing.T) {
		cmdArgs := []string{"-c", "./some_file.path"}
		args, err := NewArgs(cmdArgs)
		assert.NoError(t, err)
		assert.NotNil(t, args)
		assert.Equal(t, cmdArgs[1], args.GetConfigFilePath())
	})
}
