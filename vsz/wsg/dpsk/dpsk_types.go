package dpsk

// API Version: v8_0

type BatchGenUnbound struct {
	Amount         *int     `json:"amount,omitempty"`
	GroupDpsk      *bool    `json:"groupDpsk,omitempty"`
	PassphraseList []string `json:"passphraseList,omitempty"`
	UserName       *string  `json:"userName,omitempty"`
	UserRoleID     *string  `json:"userRoleId,omitempty"`
	VlanID         *int     `json:"vlanId,omitempty"`
}

type DeleteDpskResult struct {
	ResultCount *int `json:"resultCount,omitempty"`
}

type DeleteDPSKs struct {
	IDList []string `json:"idList,omitempty"`
}

type DeleteExpiredDpskConfig struct {
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty"`
}

type DpskInfo []*DpskInfoType

type DpskInfoType struct {
	CreationDateTime   *float64 `json:"creationDateTime,omitempty"`
	ExpirationDateTime *string  `json:"expirationDateTime,omitempty"`
	GroupDpsk          *bool    `json:"groupDpsk,omitempty"`
	ID                 *string  `json:"id,omitempty"`
	MacAddress         *string  `json:"macAddress,omitempty"`
	Passphrase         *string  `json:"passphrase,omitempty"`
	UserName           *string  `json:"userName,omitempty"`
	UserRoleID         *string  `json:"userRoleId,omitempty"`
	VlanID             *int     `json:"vlanId,omitempty"`
	WLANID             *string  `json:"wlanId,omitempty"`
}

type DpskQueryList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*DpskQueryListType `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type DpskQueryListType struct {
	CreateDateTime      *float64 `json:"createDateTime,omitempty"`
	DomainID            *string  `json:"domainId,omitempty"`
	ExpirationStartTime *float64 `json:"expirationStartTime,omitempty"`
	ExpirationTime      *float64 `json:"expirationTime,omitempty"`
	Expired             *bool    `json:"expired,omitempty"`
	Group               *bool    `json:"group,omitempty"`
	Key                 *string  `json:"key,omitempty"`
	TenantID            *string  `json:"tenantId,omitempty"`
	TTL                 *float64 `json:"ttl,omitempty"`
	UeMac               *string  `json:"ueMac,omitempty"`
	UserName            *string  `json:"userName,omitempty"`
	UserRoleID          *string  `json:"userRoleId,omitempty"`
	VlanID              *int     `json:"vlanId,omitempty"`
	WLANID              *string  `json:"wlanId,omitempty"`
	ZoneID              *string  `json:"zoneId,omitempty"`
}

type GetDpskEnabledWlans struct {
	FirstIndex *int                           `json:"firstIndex,omitempty"`
	HasMore    *bool                          `json:"hasMore,omitempty"`
	List       []*GetDpskEnabledWlansListType `json:"list,omitempty"`
	TotalCount *int                           `json:"totalCount,omitempty"`
}

type GetDpskEnabledWlansListType struct {
	Ssid     *string `json:"ssid,omitempty"`
	WLANID   *string `json:"wlanId,omitempty"`
	WLANName *string `json:"wlanName,omitempty"`
}

type GetDpskInfoList struct {
	FirstIndex *int     `json:"firstIndex,omitempty"`
	HasMore    *bool    `json:"hasMore,omitempty"`
	List       DpskInfo `json:"list,omitempty"`
	TotalCount *int     `json:"totalCount,omitempty"`
}

type GetDpskResult struct {
	DpskInfoList DpskInfo `json:"dpskInfoList,omitempty"`
	ResultCount  *int     `json:"resultCount,omitempty"`
}

type ModifyDeleteExpiredDpsk struct {
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty"`
}

type UpdateDpsk struct {
	UserName *string `json:"userName,omitempty"`
}

type WLANDpskSetting struct {
	DpskEnabled  *bool   `json:"dpskEnabled,omitempty"`
	DpskFromType *string `json:"dpskFromType,omitempty"`
	DpskType     *string `json:"dpskType,omitempty"`
	Expiration   *string `json:"expiration,omitempty"`
	Length       *int    `json:"length,omitempty"`
}

type WLANExternalDpsk struct {
	AuthService *WLANExternalDpskAuthServiceType `json:"authService,omitempty"`
	Enabled     *bool                            `json:"enabled,omitempty"`
	Encryption  *WLANExternalDpskEncryptionType  `json:"encryption,omitempty"`
}

type WLANExternalDpskAuthServiceType struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type WLANExternalDpskEncryptionType struct {
	Algorithm *string `json:"algorithm,omitempty"`
	Method    *string `json:"method,omitempty"`
}
