package wificalling

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateWifiCallingPolicy struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainID
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy
	// belongs
	DomainID *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*Epdg `json:"epdgs,omitempty" validate:"required"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	Priority *string `json:"priority,omitempty" validate:"required,oneof=BACKGROUND BEST_EFFORT VIDEO VOICE"`
}

type DeleteBulk struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type Epdg struct {
	// FQDN
	// Fully qualified domain name of ePDG
	FQDN *string `json:"fqdn,omitempty"`

	IP *string `json:"ip,omitempty"`
}

type ModifyEntireWifiCallingPolicy struct {
	*ModifyWifiCallingPolicy
}

type ModifyWifiCallingPolicy struct {
	Description *common.Description `json:"description,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*Epdg `json:"epdgs,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	Priority *string `json:"priority,omitempty" validate:"oneof=BACKGROUND BEST_EFFORT VIDEO VOICE"`
}

type WifiCallingPolicy struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainID
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy
	// belongs
	DomainID *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*Epdg `json:"epdgs,omitempty"`

	// ID
	// Identifier of the Wi-Fi calling policy
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	Priority *string `json:"priority,omitempty" validate:"oneof=BACKGROUND BEST_EFFORT VIDEO VOICE"`

	// TenantID
	// Tenant Id
	TenantID *string `json:"tenantId,omitempty"`
}

type WifiCallingPolicyList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WifiCallingPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
