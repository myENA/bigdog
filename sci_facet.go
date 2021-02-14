package bigdog

// API Version: 1.0.0

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type SCIFacetService struct {
	apiClient *SCIClient
}

func NewSCIFacetService(c *SCIClient) *SCIFacetService {
	s := new(SCIFacetService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIFacetService() *SCIFacetService {
	return NewSCIFacetService(ss.apiClient)
}

// SCIFacetGetApmacFacet200ResponseType
//
// Definition: facet_getApmacFacet200ResponseType
type SCIFacetGetApmacFacet200ResponseType []string

type SCIFacetGetApmacFacet200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIFacetGetApmacFacet200ResponseType
}

func newSCIFacetGetApmacFacet200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIFacetGetApmacFacet200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIFacetGetApmacFacet200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIFacetGetApmacFacet200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIFacetGetApmacFacet200ResponseType() SCIFacetGetApmacFacet200ResponseType {
	m := make(SCIFacetGetApmacFacet200ResponseType, 0)
	return m
}

// SCIFacetGetFacet200ResponseType
//
// Definition: facet_getFacet200ResponseType
type SCIFacetGetFacet200ResponseType []interface{}

type SCIFacetGetFacet200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIFacetGetFacet200ResponseType
}

func newSCIFacetGetFacet200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIFacetGetFacet200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIFacetGetFacet200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIFacetGetFacet200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIFacetGetFacet200ResponseType() SCIFacetGetFacet200ResponseType {
	m := make(SCIFacetGetFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSsidFacet200ResponseType
//
// Definition: facet_getSsidFacet200ResponseType
type SCIFacetGetSsidFacet200ResponseType []string

type SCIFacetGetSsidFacet200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIFacetGetSsidFacet200ResponseType
}

func newSCIFacetGetSsidFacet200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIFacetGetSsidFacet200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIFacetGetSsidFacet200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIFacetGetSsidFacet200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIFacetGetSsidFacet200ResponseType() SCIFacetGetSsidFacet200ResponseType {
	m := make(SCIFacetGetSsidFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSwitchesFacet200ResponseType
//
// Definition: facet_getSwitchesFacet200ResponseType
type SCIFacetGetSwitchesFacet200ResponseType []string

type SCIFacetGetSwitchesFacet200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIFacetGetSwitchesFacet200ResponseType
}

func newSCIFacetGetSwitchesFacet200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIFacetGetSwitchesFacet200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIFacetGetSwitchesFacet200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIFacetGetSwitchesFacet200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIFacetGetSwitchesFacet200ResponseType() SCIFacetGetSwitchesFacet200ResponseType {
	m := make(SCIFacetGetSwitchesFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSwitchHierarchyFacet200ResponseType
//
// Definition: facet_getSwitchHierarchyFacet200ResponseType
type SCIFacetGetSwitchHierarchyFacet200ResponseType []interface{}

type SCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIFacetGetSwitchHierarchyFacet200ResponseType
}

func newSCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIFacetGetSwitchHierarchyFacet200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIFacetGetSwitchHierarchyFacet200ResponseType() SCIFacetGetSwitchHierarchyFacet200ResponseType {
	m := make(SCIFacetGetSwitchHierarchyFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSystemFacet200ResponseType
//
// Definition: facet_getSystemFacet200ResponseType
type SCIFacetGetSystemFacet200ResponseType []interface{}

type SCIFacetGetSystemFacet200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIFacetGetSystemFacet200ResponseType
}

func newSCIFacetGetSystemFacet200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIFacetGetSystemFacet200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIFacetGetSystemFacet200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIFacetGetSystemFacet200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIFacetGetSystemFacet200ResponseType() SCIFacetGetSystemFacet200ResponseType {
	m := make(SCIFacetGetSystemFacet200ResponseType, 0)
	return m
}

// FacetGetApmacFacet
//
// Retrieve list of apmac facets based on query
//
// Operation ID: facet_getApmacFacet
// Operation path: /facets/apmac
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonQueryBody
func (s *SCIFacetService) FacetGetApmacFacet(ctx context.Context, body *SCICommonQueryBody, mutators ...RequestMutator) (*SCIFacetGetApmacFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetApmacFacet200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIFacetGetApmacFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetApmacFacet200ResponseTypeAPIResponse), err
}

// FacetGetFacet
//
// For the <b><code>filter</code></b> field below, an example would be
// <pre>
//   <code class="json">
//     { "type": "and", "fields": [{ "type": "selector", "dimension": "system", "value": "sys1" }]}
//   </code>
// </pre>
//
// Operation ID: facet_getFacet
// Operation path: /facets/{name}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonQueryBody
//
// Required parameters:
// - name string
//		- required
//		- oneof:[system,switchHierarchy,apmac,ssid,switches]
func (s *SCIFacetService) FacetGetFacet(ctx context.Context, body *SCICommonQueryBody, name string, mutators ...RequestMutator) (*SCIFacetGetFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetFacet200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIFacetGetFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("name", name)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetFacet200ResponseTypeAPIResponse), err
}

// FacetGetSsidFacet
//
// Retrieve list of ssid facets based on query
//
// Operation ID: facet_getSsidFacet
// Operation path: /facets/ssid
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonQueryBody
func (s *SCIFacetService) FacetGetSsidFacet(ctx context.Context, body *SCICommonQueryBody, mutators ...RequestMutator) (*SCIFacetGetSsidFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetSsidFacet200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIFacetGetSsidFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetSsidFacet200ResponseTypeAPIResponse), err
}

// FacetGetSwitchesFacet
//
// Retrieve list of switches facets based on query
//
// Operation ID: facet_getSwitchesFacet
// Operation path: /facets/switches
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonQueryBody
func (s *SCIFacetService) FacetGetSwitchesFacet(ctx context.Context, body *SCICommonQueryBody, mutators ...RequestMutator) (*SCIFacetGetSwitchesFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetSwitchesFacet200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIFacetGetSwitchesFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetSwitchesFacet200ResponseTypeAPIResponse), err
}

// FacetGetSwitchHierarchyFacet
//
// Retrieve list of switchHierarchy facets based on query
//
// Operation ID: facet_getSwitchHierarchyFacet
// Operation path: /facets/switchHierarchy
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonQueryBody
func (s *SCIFacetService) FacetGetSwitchHierarchyFacet(ctx context.Context, body *SCICommonQueryBody, mutators ...RequestMutator) (*SCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIFacetGetSwitchHierarchyFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse), err
}

// FacetGetSystemFacet
//
// Retrieve list of system facets based on query
//
// Operation ID: facet_getSystemFacet
// Operation path: /facets/system
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonQueryBody
func (s *SCIFacetService) FacetGetSystemFacet(ctx context.Context, body *SCICommonQueryBody, mutators ...RequestMutator) (*SCIFacetGetSystemFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetSystemFacet200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIFacetGetSystemFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetSystemFacet200ResponseTypeAPIResponse), err
}
