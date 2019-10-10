package zoneapmodel

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ApModel struct {
	CellularSettings *apmodel.CellularSettings `json:"cellularSettings,omitempty"`

	ExternalAntenna24 *apmodel.ExternalAntenna `json:"externalAntenna24,omitempty"`

	ExternalAntenna50 *apmodel.ExternalAntenna `json:"externalAntenna50,omitempty"`

	InternalHeaterEnabled *bool `json:"internalHeaterEnabled,omitempty"`

	Lacp *apmodel.LacpSetting `json:"lacp,omitempty"`

	LanPorts []*LanPortSetting `json:"lanPorts,omitempty"`

	LedMode *string `json:"ledMode,omitempty" validate:"omitempty,oneof=CableModem AccessPoint CableModem_AccessPoint CableModem_AccessPoint_DEFAULT ActiveSurgeProtector ActiveSurgeProtector_ModemOnline_DEFAULT Off"`

	LedStatusEnabled *bool `json:"ledStatusEnabled,omitempty"`

	Lldp *apmodel.LldpSetting `json:"lldp,omitempty"`

	PoeModeSetting *string `json:"poeModeSetting,omitempty" validate:"omitempty,oneof=Auto _802_3af _802_3at _802_3atPlus"`

	PoeOutPortEnabled *bool `json:"poeOutPortEnabled,omitempty"`

	// PoeTxChain
	// Option to use 1, 2 or 4 Tx chains while AP power source is 802.3af PoE
	PoeTxChain *int `json:"poeTxChain,omitempty"`

	// RadioBand
	// Band switch between 2.4GHz and 5GHz is provided in single radio AP ZF-7321, ZF-7321-U, and ZF-7441.
	RadioBand *string `json:"radioBand,omitempty" validate:"omitempty,oneof=2.4GHz 5GHz"`

	UsbPowerEnable *bool `json:"usbPowerEnable,omitempty"`
}

type LanPortSetting struct {
	Enabled *bool `json:"enabled" validate:"required"`

	EthPortProfile *common.GenericRef `json:"ethPortProfile,omitempty"`

	PortName *string `json:"portName" validate:"required,oneof=LAN1 LAN2 LAN3 LAN4 LAN5"`
}
