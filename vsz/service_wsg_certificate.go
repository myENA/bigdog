package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/certificate"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGCertificateService struct {
	client *Client
}

func NewWSGCertificateService(client *Client) *WSGCertificateService {
	s := new(WSGCertificateService)
	s.client = client
	return s
}

func (ss *WSGService) WSGCertificateService() *WSGCertificateService {
	serv := new(WSGCertificateService)
	serv.client = ss.client
	return serv
}

// FindCertstoreCertificate
//
// Use this API command to retrieve list of installed certificates.
func (s *WSGCertificateService) FindCertstoreCertificate(ctx context.Context, qIndex string, qListSize string) (*certificate.CertificateList, error) {
}

// FindCertstoreCertificateById
//
// Use this API command to retrieve an installed certificate.
func (s *WSGCertificateService) FindCertstoreCertificateById(ctx context.Context, pId string) (*certificate.Certificate, error) {
}

// FindCertstoreClientCert
//
// Use this API command to retrieve list of client certificates.
func (s *WSGCertificateService) FindCertstoreClientCert(ctx context.Context, qIndex string, qListSize string) (*certificate.ClientCertList, error) {
}

// FindCertstoreClientCertById
//
// Use this API command to retrieve a client certificate.
func (s *WSGCertificateService) FindCertstoreClientCertById(ctx context.Context, pId string) (*certificate.ClientCert, error) {
}

// FindCertstoreCsr
//
// Use this API command to retrieve list of certificates signing request.
func (s *WSGCertificateService) FindCertstoreCsr(ctx context.Context, qIndex string, qListSize string) (*certificate.CsrList, error) {
}

// FindCertstoreCsrById
//
// Use this API command to retrieve a certificates signing request.
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
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCert(ctx context.Context, qIndex string, qListSize string) (*certificate.TrustedCAChainCertList, error) {
}

// FindCertstoreTrustedCAChainCertById
//
// Use this API command to retrieve an installed trusted CA chain certificates.
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCertById(ctx context.Context, pId string) (*certificate.TrustedCAChain, error) {
}

// PartialUpdateCertstoreSettingServiceCertificates
//
// Use this API command to Modify serviceCertificates of the Certificate Setting.
//
// Request Body:
//	 - body certificate.ServiceCertificates
func (s *WSGCertificateService) PartialUpdateCertstoreSettingServiceCertificates(ctx context.Context, body certificate.ServiceCertificates) (*common.EmptyResult, error) {
}
