package vsz

// API Version: v9_0

type WSGDomainDevicePolicyCreateDomainDevicePolicy struct {
	// DefaultAction
	// defaultAction of the device policy cofig
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction" validate:"required,oneof=ALLOW BLOCK"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// domainId of the device policy cofig
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Rule
	// rule of the device policy cofig
	// Constraints:
	//    - required
	Rule []*WSGDomainDevicePolicyRule `json:"rule" validate:"required,dive"`
}

func NewWSGDomainDevicePolicyCreateDomainDevicePolicy() *WSGDomainDevicePolicyCreateDomainDevicePolicy {
	m := new(WSGDomainDevicePolicyCreateDomainDevicePolicy)
	return m
}

type WSGDomainDevicePolicyProfile struct {
	// DefaultAction
	// defaultAction of the device policy cofig
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction,omitempty" validate:"oneof=ALLOW BLOCK"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// identifier of the device policy cofig
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rule
	// rule of the device policy cofig
	Rule []*WSGDomainDevicePolicyRule `json:"rule"`
}

func NewWSGDomainDevicePolicyProfile() *WSGDomainDevicePolicyProfile {
	m := new(WSGDomainDevicePolicyProfile)
	return m
}

type WSGDomainDevicePolicyRule struct {
	// Action
	// defaultAction of the device policy cofig
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action,omitempty" validate:"oneof=ALLOW BLOCK"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DeviceType
	// deviceType of the device policy rule
	DeviceType *string `json:"deviceType,omitempty"`

	// Downlink
	// downlink rate Limiting of the device policy rule, rate unit is mbps, value must include be 0, 0.10, 0.25, 0.50, 0.75, 1.00, 1.25,1.50, 1.75, 2.00, 2.25, 2.50, 2.75, 3.00, 3.20, 3.50, 3.75, 4.00, 4.25, 4.50, 4.75, 5.00, 5.25, 5.50, 5.75, 6.00, 6.25, 6.50, 6.75,  7.00, 7.25, 7.50, 7.75, 8.00, 8.25, 8.50, 8.75,  9.00, 9.25, 9.50, 9.75,  10.00, 10.25, 10.50, 10.75, 11.00, 11.25, 11.50, 11.75, 12.00, 12.25, 12.50, 12.75, 13.00, 13.25, 13.50, 13.75, 14.00, 14.25, 14.50, 14.75, 15.00, 15.25, 15.50, 15.75, 16.00, 16.25, 16.50, 16.75, 17.00, 17.25, 17.50, 17.75, 18.00, 18.25, 18.50, 18.75, 19.00, 19.25, 19.50, 19.75, 20.00
	// Constraints:
	//    - min:0.000000
	//    - max:200.000000
	Downlink *float64 `json:"downlink,omitempty" validate:"gte=0.000000,lte=200.000000"`

	// OsVendor
	// osVendor of the device policy rule
	OsVendor *string `json:"osVendor,omitempty"`

	// Uplink
	// uplink rate Limiting of the device policy rule, rate unit is mbps, value must include be 0, 0.10, 0.25, 0.50, 0.75, 1.00, 1.25,1.50, 1.75, 2.00, 2.25, 2.50, 2.75, 3.00, 3.20, 3.50, 3.75, 4.00, 4.25, 4.50, 4.75, 5.00, 5.25, 5.50, 5.75, 6.00, 6.25, 6.50, 6.75,  7.00, 7.25, 7.50, 7.75, 8.00, 8.25, 8.50, 8.75,  9.00, 9.25, 9.50, 9.75,  10.00, 10.25, 10.50, 10.75, 11.00, 11.25, 11.50, 11.75, 12.00, 12.25, 12.50, 12.75, 13.00, 13.25, 13.50, 13.75, 14.00, 14.25, 14.50, 14.75, 15.00, 15.25, 15.50, 15.75, 16.00, 16.25, 16.50, 16.75, 17.00, 17.25, 17.50, 17.75, 18.00, 18.25, 18.50, 18.75, 19.00, 19.25, 19.50, 19.75, 20.00
	// Constraints:
	//    - min:0.000000
	//    - max:200.000000
	Uplink *float64 `json:"uplink,omitempty" validate:"gte=0.000000,lte=200.000000"`

	// Vlan
	// VLAN Members of the ethernet port profile
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	Vlan *int `json:"vlan,omitempty" validate:"omitempty,gte=1,lte=4094"`
}

func NewWSGDomainDevicePolicyRule() *WSGDomainDevicePolicyRule {
	m := new(WSGDomainDevicePolicyRule)
	return m
}

type WSGDomainDevicePolicyModifyDomainDevicePolicy struct {
	// DefaultAction
	// defaultAction of the device policy cofig
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction,omitempty" validate:"oneof=ALLOW BLOCK"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rule
	// rule of the device policy cofig
	Rule []*WSGDomainDevicePolicyRule `json:"rule"`
}

func NewWSGDomainDevicePolicyModifyDomainDevicePolicy() *WSGDomainDevicePolicyModifyDomainDevicePolicy {
	m := new(WSGDomainDevicePolicyModifyDomainDevicePolicy)
	return m
}

type WSGDomainDevicePolicyProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDomainDevicePolicyProfileListType `json:"list"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDomainDevicePolicyProfileList() *WSGDomainDevicePolicyProfileList {
	m := new(WSGDomainDevicePolicyProfileList)
	return m
}

type WSGDomainDevicePolicyProfileListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGDomainDevicePolicyProfileListType() *WSGDomainDevicePolicyProfileListType {
	m := new(WSGDomainDevicePolicyProfileListType)
	return m
}
