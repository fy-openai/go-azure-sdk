package outputs

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMode string

const (
	AuthenticationModeConnectionString AuthenticationMode = "ConnectionString"
	AuthenticationModeMsi              AuthenticationMode = "Msi"
	AuthenticationModeUserToken        AuthenticationMode = "UserToken"
)

func PossibleValuesForAuthenticationMode() []string {
	return []string{
		string(AuthenticationModeConnectionString),
		string(AuthenticationModeMsi),
		string(AuthenticationModeUserToken),
	}
}

func parseAuthenticationMode(input string) (*AuthenticationMode, error) {
	vals := map[string]AuthenticationMode{
		"connectionstring": AuthenticationModeConnectionString,
		"msi":              AuthenticationModeMsi,
		"usertoken":        AuthenticationModeUserToken,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMode(input)
	return &out, nil
}

type BlobWriteMode string

const (
	BlobWriteModeAppend BlobWriteMode = "Append"
	BlobWriteModeOnce   BlobWriteMode = "Once"
)

func PossibleValuesForBlobWriteMode() []string {
	return []string{
		string(BlobWriteModeAppend),
		string(BlobWriteModeOnce),
	}
}

func parseBlobWriteMode(input string) (*BlobWriteMode, error) {
	vals := map[string]BlobWriteMode{
		"append": BlobWriteModeAppend,
		"once":   BlobWriteModeOnce,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BlobWriteMode(input)
	return &out, nil
}

type Encoding string

const (
	EncodingUTFEight Encoding = "UTF8"
)

func PossibleValuesForEncoding() []string {
	return []string{
		string(EncodingUTFEight),
	}
}

func parseEncoding(input string) (*Encoding, error) {
	vals := map[string]Encoding{
		"utf8": EncodingUTFEight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Encoding(input)
	return &out, nil
}

type EventSerializationType string

const (
	EventSerializationTypeAvro      EventSerializationType = "Avro"
	EventSerializationTypeCsv       EventSerializationType = "Csv"
	EventSerializationTypeCustomClr EventSerializationType = "CustomClr"
	EventSerializationTypeJson      EventSerializationType = "Json"
	EventSerializationTypeParquet   EventSerializationType = "Parquet"
)

func PossibleValuesForEventSerializationType() []string {
	return []string{
		string(EventSerializationTypeAvro),
		string(EventSerializationTypeCsv),
		string(EventSerializationTypeCustomClr),
		string(EventSerializationTypeJson),
		string(EventSerializationTypeParquet),
	}
}

func parseEventSerializationType(input string) (*EventSerializationType, error) {
	vals := map[string]EventSerializationType{
		"avro":      EventSerializationTypeAvro,
		"csv":       EventSerializationTypeCsv,
		"customclr": EventSerializationTypeCustomClr,
		"json":      EventSerializationTypeJson,
		"parquet":   EventSerializationTypeParquet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventSerializationType(input)
	return &out, nil
}

type JsonOutputSerializationFormat string

const (
	JsonOutputSerializationFormatArray         JsonOutputSerializationFormat = "Array"
	JsonOutputSerializationFormatLineSeparated JsonOutputSerializationFormat = "LineSeparated"
)

func PossibleValuesForJsonOutputSerializationFormat() []string {
	return []string{
		string(JsonOutputSerializationFormatArray),
		string(JsonOutputSerializationFormatLineSeparated),
	}
}

func parseJsonOutputSerializationFormat(input string) (*JsonOutputSerializationFormat, error) {
	vals := map[string]JsonOutputSerializationFormat{
		"array":         JsonOutputSerializationFormatArray,
		"lineseparated": JsonOutputSerializationFormatLineSeparated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JsonOutputSerializationFormat(input)
	return &out, nil
}

type OutputWatermarkMode string

const (
	OutputWatermarkModeNone                                OutputWatermarkMode = "None"
	OutputWatermarkModeSendCurrentPartitionWatermark       OutputWatermarkMode = "SendCurrentPartitionWatermark"
	OutputWatermarkModeSendLowestWatermarkAcrossPartitions OutputWatermarkMode = "SendLowestWatermarkAcrossPartitions"
)

func PossibleValuesForOutputWatermarkMode() []string {
	return []string{
		string(OutputWatermarkModeNone),
		string(OutputWatermarkModeSendCurrentPartitionWatermark),
		string(OutputWatermarkModeSendLowestWatermarkAcrossPartitions),
	}
}

func parseOutputWatermarkMode(input string) (*OutputWatermarkMode, error) {
	vals := map[string]OutputWatermarkMode{
		"none":                                OutputWatermarkModeNone,
		"sendcurrentpartitionwatermark":       OutputWatermarkModeSendCurrentPartitionWatermark,
		"sendlowestwatermarkacrosspartitions": OutputWatermarkModeSendLowestWatermarkAcrossPartitions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputWatermarkMode(input)
	return &out, nil
}
