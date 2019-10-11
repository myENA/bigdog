package certificate

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type Certificate struct {
	CertificasSigningRequest *common.GenericRef `json:"certificasSigningRequest,omitempty"`

	// Data
	// Data of the certificate
	Data *string `json:"data,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the certificate
	Id *string `json:"id,omitempty"`

	// Information
	// Information of the certificate
	Information *string `json:"information,omitempty"`

	// IntermediateData
	// Intermediate data of the certificate
	IntermediateData []string `json:"intermediateData,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

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

func NewCertificate() *Certificate {
	certificateType := new(Certificate)
	return certificateType
}

func NewCertificateWithDefaults() *Certificate {
	certificateType := new(Certificate)
	return certificateType
}

type CertificateList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CertificateListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewCertificateList() *CertificateList {
	certificateListType := new(CertificateList)
	return certificateListType
}

func NewCertificateListWithDefaults() *CertificateList {
	certificateListType := new(CertificateList)
	return certificateListType
}

type CertificateListType struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the certificate
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewCertificateListType() *CertificateListType {
	certificateListTypeType := new(CertificateListType)
	return certificateListTypeType
}

func NewCertificateListTypeWithDefaults() *CertificateListType {
	certificateListTypeType := new(CertificateListType)
	return certificateListTypeType
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

	Description *common.Description `json:"description,omitempty"`

	// Email
	// Email of the certificates signing request
	Email *string `json:"email,omitempty"`

	// Id
	// Identifier of the certificates signing request
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

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

func NewCertificatesSigningRequest() *CertificatesSigningRequest {
	certificatesSigningRequestType := new(CertificatesSigningRequest)
	return certificatesSigningRequestType
}

func NewCertificatesSigningRequestWithDefaults() *CertificatesSigningRequest {
	certificatesSigningRequestType := new(CertificatesSigningRequest)
	return certificatesSigningRequestType
}

type CertSetting struct {
	// ServiceCertificates
	// Certificate Setting of the service
	ServiceCertificates []*ServiceCertificate `json:"serviceCertificates,omitempty"`
}

func NewCertSetting() *CertSetting {
	certSettingType := new(CertSetting)
	return certSettingType
}

func NewCertSettingWithDefaults() *CertSetting {
	certSettingType := new(CertSetting)
	return certSettingType
}

type ClientCert struct {
	// Data
	// Data of the client certificate
	Data *string `json:"data,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the client certificate
	Id *string `json:"id,omitempty"`

	// Information
	// Information of the client certificate
	Information *string `json:"information,omitempty"`

	// IntermediateData
	// Intermediate data of the client certificate
	IntermediateData []string `json:"intermediateData,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// PrivateKeyData
	// Private key data of the client certificate
	PrivateKeyData *string `json:"privateKeyData,omitempty"`

	// PublicKey
	// Public key data of the client certificate
	PublicKey *string `json:"publicKey,omitempty"`

	// RootData
	// Root data of the client certificate
	RootData *string `json:"rootData,omitempty"`
}

func NewClientCert() *ClientCert {
	clientCertType := new(ClientCert)
	return clientCertType
}

func NewClientCertWithDefaults() *ClientCert {
	clientCertType := new(ClientCert)
	return clientCertType
}

type ClientCertList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ClientCertListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewClientCertList() *ClientCertList {
	clientCertListType := new(ClientCertList)
	return clientCertListType
}

func NewClientCertListWithDefaults() *ClientCertList {
	clientCertListType := new(ClientCertList)
	return clientCertListType
}

type ClientCertListType struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the client certificate
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewClientCertListType() *ClientCertListType {
	clientCertListTypeType := new(ClientCertListType)
	return clientCertListTypeType
}

func NewClientCertListTypeWithDefaults() *ClientCertListType {
	clientCertListTypeType := new(ClientCertListType)
	return clientCertListTypeType
}

type CreateCert struct {
	CertificasSigningRequest *common.GenericRef `json:"certificasSigningRequest,omitempty"`

	// Data
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	// Constraints:
	//    - required
	Data *string `json:"data" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// IntermediateData
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	IntermediateData []string `json:"intermediateData,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Passphrase
	// Key passphrase
	Passphrase *string `json:"passphrase,omitempty"`

	// PrivateKeyData
	// Incompatible with certificasSigningRequest. Must choose one to enter. The value must be in PEM format which is a Base64 encoded DER certificate.
	PrivateKeyData *string `json:"privateKeyData,omitempty"`

	// RootData
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	RootData *string `json:"rootData,omitempty"`
}

func NewCreateCert() *CreateCert {
	createCertType := new(CreateCert)
	return createCertType
}

func NewCreateCertWithDefaults() *CreateCert {
	createCertType := new(CreateCert)
	return createCertType
}

type CreateClientCert struct {
	// Data
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	// Constraints:
	//    - required
	Data *string `json:"data" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// IntermediateData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	IntermediateData []string `json:"intermediateData,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// PrivateKeyData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	// Constraints:
	//    - required
	PrivateKeyData *string `json:"privateKeyData" validate:"required"`

	// RootData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	RootData *string `json:"rootData,omitempty"`
}

func NewCreateClientCert() *CreateClientCert {
	createClientCertType := new(CreateClientCert)
	return createClientCertType
}

func NewCreateClientCertWithDefaults() *CreateClientCert {
	createClientCertType := new(CreateClientCert)
	return createClientCertType
}

type CreateCSR struct {
	// City
	// City of the certificates signing request
	// Constraints:
	//    - required
	//    - max:128
	City *string `json:"city" validate:"required,max=128"`

	// CommonName
	// Constraints:
	//    - required
	CommonName *common.FQDN `json:"commonName" validate:"required"`

	// CountryCode
	// Country code of the certificates signing request
	// Constraints:
	//    - required
	CountryCode *string `json:"countryCode" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// Email
	// Constraints:
	//    - required
	Email *common.Email `json:"email" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Organization
	// Organization of the certificates signing request
	// Constraints:
	//    - required
	//    - max:64
	Organization *string `json:"organization" validate:"required,max=64"`

	// OrganizationUnit
	// Organization unit of the certificates signing request
	// Constraints:
	//    - nullable
	//    - max:64
	OrganizationUnit *string `json:"organizationUnit,omitempty" validate:"omitempty,max=64"`

	// State
	// State of the certificates signing request
	// Constraints:
	//    - required
	//    - max:128
	State *string `json:"state" validate:"required,max=128"`
}

func NewCreateCSR() *CreateCSR {
	createCSRType := new(CreateCSR)
	return createCSRType
}

func NewCreateCSRWithDefaults() *CreateCSR {
	createCSRType := new(CreateCSR)
	return createCSRType
}

type CreateTrustedCAChain struct {
	Description *common.Description `json:"description,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	// Constraints:
	//    - required
	RootCertData *string `json:"rootCertData" validate:"required"`
}

func NewCreateTrustedCAChain() *CreateTrustedCAChain {
	createTrustedCAChainType := new(CreateTrustedCAChain)
	return createTrustedCAChainType
}

func NewCreateTrustedCAChainWithDefaults() *CreateTrustedCAChain {
	createTrustedCAChainType := new(CreateTrustedCAChain)
	return createTrustedCAChainType
}

type CsrList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CsrListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewCsrList() *CsrList {
	csrListType := new(CsrList)
	return csrListType
}

func NewCsrListWithDefaults() *CsrList {
	csrListType := new(CsrList)
	return csrListType
}

type CsrListType struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the certificates signing request
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewCsrListType() *CsrListType {
	csrListTypeType := new(CsrListType)
	return csrListTypeType
}

func NewCsrListTypeWithDefaults() *CsrListType {
	csrListTypeType := new(CsrListType)
	return csrListTypeType
}

type DeleteBulk struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulk() *DeleteBulk {
	deleteBulkType := new(DeleteBulk)
	return deleteBulkType
}

func NewDeleteBulkWithDefaults() *DeleteBulk {
	deleteBulkType := new(DeleteBulk)
	return deleteBulkType
}

type ModifyTrustedCAChain struct {
	Description *common.Description `json:"description,omitempty"`

	// Information
	// Information of the certificates
	Information *string `json:"information,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

func NewModifyTrustedCAChain() *ModifyTrustedCAChain {
	modifyTrustedCAChainType := new(ModifyTrustedCAChain)
	return modifyTrustedCAChainType
}

func NewModifyTrustedCAChainWithDefaults() *ModifyTrustedCAChain {
	modifyTrustedCAChainType := new(ModifyTrustedCAChain)
	return modifyTrustedCAChainType
}

type ServiceCertificate struct {
	Certificate *common.GenericRef `json:"certificate,omitempty"`

	// Service
	// the service
	// Constraints:
	//    - required
	//    - oneof:[MANAGEMENT_WEB,AP_PORTAL,HOTSPOT,COMMUNICATOR]
	Service *string `json:"service" validate:"required,oneof=MANAGEMENT_WEB AP_PORTAL HOTSPOT COMMUNICATOR"`
}

func NewServiceCertificate() *ServiceCertificate {
	serviceCertificateType := new(ServiceCertificate)
	return serviceCertificateType
}

func NewServiceCertificateWithDefaults() *ServiceCertificate {
	serviceCertificateType := new(ServiceCertificate)
	return serviceCertificateType
}

type ServiceCertificates []*ServiceCertificate

func NewServiceCertificates() *ServiceCertificates {
	serviceCertificatesType := make(ServiceCertificates, 0)
	return &serviceCertificatesType
}

func NewServiceCertificatesWithDefaults() *ServiceCertificates {
	serviceCertificatesType := make(ServiceCertificates, 0)
	return &serviceCertificatesType
}

type TrustedCAChain struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the trusted CA chain certificates
	Id *string `json:"id,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

func NewTrustedCAChain() *TrustedCAChain {
	trustedCAChainType := new(TrustedCAChain)
	return trustedCAChainType
}

func NewTrustedCAChainWithDefaults() *TrustedCAChain {
	trustedCAChainType := new(TrustedCAChain)
	return trustedCAChainType
}

type TrustedCAChainCertList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*TrustedCAChainCertListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewTrustedCAChainCertList() *TrustedCAChainCertList {
	trustedCAChainCertListType := new(TrustedCAChainCertList)
	return trustedCAChainCertListType
}

func NewTrustedCAChainCertListWithDefaults() *TrustedCAChainCertList {
	trustedCAChainCertListType := new(TrustedCAChainCertList)
	return trustedCAChainCertListType
}

type TrustedCAChainCertListType struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the trusted CA chain certificate
	Id *string `json:"id,omitempty"`

	// Information
	// Information of the certificates
	Information *string `json:"information,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	ModifiedDateTime *common.NormalNameAllowBlank `json:"modifiedDateTime,omitempty"`

	ModifierUsername *common.NormalNameAllowBlank `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

func NewTrustedCAChainCertListType() *TrustedCAChainCertListType {
	trustedCAChainCertListTypeType := new(TrustedCAChainCertListType)
	return trustedCAChainCertListTypeType
}

func NewTrustedCAChainCertListTypeWithDefaults() *TrustedCAChainCertListType {
	trustedCAChainCertListTypeType := new(TrustedCAChainCertListType)
	return trustedCAChainCertListTypeType
}
