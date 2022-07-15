package provider

import (
	gen "github.com/james-milligan/flagd-provider-client-go/schemas/protobuf/gen/v1"
)

type IProvider interface {
	// getMetadata() metadata
	ResolveBooleanValue(string, bool, map[string]interface{}, interface{}) (*gen.ResolveBooleanResponse, error)
	ResolveStringValue(string, string, map[string]interface{}, interface{}) (*gen.ResolveStringResponse, error)
	ResolveNumberValue(string, float32, map[string]interface{}, interface{}) (*gen.ResolveNumberResponse, error)
	ResolveObjectValue(string, map[string]interface{}, map[string]interface{}, interface{}) (*gen.ResolveObjectResponse, error)
}
