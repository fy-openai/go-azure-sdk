package networksecurityperimeter

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterClient struct {
	Client *resourcemanager.Client
}

func NewNetworkSecurityPerimeterClientWithBaseURI(sdkApi sdkEnv.Api) (*NetworkSecurityPerimeterClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "networksecurityperimeter", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkSecurityPerimeterClient: %+v", err)
	}

	return &NetworkSecurityPerimeterClient{
		Client: client,
	}, nil
}
