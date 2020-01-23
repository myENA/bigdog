package vsz

// API Version: v8_1

type WSGAPModel struct {
	CellularSettings *WSGAPModelCellularSettings `json:"cellularSettings,omitempty"`

	ExternalAntenna24 *WSGAPModelExternalAntenna `json:"externalAntenna24,omitempty"`

	ExternalAntenna50 *WSGAPModelExternalAntenna `json:"externalAntenna50,omitempty"`

	// InternalHeaterEnabled
	// Constraints:
	//    - nullable
	InternalHeaterEnabled *bool `json:"internalHeaterEnabled,omitempty" validate:"omitempty"`

	Lacp *WSGAPModelLacpSetting `json:"lacp,omitempty"`

	LanPorts []*WSGAPModelLanPortSetting `json:"lanPorts,omitempty"`

	// LedMode
	// Constraints:
	//    - nullable
	//    - oneof:[CableModem,AccessPoint,CableModem_AccessPoint,CableModem_AccessPoint_DEFAULT,ActiveSurgeProtector,ActiveSurgeProtector_ModemOnline_DEFAULT,Off]
	LedMode *string `json:"ledMode,omitempty" validate:"omitempty,oneof=CableModem AccessPoint CableModem_AccessPoint CableModem_AccessPoint_DEFAULT ActiveSurgeProtector ActiveSurgeProtector_ModemOnline_DEFAULT Off"`

	// LedStatusEnabled
	// Constraints:
	//    - nullable
	LedStatusEnabled *bool `json:"ledStatusEnabled,omitempty" validate:"omitempty"`

	Lldp *WSGAPModelLldpSetting `json:"lldp,omitempty"`

	// PoeModeSetting
	// Constraints:
	//    - nullable
	//    - oneof:[Auto,_802_3af,_802_3at,_802_3atPlus]
	PoeModeSetting *string `json:"poeModeSetting,omitempty" validate:"omitempty,oneof=Auto _802_3af _802_3at _802_3atPlus"`

	// PoeOutPortEnabled
	// Constraints:
	//    - nullable
	PoeOutPortEnabled *bool `json:"poeOutPortEnabled,omitempty" validate:"omitempty"`

	// PoeTxChain
	// Option to use 1, 2 or 4 Tx chains while AP power source is 802.3af PoE
	// Constraints:
	//    - nullable
	//    - default:2
	PoeTxChain *int `json:"poeTxChain,omitempty" validate:"omitempty"`

	// RadioBand
	// Band switch between 2.4GHz and 5GHz is provided in single radio AP ZF-7321, ZF-7321-U, and ZF-7441.
	// Constraints:
	//    - nullable
	//    - oneof:[2.4GHz,5GHz]
	RadioBand *string `json:"radioBand,omitempty" validate:"omitempty,oneof=2.4GHz 5GHz"`

	// UsbPowerEnable
	// Constraints:
	//    - nullable
	UsbPowerEnable *bool `json:"usbPowerEnable,omitempty" validate:"omitempty"`
}

func NewWSGAPModel() *WSGAPModel {
	m := new(WSGAPModel)
	return m
}

type WSGAPModelAuthenticatorAAAServer struct {
	// EnableUseSCGasProxy
	// Constraints:
	//    - required
	EnableUseSCGasProxy *bool `json:"enableUseSCGasProxy" validate:"required"`

	Server *WSGCommonGenericRef `json:"server,omitempty"`
}

func NewWSGAPModelAuthenticatorAAAServer() *WSGAPModelAuthenticatorAAAServer {
	m := new(WSGAPModelAuthenticatorAAAServer)
	return m
}

