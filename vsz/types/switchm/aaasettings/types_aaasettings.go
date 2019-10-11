package aaasettings

// API Version: v8_1

import (
	"encoding/json"
)

type AaaSettings struct {
	// Accounting
	// Accounting service
	Accounting *AaaSettingsAccountingType `json:"accounting,omitempty"`

	// Authentication
	// Authentication service
	Authentication *AaaSettingsAuthenticationType `json:"authentication,omitempty"`

	// Authorization
	// Authorization service
	Authorization *AaaSettingsAuthorizationType `json:"authorization,omitempty"`
}

func NewAaaSettings() *AaaSettings {
	aaaSettingsType := new(AaaSettings)
	return aaaSettingsType
}

func NewDefaultAaaSettings() *AaaSettings {
	aaaSettingsType := new(AaaSettings)
	return aaaSettingsType
}

// AaaSettingsAccountingType
//
// Accounting service
type AaaSettingsAccountingType struct {
	// Commands
	// Commands service
	Commands *AaaSettingsAccountingTypeCommandsType `json:"commands,omitempty"`

	// EnabledCommandAccounting
	// Command Enabled
	EnabledCommandAccounting *bool `json:"enabledCommandAccounting,omitempty"`

	// EnabledExecAccounting
	// Exec Enabled
	EnabledExecAccounting *bool `json:"enabledExecAccounting,omitempty"`

	// Exec
	// Exec service
	Exec *AaaSettingsAccountingTypeExecType `json:"exec,omitempty"`
}

func NewAaaSettingsAccountingType() *AaaSettingsAccountingType {
	aaaSettingsAccountingTypeType := new(AaaSettingsAccountingType)
	return aaaSettingsAccountingTypeType
}

func NewDefaultAaaSettingsAccountingType() *AaaSettingsAccountingType {
	aaaSettingsAccountingTypeType := new(AaaSettingsAccountingType)
	return aaaSettingsAccountingTypeType
}

// AaaSettingsAccountingTypeCommandsType
//
// Commands service
type AaaSettingsAccountingTypeCommandsType struct {
	// Level
	// Access level of command
	// Constraints:
	//    - nullable
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty" validate:"omitempty,oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Server1
	// Primary server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`
}

func NewAaaSettingsAccountingTypeCommandsType() *AaaSettingsAccountingTypeCommandsType {
	aaaSettingsAccountingTypeCommandsTypeType := new(AaaSettingsAccountingTypeCommandsType)
	return aaaSettingsAccountingTypeCommandsTypeType
}

func NewDefaultAaaSettingsAccountingTypeCommandsType() *AaaSettingsAccountingTypeCommandsType {
	aaaSettingsAccountingTypeCommandsTypeType := new(AaaSettingsAccountingTypeCommandsType)
	return aaaSettingsAccountingTypeCommandsTypeType
}

// AaaSettingsAccountingTypeExecType
//
// Exec service
type AaaSettingsAccountingTypeExecType struct {
	// Server1
	// Primary server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`
}

func NewAaaSettingsAccountingTypeExecType() *AaaSettingsAccountingTypeExecType {
	aaaSettingsAccountingTypeExecTypeType := new(AaaSettingsAccountingTypeExecType)
	return aaaSettingsAccountingTypeExecTypeType
}

func NewDefaultAaaSettingsAccountingTypeExecType() *AaaSettingsAccountingTypeExecType {
	aaaSettingsAccountingTypeExecTypeType := new(AaaSettingsAccountingTypeExecType)
	return aaaSettingsAccountingTypeExecTypeType
}

// AaaSettingsAuthenticationType
//
// Authentication service
type AaaSettingsAuthenticationType struct {
	// EnabledSSHAuthn
	// SSH Enabled
	EnabledSSHAuthn *bool `json:"enabledSSHAuthn,omitempty"`

	// EnableTelnetAuthn
	// Telnet Enabled
	EnableTelnetAuthn *bool `json:"enableTelnetAuthn,omitempty"`

	// FirstPref
	// Primary server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	FirstPref *string `json:"firstPref,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS LOCAL"`

	// SecondPref
	// Secondary server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	SecondPref *string `json:"secondPref,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS LOCAL"`

	// ThirdPref
	// Third server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	ThirdPref *string `json:"thirdPref,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS LOCAL"`
}

