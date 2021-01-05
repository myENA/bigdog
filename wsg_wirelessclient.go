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
// Operation ID: addClientsBulkDeauth
//
// Use this API command to bulk deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClientList
func (s *WSGWirelessClientService) AddClientsBulkDeauth(ctx context.Context, body *WSGClientDeAuthClientList, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClientsBulkDeauth, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddClientsBulkDisconnect
//
// Operation ID: addClientsBulkDisconnect
//
// Use this API command to bulk disconnect client.
//
// Request Body:
//	 - body *WSGClientDisconnectClientList
func (s *WSGWirelessClientService) AddClientsBulkDisconnect(ctx context.Context, body *WSGClientDisconnectClientList, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClientsBulkDisconnect, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddClientsByWlanNameByWlanname
//
// Operation ID: addClientsByWlanNameByWlanname
//
// Use this API command to query client by wlan name.
//
// Request Body:
//	 - body *WSGClientQueryQueryCriteria
//
// Required Parameters:
// - wlanname string
//		- required
func (s *WSGWirelessClientService) AddClientsByWlanNameByWlanname(ctx context.Context, body *WSGClientQueryQueryCriteria, wlanname string, mutators ...RequestMutator) (*WSGClientQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClientQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClientsByWlanNameByWlanname, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("wlanname", wlanname)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGClientQueryList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddClientsDeauth
//
// Operation ID: addClientsDeauth
//
// Use this API command to deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClient
func (s *WSGWirelessClientService) AddClientsDeauth(ctx context.Context, body *WSGClientDeAuthClient, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClientsDeauth, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddClientsDisconnect
//
// Operation ID: addClientsDisconnect
//
// Use this API command to disconnect client.
//
// Request Body:
//	 - body *WSGClientDisconnectClient
func (s *WSGWirelessClientService) AddClientsDisconnect(ctx context.Context, body *WSGClientDisconnectClient, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClientsDisconnect, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindApsOperationalClientTotalCountByApMac
//
// Operation ID: findApsOperationalClientTotalCountByApMac
//
// Use this API command to retrieve the total client count per AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGWirelessClientService) FindApsOperationalClientTotalCountByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*RawAPIResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApsOperationalClientTotalCountByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindClientByQueryCriteria
//
// Operation ID: findClientByQueryCriteria
//
// Query clients with specified filters.
//
// Request Body:
//	 - body *WSGClientQueryQueryCriteria
func (s *WSGWirelessClientService) FindClientByQueryCriteria(ctx context.Context, body *WSGClientQueryQueryCriteria, mutators ...RequestMutator) (*WSGClientQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClientQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindClientByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGClientQueryList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindHistoricalclientByQueryCriteria
//
// Operation ID: findHistoricalclientByQueryCriteria
//
// Use this API command to retrive historical client.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWirelessClientService) FindHistoricalclientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGClientHistoricalClientList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClientHistoricalClientList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindHistoricalclientByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGClientHistoricalClientList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
