package apiversionset

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkspaceApiVersionSetId{}

// WorkspaceApiVersionSetId is a struct representing the Resource ID for a Workspace Api Version Set
type WorkspaceApiVersionSetId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	WorkspaceId       string
	VersionSetId      string
}

// NewWorkspaceApiVersionSetID returns a new WorkspaceApiVersionSetId struct
func NewWorkspaceApiVersionSetID(subscriptionId string, resourceGroupName string, serviceName string, workspaceId string, versionSetId string) WorkspaceApiVersionSetId {
	return WorkspaceApiVersionSetId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		WorkspaceId:       workspaceId,
		VersionSetId:      versionSetId,
	}
}

// ParseWorkspaceApiVersionSetID parses 'input' into a WorkspaceApiVersionSetId
func ParseWorkspaceApiVersionSetID(input string) (*WorkspaceApiVersionSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspaceApiVersionSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspaceApiVersionSetId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.WorkspaceId, ok = parsed.Parsed["workspaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceId", *parsed)
	}

	if id.VersionSetId, ok = parsed.Parsed["versionSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionSetId", *parsed)
	}

	return &id, nil
}

// ParseWorkspaceApiVersionSetIDInsensitively parses 'input' case-insensitively into a WorkspaceApiVersionSetId
// note: this method should only be used for API response data and not user input
func ParseWorkspaceApiVersionSetIDInsensitively(input string) (*WorkspaceApiVersionSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspaceApiVersionSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspaceApiVersionSetId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.WorkspaceId, ok = parsed.Parsed["workspaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceId", *parsed)
	}

	if id.VersionSetId, ok = parsed.Parsed["versionSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionSetId", *parsed)
	}

	return &id, nil
}

// ValidateWorkspaceApiVersionSetID checks that 'input' can be parsed as a Workspace Api Version Set ID
func ValidateWorkspaceApiVersionSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspaceApiVersionSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Api Version Set ID
func (id WorkspaceApiVersionSetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/workspaces/%s/apiVersionSets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.WorkspaceId, id.VersionSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Api Version Set ID
func (id WorkspaceApiVersionSetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceId", "workspaceIdValue"),
		resourceids.StaticSegment("staticApiVersionSets", "apiVersionSets", "apiVersionSets"),
		resourceids.UserSpecifiedSegment("versionSetId", "versionSetIdValue"),
	}
}

// String returns a human-readable description of this Workspace Api Version Set ID
func (id WorkspaceApiVersionSetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Workspace: %q", id.WorkspaceId),
		fmt.Sprintf("Version Set: %q", id.VersionSetId),
	}
	return fmt.Sprintf("Workspace Api Version Set (%s)", strings.Join(components, "\n"))
}
