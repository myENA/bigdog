package domain

// API Version: v8_0

type CreateDomain struct {
	Description     *string `json:"description,omitempty"`
	DomainType      *string `json:"domainType,omitempty"`
	Name            *string `json:"name,omitempty"`
	ParentDomainID  *string `json:"parentDomainId,omitempty"`
	ZeroTouchStatus *bool   `json:"zeroTouchStatus,omitempty"`
}

type DomainConfiguration struct {
	AdministratorCount *int    `json:"administratorCount,omitempty"`
	ApCount            *int    `json:"apCount,omitempty"`
	CreateDatetime     *string `json:"createDatetime,omitempty"`
	CreatedBy          *string `json:"createdBy,omitempty"`
	Description        *string `json:"description,omitempty"`
	DomainType         *string `json:"domainType,omitempty"`
	ID                 *string `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	ParentDomainID     *string `json:"parentDomainId,omitempty"`
	SubDomainCount     *int    `json:"subDomainCount,omitempty"`
	ZeroTouchStatus    *bool   `json:"zeroTouchStatus,omitempty"`
	ZoneCount          *int    `json:"zoneCount,omitempty"`
}

type DomainList struct {
	FirstIndex *int                   `json:"firstIndex,omitempty"`
	HasMore    *bool                  `json:"hasMore,omitempty"`
	List       []*DomainConfiguration `json:"list,omitempty"`
	TotalCount *int                   `json:"totalCount,omitempty"`
}

type ModifyDomain struct {
	Description     *string `json:"description,omitempty"`
	DomainType      *string `json:"domainType,omitempty"`
	Name            *string `json:"name,omitempty"`
	ParentDomainID  *string `json:"parentDomainId,omitempty"`
	ZeroTouchStatus *bool   `json:"zeroTouchStatus,omitempty"`
}
