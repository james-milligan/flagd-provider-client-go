package service

import (
	gen "github.com/james-milligan/flagd-provider-client-go/schemas/protobuf/gen/v1"
	of "github.com/open-feature/golang-sdk/pkg/openfeature"
)

type IServiceOption func(IService)

type IService interface {
	ResolveBoolean(string, of.EvaluationContext, ...IServiceOption) (*gen.ResolveBooleanResponse, error)
	ResolveString(string, of.EvaluationContext, ...IServiceOption) (*gen.ResolveStringResponse, error)
	ResolveNumber(string, of.EvaluationContext, ...IServiceOption) (*gen.ResolveNumberResponse, error)
	ResolveObject(string, of.EvaluationContext, ...IServiceOption) (*gen.ResolveObjectResponse, error)
}
