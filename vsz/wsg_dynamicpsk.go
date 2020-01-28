package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
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
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskBatchGenUnboundById(ctx context.Context, body *WSGDPSKBatchGenUnbound, id string, zoneId string) (*WSGDPSKGetDpskResult, error) {
	var (
		req      *APIRequest
		resp     *WSGDPSKGetDpskResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansDpskBatchGenUnboundById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
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
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskById(ctx context.Context, body *WSGDPSKDeleteDPSKs, id string, zoneId string) (*WSGDPSKDeleteDpskResult, error) {
	var (
		req      *APIRequest
		resp     *WSGDPSKDeleteDpskResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansDpskById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKDeleteDpskResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// AddRkszonesWlansDpskUploadById
//
// Use this API command to upload DPSK file of a WLAN (CSV file and Content-Type multipart/form-data ONLY). In curl, here is an example curl -b /tmp/headers.txt -k -X POST -v -F upload@myDpsk.csv 'https://127.0.0.1:8443/wsg/api/public/v7_0/rkszones/{zoneUuid}/wlans/{wlanId}/dpsk/upload'
//
// Request Body:
//	 - body []byte
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) AddRkszonesWlansDpskUploadById(ctx context.Context, body []byte, id string, zoneId string) (*WSGDPSKGetDpskResult, error) {
	var (
		req      *APIRequest
		resp     *WSGDPSKGetDpskResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansDpskUploadById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// FindRkszonesDeleteExpiredDpskByZoneId
//
// Use this API command to retrieve interval of delete expired DPSK of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, zoneId string) (*WSGDPSKDeleteExpiredDpskConfig, error) {
	var (
		req      *APIRequest
		resp     *WSGDPSKDeleteExpiredDpskConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDeleteExpiredDpskByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKDeleteExpiredDpskConfig()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesDownloadDpskCsvSample
//
// Use this API command to download DPSK CSV sample.
//
// Optional Parameters:
// - type_ string
//		- nullable
func (s *WSGDynamicPSKService) FindRkszonesDownloadDpskCsvSample(ctx context.Context, optionalParams map[string][]string) ([]byte, error) {
	var (
		req      *APIRequest
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDownloadDpskCsvSample, true)
	if v, ok := optionalParams["type"]; ok {
		req.AddQueryParameter("type", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// FindRkszonesDpskByZoneId
//
// Use this API command to retrieve DPSK info of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDpskByZoneId(ctx context.Context, zoneId string) (*WSGDPSKGetDpskInfoList, error) {
	var (
		req      *APIRequest
		resp     *WSGDPSKGetDpskInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDpskByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskInfoList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesDpskEnabledWlansByZoneId
//
// Use this API command to retrieve DPSK enabled WLAN info of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDynamicPSKService) FindRkszonesDpskEnabledWlansByZoneId(ctx context.Context, zoneId string) (*WSGDPSKGetDpskEnabledWlans, error) {
	var (
		req      *APIRequest
		resp     *WSGDPSKGetDpskEnabledWlans
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDpskEnabledWlansByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskEnabledWlans()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskByDpskId(ctx context.Context, dpskId string, id string, zoneId string) (*WSGDPSKGetDpskInfoList, error) {
	var (
		req      *APIRequest
		resp     *WSGDPSKGetDpskInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, dpskId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesWlansDpskByDpskId, true)
	req.SetPathParameter("dpskId", dpskId)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskInfoList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGDynamicPSKService) FindRkszonesWlansDpskById(ctx context.Context, id string, zoneId string) (*WSGDPSKGetDpskInfoList, error) {
	var (
		req      *APIRequest
		resp     *WSGDPSKGetDpskInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesWlansDpskById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPSKGetDpskInfoList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGDynamicPSKService) PartialUpdateRkszonesWlansDpskByDpskId(ctx context.Context, body *WSGDPSKUpdateDpsk, dpskId string, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, dpskId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesWlansDpskByDpskId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("dpskId", dpskId)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
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
func (s *WSGDynamicPSKService) UpdateRkszonesDeleteExpiredDpskByZoneId(ctx context.Context, body *WSGDPSKModifyDeleteExpiredDpsk, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesDeleteExpiredDpskByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
