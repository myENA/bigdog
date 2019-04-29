package sci

// API Version: v8_0

type CreateSciProfile struct {
	SciPassword   *string `json:"sciPassword,omitempty"`
	SciProfile    *string `json:"sciProfile,omitempty"`
	SciServerHost *string `json:"sciServerHost,omitempty"`
	SciServerPort *string `json:"sciServerPort,omitempty"`
	SciSystemID   *string `json:"sciSystemId,omitempty"`
	SciUser       *string `json:"sciUser,omitempty"`
}

type DeleteSciProfile struct {
	ID *string `json:"id,omitempty"`
}

type DeleteSciProfileList struct {
	List []*List `json:"list,omitempty"`
}

type ModifyEventCode struct {
	SciAcceptedEventCodes []int `json:"sciAcceptedEventCodes,omitempty"`
}

type ModifySciEnabled struct {
	SciEnabled *bool `json:"sciEnabled,omitempty"`
}

type ModifySciPriorityList struct {
	List []*List `json:"list,omitempty"`
}

type ModifySciPriorityListType struct {
	ID          *string `json:"id,omitempty"`
	SciPriority *int    `json:"sciPriority,omitempty"`
	SciProfile  *string `json:"sciProfile,omitempty"`
}

type ModifySciProfile struct {
	ID            *string `json:"id,omitempty"`
	SciPassword   *string `json:"sciPassword,omitempty"`
	SciProfile    *string `json:"sciProfile,omitempty"`
	SciServerHost *string `json:"sciServerHost,omitempty"`
	SciServerPort *string `json:"sciServerPort,omitempty"`
	SciSystemID   *string `json:"sciSystemId,omitempty"`
	SciUser       *string `json:"sciUser,omitempty"`
}

type SciEventCode struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type SciEventCodeListType struct {
	Code *int    `json:"code,omitempty"`
	Type *string `json:"type,omitempty"`
}

type SciProfile struct {
	ID            *string `json:"id,omitempty"`
	SciPassword   *string `json:"sciPassword,omitempty"`
	SciPriority   *int    `json:"sciPriority,omitempty"`
	SciProfile    *string `json:"sciProfile,omitempty"`
	SciServerHost *string `json:"sciServerHost,omitempty"`
	SciServerPort *string `json:"sciServerPort,omitempty"`
	SciSystemID   *string `json:"sciSystemId,omitempty"`
	SciUser       *string `json:"sciUser,omitempty"`
}

type SciProfileListExtraType struct {
	SciEnabled *bool `json:"sciEnabled,omitempty"`
}
