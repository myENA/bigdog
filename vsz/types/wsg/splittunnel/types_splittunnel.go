package splittunnel

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateSplitTunnelProfile struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name" validate:"required"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*SplitTunnelIpMaskRule `json:"rules" validate:"required"`
}

type ModifySplitTunnelProfile struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*SplitTunnelIpMaskRule `json:"rules,omitempty"`
}

type SplitTunnelIpMaskRule struct {
	// DestinationIp
	// Destination IP of split tunnel profile rule
	DestinationIp *string `json:"destinationIp" validate:"required"`

	// DestinationIpMask
	// Destination IP mask of split tunnel profile rule
	DestinationIpMask *string `json:"destinationIpMask" validate:"required"`
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

type SplitTunnelProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SplitTunnelProfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SplitTunnelProfileListType struct {
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type SplitTunnelProfileQuery struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SplitTunnelProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
