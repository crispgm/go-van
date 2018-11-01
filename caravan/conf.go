package caravan

import (
	"errors"
	"io/ioutil"

	"github.com/crispgm/go-van/deploy"
	yaml "gopkg.in/yaml.v2"
)

// Conf of van
type Conf struct {
	Source      string      `yaml:"src"`
	Destination string      `yaml:"dst"`
	Mode        deploy.Mode `yaml:"deploy_mode"`
	Once        bool        `yaml:"once"`
	Debug       bool        `yaml:"debug"`
	Incremental bool        `yaml:"incremental"`
	Exclude     []string    `yaml:"exclude"`
}

// DefaultConf of Caravan
var DefaultConf = Conf{
	Source:      ".",
	Destination: ".",
	Mode:        deploy.ModeRSync,
	Incremental: true,
	Once:        false,
	Debug:       false,
	Exclude:     []string{".git", ".svn", "/node_modules"},
}

// LoadFrom loads conf from path
func LoadFrom(path, spec string) (*Conf, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf map[string]Conf
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return nil, err
	}
	if spec, ok := conf[spec]; ok {
		return &spec, nil
	}
	return nil, errors.New("No spec")
}
