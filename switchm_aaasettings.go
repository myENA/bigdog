package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type SwitchMAAASettingsService struct {
	apiClient *VSZClient
}

func NewSwitchMAAASettingsService(c *VSZClient) *SwitchMAAASettingsService {
	s := new(SwitchMAAASettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMAAASettingsService() *SwitchMAAASettingsService {
	return NewSwitchMAAASettingsService(ss.apiClient)
}

// SwitchMAAASettingsAAASetting
//
// Definition: aaaSettings_AAASetting
type SwitchMAAASettingsAAASetting struct {
	// Accounting
	// Accounting service
	Accounting *SwitchMAAASettingsAAASettingAccountingType `json:"accounting,omitempty"`

	AltoId *string `json:"altoId,omitempty"`

	// Authentication
	// Authentication service
	Authentication *SwitchMAAASettingsAAASettingAuthenticationType `json:"authentication,omitempty"`

	// Authorization
	// Authorization service
	Authorization *SwitchMAAASettingsAAASettingAuthorizationType `json:"authorization,omitempty"`
}

type SwitchMAAASettingsAAASettingAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMAAASettingsAAASetting
}

func newSwitchMAAASettingsAAASettingAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMAAASettingsAAASettingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMAAASettingsAAASettingAPIResponse) Hydrate() error {
	r.Data = new(SwitchMAAASettingsAAASetting)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMAAASettingsAAASetting() *SwitchMAAASettingsAAASetting {
	m := new(SwitchMAAASettingsAAASetting)
	return m
}

// SwitchMAAASettingsAAASettingAccountingType
//
// Definition: aaaSettings_AAASettingAccountingType
//
// Accounting service
type SwitchMAAASettingsAAASettingAccountingType struct {
	// Commands
	// Commands service
	Commands *SwitchMAAASettingsAAASettingAccountingTypeCommandsType `json:"commands,omitempty"`

	// EnabledCommandAccounting
	// Command Enabled
	EnabledCommandAccounting *bool `json:"enabledCommandAccounting,omitempty"`

	// EnabledExecAccounting
	// Exec Enabled
	EnabledExecAccounting *bool `json:"enabledExecAccounting,omitempty"`

	// Exec
	// Exec service
	Exec *SwitchMAAASettingsAAASettingAccountingTypeExecType `json:"exec,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAccountingType() *SwitchMAAASettingsAAASettingAccountingType {
	m := new(SwitchMAAASettingsAAASettingAccountingType)
	return m
}

// SwitchMAAASettingsAAASettingAccountingTypeCommandsType
//
// Definition: aaaSettings_AAASettingAccountingTypeCommandsType
//
// Commands service
type SwitchMAAASettingsAAASettingAccountingTypeCommandsType struct {
	// Level
	// Access level of command
	// Constraints:
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty"`

	// Server1
	// Primary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAccountingTypeCommandsType() *SwitchMAAASettingsAAASettingAccountingTypeCommandsType {
	m := new(SwitchMAAASettingsAAASettingAccountingTypeCommandsType)
	return m
}

// SwitchMAAASettingsAAASettingAccountingTypeExecType
//
// Definition: aaaSettings_AAASettingAccountingTypeExecType
//
// Exec service
type SwitchMAAASettingsAAASettingAccountingTypeExecType struct {
	// Server1
	// Primary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAccountingTypeExecType() *SwitchMAAASettingsAAASettingAccountingTypeExecType {
	m := new(SwitchMAAASettingsAAASettingAccountingTypeExecType)
	return m
}

// SwitchMAAASettingsAAASettingAuthenticationType
//
// Definition: aaaSettings_AAASettingAuthenticationType
//
// Authentication service
type SwitchMAAASettingsAAASettingAuthenticationType struct {
	// EnabledSSHAuthn
	// SSH Enabled
	EnabledSSHAuthn *bool `json:"enabledSSHAuthn,omitempty"`

	// EnableTelnetAuthn
	// Telnet Enabled
	EnableTelnetAuthn *bool `json:"enableTelnetAuthn,omitempty"`

	// FirstPref
	// Primary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	FirstPref *string `json:"firstPref,omitempty"`

	// SecondPref
	// Secondary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	SecondPref *string `json:"secondPref,omitempty"`

	// ThirdPref
	// Third server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	ThirdPref *string `json:"thirdPref,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAuthenticationType() *SwitchMAAASettingsAAASettingAuthenticationType {
	m := new(SwitchMAAASettingsAAASettingAuthenticationType)
	return m
}

// SwitchMAAASettingsAAASettingAuthorizationType
//
// Definition: aaaSettings_AAASettingAuthorizationType
//
// Authorization service
type SwitchMAAASettingsAAASettingAuthorizationType struct {
	// Commands
	// Commands service
	Commands *SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType `json:"commands,omitempty"`

	// EnabledCommandAuthz
	// Command Enabled
	EnabledCommandAuthz *bool `json:"enabledCommandAuthz,omitempty"`

	// EnabledExecAuthz
	// Exec Enabled
	EnabledExecAuthz *bool `json:"enabledExecAuthz,omitempty"`

	// Exec
	// Exec Service
	Exec *SwitchMAAASettingsAAASettingAuthorizationTypeExecType `json:"exec,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAuthorizationType() *SwitchMAAASettingsAAASettingAuthorizationType {
	m := new(SwitchMAAASettingsAAASettingAuthorizationType)
	return m
}

// SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType
//
// Definition: aaaSettings_AAASettingAuthorizationTypeCommandsType
//
// Commands service
type SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType struct {
	// Level
	// Access level of command
	// Constraints:
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty"`

	// Server1
	// Primary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAuthorizationTypeCommandsType() *SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType {
	m := new(SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType)
	return m
}

// SwitchMAAASettingsAAASettingAuthorizationTypeExecType
//
// Definition: aaaSettings_AAASettingAuthorizationTypeExecType
//
// Exec Service
type SwitchMAAASettingsAAASettingAuthorizationTypeExecType struct {
	// Server1
	// Primary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAuthorizationTypeExecType() *SwitchMAAASettingsAAASettingAuthorizationTypeExecType {
	m := new(SwitchMAAASettingsAAASettingAuthorizationTypeExecType)
	return m
}

// FindGroupAaaSettingsByGroupId
//
// Operation ID: findGroupAaaSettingsByGroupId
//
// Use this API command to retrieve the AAA settings.
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMAAASettingsService) FindGroupAaaSettingsByGroupId(ctx context.Context, groupId string, mutators ...RequestMutator) (*SwitchMAAASettingsAAASettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMAAASettingsAAASettingAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMAAASettingsAAASettingAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindGroupAaaSettingsByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMAAASettingsAAASettingAPIResponse), err
}

// UpdateGroupAaaSettingsByGroupId
//
// Operation ID: updateGroupAaaSettingsByGroupId
//
// Use this API command to modify the AAA settings.
//
// Request Body:
//	 - body *SwitchMAAASettingsAAASetting
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMAAASettingsService) UpdateGroupAaaSettingsByGroupId(ctx context.Context, body *SwitchMAAASettingsAAASetting, groupId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateGroupAaaSettingsByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
