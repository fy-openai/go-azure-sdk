package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceEditionCapability struct {
	IsGeneralPurposeV2           *bool                              `json:"isGeneralPurposeV2,omitempty"`
	Name                         *string                            `json:"name,omitempty"`
	Reason                       *string                            `json:"reason,omitempty"`
	Status                       *CapabilityStatus                  `json:"status,omitempty"`
	SupportedFamilies            *[]ManagedInstanceFamilyCapability `json:"supportedFamilies,omitempty"`
	SupportedStorageCapabilities *[]StorageCapability               `json:"supportedStorageCapabilities,omitempty"`
}
