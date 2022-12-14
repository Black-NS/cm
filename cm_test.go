package cm

import (
	"testing"
)

type TestInformation struct {
	Server   ServerInfo                     `toml:"server" json:"server" yaml:"server"`
	Database map[string]DatabaseInformation `toml:"database" json:"database" yaml:"database"`
}

type ServerInfo struct {
	Port    int  `toml:"port" json:"port" yaml:"port"`
	Release bool `toml:"release" json:"release" yaml:"release"`
}

type DatabaseInformation struct {
	DriverURI    string `toml:"uri" json:"uri" yaml:"uri"`
	UsedDatabase string `toml:"used" json:"used" yaml:"used"`
	Table        string `toml:"table" json:"table" yaml:"table"`
	Collection   string `toml:"collection" json:"collection" yaml:"collection"`
}

func TestPrint(t *testing.T) {
	test := ServerInfo{
		Port:    8080,
		Release: false,
	}

	Print(test, true)
	Print(test, false)
}

func TestToml(t *testing.T) {
	var test TestInformation
	err := loadTOMLConfigure("test.toml", &test)

	if nil != err {
		t.Fatal(err)
	}

	Print(test, true)
}

func TestJSON(t *testing.T) {
	var test TestInformation
	err := loadJSONConfigure("test.json", &test)

	if nil != err {
		t.Fatal(err)
	}

	Print(test, true)
}

func TestYaml(t *testing.T) {
	var test TestInformation

	err := loadYAMLConfigure("test.yaml", &test)

	if nil != err {
		t.Fatal(err)
	}

	Print(test, true)
}

func TestLoadConfigure(t *testing.T) {
	var testYaml TestInformation
	var testJson TestInformation
	var testToml TestInformation

	err := loadTOMLConfigure("test.toml", &testToml)

	if nil != err {
		t.Fatal(err)
	}

	err = loadJSONConfigure("test.json", &testJson)

	if nil != err {
		t.Fatal(err)
	}

	err = loadYAMLConfigure("test.yaml", &testYaml)

	if nil != err {
		t.Fatal(err)
	}

	Print(testToml, true)
	Print(testJson, true)
	Print(testToml, true)
}
