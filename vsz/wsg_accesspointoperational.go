package vsz

// API Version: v9_0

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

// FindApByQueryCriteria
//
// Query APs with specified filters
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindApByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAPQueryList, error) {
	var (
		req      *APIRequest
		resp     *WSGAPQueryList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindApByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPQueryList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
// s API provide detailed AP status and configuration, therefore it was designed for single AP information retrieving. If you need to retrieve large number of ap states, please use "POST://query/ap" (refer to the "Query APs" section of the category "Access Point Operational").
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

// FindApWlanByQueryCriteria
//
// Use this API command to retrieve AP Wlan list with BSSID information by QueryCriteria
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindApWlanByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGWLANQueryApWlanBssidQueryList, error) {
	var (
		req      *APIRequest
		resp     *WSGWLANQueryApWlanBssidQueryList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindApWlanByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWLANQueryApWlanBssidQueryList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindIndoorMapByQueryCriteria
//
// Query indoorMap with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindIndoorMapByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGIndoorMapSummaryList, error) {
	var (
		req      *APIRequest
		resp     *WSGIndoorMapSummaryList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindIndoorMapByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIndoorMapSummaryList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindMeshNeighborByApMacByQueryCriteria
//
// Use this API command to retrieve a list of neighbor access points on mesh AP.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindMeshNeighborByApMacByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, apMac string) (*WSGMeshNeighborInfoList, error) {
	var (
		req      *APIRequest
		resp     *WSGMeshNeighborInfoList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindMeshNeighborByApMacByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGMeshNeighborInfoList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindMeshTopologyByApMacByQueryCriteria
//
// Use this API command to retrieve a list of topology on mesh AP.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindMeshTopologyByApMacByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, apMac string) (WSGMeshNodeInfoArray, error) {
	var (
		req      *APIRequest
		resp     WSGMeshNodeInfoArray
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindMeshTopologyByApMacByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeWSGMeshNodeInfoArray()
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// FindMeshTopologyByQueryCriteria
//
// Use this API command to retrieve a list of topology on zone.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindMeshTopologyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGMeshNodeInfoList, error) {
	var (
		req      *APIRequest
		resp     *WSGMeshNodeInfoList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindMeshTopologyByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGMeshNodeInfoList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRoguesInfoListByQueryCriteria
//
// Use this API command to retrieve a list of rogue access points.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindRoguesInfoListByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGRogueInfoList, error) {
	var (
		req      *APIRequest
		resp     *WSGRogueInfoList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRoguesInfoListByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGRogueInfoList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}
