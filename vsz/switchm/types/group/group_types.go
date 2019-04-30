package group

// API Version: v8_0

type AuditID struct {
	// ID
	// Audit Id
	ID *string `json:"id,omitempty"`

	// Name
	// Audit name
	Name *string `json:"name,omitempty"`
}

type ClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	ExtraValues map[string]interface{} `json:"extraValues,omitempty"`

	// ID
	// Identifier of the client
	ID *string `json:"id,omitempty"`

	// Label
	// Label of the client
	Label *string `json:"label,omitempty"`

	// Type
	// Type of the client
	Type *string `json:"type,omitempty"`
}

type DeleteSwitchGroupResult struct {
	*AuditID
}

type ErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

type GroupsByIdsQueryResultList struct {
	Data *QueryResultList `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Metadata of query result list
	MetaData map[string]interface{} `json:"metaData,omitempty"`

	// Success
	// Query result success or not
	Success *bool `json:"success,omitempty"`
}

type QueryResultList struct {
	// Extra
	// Extra information for query result list
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first query result returned out of the complete query result list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more query result after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*ClientObjectID `json:"list,omitempty"`

	// RawDataTotalCount
	// Query result count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total query result count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchGroup struct {
	// CreateDatetime
	// Create datetime of the switch group
	CreateDatetime *int `json:"createDatetime,omitempty"`

	// CreatorID
	// Creator Id of the switch group
	CreatorID *string `json:"creatorId,omitempty"`

	// Description
	// Description of the switch group
	Description *string `json:"description,omitempty"`

	// DomainID
	// Identifier of the management domain to which the switch group belong
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Identifier of the switch group
	ID *string `json:"id,omitempty"`

	// LevelOne
	// Level one  of the switch group
	LevelOne *bool `json:"levelOne,omitempty"`

	// LevelTwo
	// Level two of the switch group
	LevelTwo *bool `json:"levelTwo,omitempty"`

	// Name
	// Name of the switch group
	Name *string `json:"name,omitempty"`

	// SampledInstant
	// Sampled instant of the switch group
	SampledInstant map[string]interface{} `json:"sampledInstant,omitempty"`

	// SwitchGroupLevelOneID
	// Level one Id of the switch group
	SwitchGroupLevelOneID *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoID
	// Level two Id of the switch group
	SwitchGroupLevelTwoID *string `json:"switchGroupLevelTwoId,omitempty"`

	// TenantID
	// Tenant Id of the switch group
	TenantID *string `json:"tenantId,omitempty"`
}

type SwitchGroupQueryResult struct {
	*SwitchGroup
}

type UpdateSwitchGroup struct {
	*SwitchGroup
}

type UpdateSwitchGroupResult struct {
	*AuditID
}
