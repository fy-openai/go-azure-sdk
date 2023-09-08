package groupsitepermission

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdSiteByIdPermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PermissionCollectionResponse
}

type ListGroupByIdSiteByIdPermissionsCompleteResult struct {
	Items []models.PermissionCollectionResponse
}

// ListGroupByIdSiteByIdPermissions ...
func (c GroupSitePermissionClient) ListGroupByIdSiteByIdPermissions(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdPermissionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/permissions", id.ID()),
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
		Values *[]models.PermissionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdPermissionsComplete retrieves all the results into a single object
func (c GroupSitePermissionClient) ListGroupByIdSiteByIdPermissionsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdPermissionsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdPermissionsCompleteMatchingPredicate(ctx, id, models.PermissionCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdPermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSitePermissionClient) ListGroupByIdSiteByIdPermissionsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.PermissionCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdPermissionsCompleteResult, err error) {
	items := make([]models.PermissionCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdPermissions(ctx, id)
	if err != nil {
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

	result = ListGroupByIdSiteByIdPermissionsCompleteResult{
		Items: items,
	}
	return
}
