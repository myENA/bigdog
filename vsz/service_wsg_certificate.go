package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/certificate"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"gopkg.in/go-playground/validator.v9"
)

type WSGCertificateService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGCertificateService(c *APIClient) *WSGCertificateService {
	s := new(WSGCertificateService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGCertificateService() *WSGCertificateService {
	return NewWSGCertificateService(ss.apiClient)
}

// AddCertstoreCertificate
//
// Use this API command to create an installed certificate.
//
// Request Body:
//	 - body *certificate.CreateCert
func (s *WSGCertificateService) AddCertstoreCertificate(ctx context.Context, body *certificate.CreateCert) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddCertstoreClientCert
//
// Use this API command to create a client certificate.
//
// Request Body:
//	 - body *certificate.CreateClientCert
func (s *WSGCertificateService) AddCertstoreClientCert(ctx context.Context, body *certificate.CreateClientCert) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddCertstoreCsr
//
// Use this API command to create a certificates signing request.
//
// Request Body:
//	 - body *certificate.CreateCSR
func (s *WSGCertificateService) AddCertstoreCsr(ctx context.Context, body *certificate.CreateCSR) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddCertstoreTrustedCAChainCert
//
// Use this API command to create trusted CA chain certificates.
//
// Request Body:
//	 - body *certificate.CreateTrustedCAChain
func (s *WSGCertificateService) AddCertstoreTrustedCAChainCert(ctx context.Context, body *certificate.CreateTrustedCAChain) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// DeleteCertstoreCertificateById
//
// Use this API command to delete an installed certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) DeleteCertstoreCertificateById(ctx context.Context, pId string) (*common.EmptyResult, error) {
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
func (s *WSGCertificateService) DeleteCertstoreClientCertById(ctx context.Context, pId string) (*common.EmptyResult, error) {
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
func (s *WSGCertificateService) DeleteCertstoreCsrById(ctx context.Context, pId string) (*common.EmptyResult, error) {
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
//	 - body *certificate.DeleteBulk
func (s *WSGCertificateService) DeleteCertstoreTrustedCAChainCert(ctx context.Context, body *certificate.DeleteBulk) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// DeleteCertstoreTrustedCAChainCertById
//
// Use this API command to delete a trusted CA chain certificate.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGCertificateService) DeleteCertstoreTrustedCAChainCertById(ctx context.Context, pId string) (*common.EmptyResult, error) {
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
func (s *WSGCertificateService) FindCertstoreCertificate(ctx context.Context, qIndex string, qListSize string) (*certificate.CertificateList, error) {
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
func (s *WSGCertificateService) FindCertstoreCertificateById(ctx context.Context, pId string) (*certificate.Certificate, error) {
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
func (s *WSGCertificateService) FindCertstoreClientCert(ctx context.Context, qIndex string, qListSize string) (*certificate.ClientCertList, error) {
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
func (s *WSGCertificateService) FindCertstoreClientCertById(ctx context.Context, pId string) (*certificate.ClientCert, error) {
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
func (s *WSGCertificateService) FindCertstoreCsr(ctx context.Context, qIndex string, qListSize string) (*certificate.CsrList, error) {
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
func (s *WSGCertificateService) FindCertstoreCsrById(ctx context.Context, pId string) (*certificate.CertificatesSigningRequest, error) {
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
func (s *WSGCertificateService) FindCertstoreSetting(ctx context.Context) (*certificate.CertSetting, error) {
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
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCert(ctx context.Context, qIndex string, qListSize string) (*certificate.TrustedCAChainCertList, error) {
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
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCertById(ctx context.Context, pId string) (*certificate.TrustedCAChain, error) {
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
//	 - body *certificate.CertSetting
func (s *WSGCertificateService) PartialUpdateCertstoreSetting(ctx context.Context, body *certificate.CertSetting) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// PartialUpdateCertstoreSettingServiceCertificates
//
// Use this API command to Modify serviceCertificates of the Certificate Setting.
//
// Request Body:
//	 - body certificate.ServiceCertificates
func (s *WSGCertificateService) PartialUpdateCertstoreSettingServiceCertificates(ctx context.Context, body certificate.ServiceCertificates) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if len(body) == 0 {
		return nil, errors.New("body cannot be empty")
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}
