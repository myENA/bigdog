package vsz

// API Version: v9_0

type SwitchMDnsconfigCreateDnsConfig struct {
	Dns *SwitchMDnsconfigObject `json:"dns,omitempty"`

	// Id
	// Switch Group Id
	Id *string `json:"id,omitempty"`
}

func NewSwitchMDnsconfigCreateDnsConfig() *SwitchMDnsconfigCreateDnsConfig {
	m := new(SwitchMDnsconfigCreateDnsConfig)
	return m
}

type SwitchMDnsconfig struct {
	// CreatedTime
	// The create time of the DNS Config
	CreatedTime *int `json:"createdTime,omitempty"`

	Dns *SwitchMDnsconfigObject `json:"dns,omitempty"`

	// UpdatedTime
	// The modify time of the DNS Config
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMDnsconfig() *SwitchMDnsconfig {
	m := new(SwitchMDnsconfig)
	return m
}

type SwitchMDnsconfigObject struct {
	// Ip
	// DNS Config IP
	Ip *string `json:"ip,omitempty"`
}

func NewSwitchMDnsconfigObject() *SwitchMDnsconfigObject {
	m := new(SwitchMDnsconfigObject)
	return m
}

type SwitchMDnsconfigUpdateDnsConfig struct {
	Dns *SwitchMDnsconfigObject `json:"dns,omitempty"`
}

func NewSwitchMDnsconfigUpdateDnsConfig() *SwitchMDnsconfigUpdateDnsConfig {
	m := new(SwitchMDnsconfigUpdateDnsConfig)
	return m
}
