package ruckus

// API Version: v9_0

type SwitchMSwitchDNSConfigCreateDnsConfig struct {
	Dns *SwitchMSwitchDNSConfigDnsConfigObject `json:"dns,omitempty"`

	// Id
	// Switch Group Id
	Id *string `json:"id,omitempty"`
}

func NewSwitchMSwitchDNSConfigCreateDnsConfig() *SwitchMSwitchDNSConfigCreateDnsConfig {
	m := new(SwitchMSwitchDNSConfigCreateDnsConfig)
	return m
}

type SwitchMSwitchDNSConfigDnsConfig struct {
	// CreatedTime
	// The create time of the DNS Config
	CreatedTime *int `json:"createdTime,omitempty"`

	Dns *SwitchMSwitchDNSConfigDnsConfigObject `json:"dns,omitempty"`

	// UpdatedTime
	// The modify time of the DNS Config
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMSwitchDNSConfigDnsConfig() *SwitchMSwitchDNSConfigDnsConfig {
	m := new(SwitchMSwitchDNSConfigDnsConfig)
	return m
}

type SwitchMSwitchDNSConfigDnsConfigObject struct {
	// Ip
	// DNS Config IP
	Ip *string `json:"ip,omitempty"`
}

func NewSwitchMSwitchDNSConfigDnsConfigObject() *SwitchMSwitchDNSConfigDnsConfigObject {
	m := new(SwitchMSwitchDNSConfigDnsConfigObject)
	return m
}

type SwitchMSwitchDNSConfigUpdateDnsConfig struct {
	Dns *SwitchMSwitchDNSConfigDnsConfigObject `json:"dns,omitempty"`
}

func NewSwitchMSwitchDNSConfigUpdateDnsConfig() *SwitchMSwitchDNSConfigUpdateDnsConfig {
	m := new(SwitchMSwitchDNSConfigUpdateDnsConfig)
	return m
}
