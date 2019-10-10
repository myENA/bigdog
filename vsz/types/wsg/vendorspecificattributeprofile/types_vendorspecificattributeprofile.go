package vendorspecificattributeprofile

// API Version: v8_1

import (
	"encoding/json"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateResult interface{}

type DeleteBulk struct {
	IdList common.IdList `json:"idList,omitempty"`
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

type Get struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	Attributes []*VendorSpecificAttribute `json:"attributes,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// ZoneId
	// Zone Id
	ZoneId *string `json:"zoneId,omitempty"`
}

type List struct {
	// FirstIndex
	// Index of the first profile returned out of the profile list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more profiles after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Information list of the vendor specific attribute profile
	List []*ListType `json:"list,omitempty"`

	// TotalCount
	// Total number of the profiles
	TotalCount *int `json:"totalCount,omitempty"`
}

type ListType struct {
	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type Persist struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	Attributes []*VendorSpecificAttribute `json:"attributes,omitempty" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`
}

type QueryCriteriaResult struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*Get `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type VendorSpecificAttribute struct {
	// KeyId
	// Key ID of vendor specific attribute
	KeyId *int `json:"keyId,omitempty" validate:"required"`

	// SupportedRadiusProtocol
	// The radius protocol to which this given vendor specific attribute will attach
	SupportedRadiusProtocol *string `json:"supportedRadiusProtocol,omitempty" validate:"required"`

	// Type
	// Type of vendor specific attribute
	Type *string `json:"type,omitempty" validate:"required"`

	// Value
	// Value of vendor specific attribute
	Value *string `json:"value,omitempty" validate:"required"`

	// VendorId
	// Vendor ID of vendor specific attribute
	VendorId *int `json:"vendorId,omitempty" validate:"required"`
}
