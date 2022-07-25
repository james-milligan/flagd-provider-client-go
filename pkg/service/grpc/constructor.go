package grpc_service

type GRPCServiceOption func(*GRPCServiceConfiguration)

func NewGRPCService(opts ...GRPCServiceOption) *GRPCService {
	const (
		port = 8080
	)
	serviceConfiguration := &GRPCServiceConfiguration{
		Port: port,
	}
	svc := &GRPCService{
		Client: &GRPCClient{
			GRPCServiceConfiguration: serviceConfiguration,
		},
	}
	for _, opt := range opts {
		opt(serviceConfiguration)
	}
	return svc
}

func WithPort(port int32) GRPCServiceOption {
	return func(s *GRPCServiceConfiguration) {
		s.Port = port
	}
}
