package rpcclient

import (
	"fmt"

	"github.com/designsbysm/timber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(host string, port string) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	timber.Info(fmt.Sprintf("RPC: sending on %s", address))

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	return grpc.Dial(address, opts)
}
