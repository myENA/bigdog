package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
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

func newWSGAccessPointOperationalAccessPointWlansListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAccessPointOperationalAccessPointWlansListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGAccessPointOperationalAccessPointWlansListAPIResponse) Hydrate() error {
	r.Data = new(WSGAccessPointOperationalAccessPointWlansList)
	return json.NewDecoder(r).Decode(r.Data)
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
// Operation ID: addApsApPacketCaptureDownloadByApMac
//
// Use this API to download AP packet capture file
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureDownloadByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsApPacketCaptureDownloadByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "application/octet-stream")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddApsApPacketCaptureStartFileCaptureByApMac
//
// Operation ID: addApsApPacketCaptureStartFileCaptureByApMac
//
// Use this API to start AP packet capture
//
// Request Body:
//	 - body *WSGAPPackCaptureApPacketCaptureReq
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartFileCaptureByApMac(ctx context.Context, body *WSGAPPackCaptureApPacketCaptureReq, apMac string, mutators ...RequestMutator) (*WSGAPPackCaptureApPacketCaptureResAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPPackCaptureApPacketCaptureRes
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsApPacketCaptureStartFileCaptureByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAPPackCaptureApPacketCaptureRes()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddApsApPacketCaptureStartStreamingByApMac
//
// Operation ID: addApsApPacketCaptureStartStreamingByApMac
//
// Use this API to start AP packet streaming
//
// Request Body:
//	 - body *WSGAPPackCaptureApPacketCaptureReq
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartStreamingByApMac(ctx context.Context, body *WSGAPPackCaptureApPacketCaptureReq, apMac string, mutators ...RequestMutator) (*WSGAPPackCaptureApPacketCaptureResAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPPackCaptureApPacketCaptureRes
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsApPacketCaptureStartStreamingByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAPPackCaptureApPacketCaptureRes()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddApsApPacketCaptureStopByApMac
//
// Operation ID: addApsApPacketCaptureStopByApMac
//
// Use this API to stop AP packet capture or streaming
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStopByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsApPacketCaptureStopByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddApsOperationalBlinkLedByApMac
//
// Operation ID: addApsOperationalBlinkLedByApMac
//
// use this API to make ap blink its led to show its position.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsOperationalBlinkLedByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsOperationalBlinkLedByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddApsSwitchoverCluster
//
// Operation ID: addApsSwitchoverCluster
//
// Use this API command to switchover AP to another cluster
//
// Request Body:
//	 - body *WSGAPSwitchoverAP
func (s *WSGAccessPointOperationalService) AddApsSwitchoverCluster(ctx context.Context, body *WSGAPSwitchoverAP, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsSwitchoverCluster, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindApByQueryCriteria
//
// Operation ID: findApByQueryCriteria
//
// Query APs with specified filters
//
// Request Body:
//	 - body *WSGAPQueryQueryCriteria
func (s *WSGAccessPointOperationalService) FindApByQueryCriteria(ctx context.Context, body *WSGAPQueryQueryCriteria, mutators ...RequestMutator) (*WSGAPQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindApByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAPQueryList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApsApPacketCaptureByApMac
//
// Operation ID: findApsApPacketCaptureByApMac
//
// Use this API to get AP packet capture status
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsApPacketCaptureByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGAPPackCaptureApPacketCaptureResAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPPackCaptureApPacketCaptureRes
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApsApPacketCaptureByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAPPackCaptureApPacketCaptureRes()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApsOperationalNeighborByApMac
//
// Operation ID: findApsOperationalNeighborByApMac
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
func (s *WSGAccessPointOperationalService) FindApsOperationalNeighborByApMac(ctx context.Context, apMac string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAPNeighborAPListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPNeighborAPList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApsOperationalNeighborByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("apMac", apMac)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAPNeighborAPList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApsOperationalSummaryByApMac
//
// Operation ID: findApsOperationalSummaryByApMac
//
// s API provide detailed AP status and configuration, therefore it was designed for single AP information retrieving. If you need to retrieve large number of ap states, please use "POST://query/ap" (refer to the "Query APs" section of the category "Access Point Operational").
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalSummaryByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGAPOperationalSummaryAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPOperationalSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApsOperationalSummaryByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAPOperationalSummary()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApWlanByQueryCriteria
//
// Operation ID: findApWlanByQueryCriteria
//
// Use this API command to retrieve AP Wlan list with BSSID information by QueryCriteria
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindApWlanByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGWLANQueryApWlanBssidQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWLANQueryApWlanBssidQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindApWlanByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGWLANQueryApWlanBssidQueryList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindIndoorMapByQueryCriteria
//
// Operation ID: findIndoorMapByQueryCriteria
//
// Query indoorMap with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindIndoorMapByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGIndoorMapSummaryListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIndoorMapSummaryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindIndoorMapByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGIndoorMapSummaryList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindMeshNeighborByApMacByQueryCriteria
//
// Operation ID: findMeshNeighborByApMacByQueryCriteria
//
// Use this API command to retrieve a list of neighbor access points on mesh AP.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindMeshNeighborByApMacByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, apMac string, mutators ...RequestMutator) (*WSGMeshNeighborInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGMeshNeighborInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindMeshNeighborByApMacByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGMeshNeighborInfoList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindMeshTopologyByApMacByQueryCriteria
//
// Operation ID: findMeshTopologyByApMacByQueryCriteria
//
// Use this API command to retrieve a list of topology on mesh AP.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindMeshTopologyByApMacByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, apMac string, mutators ...RequestMutator) (WSGMeshNodeInfoArray, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     WSGMeshNodeInfoArray
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindMeshTopologyByApMacByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeWSGMeshNodeInfoArray()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindMeshTopologyByQueryCriteria
//
// Operation ID: findMeshTopologyByQueryCriteria
//
// Use this API command to retrieve a list of topology on zone.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindMeshTopologyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGMeshNodeInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGMeshNodeInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindMeshTopologyByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGMeshNodeInfoList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRoguesInfoListByQueryCriteria
//
// Operation ID: findRoguesInfoListByQueryCriteria
//
// Use this API command to retrieve a list of rogue access points.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccessPointOperationalService) FindRoguesInfoListByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRogueInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGRogueInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindRoguesInfoListByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGRogueInfoList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSpecificApWlanDetails
//
// Operation ID: findSpecificApWlanDetails
//
// Fetch detailed information on WLANs associated with a specific AP
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindSpecificApWlanDetails(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGAccessPointOperationalAccessPointWlansListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAccessPointOperationalAccessPointWlansList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSpecificApWlanDetails, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAccessPointOperationalAccessPointWlansList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
