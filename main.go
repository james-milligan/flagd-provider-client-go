package main

import (
	"fmt"

	"github.com/james-milligan/flagd-provider-client-go/pkg/provider"
	service "github.com/james-milligan/flagd-provider-client-go/pkg/service/grpc"
)

func main() {
	provider := provider.NewProvider(provider.WithService(
		service.NewGRPCService(
			service.WithPort(8080),
		),
	))
	fmt.Println(provider.ResolveBooleanValue("notMyBoolFlag", true, map[string]interface{}{}))
	fmt.Println(provider.ResolveBooleanValue("myBoolFlag", true, map[string]interface{}{}))
	fmt.Println(provider.ResolveObjectValue("myObjectFlag", map[string]interface{}{"food": "bars"}, map[string]interface{}{}))
	fmt.Println(provider.ResolveStringValue("myStringFlag", "not returned", map[string]interface{}{}))
	fmt.Println(provider.ResolveBooleanValue("isColorYellow", true, map[string]interface{}{"color": "yellow"}))
}
