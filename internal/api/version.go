package api

import (
	"context"
	"encoding/json"
	pbCommon "golayout/internal/proto/common"
	"golayout/pkg/etcd"
	"golayout/pkg/logger"
	"golayout/pkg/version"
	"net/http"
	"strings"
)

func VersionDoc()(doc, input, output string){
	versions := make(map[string]*version.Ver)
	versions["api"] = version.GetVersion()

	jsonBytes, _ := json.MarshalIndent(&versions, "", "\t")
	output = string(jsonBytes)

	return "version's document text write here", "", strings.TrimSpace(output)
}

//Version Get all backend services version
func Version(w http.ResponseWriter, r *http.Request) {
	logger.Info("calling Version")
	versions := make(map[string]*version.Ver)
	versions["api"] = version.GetVersion()

	monitorConn, err:= etcd.GetGrpcConnection(etcdCfg.Key, "monitor")
	if err != nil{
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	monitorCli := pbCommon.NewCommonClient(monitorConn)
	replay, err := monitorCli.Version(context.Background(), &pbCommon.VersionRequest{})
	if err != nil{
		logger.Error(err)
		versions["monitor"] = nil
		return
	}else{
		var monitorVersion version.Ver
		err = json.NewDecoder(strings.NewReader(replay.Message)).Decode(&monitorVersion)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		versions["monitor"] = &monitorVersion
	}

	json.NewEncoder(w).Encode(versions)
	w.WriteHeader(http.StatusOK)
}
