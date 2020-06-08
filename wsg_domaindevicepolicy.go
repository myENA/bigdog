package ruckus

// API Version: v9_0

type WSGDomaindevicepolicyCreateDomainDevicePolicy struct {
	// DefaultAction
	// defaultAction of the device policy cofig
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction"`

	Description **WSGDomaindevicepolicyCreateDomainDevicePolicy `json:"description,omitempty"`

	// DomainId
	// domainId of the device policy cofig
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name **WSGDomaindevicepolicyCreateDomainDevicePolicy `json:"name"`

	// Rule
	// rule of the device policy cofig
	// Constraints:
	//    - required
	Rule []**WSGDomaindevicepolicyCreateDomainDevicePolicy `json:"rule"`
}

func NewWSGDomaindevicepolicyCreateDomainDevicePolicy() *WSGDomaindevicepolicyCreateDomainDevicePolicy {
	m := new(WSGDomaindevicepolicyCreateDomainDevicePolicy)
	return m
}

type WSGDomaindevicepolicyProfile struct {
	// DefaultAction
	// defaultAction of the device policy cofig
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description **WSGDomaindevicepolicyProfile `json:"description,omitempty"`

	// Id
	// identifier of the device policy cofig
	Id *string `json:"id,omitempty"`

	Name **WSGDomaindevicepolicyProfile `json:"name,omitempty"`

	// Rule
	// rule of the device policy cofig
	Rule []**WSGDomaindevicepolicyProfile `json:"rule,omitempty"`
}

func NewWSGDomaindevicepolicyProfile() *WSGDomaindevicepolicyProfile {
	m := new(WSGDomaindevicepolicyProfile)
	return m
}

type WSGDomaindevicepolicyRule struct {
	// Action
	// defaultAction of the device policy cofig
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action,omitempty"`

	Description **WSGDomaindevicepolicyRule `json:"description,omitempty"`

	// DeviceType
	// deviceType of the device policy rule
	DeviceType *string `json:"deviceType,omitempty"`

	// Downlink
	// downlink rate Limiting of the device policy rule, rate unit is mbps, value must include be 0, 0.10, 0.25, 0.50, 0.75, 1.00, 1.25,1.50, 1.75, 2.00, 2.25, 2.50, 2.75, 3.00, 3.20, 3.50, 3.75, 4.00, 4.25, 4.50, 4.75, 5.00, 5.25, 5.50, 5.75, 6.00, 6.25, 6.50, 6.75,  7.00, 7.25, 7.50, 7.75, 8.00, 8.25, 8.50, 8.75,  9.00, 9.25, 9.50, 9.75,  10.00, 10.25, 10.50, 10.75, 11.00, 11.25, 11.50, 11.75, 12.00, 12.25, 12.50, 12.75, 13.00, 13.25, 13.50, 13.75, 14.00, 14.25, 14.50, 14.75, 15.00, 15.25, 15.50, 15.75, 16.00, 16.25, 16.50, 16.75, 17.00, 17.25, 17.50, 17.75, 18.00, 18.25, 18.50, 18.75, 19.00, 19.25, 19.50, 19.75, 20.00
	// Constraints:
	//    - min:0.000000
	//    - max:200.000000
	Downlink *float64 `json:"downlink,omitempty"`

	// OsVendor
	// osVendor of the device policy rule
	OsVendor *string `json:"osVendor,omitempty"`

	// Uplink
	// uplink rate Limiting of the device policy rule, rate unit is mbps, value must include be 0, 0.10, 0.25, 0.50, 0.75, 1.00, 1.25,1.50, 1.75, 2.00, 2.25, 2.50, 2.75, 3.00, 3.20, 3.50, 3.75, 4.00, 4.25, 4.50, 4.75, 5.00, 5.25, 5.50, 5.75, 6.00, 6.25, 6.50, 6.75,  7.00, 7.25, 7.50, 7.75, 8.00, 8.25, 8.50, 8.75,  9.00, 9.25, 9.50, 9.75,  10.00, 10.25, 10.50, 10.75, 11.00, 11.25, 11.50, 11.75, 12.00, 12.25, 12.50, 12.75, 13.00, 13.25, 13.50, 13.75, 14.00, 14.25, 14.50, 14.75, 15.00, 15.25, 15.50, 15.75, 16.00, 16.25, 16.50, 16.75, 17.00, 17.25, 17.50, 17.75, 18.00, 18.25, 18.50, 18.75, 19.00, 19.25, 19.50, 19.75, 20.00
	// Constraints:
	//    - min:0.000000
	//    - max:200.000000
	Uplink *float64 `json:"uplink,omitempty"`

	// Vlan
	// VLAN Members of the ethernet port profile
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	Vlan *int `json:"vlan,omitempty"`
}

func NewWSGDomaindevicepolicyRule() *WSGDomaindevicepolicyRule {
	m := new(WSGDomaindevicepolicyRule)
	return m
}

type WSGDomaindevicepolicyModifyDomainDevicePolicy struct {
	// DefaultAction
	// defaultAction of the device policy cofig
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description **WSGDomaindevicepolicyModifyDomainDevicePolicy `json:"description,omitempty"`

	Name **WSGDomaindevicepolicyModifyDomainDevicePolicy `json:"name,omitempty"`

	// Rule
	// rule of the device policy cofig
	Rule []**WSGDomaindevicepolicyModifyDomainDevicePolicy `json:"rule,omitempty"`
}

func NewWSGDomaindevicepolicyModifyDomainDevicePolicy() *WSGDomaindevicepolicyModifyDomainDevicePolicy {
	m := new(WSGDomaindevicepolicyModifyDomainDevicePolicy)
	return m
}

type WSGDomaindevicepolicyProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGDomaindevicepolicyProfileList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDomaindevicepolicyProfileList() *WSGDomaindevicepolicyProfileList {
	m := new(WSGDomaindevicepolicyProfileList)
	return m
}

type WSGDomaindevicepolicyProfileListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name **WSGDomaindevicepolicyProfileListType `json:"name,omitempty"`
}

func NewWSGDomaindevicepolicyProfileListType() *WSGDomaindevicepolicyProfileListType {
	m := new(WSGDomaindevicepolicyProfileListType)
	return m
}
