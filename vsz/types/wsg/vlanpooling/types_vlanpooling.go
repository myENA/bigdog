package vlanpooling

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	Algo *string `json:"algo,omitempty" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty" validate:"required"`
}

type DeleteBulkVlanPooling struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type ModifyVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	Algo *string `json:"algo,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

type VlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	Algo *string `json:"algo,omitempty"`

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

type VlanPoolingList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*VlanPoolingListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type VlanPoolingListType struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	Algo *string `json:"algo,omitempty"`

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
