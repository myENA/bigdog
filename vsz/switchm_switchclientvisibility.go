package vsz

// API Version: v9_0

type SwitchMSwitchClientVisibilityClient struct {
	// ClientAuthType
	// Constraints:
	//    - nullable
	ClientAuthType *string `json:"clientAuthType,omitempty"`

	// ClientDesc
	// Constraints:
	//    - nullable
	ClientDesc *string `json:"clientDesc,omitempty"`

	// ClientIpv4Addr
	// Constraints:
	//    - nullable
	ClientIpv4Addr *string `json:"clientIpv4Addr,omitempty"`

	// ClientIpv6Addr
	// Constraints:
	//    - nullable
	ClientIpv6Addr *string `json:"clientIpv6Addr,omitempty"`

	// ClientLastSeen
	// Constraints:
	//    - nullable
	ClientLastSeen *string `json:"clientLastSeen,omitempty"`

	// ClientMac
	// Constraints:
	//    - nullable
	ClientMac *string `json:"clientMac,omitempty"`

	// ClientStatus
	// Constraints:
	//    - nullable
	ClientStatus *string `json:"clientStatus,omitempty"`

	// ClientType
	// Constraints:
	//    - nullable
	ClientType *string `json:"clientType,omitempty"`

	// ClientUserName
	// Constraints:
	//    - nullable
	ClientUserName *string `json:"clientUserName,omitempty"`

	// ClientVlan
	// Constraints:
	//    - nullable
	ClientVlan *string `json:"clientVlan,omitempty"`

	// PastAuthHistory
	// Constraints:
	//    - nullable
	PastAuthHistory *string `json:"pastAuthHistory,omitempty"`

	// SwitchName
	// Constraints:
	//    - nullable
	SwitchName *string `json:"switchName,omitempty"`

	// SwitchPort
	// Constraints:
	//    - nullable
	SwitchPort *int `json:"switchPort,omitempty"`
}

func NewSwitchMSwitchClientVisibilityClient() *SwitchMSwitchClientVisibilityClient {
	m := new(SwitchMSwitchClientVisibilityClient)
	return m
}

type SwitchMSwitchClientVisibilityList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchClientVisibilityClient `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchClientVisibilityList() *SwitchMSwitchClientVisibilityList {
	m := new(SwitchMSwitchClientVisibilityList)
	return m
}
