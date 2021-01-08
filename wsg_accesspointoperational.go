package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
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

func newWSGAccessPointOperationalAccessPointWlansListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAccessPointOperationalAccessPointWlansListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsApPacketCaptureDownloadByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPPackCaptureApPacketCaptureResAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsApPacketCaptureStartFileCaptureByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
	}
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPPackCaptureApPacketCaptureResAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsApPacketCaptureStartStreamingByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
	}
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsApPacketCaptureStopByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsOperationalBlinkLedByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApsSwitchoverCluster, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPQueryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindApByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAPQueryListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPQueryListAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPPackCaptureApPacketCaptureResAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApsApPacketCaptureByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPPackCaptureApPacketCaptureResAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPNeighborAPListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPNeighborAPListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApsOperationalNeighborByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPNeighborAPListAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPOperationalSummaryAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPOperationalSummaryAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApsOperationalSummaryByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPOperationalSummaryAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGWLANQueryApWlanBssidQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGWLANQueryApWlanBssidQueryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindApWlanByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGWLANQueryApWlanBssidQueryListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWLANQueryApWlanBssidQueryListAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIndoorMapSummaryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIndoorMapSummaryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindIndoorMapByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGIndoorMapSummaryListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIndoorMapSummaryListAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGMeshNeighborInfoListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGMeshNeighborInfoListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindMeshNeighborByApMacByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGMeshNeighborInfoListAPIResponse), err
	}
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGMeshNeighborInfoListAPIResponse), err
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
func (s *WSGAccessPointOperationalService) FindMeshTopologyByApMacByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, apMac string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindMeshTopologyByApMacByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGMeshNodeInfoListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGMeshNodeInfoListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindMeshTopologyByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGMeshNodeInfoListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGMeshNodeInfoListAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGRogueInfoListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGRogueInfoListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindRoguesInfoListByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGRogueInfoListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGRogueInfoListAPIResponse), err
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
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAccessPointOperationalAccessPointWlansListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAccessPointOperationalAccessPointWlansListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSpecificApWlanDetails, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAccessPointOperationalAccessPointWlansListAPIResponse), err
}
