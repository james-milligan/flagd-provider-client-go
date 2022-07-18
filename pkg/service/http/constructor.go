package http_service

type HTTPServiceOption func(*HTTPService)

func NewHTTPService(opts ...HTTPServiceOption) *HTTPService {
	const (
		port = 8080
	)
	svc := &HTTPService{
		HTTPServiceConfiguration: &HTTPServiceConfiguration{
			Port:    port,
			Address: "localhost",
		},
	}
	for _, opt := range opts {
		opt(svc)
	}
	return svc
}

func WithPort(port int32) HTTPServiceOption {
	return func(s *HTTPService) {
		s.HTTPServiceConfiguration.Port = port
	}
}
