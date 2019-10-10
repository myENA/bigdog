package domain

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateDomain struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainType
	// domain type
	DomainType *string `json:"domainType,omitempty" validate:"oneof=PARTNER MVNO REGULAR"`

	Name *common.NormalName `json:"name" validate:"required"`

	// ParentDomainId
	// parent domain id
	ParentDomainId *string `json:"parentDomainId,omitempty"`

	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`
}

type DomainConfiguration struct {
	// AdministratorCount
	// # of Subdomains
	AdministratorCount *int `json:"administratorCount,omitempty"`

	// ApCount
	// # of Subdomains
	ApCount *int `json:"apCount,omitempty"`

	// CreateDatetime
	// Created by
	CreateDatetime *string `json:"createDatetime,omitempty"`

	// CreatedBy
	// Created by
	CreatedBy *string `json:"createdBy,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainType
	// domain type
	DomainType *string `json:"domainType,omitempty"`

	// Id
	// Identifier of the domain
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// ParentDomainId
	// Parent Domain Id
	ParentDomainId *string `json:"parentDomainId,omitempty"`

	// SubDomainCount
	// # of Subdomains
	SubDomainCount *int `json:"subDomainCount,omitempty"`

	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`

	// ZoneCount
	// # of Zones
	ZoneCount *int `json:"zoneCount,omitempty"`
}

type DomainList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DomainConfiguration `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ModifyDomain struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainType
	// domain type
	DomainType *string `json:"domainType,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// ParentDomainId
	// parent domain id
	ParentDomainId *string `json:"parentDomainId,omitempty"`

	// ZeroTouchStatus
	// Zero Touch enable/disable
	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`
}
