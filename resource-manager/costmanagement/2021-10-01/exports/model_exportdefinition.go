package exports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportDefinition struct {
	DataSet    *ExportDataset    `json:"dataSet"`
	TimePeriod *ExportTimePeriod `json:"timePeriod"`
	Timeframe  TimeframeType     `json:"timeframe"`
	Type       ExportType        `json:"type"`
}
