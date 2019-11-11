package vsz

// API Version: v8_1

import (
	"encoding/json"
)

type SwitchMDnsConfigCreateDnsConfig struct {
	Dns *SwitchMDnsConfigObject `json:"dns,omitempty"`

	// Id
	// Switch Group Id
	Id *string `json:"id,omitempty"`
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

type SwitchMDnsConfigObject struct {
	// Ip
	// DNS Config IP
	Ip *string `json:"ip,omitempty"`
}

type SwitchMDnsConfigEmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMDnsConfigEmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMDnsConfigEmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMDnsConfigEmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type SwitchMDnsConfigUpdateDnsConfig struct {
	Dns *SwitchMDnsConfigObject `json:"dns,omitempty"`
}
