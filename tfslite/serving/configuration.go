// Patrick Wieschollek <mail@patwie.com>

package serving

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Configure Ports etc.
type Configuration struct {
	TensorFlowServingLite struct {
		RestPort int `yaml:"rest_port"`
	} `yaml:"tensorflow_serving_lite"`
	Models []struct {
		Path    string `yaml:"path"`
		Name    string `yaml:"name"`
		Version int    `yaml:"version"`
	} `yaml:"models"`
}

// parse configuration from environment
func GetConfiguration(fn string) (conf *Configuration) {

	y, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(y, &conf)
	if err != nil {
		panic(err)
	}

	return conf
}
