package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetUpdateNetworkConfigurationProperties struct {
	DeleteOption                *DeleteOptions                                         `json:"deleteOption,omitempty"`
	DnsSettings                 *VirtualMachineScaleSetNetworkConfigurationDnsSettings `json:"dnsSettings"`
	EnableAcceleratedNetworking *bool                                                  `json:"enableAcceleratedNetworking,omitempty"`
	EnableFpga                  *bool                                                  `json:"enableFpga,omitempty"`
	EnableIPForwarding          *bool                                                  `json:"enableIPForwarding,omitempty"`
	IPConfigurations            *[]VirtualMachineScaleSetUpdateIPConfiguration         `json:"ipConfigurations,omitempty"`
	NetworkSecurityGroup        *SubResource                                           `json:"networkSecurityGroup"`
	Primary                     *bool                                                  `json:"primary,omitempty"`
}
