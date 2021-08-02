package api

import (
	"context"
	"encoding/json"
	pbCommon "golayout/internal/proto/common"
	"golayout/pkg/version"
	"net/http"
	"strings"
)

//Version Get all backend services version
func (s *Server)Version(w http.ResponseWriter, r *http.Request) {
	versions := make(map[string]*version.Ver)
	versions["api"] = version.GetVersion()

	monitorCli := pbCommon.NewCommonClient(s.grpcMonitorBackend)
	replay, err := monitorCli.Version(context.TODO(), &pbCommon.VersionRequest{})
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var monitorVersion version.Ver
	err = json.NewDecoder(strings.NewReader(replay.Message)).Decode(&monitorVersion)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	versions["monitor"] = &monitorVersion

	json.NewEncoder(w).Encode(versions)
	w.WriteHeader(http.StatusOK)
}
