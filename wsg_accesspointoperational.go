package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type WSGAccessPointOperationalService struct {
	apiClient *VSZClient
}

func NewWSGAccessPointOperationalService(c *VSZClient) *WSGAccessPointOperationalService {
	s := new(WSGAccessPointOperationalService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccessPointOperationalService() *WSGAccessPointOperationalService {
	return NewWSGAccessPointOperationalService(ss.apiClient)
}

// WSGAccessPointOperationalAccessPointWlanEntry
//
// Definition: accesspointoperational_accessPointWlanEntry
type WSGAccessPointOperationalAccessPointWlanEntry struct {
	ApMac *string `json:"apMac,omitempty"`

	ApMacAddress *string `json:"apMacAddress,omitempty"`

	Authmethod *string `json:"authmethod,omitempty"`

	Bssid *string `json:"bssid,omitempty"`

	Key *string `json:"key,omitempty"`

	RadioId *string `json:"radioId,omitempty"`

	RadioType *string `json:"radioType,omitempty"`

	RxBytes *string `json:"rxBytes,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	TotalNumClients *string `json:"totalNumClients,omitempty"`

	TxBytes *string `json:"txBytes,omitempty"`

	WlanId *string `json:"wlanId,omitempty"`

	WlanType *int `json:"wlanType,omitempty"`

	WsgWlanId *string `json:"wsgWlanId,omitempty"`
}

func NewWSGAccessPointOperationalAccessPointWlanEntry() *WSGAccessPointOperationalAccessPointWlanEntry {
	m := new(WSGAccessPointOperationalAccessPointWlanEntry)
	return m
}

// WSGAccessPointOperationalAccessPointWlansList
//
// Definition: accesspointoperational_accessPointWlansList
type WSGAccessPointOperationalAccessPointWlansList struct {
	Data *WSGAccessPointOperationalAccessPointWlansListData `json:"data,omitempty"`

	Error interface{} `json:"error,omitempty"`

	Success *bool `json:"success,omitempty"`
}

type WSGAccessPointOperationalAccessPointWlansListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAccessPointOperationalAccessPointWlansList
}

func newWSGAccessPointOperationalAccessPointWlansListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAccessPointOperationalAccessPointWlansListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAccessPointOperationalAccessPointWlansListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAccessPointOperationalAccessPointWlansList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAccessPointOperationalAccessPointWlansList() *WSGAccessPointOperationalAccessPointWlansList {
	m := new(WSGAccessPointOperationalAccessPointWlansList)
	return m
}

// WSGAccessPointOperationalAccessPointWlansListData
//
// Definition: accesspointoperational_accessPointWlansListData
type WSGAccessPointOperationalAccessPointWlansListData struct {
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAccessPointOperationalAccessPointWlanEntry `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAccessPointOperationalAccessPointWlansListData() *WSGAccessPointOperationalAccessPointWlansListData {
	m := new(WSGAccessPointOperationalAccessPointWlansListData)
	return m
}

// AddApsApPacketCaptureDownloadByApMac
//
// Use this API to download AP packet capture file
//
// Operation ID: addApsApPacketCaptureDownloadByApMac
// Operation path: /aps/{apMac}/apPacketCapture/download
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureDownloadByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddApsApPacketCaptureDownloadByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*FileAPIResponse), err
}

// AddApsApPacketCaptureStartFileCaptureByApMac
//
// Use this API to start AP packet capture
//
// Operation ID: addApsApPacketCaptureStartFileCaptureByApMac
// Operation path: /aps/{apMac}/apPacketCapture/startFileCapture
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGAPPackCaptureApPacketCaptureReq
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartFileCaptureByApMac(ctx context.Context, body *WSGAPPackCaptureApPacketCaptureReq, apMac string, mutators ...RequestMutator) (*WSGAPPackCaptureApPacketCaptureResAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAPPackCaptureApPacketCaptureResAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddApsApPacketCaptureStartFileCaptureByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
}

// AddApsApPacketCaptureStartStreamingByApMac
//
// Use this API to start AP packet streaming
//
// Operation ID: addApsApPacketCaptureStartStreamingByApMac
// Operation path: /aps/{apMac}/apPacketCapture/startStreaming
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGAPPackCaptureApPacketCaptureReq
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartStreamingByApMac(ctx context.Context, body *WSGAPPackCaptureApPacketCaptureReq, apMac string, mutators ...RequestMutator) (*WSGAPPackCaptureApPacketCaptureResAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAPPackCaptureApPacketCaptureResAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddApsApPacketCaptureStartStreamingByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
}

// AddApsApPacketCaptureStopByApMac
//
// Use this API to stop AP packet capture or streaming
//
// Operation ID: addApsApPacketCaptureStopByApMac
// Operation path: /aps/{apMac}/apPacketCapture/stop
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStopByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddApsApPacketCaptureStopByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// AddApsOperationalBlinkLedByApMac
//
// use this API to make ap blink its led to show its position.
//
// Operation ID: addApsOperationalBlinkLedByApMac
// Operation path: /aps/{apMac}/operational/blinkLed
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsOperationalBlinkLedByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddApsOperationalBlinkLedByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// AddApsSwitchoverCluster
//
// Use this API command to switchover AP to another cluster
//
// Operation ID: addApsSwitchoverCluster
// Operation path: /aps/switchoverCluster
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAPSwitchoverAP
func (s *WSGAccessPointOperationalService) AddApsSwitchoverCluster(ctx context.Context, body *WSGAPSwitchoverAP, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddApsSwitchoverCluster, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindApByQueryCriteria
//
// Query APs with specified filters
//
// Operation ID: findApByQueryCriteria
// Operation path: /query/ap
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGAPQueryQueryCriteria
func (s *WSGAccessPointOperationalService) FindApByQueryCriteria(ctx context.Context, body *WSGAPQueryQueryCriteria, mutators ...RequestMutator) (*WSGAPQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAPQueryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindApByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAPQueryListAPIResponse), err
}

// FindApsApPacketCaptureByApMac
//
// Use this API to get AP packet capture status
//
// Operation ID: findApsApPacketCaptureByApMac
// Operation path: /aps/{apMac}/apPacketCapture
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsApPacketCaptureByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGAPPackCaptureApPacketCaptureResAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAPPackCaptureApPacketCaptureResAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApsApPacketCaptureByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
}

