package api

import (
	"encoding/json"
)

func Unmarshal[T any](data []byte) (T, error) {
	var element T
	err := json.Unmarshal(data, &element)
	if err != nil {
		var zero T
		return zero, err
	}
	return element, nil
}