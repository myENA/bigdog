package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type SwitchMAAAServersService struct {
	apiClient *APIClient
}

func NewSwitchMAAAServersService(c *APIClient) *SwitchMAAAServersService {
	s := new(SwitchMAAAServersService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMAAAServersService() *SwitchMAAAServersService {
	return NewSwitchMAAAServersService(ss.apiClient)
}

type SwitchMAAAServersAAAServer struct {
	// AcctPort
	// AAA server accounting port
	AcctPort *int `json:"acctPort,omitempty"`

	// AuthPort
	// AAA server authentication port
	AuthPort *int `json:"authPort,omitempty"`

	// CreatedTime
	// The create time of the AAA server
	CreatedTime *int `json:"createdTime,omitempty"`

	// CreatorId
	// AAA server creator Id
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// AAA server creator name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Id
	// AAA server ID
	Id *string `json:"id,omitempty"`

	// Ip
	// AAA server IP address
	Ip *string `json:"ip,omitempty"`

	// Level
	// Access level of AAA server
	// Constraints:
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty" validate:"oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Name
	// Name of the AAA server
	Name *string `json:"name,omitempty"`

	// Password
	// Password for local user
	Password *string `json:"password,omitempty"`

	// Purpose
	// AAA server purpose
	// Constraints:
	//    - oneof:[DEFAULT,AUTHENTICATION_ONLY,AUTHORIZATION_ONLY,ACCOUNTING_ONLY]
	Purpose *string `json:"purpose,omitempty" validate:"oneof=DEFAULT AUTHENTICATION_ONLY AUTHORIZATION_ONLY ACCOUNTING_ONLY"`

	// Secret
	// AAA server secret
	Secret *string `json:"secret,omitempty"`

	// ServerType
	// The server tpye of the AAA server
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	ServerType *string `json:"serverType,omitempty" validate:"oneof=RADIUS TACACS_PLUS LOCAL"`

	// UpdatedTime
	// The modify time of the AAA server
	UpdatedTime *int `json:"updatedTime,omitempty"`

	// UpdaterId
	// AAA server updater Id
	UpdaterId *string `json:"updaterId,omitempty"`

	// UpdaterUsername
	// AAA server updater name
	UpdaterUsername *string `json:"updaterUsername,omitempty"`

	// Username
	// Username for local user
	Username *string `json:"username,omitempty"`
}

func NewSwitchMAAAServersAAAServer() *SwitchMAAAServersAAAServer {
	m := new(SwitchMAAAServersAAAServer)
	return m
}

type SwitchMAAAServersQueryResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMAAAServersQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first AAA Server returned out of the complete AAA Server list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more AAA Servers after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMAAAServersAAAServer `json:"list,omitempty"`

	// RawDataTotalCount
	// Total AAA Servers count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total AAA Servers count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMAAAServersQueryResult() *SwitchMAAAServersQueryResult {
	m := new(SwitchMAAAServersQueryResult)
	return m
}

// SwitchMAAAServersQueryResultExtraType
//
// Any additional response data
type SwitchMAAAServersQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMAAAServersQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMAAAServersQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMAAAServersQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMAAAServersQueryResultExtraType() *SwitchMAAAServersQueryResultExtraType {
	m := new(SwitchMAAAServersQueryResultExtraType)
	return m
}

type SwitchMAAAServersCreateAdminAAAServer struct {
	// AcctPort
	// AAA server accounting port
	AcctPort *int `json:"acctPort,omitempty"`

	// AuthPort
	// AAA server authentication port
	AuthPort *int `json:"authPort,omitempty"`

	// Ip
	// AAA server IP address
	Ip *string `json:"ip,omitempty"`

	// Level
	// Access level of AAA server
	// Constraints:
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty" validate:"oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Name
	// Name of the AAA server
	Name *string `json:"name,omitempty"`

	// Password
	// Password for local user
	Password *string `json:"password,omitempty"`

	// Purpose
	// AAA server purpose
	// Constraints:
	//    - oneof:[DEFAULT,AUTHENTICATION_ONLY,AUTHORIZATION_ONLY,ACCOUNTING_ONLY]
	Purpose *string `json:"purpose,omitempty" validate:"oneof=DEFAULT AUTHENTICATION_ONLY AUTHORIZATION_ONLY ACCOUNTING_ONLY"`

	// Secret
	// AAA server secret
	Secret *string `json:"secret,omitempty"`

	// ServerType
	// The server tpye of the AAA server
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	ServerType *string `json:"serverType,omitempty" validate:"oneof=RADIUS TACACS_PLUS LOCAL"`

	// Username
	// Username for local user
	Username *string `json:"username,omitempty"`
}

func NewSwitchMAAAServersCreateAdminAAAServer() *SwitchMAAAServersCreateAdminAAAServer {
	m := new(SwitchMAAAServersCreateAdminAAAServer)
	return m
}

type SwitchMAAAServersEmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMAAAServersEmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMAAAServersEmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMAAAServersEmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMAAAServersEmptyResult() *SwitchMAAAServersEmptyResult {
	m := new(SwitchMAAAServersEmptyResult)
	return m
}

// AddAaaServersAdmin
//
// Use this API command to create a new AAA server.
//
// Request Body:
//	 - body *SwitchMAAAServersCreateAdminAAAServer
func (s *SwitchMAAAServersService) AddAaaServersAdmin(ctx context.Context, body *SwitchMAAAServersCreateAdminAAAServer) (*SwitchMCommonCreateResult, error) {
	var (
		resp *SwitchMCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
}

// DeleteAaaServersAdmin
//
// Use this API command to delete AAA Servers.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMAAAServersService) DeleteAaaServersAdmin(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
}

// DeleteAaaServersAdminById
//
// Use this API command to delete a AAA server.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMAAAServersService) DeleteAaaServersAdminById(ctx context.Context, id string) (*SwitchMAAAServersEmptyResult, error) {
	var (
		resp *SwitchMAAAServersEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
}

// FindAaaServersAdmin
//
// Use this API command to retrieve a list of AAA server.
func (s *SwitchMAAAServersService) FindAaaServersAdmin(ctx context.Context) (*SwitchMAAAServersQueryResult, error) {
	var (
		resp *SwitchMAAAServersQueryResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindAaaServersAdminById
//
// Use this API command to retrieve a AAA server.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMAAAServersService) FindAaaServersAdminById(ctx context.Context, id string) (*SwitchMAAAServersAAAServer, error) {
	var (
		resp *SwitchMAAAServersAAAServer
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
}

// UpdateAaaServersAdminById
//
// Use this API command to modify the basic information on a AAA server by complete attributes.
//
// Request Body:
//	 - body *SwitchMAAAServersCreateAdminAAAServer
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMAAAServersService) UpdateAaaServersAdminById(ctx context.Context, body *SwitchMAAAServersCreateAdminAAAServer, id string) (*SwitchMAAAServersEmptyResult, error) {
	var (
		resp *SwitchMAAAServersEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
}
