package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"io"
)

// WSGEventManagementEventDataList
//
// Definition: eventManagement_eventDataList
type WSGEventManagementEventDataList struct {
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGEventManagementSingleEventSetting `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGEventManagementEventDataList() *WSGEventManagementEventDataList {
	m := new(WSGEventManagementEventDataList)
	return m
}

// WSGEventManagementEventDataResponse
//
// Definition: eventManagement_eventDataResponse
type WSGEventManagementEventDataResponse struct {
	Data *WSGEventManagementEventDataList `json:"data,omitempty"`

	// Error
	// The error message of http request
	Error *string `json:"error,omitempty"`

	// Extra
	// Extra information for event management setting
	Extra *string `json:"extra,omitempty"`

	// Success
	// The status of http request
	Success *bool `json:"success,omitempty"`
}

type WSGEventManagementEventDataResponseAPIResponse struct {
	*RawAPIResponse
	Data *WSGEventManagementEventDataResponse
}

func newWSGEventManagementEventDataResponseAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGEventManagementEventDataResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGEventManagementEventDataResponseAPIResponse) Hydrate() error {
	r.Data = new(WSGEventManagementEventDataResponse)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGEventManagementEventDataResponse() *WSGEventManagementEventDataResponse {
	m := new(WSGEventManagementEventDataResponse)
	return m
}

// WSGEventManagementEventEmailSetting
//
// Definition: eventManagement_eventEmailSetting
type WSGEventManagementEventEmailSetting struct {
	// EmailEnabled
	// Enable/Disable Email sending function
	EmailEnabled *bool `json:"emailEnabled,omitempty"`

	// MailTo
	// E-mail recipients
	MailTo *string `json:"mailTo,omitempty"`
}

type WSGEventManagementEventEmailSettingAPIResponse struct {
	*RawAPIResponse
	Data *WSGEventManagementEventEmailSetting
}

func newWSGEventManagementEventEmailSettingAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGEventManagementEventEmailSettingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGEventManagementEventEmailSettingAPIResponse) Hydrate() error {
	r.Data = new(WSGEventManagementEventEmailSetting)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGEventManagementEventEmailSetting() *WSGEventManagementEventEmailSetting {
	m := new(WSGEventManagementEventEmailSetting)
	return m
}

// WSGEventManagementEventSettingList
//
// Definition: eventManagement_eventSettingList
type WSGEventManagementEventSettingList []*WSGEventManagementSingleEventSetting

func MakeWSGEventManagementEventSettingList() WSGEventManagementEventSettingList {
	m := make(WSGEventManagementEventSettingList, 0)
	return m
}

// WSGEventManagementSingleEventSetting
//
// Definition: eventManagement_singleEventSetting
type WSGEventManagementSingleEventSetting struct {
	// Category
	// Event category
	Category *string `json:"category,omitempty"`

	// ConfigPageDesc
	// Event description
	ConfigPageDesc *string `json:"configPageDesc,omitempty"`

	// DbPersistence
	// Enable/Disable DB persistence for this event
	DbPersistence *bool `json:"dbPersistence,omitempty"`

	// EventCode
	// Event code
	EventCode *int `json:"eventCode,omitempty"`

	// Oid
	// Event OID
	Oid *string `json:"oid,omitempty"`

	// Severity
	// Event severity
	Severity *string `json:"severity,omitempty"`

	// TriggerEmail
	// Enable/Disable Email sending for this event
	TriggerEmail *bool `json:"triggerEmail,omitempty"`

	// TriggerTrap
	// Enable/Disable SNMP function for this event
	TriggerTrap *bool `json:"triggerTrap,omitempty"`

	// Type
	// Event type
	Type *string `json:"type,omitempty"`

	// ZoneOverride
	// Enable/Disable override event system settings
	ZoneOverride *bool `json:"zoneOverride,omitempty"`
}

func NewWSGEventManagementSingleEventSetting() *WSGEventManagementSingleEventSetting {
	m := new(WSGEventManagementSingleEventSetting)
	return m
}
