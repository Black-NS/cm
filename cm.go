package cm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"gopkg.in/ini.v1"
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

func loadINIContent(filename string) (*ini.File, error) {
	return ini.Load(filename)
}

func loadINIConfigure(filename string, config interface{}) error {
	content, err := loadINIContent(filename)

	if nil != err {
		return err
	}

	err = content.MapTo(config)

	if nil != err {
		return err
	}

	return nil
}

func loadINIConfigureWithSectionDefault(filename string, config interface{}) error {
	var configures INIConfigures

	configures.Configure = config
	configures.SectionName = ini.DefaultSection

	return loadINIConfigureWithSectionNames(filename, []INIConfigures{configures})
}

func loadINIConfigureWithSectionName(filename string, config interface{}, sectionName string) error {
	var configures INIConfigures

	configures.Configure = config
	configures.SectionName = sectionName

	return loadINIConfigureWithSectionNames(filename, []INIConfigures{configures})

}

func loadINIConfigureWithSectionNames(filename string, configures []INIConfigures) error {
	content, err := loadINIContent(filename)

	if nil != err {
		return err
	}

	for _, value := range configures {
		err = content.Section(value.SectionName).MapTo(value.Configure)

		if nil != err {
			return err
		}
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
	case INI:
		return loadINIConfigure(filename, config)
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
