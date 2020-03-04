package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"fmt"
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
	Level *string `json:"level,omitempty" validate:"oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Server1
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`
}

func NewSwitchMAAASettingsAAASettingAccountingTypeCommandsType() *SwitchMAAASettingsAAASettingAccountingTypeCommandsType {
	m := new(SwitchMAAASettingsAAASettingAccountingTypeCommandsType)
	return m
}

type SwitchMAAASettingsAAASettingAccountingTypeExecType struct {
	// Server1
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`
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
	FirstPref *string `json:"firstPref,omitempty" validate:"oneof=RADIUS TACACS_PLUS LOCAL"`

	// SecondPref
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	SecondPref *string `json:"secondPref,omitempty" validate:"oneof=RADIUS TACACS_PLUS LOCAL"`

	// ThirdPref
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	ThirdPref *string `json:"thirdPref,omitempty" validate:"oneof=RADIUS TACACS_PLUS LOCAL"`
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
	Level *string `json:"level,omitempty" validate:"oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Server1
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`
}

func NewSwitchMAAASettingsAAASettingAuthorizationTypeCommandsType() *SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType {
	m := new(SwitchMAAASettingsAAASettingAuthorizationTypeCommandsType)
	return m
}

type SwitchMAAASettingsAAASettingAuthorizationTypeExecType struct {
	// Server1
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`
}

func NewSwitchMAAASettingsAAASettingAuthorizationTypeExecType() *SwitchMAAASettingsAAASettingAuthorizationTypeExecType {
	m := new(SwitchMAAASettingsAAASettingAuthorizationTypeExecType)
	return m
}

type SwitchMAAASettingsEmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMAAASettingsEmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMAAASettingsEmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMAAASettingsEmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMAAASettingsEmptyResult() *SwitchMAAASettingsEmptyResult {
	m := new(SwitchMAAASettingsEmptyResult)
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.switchMPath, RouteSwitchMFindGroupAaaSettingsByGroupId), true)
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMAAASettingsAAASetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *SwitchMAAASettingsService) UpdateGroupAaaSettingsByGroupId(ctx context.Context, body *SwitchMAAASettingsAAASetting, groupId string) (*SwitchMAAASettingsEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMAAASettingsEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, fmt.Sprintf("/%s%s", s.apiClient.switchMPath, RouteSwitchMUpdateGroupAaaSettingsByGroupId), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMAAASettingsEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}
