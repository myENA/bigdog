package wificalling

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateWifiCallingPolicy struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy
	// belongs
	DomainId *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*Epdg `json:"epdgs" validate:"required"`

	Name *common.NormalName `json:"name" validate:"required"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	Priority *string `json:"priority" validate:"required,oneof=BACKGROUND BEST_EFFORT VIDEO VOICE"`
}

type DeleteBulk struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type Epdg struct {
	// Fqdn
	// Fully qualified domain name of ePDG
	Fqdn *string `json:"fqdn,omitempty"`

	Ip *string `json:"ip,omitempty"`
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
	Priority *string `json:"priority,omitempty" validate:"omitempty,oneof=BACKGROUND BEST_EFFORT VIDEO VOICE"`
}

type WifiCallingPolicy struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy
	// belongs
	DomainId *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*Epdg `json:"epdgs,omitempty"`

	// Id
	// Identifier of the Wi-Fi calling policy
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	Priority *string `json:"priority,omitempty" validate:"omitempty,oneof=BACKGROUND BEST_EFFORT VIDEO VOICE"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`
}

type WifiCallingPolicyList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WifiCallingPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
