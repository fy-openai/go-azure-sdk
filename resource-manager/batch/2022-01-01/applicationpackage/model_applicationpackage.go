package applicationpackage

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationPackage struct {
	Etag       *string                       `json:"etag,omitempty"`
	Id         *string                       `json:"id,omitempty"`
	Name       *string                       `json:"name,omitempty"`
	Properties *ApplicationPackageProperties `json:"properties"`
	Type       *string                       `json:"type,omitempty"`
}
