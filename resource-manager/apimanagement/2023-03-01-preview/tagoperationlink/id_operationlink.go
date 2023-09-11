package tagoperationlink

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = OperationLinkId{}

// OperationLinkId is a struct representing the Resource ID for a Operation Link
type OperationLinkId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	TagId             string
	OperationLinkId   string
}

// NewOperationLinkID returns a new OperationLinkId struct
func NewOperationLinkID(subscriptionId string, resourceGroupName string, serviceName string, tagId string, operationLinkId string) OperationLinkId {
	return OperationLinkId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		TagId:             tagId,
		OperationLinkId:   operationLinkId,
	}
}

// ParseOperationLinkID parses 'input' into a OperationLinkId
func ParseOperationLinkID(input string) (*OperationLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(OperationLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OperationLinkId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.TagId, ok = parsed.Parsed["tagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tagId", *parsed)
	}

	if id.OperationLinkId, ok = parsed.Parsed["operationLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationLinkId", *parsed)
	}

	return &id, nil
}

// ParseOperationLinkIDInsensitively parses 'input' case-insensitively into a OperationLinkId
// note: this method should only be used for API response data and not user input
func ParseOperationLinkIDInsensitively(input string) (*OperationLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(OperationLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OperationLinkId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.TagId, ok = parsed.Parsed["tagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tagId", *parsed)
	}

	if id.OperationLinkId, ok = parsed.Parsed["operationLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationLinkId", *parsed)
	}

	return &id, nil
}

// ValidateOperationLinkID checks that 'input' can be parsed as a Operation Link ID
func ValidateOperationLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOperationLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Operation Link ID
func (id OperationLinkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/tags/%s/operationLinks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.TagId, id.OperationLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Operation Link ID
func (id OperationLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticTags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("tagId", "tagIdValue"),
		resourceids.StaticSegment("staticOperationLinks", "operationLinks", "operationLinks"),
		resourceids.UserSpecifiedSegment("operationLinkId", "operationLinkIdValue"),
	}
}

// String returns a human-readable description of this Operation Link ID
func (id OperationLinkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Tag: %q", id.TagId),
		fmt.Sprintf("Operation Link: %q", id.OperationLinkId),
	}
	return fmt.Sprintf("Operation Link (%s)", strings.Join(components, "\n"))
}
