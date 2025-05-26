package datastar

import "encoding/json"

type Map map[string]any

func (m Map) Bytes() ([]byte, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func (m Map) String() string {
	bytes, err := m.Bytes()
	if err != nil {
		return "{}"
	}
	return string(bytes)
}
