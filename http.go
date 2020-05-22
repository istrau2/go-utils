package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RespondJson(w *http.ResponseWriter, v interface{}) error {
	(*w).Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(*w).Encode(v)
	if err != nil {
		return err
	}

	return nil
}

func RequestString(r *http.Request) (string, error) {
	bytedata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(bytedata), nil
}
