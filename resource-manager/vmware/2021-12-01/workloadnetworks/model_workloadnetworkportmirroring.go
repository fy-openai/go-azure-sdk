package workloadnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkPortMirroring struct {
	Id         *string                                 `json:"id,omitempty"`
	Name       *string                                 `json:"name,omitempty"`
	Properties *WorkloadNetworkPortMirroringProperties `json:"properties"`
	Type       *string                                 `json:"type,omitempty"`
}
