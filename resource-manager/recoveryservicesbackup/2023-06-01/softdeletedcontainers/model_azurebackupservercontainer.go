package softdeletedcontainers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionContainer = AzureBackupServerContainer{}

type AzureBackupServerContainer struct {
	CanReRegister      *bool                     `json:"canReRegister,omitempty"`
	ContainerId        *string                   `json:"containerId,omitempty"`
	DpmAgentVersion    *string                   `json:"dpmAgentVersion,omitempty"`
	DpmServers         *[]string                 `json:"dpmServers,omitempty"`
	ExtendedInfo       *DPMContainerExtendedInfo `json:"extendedInfo,omitempty"`
	ProtectedItemCount *int64                    `json:"protectedItemCount,omitempty"`
	ProtectionStatus   *string                   `json:"protectionStatus,omitempty"`
	UpgradeAvailable   *bool                     `json:"upgradeAvailable,omitempty"`

	// Fields inherited from ProtectionContainer
	BackupManagementType  *BackupManagementType `json:"backupManagementType,omitempty"`
	FriendlyName          *string               `json:"friendlyName,omitempty"`
	HealthStatus          *string               `json:"healthStatus,omitempty"`
	ProtectableObjectType *string               `json:"protectableObjectType,omitempty"`
	RegistrationStatus    *string               `json:"registrationStatus,omitempty"`
}

var _ json.Marshaler = AzureBackupServerContainer{}

func (s AzureBackupServerContainer) MarshalJSON() ([]byte, error) {
	type wrapper AzureBackupServerContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureBackupServerContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureBackupServerContainer: %+v", err)
	}
	decoded["containerType"] = "AzureBackupServerContainer"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureBackupServerContainer: %+v", err)
	}

	return encoded, nil
}
