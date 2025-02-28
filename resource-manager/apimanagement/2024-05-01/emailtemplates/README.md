
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/emailtemplates` Documentation

The `emailtemplates` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2024-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/emailtemplates"
```


### Client Initialization

```go
client := emailtemplates.NewEmailTemplatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EmailTemplatesClient.EmailTemplateCreateOrUpdate`

```go
ctx := context.TODO()
id := emailtemplates.NewTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "accountClosedDeveloper")

payload := emailtemplates.EmailTemplateUpdateParameters{
	// ...
}


read, err := client.EmailTemplateCreateOrUpdate(ctx, id, payload, emailtemplates.DefaultEmailTemplateCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmailTemplatesClient.EmailTemplateDelete`

```go
ctx := context.TODO()
id := emailtemplates.NewTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "accountClosedDeveloper")

read, err := client.EmailTemplateDelete(ctx, id, emailtemplates.DefaultEmailTemplateDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmailTemplatesClient.EmailTemplateGet`

```go
ctx := context.TODO()
id := emailtemplates.NewTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "accountClosedDeveloper")

read, err := client.EmailTemplateGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmailTemplatesClient.EmailTemplateGetEntityTag`

```go
ctx := context.TODO()
id := emailtemplates.NewTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "accountClosedDeveloper")

read, err := client.EmailTemplateGetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmailTemplatesClient.EmailTemplateUpdate`

```go
ctx := context.TODO()
id := emailtemplates.NewTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "accountClosedDeveloper")

payload := emailtemplates.EmailTemplateUpdateParameters{
	// ...
}


read, err := client.EmailTemplateUpdate(ctx, id, payload, emailtemplates.DefaultEmailTemplateUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
