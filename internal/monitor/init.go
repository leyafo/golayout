package monitor

import (
	pbCommon "golayout/internal/proto/common"
	"google.golang.org/grpc"
)

func RegisterServerStub(s *grpc.Server) {
	pbCommon.RegisterCommonServer(s, &Version{})
}
