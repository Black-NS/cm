package cm

import (
	"testing"
)

type TestInformation struct {
	Server   ServerInfo                     `toml:"server" json:"server" yaml:"server"`
	Database map[string]DatabaseInformation `toml:"database" json:"database" yaml:"database"`
}

type TestInformationINI struct {
	Server   ServerInfo          `json:"server" ini:"server"`
	Database DatabaseInformation `json:"database" ini:"database"`
}

type ServerInfo struct {
	Port    int  `toml:"port" json:"port" yaml:"port" ini:"port"`
	Release bool `toml:"release" json:"release" yaml:"release" ini:"release"`
}

type DatabaseInformation struct {
	DriverURI    string `toml:"uri" json:"uri" yaml:"uri" ini:"uri"`
	UsedDatabase string `toml:"used" json:"used" yaml:"used"`
	Table        string `toml:"table" json:"table" yaml:"table"`
	Collection   string `toml:"collection" json:"collection" yaml:"collection"`
}

type AppInformation struct {
	AppName string `json:"appName" ini:"appName"`
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

func TestINI(t *testing.T) {
	var test TestInformationINI

	err := loadINIConfigure("test.ini", &test)

	if nil != err {
		t.Fatal(err)
	}

	Print(test, true)
}

func TestINIWithDefaultSection(t *testing.T) {
	var appInfo1 AppInformation
	var appInfo2 AppInformation

	err := loadINIConfigure("test.ini", &appInfo1)

	if nil != err {
		t.Fatal(err)
	}

	Print(appInfo1, true)

	err = loadINIConfigureWithSectionDefault("test.ini", &appInfo2)

	if nil != err {
		t.Fatal(err)
	}

	Print(appInfo2, true)
}

func TestIniWithCustomSection(t *testing.T) {
	var database DatabaseInformation

	err := loadINIConfigureWithSectionName("test.ini", &database, "database")

	if nil != err {
		t.Fatal(err)
	}

	Print(database, true)
}

func TestIniWithCustomSections(t *testing.T) {
	configures := []INIConfigures{
		{SectionName: "database", Configure: new(DatabaseInformation)},
		{SectionName: "server", Configure: new(ServerInfo)},
	}

	err := loadINIConfigureWithSectionNames("test.ini", configures)

	if nil != err {
		t.Fatal(err)
	}

	for _, value := range configures {
		Print(value.Configure, true)
	}
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
