package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
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
	// Constraints:
	//    - nullable
	IntermediateData []string `json:"intermediateData,omitempty" validate:"omitempty,dive"`

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

func NewWSGCertificate() *WSGCertificate {
	m := new(WSGCertificate)
	return m
}

type WSGCertificateList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGCertificateList() *WSGCertificateList {
	m := new(WSGCertificateList)
	return m
}

type WSGCertificateListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the certificate
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGCertificateListType() *WSGCertificateListType {
	m := new(WSGCertificateListType)
	return m
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

func NewWSGCertificatesSigningRequest() *WSGCertificatesSigningRequest {
	m := new(WSGCertificatesSigningRequest)
	return m
}

type WSGCertificateCertSetting struct {
	// ServiceCertificates
	// Certificate Setting of the service
	ServiceCertificates []*WSGCertificateServiceCertificate `json:"serviceCertificates,omitempty"`
}

func NewWSGCertificateCertSetting() *WSGCertificateCertSetting {
	m := new(WSGCertificateCertSetting)
	return m
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
	// Constraints:
	//    - nullable
	IntermediateData []string `json:"intermediateData,omitempty" validate:"omitempty,dive"`

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

func NewWSGCertificateClientCert() *WSGCertificateClientCert {
	m := new(WSGCertificateClientCert)
	return m
}

type WSGCertificateClientCertList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateClientCertListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGCertificateClientCertList() *WSGCertificateClientCertList {
	m := new(WSGCertificateClientCertList)
	return m
}

type WSGCertificateClientCertListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the client certificate
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGCertificateClientCertListType() *WSGCertificateClientCertListType {
	m := new(WSGCertificateClientCertListType)
	return m
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
	// Constraints:
	//    - nullable
	IntermediateData []string `json:"intermediateData,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

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

func NewWSGCertificateCreateCert() *WSGCertificateCreateCert {
	m := new(WSGCertificateCreateCert)
	return m
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
	// Constraints:
	//    - nullable
	IntermediateData []string `json:"intermediateData,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// PrivateKeyData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	// Constraints:
	//    - required
	PrivateKeyData *string `json:"privateKeyData" validate:"required"`

	// RootData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	RootData *string `json:"rootData,omitempty"`
}

func NewWSGCertificateCreateClientCert() *WSGCertificateCreateClientCert {
	m := new(WSGCertificateCreateClientCert)
	return m
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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Organization
	// Organization of the certificates signing request
	// Constraints:
	//    - required
	//    - max:64
	Organization *string `json:"organization" validate:"required,max=64"`

	// OrganizationUnit
	// Organization unit of the certificates signing request
	// Constraints:
	//    - max:64
	OrganizationUnit *string `json:"organizationUnit,omitempty" validate:"max=64"`

	// State
	// State of the certificates signing request
	// Constraints:
	//    - required
	//    - max:128
	State *string `json:"state" validate:"required,max=128"`
}

func NewWSGCertificateCreateCSR() *WSGCertificateCreateCSR {
	m := new(WSGCertificateCreateCSR)
	return m
}

type WSGCertificateCreateTrustedCAChain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	// Constraints:
	//    - nullable
	InterCertData []string `json:"interCertData,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	// Constraints:
	//    - required
	RootCertData *string `json:"rootCertData" validate:"required"`
}

func NewWSGCertificateCreateTrustedCAChain() *WSGCertificateCreateTrustedCAChain {
	m := new(WSGCertificateCreateTrustedCAChain)
	return m
}

type WSGCertificateCsrList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateCsrListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGCertificateCsrList() *WSGCertificateCsrList {
	m := new(WSGCertificateCsrList)
	return m
}

type WSGCertificateCsrListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the certificates signing request
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGCertificateCsrListType() *WSGCertificateCsrListType {
	m := new(WSGCertificateCsrListType)
	return m
}

type WSGCertificateDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGCertificateDeleteBulk() *WSGCertificateDeleteBulk {
	m := new(WSGCertificateDeleteBulk)
	return m
}

type WSGCertificateModifyTrustedCAChain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Information
	// Information of the certificates
	Information *string `json:"information,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	// Constraints:
	//    - nullable
	InterCertData []string `json:"interCertData,omitempty" validate:"omitempty,dive"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

func NewWSGCertificateModifyTrustedCAChain() *WSGCertificateModifyTrustedCAChain {
	m := new(WSGCertificateModifyTrustedCAChain)
	return m
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

func NewWSGCertificateServiceCertificate() *WSGCertificateServiceCertificate {
	m := new(WSGCertificateServiceCertificate)
	return m
}

type WSGCertificateServiceCertificates []*WSGCertificateServiceCertificate

func MakeWSGCertificateServiceCertificates() WSGCertificateServiceCertificates {
	m := make(WSGCertificateServiceCertificates, 0)
	return m
}

type WSGCertificateTrustedCAChain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the trusted CA chain certificates
	Id *string `json:"id,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	// Constraints:
	//    - nullable
	InterCertData []string `json:"interCertData,omitempty" validate:"omitempty,dive"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

func NewWSGCertificateTrustedCAChain() *WSGCertificateTrustedCAChain {
	m := new(WSGCertificateTrustedCAChain)
	return m
}

type WSGCertificateTrustedCAChainCertList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateTrustedCAChainCertListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGCertificateTrustedCAChainCertList() *WSGCertificateTrustedCAChainCertList {
	m := new(WSGCertificateTrustedCAChainCertList)
	return m
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
	// Constraints:
	//    - nullable
	InterCertData []string `json:"interCertData,omitempty" validate:"omitempty,dive"`

	ModifiedDateTime *WSGCommonNormalNameAllowBlank `json:"modifiedDateTime,omitempty"`

	ModifierUsername *WSGCommonNormalNameAllowBlank `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

func NewWSGCertificateTrustedCAChainCertListType() *WSGCertificateTrustedCAChainCertListType {
	m := new(WSGCertificateTrustedCAChainCertListType)
	return m
}

// AddCertstoreCertificate
//
// Use this API command to create an installed certificate.
//
// Request Body:
//	 - body *WSGCertificateCreateCert
func (s *WSGCertificateService) AddCertstoreCertificate(ctx context.Context, body *WSGCertificateCreateCert) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddCertstoreCertificate, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddCertstoreClientCert
//
// Use this API command to create a client certificate.
//
// Request Body:
//	 - body *WSGCertificateCreateClientCert
func (s *WSGCertificateService) AddCertstoreClientCert(ctx context.Context, body *WSGCertificateCreateClientCert) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddCertstoreClientCert, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddCertstoreCsr
//
// Use this API command to create a certificates signing request.
//
// Request Body:
//	 - body *WSGCertificateCreateCSR
func (s *WSGCertificateService) AddCertstoreCsr(ctx context.Context, body *WSGCertificateCreateCSR) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddCertstoreCsr, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddCertstoreTrustedCAChainCert
//
// Use this API command to create trusted CA chain certificates.
//
// Request Body:
//	 - body *WSGCertificateCreateTrustedCAChain
func (s *WSGCertificateService) AddCertstoreTrustedCAChainCert(ctx context.Context, body *WSGCertificateCreateTrustedCAChain) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddCertstoreTrustedCAChainCert, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteCertstoreCertificateById
//
// Use this API command to delete an installed certificate.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) DeleteCertstoreCertificateById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteCertstoreCertificateById, true)
	req.SetPathParameter("id", id)
}

// DeleteCertstoreClientCertById
//
// Use this API command to delete a client certificate.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) DeleteCertstoreClientCertById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteCertstoreClientCertById, true)
	req.SetPathParameter("id", id)
}

