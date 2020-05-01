package utils

import (
	"encoding/json"
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
