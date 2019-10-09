package calea

// API Version: v8_1

type CaleaCommonSettingRq struct {
	// CaleaServerIp
	// CARLEA Server IP
	CaleaServerIp *string `json:"caleaServerIp,omitempty"`

	// DcIp
	// DP IP in Data Center
	DcIp *string `json:"dc_ip,omitempty"`
}

type CaleaCommonSettingRsp struct {
	// CaleaServerIp
	// CARLEA Server IP
	CaleaServerIp *string `json:"caleaServerIp,omitempty"`

	// DcIp
	// DP IP in Data Center
	DcIp *string `json:"dc_ip,omitempty"`
}

type CaleaMacListRq struct {
	MacList []string `json:"macList,omitempty"`
}

type CaleaMacListRsp struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []string `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
