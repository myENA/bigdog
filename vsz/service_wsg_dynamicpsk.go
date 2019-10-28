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

func NewWSGDynamicPSKService(client *Client) *WSGDynamicPSKService {
	s := new(WSGDynamicPSKService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDynamicPSKService() *WSGDynamicPSKService {
	serv := new(WSGDynamicPSKService)
	serv.client = ss.client
	return serv
}

// AddRkszonesWlansDpskBatchGenUnboundById
//
// Use this API command to batch generate DPSKs of a WLAN. You can either specify passphrases or not. If the amount is bigger than 1, system will generate usernames with index. e.g. student-1, student-2, ...etc.
//
// Request Body:
//	 - body *dpsk.BatchGenUnbound
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskBatchGenUnboundById(ctx context.Context, body *dpsk.BatchGenUnbound, pId string, pZoneId string) (*dpsk.GetDpskResult, error) {
}

// AddRkszonesWlansDpskUploadById
//
// Use this API command to upload DPSK file of a WLAN (CSV file and Content-Type multipart/form-data ONLY). In curl, here is an example curl -b /tmp/headers.txt -k -X POST -v -F upload@myDpsk.csv 'https://127.0.0.1:8443/wsg/api/public/v7_0/rkszones/{zoneUuid}/wlans/{wlanId}/dpsk/upload'
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskUploadById(ctx context.Context, pId string, pZoneId string) (*dpsk.GetDpskResult, error) {
}

// FindRkszonesDeleteExpiredDpskByZoneId
//
// Use this API command to retrieve interval of delete expired DPSK of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, pZoneId string) (*dpsk.DeleteExpiredDpskConfig, error) {
}

// FindRkszonesDownloadDpskCsvSample
//
// Use this API command to download DPSK CSV sample.
//
// Query Parameters:
// - qType string
func (s *WSGDynamicPSKService) FindRkszonesDownloadDpskCsvSample(ctx context.Context, qType string) (json.RawMessage, error) {
}

// FindRkszonesDpskByZoneId
//
// Use this API command to retrieve DPSK info of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDpskByZoneId(ctx context.Context, pZoneId string) (*dpsk.GetDpskInfoList, error) {
}

// FindRkszonesDpskEnabledWlansByZoneId
//
// Use this API command to retrieve DPSK enabled WLAN info of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDpskEnabledWlansByZoneId(ctx context.Context, pZoneId string) (*dpsk.GetDpskEnabledWlans, error) {
}

// FindRkszonesWlansDpskByDpskId
//
// Use this API command to retrieve DPSK info.
//
// Path Parameters:
// - pDpskId string
//		- required
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskByDpskId(ctx context.Context, pDpskId string, pId string, pZoneId string) (*dpsk.GetDpskInfoList, error) {
}

// FindRkszonesWlansDpskById
//
// Use this API command to retrieve DPSK info of a WLAN.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskById(ctx context.Context, pId string, pZoneId string) (*dpsk.GetDpskInfoList, error) {
}
