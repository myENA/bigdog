package vendorspecificattributeprofile

// API Version: v8_1

import (
	"encoding/json"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateResult interface{}

func NewCreateResult() *CreateResult {
	createResultType := new(CreateResult)
	return createResultType
}

func NewCreateResultWithDefaults() *CreateResult {
	createResultType := new(CreateResult)
	return createResultType
}

type DeleteBulk struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulk() *DeleteBulk {
	deleteBulkType := new(DeleteBulk)
	return deleteBulkType
}

func NewDeleteBulkWithDefaults() *DeleteBulk {
	deleteBulkType := new(DeleteBulk)
	return deleteBulkType
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

func NewGet() *Get {
	getType := new(Get)
	return getType
}

func NewGetWithDefaults() *Get {
	getType := new(Get)
	return getType
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

func NewList() *List {
	listType := new(List)
	return listType
}

func NewListWithDefaults() *List {
	listType := new(List)
	return listType
}

type ListType struct {
	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewListType() *ListType {
	listTypeType := new(ListType)
	return listTypeType
}

func NewListTypeWithDefaults() *ListType {
	listTypeType := new(ListType)
	return listTypeType
}

type Persist struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	// Constraints:
	//    - required
	Attributes []*VendorSpecificAttribute `json:"attributes" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`
}

func NewPersist() *Persist {
	persistType := new(Persist)
	return persistType
}

func NewPersistWithDefaults() *Persist {
	persistType := new(Persist)
	return persistType
}

type QueryCriteriaResult struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*Get `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewQueryCriteriaResult() *QueryCriteriaResult {
	queryCriteriaResultType := new(QueryCriteriaResult)
	return queryCriteriaResultType
}

func NewQueryCriteriaResultWithDefaults() *QueryCriteriaResult {
	queryCriteriaResultType := new(QueryCriteriaResult)
	return queryCriteriaResultType
}

type VendorSpecificAttribute struct {
	// KeyId
	// Key ID of vendor specific attribute
	// Constraints:
	//    - required
	KeyId *int `json:"keyId" validate:"required"`

	// SupportedRadiusProtocol
	// The radius protocol to which this given vendor specific attribute will attach
	// Constraints:
	//    - required
	//    - oneof:[ACCOUNTING,AUTHENTICATION,BOTH_AUTHENTICATION_AND_ACCOUNTING]
	SupportedRadiusProtocol *string `json:"supportedRadiusProtocol" validate:"required,oneof=ACCOUNTING AUTHENTICATION BOTH_AUTHENTICATION_AND_ACCOUNTING"`

	// Type
	// Type of vendor specific attribute
	// Constraints:
	//    - required
	//    - oneof:[INTEGER,STRING]
	Type *string `json:"type" validate:"required,oneof=INTEGER STRING"`

	// Value
	// Value of vendor specific attribute
	// Constraints:
	//    - required
	Value *string `json:"value" validate:"required"`

	// VendorId
	// Vendor ID of vendor specific attribute
	// Constraints:
	//    - required
	VendorId *int `json:"vendorId" validate:"required"`
}

func NewVendorSpecificAttribute() *VendorSpecificAttribute {
	vendorSpecificAttributeType := new(VendorSpecificAttribute)
	return vendorSpecificAttributeType
}

func NewVendorSpecificAttributeWithDefaults() *VendorSpecificAttribute {
	vendorSpecificAttributeType := new(VendorSpecificAttribute)
	return vendorSpecificAttributeType
}
