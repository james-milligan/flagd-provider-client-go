package grpc_service

type GRPCServiceOption func(*GRPCService)

func NewGRPCService(opts ...GRPCServiceOption) *GRPCService {
	const (
		port = 8080
	)
	svc := &GRPCService{
		GRPCServiceConfiguration: &GRPCServiceConfiguration{
			Port: port,
		},
	}
	for _, opt := range opts {
		opt(svc)
	}
	return svc
}

func WithPort(port int32) GRPCServiceOption {
	return func(s *GRPCService) {
		s.GRPCServiceConfiguration.Port = port
	}
}
