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
// Operation ID: healthExtendGroupBarByType
//
// Retrieve connection failure legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGHealthExtendGroupBarByType, true)
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
// Operation ID: healthExtendLineByType
//
// Retrieve graph values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGHealthExtendLineByType, true)
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
// Operation ID: perfGroupBarByType
//
// Retrieve graph type legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfGroupBarByType, true)
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
// Operation ID: perfLineByType
//
// Retrieve graph values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfLineByType, true)
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
