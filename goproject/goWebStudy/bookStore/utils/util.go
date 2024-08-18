package utils

import "encoding/json"

func JsonMarshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return data, err
}

func JsonUnmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	return err
}
