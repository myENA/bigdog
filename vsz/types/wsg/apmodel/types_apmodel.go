package apmodel

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ApModel struct {
	CellularSettings *CellularSettings `json:"cellularSettings,omitempty"`

	ExternalAntenna24 *ExternalAntenna `json:"externalAntenna24,omitempty"`

	ExternalAntenna50 *ExternalAntenna `json:"externalAntenna50,omitempty"`

	InternalHeaterEnabled *bool `json:"internalHeaterEnabled,omitempty"`

	Lacp *LacpSetting `json:"lacp,omitempty"`

	LanPorts []*LanPortSetting `json:"lanPorts,omitempty"`

	// LedMode
	// Constraints:
	//    - nullable
	//    - oneof:[CableModem,AccessPoint,CableModem_AccessPoint,CableModem_AccessPoint_DEFAULT,ActiveSurgeProtector,ActiveSurgeProtector_ModemOnline_DEFAULT,Off]
	LedMode *string `json:"ledMode,omitempty" validate:"omitempty,oneof=CableModem AccessPoint CableModem_AccessPoint CableModem_AccessPoint_DEFAULT ActiveSurgeProtector ActiveSurgeProtector_ModemOnline_DEFAULT Off"`

	LedStatusEnabled *bool `json:"ledStatusEnabled,omitempty"`

	Lldp *LldpSetting `json:"lldp,omitempty"`

	// PoeModeSetting
	// Constraints:
	//    - nullable
	//    - oneof:[Auto,_802_3af,_802_3at,_802_3atPlus]
	PoeModeSetting *string `json:"poeModeSetting,omitempty" validate:"omitempty,oneof=Auto _802_3af _802_3at _802_3atPlus"`

	PoeOutPortEnabled *bool `json:"poeOutPortEnabled,omitempty"`

	// PoeTxChain
	// Option to use 1, 2 or 4 Tx chains while AP power source is 802.3af PoE
	PoeTxChain *int `json:"poeTxChain,omitempty"`

	// RadioBand
	// Band switch between 2.4GHz and 5GHz is provided in single radio AP ZF-7321, ZF-7321-U, and ZF-7441.
	// Constraints:
	//    - nullable
	//    - oneof:[2.4GHz,5GHz]
	RadioBand *string `json:"radioBand,omitempty" validate:"omitempty,oneof=2.4GHz 5GHz"`

	UsbPowerEnable *bool `json:"usbPowerEnable,omitempty"`
}

func NewApModel() *ApModel {
	apModelType := new(ApModel)
	return apModelType
}

func NewDefaultApModel() *ApModel {
	apModelType := new(ApModel)
	poeTxChainField := 2
	apModelType.PoeTxChain = &poeTxChainField
	return apModelType
}

type AuthenticatorAAAServer struct {
	// EnableUseSCGasProxy
	// Constraints:
	//    - required
	EnableUseSCGasProxy *bool `json:"enableUseSCGasProxy" validate:"required"`

	Server *common.GenericRef `json:"server,omitempty"`
}

func NewAuthenticatorAAAServer() *AuthenticatorAAAServer {
	authenticatorAAAServerType := new(AuthenticatorAAAServer)
	return authenticatorAAAServerType
}

func NewDefaultAuthenticatorAAAServer() *AuthenticatorAAAServer {
	authenticatorAAAServerType := new(AuthenticatorAAAServer)
	return authenticatorAAAServerType
}

type CellularSettings struct {
	// DataRoaming
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:1
	DataRoaming *int `json:"dataRoaming,omitempty" validate:"omitempty,gte=0,lte=1"`

	// DataRoaming2
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:1
	DataRoaming2 *int `json:"dataRoaming2,omitempty" validate:"omitempty,gte=0,lte=1"`

	// MobileAPName
	// Constraints:
	//    - nullable
	//    - max:100
	MobileAPName *string `json:"mobileAPName,omitempty" validate:"omitempty,max=100"`

	// MobileAPName2
	// Constraints:
	//    - nullable
	//    - max:100
	MobileAPName2 *string `json:"mobileAPName2,omitempty" validate:"omitempty,max=100"`

	// Select3g4g
	// Constraints:
	//    - required
	//    - min:0
	//    - max:2
	Select3g4g *int `json:"select3g4g" validate:"required,gte=0,lte=2"`

	// Select3g4g2
	// Constraints:
	//    - required
	//    - min:0
	//    - max:2
	Select3g4g2 *int `json:"select3g4g2" validate:"required,gte=0,lte=2"`

	// SimCardUsage
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:2
	SimCardUsage *int `json:"simCardUsage,omitempty" validate:"omitempty,gte=0,lte=2"`

	// WanConnection
	// Constraints:
	//    - required
	//    - min:0
	//    - max:3
	WanConnection *int `json:"wanConnection" validate:"required,gte=0,lte=3"`

	// WanRecoveryTimer
	// Constraints:
	//    - required
	//    - min:10
	//    - max:300
	WanRecoveryTimer *int `json:"wanRecoveryTimer" validate:"required,gte=10,lte=300"`
}

