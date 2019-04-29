package flexivpn

// API Version: v8_0

type FlexiVpnGlobalSetting struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type FlexiVpnSetting struct {
	WLANID         *string `json:"wlanId,omitempty"`
	ZoneAffinityID *string `json:"zoneAffinityId,omitempty"`
}
