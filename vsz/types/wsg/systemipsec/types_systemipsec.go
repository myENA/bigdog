package systemipsec

// API Version: v8_1

type GetResult struct {
	// AuthType
	// Authentication type
	AuthType *string `json:"authType,omitempty"`

	CertSubject *string `json:"certSubject,omitempty"`

	// EspProposals
	// Proposal of Encapsulating Security Payload
	EspProposals []*Proposal `json:"espProposals,omitempty"`

	// EspRekeyDisabled
	// Disable rekey mechanisam of Encapsulating Security Payload
	EspRekeyDisabled *bool `json:"espRekeyDisabled,omitempty"`

	// EspRekeyTime
	// Rekey time of Encapsulating Security Payload
	EspRekeyTime *int `json:"espRekeyTime,omitempty"`

	// EspRekeyTimeUnit
	// Rekey time unit of Encapsulating Security Payload
	EspRekeyTimeUnit *string `json:"espRekeyTimeUnit,omitempty"`

	// IkeProposals
	// Proposal of Internet Key Exchange
	IkeProposals []*Proposal `json:"ikeProposals,omitempty"`

	// IkeRekeyDisabled
	// Disable rekey mechanisam of Internet Key Exchange
	IkeRekeyDisabled *bool `json:"ikeRekeyDisabled,omitempty"`

	// IkeRekeyTime
	// Rekey time of Internet Key Exchange
	IkeRekeyTime *int `json:"ikeRekeyTime,omitempty"`

	// IkeRekeyTimeUnit
	// Rekey time unit of Internet Key Exchange
	IkeRekeyTimeUnit *string `json:"ikeRekeyTimeUnit,omitempty"`

	// IPSecEnabled
	// Enable System IPSec
	IPSecEnabled *bool `json:"ipSecEnabled,omitempty"`

	OcspEnabled *bool `json:"ocspEnabled,omitempty"`

	OcspServerURI *string `json:"ocspServerUri,omitempty"`

	// PreSharedKey
	// Pre-shared key
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// RemoteID
	// The IP of IPSec receiver
	RemoteID *string `json:"remoteId,omitempty"`

	// SCGCertID
	// SCG client certification id.
	SCGCertID *string `json:"scgCertId,omitempty"`

	// SecurityGateway
	// Security gateway IP
	SecurityGateway *string `json:"securityGateway,omitempty"`

	// SubnetMask
	// Subnet Mask of security gateway
	SubnetMask *string `json:"subnetMask,omitempty"`

	TrustChainProfileID *string `json:"trustChainProfileId,omitempty"`
}

type Proposal struct {
	// AuthAlg
	// Authentication algorithm
	AuthAlg *string `json:"authAlg,omitempty" validate:"required"`

	// EncAlg
	// Encrytion algorithm
	EncAlg *string `json:"encAlg,omitempty" validate:"required"`
}

type Update struct {
	// AuthType
	// Authentication type
	AuthType *string `json:"authType,omitempty"`

	CertSubject *string `json:"certSubject,omitempty"`

	// EspProposals
	// Proposal of Encapsulating Security Payload
	EspProposals []*Proposal `json:"espProposals,omitempty"`

	// EspRekeyDisabled
	// Disable rekey mechanisam of Encapsulating Security Payload
	EspRekeyDisabled *bool `json:"espRekeyDisabled,omitempty"`

	// EspRekeyTime
	// Rekey time of Encapsulating Security Payload
	EspRekeyTime *int `json:"espRekeyTime,omitempty"`

	// EspRekeyTimeUnit
	// Rekey time unit of Encapsulating Security Payload
	EspRekeyTimeUnit *string `json:"espRekeyTimeUnit,omitempty"`

	// IkeProposals
	// Proposal of Internet Key Exchange
	IkeProposals []*Proposal `json:"ikeProposals,omitempty"`

	// IkeRekeyDisabled
	// Disable rekey mechanisam of Internet Key Exchange
	IkeRekeyDisabled *bool `json:"ikeRekeyDisabled,omitempty"`

	// IkeRekeyTime
	// Rekey time of Internet Key Exchange
	IkeRekeyTime *int `json:"ikeRekeyTime,omitempty"`

	// IkeRekeyTimeUnit
	// Rekey time unit of Internet Key Exchange
	IkeRekeyTimeUnit *string `json:"ikeRekeyTimeUnit,omitempty"`

	// IPSecEnabled
	// Enable System IPSec
	IPSecEnabled *bool `json:"ipSecEnabled,omitempty" validate:"required"`

	OcspEnabled *bool `json:"ocspEnabled,omitempty"`

	OcspServerURI *string `json:"ocspServerUri,omitempty"`

	// PreSharedKey
	// Pre-shared key
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// RemoteID
	// The IP of IPSec receiver
	RemoteID *string `json:"remoteId,omitempty"`

	// SCGCertID
	// SCG client certification id.
	SCGCertID *string `json:"scgCertId,omitempty"`

	// SecurityGateway
	// Security gateway IP
	SecurityGateway *string `json:"securityGateway,omitempty"`

	// SubnetMask
	// Subnet Mask of security gateway
	SubnetMask *string `json:"subnetMask,omitempty"`

	TrustChainProfileID *string `json:"trustChainProfileId,omitempty"`
}
