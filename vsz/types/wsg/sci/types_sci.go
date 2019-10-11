package sci

// API Version: v8_1

import (
	"encoding/json"
)

type CreateSciProfile struct {
	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPassword *string `json:"sciPassword" validate:"required"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile" validate:"required"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerHost *string `json:"sciServerHost" validate:"required"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerPort *string `json:"sciServerPort" validate:"required"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciSystemId *string `json:"sciSystemId" validate:"required"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciUser *string `json:"sciUser" validate:"required"`
}

func NewCreateSciProfile() *CreateSciProfile {
	createSciProfileType := new(CreateSciProfile)
	return createSciProfileType
}

func NewCreateSciProfileWithDefaults() *CreateSciProfile {
	createSciProfileType := new(CreateSciProfile)
	return createSciProfileType
}

type DeleteSciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	Id *string `json:"id" validate:"required"`
}

func NewDeleteSciProfile() *DeleteSciProfile {
	deleteSciProfileType := new(DeleteSciProfile)
	return deleteSciProfileType
}

func NewDeleteSciProfileWithDefaults() *DeleteSciProfile {
	deleteSciProfileType := new(DeleteSciProfile)
	return deleteSciProfileType
}

type DeleteSciProfileList struct {
	List []*DeleteSciProfile `json:"list,omitempty"`
}

func NewDeleteSciProfileList() *DeleteSciProfileList {
	deleteSciProfileListType := new(DeleteSciProfileList)
	return deleteSciProfileListType
}

func NewDeleteSciProfileListWithDefaults() *DeleteSciProfileList {
	deleteSciProfileListType := new(DeleteSciProfileList)
	return deleteSciProfileListType
}

type ModifyEventCode struct {
	// SciAcceptedEventCodes
	// Constraints:
	//    - required
	SciAcceptedEventCodes []int `json:"sciAcceptedEventCodes" validate:"required"`
}

func NewModifyEventCode() *ModifyEventCode {
	modifyEventCodeType := new(ModifyEventCode)
	return modifyEventCodeType
}

func NewModifyEventCodeWithDefaults() *ModifyEventCode {
	modifyEventCodeType := new(ModifyEventCode)
	return modifyEventCodeType
}

type ModifySciEnabled struct {
	// SciEnabled
	// Is SZ/SCI interface enabled or disabled
	// Constraints:
	//    - required
	SciEnabled *bool `json:"sciEnabled" validate:"required"`
}

func NewModifySciEnabled() *ModifySciEnabled {
	modifySciEnabledType := new(ModifySciEnabled)
	return modifySciEnabledType
}

func NewModifySciEnabledWithDefaults() *ModifySciEnabled {
	modifySciEnabledType := new(ModifySciEnabled)
	return modifySciEnabledType
}

type ModifySciPriorityList struct {
	List []*ModifySciPriorityListType `json:"list,omitempty"`
}

func NewModifySciPriorityList() *ModifySciPriorityList {
	modifySciPriorityListType := new(ModifySciPriorityList)
	return modifySciPriorityListType
}

func NewModifySciPriorityListWithDefaults() *ModifySciPriorityList {
	modifySciPriorityListType := new(ModifySciPriorityList)
	return modifySciPriorityListType
}

type ModifySciPriorityListType struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	Id *string `json:"id" validate:"required"`

	// SciPriority
	// Priority of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPriority *int `json:"sciPriority" validate:"required"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile" validate:"required"`
}

func NewModifySciPriorityListType() *ModifySciPriorityListType {
	modifySciPriorityListTypeType := new(ModifySciPriorityListType)
	return modifySciPriorityListTypeType
}

func NewModifySciPriorityListTypeWithDefaults() *ModifySciPriorityListType {
	modifySciPriorityListTypeType := new(ModifySciPriorityListType)
	return modifySciPriorityListTypeType
}

type ModifySciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	Id *string `json:"id,omitempty"`

	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPassword *string `json:"sciPassword" validate:"required"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile" validate:"required"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerHost *string `json:"sciServerHost" validate:"required"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerPort *string `json:"sciServerPort" validate:"required"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciSystemId *string `json:"sciSystemId" validate:"required"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciUser *string `json:"sciUser" validate:"required"`
}

func NewModifySciProfile() *ModifySciProfile {
	modifySciProfileType := new(ModifySciProfile)
	return modifySciProfileType
}

func NewModifySciProfileWithDefaults() *ModifySciProfile {
	modifySciProfileType := new(ModifySciProfile)
	return modifySciProfileType
}

type SciEventCode struct {
	// FirstIndex
	// Index of the first event code returned from the complete event code set
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more SCI accepted event codes after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SciEventCodeListType `json:"list,omitempty"`

	// TotalCount
	// Total SCI accepted event code count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSciEventCode() *SciEventCode {
	sciEventCodeType := new(SciEventCode)
	return sciEventCodeType
}

func NewSciEventCodeWithDefaults() *SciEventCode {
	sciEventCodeType := new(SciEventCode)
	return sciEventCodeType
}

type SciEventCodeListType struct {
	// Code
	// SCI accepted event code
	Code *int `json:"code,omitempty"`

	// Type
	// SCI accepted event type
	Type *string `json:"type,omitempty"`
}

func NewSciEventCodeListType() *SciEventCodeListType {
	sciEventCodeListTypeType := new(SciEventCodeListType)
	return sciEventCodeListTypeType
}

func NewSciEventCodeListTypeWithDefaults() *SciEventCodeListType {
	sciEventCodeListTypeType := new(SciEventCodeListType)
	return sciEventCodeListTypeType
}

type SciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	Id *string `json:"id,omitempty"`

	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	SciPassword *string `json:"sciPassword,omitempty"`

	// SciPriority
	// Priority of the SCI profile for SZ/SCI interface
	SciPriority *int `json:"sciPriority,omitempty"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	SciProfile *string `json:"sciProfile,omitempty"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	SciServerHost *string `json:"sciServerHost,omitempty"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	SciServerPort *string `json:"sciServerPort,omitempty"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	SciSystemId *string `json:"sciSystemId,omitempty"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	SciUser *string `json:"sciUser,omitempty"`
}

func NewSciProfile() *SciProfile {
	sciProfileType := new(SciProfile)
	return sciProfileType
}

func NewSciProfileWithDefaults() *SciProfile {
	sciProfileType := new(SciProfile)
	return sciProfileType
}

type SciProfileList struct {
	Extra *SciProfileListExtraType `json:"extra,omitempty"`

	List []*SciProfile `json:"list,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SciProfileList) UnmarshalJSON(b []byte) error {
	tmpt := new(SciProfileList)
	if err := json.Unmarshal(b, tmpt); err != nil {
		return err
	}
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	delete(tmp, "extra")
	delete(tmp, "list")
	tmpt.XAdditionalProperties = tmp
	*t = *tmpt
	return nil
}

func (t *SciProfileList) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Extra != nil {
		tmp["extra"] = t.Extra
	}
	if t.List != nil {
		tmp["list"] = t.List
	}
	return json.Marshal(tmp)
}

func NewSciProfileList() *SciProfileList {
	sciProfileListType := new(SciProfileList)
	return sciProfileListType
}

func NewSciProfileListWithDefaults() *SciProfileList {
	sciProfileListType := new(SciProfileList)
	return sciProfileListType
}

type SciProfileListExtraType struct {
	// SciEnabled
	// SCI password of the SCI profile for SZ/SCI interface
	SciEnabled *bool `json:"sciEnabled,omitempty"`
}

func NewSciProfileListExtraType() *SciProfileListExtraType {
	sciProfileListExtraTypeType := new(SciProfileListExtraType)
	return sciProfileListExtraTypeType
}

func NewSciProfileListExtraTypeWithDefaults() *SciProfileListExtraType {
	sciProfileListExtraTypeType := new(SciProfileListExtraType)
	return sciProfileListExtraTypeType
}
