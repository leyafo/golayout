package api

import (
	"encoding/json"
	"golayout/pkg/logger"
	"golayout/pkg/version"
	"net/http"
	"strings"
)

func VersionDoc() (doc, input, output string) {
	versions := make(map[string]*version.Ver)
	versions["api"] = version.GetVersion()

	jsonBytes, _ := json.MarshalIndent(&versions, "", "\t")
	output = string(jsonBytes)

	return "version's document text write here", "nil", strings.TrimSpace(output)
}

//Version Get all backend services version
func Version(w http.ResponseWriter, r *http.Request) {
	logger.Info("calling Version")
	versions := make(map[string]*version.Ver)
	versions["api"] = version.GetVersion()

	json.NewEncoder(w).Encode(versions)
	w.WriteHeader(http.StatusOK)
}
