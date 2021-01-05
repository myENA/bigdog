package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
)

// WSGAPModel
//
// Definition: apmodel_apModel
type WSGAPModel struct {
	CellularSettings *WSGAPModelCellularSettings `json:"cellularSettings,omitempty"`

	ExternalAntenna24 *WSGAPModelExternalAntenna `json:"externalAntenna24,omitempty"`

	ExternalAntenna50 *WSGAPModelExternalAntenna `json:"externalAntenna50,omitempty"`

	// InternalHeaterEnabled
	// Constraints:
	//    - nullable
	InternalHeaterEnabled *bool `json:"internalHeaterEnabled,omitempty"`

	Lacp *WSGAPModelLacpSetting `json:"lacp,omitempty"`

	LanPorts []*WSGAPModelLanPortSetting `json:"lanPorts,omitempty"`

	// LedMode
	// Constraints:
	//    - nullable
	//    - oneof:[CableModem,AccessPoint,CableModem_AccessPoint,CableModem_AccessPoint_DEFAULT,ActiveSurgeProtector,ActiveSurgeProtector_ModemOnline_DEFAULT,Off]
	LedMode *string `json:"ledMode,omitempty"`

	// LedStatusEnabled
	// Constraints:
	//    - nullable
	LedStatusEnabled *bool `json:"ledStatusEnabled,omitempty"`

	Lldp *WSGAPModelLldpSetting `json:"lldp,omitempty"`

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
	RadioBand *string `json:"radioBand,omitempty"`

	// UsbPowerEnable
	// Constraints:
	//    - nullable
	UsbPowerEnable *bool `json:"usbPowerEnable,omitempty"`
}

func NewWSGAPModel() *WSGAPModel {
	m := new(WSGAPModel)
	return m
}

// WSGAPModelAuthenticatorAAAServer
//
// Definition: apmodel_authenticatorAAAServer
type WSGAPModelAuthenticatorAAAServer struct {
	// EnableUseSCGasProxy
	// Constraints:
	//    - required
	EnableUseSCGasProxy *bool `json:"enableUseSCGasProxy"`

	Server *WSGCommonGenericRef `json:"server,omitempty"`
}

func NewWSGAPModelAuthenticatorAAAServer() *WSGAPModelAuthenticatorAAAServer {
	m := new(WSGAPModelAuthenticatorAAAServer)
	return m
}

// WSGAPModelCellularSettings
//
// Definition: apmodel_cellularSettings
type WSGAPModelCellularSettings struct {
	// DataRoaming
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:1
	DataRoaming *int `json:"dataRoaming,omitempty"`

	// DataRoaming2
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:1
	DataRoaming2 *int `json:"dataRoaming2,omitempty"`

	// MobileAPName
	// Constraints:
	//    - nullable
	//    - max:100
	MobileAPName *string `json:"mobileAPName,omitempty"`

	// MobileAPName2
	// Constraints:
	//    - nullable
	//    - max:100
	MobileAPName2 *string `json:"mobileAPName2,omitempty"`

	// Select3g4g
	// Constraints:
	//    - required
	//    - min:0
	//    - max:2
	Select3g4g *int `json:"select3g4g"`

	// Select3g4g2
	// Constraints:
	//    - required
	//    - min:0
	//    - max:2
	Select3g4g2 *int `json:"select3g4g2"`

	// SimCardUsage
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:2
	SimCardUsage *int `json:"simCardUsage,omitempty"`

	// WanConnection
	// Constraints:
	//    - required
	//    - min:0
	//    - max:3
	WanConnection *int `json:"wanConnection"`

	// WanRecoveryTimer
	// Constraints:
	//    - required
	//    - min:10
	//    - max:300
	WanRecoveryTimer *int `json:"wanRecoveryTimer"`
}

func NewWSGAPModelCellularSettings() *WSGAPModelCellularSettings {
	m := new(WSGAPModelCellularSettings)
	return m
}

