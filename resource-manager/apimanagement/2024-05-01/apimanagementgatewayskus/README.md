
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/apimanagementgatewayskus` Documentation

The `apimanagementgatewayskus` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2024-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/apimanagementgatewayskus"
```


### Client Initialization

```go
client := apimanagementgatewayskus.NewApiManagementGatewaySkusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApiManagementGatewaySkusClient.ListAvailableSkus`

```go
ctx := context.TODO()
id := apimanagementgatewayskus.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "gatewayValue")

// alternatively `client.ListAvailableSkus(ctx, id)` can be used to do batched pagination
items, err := client.ListAvailableSkusComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
