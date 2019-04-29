package vlanpooling

// API Version: v8_0

type CreateVlanPooling struct {
	Algo        *string `json:"algo,omitempty"`
	Description *string `json:"description,omitempty"`
	DomainID    *string `json:"domainId,omitempty"`
	Name        *string `json:"name,omitempty"`
	Pool        *string `json:"pool,omitempty"`
}

type DeleteBulkVlanPooling struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type ModifyVlanPooling struct {
	Algo        *string `json:"algo,omitempty"`
	Description *string `json:"description,omitempty"`
	DomainID    *string `json:"domainId,omitempty"`
	Name        *string `json:"name,omitempty"`
	Pool        *string `json:"pool,omitempty"`
}

type VlanPooling struct {
	Algo        *string `json:"algo,omitempty"`
	Description *string `json:"description,omitempty"`
	DomainID    *string `json:"domainId,omitempty"`
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Pool        *string `json:"pool,omitempty"`
}

type VlanPoolingList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type VlanPoolingListType struct {
	Algo        *string `json:"algo,omitempty"`
	Description *string `json:"description,omitempty"`
	DomainID    *string `json:"domainId,omitempty"`
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Pool        *string `json:"pool,omitempty"`
}
