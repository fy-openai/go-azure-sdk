package resources

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenericResource struct {
	Id         *string                            `json:"id,omitempty"`
	Identity   *identity.SystemAndUserAssignedMap `json:"identity,omitempty"`
	Kind       *string                            `json:"kind,omitempty"`
	Location   *string                            `json:"location,omitempty"`
	ManagedBy  *string                            `json:"managedBy,omitempty"`
	Name       *string                            `json:"name,omitempty"`
	Plan       *Plan                              `json:"plan"`
	Properties *interface{}                       `json:"properties,omitempty"`
	Sku        *Sku                               `json:"sku"`
	Tags       *map[string]string                 `json:"tags,omitempty"`
	Type       *string                            `json:"type,omitempty"`
}
