package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpsk"
)

type WSGDynamicPSKService struct {
    client *Client
}

func NewWSGDynamicPSKService (client *Client) *WSGDynamicPSKService {
    s := new(WSGDynamicPSKService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDynamicPSKService () *WSGDynamicPSKService {
    serv := new(WSGDynamicPSKService)
    serv.client = ss.client
    return serv
}

func (s *WSGDynamicPSKService) AddRkszonesWlansDpskBatchGenUnboundById (ctx context.Context, id string, zoneId string) (*dpsk.GetDpskResult, error) {
}

func (s *WSGDynamicPSKService) AddRkszonesWlansDpskUploadById (ctx context.Context, id string, zoneId string) (*dpsk.GetDpskResult, error) {
}

func (s *WSGDynamicPSKService) FindRkszonesDeleteExpiredDpskByZoneId (ctx context.Context, zoneId string) (*dpsk.DeleteExpiredDpskConfig, error) {
}

func (s *WSGDynamicPSKService) FindRkszonesDownloadDpskCsvSample (ctx context.Context) (json.RawMessage, error) {
}

func (s *WSGDynamicPSKService) FindRkszonesDpskByZoneId (ctx context.Context, zoneId string) (*dpsk.GetDpskInfoList, error) {
}

func (s *WSGDynamicPSKService) FindRkszonesDpskEnabledWlansByZoneId (ctx context.Context, zoneId string) (*dpsk.GetDpskEnabledWlans, error) {
}

func (s *WSGDynamicPSKService) FindRkszonesWlansDpskByDpskId (ctx context.Context, dpskId string, id string, zoneId string) (*dpsk.GetDpskInfoList, error) {
}

func (s *WSGDynamicPSKService) FindRkszonesWlansDpskById (ctx context.Context, id string, zoneId string) (*dpsk.GetDpskInfoList, error) {
}

