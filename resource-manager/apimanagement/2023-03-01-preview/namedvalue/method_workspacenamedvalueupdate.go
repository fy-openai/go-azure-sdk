package namedvalue

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceNamedValueUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type WorkspaceNamedValueUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultWorkspaceNamedValueUpdateOperationOptions() WorkspaceNamedValueUpdateOperationOptions {
	return WorkspaceNamedValueUpdateOperationOptions{}
}

func (o WorkspaceNamedValueUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o WorkspaceNamedValueUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o WorkspaceNamedValueUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// WorkspaceNamedValueUpdate ...
func (c NamedValueClient) WorkspaceNamedValueUpdate(ctx context.Context, id WorkspaceNamedValueId, input NamedValueUpdateParameters, options WorkspaceNamedValueUpdateOperationOptions) (result WorkspaceNamedValueUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		Path:          id.ID(),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// WorkspaceNamedValueUpdateThenPoll performs WorkspaceNamedValueUpdate then polls until it's completed
func (c NamedValueClient) WorkspaceNamedValueUpdateThenPoll(ctx context.Context, id WorkspaceNamedValueId, input NamedValueUpdateParameters, options WorkspaceNamedValueUpdateOperationOptions) error {
	result, err := c.WorkspaceNamedValueUpdate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing WorkspaceNamedValueUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after WorkspaceNamedValueUpdate: %+v", err)
	}

	return nil
}
