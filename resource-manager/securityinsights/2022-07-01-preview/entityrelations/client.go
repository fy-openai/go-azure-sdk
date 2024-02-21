package entityrelations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityRelationsClient struct {
	Client *resourcemanager.Client
}

func NewEntityRelationsClientWithBaseURI(sdkApi sdkEnv.Api) (*EntityRelationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "entityrelations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntityRelationsClient: %+v", err)
	}

	return &EntityRelationsClient{
		Client: client,
	}, nil
}
