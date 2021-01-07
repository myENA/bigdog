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
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskBatchGenUnboundById(ctx context.Context, body *WSGDPSKBatchGenUnbound, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPSKGetDpskResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansDpskBatchGenUnboundById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPSKGetDpskResultAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKGetDpskResultAPIResponse), err
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
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskById(ctx context.Context, body *WSGDPSKDeleteDPSKs, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKDeleteDpskResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKDeleteDpskResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPSKDeleteDpskResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansDpskById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPSKDeleteDpskResultAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKDeleteDpskResultAPIResponse), err
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
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskUploadById(ctx context.Context, filename string, uploadFile io.Reader, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPSKGetDpskResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansDpskUploadById, true)
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
// Operation ID: findDpskByQueryCriteria
//
// Query DPSKs with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDynamicPSKService) FindDpskByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGDPSKQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPSKQueryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindDpskByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPSKQueryListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKQueryListAPIResponse), err
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
func (s *WSGDynamicPSKService) FindRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDPSKDeleteExpiredDpskConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKDeleteExpiredDpskConfigAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPSKDeleteExpiredDpskConfigAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDeleteExpiredDpskByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKDeleteExpiredDpskConfigAPIResponse), err
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
func (s *WSGDynamicPSKService) FindRkszonesDownloadDpskCsvSample(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDownloadDpskCsvSample, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["type"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("type", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
func (s *WSGDynamicPSKService) FindRkszonesDpskByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskInfoListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPSKGetDpskInfoListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDpskByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKGetDpskInfoListAPIResponse), err
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
func (s *WSGDynamicPSKService) FindRkszonesDpskEnabledWlansByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskEnabledWlansAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskEnabledWlansAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPSKGetDpskEnabledWlansAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDpskEnabledWlansByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPSKGetDpskEnabledWlansAPIResponse), err
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
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskByDpskId(ctx context.Context, dpskId string, id string, zoneId string, mutators ...RequestMutator) (*WSGDPSKGetDpskInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPSKGetDpskInfoListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPSKGetDpskInfoListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesWlansDpskByDpskId, true)
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
// Operation ID: findRkszonesWlansDpskById
//
// Use this API command to retrieve DPSK info of a WLAN.
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPSKGetDpskInfoListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesWlansDpskById, true)
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
func (s *WSGDynamicPSKService) PartialUpdateRkszonesWlansDpskByDpskId(ctx context.Context, body *WSGDPSKUpdateDpsk, dpskId string, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesWlansDpskByDpskId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("dpskId", dpskId)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
func (s *WSGDynamicPSKService) UpdateRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, body *WSGDPSKModifyDeleteExpiredDpsk, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesDeleteExpiredDpskByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
