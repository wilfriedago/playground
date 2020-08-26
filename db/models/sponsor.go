package models

import (
	"encoding/json"
)

type Sponsor struct {
	Name  string `json:"name" redis:"name"`
	Zoom    string `json:"zoom" redis:"zoom"`
	ID    string `json:"id" redis:"-"`
}

func (s Sponsor) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s Sponsor) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
