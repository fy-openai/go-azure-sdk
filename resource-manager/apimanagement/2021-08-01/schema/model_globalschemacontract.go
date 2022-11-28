package schema

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalSchemaContract struct {
	Id         *string                         `json:"id,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Properties *GlobalSchemaContractProperties `json:"properties"`
	Type       *string                         `json:"type,omitempty"`
}
