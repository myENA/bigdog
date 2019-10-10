package aaaservers

// API Version: v8_1

import (
	"encoding/json"
)

type AAAServer struct {
	// AcctPort
	// AAA server accounting port
	AcctPort *int `json:"acctPort,omitempty"`

	// AuthPort
	// AAA server authentication port
	AuthPort *int `json:"authPort,omitempty"`

	// CreatedTime
	// The create time of the AAA server
	CreatedTime *int `json:"createdTime,omitempty"`

	// CreatorID
	// AAA server creator Id
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// AAA server creator name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// ID
	// AAA server ID
	ID *string `json:"id,omitempty"`

	// IP
	// AAA server IP address
	IP *string `json:"ip,omitempty"`

	// Level
	// Access level of AAA server
	Level *string `json:"level,omitempty" validate:"oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Name
	// Name of the AAA server
	Name *string `json:"name,omitempty"`

	// Password
	// Password for local user
	Password *string `json:"password,omitempty"`

	// Purpose
	// AAA server purpose
	Purpose *string `json:"purpose,omitempty" validate:"oneof=DEFAULT AUTHENTICATION_ONLY AUTHORIZATION_ONLY ACCOUNTING_ONLY"`

	// Secret
	// AAA server secret
	Secret *string `json:"secret,omitempty"`

	// ServerType
	// The server tpye of the AAA server
	ServerType *string `json:"serverType,omitempty" validate:"oneof=RADIUS TACACS_PLUS LOCAL"`

	// UpdatedTime
	// The modify time of the AAA server
	UpdatedTime *int `json:"updatedTime,omitempty"`

	// UpdaterID
	// AAA server updater Id
	UpdaterID *string `json:"updaterId,omitempty"`

	// UpdaterUsername
	// AAA server updater name
	UpdaterUsername *string `json:"updaterUsername,omitempty"`

	// Username
	// Username for local user
	Username *string `json:"username,omitempty"`
}

type AaaServersQueryResult struct {
	// Extra
	// Any additional response data
	Extra *AaaServersQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first AAA Server returned out of the complete AAA Server list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more AAA Servers after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*AAAServer `json:"list,omitempty"`

	// RawDataTotalCount
	// Total AAA Servers count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total AAA Servers count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

// AaaServersQueryResultExtraType
//
// Any additional response data
type AaaServersQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *AaaServersQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = AaaServersQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *AaaServersQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type CreateAdminAAAServer struct {
	// AcctPort
	// AAA server accounting port
	AcctPort *int `json:"acctPort,omitempty"`

	// AuthPort
	// AAA server authentication port
	AuthPort *int `json:"authPort,omitempty"`

	// IP
	// AAA server IP address
	IP *string `json:"ip,omitempty"`

	// Level
	// Access level of AAA server
	Level *string `json:"level,omitempty" validate:"oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Name
	// Name of the AAA server
	Name *string `json:"name,omitempty"`

	// Password
	// Password for local user
	Password *string `json:"password,omitempty"`

	// Purpose
	// AAA server purpose
	Purpose *string `json:"purpose,omitempty" validate:"oneof=DEFAULT AUTHENTICATION_ONLY AUTHORIZATION_ONLY ACCOUNTING_ONLY"`

	// Secret
	// AAA server secret
	Secret *string `json:"secret,omitempty"`

	// ServerType
	// The server tpye of the AAA server
	ServerType *string `json:"serverType,omitempty" validate:"oneof=RADIUS TACACS_PLUS LOCAL"`

	// Username
	// Username for local user
	Username *string `json:"username,omitempty"`
}

type EmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *EmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = EmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *EmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}
