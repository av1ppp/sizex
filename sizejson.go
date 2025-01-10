package sizex

import "encoding/json"

func (self Size) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.String())
}

func (self *Size) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	s, err := ParseSize(v)
	if err != nil {
		return err
	}
	*self = s

	return nil
}
