package devicepolicy

// API Version: v8_0

type CreateDevicePolicy struct {
	DefaultAction *string `json:"defaultAction,omitempty"`
	Description   *string `json:"description,omitempty"`
	Name          *string `json:"name,omitempty"`
}

type DevicePolicyPorfile struct {
	DefaultAction *string             `json:"defaultAction,omitempty"`
	Description   *string             `json:"description,omitempty"`
	ID            *string             `json:"id,omitempty"`
	Name          *string             `json:"name,omitempty"`
	Rule          []*DevicePolicyRule `json:"rule,omitempty"`
}

type DevicePolicyRule struct {
	Action      *string  `json:"action,omitempty"`
	Description *string  `json:"description,omitempty"`
	DeviceType  *string  `json:"deviceType,omitempty"`
	Downlink    *float64 `json:"downlink,omitempty"`
	Uplink      *float64 `json:"uplink,omitempty"`
	Vlan        *int     `json:"vlan,omitempty"`
}

type ModifyDevicePolicy struct {
	DefaultAction *string             `json:"defaultAction,omitempty"`
	Description   *string             `json:"description,omitempty"`
	Name          *string             `json:"name,omitempty"`
	Rule          []*DevicePolicyRule `json:"rule,omitempty"`
}

type PorfileList struct {
	FirstIndex *int               `json:"firstIndex,omitempty"`
	HasMore    *bool              `json:"hasMore,omitempty"`
	List       []*PorfileListType `json:"list,omitempty"`
	TotalCount *int               `json:"totalCount,omitempty"`
}

type PorfileListType struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
