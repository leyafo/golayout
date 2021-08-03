package api

import (
	"bytes"
	"context"
	"encoding/json"
	pbCommon "golayout/internal/proto/common"
	"golayout/pkg/version"
	"net/http"
	"strings"
)

func VersionDoc()(doc, input, output string){
	exampleOutput := new(bytes.Buffer)
	versions := make(map[string]*version.Ver)
	versions["api"] = version.GetVersion()

	json.NewEncoder(exampleOutput).Encode(versions)

	return "version's document text write here", "", exampleOutput.String()
}

//Version Get all backend services version
func Version(w http.ResponseWriter, r *http.Request) {
	versions := make(map[string]*version.Ver)
	versions["api"] = version.GetVersion()

	monitorCli := pbCommon.NewCommonClient(nil)
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
