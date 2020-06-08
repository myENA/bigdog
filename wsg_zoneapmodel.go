package ruckus

// API Version: v9_0

type WSGZoneAPModelApModel struct {
	CellularSettings *WSGZoneAPModelApModel `json:"cellularSettings,omitempty"`

	ExternalAntenna24 *WSGZoneAPModelApModel `json:"externalAntenna24,omitempty"`

	ExternalAntenna50 *WSGZoneAPModelApModel `json:"externalAntenna50,omitempty"`

	// InternalHeaterEnabled
	// Constraints:
	//    - nullable
	InternalHeaterEnabled *bool `json:"internalHeaterEnabled,omitempty"`

	Lacp *WSGZoneAPModelApModel `json:"lacp,omitempty"`

	LanPorts []*WSGZoneAPModelApModel `json:"lanPorts,omitempty"`

	// LedMode
	// Constraints:
	//    - nullable
	//    - oneof:[CableModem,AccessPoint,CableModem_AccessPoint,CableModem_AccessPoint_DEFAULT,ActiveSurgeProtector,ActiveSurgeProtector_ModemOnline_DEFAULT,Off]
	LedMode *string `json:"ledMode,omitempty"`

	// LedStatusEnabled
	// Constraints:
	//    - nullable
	LedStatusEnabled *bool `json:"ledStatusEnabled,omitempty"`

	Lldp *WSGZoneAPModelApModel `json:"lldp,omitempty"`

	PoeModeSetting *WSGZoneAPModelApModel `json:"poeModeSetting,omitempty"`

	// PoeOutPortEnabled
	// Constraints:
	//    - nullable
	PoeOutPortEnabled *bool `json:"poeOutPortEnabled,omitempty"`

	// PoeTxChain
	// Option to use 1, 2 or 4 Tx chains while AP power source is 802.3af PoE
	// Constraints:
	//    - nullable
	//    - default:2
	PoeTxChain *int `json:"poeTxChain,omitempty"`

	// RadioBand
	// Band switch between 2.4GHz and 5GHz is provided in single radio AP ZF-7321, ZF-7321-U, and ZF-7441.
	// Constraints:
	//    - nullable
	//    - oneof:[2.4GHz,5GHz]
	RadioBand *string `json:"radioBand,omitempty"`

	// UsbPowerEnable
	// Constraints:
	//    - nullable
	UsbPowerEnable *bool `json:"usbPowerEnable,omitempty"`
}

func NewWSGZoneAPModelApModel() *WSGZoneAPModelApModel {
	m := new(WSGZoneAPModelApModel)
	return m
}

type WSGZoneAPModelLanPortSetting struct {
	// Enabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled"`

	EthPortProfile *WSGZoneAPModelLanPortSetting `json:"ethPortProfile,omitempty"`

	// PortName
	// Constraints:
	//    - required
	//    - oneof:[LAN1,LAN2,LAN3,LAN4,LAN5]
	PortName *string `json:"portName"`
}

func NewWSGZoneAPModelLanPortSetting() *WSGZoneAPModelLanPortSetting {
	m := new(WSGZoneAPModelLanPortSetting)
	return m
}
