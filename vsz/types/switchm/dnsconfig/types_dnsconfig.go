package dnsconfig

// API Version: v8_1

import (
	"encoding/json"
)

type CreateDNSConfig struct {
	DNS *DNSConfigObject `json:"dns,omitempty"`

	// ID
	// Switch Group Id
	ID *string `json:"id,omitempty"`
}

type DNSConfig struct {
	// CreatedTime
	// The create time of the DNS Config
	CreatedTime *int `json:"createdTime,omitempty"`

	DNS *DNSConfigObject `json:"dns,omitempty"`

	// UpdatedTime
	// The modify time of the DNS Config
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

type DNSConfigObject struct {
	// IP
	// DNS Config IP
	IP *string `json:"ip,omitempty"`
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

type UpdateDNSConfig struct {
	DNS *DNSConfigObject `json:"dns,omitempty"`
}
