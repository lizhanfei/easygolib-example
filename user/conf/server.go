package conf

import "time"

type ServerConf struct {
	Address      string        `yaml:"address"`
	ReadTimeout  time.Duration `yaml:"readtimeout"`
	WriteTimeout time.Duration `yaml:"writetimeout"`
	CloseTimeout time.Duration `yaml:"closetimeout"`
}
