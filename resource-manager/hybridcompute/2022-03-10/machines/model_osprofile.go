package machines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OSProfile struct {
	ComputerName         *string                        `json:"computerName,omitempty"`
	LinuxConfiguration   *OSProfileLinuxConfiguration   `json:"linuxConfiguration"`
	WindowsConfiguration *OSProfileWindowsConfiguration `json:"windowsConfiguration"`
}
