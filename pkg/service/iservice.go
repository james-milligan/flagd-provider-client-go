package service

import (
	gen "github.com/james-milligan/flagd-provider-client-go/schemas/protobuf/gen/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

type ISercviceOption func(IService)

type IService interface {
	ResolveBoolean(string, bool, map[string]interface{}, ...ISercviceOption) (*gen.ResolveBooleanResponse, error)
	ResolveString(string, string, map[string]interface{}, ...ISercviceOption) (*gen.ResolveStringResponse, error)
	ResolveNumber(string, float32, map[string]interface{}, ...ISercviceOption) (*gen.ResolveNumberResponse, error)
	ResolveObject(string, *structpb.Struct, map[string]interface{}, ...ISercviceOption) (*gen.ResolveObjectResponse, error)
}
