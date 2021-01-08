package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
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

func MakeSCIFacetGetApmacFacet200ResponseType() SCIFacetGetApmacFacet200ResponseType {
	m := make(SCIFacetGetApmacFacet200ResponseType, 0)
	return m
}

// SCIFacetGetFacet200ResponseType
//
// Definition: facet_getFacet200ResponseType
type SCIFacetGetFacet200ResponseType []interface{}

func MakeSCIFacetGetFacet200ResponseType() SCIFacetGetFacet200ResponseType {
	m := make(SCIFacetGetFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSsidFacet200ResponseType
//
// Definition: facet_getSsidFacet200ResponseType
type SCIFacetGetSsidFacet200ResponseType []string

func MakeSCIFacetGetSsidFacet200ResponseType() SCIFacetGetSsidFacet200ResponseType {
	m := make(SCIFacetGetSsidFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSwitchesFacet200ResponseType
//
// Definition: facet_getSwitchesFacet200ResponseType
type SCIFacetGetSwitchesFacet200ResponseType []string

func MakeSCIFacetGetSwitchesFacet200ResponseType() SCIFacetGetSwitchesFacet200ResponseType {
	m := make(SCIFacetGetSwitchesFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSwitchHierarchyFacet200ResponseType
//
// Definition: facet_getSwitchHierarchyFacet200ResponseType
type SCIFacetGetSwitchHierarchyFacet200ResponseType []interface{}

func MakeSCIFacetGetSwitchHierarchyFacet200ResponseType() SCIFacetGetSwitchHierarchyFacet200ResponseType {
	m := make(SCIFacetGetSwitchHierarchyFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSystemFacet200ResponseType
//
// Definition: facet_getSystemFacet200ResponseType
type SCIFacetGetSystemFacet200ResponseType []interface{}

func MakeSCIFacetGetSystemFacet200ResponseType() SCIFacetGetSystemFacet200ResponseType {
	m := make(SCIFacetGetSystemFacet200ResponseType, 0)
	return m
}

// FacetGetApmacFacet
//
// Operation ID: facet_getApmacFacet
//
// Retrieve list of apmac facets based on query
//
// Request Body:
//	 - body *SCICommonQueryBody
func (s *SCIFacetService) FacetGetApmacFacet(ctx context.Context, body *SCICommonQueryBody, mutators ...RequestMutator) (*SCIFacetGetApmacFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetApmacFacet200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIFacetGetApmacFacet200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIFacetGetApmacFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIFacetGetApmacFacet200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetApmacFacet200ResponseTypeAPIResponse), err
}

// FacetGetFacet
//
// Operation ID: facet_getFacet
//
// For the <b><code>filter</code></b> field below, an example would be
// <pre>
//   <code class="json">
//     { "type": "and", "fields": [{ "type": "selector", "dimension": "system", "value": "sys1" }]}
//   </code>
// </pre>
//
// Request Body:
//	 - body *SCICommonQueryBody
//
// Required Parameters:
// - name string
//		- required
//		- oneof:[system,switchHierarchy,apmac,ssid,switches]
func (s *SCIFacetService) FacetGetFacet(ctx context.Context, body *SCICommonQueryBody, name string, mutators ...RequestMutator) (*SCIFacetGetFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetFacet200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIFacetGetFacet200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIFacetGetFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIFacetGetFacet200ResponseTypeAPIResponse), err
	}
	req.PathParams.Set("name", name)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetFacet200ResponseTypeAPIResponse), err
}

// FacetGetSsidFacet
//
// Operation ID: facet_getSsidFacet
//
// Retrieve list of ssid facets based on query
//
// Request Body:
//	 - body *SCICommonQueryBody
func (s *SCIFacetService) FacetGetSsidFacet(ctx context.Context, body *SCICommonQueryBody, mutators ...RequestMutator) (*SCIFacetGetSsidFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetSsidFacet200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIFacetGetSsidFacet200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIFacetGetSsidFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIFacetGetSsidFacet200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetSsidFacet200ResponseTypeAPIResponse), err
}

// FacetGetSwitchesFacet
//
// Operation ID: facet_getSwitchesFacet
//
// Retrieve list of switches facets based on query
//
// Request Body:
//	 - body *SCICommonQueryBody
func (s *SCIFacetService) FacetGetSwitchesFacet(ctx context.Context, body *SCICommonQueryBody, mutators ...RequestMutator) (*SCIFacetGetSwitchesFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetSwitchesFacet200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIFacetGetSwitchesFacet200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIFacetGetSwitchesFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIFacetGetSwitchesFacet200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetSwitchesFacet200ResponseTypeAPIResponse), err
}

// FacetGetSwitchHierarchyFacet
//
// Operation ID: facet_getSwitchHierarchyFacet
//
// Retrieve list of switchHierarchy facets based on query
//
// Request Body:
//	 - body *SCICommonQueryBody
func (s *SCIFacetService) FacetGetSwitchHierarchyFacet(ctx context.Context, body *SCICommonQueryBody, mutators ...RequestMutator) (*SCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIFacetGetSwitchHierarchyFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetSwitchHierarchyFacet200ResponseTypeAPIResponse), err
}

// FacetGetSystemFacet
//
// Operation ID: facet_getSystemFacet
//
// Retrieve list of system facets based on query
//
// Request Body:
//	 - body *SCICommonQueryBody
func (s *SCIFacetService) FacetGetSystemFacet(ctx context.Context, body *SCICommonQueryBody, mutators ...RequestMutator) (*SCIFacetGetSystemFacet200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIFacetGetSystemFacet200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIFacetGetSystemFacet200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIFacetGetSystemFacet, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIFacetGetSystemFacet200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIFacetGetSystemFacet200ResponseTypeAPIResponse), err
}
