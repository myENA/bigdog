package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/calea"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGCALEAService struct {
	client *Client
}

func NewWSGCALEAService(client *Client) *WSGCALEAService {
	s := new(WSGCALEAService)
	s.client = client
	return s
}

func (ss *WSGService) WSGCALEAService() *WSGCALEAService {
	serv := new(WSGCALEAService)
	serv.client = ss.client
	return serv
}

// AddSystemCaleaMac
//
// Use this API command to add specified CALEA UE MACs
//
// Request Body:
//	 - body *calea.CaleaMacListRq
func (s *WSGCALEAService) AddSystemCaleaMac(ctx context.Context, body *calea.CaleaMacListRq) (*common.EmptyResult, error) {
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
