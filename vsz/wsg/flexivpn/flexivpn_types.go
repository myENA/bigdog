package flexivpn

// API Version: v8_0

type FlexiVpnGlobalSetting struct {
	// Enabled
	// whether Flexi-VPN enabled in system
	Enabled *bool `json:"enabled,omitempty"`
}

type FlexiVpnSetting struct {
	// WLANID
	// Wlan ID
	WLANID *string `json:"wlanId,omitempty"`

	// ZoneAffinityID
	// Zone Affinity ID
	ZoneAffinityID *string `json:"zoneAffinityId,omitempty"`
}
