package vlanpooling

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - required
	//    - default:'MAC_HASH'
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo" validate:"required,oneof=MAC_HASH"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Pool
	// VLANs of the VLAN pooling profile
	// Constraints:
	//    - required
	Pool *string `json:"pool" validate:"required"`
}

func NewCreateVlanPooling() *CreateVlanPooling {
	createVlanPoolingType := new(CreateVlanPooling)
	return createVlanPoolingType
}

func NewCreateVlanPoolingWithDefaults() *CreateVlanPooling {
	createVlanPoolingType := new(CreateVlanPooling)
	algoField := `MAC_HASH`
	createVlanPoolingType.Algo = &algoField
	return createVlanPoolingType
}

type DeleteBulkVlanPooling struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulkVlanPooling() *DeleteBulkVlanPooling {
	deleteBulkVlanPoolingType := new(DeleteBulkVlanPooling)
	return deleteBulkVlanPoolingType
}

func NewDeleteBulkVlanPoolingWithDefaults() *DeleteBulkVlanPooling {
	deleteBulkVlanPoolingType := new(DeleteBulkVlanPooling)
	return deleteBulkVlanPoolingType
}

type ModifyVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - nullable
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty" validate:"omitempty,oneof=MAC_HASH"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

func NewModifyVlanPooling() *ModifyVlanPooling {
	modifyVlanPoolingType := new(ModifyVlanPooling)
	return modifyVlanPoolingType
}

func NewModifyVlanPoolingWithDefaults() *ModifyVlanPooling {
	modifyVlanPoolingType := new(ModifyVlanPooling)
	return modifyVlanPoolingType
}

type VlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - nullable
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty" validate:"omitempty,oneof=MAC_HASH"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Identifier of the domain which the VLAN pooling profile belongs to
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the VLAN pooling profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

func NewVlanPooling() *VlanPooling {
	vlanPoolingType := new(VlanPooling)
	return vlanPoolingType
}

func NewVlanPoolingWithDefaults() *VlanPooling {
	vlanPoolingType := new(VlanPooling)
	return vlanPoolingType
}

type VlanPoolingList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*VlanPoolingListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewVlanPoolingList() *VlanPoolingList {
	vlanPoolingListType := new(VlanPoolingList)
	return vlanPoolingListType
}

func NewVlanPoolingListWithDefaults() *VlanPoolingList {
	vlanPoolingListType := new(VlanPoolingList)
	return vlanPoolingListType
}

type VlanPoolingListType struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - nullable
	//    - default:'MAC_HASH'
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty" validate:"omitempty,oneof=MAC_HASH"`

	// Description
	// Description of the service
	Description *string `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

func NewVlanPoolingListType() *VlanPoolingListType {
	vlanPoolingListTypeType := new(VlanPoolingListType)
	return vlanPoolingListTypeType
}

func NewVlanPoolingListTypeWithDefaults() *VlanPoolingListType {
	vlanPoolingListTypeType := new(VlanPoolingListType)
	algoField := `MAC_HASH`
	vlanPoolingListTypeType.Algo = &algoField
	return vlanPoolingListTypeType
}
