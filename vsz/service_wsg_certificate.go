package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/certificate"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
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
	serv := new(WSGCertificateService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddCertstoreCertificate
//
// Use this API command to create an installed certificate.
//
// Request Body:
//	 - body *certificate.CreateCert
func (s *WSGCertificateService) AddCertstoreCertificate(ctx context.Context, body *certificate.CreateCert) (*common.CreateResult, error) {
}

// AddCertstoreClientCert
//
// Use this API command to create a client certificate.
//
// Request Body:
//	 - body *certificate.CreateClientCert
func (s *WSGCertificateService) AddCertstoreClientCert(ctx context.Context, body *certificate.CreateClientCert) (*common.CreateResult, error) {
}

// AddCertstoreCsr
//
// Use this API command to create a certificates signing request.
//
// Request Body:
//	 - body *certificate.CreateCSR
func (s *WSGCertificateService) AddCertstoreCsr(ctx context.Context, body *certificate.CreateCSR) (*common.CreateResult, error) {
}

// AddCertstoreTrustedCAChainCert
//
// Use this API command to create trusted CA chain certificates.
//
// Request Body:
//	 - body *certificate.CreateTrustedCAChain
func (s *WSGCertificateService) AddCertstoreTrustedCAChainCert(ctx context.Context, body *certificate.CreateTrustedCAChain) (*common.CreateResult, error) {
}

// DeleteCertstoreCertificateById
//
// Use this API command to delete an installed certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) DeleteCertstoreCertificateById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteCertstoreClientCertById
//
// Use this API command to delete a client certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) DeleteCertstoreClientCertById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteCertstoreCsrById
//
// Use this API command to delete a certificates signing request.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) DeleteCertstoreCsrById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteCertstoreTrustedCAChainCert
//
// Use this API command to delete bulk trusted CA chain certificates.
//
// Request Body:
//	 - body *certificate.DeleteBulk
func (s *WSGCertificateService) DeleteCertstoreTrustedCAChainCert(ctx context.Context, body *certificate.DeleteBulk) (*common.EmptyResult, error) {
}

// DeleteCertstoreTrustedCAChainCertById
//
// Use this API command to delete a trusted CA chain certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) DeleteCertstoreTrustedCAChainCertById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindCertstoreCertificate
//
// Use this API command to retrieve list of installed certificates.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGCertificateService) FindCertstoreCertificate(ctx context.Context, qIndex string, qListSize string) (*certificate.CertificateList, error) {
}

// FindCertstoreCertificateById
//
// Use this API command to retrieve an installed certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) FindCertstoreCertificateById(ctx context.Context, pId string) (*certificate.Certificate, error) {
}

// FindCertstoreClientCert
//
// Use this API command to retrieve list of client certificates.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGCertificateService) FindCertstoreClientCert(ctx context.Context, qIndex string, qListSize string) (*certificate.ClientCertList, error) {
}

// FindCertstoreClientCertById
//
// Use this API command to retrieve a client certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) FindCertstoreClientCertById(ctx context.Context, pId string) (*certificate.ClientCert, error) {
}

// FindCertstoreCsr
//
// Use this API command to retrieve list of certificates signing request.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGCertificateService) FindCertstoreCsr(ctx context.Context, qIndex string, qListSize string) (*certificate.CsrList, error) {
}

// FindCertstoreCsrById
//
// Use this API command to retrieve a certificates signing request.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) FindCertstoreCsrById(ctx context.Context, pId string) (*certificate.CertificatesSigningRequest, error) {
}

// FindCertstoreSetting
//
// Use this API command to retrieve certificate setting.
func (s *WSGCertificateService) FindCertstoreSetting(ctx context.Context) (*certificate.CertSetting, error) {
}

// FindCertstoreTrustedCAChainCert
//
// Use this API command to retrieve list of installed trusted CA chain certificates.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCert(ctx context.Context, qIndex string, qListSize string) (*certificate.TrustedCAChainCertList, error) {
}

// FindCertstoreTrustedCAChainCertById
//
// Use this API command to retrieve an installed trusted CA chain certificates.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCertById(ctx context.Context, pId string) (*certificate.TrustedCAChain, error) {
}

// PartialUpdateCertstoreSetting
//
// Use this API command to Modify the Certificate Setting.
//
// Request Body:
//	 - body *certificate.CertSetting
func (s *WSGCertificateService) PartialUpdateCertstoreSetting(ctx context.Context, body *certificate.CertSetting) (*common.EmptyResult, error) {
}

// PartialUpdateCertstoreSettingServiceCertificates
//
// Use this API command to Modify serviceCertificates of the Certificate Setting.
//
// Request Body:
//	 - body certificate.ServiceCertificates
func (s *WSGCertificateService) PartialUpdateCertstoreSettingServiceCertificates(ctx context.Context, body certificate.ServiceCertificates) (*common.EmptyResult, error) {
}

// PartialUpdateCertstoreTrustedCAChainCertById
//
// Use this API command to patch a trusted CA chain certificates.
//
// Request Body:
//	 - body *certificate.ModifyTrustedCAChain
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) PartialUpdateCertstoreTrustedCAChainCertById(ctx context.Context, body *certificate.ModifyTrustedCAChain, pId string) (*common.EmptyResult, error) {
}
