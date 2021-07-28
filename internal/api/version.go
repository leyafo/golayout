package api

import (
	"encoding/json"
	"golayout/pkg/version"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(version.Version)
	w.WriteHeader(http.StatusOK)
}
