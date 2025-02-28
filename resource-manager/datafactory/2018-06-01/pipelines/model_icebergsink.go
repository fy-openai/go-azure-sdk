package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CopySink = IcebergSink{}

type IcebergSink struct {
	FormatSettings FormatWriteSettings `json:"formatSettings"`
	StoreSettings  StoreWriteSettings  `json:"storeSettings"`

	// Fields inherited from CopySink
	DisableMetricsCollection *bool   `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *int64  `json:"maxConcurrentConnections,omitempty"`
	SinkRetryCount           *int64  `json:"sinkRetryCount,omitempty"`
	SinkRetryWait            *string `json:"sinkRetryWait,omitempty"`
	WriteBatchSize           *int64  `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout        *string `json:"writeBatchTimeout,omitempty"`
}

var _ json.Marshaler = IcebergSink{}

func (s IcebergSink) MarshalJSON() ([]byte, error) {
	type wrapper IcebergSink
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IcebergSink: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IcebergSink: %+v", err)
	}
	decoded["type"] = "IcebergSink"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IcebergSink: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IcebergSink{}

func (s *IcebergSink) UnmarshalJSON(bytes []byte) error {
	type alias IcebergSink
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into IcebergSink: %+v", err)
	}

	s.DisableMetricsCollection = decoded.DisableMetricsCollection
	s.MaxConcurrentConnections = decoded.MaxConcurrentConnections
	s.SinkRetryCount = decoded.SinkRetryCount
	s.SinkRetryWait = decoded.SinkRetryWait
	s.WriteBatchSize = decoded.WriteBatchSize
	s.WriteBatchTimeout = decoded.WriteBatchTimeout

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IcebergSink into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["formatSettings"]; ok {
		impl, err := unmarshalFormatWriteSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'FormatSettings' for 'IcebergSink': %+v", err)
		}
		s.FormatSettings = impl
	}

	if v, ok := temp["storeSettings"]; ok {
		impl, err := unmarshalStoreWriteSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StoreSettings' for 'IcebergSink': %+v", err)
		}
		s.StoreSettings = impl
	}
	return nil
}
