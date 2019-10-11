package dnsconfig

// API Version: v8_1

import (
	"encoding/json"
)

type CreateDnsConfig struct {
	Dns *DnsConfigObject `json:"dns,omitempty"`

	// Id
	// Switch Group Id
	Id *string `json:"id,omitempty"`
}

func NewCreateDnsConfig() *CreateDnsConfig {
	createDnsConfigType := new(CreateDnsConfig)
	return createDnsConfigType
}

func NewCreateDnsConfigWithDefaults() *CreateDnsConfig {
	createDnsConfigType := new(CreateDnsConfig)
	return createDnsConfigType
}

type DnsConfig struct {
	// CreatedTime
	// The create time of the DNS Config
	CreatedTime *int `json:"createdTime,omitempty"`

	Dns *DnsConfigObject `json:"dns,omitempty"`

	// UpdatedTime
	// The modify time of the DNS Config
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewDnsConfig() *DnsConfig {
	dnsConfigType := new(DnsConfig)
	return dnsConfigType
}

func NewDnsConfigWithDefaults() *DnsConfig {
	dnsConfigType := new(DnsConfig)
	return dnsConfigType
}

type DnsConfigObject struct {
	// Ip
	// DNS Config IP
	Ip *string `json:"ip,omitempty"`
}

func NewDnsConfigObject() *DnsConfigObject {
	dnsConfigObjectType := new(DnsConfigObject)
	return dnsConfigObjectType
}

func NewDnsConfigObjectWithDefaults() *DnsConfigObject {
	dnsConfigObjectType := new(DnsConfigObject)
	return dnsConfigObjectType
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

type UpdateDnsConfig struct {
	Dns *DnsConfigObject `json:"dns,omitempty"`
}

func NewUpdateDnsConfig() *UpdateDnsConfig {
	updateDnsConfigType := new(UpdateDnsConfig)
	return updateDnsConfigType
}

func NewUpdateDnsConfigWithDefaults() *UpdateDnsConfig {
	updateDnsConfigType := new(UpdateDnsConfig)
	return updateDnsConfigType
}