func NewAaaSettingsAuthenticationType() *AaaSettingsAuthenticationType {
	aaaSettingsAuthenticationTypeType := new(AaaSettingsAuthenticationType)
	return aaaSettingsAuthenticationTypeType
}

func NewDefaultAaaSettingsAuthenticationType() *AaaSettingsAuthenticationType {
	aaaSettingsAuthenticationTypeType := new(AaaSettingsAuthenticationType)
	return aaaSettingsAuthenticationTypeType
}

// AaaSettingsAuthorizationType
//
// Authorization service
type AaaSettingsAuthorizationType struct {
	// Commands
	// Commands service
	Commands *AaaSettingsAuthorizationTypeCommandsType `json:"commands,omitempty"`

	// EnabledCommandAuthz
	// Command Enabled
	EnabledCommandAuthz *bool `json:"enabledCommandAuthz,omitempty"`

	// EnabledExecAuthz
	// Exec Enabled
	EnabledExecAuthz *bool `json:"enabledExecAuthz,omitempty"`

	// Exec
	// Exec Service
	Exec *AaaSettingsAuthorizationTypeExecType `json:"exec,omitempty"`
}

func NewAaaSettingsAuthorizationType() *AaaSettingsAuthorizationType {
	aaaSettingsAuthorizationTypeType := new(AaaSettingsAuthorizationType)
	return aaaSettingsAuthorizationTypeType
}

func NewDefaultAaaSettingsAuthorizationType() *AaaSettingsAuthorizationType {
	aaaSettingsAuthorizationTypeType := new(AaaSettingsAuthorizationType)
	return aaaSettingsAuthorizationTypeType
}

// AaaSettingsAuthorizationTypeCommandsType
//
// Commands service
type AaaSettingsAuthorizationTypeCommandsType struct {
	// Level
	// Access level of command
	// Constraints:
	//    - nullable
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty" validate:"omitempty,oneof=READ_WRITE PORT_CONFIG READ_ONLY"`

	// Server1
	// Primary server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`
}

func NewAaaSettingsAuthorizationTypeCommandsType() *AaaSettingsAuthorizationTypeCommandsType {
	aaaSettingsAuthorizationTypeCommandsTypeType := new(AaaSettingsAuthorizationTypeCommandsType)
	return aaaSettingsAuthorizationTypeCommandsTypeType
}

func NewDefaultAaaSettingsAuthorizationTypeCommandsType() *AaaSettingsAuthorizationTypeCommandsType {
	aaaSettingsAuthorizationTypeCommandsTypeType := new(AaaSettingsAuthorizationTypeCommandsType)
	return aaaSettingsAuthorizationTypeCommandsTypeType
}

// AaaSettingsAuthorizationTypeExecType
//
// Exec Service
type AaaSettingsAuthorizationTypeExecType struct {
	// Server1
	// Primary server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server1 *string `json:"server1,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`

	// Server2
	// Secondary server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS_PLUS]
	Server2 *string `json:"server2,omitempty" validate:"omitempty,oneof=RADIUS TACACS_PLUS"`
}

func NewAaaSettingsAuthorizationTypeExecType() *AaaSettingsAuthorizationTypeExecType {
	aaaSettingsAuthorizationTypeExecTypeType := new(AaaSettingsAuthorizationTypeExecType)
	return aaaSettingsAuthorizationTypeExecTypeType
}

func NewDefaultAaaSettingsAuthorizationTypeExecType() *AaaSettingsAuthorizationTypeExecType {
	aaaSettingsAuthorizationTypeExecTypeType := new(AaaSettingsAuthorizationTypeExecType)
	return aaaSettingsAuthorizationTypeExecTypeType
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

func NewEmptyResult() *EmptyResult {
	emptyResultType := new(EmptyResult)
	return emptyResultType
}

func NewDefaultEmptyResult() *EmptyResult {
	emptyResultType := new(EmptyResult)
	return emptyResultType
}
