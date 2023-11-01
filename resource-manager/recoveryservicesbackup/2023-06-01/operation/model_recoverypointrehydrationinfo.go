package operation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointRehydrationInfo struct {
	RehydrationPriority          *RehydrationPriority `json:"rehydrationPriority,omitempty"`
	RehydrationRetentionDuration *string              `json:"rehydrationRetentionDuration,omitempty"`
}
