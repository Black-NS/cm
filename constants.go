package cm

const (
	_ = iota
	JSON
	TOML
	YAML
	INI
)

type ConfigureType int

/**
Mapper
*/

type INIConfigures struct {
	SectionName string
	Configure   interface{}
}
