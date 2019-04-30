package zoneapmodel

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/types/apmodel"
	"github.com/myENA/ruckus-client/vsz/wsg/types/common"
)

type ApModel struct {
	CellularSettings *apmodel.CellularSettings `json:"cellularSettings,omitempty"`

	ExternalAntenna24 *apmodel.ExternalAntenna `json:"externalAntenna24,omitempty"`

	ExternalAntenna50 *apmodel.ExternalAntenna `json:"externalAntenna50,omitempty"`

	InternalHeaterEnabled *bool `json:"internalHeaterEnabled,omitempty"`

	LanPorts []*LanPortSetting `json:"lanPorts,omitempty"`

	LedMode *string `json:"ledMode,omitempty"`

	LedStatusEnabled *bool `json:"ledStatusEnabled,omitempty"`

	Lldp *apmodel.LldpSetting `json:"lldp,omitempty"`

	PoeModeSetting *string `json:"poeModeSetting,omitempty"`

	PoeOutPortEnabled *bool `json:"poeOutPortEnabled,omitempty"`

	// PoeTxChain
	// Option to use 1, 2 or 4 Tx chains while AP power source is 802.3af PoE
	PoeTxChain *int `json:"poeTxChain,omitempty"`

	// RadioBand
	// Band switch between 2.4GHz and 5GHz is provided in single radio AP ZF-7321, ZF-7321-U, and ZF-7441.
	RadioBand *string `json:"radioBand,omitempty"`

	UsbPowerEnable *bool `json:"usbPowerEnable,omitempty"`
}

type LanPortSetting struct {
	Enabled *bool `json:"enabled,omitempty"`

	EthPortProfile *common.GenericRef `json:"ethPortProfile,omitempty"`

	PortName *string `json:"portName,omitempty"`
}
