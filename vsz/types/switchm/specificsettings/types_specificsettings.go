package specificsettings

// API Version: v8_1

import (
	"encoding/json"
)

type DHCPOption struct {
	// Seq
	// Seq of Option
	Seq *int `json:"seq,omitempty"`

	// Type
	// Type of Option
	Type *string `json:"type,omitempty"`

	// Value
	// Value of Option
	Value *string `json:"value,omitempty"`
}

func NewDHCPOption() *DHCPOption {
	dHCPOptionType := new(DHCPOption)
	return dHCPOptionType
}

func NewDHCPOptionWithDefaults() *DHCPOption {
	dHCPOptionType := new(DHCPOption)
	return dHCPOptionType
}

type DHCPServer struct {
	// DefaultRouterIp
	// Default Router Ip
	DefaultRouterIp *string `json:"defaultRouterIp,omitempty"`

	DhcpOptions []*DHCPOption `json:"dhcpOptions,omitempty"`

	// ExcludedEnd
	// Excluded range end
	ExcludedEnd *string `json:"excludedEnd,omitempty"`

	// ExcludedStart
	// Excluded range start
	ExcludedStart *string `json:"excludedStart,omitempty"`

	// LeaseDays
	// Lease Days
	LeaseDays *int `json:"leaseDays,omitempty"`

	// LeaseHrs
	// Lease Hours
	LeaseHrs *int `json:"leaseHrs,omitempty"`

	// LeaseMins
	// Lease Mins
	LeaseMins *int `json:"leaseMins,omitempty"`

	// Network
	// Network/Mask
	Network *string `json:"network,omitempty"`

	// PoolName
	// Pool Name
	PoolName *string `json:"poolName,omitempty"`
}

func NewDHCPServer() *DHCPServer {
	dHCPServerType := new(DHCPServer)
	return dHCPServerType
}

func NewDHCPServerWithDefaults() *DHCPServer {
	dHCPServerType := new(DHCPServer)
	return dHCPServerType
}

type EmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *EmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = EmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *EmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewEmptyResult() *EmptyResult {
	emptyResultType := new(EmptyResult)
	return emptyResultType
}

func NewEmptyResultWithDefaults() *EmptyResult {
	emptyResultType := new(EmptyResult)
	return emptyResultType
}

type IdList struct {
	// Hostname
	// Hostname
	Hostname *string `json:"hostname,omitempty"`

	// Id
	// Id
	Id *string `json:"id,omitempty"`
}

func NewIdList() *IdList {
	idListType := new(IdList)
	return idListType
}

func NewIdListWithDefaults() *IdList {
	idListType := new(IdList)
	return idListType
}

type SpecificSettings struct {
	// CreatedTime
	// The create time of the Specific Settings
	CreatedTime *int `json:"createdTime,omitempty"`

	// DhcpServerEnabled
	// DHCP server enabled
	DhcpServerEnabled *bool `json:"dhcpServerEnabled,omitempty"`

	DhcpServers []*DHCPServer `json:"dhcpServers,omitempty"`

	// Hostname
	// Hostname
	Hostname *string `json:"hostname,omitempty"`

	// Id
	// Id
	Id *string `json:"id,omitempty"`

	// IgmpSnooping
	// IGMP snopping
	IgmpSnooping *string `json:"igmpSnooping,omitempty"`

	// JumboMode
	// JumboMode
	JumboMode *bool `json:"jumboMode,omitempty"`

	// UpdatedTime
	// The modify time of the Specific Settings
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSpecificSettings() *SpecificSettings {
	specificSettingsType := new(SpecificSettings)
	return specificSettingsType
}

func NewSpecificSettingsWithDefaults() *SpecificSettings {
	specificSettingsType := new(SpecificSettings)
	return specificSettingsType
}

type SpecificSettingsAllResult struct {
	// Extra
	// Any additional response data
	Extra *SpecificSettingsAllResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Specific Settings returned out of the complete Specific Settings list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Specific Settings after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*IdList `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Specific Settings count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Specific Settings count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSpecificSettingsAllResult() *SpecificSettingsAllResult {
	specificSettingsAllResultType := new(SpecificSettingsAllResult)
	return specificSettingsAllResultType
}

func NewSpecificSettingsAllResultWithDefaults() *SpecificSettingsAllResult {
	specificSettingsAllResultType := new(SpecificSettingsAllResult)
	return specificSettingsAllResultType
}

// SpecificSettingsAllResultExtraType
//
// Any additional response data
type SpecificSettingsAllResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SpecificSettingsAllResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SpecificSettingsAllResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SpecificSettingsAllResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSpecificSettingsAllResultExtraType() *SpecificSettingsAllResultExtraType {
	specificSettingsAllResultExtraTypeType := new(SpecificSettingsAllResultExtraType)
	return specificSettingsAllResultExtraTypeType
}

func NewSpecificSettingsAllResultExtraTypeWithDefaults() *SpecificSettingsAllResultExtraType {
	specificSettingsAllResultExtraTypeType := new(SpecificSettingsAllResultExtraType)
	return specificSettingsAllResultExtraTypeType
}

type UpdateSpecificSettings struct {
	// DhcpServerEnabled
	// DHCP server enabled
	DhcpServerEnabled *bool `json:"dhcpServerEnabled,omitempty"`

	DhcpServers []*DHCPServer `json:"dhcpServers,omitempty"`

	// Hostname
	// Hostname
	Hostname *string `json:"hostname,omitempty"`

	// IgmpSnooping
	// IGMP snopping
	IgmpSnooping *string `json:"igmpSnooping,omitempty"`

	// JumboMode
	// JumboMode
	JumboMode *bool `json:"jumboMode,omitempty"`
}

func NewUpdateSpecificSettings() *UpdateSpecificSettings {
	updateSpecificSettingsType := new(UpdateSpecificSettings)
	return updateSpecificSettingsType
}

func NewUpdateSpecificSettingsWithDefaults() *UpdateSpecificSettings {
	updateSpecificSettingsType := new(UpdateSpecificSettings)
	return updateSpecificSettingsType
}
