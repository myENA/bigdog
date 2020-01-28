package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGAccessPointOperationalService struct {
	apiClient *APIClient
}

func NewWSGAccessPointOperationalService(c *APIClient) *WSGAccessPointOperationalService {
	s := new(WSGAccessPointOperationalService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccessPointOperationalService() *WSGAccessPointOperationalService {
	return NewWSGAccessPointOperationalService(ss.apiClient)
}

// AddApsApPacketCaptureDownloadByApMac
//
// Use this API to download AP packet capture file
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureDownloadByApMac(ctx context.Context, apMac string) ([]byte, error) {
	var (
		req      *APIRequest
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsApPacketCaptureDownloadByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// AddApsApPacketCaptureStartFileCaptureByApMac
//
// Use this API to start AP packet capture
//
// Request Body:
//	 - body *WSGAPPackCaptureApPacketCaptureReq
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartFileCaptureByApMac(ctx context.Context, body *WSGAPPackCaptureApPacketCaptureReq, apMac string) (*WSGAPPackCaptureApPacketCaptureRes, error) {
	var (
		req      *APIRequest
		resp     *WSGAPPackCaptureApPacketCaptureRes
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
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsApPacketCaptureStartFileCaptureByApMac, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPPackCaptureApPacketCaptureRes()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// AddApsApPacketCaptureStartStreamingByApMac
//
// Use this API to start AP packet streaming
//
// Request Body:
//	 - body *WSGAPPackCaptureApPacketCaptureReq
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartStreamingByApMac(ctx context.Context, body *WSGAPPackCaptureApPacketCaptureReq, apMac string) (*WSGAPPackCaptureApPacketCaptureRes, error) {
	var (
		req      *APIRequest
		resp     *WSGAPPackCaptureApPacketCaptureRes
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
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsApPacketCaptureStartStreamingByApMac, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPPackCaptureApPacketCaptureRes()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// AddApsApPacketCaptureStopByApMac
//
// Use this API to stop AP packet capture or streaming
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStopByApMac(ctx context.Context, apMac string) (interface{}, error) {
	var (
		req      *APIRequest
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsApPacketCaptureStopByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// AddApsOperationalBlinkLedByApMac
//
// use this API to make ap blink its led to show its position.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsOperationalBlinkLedByApMac(ctx context.Context, apMac string) (interface{}, error) {
	var (
		req      *APIRequest
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsOperationalBlinkLedByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// AddApsSwitchoverCluster
//
// Use this API command to switchover AP to another cluster
//
// Request Body:
//	 - body *WSGAPSwitchoverAP
func (s *WSGAccessPointOperationalService) AddApsSwitchoverCluster(ctx context.Context, body *WSGAPSwitchoverAP) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsSwitchoverCluster, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindApsApPacketCaptureByApMac
//
// Use this API to get AP packet capture status
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsApPacketCaptureByApMac(ctx context.Context, apMac string) (*WSGAPPackCaptureApPacketCaptureRes, error) {
	var (
		req      *APIRequest
		resp     *WSGAPPackCaptureApPacketCaptureRes
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsApPacketCaptureByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPPackCaptureApPacketCaptureRes()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindApsOperationalAlarmsByApMac
//
// Use this API command to retrieve the list of alarms on an AP.
//
// Required Parameters:
// - apMac string
//		- required
//
// Optional Parameters:
// - category string
//		- nullable
// - code float64
//		- nullable
// - endTime string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
// - severity string
//		- nullable
// - startTime string
//		- nullable
// - status string
//		- nullable
func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmsByApMac(ctx context.Context, apMac string, optionalParams map[string][]string) (*WSGAPAlarmList, error) {
	var (
		req      *APIRequest
		resp     *WSGAPAlarmList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalAlarmsByApMac, true)
	req.SetPathParameter("apMac", apMac)
	if v, ok := optionalParams["category"]; ok {
		req.AddQueryParameter("category", v)
	}
	if v, ok := optionalParams["code"]; ok {
		req.AddQueryParameter("code", v)
	}
	if v, ok := optionalParams["endTime"]; ok {
		req.AddQueryParameter("endTime", v)
	}
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["severity"]; ok {
		req.AddQueryParameter("severity", v)
	}
	if v, ok := optionalParams["startTime"]; ok {
		req.AddQueryParameter("startTime", v)
	}
	if v, ok := optionalParams["status"]; ok {
		req.AddQueryParameter("status", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPAlarmList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindApsOperationalAlarmSummaryByApMac
//
// Use this API command to retrieve the alarm summary of an AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmSummaryByApMac(ctx context.Context, apMac string) (*WSGAPAlarmSummary, error) {
	var (
		req      *APIRequest
		resp     *WSGAPAlarmSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalAlarmSummaryByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPAlarmSummary()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindApsOperationalEventSummaryByApMac
//
// Use this API command to retrieve the event summary of an AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalEventSummaryByApMac(ctx context.Context, apMac string) (*WSGAPEventSummary, error) {
	var (
		req      *APIRequest
		resp     *WSGAPEventSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalEventSummaryByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPEventSummary()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindApsOperationalNeighborByApMac
//
// Use this API command to retrieve a list of neighbor access points on mesh AP.
//
// Required Parameters:
// - apMac string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGAccessPointOperationalService) FindApsOperationalNeighborByApMac(ctx context.Context, apMac string, optionalParams map[string][]string) (*WSGAPNeighborAPList, error) {
	var (
		req      *APIRequest
		resp     *WSGAPNeighborAPList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalNeighborByApMac, true)
	req.SetPathParameter("apMac", apMac)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPNeighborAPList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindApsOperationalSummaryByApMac
//
// Use this API command to retrieve the operational information of an AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalSummaryByApMac(ctx context.Context, apMac string) (*WSGAPOperationalSummary, error) {
	var (
		req      *APIRequest
		resp     *WSGAPOperationalSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalSummaryByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPOperationalSummary()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}
