package aad

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// ExternalAccess enumerates the values for external access.
type ExternalAccess string

const (
	// Disabled specifies the disabled state for external access.
	Disabled ExternalAccess = "Disabled"
	// Enabled specifies the enabled state for external access.
	Enabled ExternalAccess = "Enabled"
)

// Ldaps enumerates the values for ldaps.
type Ldaps string

const (
	// LdapsDisabled specifies the ldaps disabled state for ldaps.
	LdapsDisabled Ldaps = "Disabled"
	// LdapsEnabled specifies the ldaps enabled state for ldaps.
	LdapsEnabled Ldaps = "Enabled"
)

// DomainService is domain service.
type DomainService struct {
	autorest.Response        `json:"-"`
	ID                       *string             `json:"id,omitempty"`
	Name                     *string             `json:"name,omitempty"`
	Type                     *string             `json:"type,omitempty"`
	Location                 *string             `json:"location,omitempty"`
	Tags                     *map[string]*string `json:"tags,omitempty"`
	Etag                     *string             `json:"etag,omitempty"`
	*DomainServiceProperties `json:"properties,omitempty"`
}

// DomainServiceListResult is the response from the List Domain Services operation.
type DomainServiceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]DomainService `json:"value,omitempty"`
}

// DomainServicePatchProperties is update Properties of the Domain Service.
type DomainServicePatchProperties struct {
	LdapsSettings *LdapsSettings `json:"ldapsSettings,omitempty"`
}

// DomainServiceProperties is properties of the Domain Service.
type DomainServiceProperties struct {
	TenantID                  *string        `json:"tenantId,omitempty"`
	DomainName                *string        `json:"domainName,omitempty"`
	VnetSiteID                *string        `json:"vnetSiteId,omitempty"`
	SubnetID                  *string        `json:"subnetId,omitempty"`
	LdapsSettings             *LdapsSettings `json:"ldapsSettings,omitempty"`
	DomainControllerIPAddress *[]string      `json:"domainControllerIpAddress,omitempty"`
	ServiceStatus             *string        `json:"serviceStatus,omitempty"`
	ProvisioningState         *string        `json:"provisioningState,omitempty"`
}

// LdapsSettings is secure LDAP Settings
type LdapsSettings struct {
	Ldaps                   Ldaps          `json:"ldaps,omitempty"`
	PfxCertificate          *string        `json:"pfxCertificate,omitempty"`
	PfxCertificatePassword  *string        `json:"pfxCertificatePassword,omitempty"`
	PublicCertificate       *string        `json:"publicCertificate,omitempty"`
	CertificateThumbprint   *string        `json:"certificateThumbprint,omitempty"`
	CertificateNotAfter     *date.Time     `json:"certificateNotAfter,omitempty"`
	ExternalAccess          ExternalAccess `json:"externalAccess,omitempty"`
	ExternalAccessIPAddress *string        `json:"externalAccessIpAddress,omitempty"`
}

// OperationDisplayInfo is the operation supported by Domain Services.
type OperationDisplayInfo struct {
	Description *string `json:"description,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
}

// OperationEntity is the operation supported by Domain Services.
type OperationEntity struct {
	Name    *string               `json:"name,omitempty"`
	Display *OperationDisplayInfo `json:"display,omitempty"`
	Origin  *string               `json:"origin,omitempty"`
}

// OperationEntityListResult is the list of domain service operation response.
type OperationEntityListResult struct {
	autorest.Response `json:"-"`
	Value             *[]OperationEntity `json:"value,omitempty"`
}

// Resource is the Resource model definition.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
	Etag     *string             `json:"etag,omitempty"`
}
