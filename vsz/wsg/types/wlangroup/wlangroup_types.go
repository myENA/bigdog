package wlangroup

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/types/common"
)

type CreateWLANGroup struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type ModifyWLANGroup struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type ModifyWLANGroupMember struct {
	// AccessVlan
	// Access VLAN
	AccessVlan *int `json:"accessVlan,omitempty"`

	// NasID
	// NAS-ID
	NasID *string `json:"nasId,omitempty"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}

type WLANGroup struct {
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

	// ID
	// Identifier of the WLAN group
	ID *string `json:"id,omitempty"`

	// Members
	// Members of the WLAN group
	Members []*WLANMember `json:"members,omitempty"`

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

	// ZoneID
	// Identifier of the zone to which the WLAN group belongs
	ZoneID *string `json:"zoneId,omitempty"`
}

type WLANGroupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WLANGroup `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WLANMember struct {
	// AccessVlan
	// Access VLAN
	AccessVlan *int `json:"accessVlan,omitempty"`

	// ID
	// Identifier of the WLAN
	ID *string `json:"id,omitempty"`

	// NasID
	// NAS-ID
	NasID *string `json:"nasId,omitempty"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}
