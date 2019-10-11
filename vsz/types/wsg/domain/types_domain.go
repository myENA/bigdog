package domain

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateDomain struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainType
	// domain type
	// Constraints:
	//    - nullable
	//    - default:'REGULAR'
	//    - oneof:[PARTNER,MVNO,REGULAR]
	DomainType *string `json:"domainType,omitempty" validate:"omitempty,oneof=PARTNER MVNO REGULAR"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// ParentDomainId
	// parent domain id
	ParentDomainId *string `json:"parentDomainId,omitempty"`

	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`
}

func NewCreateDomain() *CreateDomain {
	createDomainType := new(CreateDomain)
	return createDomainType
}

func NewDefaultCreateDomain() *CreateDomain {
	createDomainType := new(CreateDomain)
	domainTypeField := `REGULAR`
	createDomainType.DomainType = &domainTypeField
	zeroTouchStatusField := false
	createDomainType.ZeroTouchStatus = &zeroTouchStatusField
	return createDomainType
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

func NewDomainConfiguration() *DomainConfiguration {
	domainConfigurationType := new(DomainConfiguration)
	return domainConfigurationType
}

func NewDefaultDomainConfiguration() *DomainConfiguration {
	domainConfigurationType := new(DomainConfiguration)
	return domainConfigurationType
}

type DomainList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DomainConfiguration `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDomainList() *DomainList {
	domainListType := new(DomainList)
	return domainListType
}

func NewDefaultDomainList() *DomainList {
	domainListType := new(DomainList)
	return domainListType
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

func NewModifyDomain() *ModifyDomain {
	modifyDomainType := new(ModifyDomain)
	return modifyDomainType
}

func NewDefaultModifyDomain() *ModifyDomain {
	modifyDomainType := new(ModifyDomain)
	return modifyDomainType
}
