package apidiagnostic

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApiDiagnosticId{}

// ApiDiagnosticId is a struct representing the Resource ID for a Api Diagnostic
type ApiDiagnosticId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ApiId             string
	DiagnosticId      string
}

// NewApiDiagnosticID returns a new ApiDiagnosticId struct
func NewApiDiagnosticID(subscriptionId string, resourceGroupName string, serviceName string, apiId string, diagnosticId string) ApiDiagnosticId {
	return ApiDiagnosticId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ApiId:             apiId,
		DiagnosticId:      diagnosticId,
	}
}

// ParseApiDiagnosticID parses 'input' into a ApiDiagnosticId
func ParseApiDiagnosticID(input string) (*ApiDiagnosticId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApiDiagnosticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApiDiagnosticId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.ApiId, ok = parsed.Parsed["apiId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "apiId", *parsed)
	}

	if id.DiagnosticId, ok = parsed.Parsed["diagnosticId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "diagnosticId", *parsed)
	}

	return &id, nil
}

// ParseApiDiagnosticIDInsensitively parses 'input' case-insensitively into a ApiDiagnosticId
// note: this method should only be used for API response data and not user input
func ParseApiDiagnosticIDInsensitively(input string) (*ApiDiagnosticId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApiDiagnosticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApiDiagnosticId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.ApiId, ok = parsed.Parsed["apiId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "apiId", *parsed)
	}

	if id.DiagnosticId, ok = parsed.Parsed["diagnosticId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "diagnosticId", *parsed)
	}

	return &id, nil
}

// ValidateApiDiagnosticID checks that 'input' can be parsed as a Api Diagnostic ID
func ValidateApiDiagnosticID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApiDiagnosticID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Api Diagnostic ID
func (id ApiDiagnosticId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/apis/%s/diagnostics/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ApiId, id.DiagnosticId)
}

// Segments returns a slice of Resource ID Segments which comprise this Api Diagnostic ID
func (id ApiDiagnosticId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticApis", "apis", "apis"),
		resourceids.UserSpecifiedSegment("apiId", "apiIdValue"),
		resourceids.StaticSegment("staticDiagnostics", "diagnostics", "diagnostics"),
		resourceids.UserSpecifiedSegment("diagnosticId", "diagnosticIdValue"),
	}
}

// String returns a human-readable description of this Api Diagnostic ID
func (id ApiDiagnosticId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Api: %q", id.ApiId),
		fmt.Sprintf("Diagnostic: %q", id.DiagnosticId),
	}
	return fmt.Sprintf("Api Diagnostic (%s)", strings.Join(components, "\n"))
}
