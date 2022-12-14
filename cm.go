package cm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

func loadTOMLConfigure(filename string, config interface{}) error {
	_, err := toml.DecodeFile(filename, config)

	return err
}

func loadJSONConfigure(filename string, config interface{}) error {
	content, err := ioutil.ReadFile(filename)

	if nil != err {
		return err
	}

	err = json.Unmarshal(content, config)

	if nil != err {
		return err
	}

	return nil
}

func loadYAMLConfigure(filename string, config interface{}) error {
	content, err := ioutil.ReadFile(filename)

	if nil != err {
		return err
	}

	err = yaml.Unmarshal(content, config)

	if nil != err {
		return err
	}

	return nil
}

func LoadConfigure(filename string, config interface{}, configureType ConfigureType) error {
	switch configureType {
	case TOML:
		return loadTOMLConfigure(filename, config)
	case JSON:
		return loadJSONConfigure(filename, config)
	case YAML:
		return loadYAMLConfigure(filename, config)
	default:
		return errors.New("Unsupported Type")
	}
}

func Print(config interface{}, isPretty bool) {
	if isPretty {
		jsonObject, err := json.MarshalIndent(config, "", "    ")

		if nil != err {
			fmt.Printf("%+v\n", config)
		}

		fmt.Printf("%s\n", string(jsonObject))
	} else {
		fmt.Printf("%+v\n", config)
	}
}
