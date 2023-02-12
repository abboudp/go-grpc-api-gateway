package order

import (
	"fmt"

	"github.com/abboudp/go-grpc-api-gateway/pkg/config"
	"github.com/abboudp/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderSerivceClient
}

func InitServiceClient(c *config.Config) pb.OrderSerivceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Count not connect", err)
	}

	return pb.NewOrderSerivceClient(cc)
}
