package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/calea"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGCALEAService struct {
	apiClient *APIClient
}

func NewWSGCALEAService(c *APIClient) *WSGCALEAService {
	s := new(WSGCALEAService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGCALEAService() *WSGCALEAService {
	serv := new(WSGCALEAService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddSystemCaleaCommonSetting
//
// Use this API command to set CALEA common setting.
//
// Request Body:
//	 - body *calea.CaleaCommonSettingRq
func (s *WSGCALEAService) AddSystemCaleaCommonSetting(ctx context.Context, body *calea.CaleaCommonSettingRq) (*common.EmptyResult, error) {
}

// AddSystemCaleaMac
//
// Use this API command to add specified CALEA UE MACs
//
// Request Body:
//	 - body *calea.CaleaMacListRq
func (s *WSGCALEAService) AddSystemCaleaMac(ctx context.Context, body *calea.CaleaMacListRq) (*common.EmptyResult, error) {
}

// AddSystemCaleaMacList
//
// Use this API command to upload csv file of CALEA UE MACs.
func (s *WSGCALEAService) AddSystemCaleaMacList(ctx context.Context) error {
}

// DeleteSystemCaleaMac
//
// Use this API command to delete specified CALEA UE MACs.
//
// Request Body:
//	 - body *calea.CaleaMacListRq
func (s *WSGCALEAService) DeleteSystemCaleaMac(ctx context.Context, body *calea.CaleaMacListRq) (*common.EmptyResult, error) {
}

// DeleteSystemCaleaMacList
//
// Use this API command to delete all CALEA UE MACs.
func (s *WSGCALEAService) DeleteSystemCaleaMacList(ctx context.Context) (*common.EmptyResult, error) {
}

// FindSystemCaleaCommonSetting
//
// Use this API command to get CALEA common setting.
func (s *WSGCALEAService) FindSystemCaleaCommonSetting(ctx context.Context) (*calea.CaleaCommonSettingRsp, error) {
}

// FindSystemCaleaMacList
//
// Use this API command to get all CALEA UE MACs.
func (s *WSGCALEAService) FindSystemCaleaMacList(ctx context.Context) (*calea.CaleaMacListRsp, error) {
}
