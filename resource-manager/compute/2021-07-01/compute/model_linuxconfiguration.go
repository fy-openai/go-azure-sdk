package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinuxConfiguration struct {
	DisablePasswordAuthentication *bool               `json:"disablePasswordAuthentication,omitempty"`
	PatchSettings                 *LinuxPatchSettings `json:"patchSettings"`
	ProvisionVMAgent              *bool               `json:"provisionVMAgent,omitempty"`
	Ssh                           *SshConfiguration   `json:"ssh"`
}
