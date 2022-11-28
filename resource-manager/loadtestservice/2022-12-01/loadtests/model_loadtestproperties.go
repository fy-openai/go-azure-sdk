package loadtests

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoadTestProperties struct {
	DataPlaneURI      *string               `json:"dataPlaneURI,omitempty"`
	Description       *string               `json:"description,omitempty"`
	Encryption        *EncryptionProperties `json:"encryption"`
	ProvisioningState *ResourceState        `json:"provisioningState,omitempty"`
}
