package splittunnel

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateSplitTunnelProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Rules
	// Destination rule of split tunnel profile
	// Constraints:
	//    - required
	Rules []*SplitTunnelIpMaskRule `json:"rules" validate:"required"`
}

func NewCreateSplitTunnelProfile() *CreateSplitTunnelProfile {
	createSplitTunnelProfileType := new(CreateSplitTunnelProfile)
	return createSplitTunnelProfileType
}

func NewCreateSplitTunnelProfileWithDefaults() *CreateSplitTunnelProfile {
	createSplitTunnelProfileType := new(CreateSplitTunnelProfile)
	return createSplitTunnelProfileType
}

type ModifySplitTunnelProfile struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*SplitTunnelIpMaskRule `json:"rules,omitempty"`
}

func NewModifySplitTunnelProfile() *ModifySplitTunnelProfile {
	modifySplitTunnelProfileType := new(ModifySplitTunnelProfile)
	return modifySplitTunnelProfileType
}

func NewModifySplitTunnelProfileWithDefaults() *ModifySplitTunnelProfile {
	modifySplitTunnelProfileType := new(ModifySplitTunnelProfile)
	return modifySplitTunnelProfileType
}

type SplitTunnelIpMaskRule struct {
	// DestinationIp
	// Destination IP of split tunnel profile rule
	// Constraints:
	//    - required
	DestinationIp *string `json:"destinationIp" validate:"required"`

	// DestinationIpMask
	// Destination IP mask of split tunnel profile rule
	// Constraints:
	//    - required
	DestinationIpMask *string `json:"destinationIpMask" validate:"required"`
}

func NewSplitTunnelIpMaskRule() *SplitTunnelIpMaskRule {
	splitTunnelIpMaskRuleType := new(SplitTunnelIpMaskRule)
	return splitTunnelIpMaskRuleType
}

func NewSplitTunnelIpMaskRuleWithDefaults() *SplitTunnelIpMaskRule {
	splitTunnelIpMaskRuleType := new(SplitTunnelIpMaskRule)
	return splitTunnelIpMaskRuleType
}

type SplitTunnelProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*SplitTunnelIpMaskRule `json:"rules,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`
}

func NewSplitTunnelProfile() *SplitTunnelProfile {
	splitTunnelProfileType := new(SplitTunnelProfile)
	return splitTunnelProfileType
}

func NewSplitTunnelProfileWithDefaults() *SplitTunnelProfile {
	splitTunnelProfileType := new(SplitTunnelProfile)
	return splitTunnelProfileType
}

type SplitTunnelProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SplitTunnelProfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSplitTunnelProfileList() *SplitTunnelProfileList {
	splitTunnelProfileListType := new(SplitTunnelProfileList)
	return splitTunnelProfileListType
}

func NewSplitTunnelProfileListWithDefaults() *SplitTunnelProfileList {
	splitTunnelProfileListType := new(SplitTunnelProfileList)
	return splitTunnelProfileListType
}

type SplitTunnelProfileListType struct {
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewSplitTunnelProfileListType() *SplitTunnelProfileListType {
	splitTunnelProfileListTypeType := new(SplitTunnelProfileListType)
	return splitTunnelProfileListTypeType
}

func NewSplitTunnelProfileListTypeWithDefaults() *SplitTunnelProfileListType {
	splitTunnelProfileListTypeType := new(SplitTunnelProfileListType)
	return splitTunnelProfileListTypeType
}

type SplitTunnelProfileQuery struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SplitTunnelProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSplitTunnelProfileQuery() *SplitTunnelProfileQuery {
	splitTunnelProfileQueryType := new(SplitTunnelProfileQuery)
	return splitTunnelProfileQueryType
}

func NewSplitTunnelProfileQueryWithDefaults() *SplitTunnelProfileQuery {
	splitTunnelProfileQueryType := new(SplitTunnelProfileQuery)
	return splitTunnelProfileQueryType
}