// DeleteCertstoreCsrById
//
// Use this API command to delete a certificates signing request.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) DeleteCertstoreCsrById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteCertstoreCsrById, true)
	req.SetPathParameter("id", id)
}

// DeleteCertstoreTrustedCAChainCert
//
// Use this API command to delete bulk trusted CA chain certificates.
//
// Request Body:
//	 - body *WSGCertificateDeleteBulk
func (s *WSGCertificateService) DeleteCertstoreTrustedCAChainCert(ctx context.Context, body *WSGCertificateDeleteBulk) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteCertstoreTrustedCAChainCert, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteCertstoreTrustedCAChainCertById
//
// Use this API command to delete a trusted CA chain certificate.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) DeleteCertstoreTrustedCAChainCertById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteCertstoreTrustedCAChainCertById, true)
	req.SetPathParameter("id", id)
}

// FindCertstoreCertificate
//
// Use this API command to retrieve list of installed certificates.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGCertificateService) FindCertstoreCertificate(ctx context.Context, optionalParams map[string]interface{}) (*WSGCertificateList, error) {
	var (
		req  *APIRequest
		resp *WSGCertificateList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindCertstoreCertificate, true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
}

// FindCertstoreCertificateById
//
// Use this API command to retrieve an installed certificate.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) FindCertstoreCertificateById(ctx context.Context, id string) (*WSGCertificate, error) {
	var (
		req  *APIRequest
		resp *WSGCertificate
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindCertstoreCertificateById, true)
	req.SetPathParameter("id", id)
}

