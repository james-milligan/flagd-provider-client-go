package provider

import (
	"github.com/james-milligan/flagd-provider-client-go/pkg/service"
	gen "github.com/james-milligan/flagd-provider-client-go/schemas/protobuf/gen/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

type Provider struct {
	service service.IService
}

func (p *Provider) ResolveBooleanValue(flagKey string, defaultValue bool, context map[string]interface{}, options ...service.ISercviceOption) (*gen.ResolveBooleanResponse, error) {
	return p.service.ResolveBoolean(flagKey, defaultValue, context, options...)
}

func (p *Provider) ResolveStringValue(flagKey string, defaultValue string, context map[string]interface{}, options ...service.ISercviceOption) (*gen.ResolveStringResponse, error) {
	return p.service.ResolveString(flagKey, defaultValue, context, options...)
}

func (p *Provider) ResolveNumberValue(flagKey string, defaultValue float32, context map[string]interface{}, options ...service.ISercviceOption) (*gen.ResolveNumberResponse, error) {
	return p.service.ResolveNumber(flagKey, defaultValue, context, options...)
}

func (p *Provider) ResolveObjectValue(flagKey string, defaultValue map[string]interface{}, context map[string]interface{}, options ...service.ISercviceOption) (*gen.ResolveObjectResponse, error) {
	pbstruct, err := structpb.NewStruct(defaultValue)
	if err != nil {
		return &gen.ResolveObjectResponse{
			Value:   nil,
			Reason:  "INTERNAL ERROR",
			Variant: "null",
		}, err
	}
	return p.service.ResolveObject(flagKey, pbstruct, context, options...)
}
