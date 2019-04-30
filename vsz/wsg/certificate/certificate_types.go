package certificate

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/common"
)

type Certificate struct {
	CertificasSigningRequest *common.GenericRef `json:"certificasSigningRequest,omitempty"`

	// Data
	// Data of the certificate
	Data *string `json:"data,omitempty"`

	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the certificate
	ID *string `json:"id,omitempty"`

	// Information
	// Information of the certificate
	Information *string `json:"information,omitempty"`

	// IntermediateData
	// Intermediate data of the certificate
	IntermediateData []string `json:"intermediateData,omitempty"`

	Name *string `json:"name,omitempty"`

	// Passphrase
	// Key passphrase
	Passphrase *string `json:"passphrase,omitempty"`

	// PrivateKeyData
	// Private key data of the certificate
	PrivateKeyData *string `json:"privateKeyData,omitempty"`

	// PublicKey
	// Public key data of the certificate
	PublicKey *string `json:"publicKey,omitempty"`

	// RootData
	// Root data of the certificate
	RootData *string `json:"rootData,omitempty"`
}

type CertificateList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CertificateListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CertificateListType struct {
	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the certificate
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type CertificatesSigningRequest struct {
	// City
	// City of the certificates signing request
	City *string `json:"city,omitempty"`

	// CommonName
	// Common name of the certificates signing request
	CommonName *string `json:"commonName,omitempty"`

	// CountryCode
	// Country code of the certificates signing request
	CountryCode *string `json:"countryCode,omitempty"`

	Description *string `json:"description,omitempty"`

	// Email
	// Email of the certificates signing request
	Email *string `json:"email,omitempty"`

	// ID
	// Identifier of the certificates signing request
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	// Organization
	// Organization of the certificates signing request
	Organization *string `json:"organization,omitempty"`

	// OrganizationUnit
	// Organization unit of the certificates signing request
	OrganizationUnit *string `json:"organizationUnit,omitempty"`

	// State
	// State of the certificates signing request
	State *string `json:"state,omitempty"`
}

type CertSetting struct {
	// ServiceCertificates
	// Certificate Setting of the service
	ServiceCertificates []*ServiceCertificate `json:"serviceCertificates,omitempty"`
}

type CreateCert struct {
	CertificasSigningRequest *common.GenericRef `json:"certificasSigningRequest,omitempty"`

	// Data
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	Data *string `json:"data,omitempty"`

	Description *string `json:"description,omitempty"`

	// IntermediateData
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	IntermediateData []string `json:"intermediateData,omitempty"`

	Name *string `json:"name,omitempty"`

	// Passphrase
	// Key passphrase
	Passphrase *string `json:"passphrase,omitempty"`

	// PrivateKeyData
	// Incompatible with certificasSigningRequest. Must choose one to enter. The value must be in PEM format
	// which is a Base64 encoded DER certificate.
	PrivateKeyData *string `json:"privateKeyData,omitempty"`

	// RootData
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	RootData *string `json:"rootData,omitempty"`
}

type CreateCSR struct {
	// City
	// City of the certificates signing request
	City *string `json:"city,omitempty"`

	CommonName *string `json:"commonName,omitempty"`

	// CountryCode
	// Country code of the certificates signing request
	CountryCode *string `json:"countryCode,omitempty"`

	Description *string `json:"description,omitempty"`

	Email *string `json:"email,omitempty"`

	Name *string `json:"name,omitempty"`

	// Organization
	// Organization of the certificates signing request
	Organization *string `json:"organization,omitempty"`

	// OrganizationUnit
	// Organization unit of the certificates signing request
	OrganizationUnit *string `json:"organizationUnit,omitempty"`

	// State
	// State of the certificates signing request
	State *string `json:"state,omitempty"`
}

type CreateTrustedCAChain struct {
	Description *string `json:"description,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	Name *string `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

type CsrList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CsrListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CsrListType struct {
	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the certificates signing request
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type DeleteBulk struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type ModifyTrustedCAChain struct {
	Description *string `json:"description,omitempty"`

	// Information
	// Information of the certificates
	Information *string `json:"information,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	Name *string `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

type ServiceCertificate struct {
	Certificate *common.GenericRef `json:"certificate,omitempty"`

	// Service
	// the service
	Service *string `json:"service,omitempty"`
}

type ServiceCertificates []*ServiceCertificate

type TrustedCAChain struct {
	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the trusted CA chain certificates
	ID *string `json:"id,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	Name *string `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

type TrustedCAChainCertList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*TrustedCAChainCertListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type TrustedCAChainCertListType struct {
	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the trusted CA chain certificate
	ID *string `json:"id,omitempty"`

	// Information
	// Information of the certificates
	Information *string `json:"information,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	ModifiedDateTime *string `json:"modifiedDateTime,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}
