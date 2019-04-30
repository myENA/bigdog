package zoneapmodel

// API Version: v8_0

type ApModel struct {
	CellularSettings      *apmodel.CellularSettings `json:"cellularSettings,omitempty"`
	ExternalAntenna24     *apmodel.ExternalAntenna  `json:"externalAntenna24,omitempty"`
	ExternalAntenna50     *apmodel.ExternalAntenna  `json:"externalAntenna50,omitempty"`
	InternalHeaterEnabled *bool                     `json:"internalHeaterEnabled,omitempty"`
	LanPorts              []*LanPortSetting         `json:"lanPorts,omitempty"`
	LedMode               *string                   `json:"ledMode,omitempty"`
	LedStatusEnabled      *bool                     `json:"ledStatusEnabled,omitempty"`
	Lldp                  *apmodel.LldpSetting      `json:"lldp,omitempty"`
	PoeModeSetting        *string                   `json:"poeModeSetting,omitempty"`
	PoeOutPortEnabled     *bool                     `json:"poeOutPortEnabled,omitempty"`
	PoeTxChain            *int                      `json:"poeTxChain,omitempty"`
	RadioBand             *string                   `json:"radioBand,omitempty"`
	UsbPowerEnable        *bool                     `json:"usbPowerEnable,omitempty"`
}

type LanPortSetting struct {
	Enabled        *bool              `json:"enabled,omitempty"`
	EthPortProfile *common.GenericRef `json:"ethPortProfile,omitempty"`
	PortName       *string            `json:"portName,omitempty"`
}
