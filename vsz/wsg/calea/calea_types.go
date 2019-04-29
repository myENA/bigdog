package calea

// API Version: v8_0

type CaleaCommonSettingRq struct {
	CaleaServerIP *string `json:"caleaServerIp,omitempty"`
	DcIP          *string `json:"dc_ip,omitempty"`
}

type CaleaCommonSettingRsp struct {
	CaleaServerIP *string `json:"caleaServerIp,omitempty"`
	DcIP          *string `json:"dc_ip,omitempty"`
}

type CaleaMacListRq struct {
	MacList []string `json:"macList,omitempty"`
}

type CaleaMacListRsp struct {
	FirstIndex *int     `json:"firstIndex,omitempty"`
	HasMore    *bool    `json:"hasMore,omitempty"`
	List       []string `json:"list,omitempty"`
	TotalCount *int     `json:"totalCount,omitempty"`
}
