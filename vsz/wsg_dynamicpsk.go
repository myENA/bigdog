package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGDynamicPSKService struct {
	apiClient *APIClient
}

func NewWSGDynamicPSKService(c *APIClient) *WSGDynamicPSKService {
	s := new(WSGDynamicPSKService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDynamicPSKService() *WSGDynamicPSKService {
	return NewWSGDynamicPSKService(ss.apiClient)
}

// AddRkszonesWlansDpskBatchGenUnboundById
//
// Use this API command to batch generate DPSKs of a WLAN. You can either specify passphrases or not. If the amount is bigger than 1, system will generate usernames with index. e.g. student-1, student-2, ...etc.
//
// Request Body:
//	 - body *WSGDPSKBatchGenUnbound
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskBatchGenUnboundById(ctx context.Context, body *WSGDPSKBatchGenUnbound, pId string, pZoneId string) (*WSGDPSKGetDpskResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRkszonesWlansDpskById
//
// Use this API command to delete DPSKs of a WLAN.
//
// Request Body:
//	 - body *WSGDPSKDeleteDPSKs
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskById(ctx context.Context, body *WSGDPSKDeleteDPSKs, pId string, pZoneId string) (*WSGDPSKDeleteDpskResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskUploadById(ctx context.Context, pId string, pZoneId string) (*WSGDPSKGetDpskResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesDeleteExpiredDpskByZoneId
//
// Use this API command to retrieve interval of delete expired DPSK of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, pZoneId string) (*WSGDPSKDeleteExpiredDpskConfig, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesDownloadDpskCsvSample
//
// Use this API command to download DPSK CSV sample.
//
// Query Parameters:
// - qType string
//		- nullable
func (s *WSGDynamicPSKService) FindRkszonesDownloadDpskCsvSample(ctx context.Context, qType string) ([]byte, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesDpskByZoneId
//
// Use this API command to retrieve DPSK info of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDpskByZoneId(ctx context.Context, pZoneId string) (*WSGDPSKGetDpskInfoList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesDpskEnabledWlansByZoneId
//
// Use this API command to retrieve DPSK enabled WLAN info of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDpskEnabledWlansByZoneId(ctx context.Context, pZoneId string) (*WSGDPSKGetDpskEnabledWlans, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskByDpskId(ctx context.Context, pDpskId string, pId string, pZoneId string) (*WSGDPSKGetDpskInfoList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskById(ctx context.Context, pId string, pZoneId string) (*WSGDPSKGetDpskInfoList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesWlansDpskByDpskId
//
// Use this API command to update DPSK info.
//
// Request Body:
//	 - body *WSGDPSKUpdateDpsk
//
// Path Parameters:
// - pDpskId string
//		- required
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) PartialUpdateRkszonesWlansDpskByDpskId(ctx context.Context, body *WSGDPSKUpdateDpsk, pDpskId string, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesDeleteExpiredDpskByZoneId
//
// Use this API command to modify interval of delete expired DPSK of a zone.
//
// Request Body:
//	 - body *WSGDPSKModifyDeleteExpiredDpsk
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDynamicPSKService) UpdateRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, body *WSGDPSKModifyDeleteExpiredDpsk, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
