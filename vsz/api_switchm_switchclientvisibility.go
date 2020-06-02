package vsz

// API Version: v9_0

type SwitchMSwitchclientvisibilityClient struct {
	ClientAuthType *string `json:"clientAuthType,omitempty"`

	ClientDesc *string `json:"clientDesc,omitempty"`

	ClientIpv4Addr *string `json:"clientIpv4Addr,omitempty"`

	ClientIpv6Addr *string `json:"clientIpv6Addr,omitempty"`

	ClientLastSeen *string `json:"clientLastSeen,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	ClientStatus *string `json:"clientStatus,omitempty"`

	ClientType *string `json:"clientType,omitempty"`

	ClientUserName *string `json:"clientUserName,omitempty"`

	ClientVlan *string `json:"clientVlan,omitempty"`

	PastAuthHistory *string `json:"pastAuthHistory,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`

	SwitchPort *int `json:"switchPort,omitempty"`
}

func NewSwitchMSwitchclientvisibilityClient() *SwitchMSwitchclientvisibilityClient {
	m := new(SwitchMSwitchclientvisibilityClient)
	return m
}

type SwitchMSwitchclientvisibilityList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchclientvisibilityClient `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchclientvisibilityList() *SwitchMSwitchclientvisibilityList {
	m := new(SwitchMSwitchclientvisibilityList)
	return m
}