// WSGAPModelCommonAttribute
//
// Definition: apmodel_commonAttribute
type WSGAPModelCommonAttribute struct {
	AllowDfsCountry *string `json:"allowDfsCountry,omitempty"`

	CapabilityScore *int `json:"capabilityScore,omitempty"`

	CpuFrequency *int `json:"cpuFrequency,omitempty"`

	// HasCablemodem
	// Constraints:
	//    - nullable
	HasCablemodem *bool `json:"hasCablemodem,omitempty"`

	// HasGps
	// Constraints:
	//    - nullable
	HasGps *bool `json:"hasGps,omitempty"`

	// HasScanRadio
	// Constraints:
	//    - nullable
	HasScanRadio *bool `json:"hasScanRadio,omitempty"`

	// IsAllowDisableExtAnt
	// Constraints:
	//    - nullable
	IsAllowDisableExtAnt *bool `json:"isAllowDisableExtAnt,omitempty"`

	// IsDualRadio
	// Constraints:
	//    - nullable
	IsDualRadio *bool `json:"isDualRadio,omitempty"`

	// IsOutdoor
	// Constraints:
	//    - nullable
	IsOutdoor *bool `json:"isOutdoor,omitempty"`

	MaxChannelization5G *int `json:"maxChannelization5G,omitempty"`

	MaxChannelization24G *int `json:"maxChannelization24G,omitempty"`

	MaxClientsUpper *int `json:"maxClientsUpper,omitempty"`

	MaxWlanNum5G *int `json:"maxWlanNum5G,omitempty"`

	MaxWlanNum24G *int `json:"maxWlanNum24G,omitempty"`

	MeshRadioCaps *string `json:"meshRadioCaps,omitempty"`

	// NoAvc
	// Constraints:
	//    - nullable
	NoAvc *bool `json:"noAvc,omitempty"`

	// NoMesh
	// Constraints:
	//    - nullable
	NoMesh *bool `json:"noMesh,omitempty"`

	NonEditablePorts []int `json:"nonEditablePorts,omitempty"`

	NonVisiblePorts []int `json:"nonVisiblePorts,omitempty"`

	NumOfCores *int `json:"numOfCores,omitempty"`

	PoeModeCaps *string `json:"poeModeCaps,omitempty"`

	Ram *int `json:"ram,omitempty"`

	Reserved5GWlanForMesh *int `json:"reserved5GWlanForMesh,omitempty"`

	ScalingFactor *int `json:"scalingFactor,omitempty"`

	// Support11AC
	// Constraints:
	//    - nullable
	Support11AC *bool `json:"support11AC,omitempty"`

	// SupportAPUsbSoftwarePackage
	// Constraints:
	//    - nullable
	SupportAPUsbSoftwarePackage *bool `json:"supportAPUsbSoftwarePackage,omitempty"`

	// SupportBandSwitch
	// Constraints:
	//    - nullable
	SupportBandSwitch *bool `json:"supportBandSwitch,omitempty"`

	// SupportBonjour
	// Constraints:
	//    - nullable
	SupportBonjour *bool `json:"supportBonjour,omitempty"`

	// SupportChannelization160
	// Constraints:
	//    - nullable
	SupportChannelization160 *bool `json:"supportChannelization160,omitempty"`

	// SupportIpsec
	// Constraints:
	//    - nullable
	SupportIpsec *bool `json:"supportIpsec,omitempty"`

	// SupportLBS
	// Constraints:
	//    - nullable
	SupportLBS *bool `json:"supportLBS,omitempty"`

	// SupportResetCablemodem
	// Constraints:
	//    - nullable
	SupportResetCablemodem *bool `json:"supportResetCablemodem,omitempty"`
}

type WSGAPModelCommonAttributeAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPModelCommonAttribute
}

func newWSGAPModelCommonAttributeAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAPModelCommonAttributeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGAPModelCommonAttributeAPIResponse) Hydrate() error {
	r.Data = new(WSGAPModelCommonAttribute)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPModelCommonAttribute() *WSGAPModelCommonAttribute {
	m := new(WSGAPModelCommonAttribute)
	return m
}

// WSGAPModelExternalAntenna
//
// Definition: apmodel_externalAntenna
type WSGAPModelExternalAntenna struct {
	// ChainMask
	// Constraints:
	//    - nullable
	//    - oneof:[Two,Three]
	ChainMask *string `json:"chainMask,omitempty"`

	// Dbi
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:90
	Dbi *int `json:"dbi,omitempty"`

	// Enabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled"`
}

func NewWSGAPModelExternalAntenna() *WSGAPModelExternalAntenna {
	m := new(WSGAPModelExternalAntenna)
	return m
}

// WSGAPModelLacpSetting
//
// Definition: apmodel_lacpSetting
type WSGAPModelLacpSetting struct {
	Enabled *bool `json:"enabled,omitempty"`

	KeepApSetting *bool `json:"keepApSetting,omitempty"`
}

func NewWSGAPModelLacpSetting() *WSGAPModelLacpSetting {
	m := new(WSGAPModelLacpSetting)
	return m
}

// WSGAPModelLanPort8021X
//
// Definition: apmodel_lanPort8021X
type WSGAPModelLanPort8021X struct {
	Authenticator *WSGAPModelLanPortAuthenticator `json:"authenticator,omitempty"`

	Supplicant *WSGAPModelLanPortSupplicant `json:"supplicant,omitempty"`

	// Type
	// Constraints:
	//    - required
	//    - oneof:[Disable,Supplicant,PortBasedAuthenticator,MACBasedAuthenticator]
	Type *string `json:"type"`
}

func NewWSGAPModelLanPort8021X() *WSGAPModelLanPort8021X {
	m := new(WSGAPModelLanPort8021X)
	return m
}

// WSGAPModelLanPortAuthenticator
//
// Definition: apmodel_lanPortAuthenticator
type WSGAPModelLanPortAuthenticator struct {
	Accounting *WSGAPModelAuthenticatorAAAServer `json:"accounting,omitempty"`

	Authentication *WSGAPModelAuthenticatorAAAServer `json:"authentication,omitempty"`

	DisabledAccounting *bool `json:"disabledAccounting,omitempty"`

	// MacAuthByPassEnabled
	// Constraints:
	//    - required
	MacAuthByPassEnabled *bool `json:"macAuthByPassEnabled"`
}

func NewWSGAPModelLanPortAuthenticator() *WSGAPModelLanPortAuthenticator {
	m := new(WSGAPModelLanPortAuthenticator)
	return m
}

// WSGAPModelLanPortSetting
//
// Definition: apmodel_lanPortSetting
type WSGAPModelLanPortSetting struct {
	// Enabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled"`

	EthPortProfile *WSGCommonGenericRef `json:"ethPortProfile,omitempty"`

	// Members
	// Constraints:
	//    - nullable
	Members *string `json:"members,omitempty"`

	OverwriteVlanEnabled *bool `json:"overwriteVlanEnabled,omitempty"`

	// PortName
	// Constraints:
	//    - required
	//    - oneof:[LAN1,LAN2,LAN3,LAN4,LAN5]
	PortName *string `json:"portName"`

	// VlanUntagId
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:4094
	VlanUntagId *int `json:"vlanUntagId,omitempty"`
}

func NewWSGAPModelLanPortSetting() *WSGAPModelLanPortSetting {
	m := new(WSGAPModelLanPortSetting)
	return m
}

// WSGAPModelLanPortSupplicant
//
// Definition: apmodel_lanPortSupplicant
type WSGAPModelLanPortSupplicant struct {
	// Password
	// Constraints:
	//    - nullable
	//    - max:64
	Password *string `json:"password,omitempty"`

	// Type
	// Constraints:
	//    - required
	//    - oneof:[MACAddress,Custom]
	Type *string `json:"type"`

	// UserName
	// Constraints:
	//    - nullable
	//    - max:64
	UserName *string `json:"userName,omitempty"`
}

func NewWSGAPModelLanPortSupplicant() *WSGAPModelLanPortSupplicant {
	m := new(WSGAPModelLanPortSupplicant)
	return m
}

// WSGAPModelLldpSetting
//
// Definition: apmodel_lldpSetting
type WSGAPModelLldpSetting struct {
	// AdvertiseIntervalInSec
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:300
	AdvertiseIntervalInSec *int `json:"advertiseIntervalInSec,omitempty"`

	// Enabled
	// Constraints:
	//    - required
	//    - default:true
	Enabled *bool `json:"enabled"`

	// HoldTimeInSec
	// Constraints:
	//    - nullable
	//    - min:60
	//    - max:1200
	HoldTimeInSec *int `json:"holdTimeInSec,omitempty"`

	// ManagementIPTLVEnabled
	// Constraints:
	//    - nullable
	ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty"`
}

func NewWSGAPModelLldpSetting() *WSGAPModelLldpSetting {
	m := new(WSGAPModelLldpSetting)
	return m
}
