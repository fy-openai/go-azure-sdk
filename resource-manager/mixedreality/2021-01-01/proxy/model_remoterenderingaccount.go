package proxy

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteRenderingAccount struct {
	Id         *string                        `json:"id,omitempty"`
	Identity   *identity.SystemAssigned       `json:"identity,omitempty"`
	Kind       *Sku                           `json:"kind"`
	Location   string                         `json:"location"`
	Name       *string                        `json:"name,omitempty"`
	Plan       *identity.SystemAssigned       `json:"plan,omitempty"`
	Properties *MixedRealityAccountProperties `json:"properties"`
	Sku        *Sku                           `json:"sku"`
	SystemData *systemdata.SystemData         `json:"systemData,omitempty"`
	Tags       *map[string]string             `json:"tags,omitempty"`
	Type       *string                        `json:"type,omitempty"`
}
