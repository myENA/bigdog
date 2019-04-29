package group

// API Version: v8_0

type AuditID struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ClientObjectID struct {
	ExtraValues map[string]interface{} `json:"extraValues,omitempty"`
	ID          *string                `json:"id,omitempty"`
	Label       *string                `json:"label,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

type DeleteSwitchGroupResult struct {
	*AuditID
}

type ErrorObject struct {
	List    []string `json:"list,omitempty"`
	Message *string  `json:"message,omitempty"`
	MsgKey  *string  `json:"msgKey,omitempty"`
}

type GroupsByIdsQueryResultList struct {
	Data     *Data                  `json:"data,omitempty"`
	Error    *Error                 `json:"error,omitempty"`
	Extra    interface{}            `json:"extra,omitempty"`
	MetaData map[string]interface{} `json:"metaData,omitempty"`
	Success  *bool                  `json:"success,omitempty"`
}

type QueryResultList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type SwitchGroup struct {
	CreateDatetime        *int                   `json:"createDatetime,omitempty"`
	CreatorID             *string                `json:"creatorId,omitempty"`
	Description           *string                `json:"description,omitempty"`
	DomainID              *string                `json:"domainId,omitempty"`
	ID                    *string                `json:"id,omitempty"`
	LevelOne              *bool                  `json:"levelOne,omitempty"`
	LevelTwo              *bool                  `json:"levelTwo,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	SampledInstant        map[string]interface{} `json:"sampledInstant,omitempty"`
	SwitchGroupLevelOneID *string                `json:"switchGroupLevelOneId,omitempty"`
	SwitchGroupLevelTwoID *string                `json:"switchGroupLevelTwoId,omitempty"`
	TenantID              *string                `json:"tenantId,omitempty"`
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
