
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/portalrevision` Documentation

The `portalrevision` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2024-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/portalrevision"
```


### Client Initialization

```go
client := portalrevision.NewPortalRevisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PortalRevisionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := portalrevision.NewPortalRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "portalRevisionIdValue")

payload := portalrevision.PortalRevisionContract{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PortalRevisionClient.Get`

```go
ctx := context.TODO()
id := portalrevision.NewPortalRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "portalRevisionIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PortalRevisionClient.GetEntityTag`

```go
ctx := context.TODO()
id := portalrevision.NewPortalRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "portalRevisionIdValue")

read, err := client.GetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PortalRevisionClient.ListByService`

```go
ctx := context.TODO()
id := portalrevision.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

// alternatively `client.ListByService(ctx, id, portalrevision.DefaultListByServiceOperationOptions())` can be used to do batched pagination
items, err := client.ListByServiceComplete(ctx, id, portalrevision.DefaultListByServiceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PortalRevisionClient.Update`

```go
ctx := context.TODO()
id := portalrevision.NewPortalRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "portalRevisionIdValue")

payload := portalrevision.PortalRevisionContract{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, portalrevision.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
