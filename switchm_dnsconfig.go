package ruckus

// API Version: v9_0

type SwitchMDNSConfigCreateDnsConfig struct {
	Dns *SwitchMDNSConfigObject `json:"dns,omitempty"`

	// Id
	// Switch Group Id
	Id *string `json:"id,omitempty"`
}

func NewSwitchMDNSConfigCreateDnsConfig() *SwitchMDNSConfigCreateDnsConfig {
	m := new(SwitchMDNSConfigCreateDnsConfig)
	return m
}

type SwitchMDNSConfig struct {
	// CreatedTime
	// The create time of the DNS Config
	CreatedTime *int `json:"createdTime,omitempty"`

	Dns *SwitchMDNSConfigObject `json:"dns,omitempty"`

	// UpdatedTime
	// The modify time of the DNS Config
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMDNSConfig() *SwitchMDNSConfig {
	m := new(SwitchMDNSConfig)
	return m
}

type SwitchMDNSConfigObject struct {
	// Ip
	// DNS Config IP
	Ip *string `json:"ip,omitempty"`
}

func NewSwitchMDNSConfigObject() *SwitchMDNSConfigObject {
	m := new(SwitchMDNSConfigObject)
	return m
}

type SwitchMDNSConfigUpdateDnsConfig struct {
	Dns *SwitchMDNSConfigObject `json:"dns,omitempty"`
}

func NewSwitchMDNSConfigUpdateDnsConfig() *SwitchMDNSConfigUpdateDnsConfig {
	m := new(SwitchMDNSConfigUpdateDnsConfig)
	return m
}