func NewCellularSettings() *CellularSettings {
	cellularSettingsType := new(CellularSettings)
	return cellularSettingsType
}

func NewDefaultCellularSettings() *CellularSettings {
	cellularSettingsType := new(CellularSettings)
	return cellularSettingsType
}

type CommonAttribute struct {
	AllowDfsCountry *string `json:"allowDfsCountry,omitempty"`

	CapabilityScore *float64 `json:"capabilityScore,omitempty"`

	CpuFrequency *int `json:"cpuFrequency,omitempty"`

	HasCablemodem *bool `json:"hasCablemodem,omitempty"`

	HasGps *bool `json:"hasGps,omitempty"`

	HasScanRadio *bool `json:"hasScanRadio,omitempty"`

	IsAllowDisableExtAnt *bool `json:"isAllowDisableExtAnt,omitempty"`

	IsDualRadio *bool `json:"isDualRadio,omitempty"`

	IsOutdoor *bool `json:"isOutdoor,omitempty"`

	MaxChannelization5G *int `json:"maxChannelization5G,omitempty"`

	MaxChannelization24G *int `json:"maxChannelization24G,omitempty"`

	MaxClientsUpper *int `json:"maxClientsUpper,omitempty"`

	MaxWlanNum5G *int `json:"maxWlanNum5G,omitempty"`

	MaxWlanNum24G *int `json:"maxWlanNum24G,omitempty"`

	MeshRadioCaps *string `json:"meshRadioCaps,omitempty"`

	NoAvc *bool `json:"noAvc,omitempty"`

	NoMesh *bool `json:"noMesh,omitempty"`

	NonEditablePorts []int `json:"nonEditablePorts,omitempty"`

	NonVisiblePorts []int `json:"nonVisiblePorts,omitempty"`

	NumOfCores *int `json:"numOfCores,omitempty"`

	PoeModeCaps *string `json:"poeModeCaps,omitempty"`

	Ram *int `json:"ram,omitempty"`

	Reserved5GWlanForMesh *int `json:"reserved5GWlanForMesh,omitempty"`

	ScalingFactor *int `json:"scalingFactor,omitempty"`

	Support11AC *bool `json:"support11AC,omitempty"`

	SupportAPUsbSoftwarePackage *bool `json:"supportAPUsbSoftwarePackage,omitempty"`

	SupportBandSwitch *bool `json:"supportBandSwitch,omitempty"`

	SupportBonjour *bool `json:"supportBonjour,omitempty"`

	SupportChannelization160 *bool `json:"supportChannelization160,omitempty"`

	SupportIpsec *bool `json:"supportIpsec,omitempty"`

	SupportLBS *bool `json:"supportLBS,omitempty"`

	SupportResetCablemodem *bool `json:"supportResetCablemodem,omitempty"`
}

func NewCommonAttribute() *CommonAttribute {
	commonAttributeType := new(CommonAttribute)
	return commonAttributeType
}

func NewDefaultCommonAttribute() *CommonAttribute {
	commonAttributeType := new(CommonAttribute)
	return commonAttributeType
}

type ExternalAntenna struct {
	// ChainMask
	// Constraints:
	//    - nullable
	//    - oneof:[Two,Three]
	ChainMask *string `json:"chainMask,omitempty" validate:"omitempty,oneof=Two Three"`

	// Dbi
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:90
	Dbi *int `json:"dbi,omitempty" validate:"omitempty,gte=0,lte=90"`

	// Enabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled" validate:"required"`
}

func NewExternalAntenna() *ExternalAntenna {
	externalAntennaType := new(ExternalAntenna)
	return externalAntennaType
}

func NewDefaultExternalAntenna() *ExternalAntenna {
	externalAntennaType := new(ExternalAntenna)
	return externalAntennaType
}

type LacpSetting struct {
	Enabled *bool `json:"enabled,omitempty"`
}

func NewLacpSetting() *LacpSetting {
	lacpSettingType := new(LacpSetting)
	return lacpSettingType
}

