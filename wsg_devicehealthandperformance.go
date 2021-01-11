package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGDeviceHealthAndPerformanceService struct {
	apiClient *VSZClient
}

func NewWSGDeviceHealthAndPerformanceService(c *VSZClient) *WSGDeviceHealthAndPerformanceService {
	s := new(WSGDeviceHealthAndPerformanceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDeviceHealthAndPerformanceService() *WSGDeviceHealthAndPerformanceService {
	return NewWSGDeviceHealthAndPerformanceService(ss.apiClient)
}

// HealthExtendGroupBarByType
//
// Retrieve connection failure legend values based on query
//
// Operation ID: healthExtendGroupBarByType
// Operation path: /healthExtend/groupBar/{type}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required parameters:
// - type_ string
//		- required
//		- oneof:[totalFailure,authFailure,assocFailure,eapFailure,radiusFailure,dhcpFailure,userAuthFailure]
func (s *WSGDeviceHealthAndPerformanceService) HealthExtendGroupBarByType(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, type_ string, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPerformanceAndHealthExtensionsGroupBarListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPerformanceAndHealthExtensionsGroupBarListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGHealthExtendGroupBarByType, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGPerformanceAndHealthExtensionsGroupBarListAPIResponse), err
	}
	req.PathParams.Set("type", type_)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPerformanceAndHealthExtensionsGroupBarListAPIResponse), err
}

// HealthExtendLineByType
//
// Retrieve graph values based on query
//
// Operation ID: healthExtendLineByType
// Operation path: /healthExtend/line/{type}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required parameters:
// - type_ string
//		- required
//		- oneof:[totalFailure,authFailure,assocFailure,eapFailure,radiusFailure,dhcpFailure,userAuthFailure]
func (s *WSGDeviceHealthAndPerformanceService) HealthExtendLineByType(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, type_ string, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPerformanceAndHealthExtensionsLineListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPerformanceAndHealthExtensionsLineListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGHealthExtendLineByType, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGPerformanceAndHealthExtensionsLineListAPIResponse), err
	}
	req.PathParams.Set("type", type_)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPerformanceAndHealthExtensionsLineListAPIResponse), err
}

// PerfGroupBarByType
//
// Retrieve graph type legend values based on query
//
// Operation ID: perfGroupBarByType
// Operation path: /perf/groupBar/{type}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required parameters:
// - type_ string
//		- required
//		- oneof:[latency,airtimeUtilization,capacity]
func (s *WSGDeviceHealthAndPerformanceService) PerfGroupBarByType(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, type_ string, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPerformanceAndHealthExtensionsGroupBarListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPerformanceAndHealthExtensionsGroupBarListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGPerfGroupBarByType, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGPerformanceAndHealthExtensionsGroupBarListAPIResponse), err
	}
	req.PathParams.Set("type", type_)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPerformanceAndHealthExtensionsGroupBarListAPIResponse), err
}

// PerfLineByType
//
// Retrieve graph values based on query
//
// Operation ID: perfLineByType
// Operation path: /perf/line/{type}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required parameters:
// - type_ string
//		- required
//		- oneof:[latency,airtimeUtilization,capacity]
func (s *WSGDeviceHealthAndPerformanceService) PerfLineByType(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, type_ string, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPerformanceAndHealthExtensionsLineListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPerformanceAndHealthExtensionsLineListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGPerfLineByType, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGPerformanceAndHealthExtensionsLineListAPIResponse), err
	}
	req.PathParams.Set("type", type_)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPerformanceAndHealthExtensionsLineListAPIResponse), err
}
