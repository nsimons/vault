package plugin

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hashicorp/vault/sdk/plugin/pb"
)

func TestStorage_GRPC_ReturnsErrIfStorageNil(t *testing.T) {
	_, err := new(GRPCStorageServer).Get(context.Background(), nil)
	if err == nil {
		t.Error("Expected error when using server with no impl")
	}
}

func TestStorage_impl(t *testing.T) {
	var _ logical.Storage = new(GRPCStorageClient)
}

func TestStorage_GRPC(t *testing.T) {
	storage := &logical.InmemStorage{}
	client, _ := plugin.TestGRPCConn(t, func(s *grpc.Server) {
		pb.RegisterStorageServer(s, &GRPCStorageServer{
			impl: storage,
		})
	})
	defer client.Close()

	testStorage := &GRPCStorageClient{client: pb.NewStorageClient(client)}

	logical.TestStorage(t, testStorage)
}
