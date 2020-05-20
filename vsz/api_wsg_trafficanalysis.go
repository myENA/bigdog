package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGTrafficAnalysisService struct {
	apiClient *APIClient
}

func NewWSGTrafficAnalysisService(c *APIClient) *WSGTrafficAnalysisService {
	s := new(WSGTrafficAnalysisService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGTrafficAnalysisService() *WSGTrafficAnalysisService {
	return NewWSGTrafficAnalysisService(ss.apiClient)
}

type WSGTrafficAnalysisResults struct {
	Extra *WSGTrafficAnalysisResultsExtraType `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGTrafficAnalysisResultsListType `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGTrafficAnalysisResults() *WSGTrafficAnalysisResults {
	m := new(WSGTrafficAnalysisResults)
	return m
}

type WSGTrafficAnalysisResultsExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGTrafficAnalysisResultsExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGTrafficAnalysisResultsExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGTrafficAnalysisResultsExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGTrafficAnalysisResultsExtraType() *WSGTrafficAnalysisResultsExtraType {
	m := new(WSGTrafficAnalysisResultsExtraType)
	return m
}

type WSGTrafficAnalysisResultsListType struct {
	AppName *string `json:"appName,omitempty"`

	ExtraValues24 *WSGTrafficAnalysisResultsListTypeExtraValues24Type `json:"extraValues24,omitempty"`

	ExtraValues50 *WSGTrafficAnalysisResultsListTypeExtraValues50Type `json:"extraValues50,omitempty"`

	ExtraValuesTotal *WSGTrafficAnalysisResultsListTypeExtraValuesTotalType `json:"extraValuesTotal,omitempty"`

	Key *string `json:"key,omitempty"`

	Total *float64 `json:"total,omitempty"`

	Value *float64 `json:"value,omitempty"`

	Value24 *float64 `json:"value24,omitempty"`

	Value50 *float64 `json:"value50,omitempty"`
}

func NewWSGTrafficAnalysisResultsListType() *WSGTrafficAnalysisResultsListType {
	m := new(WSGTrafficAnalysisResultsListType)
	return m
}

type WSGTrafficAnalysisResultsListTypeExtraValues24Type struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGTrafficAnalysisResultsListTypeExtraValues24Type) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGTrafficAnalysisResultsListTypeExtraValues24Type{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGTrafficAnalysisResultsListTypeExtraValues24Type) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGTrafficAnalysisResultsListTypeExtraValues24Type() *WSGTrafficAnalysisResultsListTypeExtraValues24Type {
	m := new(WSGTrafficAnalysisResultsListTypeExtraValues24Type)
	return m
}

type WSGTrafficAnalysisResultsListTypeExtraValues50Type struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGTrafficAnalysisResultsListTypeExtraValues50Type) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGTrafficAnalysisResultsListTypeExtraValues50Type{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGTrafficAnalysisResultsListTypeExtraValues50Type) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGTrafficAnalysisResultsListTypeExtraValues50Type() *WSGTrafficAnalysisResultsListTypeExtraValues50Type {
	m := new(WSGTrafficAnalysisResultsListTypeExtraValues50Type)
	return m
}

type WSGTrafficAnalysisResultsListTypeExtraValuesTotalType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGTrafficAnalysisResultsListTypeExtraValuesTotalType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGTrafficAnalysisResultsListTypeExtraValuesTotalType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGTrafficAnalysisResultsListTypeExtraValuesTotalType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGTrafficAnalysisResultsListTypeExtraValuesTotalType() *WSGTrafficAnalysisResultsListTypeExtraValuesTotalType {
	m := new(WSGTrafficAnalysisResultsListTypeExtraValuesTotalType)
	return m
}

// FindTrafficAnalysisAggregatesByQueryCriteria
//
// View traffic analysis aggregates
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[clients,usage]
// - source string
//		- required
//		- oneof:[ap,wlan]
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisAggregatesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGTrafficAnalysisResults
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindTrafficAnalysisAggregatesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindTrafficAnalysisClientResourceByQueryCriteria
//
// View client resource analytics
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[usage,os,app]
// - source string
//		- required
//		- oneof:[ap,wlan]
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisClientResourceByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGTrafficAnalysisResults
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindTrafficAnalysisClientResourceByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindTrafficAnalysisLineRatesByQueryCriteria
//
// View line rate aggregation data
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[usage,clients,aggclients]
// - source string
//		- required
//		- oneof:[ap,wlan]
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisLineRatesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGTrafficAnalysisResults
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindTrafficAnalysisLineRatesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
