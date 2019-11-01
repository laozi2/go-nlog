package conf

import (
	"io/ioutil"

	"github.com/laozi2/go-nlog"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Nlog *log.NlogConf `yaml:"nlog"`
}

func ParseConf(conf *Config, confFile string) error {
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(data), conf)
	if err != nil {
		return err
	}

	return nil
}
