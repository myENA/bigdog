package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGCertificateService struct {
	apiClient *APIClient
}

func NewWSGCertificateService(c *APIClient) *WSGCertificateService {
	s := new(WSGCertificateService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGCertificateService() *WSGCertificateService {
	return NewWSGCertificateService(ss.apiClient)
}

type WSGCertificate struct {
	CertificasSigningRequest *WSGCommonGenericRef `json:"certificasSigningRequest,omitempty"`

	// Data
	// Data of the certificate
	Data *string `json:"data,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the certificate
	Id *string `json:"id,omitempty"`

	// Information
	// Information of the certificate
	Information *string `json:"information,omitempty"`

	// IntermediateData
	// Intermediate data of the certificate
	IntermediateData []string `json:"intermediateData,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

type WSGCertificateList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGCertificateListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the certificate
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

type WSGCertificatesSigningRequest struct {
	// City
	// City of the certificates signing request
	City *string `json:"city,omitempty"`

	// CommonName
	// Common name of the certificates signing request
	CommonName *string `json:"commonName,omitempty"`

	// CountryCode
	// Country code of the certificates signing request
	CountryCode *string `json:"countryCode,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Email
	// Email of the certificates signing request
	Email *string `json:"email,omitempty"`

	// Id
	// Identifier of the certificates signing request
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

type WSGCertificateCertSetting struct {
	// ServiceCertificates
	// Certificate Setting of the service
	ServiceCertificates []*WSGCertificateServiceCertificate `json:"serviceCertificates,omitempty"`
}

type WSGCertificateClientCert struct {
	// Data
	// Data of the client certificate
	Data *string `json:"data,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the client certificate
	Id *string `json:"id,omitempty"`

	// Information
	// Information of the client certificate
	Information *string `json:"information,omitempty"`

	// IntermediateData
	// Intermediate data of the client certificate
	IntermediateData []string `json:"intermediateData,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

type WSGCertificateClientCertList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateClientCertListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGCertificateClientCertListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the client certificate
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

type WSGCertificateCreateCert struct {
	CertificasSigningRequest *WSGCommonGenericRef `json:"certificasSigningRequest,omitempty"`

	// Data
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	// Constraints:
	//    - required
	Data *string `json:"data" validate:"required"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// IntermediateData
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	IntermediateData []string `json:"intermediateData,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

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

type WSGCertificateCreateClientCert struct {
	// Data
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	// Constraints:
	//    - required
	Data *string `json:"data" validate:"required"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// IntermediateData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	IntermediateData []string `json:"intermediateData,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PrivateKeyData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	// Constraints:
	//    - required
	PrivateKeyData *string `json:"privateKeyData" validate:"required"`

	// RootData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	RootData *string `json:"rootData,omitempty"`
}

type WSGCertificateCreateCSR struct {
	// City
	// City of the certificates signing request
	// Constraints:
	//    - required
	//    - max:128
	City *string `json:"city" validate:"required,max=128"`

	// CommonName
	// Constraints:
	//    - required
	CommonName *WSGCommonFQDN `json:"commonName" validate:"required"`

	// CountryCode
	// Country code of the certificates signing request
	// Constraints:
	//    - required
	CountryCode *string `json:"countryCode" validate:"required"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Email
	// Constraints:
	//    - required
	Email *WSGCommonEmail `json:"email" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

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

type WSGCertificateCreateTrustedCAChain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	// Constraints:
	//    - required
	RootCertData *string `json:"rootCertData" validate:"required"`
}

type WSGCertificateCsrList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateCsrListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGCertificateCsrListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the certificates signing request
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

type WSGCertificateDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

type WSGCertificateModifyTrustedCAChain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Information
	// Information of the certificates
	Information *string `json:"information,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

type WSGCertificateServiceCertificate struct {
	Certificate *WSGCommonGenericRef `json:"certificate,omitempty"`

	// Service
	// the service
	// Constraints:
	//    - required
	//    - oneof:[MANAGEMENT_WEB,AP_PORTAL,HOTSPOT,COMMUNICATOR]
	Service *string `json:"service" validate:"required,oneof=MANAGEMENT_WEB AP_PORTAL HOTSPOT COMMUNICATOR"`
}

type WSGCertificateServiceCertificates []*WSGCertificateServiceCertificate

type WSGCertificateTrustedCAChain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the trusted CA chain certificates
	Id *string `json:"id,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

type WSGCertificateTrustedCAChainCertList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateTrustedCAChainCertListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGCertificateTrustedCAChainCertListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the trusted CA chain certificate
	Id *string `json:"id,omitempty"`

	// Information
	// Information of the certificates
	Information *string `json:"information,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	InterCertData []string `json:"interCertData,omitempty"`

	ModifiedDateTime *WSGCommonNormalNameAllowBlank `json:"modifiedDateTime,omitempty"`

	ModifierUsername *WSGCommonNormalNameAllowBlank `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

// AddCertstoreCertificate
//
// Use this API command to create an installed certificate.
//
// Request Body:
//	 - body *WSGCertificateCreateCert
func (s *WSGCertificateService) AddCertstoreCertificate(ctx context.Context, body *WSGCertificateCreateCert) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddCertstoreClientCert
//
// Use this API command to create a client certificate.
//
// Request Body:
//	 - body *WSGCertificateCreateClientCert
func (s *WSGCertificateService) AddCertstoreClientCert(ctx context.Context, body *WSGCertificateCreateClientCert) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddCertstoreCsr
//
// Use this API command to create a certificates signing request.
//
// Request Body:
//	 - body *WSGCertificateCreateCSR
func (s *WSGCertificateService) AddCertstoreCsr(ctx context.Context, body *WSGCertificateCreateCSR) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddCertstoreTrustedCAChainCert
//
// Use this API command to create trusted CA chain certificates.
//
// Request Body:
//	 - body *WSGCertificateCreateTrustedCAChain
func (s *WSGCertificateService) AddCertstoreTrustedCAChainCert(ctx context.Context, body *WSGCertificateCreateTrustedCAChain) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteCertstoreCertificateById
//
// Use this API command to delete an installed certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) DeleteCertstoreCertificateById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteCertstoreClientCertById
//
// Use this API command to delete a client certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) DeleteCertstoreClientCertById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteCertstoreCsrById
//
// Use this API command to delete a certificates signing request.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) DeleteCertstoreCsrById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteCertstoreTrustedCAChainCert
//
// Use this API command to delete bulk trusted CA chain certificates.
//
// Request Body:
//	 - body *WSGCertificateDeleteBulk
func (s *WSGCertificateService) DeleteCertstoreTrustedCAChainCert(ctx context.Context, body *WSGCertificateDeleteBulk) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteCertstoreTrustedCAChainCertById
//
// Use this API command to delete a trusted CA chain certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) DeleteCertstoreTrustedCAChainCertById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCertstoreCertificate
//
// Use this API command to retrieve list of installed certificates.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGCertificateService) FindCertstoreCertificate(ctx context.Context, qIndex string, qListSize string) (*WSGCertificateList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCertstoreCertificateById
//
// Use this API command to retrieve an installed certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) FindCertstoreCertificateById(ctx context.Context, pId string) (*WSGCertificate, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCertstoreClientCert
//
// Use this API command to retrieve list of client certificates.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGCertificateService) FindCertstoreClientCert(ctx context.Context, qIndex string, qListSize string) (*WSGCertificateClientCertList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCertstoreClientCertById
//
// Use this API command to retrieve a client certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) FindCertstoreClientCertById(ctx context.Context, pId string) (*WSGCertificateClientCert, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCertstoreCsr
//
// Use this API command to retrieve list of certificates signing request.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGCertificateService) FindCertstoreCsr(ctx context.Context, qIndex string, qListSize string) (*WSGCertificateCsrList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCertstoreCsrById
//
// Use this API command to retrieve a certificates signing request.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) FindCertstoreCsrById(ctx context.Context, pId string) (*WSGCertificatesSigningRequest, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCertstoreSetting
//
// Use this API command to retrieve certificate setting.
func (s *WSGCertificateService) FindCertstoreSetting(ctx context.Context) (*WSGCertificateCertSetting, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCertstoreTrustedCAChainCert
//
// Use this API command to retrieve list of installed trusted CA chain certificates.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCert(ctx context.Context, qIndex string, qListSize string) (*WSGCertificateTrustedCAChainCertList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCertstoreTrustedCAChainCertById
//
// Use this API command to retrieve an installed trusted CA chain certificates.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCertById(ctx context.Context, pId string) (*WSGCertificateTrustedCAChain, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateCertstoreSetting
//
// Use this API command to Modify the Certificate Setting.
//
// Request Body:
//	 - body *WSGCertificateCertSetting
func (s *WSGCertificateService) PartialUpdateCertstoreSetting(ctx context.Context, body *WSGCertificateCertSetting) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateCertstoreSettingServiceCertificates
//
// Use this API command to Modify serviceCertificates of the Certificate Setting.
//
// Request Body:
//	 - body WSGCertificateServiceCertificates
func (s *WSGCertificateService) PartialUpdateCertstoreSettingServiceCertificates(ctx context.Context, body WSGCertificateServiceCertificates) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateCertstoreTrustedCAChainCertById
//
// Use this API command to patch a trusted CA chain certificates.
//
// Request Body:
//	 - body *WSGCertificateModifyTrustedCAChain
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) PartialUpdateCertstoreTrustedCAChainCertById(ctx context.Context, body *WSGCertificateModifyTrustedCAChain, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
