package calea

// API Version: v8_1

type CaleaCommonSettingRq struct {
	// CaleaServerIp
	// CARLEA Server IP
	CaleaServerIp *string `json:"caleaServerIp,omitempty"`

	// Dcip
	// DP IP in Data Center
	Dcip *string `json:"dc_ip,omitempty"`
}

func NewCaleaCommonSettingRq() *CaleaCommonSettingRq {
	caleaCommonSettingRqType := new(CaleaCommonSettingRq)
	return caleaCommonSettingRqType
}

func NewDefaultCaleaCommonSettingRq() *CaleaCommonSettingRq {
	caleaCommonSettingRqType := new(CaleaCommonSettingRq)
	return caleaCommonSettingRqType
}

type CaleaCommonSettingRsp struct {
	// CaleaServerIp
	// CARLEA Server IP
	CaleaServerIp *string `json:"caleaServerIp,omitempty"`

	// Dcip
	// DP IP in Data Center
	Dcip *string `json:"dc_ip,omitempty"`
}

func NewCaleaCommonSettingRsp() *CaleaCommonSettingRsp {
	caleaCommonSettingRspType := new(CaleaCommonSettingRsp)
	return caleaCommonSettingRspType
}

func NewDefaultCaleaCommonSettingRsp() *CaleaCommonSettingRsp {
	caleaCommonSettingRspType := new(CaleaCommonSettingRsp)
	return caleaCommonSettingRspType
}

type CaleaMacListRq struct {
	MacList []string `json:"macList,omitempty"`
}

func NewCaleaMacListRq() *CaleaMacListRq {
	caleaMacListRqType := new(CaleaMacListRq)
	return caleaMacListRqType
}

func NewDefaultCaleaMacListRq() *CaleaMacListRq {
	caleaMacListRqType := new(CaleaMacListRq)
	return caleaMacListRqType
}

type CaleaMacListRsp struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []string `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewCaleaMacListRsp() *CaleaMacListRsp {
	caleaMacListRspType := new(CaleaMacListRsp)
	return caleaMacListRspType
}

func NewDefaultCaleaMacListRsp() *CaleaMacListRsp {
	caleaMacListRspType := new(CaleaMacListRsp)
	return caleaMacListRspType
}
