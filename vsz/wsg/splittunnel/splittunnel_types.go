package splittunnel

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/common"
)

type CreateSplitTunnelProfile struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*SplitTunnelIPMaskRule `json:"rules,omitempty"`
}

type ModifySplitTunnelProfile struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*SplitTunnelIPMaskRule `json:"rules,omitempty"`
}

type SplitTunnelIPMaskRule struct {
	// DestinationIP
	// Destination IP of split tunnel profile rule
	DestinationIP *string `json:"destinationIp,omitempty"`

	// DestinationIPMask
	// Destination IP mask of split tunnel profile rule
	DestinationIPMask *string `json:"destinationIpMask,omitempty"`
}

type SplitTunnelProfile struct {
	Description *string `json:"description,omitempty"`

	DomainID *string `json:"domainId,omitempty"`

	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

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

	Name *string `json:"name,omitempty"`
}

type SplitTunnelProfileQuery struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SplitTunnelProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
