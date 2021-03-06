package bigdog

// API Version: v9_1

// SwitchMSwitchClientVisibilityClient
//
// Definition: switchClientVisibility_client
type SwitchMSwitchClientVisibilityClient struct {
	ClientAuthType *string `json:"clientAuthType,omitempty"`

	ClientDesc *string `json:"clientDesc,omitempty"`

	ClientIpv4Addr *string `json:"clientIpv4Addr,omitempty"`

	ClientIpv6Addr *string `json:"clientIpv6Addr,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	ClientStatus *string `json:"clientStatus,omitempty"`

	ClientType *string `json:"clientType,omitempty"`

	ClientUserName *string `json:"clientUserName,omitempty"`

	ClientVlan *string `json:"clientVlan,omitempty"`

	PastAuthHistory *string `json:"pastAuthHistory,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`

	SwitchPort *int `json:"switchPort,omitempty"`

	UpdatedTime *string `json:"updatedTime,omitempty"`
}

func NewSwitchMSwitchClientVisibilityClient() *SwitchMSwitchClientVisibilityClient {
	m := new(SwitchMSwitchClientVisibilityClient)
	return m
}

// SwitchMSwitchClientVisibilityList
//
// Definition: switchClientVisibility_list
type SwitchMSwitchClientVisibilityList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchClientVisibilityClient `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchClientVisibilityList() *SwitchMSwitchClientVisibilityList {
	m := new(SwitchMSwitchClientVisibilityList)
	return m
}
