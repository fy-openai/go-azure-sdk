package apidiagnostic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceApiDiagnosticListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DiagnosticContract
}

type WorkspaceApiDiagnosticListByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DiagnosticContract
}

type WorkspaceApiDiagnosticListByWorkspaceOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultWorkspaceApiDiagnosticListByWorkspaceOperationOptions() WorkspaceApiDiagnosticListByWorkspaceOperationOptions {
	return WorkspaceApiDiagnosticListByWorkspaceOperationOptions{}
}

func (o WorkspaceApiDiagnosticListByWorkspaceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WorkspaceApiDiagnosticListByWorkspaceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o WorkspaceApiDiagnosticListByWorkspaceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type WorkspaceApiDiagnosticListByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkspaceApiDiagnosticListByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkspaceApiDiagnosticListByWorkspace ...
func (c ApiDiagnosticClient) WorkspaceApiDiagnosticListByWorkspace(ctx context.Context, id WorkspaceApiId, options WorkspaceApiDiagnosticListByWorkspaceOperationOptions) (result WorkspaceApiDiagnosticListByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &WorkspaceApiDiagnosticListByWorkspaceCustomPager{},
		Path:          fmt.Sprintf("%s/diagnostics", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]DiagnosticContract `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkspaceApiDiagnosticListByWorkspaceComplete retrieves all the results into a single object
func (c ApiDiagnosticClient) WorkspaceApiDiagnosticListByWorkspaceComplete(ctx context.Context, id WorkspaceApiId, options WorkspaceApiDiagnosticListByWorkspaceOperationOptions) (WorkspaceApiDiagnosticListByWorkspaceCompleteResult, error) {
	return c.WorkspaceApiDiagnosticListByWorkspaceCompleteMatchingPredicate(ctx, id, options, DiagnosticContractOperationPredicate{})
}

// WorkspaceApiDiagnosticListByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApiDiagnosticClient) WorkspaceApiDiagnosticListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceApiId, options WorkspaceApiDiagnosticListByWorkspaceOperationOptions, predicate DiagnosticContractOperationPredicate) (result WorkspaceApiDiagnosticListByWorkspaceCompleteResult, err error) {
	items := make([]DiagnosticContract, 0)

	resp, err := c.WorkspaceApiDiagnosticListByWorkspace(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = WorkspaceApiDiagnosticListByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
