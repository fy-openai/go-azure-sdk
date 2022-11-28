package dataconnectors

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnector = ASCDataConnector{}

type ASCDataConnector struct {
	Properties *ASCDataConnectorProperties `json:"properties"`

	// Fields inherited from DataConnector
	Etag       *string                `json:"etag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = ASCDataConnector{}

func (s ASCDataConnector) MarshalJSON() ([]byte, error) {
	type wrapper ASCDataConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ASCDataConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ASCDataConnector: %+v", err)
	}
	decoded["kind"] = "AzureSecurityCenter"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ASCDataConnector: %+v", err)
	}

	return encoded, nil
}
