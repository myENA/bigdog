package vsz

// API Version: v9_0

type SwitchMDnsConfigCreateDnsConfig struct {
	Dns *SwitchMDnsConfigObject `json:"dns,omitempty"`

	// Id
	// Switch Group Id
	Id *string `json:"id,omitempty"`
}

func NewSwitchMDnsConfigCreateDnsConfig() *SwitchMDnsConfigCreateDnsConfig {
	m := new(SwitchMDnsConfigCreateDnsConfig)
	return m
}

type SwitchMDnsConfig struct {
	// CreatedTime
	// The create time of the DNS Config
	CreatedTime *int `json:"createdTime,omitempty"`

	Dns *SwitchMDnsConfigObject `json:"dns,omitempty"`

	// UpdatedTime
	// The modify time of the DNS Config
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMDnsConfig() *SwitchMDnsConfig {
	m := new(SwitchMDnsConfig)
	return m
}

type SwitchMDnsConfigObject struct {
	// Ip
	// DNS Config IP
	Ip *string `json:"ip,omitempty"`
}

func NewSwitchMDnsConfigObject() *SwitchMDnsConfigObject {
	m := new(SwitchMDnsConfigObject)
	return m
}

type SwitchMDnsConfigUpdateDnsConfig struct {
	Dns *SwitchMDnsConfigObject `json:"dns,omitempty"`
}

func NewSwitchMDnsConfigUpdateDnsConfig() *SwitchMDnsConfigUpdateDnsConfig {
	m := new(SwitchMDnsConfigUpdateDnsConfig)
	return m
}
