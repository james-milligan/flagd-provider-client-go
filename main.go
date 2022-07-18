package main

import (
	"fmt"

	"github.com/james-milligan/flagd-provider-client-go/pkg/provider"
	service "github.com/james-milligan/flagd-provider-client-go/pkg/service/grpc"
	gosdk "github.com/open-feature/golang-sdk/pkg/openfeature"
)

func main() {
	gosdk.SetProvider(provider.NewProvider(provider.WithService(
		service.NewGRPCService(
			service.WithPort(8080),
		),
	)))
	client := gosdk.GetClient("test-client")
	count := 1
	fmt.Println(count)
	count++
	fmt.Println(client.GetBooleanValue("myBoolFlag", true, map[string]interface{}{}))
	fmt.Println(count)
	count++
	fmt.Println(client.GetBooleanValueDetails("notMyBoolFlag", true, map[string]interface{}{}))
	fmt.Println(count)
	count++
	fmt.Println(client.GetStringValue("myStringFlag", "default", map[string]interface{}{}))
	fmt.Println(count)
	count++
	fmt.Println(client.GetNumberValue("myNumberFlag", 12, map[string]interface{}{}))
	fmt.Println(count)
	count++
	fmt.Println(client.GetObjectValue("myObjectFlag", true, map[string]interface{}{}))
	fmt.Println(count)
	count++
	fmt.Println(client.GetBooleanValueDetails("isColorYellow", false, map[string]interface{}{
		"color": "yellow",
	}))
	fmt.Println(client.GetBooleanValue("isColorYellow", false, map[string]interface{}{
		"color": "yellow",
	}))
	fmt.Println(count)
	count++
}
