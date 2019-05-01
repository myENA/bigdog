package registration

// API Version: v8_0

type ClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	ExtraValues map[string]interface{} `json:"extraValues,omitempty"`

	// Id
	// Identifier of the client
	Id *string `json:"id,omitempty"`

	// Label
	// Label of the client
	Label *string `json:"label,omitempty"`

	// Type
	// Type of the client
	Type *string `json:"type,omitempty"`
}

type CreateResult struct {
	Data *ClientObjectID `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule create result
	MetaData map[string]interface{} `json:"metaData,omitempty"`

	// Success
	// Create result success or not
	Success *bool `json:"success,omitempty"`
}

type DeleteMultipleResult struct {
	Data *List `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Matadata of delete multiple rules result
	MetaData map[string]interface{} `json:"metaData,omitempty"`

	// Success
	// Delete multiple result success or not
	Success *bool `json:"success,omitempty"`
}

type DeleteResult struct {
	Data *ClientObjectID `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule delete result
	MetaData map[string]interface{} `json:"metaData,omitempty"`

	// Success
	// Delete result success or not
	Success *bool `json:"success,omitempty"`
}

type ErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

type List struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first registration rule returned out of the complete registration rule list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more registration rule after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*RegistrationRule `json:"list,omitempty"`

	// RawDataTotalCount
	// Registration rule list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Registration rule list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type ModifyResult struct {
	Data *ClientObjectID `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Matadata of rule modify result
	MetaData map[string]interface{} `json:"metaData,omitempty"`

	// Success
	// Modify result success or not
	Success *bool `json:"success,omitempty"`
}

type RegistrationRule struct {
	// CreateDatetime
	// Create datetime of the registration rule
	CreateDatetime *int `json:"createDatetime,omitempty"`

	// CreatorId
	// Creator Id of the registration rule
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorName
	// Creator name of the registration rule
	CreatorName *string `json:"creatorName,omitempty"`

	// Description
	// Description of the registration rule
	Description *string `json:"description,omitempty"`

	// GroupName
	// Switch group name of the registration rule
	GroupName *string `json:"groupName,omitempty"`

	// Id
	// Identifier of the registration rule
	Id *string `json:"id,omitempty"`

	// IpFrom
	// Start IP range of the registration rule
	IpFrom *string `json:"ipFrom,omitempty"`

	// IpRange
	// IP range of the registration rule
	IpRange *string `json:"ipRange,omitempty"`

	// IpTo
	// End IP range of the registration rule
	IpTo *string `json:"ipTo,omitempty"`

	// Label
	// Lable of the registration rule
	Label *string `json:"label,omitempty"`

	// ModelNumber
	// Switch Model number of the registration rule
	ModelNumber *string `json:"modelNumber,omitempty"`

	// Network
	// Network of the registration rule
	Network *string `json:"network,omitempty"`

	// Rank
	// Rank of the registration rule
	Rank *int `json:"rank,omitempty"`

	// SubnetMask
	// Subnet mask of the registration rule
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchGroupId
	// Switch group Id of the registration rule
	SwitchGroupId *string `json:"switchGroupId,omitempty"`

	// Type
	// Type of the registration rule
	Type *string `json:"type,omitempty"`
}

type RuleQueryResultList struct {
	Data *List `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule query result
	MetaData map[string]interface{} `json:"metaData,omitempty"`

	// Success
	// Rule query result success or not
	Success *bool `json:"success,omitempty"`
}

type RuleUUIDs []string
