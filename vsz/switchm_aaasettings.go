package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMAAASettingsService struct {
	apiClient *APIClient
}

func NewSwitchMAAASettingsService(c *APIClient) *SwitchMAAASettingsService {
	s := new(SwitchMAAASettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMAAASettingsService() *SwitchMAAASettingsService {
	return NewSwitchMAAASettingsService(ss.apiClient)
}

type SwitchMAAASettingsAAASetting struct {
	// Accounting
	// Constraints:
	//    - nullable
	Accounting *SwitchMAAASettingsAAASettingAccountingType `json:"accounting,omitempty"`

	// Authentication
	// Constraints:
	//    - nullable
	Authentication *SwitchMAAASettingsAAASettingAuthenticationType `json:"authentication,omitempty"`

	// Authorization
	// Constraints:
	//    - nullable
	Authorization *SwitchMAAASettingsAAASettingAuthorizationType `json:"authorization,omitempty"`
}

func NewSwitchMAAASettingsAAASetting() *SwitchMAAASettingsAAASetting {
	m := new(SwitchMAAASettingsAAASetting)
	return m
}

type SwitchMAAASettingsAAASettingAccountingType struct {
	// Commands
	// Constraints:
	//    - nullable
	Commands *SwitchMAAASettingsAAASettingAccountingTypeCommandsType `json:"commands,omitempty"`

	// EnabledCommandAccounting
	// Constraints:
	//    - nullable
	EnabledCommandAccounting *bool `json:"enabledCommandAccounting,omitempty"`

	// EnabledExecAccounting
	// Constraints:
	//    - nullable
	EnabledExecAccounting *bool `json:"enabledExecAccounting,omitempty"`

	// Exec
	// Constraints:
	//    - nullable
	Exec *SwitchMAAASettingsAAASettingAccountingTypeExecType `json:"exec,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAccountingType() *SwitchMAAASettingsAAASettingAccountingType {
	m := new(SwitchMAAASettingsAAASettingAccountingType)
	return m
}

type SwitchMAAASettingsAAASettingAccountingTypeCommandsType struct {
	// Level
	// Constraints:
	//    - nullable
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty" validate:"omitempty,oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Server1
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`
}

func NewSwitchMAAASettingsAAASettingAccountingTypeCommandsType() *SwitchMAAASettingsAAASettingAccountingTypeCommandsType {
	m := new(SwitchMAAASettingsAAASettingAccountingTypeCommandsType)
	return m
}

type SwitchMAAASettingsAAASettingAccountingTypeExecType struct {
	// Server1
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`
}

func NewSwitchMAAASettingsAAASettingAccountingTypeExecType() *SwitchMAAASettingsAAASettingAccountingTypeExecType {
	m := new(SwitchMAAASettingsAAASettingAccountingTypeExecType)
	return m
}

type SwitchMAAASettingsAAASettingAuthenticationType struct {
	// EnabledSSHAuthn
	// Constraints:
	//    - nullable
	EnabledSSHAuthn *bool `json:"enabledSSHAuthn,omitempty"`

	// EnableTelnetAuthn
	// Constraints:
	//    - nullable
	EnableTelnetAuthn *bool `json:"enableTelnetAuthn,omitempty"`

	// FirstPref
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	FirstPref *string `json:"firstPref,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS LOCAL"`

	// SecondPref
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	SecondPref *string `json:"secondPref,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS LOCAL"`

	// ThirdPref
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	ThirdPref *string `json:"thirdPref,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS LOCAL"`
}

func NewSwitchMAAASettingsAAASettingAuthenticationType() *SwitchMAAASettingsAAASettingAuthenticationType {
	m := new(SwitchMAAASettingsAAASettingAuthenticationType)
	return m
}

type SwitchMAAASettingsAAASettingAuthorizationType struct {
	// Commands
	// Constraints:
	//    - nullable
	Commands *SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType `json:"commands,omitempty"`

	// EnabledCommandAuthz
	// Constraints:
	//    - nullable
	EnabledCommandAuthz *bool `json:"enabledCommandAuthz,omitempty"`

	// EnabledExecAuthz
	// Constraints:
	//    - nullable
	EnabledExecAuthz *bool `json:"enabledExecAuthz,omitempty"`

	// Exec
	// Constraints:
	//    - nullable
	Exec *SwitchMAAASettingsAAASettingAuthorizationTypeExecType `json:"exec,omitempty"`
}

func NewSwitchMAAASettingsAAASettingAuthorizationType() *SwitchMAAASettingsAAASettingAuthorizationType {
	m := new(SwitchMAAASettingsAAASettingAuthorizationType)
	return m
}

type SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType struct {
	// Level
	// Constraints:
	//    - nullable
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty" validate:"omitempty,oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Server1
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`
}

func NewSwitchMAAASettingsAAASettingAuthorizationTypeCommandsType() *SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType {
	m := new(SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType)
	return m
}

type SwitchMAAASettingsAAASettingAuthorizationTypeExecType struct {
	// Server1
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`
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
func (s *SwitchMAAASettingsService) FindGroupAaaSettingsByGroupId(ctx context.Context, groupId string) (*SwitchMAAASettingsAAASetting, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindGroupAaaSettingsByGroupId, true)
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req)
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
func (s *SwitchMAAASettingsService) UpdateGroupAaaSettingsByGroupId(ctx context.Context, body *SwitchMAAASettingsAAASetting, groupId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateGroupAaaSettingsByGroupId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
