package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDynamicpskService struct {
	apiClient *APIClient
}

func NewWSGDynamicpskService(c *APIClient) *WSGDynamicpskService {
	s := new(WSGDynamicpskService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDynamicpskService() *WSGDynamicpskService {
	return NewWSGDynamicpskService(ss.apiClient)
}

// AddRkszonesWlansDpskBatchGenUnboundById
//
// Use this API command to batch generate DPSKs of a WLAN. You can either specify passphrases or not. If the amount is bigger than 1, system will generate usernames with index. e.g. student-1, student-2, ...etc.
//
// Request Body:
//	 - body *WSGDPSKBatchGenUnbound
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicpskService) AddRkszonesWlansDpskBatchGenUnboundById(ctx context.Context, body *WSGDPSKBatchGenUnbound, id string, zoneId string) (*WSGDPSKGetDpskResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPSKGetDpskResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansDpskBatchGenUnboundById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansDpskById
//
// Use this API command to delete DPSKs of a WLAN.
//
// Request Body:
//	 - body *WSGDPSKDeleteDPSKs
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicpskService) AddRkszonesWlansDpskById(ctx context.Context, body *WSGDPSKDeleteDPSKs, id string, zoneId string) (*WSGDPSKDeleteDpskResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPSKDeleteDpskResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansDpskById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKDeleteDpskResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansDpskUploadById
//
// Use this API command to upload DPSK file of a WLAN (CSV file and Content-Type multipart/form-data ONLY).
//
// Request Body:
//	 - body []byte
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicpskService) AddRkszonesWlansDpskUploadById(ctx context.Context, body []byte, id string, zoneId string) (*WSGDPSKGetDpskResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPSKGetDpskResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansDpskUploadById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// FindDpskByQueryCriteria
//
// Query DPSKs with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDynamicpskService) FindDpskByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGDPSKQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPSKQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindDpskByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDeleteExpiredDpskByZoneId
//
// Use this API command to retrieve interval of delete expired DPSK of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicpskService) FindRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, zoneId string) (*WSGDPSKDeleteExpiredDpskConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPSKDeleteExpiredDpskConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDeleteExpiredDpskByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKDeleteExpiredDpskConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDownloadDpskCsvSample
//
// Use this API command to download DPSK CSV sample.
//
// Optional Parameters:
// - type_ string
//		- nullable
func (s *WSGDynamicpskService) FindRkszonesDownloadDpskCsvSample(ctx context.Context, optionalParams map[string][]string) ([]byte, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDownloadDpskCsvSample, true)
	if v, ok := optionalParams["type"]; ok && len(v) > 0 {
		req.SetQueryParameter("type", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindRkszonesDpskByZoneId
//
// Use this API command to retrieve DPSK info of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicpskService) FindRkszonesDpskByZoneId(ctx context.Context, zoneId string) (*WSGDPSKGetDpskInfoList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPSKGetDpskInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDpskByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskInfoList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDpskEnabledWlansByZoneId
//
// Use this API command to retrieve DPSK enabled WLAN info of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicpskService) FindRkszonesDpskEnabledWlansByZoneId(ctx context.Context, zoneId string) (*WSGDPSKGetDpskEnabledWlans, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPSKGetDpskEnabledWlans
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDpskEnabledWlansByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskEnabledWlans()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesWlansDpskByDpskId
//
// Use this API command to retrieve DPSK info.
//
// Required Parameters:
// - dpskId string
//		- required
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicpskService) FindRkszonesWlansDpskByDpskId(ctx context.Context, dpskId string, id string, zoneId string) (*WSGDPSKGetDpskInfoList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPSKGetDpskInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesWlansDpskByDpskId, true)
	req.SetPathParameter("dpskId", dpskId)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskInfoList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesWlansDpskById
//
// Use this API command to retrieve DPSK info of a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicpskService) FindRkszonesWlansDpskById(ctx context.Context, id string, zoneId string) (*WSGDPSKGetDpskInfoList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPSKGetDpskInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesWlansDpskById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskInfoList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesWlansDpskByDpskId
//
// Use this API command to update DPSK info.
//
// Request Body:
//	 - body *WSGDPSKUpdateDpsk
//
// Required Parameters:
// - dpskId string
//		- required
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicpskService) PartialUpdateRkszonesWlansDpskByDpskId(ctx context.Context, body *WSGDPSKUpdateDpsk, dpskId string, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesWlansDpskByDpskId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("dpskId", dpskId)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateRkszonesDeleteExpiredDpskByZoneId
//
// Use this API command to modify interval of delete expired DPSK of a zone.
//
// Request Body:
//	 - body *WSGDPSKModifyDeleteExpiredDpsk
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicpskService) UpdateRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, body *WSGDPSKModifyDeleteExpiredDpsk, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesDeleteExpiredDpskByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
