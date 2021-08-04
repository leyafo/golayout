package etcd

import (
	"google.golang.org/grpc"
	"testing"
)

func TestGetGrpcConnection(t *testing.T) {
	type args struct {
		endpointsKey string
		serviceName  string
	}
	tests := []struct {
		name     string
		args     args
		wantConn *grpc.ClientConn
		wantErr  bool
	}{
		{
			name: "testGetGrpcConnection",
			args: args{endpointsKey: "service/golayout/", serviceName: "monitor"},
			wantConn: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConn, err := GetGrpcConnection(tt.args.endpointsKey, tt.args.serviceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGrpcConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotConn.GetState())
			if gotConn == nil{
				t.Errorf("GetGrpcConnection() gotConn = %v, want %v", gotConn, tt.wantConn)
			}
		})
	}
}