package machines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineUpdateProperties struct {
	CloudMetadata              *CloudMetadata `json:"cloudMetadata"`
	LocationData               *LocationData  `json:"locationData"`
	OsProfile                  *OSProfile     `json:"osProfile"`
	ParentClusterResourceId    *string        `json:"parentClusterResourceId,omitempty"`
	PrivateLinkScopeResourceId *string        `json:"privateLinkScopeResourceId,omitempty"`
}
