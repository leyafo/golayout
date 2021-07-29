package api

import (
	"encoding/json"
	"golayout/pkg/version"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {
	var versions []*version.Ver
	versions = append(versions, version.GetVersion())


	json.NewEncoder(w).Encode(versions)
	w.WriteHeader(http.StatusOK)
}
