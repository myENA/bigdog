package ruckus

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchTrafficService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchTrafficService(c *VSZClient) *SwitchMSwitchTrafficService {
	s := new(SwitchMSwitchTrafficService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchTrafficService() *SwitchMSwitchTrafficService {
	return NewSwitchMSwitchTrafficService(ss.apiClient)
}

type SwitchMSwitchTrafficTopPortErrorQueryResultList struct {
	// Extra
	// Extra information for top port error
	Extra **SwitchMSwitchTrafficTopPortErrorQueryResultList `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top port error returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top port error after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []**SwitchMSwitchTrafficTopPortErrorQueryResultList `json:"list,omitempty"`

	// RawDataTotalCount
	// Top port error count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top port error count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchTrafficTopPortErrorQueryResultList() *SwitchMSwitchTrafficTopPortErrorQueryResultList {
	m := new(SwitchMSwitchTrafficTopPortErrorQueryResultList)
	return m
}

// SwitchMSwitchTrafficTopPortErrorQueryResultListExtraType
//
// Extra information for top port error
type SwitchMSwitchTrafficTopPortErrorQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTrafficTopPortErrorQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTrafficTopPortErrorQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTrafficTopPortErrorQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTrafficTopPortErrorQueryResultListExtraType() *SwitchMSwitchTrafficTopPortErrorQueryResultListExtraType {
	m := new(SwitchMSwitchTrafficTopPortErrorQueryResultListExtraType)
	return m
}

type SwitchMSwitchTrafficTopPortTrafficUsageQueryResultList struct {
	// Extra
	// Extra information for top port traffic usage
	Extra **SwitchMSwitchTrafficTopPortTrafficUsageQueryResultList `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top port traffic usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top port traffic usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []**SwitchMSwitchTrafficTopPortTrafficUsageQueryResultList `json:"list,omitempty"`

	// RawDataTotalCount
	// Top port traffic usage count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top port traffic usage count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchTrafficTopPortTrafficUsageQueryResultList() *SwitchMSwitchTrafficTopPortTrafficUsageQueryResultList {
	m := new(SwitchMSwitchTrafficTopPortTrafficUsageQueryResultList)
	return m
}

// SwitchMSwitchTrafficTopPortTrafficUsageQueryResultListExtraType
//
// Extra information for top port traffic usage
type SwitchMSwitchTrafficTopPortTrafficUsageQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTrafficTopPortTrafficUsageQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTrafficTopPortTrafficUsageQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTrafficTopPortTrafficUsageQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTrafficTopPortTrafficUsageQueryResultListExtraType() *SwitchMSwitchTrafficTopPortTrafficUsageQueryResultListExtraType {
	m := new(SwitchMSwitchTrafficTopPortTrafficUsageQueryResultListExtraType)
	return m
}

type SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultList struct {
	// Extra
	// Extra information for top PoE utilization
	Extra **SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultList `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top PoE usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top PoE usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []**SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultList `json:"list,omitempty"`

	// RawDataTotalCount
	// PoE utilization count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total PoE utilization count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultList() *SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultList {
	m := new(SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultList)
	return m
}

// SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultListExtraType
//
// Extra information for top PoE utilization
type SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultListExtraType() *SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultListExtraType {
	m := new(SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultListExtraType)
	return m
}

type SwitchMSwitchTrafficTopTrafficUsageQueryResultList struct {
	// Extra
	// Extra information for top traffic usage
	Extra **SwitchMSwitchTrafficTopTrafficUsageQueryResultList `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top traffic usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top traffic usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []**SwitchMSwitchTrafficTopTrafficUsageQueryResultList `json:"list,omitempty"`

	// RawDataTotalCount
	// Top traffic usage count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top traffic usage count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchTrafficTopTrafficUsageQueryResultList() *SwitchMSwitchTrafficTopTrafficUsageQueryResultList {
	m := new(SwitchMSwitchTrafficTopTrafficUsageQueryResultList)
	return m
}

// SwitchMSwitchTrafficTopTrafficUsageQueryResultListExtraType
//
// Extra information for top traffic usage
type SwitchMSwitchTrafficTopTrafficUsageQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTrafficTopTrafficUsageQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTrafficTopTrafficUsageQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTrafficTopTrafficUsageQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTrafficTopTrafficUsageQueryResultListExtraType() *SwitchMSwitchTrafficTopTrafficUsageQueryResultListExtraType {
	m := new(SwitchMSwitchTrafficTopTrafficUsageQueryResultListExtraType)
	return m
}

type SwitchMSwitchTrafficTraffic struct {
	// Rx
	// RX traffic of the switch
	Rx *string `json:"rx,omitempty"`

	// Timestamp
	// Timestamp of the switch traffic
	Timestamp *string `json:"timestamp,omitempty"`

	// Total
	// Total traffic of the switch
	Total *int `json:"total,omitempty"`

	// Tx
	// TX traffic of the switch
	Tx *string `json:"tx,omitempty"`
}

func NewSwitchMSwitchTrafficTraffic() *SwitchMSwitchTrafficTraffic {
	m := new(SwitchMSwitchTrafficTraffic)
	return m
}

type SwitchMSwitchTrafficTrafficQueryResultList struct {
	// Extra
	// Extra information for traffic list
	Extra **SwitchMSwitchTrafficTrafficQueryResultList `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first traffic list returned out of the complete traffic list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more traffic list after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []**SwitchMSwitchTrafficTrafficQueryResultList `json:"list,omitempty"`

	// RawDataTotalCount
	// Total traffic count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of traffic list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchTrafficTrafficQueryResultList() *SwitchMSwitchTrafficTrafficQueryResultList {
	m := new(SwitchMSwitchTrafficTrafficQueryResultList)
	return m
}

// SwitchMSwitchTrafficTrafficQueryResultListExtraType
//
// Extra information for traffic list
type SwitchMSwitchTrafficTrafficQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTrafficTrafficQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTrafficTrafficQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTrafficTrafficQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTrafficTrafficQueryResultListExtraType() *SwitchMSwitchTrafficTrafficQueryResultListExtraType {
	m := new(SwitchMSwitchTrafficTrafficQueryResultListExtraType)
	return m
}

type SwitchMSwitchTrafficTrafficUsage struct {
	// Id
	// Identifier of the Traffic Usage
	Id *string `json:"id,omitempty"`

	// Key
	// Interface of the Traffic Usage
	Key *string `json:"key,omitempty"`

	// Value
	// Total Tx and Rx of Traffic Usage
	Value *float64 `json:"value,omitempty"`
}

func NewSwitchMSwitchTrafficTrafficUsage() *SwitchMSwitchTrafficTrafficUsage {
	m := new(SwitchMSwitchTrafficTrafficUsage)
	return m
}

// AddTrafficTopPoeutilization
//
// Use this API command retrieve the top 10 switches by the PoE utilization.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchTrafficService) AddTrafficTopPoeutilization(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopPoeutilization, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTrafficTopSwitchPoEUtilizationQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddTrafficTopPorterror
//
// Use this API command to get the top 10 switches by the porterror.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchTrafficService) AddTrafficTopPorterror(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTrafficTopPortErrorQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTrafficTopPortErrorQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopPorterror, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTrafficTopPortErrorQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddTrafficTopPortusage
//
// Use this API command to get the top 10 ports by the traffic.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchTrafficService) AddTrafficTopPortusage(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTrafficTopPortTrafficUsageQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTrafficTopPortTrafficUsageQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopPortusage, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTrafficTopPortTrafficUsageQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddTrafficTopUsage
//
// Use this API command to retrieve Top Swich/Port usage data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchTrafficService) AddTrafficTopUsage(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTrafficTopTrafficUsageQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTrafficTopTrafficUsageQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopUsage, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTrafficTopTrafficUsageQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddTrafficTotalTrend
//
// Use this API command to retrieve Swich/Port trend data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchTrafficService) AddTrafficTotalTrend(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTrafficTrafficQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTrafficTrafficQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTotalTrend, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTrafficTrafficQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
