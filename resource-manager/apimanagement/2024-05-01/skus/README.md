
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/skus` Documentation

The `skus` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2024-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/skus"
```


### Client Initialization

```go
client := skus.NewSkusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SkusClient.ApiManagementSkusList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ApiManagementSkusList(ctx, id)` can be used to do batched pagination
items, err := client.ApiManagementSkusListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
