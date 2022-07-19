package main

import (
	"fmt"

	"github.com/james-milligan/flagd-provider-client-go/pkg/provider"
	gosdk "github.com/open-feature/golang-sdk/pkg/openfeature"
)

func main() {
	gosdk.SetProvider(provider.NewProvider())
	client := gosdk.GetClient("test-client")

	fmt.Printf("fetching value %s, %s\n", "myBoolFlag", "flag exists and will not fail")
	fmt.Println(client.GetBooleanValueDetails("myBoolFlag", true, map[string]interface{}{}))

	fmt.Printf("\nfetching value %s, %s\n", "notMyBoolFlag", "flag does not exist and will fail with error code FLAG_NOT_FOUND")
	fmt.Println(client.GetBooleanValueDetails("notMyBoolFlag", true, map[string]interface{}{}))

	fmt.Printf("\nfetching value %s, %s\n", "myStringFlag", "flag exists and will not fail")
	fmt.Println(client.GetStringValue("myStringFlag", "default", map[string]interface{}{}))

	fmt.Printf("\nfetching value %s, %s\n", "myNumberFlag", "TYPE_MISMATCH will result in the default value being returned")
	fmt.Println(client.GetStringValueDetails("myNumberFlag", "12", map[string]interface{}{}))

	fmt.Printf("\nfetching value %s, %s\n", "myObjectFlag", "flag exists and will not fail")
	fmt.Println(client.GetObjectValue("myObjectFlag", true, map[string]interface{}{}))

	fmt.Printf("\nfetching value %s, %s\n", "isColorYellow", "TARGETING_MATCH will return on variant with value true")
	fmt.Println(client.GetBooleanValueDetails("isColorYellow", false, map[string]interface{}{
		"color": "yellow",
	}))

	fmt.Printf("\nfetching value %s, %s\n", "myBoolFlag", "TARGETING_MATCH will return on variant with value false")
	fmt.Println(client.GetBooleanValueDetails("isColorYellow", false, map[string]interface{}{
		"color": "not-yellow",
	}))
}
