package apmodel

// API Version: v8_0

type ApModel struct {
	CellularSettings      *CellularSettings  `json:"cellularSettings,omitempty"`
	ExternalAntenna24     *ExternalAntenna24 `json:"externalAntenna24,omitempty"`
	ExternalAntenna50     *ExternalAntenna50 `json:"externalAntenna50,omitempty"`
	InternalHeaterEnabled *bool              `json:"internalHeaterEnabled,omitempty"`
	LanPorts              []*LanPorts        `json:"lanPorts,omitempty"`
	LedMode               *string            `json:"ledMode,omitempty"`
	LedStatusEnabled      *bool              `json:"ledStatusEnabled,omitempty"`
	Lldp                  *Lldp              `json:"lldp,omitempty"`
	PoeModeSetting        *string            `json:"poeModeSetting,omitempty"`
	PoeOutPortEnabled     *bool              `json:"poeOutPortEnabled,omitempty"`
	PoeTxChain            *int               `json:"poeTxChain,omitempty"`
	RadioBand             *string            `json:"radioBand,omitempty"`
	UsbPowerEnable        *bool              `json:"usbPowerEnable,omitempty"`
}

type AuthenticatorAAAServer struct {
	EnableUseSCGasProxy *bool              `json:"enableUseSCGasProxy,omitempty"`
	Server              *common.GenericRef `json:"server,omitempty"`
}

type CellularSettings struct {
	DataRoaming      *int    `json:"dataRoaming,omitempty"`
	DataRoaming2     *int    `json:"dataRoaming2,omitempty"`
	MobileAPName     *string `json:"mobileAPName,omitempty"`
	MobileAPName2    *string `json:"mobileAPName2,omitempty"`
	SelectGG         *int    `json:"select3g4g,omitempty"`
	SelectGG2        *int    `json:"select3g4g2,omitempty"`
	SimCardUsage     *int    `json:"simCardUsage,omitempty"`
	WanConnection    *int    `json:"wanConnection,omitempty"`
	WanRecoveryTimer *int    `json:"wanRecoveryTimer,omitempty"`
}

type CommonAttribute struct {
	AllowDfsCountry             *string  `json:"allowDfsCountry,omitempty"`
	CapabilityScore             *float64 `json:"capabilityScore,omitempty"`
	CPUFrequency                *int     `json:"cpuFrequency,omitempty"`
	HasCablemodem               *bool    `json:"hasCablemodem,omitempty"`
	HasGps                      *bool    `json:"hasGps,omitempty"`
	HasScanRadio                *bool    `json:"hasScanRadio,omitempty"`
	IsAllowDisableExtAnt        *bool    `json:"isAllowDisableExtAnt,omitempty"`
	IsDualRadio                 *bool    `json:"isDualRadio,omitempty"`
	IsOutdoor                   *bool    `json:"isOutdoor,omitempty"`
	MaxChannelization5G         *int     `json:"maxChannelization5G,omitempty"`
	MaxChannelization24G        *int     `json:"maxChannelization24G,omitempty"`
	MaxClientsUpper             *int     `json:"maxClientsUpper,omitempty"`
	MaxWLANNum5G                *int     `json:"maxWlanNum5G,omitempty"`
	MaxWLANNum24G               *int     `json:"maxWlanNum24G,omitempty"`
	NoAvc                       *bool    `json:"noAvc,omitempty"`
	NoMesh                      *bool    `json:"noMesh,omitempty"`
	NonEditablePorts            []int    `json:"nonEditablePorts,omitempty"`
	NonVisiblePorts             []int    `json:"nonVisiblePorts,omitempty"`
	NumOfCores                  *int     `json:"numOfCores,omitempty"`
	RAM                         *int     `json:"ram,omitempty"`
	Reserved5GWLANForMesh       *int     `json:"reserved5GWlanForMesh,omitempty"`
	ScalingFactor               *int     `json:"scalingFactor,omitempty"`
	Support11AC                 *bool    `json:"support11AC,omitempty"`
	SupportAPUsbSoftwarePackage *bool    `json:"supportAPUsbSoftwarePackage,omitempty"`
	SupportBandSwitch           *bool    `json:"supportBandSwitch,omitempty"`
	SupportBonjour              *bool    `json:"supportBonjour,omitempty"`
	SupportChannelization160    *bool    `json:"supportChannelization160,omitempty"`
	SupportIpsec                *bool    `json:"supportIpsec,omitempty"`
	SupportLBS                  *bool    `json:"supportLBS,omitempty"`
	SupportResetCablemodem      *bool    `json:"supportResetCablemodem,omitempty"`
}

type ExternalAntenna struct {
	ChainMask *string `json:"chainMask,omitempty"`
	Dbi       *int    `json:"dbi,omitempty"`
	Enabled   *bool   `json:"enabled,omitempty"`
}

type LanPort8021X struct {
	Authenticator *Authenticator `json:"authenticator,omitempty"`
	Supplicant    *Supplicant    `json:"supplicant,omitempty"`
	Type          *string        `json:"type,omitempty"`
}

type LanPortAuthenticator struct {
	Accounting           *Accounting     `json:"accounting,omitempty"`
	Authentication       *Authentication `json:"authentication,omitempty"`
	DisabledAccounting   *bool           `json:"disabledAccounting,omitempty"`
	MacAuthByPassEnabled *bool           `json:"macAuthByPassEnabled,omitempty"`
}

type LanPortSetting struct {
	Enabled              *bool              `json:"enabled,omitempty"`
	EthPortProfile       *common.GenericRef `json:"ethPortProfile,omitempty"`
	Members              *string            `json:"members,omitempty"`
	OverwriteVlanEnabled *bool              `json:"overwriteVlanEnabled,omitempty"`
	PortName             *string            `json:"portName,omitempty"`
	VlanUntagID          *int               `json:"vlanUntagId,omitempty"`
}

type LanPortSupplicant struct {
	Password *string `json:"password,omitempty"`
	Type     *string `json:"type,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

type LldpSetting struct {
	AdvertiseIntervalInSec *int  `json:"advertiseIntervalInSec,omitempty"`
	Enabled                *bool `json:"enabled,omitempty"`
	HoldTimeInSec          *int  `json:"holdTimeInSec,omitempty"`
	ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty"`
}
