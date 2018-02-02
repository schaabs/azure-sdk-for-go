package compute

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
	"encoding/xml"
	"reflect"
	"time"
	"unsafe"
)

const (
	rfc3339Format = "2006-01-02T15:04:05.0000000Z07:00"
)

// used to convert times from UTC to GMT before sending across the wire
var gmt = time.FixedZone("GMT", 0)

// internal type used for marshalling time in RFC1123 format
type timeRFC1123 struct {
	time.Time
}

// MarshalText implements the encoding.TextMarshaler interface for timeRFC1123.
func (t timeRFC1123) MarshalText() ([]byte, error) {
	return []byte(t.Format(time.RFC1123)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for timeRFC1123.
func (t *timeRFC1123) UnmarshalText(data []byte) (err error) {
	t.Time, err = time.Parse(time.RFC1123, string(data))
	return
}

// internal type used for marshalling time in RFC3339 format
type timeRFC3339 struct {
	time.Time
}

// MarshalText implements the encoding.TextMarshaler interface for timeRFC3339.
func (t timeRFC3339) MarshalText() ([]byte, error) {
	return []byte(t.Format(rfc3339Format)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for timeRFC3339.
func (t *timeRFC3339) UnmarshalText(data []byte) (err error) {
	t.Time, err = time.Parse(rfc3339Format, string(data))
	return
}

// internal type used for marshalling
type instanceViewStatus struct {
	Code          *string              `json:"code,omitempty"`
	Level         StatusLevelTypesType `json:"level,omitempty"`
	DisplayStatus *string              `json:"displayStatus,omitempty"`
	Message       *string              `json:"message,omitempty"`
	Time          *timeRFC3339         `json:"time,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for InstanceViewStatus.
func (ivs InstanceViewStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*InstanceViewStatus)(nil)).Elem().Size() != reflect.TypeOf((*instanceViewStatus)(nil)).Elem().Size() {
		panic("size mismatch between InstanceViewStatus and instanceViewStatus")
	}
	ivs2 := (*instanceViewStatus)(unsafe.Pointer(&ivs))
	return e.EncodeElement(*ivs2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for InstanceViewStatus.
func (ivs *InstanceViewStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*InstanceViewStatus)(nil)).Elem().Size() != reflect.TypeOf((*instanceViewStatus)(nil)).Elem().Size() {
		panic("size mismatch between InstanceViewStatus and instanceViewStatus")
	}
	ivs2 := (*instanceViewStatus)(unsafe.Pointer(ivs))
	return d.DecodeElement(ivs2, &start)
}

// internal type used for marshalling
type maintenanceRedeployStatus struct {
	IsCustomerInitiatedMaintenanceAllowed *bool                                   `json:"isCustomerInitiatedMaintenanceAllowed,omitempty"`
	PreMaintenanceWindowStartTime         *timeRFC3339                            `json:"preMaintenanceWindowStartTime,omitempty"`
	PreMaintenanceWindowEndTime           *timeRFC3339                            `json:"preMaintenanceWindowEndTime,omitempty"`
	MaintenanceWindowStartTime            *timeRFC3339                            `json:"maintenanceWindowStartTime,omitempty"`
	MaintenanceWindowEndTime              *timeRFC3339                            `json:"maintenanceWindowEndTime,omitempty"`
	LastOperationResultCode               MaintenanceOperationResultCodeTypesType `json:"lastOperationResultCode,omitempty"`
	LastOperationMessage                  *string                                 `json:"lastOperationMessage,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for MaintenanceRedeployStatus.
func (mrs MaintenanceRedeployStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*MaintenanceRedeployStatus)(nil)).Elem().Size() != reflect.TypeOf((*maintenanceRedeployStatus)(nil)).Elem().Size() {
		panic("size mismatch between MaintenanceRedeployStatus and maintenanceRedeployStatus")
	}
	mrs2 := (*maintenanceRedeployStatus)(unsafe.Pointer(&mrs))
	return e.EncodeElement(*mrs2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for MaintenanceRedeployStatus.
func (mrs *MaintenanceRedeployStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*MaintenanceRedeployStatus)(nil)).Elem().Size() != reflect.TypeOf((*maintenanceRedeployStatus)(nil)).Elem().Size() {
		panic("size mismatch between MaintenanceRedeployStatus and maintenanceRedeployStatus")
	}
	mrs2 := (*maintenanceRedeployStatus)(unsafe.Pointer(mrs))
	return d.DecodeElement(mrs2, &start)
}

// internal type used for marshalling
type rollingUpgradeRunningStatus struct {
	Code           RollingUpgradeStatusCodeType `json:"code,omitempty"`
	StartTime      *timeRFC3339                 `json:"startTime,omitempty"`
	LastAction     RollingUpgradeActionType     `json:"lastAction,omitempty"`
	LastActionTime *timeRFC3339                 `json:"lastActionTime,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for RollingUpgradeRunningStatus.
func (rurs RollingUpgradeRunningStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*RollingUpgradeRunningStatus)(nil)).Elem().Size() != reflect.TypeOf((*rollingUpgradeRunningStatus)(nil)).Elem().Size() {
		panic("size mismatch between RollingUpgradeRunningStatus and rollingUpgradeRunningStatus")
	}
	rurs2 := (*rollingUpgradeRunningStatus)(unsafe.Pointer(&rurs))
	return e.EncodeElement(*rurs2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for RollingUpgradeRunningStatus.
func (rurs *RollingUpgradeRunningStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*RollingUpgradeRunningStatus)(nil)).Elem().Size() != reflect.TypeOf((*rollingUpgradeRunningStatus)(nil)).Elem().Size() {
		panic("size mismatch between RollingUpgradeRunningStatus and rollingUpgradeRunningStatus")
	}
	rurs2 := (*rollingUpgradeRunningStatus)(unsafe.Pointer(rurs))
	return d.DecodeElement(rurs2, &start)
}

// internal type used for marshalling
type operationStatusResponse struct {
	Name      *string      `json:"name,omitempty"`
	Status    *string      `json:"status,omitempty"`
	StartTime *timeRFC3339 `json:"startTime,omitempty"`
	EndTime   *timeRFC3339 `json:"endTime,omitempty"`
	Error     *APIError    `json:"error,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for OperationStatusResponse.
func (osr OperationStatusResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*OperationStatusResponse)(nil)).Elem().Size() != reflect.TypeOf((*operationStatusResponse)(nil)).Elem().Size() {
		panic("size mismatch between OperationStatusResponse and operationStatusResponse")
	}
	osr2 := (*operationStatusResponse)(unsafe.Pointer(&osr))
	return e.EncodeElement(*osr2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for OperationStatusResponse.
func (osr *OperationStatusResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*OperationStatusResponse)(nil)).Elem().Size() != reflect.TypeOf((*operationStatusResponse)(nil)).Elem().Size() {
		panic("size mismatch between OperationStatusResponse and operationStatusResponse")
	}
	osr2 := (*operationStatusResponse)(unsafe.Pointer(osr))
	return d.DecodeElement(osr2, &start)
}

// internal type used for marshalling
type diskProperties struct {
	TimeCreated        *timeRFC3339             `json:"timeCreated,omitempty"`
	OsType             OperatingSystemTypesType `json:"osType,omitempty"`
	CreationData       CreationData             `json:"creationData,omitempty"`
	DiskSizeGB         *int32                   `json:"diskSizeGB,omitempty"`
	EncryptionSettings *EncryptionSettings      `json:"encryptionSettings,omitempty"`
	ProvisioningState  *string                  `json:"provisioningState,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for DiskProperties.
func (dp DiskProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*DiskProperties)(nil)).Elem().Size() != reflect.TypeOf((*diskProperties)(nil)).Elem().Size() {
		panic("size mismatch between DiskProperties and diskProperties")
	}
	dp2 := (*diskProperties)(unsafe.Pointer(&dp))
	return e.EncodeElement(*dp2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for DiskProperties.
func (dp *DiskProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*DiskProperties)(nil)).Elem().Size() != reflect.TypeOf((*diskProperties)(nil)).Elem().Size() {
		panic("size mismatch between DiskProperties and diskProperties")
	}
	dp2 := (*diskProperties)(unsafe.Pointer(dp))
	return d.DecodeElement(dp2, &start)
}
