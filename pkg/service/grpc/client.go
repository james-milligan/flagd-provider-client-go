package grpc

import (
	"fmt"
	"time"

	gen "github.com/james-milligan/flagd-provider-client-go/schemas/protobuf/gen/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCServiceConfiguration struct {
	Port int32
}

type GRPCService struct {
	GRPCServiceConfiguration *GRPCServiceConfiguration
	conn                     *grpc.ClientConn
}

func (s *GRPCService) Connect() {
	if s.conn == nil {
		conn, err := grpc.Dial(
			fmt.Sprintf("localhost:%d", s.GRPCServiceConfiguration.Port),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
			grpc.WithTimeout(2*time.Second),
		)
		if err != nil {
			log.Errorf("grpc - fail to dial: %v", err)
			return
		}
		s.conn = conn
	}
}

func (s *GRPCService) GetInstance() gen.ServiceClient {
	s.Connect()
	if s.conn == nil {
		return nil
	}
	return gen.NewServiceClient(s.conn)
}
