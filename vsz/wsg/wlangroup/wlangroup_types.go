package wlangroup

// API Version: v8_0

type CreateWLANGroup struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type ModifyWLANGroup struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type ModifyWLANGroupMember struct {
	AccessVlan  *int               `json:"accessVlan,omitempty"`
	NasID       *string            `json:"nasId,omitempty"`
	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}

type WLANGroup struct {
	CreateDateTime   *int       `json:"createDateTime,omitempty"`
	CreatorID        *string    `json:"creatorId,omitempty"`
	CreatorUsername  *string    `json:"creatorUsername,omitempty"`
	Description      *string    `json:"description,omitempty"`
	ID               *string    `json:"id,omitempty"`
	Members          []*Members `json:"members,omitempty"`
	ModifiedDateTime *int       `json:"modifiedDateTime,omitempty"`
	ModifierID       *string    `json:"modifierId,omitempty"`
	ModifierUsername *string    `json:"modifierUsername,omitempty"`
	Name             *string    `json:"name,omitempty"`
	ZoneID           *string    `json:"zoneId,omitempty"`
}

type WLANGroupList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type WLANMember struct {
	AccessVlan  *int               `json:"accessVlan,omitempty"`
	ID          *string            `json:"id,omitempty"`
	NasID       *string            `json:"nasId,omitempty"`
	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}
