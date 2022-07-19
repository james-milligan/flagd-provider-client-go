package grpc_service

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	schemaV1 "go.buf.build/grpc/go/james-milligan/flagd-schema-go/schema/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

func (s *GRPCService) GetInstance() schemaV1.ServiceClient {
	s.Connect()
	if s.conn == nil {
		return nil
	}
	return schemaV1.NewServiceClient(s.conn)
}
