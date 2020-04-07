package vsz

// API Version: v9_0

type SwitchMDnsConfigCreateDnsConfig struct {
	// Dns
	// Constraints:
	//    - nullable
	Dns *SwitchMDnsConfigObject `json:"dns,omitempty"`

	// Id
	// Switch Group Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`
}

func NewSwitchMDnsConfigCreateDnsConfig() *SwitchMDnsConfigCreateDnsConfig {
	m := new(SwitchMDnsConfigCreateDnsConfig)
	return m
}

type SwitchMDnsConfig struct {
	// CreatedTime
	// The create time of the DNS Config
	// Constraints:
	//    - nullable
	CreatedTime *int `json:"createdTime,omitempty"`

	// Dns
	// Constraints:
	//    - nullable
	Dns *SwitchMDnsConfigObject `json:"dns,omitempty"`

	// UpdatedTime
	// The modify time of the DNS Config
	// Constraints:
	//    - nullable
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMDnsConfig() *SwitchMDnsConfig {
	m := new(SwitchMDnsConfig)
	return m
}

type SwitchMDnsConfigObject struct {
	// Ip
	// DNS Config IP
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`
}

func NewSwitchMDnsConfigObject() *SwitchMDnsConfigObject {
	m := new(SwitchMDnsConfigObject)
	return m
}

type SwitchMDnsConfigUpdateDnsConfig struct {
	// Dns
	// Constraints:
	//    - nullable
	Dns *SwitchMDnsConfigObject `json:"dns,omitempty"`
}

func NewSwitchMDnsConfigUpdateDnsConfig() *SwitchMDnsConfigUpdateDnsConfig {
	m := new(SwitchMDnsConfigUpdateDnsConfig)
	return m
}
