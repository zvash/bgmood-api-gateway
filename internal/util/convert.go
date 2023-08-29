package util

import "encoding/json"

func ConvertTypes[F interface{}, T interface{}](from F, to T) (T, error) {
	fromBytes, err := json.Marshal(from)
	if err != nil {
		return to, err
	}
	err = json.Unmarshal(fromBytes, to)
	if err != nil {
		return to, err
	}
	return to, nil
}
