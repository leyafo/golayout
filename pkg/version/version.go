package version

import (
	"fmt"
)

var (
	// RELEASE returns the release version
	RELEASE = "UNKNOWN"
	// REPO returns the git repository URL
	REPO = "UNKNOWN"
	// COMMIT returns the short sha from git
	COMMIT = "UNKNOWN"

	SERVICENAME= "UNKNOWN"

	BUILDTIME = "UNKNOWN"

	API   = "v1"

	Version Ver
)

type Ver struct {
	Release string `json:"release"`
	Repository string `json:"repository"`
	Commit string `json:"commit"`
	ServiceName string `json:"service_name"`
	BuildTime string `json:"build_time"`
	API string `json:"api"`
}

func init(){
	Version = Ver{
		Release:     RELEASE,
		Repository:  REPO,
		Commit:      COMMIT,
		ServiceName: SERVICENAME,
		BuildTime:   BUILDTIME,
		API: API,
	}
}
func Short()string{
	return Version.ServiceName + " "+ Version.Release
}
func Long()string{
	return fmt.Sprintf("%+v", Version)
}
