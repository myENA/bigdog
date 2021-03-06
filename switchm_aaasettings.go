package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
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

func newSwitchMAAASettingsAAASettingAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMAAASettingsAAASettingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMAAASettingsAAASettingAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMAAASettingsAAASetting)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
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
// Use this API command to retrieve the AAA settings.
//
// Operation ID: findGroupAaaSettingsByGroupId
// Operation path: /group/{groupId}/aaaSettings
// Success code: 200 (OK)
//
// Required parameters:
// - groupId string
//		- required
func (s *SwitchMAAASettingsService) FindGroupAaaSettingsByGroupId(ctx context.Context, groupId string, mutators ...RequestMutator) (*SwitchMAAASettingsAAASettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMAAASettingsAAASettingAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindGroupAaaSettingsByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMAAASettingsAAASettingAPIResponse), err
}

// UpdateGroupAaaSettingsByGroupId
//
// Use this API command to modify the AAA settings.
//
// Operation ID: updateGroupAaaSettingsByGroupId
// Operation path: /group/{groupId}/aaaSettings
// Success code: 204 (No Content)
//
// Request body:
//	 - body *SwitchMAAASettingsAAASetting
//
// Required parameters:
// - groupId string
//		- required
func (s *SwitchMAAASettingsService) UpdateGroupAaaSettingsByGroupId(ctx context.Context, body *SwitchMAAASettingsAAASetting, groupId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateGroupAaaSettingsByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("groupId", groupId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
