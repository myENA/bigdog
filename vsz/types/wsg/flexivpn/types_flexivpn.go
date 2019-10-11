package flexivpn

// API Version: v8_1

type FlexiVpnSetting struct {
	// ZoneAffinityId
	// Zone Affinity ID
	// Constraints:
	//    - required
	ZoneAffinityId *string `json:"zoneAffinityId" validate:"required"`
}

func NewFlexiVpnSetting() *FlexiVpnSetting {
	flexiVpnSettingType := new(FlexiVpnSetting)
	return flexiVpnSettingType
}

func NewFlexiVpnSettingWithDefaults() *FlexiVpnSetting {
	flexiVpnSettingType := new(FlexiVpnSetting)
	return flexiVpnSettingType
}
