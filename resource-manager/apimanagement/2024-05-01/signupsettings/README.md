
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/signupsettings` Documentation

The `signupsettings` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2024-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/signupsettings"
```


### Client Initialization

```go
client := signupsettings.NewSignUpSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SignUpSettingsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := signupsettings.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

payload := signupsettings.PortalSignupSettings{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, signupsettings.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignUpSettingsClient.Get`

```go
ctx := context.TODO()
id := signupsettings.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignUpSettingsClient.GetEntityTag`

```go
ctx := context.TODO()
id := signupsettings.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

read, err := client.GetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignUpSettingsClient.Update`

```go
ctx := context.TODO()
id := signupsettings.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

payload := signupsettings.PortalSignupSettings{
	// ...
}


read, err := client.Update(ctx, id, payload, signupsettings.DefaultUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