// FindApsOperationalNeighborByApMac
//
// Use this API command to retrieve a list of neighbor access points on mesh AP.
//
// Operation ID: findApsOperationalNeighborByApMac
// Operation path: /aps/{apMac}/operational/neighbor
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGAccessPointOperationalService) FindApsOperationalNeighborByApMac(ctx context.Context, apMac string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAPNeighborAPListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAPNeighborAPListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApsOperationalNeighborByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAPNeighborAPListAPIResponse), err
}

// FindApsOperationalSummaryByApMac
//
// s API provide detailed AP status and configuration, therefore it was designed for single AP information retrieving. If you need to retrieve large number of ap states, please use "POST://query/ap" (refer to the "Query APs" section of the category "Access Point Operational").
//
// Operation ID: findApsOperationalSummaryByApMac
// Operation path: /aps/{apMac}/operational/summary
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalSummaryByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGAPOperationalSummaryAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAPOperationalSummaryAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApsOperationalSummaryByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAPOperationalSummaryAPIResponse), err
}

// FindApWlanByQueryCriteria
//
// Use this API command to retrieve AP Wlan list with BSSID information by QueryCriteria
//
// Operation ID: findApWlanByQueryCriteria
// Operation path: /query/ap/wlan
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindApWlanByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGWLANQueryApWlanBssidQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGWLANQueryApWlanBssidQueryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindApWlanByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGWLANQueryApWlanBssidQueryListAPIResponse), err
}

// FindIndoorMapByQueryCriteria
//
// Query indoorMap with specified filters.
//
// Operation ID: findIndoorMapByQueryCriteria
// Operation path: /query/indoorMap
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindIndoorMapByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGIndoorMapSummaryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIndoorMapSummaryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindIndoorMapByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIndoorMapSummaryListAPIResponse), err
}

// FindMeshNeighborByApMacByQueryCriteria
//
// Use this API command to retrieve a list of neighbor access points on mesh AP.
//
// Operation ID: findMeshNeighborByApMacByQueryCriteria
// Operation path: /query/mesh/{apMac}/neighbor
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindMeshNeighborByApMacByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, apMac string, mutators ...RequestMutator) (*WSGMeshNeighborInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGMeshNeighborInfoListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindMeshNeighborByApMacByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGMeshNeighborInfoListAPIResponse), err
}

// FindMeshTopologyByApMacByQueryCriteria
//
// Use this API command to retrieve a list of topology on mesh AP.
//
// Operation ID: findMeshTopologyByApMacByQueryCriteria
// Operation path: /query/mesh/{apMac}/topology
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindMeshTopologyByApMacByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, apMac string, mutators ...RequestMutator) (*WSGMeshNodeInfoArrayAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGMeshNodeInfoArrayAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindMeshTopologyByApMacByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGMeshNodeInfoArrayAPIResponse), err
}

// FindMeshTopologyByQueryCriteria
//
// Use this API command to retrieve a list of topology on zone.
//
// Operation ID: findMeshTopologyByQueryCriteria
// Operation path: /query/mesh/topology
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindMeshTopologyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGMeshNodeInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGMeshNodeInfoListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindMeshTopologyByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGMeshNodeInfoListAPIResponse), err
}

// FindRoguesInfoListByQueryCriteria
//
// Use this API command to retrieve a list of rogue access points.
//
// Operation ID: findRoguesInfoListByQueryCriteria
// Operation path: /query/roguesInfoList
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindRoguesInfoListByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRogueInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGRogueInfoListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindRoguesInfoListByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGRogueInfoListAPIResponse), err
}

// FindSpecificApWlanDetails
//
// Fetch detailed information on WLANs associated with a specific AP
//
// Operation ID: findSpecificApWlanDetails
// Operation path: /aps/{apMac}/wlan
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindSpecificApWlanDetails(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGAccessPointOperationalAccessPointWlansListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAccessPointOperationalAccessPointWlansListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSpecificApWlanDetails, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAccessPointOperationalAccessPointWlansListAPIResponse), err
}
