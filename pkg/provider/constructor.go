package provider

import "github.com/james-milligan/flagd-provider-client-go/pkg/service"

type ProviderOption func(*Provider)

func NewProvider(opts ...ProviderOption) *Provider {
	provider := &Provider{}
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
