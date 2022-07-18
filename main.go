package main

import (
	"fmt"

	"github.com/james-milligan/flagd-provider-client-go/pkg/provider"
	gosdk "github.com/open-feature/golang-sdk/pkg/openfeature"
)

func main() {
	gosdk.SetProvider(provider.NewProvider())
	client := gosdk.GetClient("test-client")
	fmt.Println(client.GetBooleanValueDetails("myBoolFlag", true, map[string]interface{}{}))
	fmt.Println(client.GetBooleanValueDetails("notMyBoolFlag", true, map[string]interface{}{}))
	fmt.Println(client.GetStringValue("myStringFlag", "default", map[string]interface{}{}))
	fmt.Println(client.GetNumberValue("myNumberFlag", 12, map[string]interface{}{}))
	fmt.Println(client.GetObjectValue("myObjectFlag", true, map[string]interface{}{}))
	fmt.Println(client.GetBooleanValueDetails("isColorYellow", false, map[string]interface{}{
		"color": "yellow",
	}))
	fmt.Println(client.GetBooleanValue("isColorYellow", false, map[string]interface{}{
		"color": "yellow",
	}))
}
