package provider

import (
	"github.com/james-milligan/flagd-provider-client-go/pkg/service"
	HTTPService "github.com/james-milligan/flagd-provider-client-go/pkg/service/http"
)

type ProviderOption func(*Provider)

func NewProvider(opts ...ProviderOption) *Provider {
	provider := &Provider{
		service: &HTTPService.HTTPService{
			HTTPServiceConfiguration: &HTTPService.HTTPServiceConfiguration{
				Port: 8080,
			},
		},
	}
	for _, opt := range opts {
		opt(provider)
	}
	return provider
}

func WithService(service service.IService) ProviderOption {
	return func(p *Provider) {
		p.service = service
	}
}
