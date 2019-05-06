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

func NewWSGCertificateService (client *Client) *WSGCertificateService {
    s := new(WSGCertificateService)
    s.client = client
    return s
}

func (ss *WSGService) WSGCertificateService () *WSGCertificateService {
    serv := new(WSGCertificateService)
    serv.client = ss.client
    return serv
}

func (s *WSGCertificateService) FindCertstoreCertificate (ctx context.Context) (*certificate.CertificateList, error) {
}

func (s *WSGCertificateService) FindCertstoreCertificateById (ctx context.Context, id string) (*certificate.Certificate, error) {
}

func (s *WSGCertificateService) FindCertstoreClientCert (ctx context.Context) (*certificate.ClientCertList, error) {
}

func (s *WSGCertificateService) FindCertstoreClientCertById (ctx context.Context, id string) (*certificate.ClientCert, error) {
}

func (s *WSGCertificateService) FindCertstoreCsr (ctx context.Context) (*certificate.CsrList, error) {
}

func (s *WSGCertificateService) FindCertstoreCsrById (ctx context.Context, id string) (*certificate.CertificatesSigningRequest, error) {
}

func (s *WSGCertificateService) FindCertstoreSetting (ctx context.Context) (*certificate.CertSetting, error) {
}

func (s *WSGCertificateService) FindCertstoreTrustedCAChainCert (ctx context.Context) (*certificate.TrustedCAChainCertList, error) {
}

func (s *WSGCertificateService) FindCertstoreTrustedCAChainCertById (ctx context.Context, id string) (*certificate.TrustedCAChain, error) {
}

func (s *WSGCertificateService) PartialUpdateCertstoreSettingServiceCertificates (ctx context.Context) (*common.EmptyResult, error) {
}

