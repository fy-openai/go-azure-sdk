package signalr

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpstreamAuthSettings struct {
	ManagedIdentity *ManagedIdentitySettings `json:"managedIdentity"`
	Type            *UpstreamAuthType        `json:"type,omitempty"`
}
