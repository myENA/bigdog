package bigdog

// API Version: v9_1

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
	type _WSGTrafficAnalysisResultItem WSGTrafficAnalysisResultItem
	tmpType := new(_WSGTrafficAnalysisResultItem)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "appName")
	delete(tmpType.XAdditionalProperties, "extraValues24")
	delete(tmpType.XAdditionalProperties, "extraValues50")
	delete(tmpType.XAdditionalProperties, "extraValuesTotal")
	delete(tmpType.XAdditionalProperties, "key")
	delete(tmpType.XAdditionalProperties, "total")
	delete(tmpType.XAdditionalProperties, "value")
	delete(tmpType.XAdditionalProperties, "value24")
	delete(tmpType.XAdditionalProperties, "value50")
	*t = WSGTrafficAnalysisResultItem(*tmpType)
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

type WSGTrafficAnalysisResultsAPIResponse struct {
	*RawAPIResponse
	Data *WSGTrafficAnalysisResults
}

func newWSGTrafficAnalysisResultsAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGTrafficAnalysisResultsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGTrafficAnalysisResultsAPIResponse) Hydrate() error {
	r.Data = new(WSGTrafficAnalysisResults)
	return json.NewDecoder(r).Decode(r.Data)
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
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisAggregatesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResultsAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindTrafficAnalysisAggregatesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("resource", resource)
	req.PathParams.Set("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
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
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisAggregatesGroupedByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResultsAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindTrafficAnalysisAggregatesGroupedByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("resource", resource)
	req.PathParams.Set("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
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
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisClientResourceByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResultsAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindTrafficAnalysisClientResourceByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("resource", resource)
	req.PathParams.Set("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
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
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisLineRatesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResultsAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindTrafficAnalysisLineRatesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("resource", resource)
	req.PathParams.Set("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}
