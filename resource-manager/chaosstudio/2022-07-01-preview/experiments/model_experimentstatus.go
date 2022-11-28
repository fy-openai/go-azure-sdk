package experiments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExperimentStatus struct {
	Id         *string                     `json:"id,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Properties *ExperimentStatusProperties `json:"properties"`
	Type       *string                     `json:"type,omitempty"`
}
