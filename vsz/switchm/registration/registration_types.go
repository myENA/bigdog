package registration

// API Version: v8_0

type ClientObjectID struct {
	ExtraValues map[string]interface{} `json:"extraValues,omitempty"`
	ID          *string                `json:"id,omitempty"`
	Label       *string                `json:"label,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

type CreateResult struct {
	Data     *Data                  `json:"data,omitempty"`
	Error    *Error                 `json:"error,omitempty"`
	Extra    interface{}            `json:"extra,omitempty"`
	MetaData map[string]interface{} `json:"metaData,omitempty"`
	Success  *bool                  `json:"success,omitempty"`
}

type DeleteMultipleResult struct {
	Data     *Data                  `json:"data,omitempty"`
	Error    *Error                 `json:"error,omitempty"`
	Extra    interface{}            `json:"extra,omitempty"`
	MetaData map[string]interface{} `json:"metaData,omitempty"`
	Success  *bool                  `json:"success,omitempty"`
}

type DeleteResult struct {
	Data     *Data                  `json:"data,omitempty"`
	Error    *Error                 `json:"error,omitempty"`
	Extra    interface{}            `json:"extra,omitempty"`
	MetaData map[string]interface{} `json:"metaData,omitempty"`
	Success  *bool                  `json:"success,omitempty"`
}

type ErrorObject struct {
	List    []string `json:"list,omitempty"`
	Message *string  `json:"message,omitempty"`
	MsgKey  *string  `json:"msgKey,omitempty"`
}

type List struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type ModifyResult struct {
	Data     *Data                  `json:"data,omitempty"`
	Error    *Error                 `json:"error,omitempty"`
	Extra    interface{}            `json:"extra,omitempty"`
	MetaData map[string]interface{} `json:"metaData,omitempty"`
	Success  *bool                  `json:"success,omitempty"`
}

type RegistrationRule struct {
	CreateDatetime *int    `json:"createDatetime,omitempty"`
	CreatorID      *string `json:"creatorId,omitempty"`
	CreatorName    *string `json:"creatorName,omitempty"`
	Description    *string `json:"description,omitempty"`
	GroupName      *string `json:"groupName,omitempty"`
	ID             *string `json:"id,omitempty"`
	IPFrom         *string `json:"ipFrom,omitempty"`
	IPRange        *string `json:"ipRange,omitempty"`
	IPTo           *string `json:"ipTo,omitempty"`
	Label          *string `json:"label,omitempty"`
	ModelNumber    *string `json:"modelNumber,omitempty"`
	Network        *string `json:"network,omitempty"`
	Rank           *int    `json:"rank,omitempty"`
	SubnetMask     *string `json:"subnetMask,omitempty"`
	SwitchGroupID  *string `json:"switchGroupId,omitempty"`
	Type           *string `json:"type,omitempty"`
}

type RuleQueryResultList struct {
	Data     *Data                  `json:"data,omitempty"`
	Error    *Error                 `json:"error,omitempty"`
	Extra    interface{}            `json:"extra,omitempty"`
	MetaData map[string]interface{} `json:"metaData,omitempty"`
	Success  *bool                  `json:"success,omitempty"`
}
