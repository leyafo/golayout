package monitor

import (
	"context"
	"encoding/json"
	pbCommon "golayout/internal/proto/common"
	"golayout/pkg/version"
	"log"
	"strings"
)

type Version struct{
	pbCommon.UnimplementedCommonServer
}

func (s *Version)Version(ctx context.Context, in *pbCommon.VersionRequest) (*pbCommon.VersionReply, error) {
	log.Printf("Received: %v", in.GetName())
	v := version.GetVersion()
	b := &strings.Builder{}
	err := json.NewEncoder(b).Encode(&v)
	if err != nil{
		return nil, err
	}
	return &pbCommon.VersionReply{Message: b.String()}, nil
}
