package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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

type SwitchMAAASettings struct {
	// Accounting
	// Accounting service
	Accounting *SwitchMAAASettingsAccountingType `json:"accounting,omitempty"`

	// Authentication
	// Authentication service
	Authentication *SwitchMAAASettingsAuthenticationType `json:"authentication,omitempty"`

	// Authorization
	// Authorization service
	Authorization *SwitchMAAASettingsAuthorizationType `json:"authorization,omitempty"`
}

// SwitchMAAASettingsAccountingType
//
// Accounting service
type SwitchMAAASettingsAccountingType struct {
	// Commands
	// Commands service
	Commands *SwitchMAAASettingsAccountingTypeCommandsType `json:"commands,omitempty"`

	// EnabledCommandAccounting
	// Command Enabled
	EnabledCommandAccounting *bool `json:"enabledCommandAccounting,omitempty"`

	// EnabledExecAccounting
	// Exec Enabled
	EnabledExecAccounting *bool `json:"enabledExecAccounting,omitempty"`

	// Exec
	// Exec service
	Exec *SwitchMAAASettingsAccountingTypeExecType `json:"exec,omitempty"`
}

// SwitchMAAASettingsAccountingTypeCommandsType
//
// Commands service
type SwitchMAAASettingsAccountingTypeCommandsType struct {
	// Level
	// Access level of command
	// Constraints:
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty" validate:"oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Server1
	// Primary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`
}

// SwitchMAAASettingsAccountingTypeExecType
//
// Exec service
type SwitchMAAASettingsAccountingTypeExecType struct {
	// Server1
	// Primary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`
}

// SwitchMAAASettingsAuthenticationType
//
// Authentication service
type SwitchMAAASettingsAuthenticationType struct {
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
	FirstPref *string `json:"firstPref,omitempty" validate:"oneof=RADIUS TACACS_PLUS LOCAL"`

	// SecondPref
	// Secondary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	SecondPref *string `json:"secondPref,omitempty" validate:"oneof=RADIUS TACACS_PLUS LOCAL"`

	// ThirdPref
	// Third server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	ThirdPref *string `json:"thirdPref,omitempty" validate:"oneof=RADIUS TACACS_PLUS LOCAL"`
}

// SwitchMAAASettingsAuthorizationType
//
// Authorization service
type SwitchMAAASettingsAuthorizationType struct {
	// Commands
	// Commands service
	Commands *SwitchMAAASettingsAuthorizationTypeCommandsType `json:"commands,omitempty"`

	// EnabledCommandAuthz
	// Command Enabled
	EnabledCommandAuthz *bool `json:"enabledCommandAuthz,omitempty"`

	// EnabledExecAuthz
	// Exec Enabled
	EnabledExecAuthz *bool `json:"enabledExecAuthz,omitempty"`

	// Exec
	// Exec Service
	Exec *SwitchMAAASettingsAuthorizationTypeExecType `json:"exec,omitempty"`
}

// SwitchMAAASettingsAuthorizationTypeCommandsType
//
// Commands service
type SwitchMAAASettingsAuthorizationTypeCommandsType struct {
	// Level
	// Access level of command
	// Constraints:
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty" validate:"oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Server1
	// Primary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`
}

// SwitchMAAASettingsAuthorizationTypeExecType
//
// Exec Service
type SwitchMAAASettingsAuthorizationTypeExecType struct {
	// Server1
	// Primary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"oneof=RADIUS TACACS_PLUS"`
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

// FindAaaSettings
//
// Use this API command to retrieve the AAA settings.
func (s *SwitchMAAASettingsService) FindAaaSettings(ctx context.Context) (*SwitchMAAASettings, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAaaSettings
//
// Use this API command to modify the AAA settings.
//
// Request Body:
//	 - body *SwitchMAAASettings
func (s *SwitchMAAASettingsService) UpdateAaaSettings(ctx context.Context, body *SwitchMAAASettings) (*SwitchMAAASettingsEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
