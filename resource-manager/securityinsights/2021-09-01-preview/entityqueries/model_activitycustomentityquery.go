package entityqueries

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomEntityQuery = ActivityCustomEntityQuery{}

type ActivityCustomEntityQuery struct {
	Properties *ActivityEntityQueriesProperties `json:"properties"`

	// Fields inherited from CustomEntityQuery
	Etag       *string                `json:"etag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = ActivityCustomEntityQuery{}

func (s ActivityCustomEntityQuery) MarshalJSON() ([]byte, error) {
	type wrapper ActivityCustomEntityQuery
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ActivityCustomEntityQuery: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ActivityCustomEntityQuery: %+v", err)
	}
	decoded["kind"] = "Activity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ActivityCustomEntityQuery: %+v", err)
	}

	return encoded, nil
}
