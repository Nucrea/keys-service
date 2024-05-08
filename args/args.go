package args

import "github.com/jessevdk/go-flags"

type Args interface {
	GetConfigFilePath() string
}

type args struct {
	ConfigFilePath string `short:"c" long:"config" description:"A path to the config file" required:"true"`
}

func (a *args) GetConfigFilePath() string {
	return a.ConfigFilePath
}

func NewArgs(cmdArgs []string) (Args, error) {
	a := &args{}
	if _, err := flags.ParseArgs(a, cmdArgs); err != nil {
		return nil, err
	}

	return a, nil
}
