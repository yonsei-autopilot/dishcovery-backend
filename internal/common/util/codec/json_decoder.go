package codec

import (
	"encoding/json"
	"net/http"
)

func DecodeReq[T any](r *http.Request) (*T, error) {
	defer r.Body.Close()
	var target T
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		return nil, err
	}
	return &target, nil
}

func DecodeRes[T any](r *http.Response) (*T, error) {
	defer r.Body.Close()
	var target T
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		return nil, err
	}
	return &target, nil
}
