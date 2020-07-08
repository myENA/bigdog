package bigdog

// API Version: v9_0

import (
	"context"
	"io"
	"net/http"
)

type WSGDynamicPSKService struct {
	apiClient *VSZClient
}

func NewWSGDynamicPSKService(c *VSZClient) *WSGDynamicPSKService {
	s := new(WSGDynamicPSKService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDynamicPSKService() *WSGDynamicPSKService {
	return NewWSGDynamicPSKService(ss.apiClient)
}

// AddRkszonesWlansDpskBatchGenUnboundById
//
// Operation ID: addRkszonesWlansDpskBatchGenUnboundById
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
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskBatchGenUnboundById(ctx context.Context, body *WSGDPSKBatchGenUnbound, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskResult, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPSKGetDpskResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansDpskById
//
// Operation ID: addRkszonesWlansDpskById
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
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskById(ctx context.Context, body *WSGDPSKDeleteDPSKs, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKDeleteDpskResult, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPSKDeleteDpskResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansDpskUploadById
//
// Operation ID: addRkszonesWlansDpskUploadById
//
// Use this API command to upload DPSK file of a WLAN (CSV file and Content-Type multipart/form-data ONLY).
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskUploadById(ctx context.Context, filename string, uploadFile io.Reader, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskResult, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueMultipartFormData)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPSKGetDpskResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// FindDpskByQueryCriteria
//
// Operation ID: findDpskByQueryCriteria
//
// Query DPSKs with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDynamicPSKService) FindDpskByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGDPSKQueryList, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPSKQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDeleteExpiredDpskByZoneId
//
// Operation ID: findRkszonesDeleteExpiredDpskByZoneId
//
// Use this API command to retrieve interval of delete expired DPSK of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDPSKDeleteExpiredDpskConfig, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPSKDeleteExpiredDpskConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDownloadDpskCsvSample
//
// Operation ID: findRkszonesDownloadDpskCsvSample
//
// Use this API command to download DPSK CSV sample.
//
// Optional Parameters:
// - type_ string
//		- nullable
func (s *WSGDynamicPSKService) FindRkszonesDownloadDpskCsvSample(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDownloadDpskCsvSample, true)
	req.SetHeader(headerKeyAccept, "application/octet-stream")
	if v, ok := optionalParams["type"]; ok && len(v) > 0 {
		req.SetQueryParameter("type", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindRkszonesDpskByZoneId
//
// Operation ID: findRkszonesDpskByZoneId
//
// Use this API command to retrieve DPSK info of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDpskByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskInfoList, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPSKGetDpskInfoList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDpskEnabledWlansByZoneId
//
// Operation ID: findRkszonesDpskEnabledWlansByZoneId
//
// Use this API command to retrieve DPSK enabled WLAN info of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDpskEnabledWlansByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskEnabledWlans, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPSKGetDpskEnabledWlans()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesWlansDpskByDpskId
//
// Operation ID: findRkszonesWlansDpskByDpskId
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
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskByDpskId(ctx context.Context, dpskId string, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskInfoList, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("dpskId", dpskId)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPSKGetDpskInfoList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesWlansDpskById
//
// Operation ID: findRkszonesWlansDpskById
//
// Use this API command to retrieve DPSK info of a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskInfoList, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPSKGetDpskInfoList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesWlansDpskByDpskId
//
// Operation ID: partialUpdateRkszonesWlansDpskByDpskId
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
func (s *WSGDynamicPSKService) PartialUpdateRkszonesWlansDpskByDpskId(ctx context.Context, body *WSGDPSKUpdateDpsk, dpskId string, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("dpskId", dpskId)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateRkszonesDeleteExpiredDpskByZoneId
//
// Operation ID: updateRkszonesDeleteExpiredDpskByZoneId
//
// Use this API command to modify interval of delete expired DPSK of a zone.
//
// Request Body:
//	 - body *WSGDPSKModifyDeleteExpiredDpsk
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) UpdateRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, body *WSGDPSKModifyDeleteExpiredDpsk, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
