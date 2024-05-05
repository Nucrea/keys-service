package app

import (
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type Config interface {
	GetThreadsCount() uint
	GetStackSize() uint
	GetPort() uint16
}

type config struct {
	Threads   uint   `yaml:"threads" validate:"required"`
	StackSize uint   `yaml:"stackSize" validate:"required"`
	Port      uint16 `yaml:"port" validate:"required"`
}

func (c *config) GetPort() uint16 {
	return c.Port
}

func (c *config) GetStackSize() uint {
	return c.StackSize
}

func (c *config) GetThreadsCount() uint {
	return c.Threads
}

func NewConfig(path string) (Config, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &config{}
	if err := yaml.Unmarshal(fileBytes, conf); err != nil {
		return nil, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
