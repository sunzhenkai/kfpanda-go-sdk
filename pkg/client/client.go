package client

import (
	"time"

	"github.com/sunzhenkai/kfpanda-go-sdk/pkg/logging"
	"github.com/sunzhenkai/kung-fu-panda-protocols/go/service/kfpanda"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type ClientConfig struct {
	ServerUri string
	Timeout   time.Duration
}

func NewKfPandaClient(cfg *ClientConfig) (kfpanda.KfPandaServiceClient, error) {
	opts := []grpc.DialOption{
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    30 * time.Second,
			Timeout: cfg.Timeout,
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(cfg.ServerUri, opts...)
	if err != nil {
		logging.Logger.Infoln("create ranker grpc client failed.")
		return nil, err
	}
	return kfpanda.NewKfPandaServiceClient(conn), nil
}
