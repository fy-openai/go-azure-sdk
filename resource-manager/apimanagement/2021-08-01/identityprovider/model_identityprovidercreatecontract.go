package identityprovider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityProviderCreateContract struct {
	Id         *string                                   `json:"id,omitempty"`
	Name       *string                                   `json:"name,omitempty"`
	Properties *IdentityProviderCreateContractProperties `json:"properties"`
	Type       *string                                   `json:"type,omitempty"`
}
