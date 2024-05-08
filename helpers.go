package main

import (
	"encoding/json"
	"net/http"
)

func decodeRequestParams(r *http.Request, params any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(params); err != nil {
		return err
	}
	return nil
}
