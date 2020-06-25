package bigdog

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGTrafficAnalysisService struct {
	apiClient *VSZClient
}

func NewWSGTrafficAnalysisService(c *VSZClient) *WSGTrafficAnalysisService {
	s := new(WSGTrafficAnalysisService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGTrafficAnalysisService() *WSGTrafficAnalysisService {
	return NewWSGTrafficAnalysisService(ss.apiClient)
}

// WSGTrafficAnalysisResultItem
//
// Definition: trafficanalysis_trafficAnalysisResultItem
//
// Single traffic result item
type WSGTrafficAnalysisResultItem struct {
	AppName *string `json:"appName,omitempty"`

	// ExtraValues24
	// Value seems to usually be a map of other key => values, but is probably other things, too
	ExtraValues24 interface{} `json:"extraValues24,omitempty"`

	// ExtraValues50
	// Value seems to usually be a map of other key => values, but is probably other things, too
	ExtraValues50 interface{} `json:"extraValues50,omitempty"`

	// ExtraValuesTotal
	// Value seems to usually be a map of other key => values, but is probably other things, too
	ExtraValuesTotal interface{} `json:"extraValuesTotal,omitempty"`

	Key *string `json:"key,omitempty"`

	// Total
	// Value may be string, float, or integer type
	Total interface{} `json:"total,omitempty"`

	// Value
	// Value may be string, float, or integer type
	Value interface{} `json:"value,omitempty"`

	// Value24
	// Value may be string, float, or integer type
	Value24 interface{} `json:"value24,omitempty"`

	// Value50
	// Value may be string, float, or integer type
	Value50 interface{} `json:"value50,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGTrafficAnalysisResultItem) UnmarshalJSON(b []byte) error {
	tmpt := new(WSGTrafficAnalysisResultItem)
	if err := json.Unmarshal(b, tmpt); err != nil {
		return err
	}
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	delete(tmp, "appName")
	delete(tmp, "extraValues24")
	delete(tmp, "extraValues50")
	delete(tmp, "extraValuesTotal")
	delete(tmp, "key")
	delete(tmp, "total")
	delete(tmp, "value")
	delete(tmp, "value24")
	delete(tmp, "value50")
	tmpt.XAdditionalProperties = tmp
	*t = *tmpt
	return nil
}

func (t *WSGTrafficAnalysisResultItem) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AppName != nil {
		tmp["appName"] = t.AppName
	}
	if t.ExtraValues24 != nil {
		tmp["extraValues24"] = t.ExtraValues24
	}
	if t.ExtraValues50 != nil {
		tmp["extraValues50"] = t.ExtraValues50
	}
	if t.ExtraValuesTotal != nil {
		tmp["extraValuesTotal"] = t.ExtraValuesTotal
	}
	if t.Key != nil {
		tmp["key"] = t.Key
	}
	if t.Total != nil {
		tmp["total"] = t.Total
	}
	if t.Value != nil {
		tmp["value"] = t.Value
	}
	if t.Value24 != nil {
		tmp["value24"] = t.Value24
	}
	if t.Value50 != nil {
		tmp["value50"] = t.Value50
	}
	return json.Marshal(tmp)
}

func NewWSGTrafficAnalysisResultItem() *WSGTrafficAnalysisResultItem {
	m := new(WSGTrafficAnalysisResultItem)
	return m
}

// WSGTrafficAnalysisResults
//
// Definition: trafficanalysis_trafficAnalysisResults
//
// Results from a Traffic Analysis query.  Values may be float, integer, or string in type
type WSGTrafficAnalysisResults struct {
	// FirstIndex
	// Can probably be ignored
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Can probably be ignored
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGTrafficAnalysisResultItem `json:"list,omitempty"`

	// RawDataTotalCount
	// Always seems to be zero, not sure.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Number of items in "list"
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGTrafficAnalysisResults() *WSGTrafficAnalysisResults {
	m := new(WSGTrafficAnalysisResults)
	return m
}

// FindTrafficAnalysisAggregatesByQueryCriteria
//
// Operation ID: findTrafficAnalysisAggregatesByQueryCriteria
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
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisAggregatesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindTrafficAnalysisAggregatesGroupedByQueryCriteria
//
// Operation ID: findTrafficAnalysisAggregatesGroupedByQueryCriteria
//
// View grouped traffic analysis aggregates
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
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisAggregatesGroupedByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindTrafficAnalysisAggregatesGroupedByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindTrafficAnalysisClientResourceByQueryCriteria
//
// Operation ID: findTrafficAnalysisClientResourceByQueryCriteria
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
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisClientResourceByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindTrafficAnalysisLineRatesByQueryCriteria
//
// Operation ID: findTrafficAnalysisLineRatesByQueryCriteria
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
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisLineRatesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
