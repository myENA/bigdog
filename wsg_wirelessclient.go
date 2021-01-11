package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGWirelessClientService struct {
	apiClient *VSZClient
}

func NewWSGWirelessClientService(c *VSZClient) *WSGWirelessClientService {
	s := new(WSGWirelessClientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWirelessClientService() *WSGWirelessClientService {
	return NewWSGWirelessClientService(ss.apiClient)
}

// AddClientsBulkDeauth
//
// Use this API command to bulk deauth client.
//
// Operation ID: addClientsBulkDeauth
// Operation path: /clients/bulkDeauth
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGClientDeAuthClientList
func (s *WSGWirelessClientService) AddClientsBulkDeauth(ctx context.Context, body *WSGClientDeAuthClientList, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClientsBulkDeauth, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddClientsBulkDisconnect
//
// Use this API command to bulk disconnect client.
//
// Operation ID: addClientsBulkDisconnect
// Operation path: /clients/bulkDisconnect
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGClientDisconnectClientList
func (s *WSGWirelessClientService) AddClientsBulkDisconnect(ctx context.Context, body *WSGClientDisconnectClientList, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClientsBulkDisconnect, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddClientsByWlanNameByWlanname
//
// Use this API command to query client by wlan name.
//
// Operation ID: addClientsByWlanNameByWlanname
// Operation path: /clients/byWlanName/{wlanname}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGClientQueryQueryCriteria
//
// Required parameters:
// - wlanname string
//		- required
func (s *WSGWirelessClientService) AddClientsByWlanNameByWlanname(ctx context.Context, body *WSGClientQueryQueryCriteria, wlanname string, mutators ...RequestMutator) (*WSGClientQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGClientQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGClientQueryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClientsByWlanNameByWlanname, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGClientQueryListAPIResponse), err
	}
	req.PathParams.Set("wlanname", wlanname)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGClientQueryListAPIResponse), err
}

// AddClientsDeauth
//
// Use this API command to deauth client.
//
// Operation ID: addClientsDeauth
// Operation path: /clients/deauth
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGClientDeAuthClient
func (s *WSGWirelessClientService) AddClientsDeauth(ctx context.Context, body *WSGClientDeAuthClient, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClientsDeauth, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddClientsDisconnect
//
// Use this API command to disconnect client.
//
// Operation ID: addClientsDisconnect
// Operation path: /clients/disconnect
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGClientDisconnectClient
func (s *WSGWirelessClientService) AddClientsDisconnect(ctx context.Context, body *WSGClientDisconnectClient, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClientsDisconnect, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindApsOperationalClientTotalCountByApMac
//
// Use this API command to retrieve the total client count per AP.
//
// Operation ID: findApsOperationalClientTotalCountByApMac
// Operation path: /aps/{apMac}/operational/client/totalCount
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGWirelessClientService) FindApsOperationalClientTotalCountByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApsOperationalClientTotalCountByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindClientByQueryCriteria
//
// Query clients with specified filters.
//
// Operation ID: findClientByQueryCriteria
// Operation path: /query/client
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGClientQueryQueryCriteria
func (s *WSGWirelessClientService) FindClientByQueryCriteria(ctx context.Context, body *WSGClientQueryQueryCriteria, mutators ...RequestMutator) (*WSGClientQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGClientQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGClientQueryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindClientByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGClientQueryListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGClientQueryListAPIResponse), err
}

// FindHistoricalclientByQueryCriteria
//
// Use this API command to retrive historical client.
//
// Operation ID: findHistoricalclientByQueryCriteria
// Operation path: /query/historicalclient
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWirelessClientService) FindHistoricalclientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGClientHistoricalClientListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGClientHistoricalClientListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGClientHistoricalClientListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindHistoricalclientByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGClientHistoricalClientListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGClientHistoricalClientListAPIResponse), err
}
