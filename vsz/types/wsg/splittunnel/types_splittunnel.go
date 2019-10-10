package splittunnel

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateSplitTunnelProfile struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*SplitTunnelIPMaskRule `json:"rules,omitempty" validate:"required"`
}

type ModifySplitTunnelProfile struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*SplitTunnelIPMaskRule `json:"rules,omitempty"`
}

type SplitTunnelIPMaskRule struct {
	// DestinationIP
	// Destination IP of split tunnel profile rule
	DestinationIP *string `json:"destinationIp,omitempty" validate:"required"`

	// DestinationIPMask
	// Destination IP mask of split tunnel profile rule
	DestinationIPMask *string `json:"destinationIpMask,omitempty" validate:"required"`
}

type SplitTunnelProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DomainID *string `json:"domainId,omitempty"`

	ID *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*SplitTunnelIPMaskRule `json:"rules,omitempty"`

	TenantID *string `json:"tenantId,omitempty"`

	ZoneID *string `json:"zoneId,omitempty"`
}

type SplitTunnelProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SplitTunnelProfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SplitTunnelProfileListType struct {
	ID *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type SplitTunnelProfileQuery struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SplitTunnelProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
