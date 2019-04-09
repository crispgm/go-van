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
	ExtraArgs   []string    `yaml:"extra_args"`
	LogFormat   string      `yaml:"log_format"`
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
	ExtraArgs:   []string{"--delete"},
	LogFormat:   "[%t] EVENT %e: %p",
}

const (
	// DefaultSpec ...
	DefaultSpec = "master"
	// DefaultConfName ...
	DefaultConfName = "caravan.yml"
)

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
	if specConf, ok := conf[spec]; ok {
		return &specConf, nil
	}
	if spec == DefaultSpec && len(conf) == 1 {
		for specName, specConf := range conf {
			PrintNotice("No spec name specified, choose", specName, "instead")
			return &specConf, nil
		}
	}
	return nil, errors.New("No spec")
}

// CreateDefault conf
func CreateDefault(path string) error {
	var defaultConf = make(map[string]Conf, 0)
	defaultConf[DefaultSpec] = DefaultConf
	out, err := yaml.Marshal(defaultConf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, out, 0666)
}
