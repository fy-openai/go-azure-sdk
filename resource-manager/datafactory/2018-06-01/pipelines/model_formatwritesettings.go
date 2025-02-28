package pipelines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FormatWriteSettings interface {
}

// RawFormatWriteSettingsImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawFormatWriteSettingsImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalFormatWriteSettingsImplementation(input []byte) (FormatWriteSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling FormatWriteSettings into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AvroWriteSettings") {
		var out AvroWriteSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AvroWriteSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DelimitedTextWriteSettings") {
		var out DelimitedTextWriteSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelimitedTextWriteSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IcebergWriteSettings") {
		var out IcebergWriteSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IcebergWriteSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "JsonWriteSettings") {
		var out JsonWriteSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JsonWriteSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OrcWriteSettings") {
		var out OrcWriteSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OrcWriteSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ParquetWriteSettings") {
		var out ParquetWriteSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ParquetWriteSettings: %+v", err)
		}
		return out, nil
	}

	out := RawFormatWriteSettingsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
