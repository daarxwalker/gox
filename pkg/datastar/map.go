package datastar

import "encoding/json"

type Map map[string]any

func (m Map) Bytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func (m Map) String() string {
	bytes := m.Bytes()
	if len(bytes) == 0 {
		return "{}"
	}
	return string(bytes)
}
