package flexivpn

// API Version: v8_0

type FlexiVpnGlobalSetting struct {
	// Enabled
	// whether Flexi-VPN enabled in system
	Enabled *bool `json:"enabled,omitempty"`
}

type FlexiVpnSetting struct {
	// WlanId
	// Wlan ID
	WlanId *string `json:"wlanId,omitempty"`

	// ZoneAffinityId
	// Zone Affinity ID
	ZoneAffinityId *string `json:"zoneAffinityId,omitempty"`
}