type WSGAPModelCellularSettings struct {
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

func NewWSGAPModelCellularSettings() *WSGAPModelCellularSettings {
	m := new(WSGAPModelCellularSettings)
	return m
}

type WSGAPModelCommonAttribute struct {
	AllowDfsCountry *string `json:"allowDfsCountry,omitempty"`

	CapabilityScore *float64 `json:"capabilityScore,omitempty"`

	CpuFrequency *int `json:"cpuFrequency,omitempty"`

	// HasCablemodem
	// Constraints:
	//    - nullable
	HasCablemodem *bool `json:"hasCablemodem,omitempty" validate:"omitempty"`

	// HasGps
	// Constraints:
	//    - nullable
	HasGps *bool `json:"hasGps,omitempty" validate:"omitempty"`

	// HasScanRadio
	// Constraints:
	//    - nullable
	HasScanRadio *bool `json:"hasScanRadio,omitempty" validate:"omitempty"`

	// IsAllowDisableExtAnt
	// Constraints:
	//    - nullable
	IsAllowDisableExtAnt *bool `json:"isAllowDisableExtAnt,omitempty" validate:"omitempty"`

	// IsDualRadio
	// Constraints:
	//    - nullable
	IsDualRadio *bool `json:"isDualRadio,omitempty" validate:"omitempty"`

	// IsOutdoor
	// Constraints:
	//    - nullable
	IsOutdoor *bool `json:"isOutdoor,omitempty" validate:"omitempty"`

	MaxChannelization5G *int `json:"maxChannelization5G,omitempty"`

	MaxChannelization24G *int `json:"maxChannelization24G,omitempty"`

	MaxClientsUpper *int `json:"maxClientsUpper,omitempty"`

	MaxWlanNum5G *int `json:"maxWlanNum5G,omitempty"`

	MaxWlanNum24G *int `json:"maxWlanNum24G,omitempty"`

	MeshRadioCaps *string `json:"meshRadioCaps,omitempty"`

	// NoAvc
	// Constraints:
	//    - nullable
	NoAvc *bool `json:"noAvc,omitempty" validate:"omitempty"`

	// NoMesh
	// Constraints:
	//    - nullable
	NoMesh *bool `json:"noMesh,omitempty" validate:"omitempty"`

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
	Support11AC *bool `json:"support11AC,omitempty" validate:"omitempty"`

	// SupportAPUsbSoftwarePackage
	// Constraints:
	//    - nullable
	SupportAPUsbSoftwarePackage *bool `json:"supportAPUsbSoftwarePackage,omitempty" validate:"omitempty"`

	// SupportBandSwitch
	// Constraints:
	//    - nullable
	SupportBandSwitch *bool `json:"supportBandSwitch,omitempty" validate:"omitempty"`

	// SupportBonjour
	// Constraints:
	//    - nullable
	SupportBonjour *bool `json:"supportBonjour,omitempty" validate:"omitempty"`

	// SupportChannelization160
	// Constraints:
	//    - nullable
	SupportChannelization160 *bool `json:"supportChannelization160,omitempty" validate:"omitempty"`

	// SupportIpsec
	// Constraints:
	//    - nullable
	SupportIpsec *bool `json:"supportIpsec,omitempty" validate:"omitempty"`

	// SupportLBS
	// Constraints:
	//    - nullable
	SupportLBS *bool `json:"supportLBS,omitempty" validate:"omitempty"`

	// SupportResetCablemodem
	// Constraints:
	//    - nullable
	SupportResetCablemodem *bool `json:"supportResetCablemodem,omitempty" validate:"omitempty"`
}

func NewWSGAPModelCommonAttribute() *WSGAPModelCommonAttribute {
	m := new(WSGAPModelCommonAttribute)
	return m
}

type WSGAPModelExternalAntenna struct {
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

func NewWSGAPModelExternalAntenna() *WSGAPModelExternalAntenna {
	m := new(WSGAPModelExternalAntenna)
	return m
}

type WSGAPModelLacpSetting struct {
	Enabled *bool `json:"enabled,omitempty"`
}

func NewWSGAPModelLacpSetting() *WSGAPModelLacpSetting {
	m := new(WSGAPModelLacpSetting)
	return m
}

type WSGAPModelLanPort8021X struct {
	Authenticator *WSGAPModelLanPortAuthenticator `json:"authenticator,omitempty"`

	Supplicant *WSGAPModelLanPortSupplicant `json:"supplicant,omitempty"`

	// Type
	// Constraints:
	//    - required
	//    - oneof:[Disable,Supplicant,PortBasedAuthenticator,MACBasedAuthenticator]
	Type *string `json:"type" validate:"required,oneof=Disable Supplicant PortBasedAuthenticator MACBasedAuthenticator"`
}

func NewWSGAPModelLanPort8021X() *WSGAPModelLanPort8021X {
	m := new(WSGAPModelLanPort8021X)
	return m
}

type WSGAPModelLanPortAuthenticator struct {
	Accounting *WSGAPModelAuthenticatorAAAServer `json:"accounting,omitempty"`

	Authentication *WSGAPModelAuthenticatorAAAServer `json:"authentication,omitempty"`

	DisabledAccounting *bool `json:"disabledAccounting,omitempty"`

	// MacAuthByPassEnabled
	// Constraints:
	//    - required
	MacAuthByPassEnabled *bool `json:"macAuthByPassEnabled" validate:"required"`
}

func NewWSGAPModelLanPortAuthenticator() *WSGAPModelLanPortAuthenticator {
	m := new(WSGAPModelLanPortAuthenticator)
	return m
}

type WSGAPModelLanPortSetting struct {
	// Enabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled" validate:"required"`

	EthPortProfile *WSGCommonGenericRef `json:"ethPortProfile,omitempty"`

	// Members
	// Constraints:
	//    - nullable
	Members *string `json:"members,omitempty" validate:"omitempty"`

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

func NewWSGAPModelLanPortSetting() *WSGAPModelLanPortSetting {
	m := new(WSGAPModelLanPortSetting)
	return m
}

type WSGAPModelLanPortSupplicant struct {
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

func NewWSGAPModelLanPortSupplicant() *WSGAPModelLanPortSupplicant {
	m := new(WSGAPModelLanPortSupplicant)
	return m
}

type WSGAPModelLldpSetting struct {
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

	// ManagementIPTLVEnabled
	// Constraints:
	//    - nullable
	ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty" validate:"omitempty"`
}

func NewWSGAPModelLldpSetting() *WSGAPModelLldpSetting {
	m := new(WSGAPModelLldpSetting)
	return m
}
