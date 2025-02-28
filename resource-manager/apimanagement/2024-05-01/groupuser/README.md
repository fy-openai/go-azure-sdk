
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/groupuser` Documentation

The `groupuser` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2024-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/groupuser"
```


### Client Initialization

```go
client := groupuser.NewGroupUserClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupUserClient.CheckEntityExists`

```go
ctx := context.TODO()
id := groupuser.NewGroupUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "groupIdValue", "userIdValue")

read, err := client.CheckEntityExists(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupUserClient.Create`

```go
ctx := context.TODO()
id := groupuser.NewGroupUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "groupIdValue", "userIdValue")

read, err := client.Create(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupUserClient.Delete`

```go
ctx := context.TODO()
id := groupuser.NewGroupUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "groupIdValue", "userIdValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupUserClient.List`

```go
ctx := context.TODO()
id := groupuser.NewGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "groupIdValue")

// alternatively `client.List(ctx, id, groupuser.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, groupuser.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupUserClient.WorkspaceGroupUserCheckEntityExists`

```go
ctx := context.TODO()
id := groupuser.NewWorkspaceGroupUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "groupIdValue", "userIdValue")

read, err := client.WorkspaceGroupUserCheckEntityExists(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupUserClient.WorkspaceGroupUserCreate`

```go
ctx := context.TODO()
id := groupuser.NewWorkspaceGroupUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "groupIdValue", "userIdValue")

read, err := client.WorkspaceGroupUserCreate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupUserClient.WorkspaceGroupUserDelete`

```go
ctx := context.TODO()
id := groupuser.NewWorkspaceGroupUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "groupIdValue", "userIdValue")

read, err := client.WorkspaceGroupUserDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupUserClient.WorkspaceGroupUserList`

```go
ctx := context.TODO()
id := groupuser.NewWorkspaceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "groupIdValue")

// alternatively `client.WorkspaceGroupUserList(ctx, id, groupuser.DefaultWorkspaceGroupUserListOperationOptions())` can be used to do batched pagination
items, err := client.WorkspaceGroupUserListComplete(ctx, id, groupuser.DefaultWorkspaceGroupUserListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
