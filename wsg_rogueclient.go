package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type WSGRogueClientService struct {
	apiClient *VSZClient
}

func NewWSGRogueClientService(c *VSZClient) *WSGRogueClientService {
	s := new(WSGRogueClientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRogueClientService() *WSGRogueClientService {
	return NewWSGRogueClientService(ss.apiClient)
}

// WSGRogueClientRogueInfo
//
// Definition: rogueclient_rogueInfo
type WSGRogueClientRogueInfo struct {
	// Channel
	// Channel of the rogue client
	Channel *int `json:"channel,omitempty"`

	// Classification
	// Rogue classification by policy
	Classification *string `json:"classification,omitempty"`

	// DetectedByAP
	// The list of APs that found the rogue client
	DetectedByAP []*WSGAPInfo `json:"detectedByAP,omitempty"`

	// Encryption
	// Encryption of the rogue client
	Encryption *string `json:"encryption,omitempty"`

	// LastDetected
	// Timestamp of the rogue client
	LastDetected *int `json:"lastDetected,omitempty"`

	// MatchResult
	// What policy and rule matched when system doing classification by rogue policy
	MatchResult *string `json:"matchResult,omitempty"`

	// Radio
	// Radio of the rogue client
	Radio *string `json:"radio,omitempty"`

	RogueAPMac *WSGCommonMac `json:"rogueAPMac,omitempty"`

	RogueMac *WSGCommonMac `json:"rogueMac,omitempty"`

	// Ssid
	// SSID of the rogue client
	Ssid *string `json:"ssid,omitempty"`

	// Type
	// Type of the rogue client
	Type *string `json:"type,omitempty"`
}

func NewWSGRogueClientRogueInfo() *WSGRogueClientRogueInfo {
	m := new(WSGRogueClientRogueInfo)
	return m
}

// WSGRogueClientRogueInfoList
//
// Definition: rogueclient_rogueInfoList
type WSGRogueClientRogueInfoList struct {
	// Extra
	// Any additional response data.
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Rogue AP returned out of the complete Rogue Client list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Rogue Clients after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGRogueClientRogueInfo `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Rogue Clients count.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Rogue Clients count in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGRogueClientRogueInfoListAPIResponse struct {
	*RawAPIResponse
	Data *WSGRogueClientRogueInfoList
}

func newWSGRogueClientRogueInfoListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGRogueClientRogueInfoListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGRogueClientRogueInfoListAPIResponse) Hydrate() error {
	r.Data = new(WSGRogueClientRogueInfoList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGRogueClientRogueInfoList() *WSGRogueClientRogueInfoList {
	m := new(WSGRogueClientRogueInfoList)
	return m
}

// FindRogueclientsByQueryCriteria
//
// Use this API command to retrieve a list of rogue clients.
//
// Operation ID: findRogueclientsByQueryCriteria
// Operation path: /rogueclients/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGRogueClientService) FindRogueclientsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRogueClientRogueInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGRogueClientRogueInfoListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGRogueClientRogueInfoListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindRogueclientsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGRogueClientRogueInfoListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGRogueClientRogueInfoListAPIResponse), err
}
