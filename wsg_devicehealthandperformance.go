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

// PerfHealthExtendGroupBarAirtimeUtilization
//
// Operation ID: perfHealthExtendGroupBarAirtimeUtilization
//
// Retrieve device airtime utilization legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendGroupBarAirtimeUtilization(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsGroupBarList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendGroupBarAirtimeUtilization, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsGroupBarList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendGroupBarAssocFailure
//
// Operation ID: perfHealthExtendGroupBarAssocFailure
//
// Retrieve client association failure legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendGroupBarAssocFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsGroupBarList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendGroupBarAssocFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsGroupBarList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendGroupBarAuthFailure
//
// Operation ID: perfHealthExtendGroupBarAuthFailure
//
// Retrieve client auth failure legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendGroupBarAuthFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsGroupBarList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendGroupBarAuthFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsGroupBarList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendGroupBarCapacity
//
// Operation ID: perfHealthExtendGroupBarCapacity
//
// Retrieve device capacity legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendGroupBarCapacity(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsGroupBarList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendGroupBarCapacity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsGroupBarList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendGroupBarDHCPFailure
//
// Operation ID: perfHealthExtendGroupBarDHCPFailure
//
// Retrieve client DHCP lease acquisition failure legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendGroupBarDHCPFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsGroupBarList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendGroupBarDHCPFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsGroupBarList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendGroupBarEAPFailure
//
// Operation ID: perfHealthExtendGroupBarEAPFailure
//
// Retrieve client eap authentication failure legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendGroupBarEAPFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsGroupBarList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendGroupBarEAPFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsGroupBarList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendGroupBarLatency
//
// Operation ID: perfHealthExtendGroupBarLatency
//
// Retrieve device latency legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendGroupBarLatency(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsGroupBarList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendGroupBarLatency, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsGroupBarList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendGroupBarRadiusFailure
//
// Operation ID: perfHealthExtendGroupBarRadiusFailure
//
// Retrieve client radius authentication failure legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendGroupBarRadiusFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsGroupBarList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendGroupBarRadiusFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsGroupBarList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendGroupBarTotalFailure
//
// Operation ID: perfHealthExtendGroupBarTotalFailure
//
// Retrieve connection failure legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendGroupBarTotalFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsGroupBarList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendGroupBarTotalFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsGroupBarList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendGroupBarUserAuthFailure
//
// Operation ID: perfHealthExtendGroupBarUserAuthFailure
//
// Retrieve client open auth failure legend values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendGroupBarUserAuthFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsGroupBarList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsGroupBarList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendGroupBarUserAuthFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsGroupBarList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendLineAirtimeUtilization
//
// Operation ID: perfHealthExtendLineAirtimeUtilization
//
// Retrieve percentage of airtime utilization of access points
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendLineAirtimeUtilization(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsLineList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendLineAirtimeUtilization, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsLineList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendLineAssocFailure
//
// Operation ID: perfHealthExtendLineAssocFailure
//
// Retrieve client association failures graph data
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendLineAssocFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsLineList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendLineAssocFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsLineList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendLineAuthFailure
//
// Operation ID: perfHealthExtendLineAuthFailure
//
// Retrieve client auth failure graph values
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendLineAuthFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsLineList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendLineAuthFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsLineList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendLineCapacity
//
// Operation ID: perfHealthExtendLineCapacity
//
// Retrieve theoretical capacity of devices
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendLineCapacity(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsLineList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendLineCapacity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsLineList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendLineDHCPFailure
//
// Operation ID: perfHealthExtendLineDHCPFailure
//
// Retrieve client DHCP lease acquisition failure graph data
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendLineDHCPFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsLineList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendLineDHCPFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsLineList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendLineEAPFailure
//
// Operation ID: perfHealthExtendLineEAPFailure
//
// Retrieve client eap authentication failures graph data
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendLineEAPFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsLineList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendLineEAPFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsLineList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendLineLatency
//
// Operation ID: perfHealthExtendLineLatency
//
// Retrieve device latency graph values based on query
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendLineLatency(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsLineList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendLineLatency, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsLineList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendLineRadiusFailure
//
// Operation ID: perfHealthExtendLineRadiusFailure
//
// Retrieve client radius authentication failure graph data
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendLineRadiusFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsLineList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendLineRadiusFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsLineList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendLineTotalFailure
//
// Operation ID: perfHealthExtendLineTotalFailure
//
// Retrieve connection failure line values
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendLineTotalFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsLineList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendLineTotalFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsLineList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PerfHealthExtendLineUserAuthFailure
//
// Operation ID: perfHealthExtendLineUserAuthFailure
//
// Retrieve client open auth failures graph data
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDeviceHealthAndPerformanceService) PerfHealthExtendLineUserAuthFailure(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPerformanceAndHealthExtensionsLineList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPerformanceAndHealthExtensionsLineList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGPerfHealthExtendLineUserAuthFailure, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPerformanceAndHealthExtensionsLineList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
