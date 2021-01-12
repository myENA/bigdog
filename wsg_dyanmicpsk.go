package bigdog

// API Version: v9_1

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
// Use this API command to batch generate DPSKs of a WLAN. You can either specify passphrases or not. If the amount is bigger than 1, system will generate usernames with index. e.g. student-1, student-2, ...etc.
//
// Operation ID: addRkszonesWlansDpskBatchGenUnboundById
// Operation path: /rkszones/{zoneId}/wlans/{id}/dpsk/batchGenUnbound
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGDPSKBatchGenUnbound
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskBatchGenUnboundById(ctx context.Context, body *WSGDPSKBatchGenUnbound, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesWlansDpskBatchGenUnboundById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKGetDpskResultAPIResponse), err
}

// AddRkszonesWlansDpskById
//
// Use this API command to delete DPSKs of a WLAN.
//
// Operation ID: addRkszonesWlansDpskById
// Operation path: /rkszones/{zoneId}/wlans/{id}/dpsk
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPSKDeleteDPSKs
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskById(ctx context.Context, body *WSGDPSKDeleteDPSKs, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKDeleteDpskResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKDeleteDpskResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesWlansDpskById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKDeleteDpskResultAPIResponse), err
}

// AddRkszonesWlansDpskUploadById
//
// Use this API command to upload DPSK file of a WLAN (CSV file and Content-Type multipart/form-data ONLY).
//
// Operation ID: addRkszonesWlansDpskUploadById
// Operation path: /rkszones/{zoneId}/wlans/{id}/dpsk/upload
// Success code: 201 (Created)
//
// Form data parameters:
// - uploadFile io.Reader
//		- required
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskUploadById(ctx context.Context, filename string, uploadFile io.Reader, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesWlansDpskUploadById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*WSGDPSKGetDpskResultAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKGetDpskResultAPIResponse), err
}

// FindDpskByQueryCriteria
//
// Query DPSKs with specified filters.
//
// Operation ID: findDpskByQueryCriteria
// Operation path: /query/dpsk
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDynamicPSKService) FindDpskByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGDPSKQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKQueryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindDpskByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKQueryListAPIResponse), err
}

// FindRkszonesDeleteExpiredDpskByZoneId
//
// Use this API command to retrieve interval of delete expired DPSK of a zone.
//
// Operation ID: findRkszonesDeleteExpiredDpskByZoneId
// Operation path: /rkszones/{zoneId}/deleteExpiredDpsk
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDPSKDeleteExpiredDpskConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKDeleteExpiredDpskConfigAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesDeleteExpiredDpskByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKDeleteExpiredDpskConfigAPIResponse), err
}

// FindRkszonesDownloadDpskCsvSample
//
// Use this API command to download DPSK CSV sample.
//
// Operation ID: findRkszonesDownloadDpskCsvSample
// Operation path: /rkszones/downloadDpskCsvSample
// Success code: 200 (OK)
//
// Optional parameters:
// - type_ string
//		- nullable
func (s *WSGDynamicPSKService) FindRkszonesDownloadDpskCsvSample(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesDownloadDpskCsvSample, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["type"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("type", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*FileAPIResponse), err
}

// FindRkszonesDpskByZoneId
//
// Use this API command to retrieve DPSK info of a zone.
//
// Operation ID: findRkszonesDpskByZoneId
// Operation path: /rkszones/{zoneId}/dpsk
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDpskByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskInfoListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesDpskByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKGetDpskInfoListAPIResponse), err
}

// FindRkszonesDpskEnabledWlansByZoneId
//
// Use this API command to retrieve DPSK enabled WLAN info of a zone.
//
// Operation ID: findRkszonesDpskEnabledWlansByZoneId
// Operation path: /rkszones/{zoneId}/dpskEnabledWlans
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDpskEnabledWlansByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskEnabledWlansAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskEnabledWlansAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesDpskEnabledWlansByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKGetDpskEnabledWlansAPIResponse), err
}

// FindRkszonesWlansDpskByDpskId
//
// Use this API command to retrieve DPSK info.
//
// Operation ID: findRkszonesWlansDpskByDpskId
// Operation path: /rkszones/{zoneId}/wlans/{id}/dpsk/{dpskId}
// Success code: 200 (OK)
//
// Required parameters:
// - dpskId string
//		- required
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskByDpskId(ctx context.Context, dpskId string, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskInfoListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesWlansDpskByDpskId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("dpskId", dpskId)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKGetDpskInfoListAPIResponse), err
}

// FindRkszonesWlansDpskById
//
// Use this API command to retrieve DPSK info of a WLAN.
//
// Operation ID: findRkszonesWlansDpskById
// Operation path: /rkszones/{zoneId}/wlans/{id}/dpsk
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskInfoListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesWlansDpskById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKGetDpskInfoListAPIResponse), err
}

// PartialUpdateRkszonesWlansDpskByDpskId
//
// Use this API command to update DPSK info.
//
// Operation ID: partialUpdateRkszonesWlansDpskByDpskId
// Operation path: /rkszones/{zoneId}/wlans/{id}/dpsk/{dpskId}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGDPSKUpdateDpsk
//
// Required parameters:
// - dpskId string
//		- required
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) PartialUpdateRkszonesWlansDpskByDpskId(ctx context.Context, body *WSGDPSKUpdateDpsk, dpskId string, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesWlansDpskByDpskId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("dpskId", dpskId)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesDeleteExpiredDpskByZoneId
//
// Use this API command to modify interval of delete expired DPSK of a zone.
//
// Operation ID: updateRkszonesDeleteExpiredDpskByZoneId
// Operation path: /rkszones/{zoneId}/deleteExpiredDpsk
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGDPSKModifyDeleteExpiredDpsk
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) UpdateRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, body *WSGDPSKModifyDeleteExpiredDpsk, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesDeleteExpiredDpskByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
