package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGRogueClientService struct {
	apiClient *APIClient
}

func NewWSGRogueClientService(c *APIClient) *WSGRogueClientService {
	s := new(WSGRogueClientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRogueClientService() *WSGRogueClientService {
	return NewWSGRogueClientService(ss.apiClient)
}

type WSGRogueClientRogueInfo struct {
	// Channel
	// Channel of the rogue client
	// Constraints:
	//    - nullable
	Channel *int `json:"channel,omitempty"`

	// Classification
	// Rogue classification by policy
	// Constraints:
	//    - nullable
	Classification *string `json:"classification,omitempty"`

	// DetectedByAP
	// The list of APs that found the rogue client
	// Constraints:
	//    - nullable
	DetectedByAP []*WSGAPInfo `json:"detectedByAP,omitempty" validate:"omitempty,dive"`

	// Encryption
	// Encryption of the rogue client
	// Constraints:
	//    - nullable
	Encryption *string `json:"encryption,omitempty"`

	// LastDetected
	// Timestamp of the rogue client
	// Constraints:
	//    - nullable
	LastDetected *int `json:"lastDetected,omitempty"`

	// MatchResult
	// What policy and rule matched when system doing classification by rogue policy
	// Constraints:
	//    - nullable
	MatchResult *string `json:"matchResult,omitempty"`

	// Radio
	// Radio of the rogue client
	// Constraints:
	//    - nullable
	Radio *string `json:"radio,omitempty"`

	// RogueAPMac
	// Constraints:
	//    - nullable
	RogueAPMac *WSGCommonMac `json:"rogueAPMac,omitempty"`

	// RogueMac
	// Constraints:
	//    - nullable
	RogueMac *WSGCommonMac `json:"rogueMac,omitempty"`

	// Ssid
	// SSID of the rogue client
	// Constraints:
	//    - nullable
	Ssid *string `json:"ssid,omitempty"`

	// Type
	// Type of the rogue client
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`
}

func NewWSGRogueClientRogueInfo() *WSGRogueClientRogueInfo {
	m := new(WSGRogueClientRogueInfo)
	return m
}

type WSGRogueClientRogueInfoList struct {
	// Extra
	// Any additional response data.
	// Constraints:
	//    - nullable
	Extra *WSGRogueClientRogueInfoListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Rogue AP returned out of the complete Rogue Client list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Rogue Clients after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGRogueClientRogueInfo `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Total Rogue Clients count.
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Rogue Clients count in this response.
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGRogueClientRogueInfoList() *WSGRogueClientRogueInfoList {
	m := new(WSGRogueClientRogueInfoList)
	return m
}

// WSGRogueClientRogueInfoListExtraType
//
// Any additional response data.
// Constraints:
//    - nullable
type WSGRogueClientRogueInfoListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGRogueClientRogueInfoListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGRogueClientRogueInfoListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGRogueClientRogueInfoListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGRogueClientRogueInfoListExtraType() *WSGRogueClientRogueInfoListExtraType {
	m := new(WSGRogueClientRogueInfoListExtraType)
	return m
}

// FindRogueclientsByQueryCriteria
//
// Use this API command to retrieve a list of rogue clients.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGRogueClientService) FindRogueclientsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGRogueClientRogueInfoList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGRogueClientRogueInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRogueclientsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGRogueClientRogueInfoList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
