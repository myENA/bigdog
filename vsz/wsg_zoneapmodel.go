package vsz

// API Version: v9_0

type WSGZoneAPModelApModel struct {
	// CellularSettings
	// Constraints:
	//    - nullable
	CellularSettings *WSGAPModelCellularSettings `json:"cellularSettings,omitempty"`

	// ExternalAntenna24
	// Constraints:
	//    - nullable
	ExternalAntenna24 *WSGAPModelExternalAntenna `json:"externalAntenna24,omitempty"`

	// ExternalAntenna50
	// Constraints:
	//    - nullable
	ExternalAntenna50 *WSGAPModelExternalAntenna `json:"externalAntenna50,omitempty"`

	// InternalHeaterEnabled
	// Constraints:
	//    - nullable
	InternalHeaterEnabled *bool `json:"internalHeaterEnabled,omitempty"`

	// Lacp
	// Constraints:
	//    - nullable
	Lacp *WSGAPModelLacpSetting `json:"lacp,omitempty"`

	// LanPorts
	// Constraints:
	//    - nullable
	LanPorts []*WSGZoneAPModelLanPortSetting `json:"lanPorts,omitempty" validate:"omitempty,dive"`

	// LedMode
	// Constraints:
	//    - nullable
	//    - oneof:[CableModem,AccessPoint,CableModem_AccessPoint,CableModem_AccessPoint_DEFAULT,ActiveSurgeProtector,ActiveSurgeProtector_ModemOnline_DEFAULT,Off]
	LedMode *string `json:"ledMode,omitempty" validate:"omitempty,oneof=CableModem AccessPoint CableModem_AccessPoint CableModem_AccessPoint_DEFAULT ActiveSurgeProtector ActiveSurgeProtector_ModemOnline_DEFAULT Off"`

	// LedStatusEnabled
	// Constraints:
	//    - nullable
	LedStatusEnabled *bool `json:"ledStatusEnabled,omitempty"`

	// Lldp
	// Constraints:
	//    - nullable
	Lldp *WSGAPModelLldpSetting `json:"lldp,omitempty"`

	// PoeModeSetting
	// Constraints:
	//    - nullable
	PoeModeSetting *WSGCommonPoeModeSetting `json:"poeModeSetting,omitempty"`

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
	RadioBand *string `json:"radioBand,omitempty" validate:"omitempty,oneof=2.4GHz 5GHz"`

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
	Enabled *bool `json:"enabled" validate:"required"`

	// EthPortProfile
	// Constraints:
	//    - nullable
	EthPortProfile *WSGCommonGenericRef `json:"ethPortProfile,omitempty"`

	// PortName
	// Constraints:
	//    - required
	//    - oneof:[LAN1,LAN2,LAN3,LAN4,LAN5]
	PortName *string `json:"portName" validate:"required,oneof=LAN1 LAN2 LAN3 LAN4 LAN5"`
}

func NewWSGZoneAPModelLanPortSetting() *WSGZoneAPModelLanPortSetting {
	m := new(WSGZoneAPModelLanPortSetting)
	return m
}
