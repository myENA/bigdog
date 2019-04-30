package vlanpooling

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/common"
)

type CreateVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	Algo *string `json:"algo,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// identifier of the domain
	DomainID *string `json:"domainId,omitempty"`

	Name *string `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

type DeleteBulkVlanPooling struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type ModifyVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	Algo *string `json:"algo,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// identifier of the domain
	DomainID *string `json:"domainId,omitempty"`

	Name *string `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

type VlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	Algo *string `json:"algo,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Identifier of the domain which the VLAN pooling profile belongs to
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Identifier of the VLAN pooling profile
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

type VlanPoolingList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

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

	// DomainID
	// identifier of the domain
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Identifier of the service
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}