// FindCertstoreClientCert
//
// Use this API command to retrieve list of client certificates.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGCertificateService) FindCertstoreClientCert(ctx context.Context, optionalParams map[string]interface{}) (*WSGCertificateClientCertList, error) {
	var (
		req  *APIRequest
		resp *WSGCertificateClientCertList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindCertstoreClientCert, true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
}

// FindCertstoreClientCertById
//
// Use this API command to retrieve a client certificate.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) FindCertstoreClientCertById(ctx context.Context, id string) (*WSGCertificateClientCert, error) {
	var (
		req  *APIRequest
		resp *WSGCertificateClientCert
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindCertstoreClientCertById, true)
	req.SetPathParameter("id", id)
}

// FindCertstoreCsr
//
// Use this API command to retrieve list of certificates signing request.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGCertificateService) FindCertstoreCsr(ctx context.Context, optionalParams map[string]interface{}) (*WSGCertificateCsrList, error) {
	var (
		req  *APIRequest
		resp *WSGCertificateCsrList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindCertstoreCsr, true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
}

// FindCertstoreCsrById
//
// Use this API command to retrieve a certificates signing request.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) FindCertstoreCsrById(ctx context.Context, id string) (*WSGCertificatesSigningRequest, error) {
	var (
		req  *APIRequest
		resp *WSGCertificatesSigningRequest
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindCertstoreCsrById, true)
	req.SetPathParameter("id", id)
}

// FindCertstoreSetting
//
// Use this API command to retrieve certificate setting.
func (s *WSGCertificateService) FindCertstoreSetting(ctx context.Context) (*WSGCertificateCertSetting, error) {
	var (
		req  *APIRequest
		resp *WSGCertificateCertSetting
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindCertstoreSetting, true)
}

// FindCertstoreTrustedCAChainCert
//
// Use this API command to retrieve list of installed trusted CA chain certificates.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCert(ctx context.Context, optionalParams map[string]interface{}) (*WSGCertificateTrustedCAChainCertList, error) {
	var (
		req  *APIRequest
		resp *WSGCertificateTrustedCAChainCertList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindCertstoreTrustedCAChainCert, true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
}

// FindCertstoreTrustedCAChainCertById
//
// Use this API command to retrieve an installed trusted CA chain certificates.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCertById(ctx context.Context, id string) (*WSGCertificateTrustedCAChain, error) {
	var (
		req  *APIRequest
		resp *WSGCertificateTrustedCAChain
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindCertstoreTrustedCAChainCertById, true)
	req.SetPathParameter("id", id)
}

// PartialUpdateCertstoreSetting
//
// Use this API command to Modify the Certificate Setting.
//
// Request Body:
//	 - body *WSGCertificateCertSetting
func (s *WSGCertificateService) PartialUpdateCertstoreSetting(ctx context.Context, body *WSGCertificateCertSetting) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateCertstoreSetting, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateCertstoreSettingServiceCertificates
//
// Use this API command to Modify serviceCertificates of the Certificate Setting.
//
// Request Body:
//	 - body WSGCertificateServiceCertificates
func (s *WSGCertificateService) PartialUpdateCertstoreSettingServiceCertificates(ctx context.Context, body WSGCertificateServiceCertificates) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateCertstoreSettingServiceCertificates, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateCertstoreTrustedCAChainCertById
//
// Use this API command to patch a trusted CA chain certificates.
//
// Request Body:
//	 - body *WSGCertificateModifyTrustedCAChain
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) PartialUpdateCertstoreTrustedCAChainCertById(ctx context.Context, body *WSGCertificateModifyTrustedCAChain, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateCertstoreTrustedCAChainCertById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
}