func NewDefaultLacpSetting() *LacpSetting {
	lacpSettingType := new(LacpSetting)
	enabledField := false
	lacpSettingType.Enabled = &enabledField
	return lacpSettingType
}

type LanPort8021X struct {
	Authenticator *LanPortAuthenticator `json:"authenticator,omitempty"`

	Supplicant *LanPortSupplicant `json:"supplicant,omitempty"`

	// Type
	// Constraints:
	//    - required
	//    - oneof:[Disable,Supplicant,PortBasedAuthenticator,MACBasedAuthenticator]
	Type *string `json:"type" validate:"required,oneof=Disable Supplicant PortBasedAuthenticator MACBasedAuthenticator"`
}

func NewLanPort8021X() *LanPort8021X {
	lanPort8021XType := new(LanPort8021X)
	return lanPort8021XType
}

func NewDefaultLanPort8021X() *LanPort8021X {
	lanPort8021XType := new(LanPort8021X)
	return lanPort8021XType
}

type LanPortAuthenticator struct {
	Accounting *AuthenticatorAAAServer `json:"accounting,omitempty"`

	Authentication *AuthenticatorAAAServer `json:"authentication,omitempty"`

	DisabledAccounting *bool `json:"disabledAccounting,omitempty"`

	// MacAuthByPassEnabled
	// Constraints:
	//    - required
	MacAuthByPassEnabled *bool `json:"macAuthByPassEnabled" validate:"required"`
}

func NewLanPortAuthenticator() *LanPortAuthenticator {
	lanPortAuthenticatorType := new(LanPortAuthenticator)
	return lanPortAuthenticatorType
}

func NewDefaultLanPortAuthenticator() *LanPortAuthenticator {
	lanPortAuthenticatorType := new(LanPortAuthenticator)
	return lanPortAuthenticatorType
}

type LanPortSetting struct {
	// Enabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled" validate:"required"`

	EthPortProfile *common.GenericRef `json:"ethPortProfile,omitempty"`

	Members *string `json:"members,omitempty"`

	OverwriteVlanEnabled *bool `json:"overwriteVlanEnabled,omitempty"`

	// PortName
	// Constraints:
	//    - required
	//    - oneof:[LAN1,LAN2,LAN3,LAN4,LAN5]
	PortName *string `json:"portName" validate:"required,oneof=LAN1 LAN2 LAN3 LAN4 LAN5"`

	// VlanUntagId
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:4094
	VlanUntagId *int `json:"vlanUntagId,omitempty" validate:"omitempty,gte=0,lte=4094"`
}

func NewLanPortSetting() *LanPortSetting {
	lanPortSettingType := new(LanPortSetting)
	return lanPortSettingType
}

func NewDefaultLanPortSetting() *LanPortSetting {
	lanPortSettingType := new(LanPortSetting)
	return lanPortSettingType
}

type LanPortSupplicant struct {
	// Password
	// Constraints:
	//    - nullable
	//    - max:64
	Password *string `json:"password,omitempty" validate:"omitempty,max=64"`

	// Type
	// Constraints:
	//    - required
	//    - oneof:[MACAddress,Custom]
	Type *string `json:"type" validate:"required,oneof=MACAddress Custom"`

	// UserName
	// Constraints:
	//    - nullable
	//    - max:64
	UserName *string `json:"userName,omitempty" validate:"omitempty,max=64"`
}

func NewLanPortSupplicant() *LanPortSupplicant {
	lanPortSupplicantType := new(LanPortSupplicant)
	return lanPortSupplicantType
}

func NewDefaultLanPortSupplicant() *LanPortSupplicant {
	lanPortSupplicantType := new(LanPortSupplicant)
	return lanPortSupplicantType
}

type LldpSetting struct {
	// AdvertiseIntervalInSec
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:300
	AdvertiseIntervalInSec *int `json:"advertiseIntervalInSec,omitempty" validate:"omitempty,gte=1,lte=300"`

	// Enabled
	// Constraints:
	//    - required
	//    - default:true
	Enabled *bool `json:"enabled" validate:"required"`

	// HoldTimeInSec
	// Constraints:
	//    - nullable
	//    - min:60
	//    - max:1200
	HoldTimeInSec *int `json:"holdTimeInSec,omitempty" validate:"omitempty,gte=60,lte=1200"`

	ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty"`
}

func NewLldpSetting() *LldpSetting {
	lldpSettingType := new(LldpSetting)
	return lldpSettingType
}

func NewDefaultLldpSetting() *LldpSetting {
	lldpSettingType := new(LldpSetting)
	enabledField := false
	lldpSettingType.Enabled = &enabledField
	return lldpSettingType
}
