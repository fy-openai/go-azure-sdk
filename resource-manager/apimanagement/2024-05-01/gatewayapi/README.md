
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/gatewayapi` Documentation

The `gatewayapi` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2024-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/gatewayapi"
```


### Client Initialization

```go
client := gatewayapi.NewGatewayApiClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GatewayApiClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := gatewayapi.NewGatewayApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "gatewayIdValue", "apiIdValue")

payload := gatewayapi.AssociationContract{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GatewayApiClient.Delete`

```go
ctx := context.TODO()
id := gatewayapi.NewGatewayApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "gatewayIdValue", "apiIdValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GatewayApiClient.GetEntityTag`

```go
ctx := context.TODO()
id := gatewayapi.NewGatewayApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "gatewayIdValue", "apiIdValue")

read, err := client.GetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GatewayApiClient.ListByService`

```go
ctx := context.TODO()
id := gatewayapi.NewServiceGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "gatewayIdValue")

// alternatively `client.ListByService(ctx, id, gatewayapi.DefaultListByServiceOperationOptions())` can be used to do batched pagination
items, err := client.ListByServiceComplete(ctx, id, gatewayapi.DefaultListByServiceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
