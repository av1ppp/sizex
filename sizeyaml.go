package sizex

import (
	"strings"

	"gopkg.in/yaml.v3"
)

func (self Size) MarshalYAML() (any, error) {
	return self.String(), nil
}

func (self *Size) UnmarshalYAML(value *yaml.Node) error {
	s, err := ParseSize(strings.Trim(value.Value, "\""))
	if err != nil {
		return err
	}
	*self = s

	return nil
}
