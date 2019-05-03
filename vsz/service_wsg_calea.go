package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/calea"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGCALEAService struct {
    client *Client
}

func NewWSGCALEAService (client *Client) *WSGCALEAService {
    s := new(WSGCALEAService)
    s.client = client
    return s
}

func (ss *WSGService) WSGCALEAService () *WSGCALEAService {
    serv := new(WSGCALEAService)
    serv.client = ss.client
    return serv
}

func (s *WSGCALEAService) AddSystemCaleaMac (ctx context.Context) error {
}

func (s *WSGCALEAService) FindSystemCaleaCommonSetting (ctx context.Context) (calea.CaleaCommonSettingRsp, error) {
}

func (s *WSGCALEAService) FindSystemCaleaMacList (ctx context.Context) (calea.CaleaMacListRsp, error) {
}

