package bigdog

// API Version: v9_0

import (
	"context"
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

type SwitchMAAASettingsAAASetting struct {
	Accounting *SwitchMAAASettingsAAASettingAccountingType `json:"accounting,omitempty"`

	Authentication *SwitchMAAASettingsAAASettingAuthenticationType `json:"authentication,omitempty"`

	Authorization *SwitchMAAASettingsAAASettingAuthorizationType `json:"authorization,omitempty"`
}

func NewSwitchMAAASettingsAAASetting() *SwitchMAAASettingsAAASetting {
	m := new(SwitchMAAASettingsAAASetting)
	return m
}

type SwitchMAAASettingsAAASettingAccountingType struct {
	Commands *SwitchMAAASettingsAAASettingAccountingTypeCommandsType `json:"commands,omitempty"`

	EnabledCommandAccounting *bool `json:"enabledCommandAccounting,omitempty"`

	EnabledExecAccounting *bool `json:"enabledExecAccounting,omitempty"`

	Exec *SwitchMAAASettingsAAASettingAccountingTypeExecType `json:"exec,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAccountingType() *SwitchMAAASettingsAAASettingAccountingType {
	m := new(SwitchMAAASettingsAAASettingAccountingType)
	return m
}

type SwitchMAAASettingsAAASettingAccountingTypeCommandsType struct {
	// Level
	// Constraints:
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty"`

	// Server1
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty"`

	// Server2
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAccountingTypeCommandsType() *SwitchMAAASettingsAAASettingAccountingTypeCommandsType {
	m := new(SwitchMAAASettingsAAASettingAccountingTypeCommandsType)
	return m
}

type SwitchMAAASettingsAAASettingAccountingTypeExecType struct {
	// Server1
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty"`

	// Server2
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAccountingTypeExecType() *SwitchMAAASettingsAAASettingAccountingTypeExecType {
	m := new(SwitchMAAASettingsAAASettingAccountingTypeExecType)
	return m
}

type SwitchMAAASettingsAAASettingAuthenticationType struct {
	EnabledSSHAuthn *bool `json:"enabledSSHAuthn,omitempty"`

	EnableTelnetAuthn *bool `json:"enableTelnetAuthn,omitempty"`

	// FirstPref
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	FirstPref *string `json:"firstPref,omitempty"`

	// SecondPref
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	SecondPref *string `json:"secondPref,omitempty"`

	// ThirdPref
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	ThirdPref *string `json:"thirdPref,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAuthenticationType() *SwitchMAAASettingsAAASettingAuthenticationType {
	m := new(SwitchMAAASettingsAAASettingAuthenticationType)
	return m
}

type SwitchMAAASettingsAAASettingAuthorizationType struct {
	Commands *SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType `json:"commands,omitempty"`

	EnabledCommandAuthz *bool `json:"enabledCommandAuthz,omitempty"`

	EnabledExecAuthz *bool `json:"enabledExecAuthz,omitempty"`

	Exec *SwitchMAAASettingsAAASettingAuthorizationTypeExecType `json:"exec,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAuthorizationType() *SwitchMAAASettingsAAASettingAuthorizationType {
	m := new(SwitchMAAASettingsAAASettingAuthorizationType)
	return m
}

type SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType struct {
	// Level
	// Constraints:
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty"`

	// Server1
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty"`

	// Server2
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAuthorizationTypeCommandsType() *SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType {
	m := new(SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType)
	return m
}

type SwitchMAAASettingsAAASettingAuthorizationTypeExecType struct {
	// Server1
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty"`

	// Server2
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
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMAAASettingsService) FindGroupAaaSettingsByGroupId(ctx context.Context, groupId string, mutators ...RequestMutator) (*SwitchMAAASettingsAAASetting, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMAAASettingsAAASetting
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindGroupAaaSettingsByGroupId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMAAASettingsAAASetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateGroupAaaSettingsByGroupId
//
// Use this API command to modify the AAA settings.
//
// Request Body:
//	 - body *SwitchMAAASettingsAAASetting
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMAAASettingsService) UpdateGroupAaaSettingsByGroupId(ctx context.Context, body *SwitchMAAASettingsAAASetting, groupId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateGroupAaaSettingsByGroupId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
