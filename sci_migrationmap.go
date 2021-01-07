package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type SCIMigrationMapService struct {
	apiClient *SCIClient
}

func NewSCIMigrationMapService(c *SCIClient) *SCIMigrationMapService {
	s := new(SCIMigrationMapService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIMigrationMapService() *SCIMigrationMapService {
	return NewSCIMigrationMapService(ss.apiClient)
}

// SCIMigrationMapCount200ResponseType
//
// Definition: MigrationMap_count200ResponseType
type SCIMigrationMapCount200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

type SCIMigrationMapCount200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIMigrationMapCount200ResponseType
}

func newSCIMigrationMapCount200ResponseTypeAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SCIMigrationMapCount200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SCIMigrationMapCount200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIMigrationMapCount200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIMigrationMapCount200ResponseType() *SCIMigrationMapCount200ResponseType {
	m := new(SCIMigrationMapCount200ResponseType)
	return m
}

// SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType
//
// Definition: MigrationMap_exists_get_MigrationMaps_{id}_exists200ResponseType
type SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

type SCIMigrationMapExistsgetMigrationMapsidexists200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType
}

func newSCIMigrationMapExistsgetMigrationMapsidexists200ResponseTypeAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SCIMigrationMapExistsgetMigrationMapsidexists200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SCIMigrationMapExistsgetMigrationMapsidexists200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIMigrationMapExistsgetMigrationMapsidexists200ResponseType() *SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType {
	m := new(SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType)
	return m
}

// SCIMigrationMapExistsheadMigrationMapsid200ResponseType
//
// Definition: MigrationMap_exists_head_MigrationMaps_{id}200ResponseType
type SCIMigrationMapExistsheadMigrationMapsid200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationMapExistsheadMigrationMapsid200ResponseType() *SCIMigrationMapExistsheadMigrationMapsid200ResponseType {
	m := new(SCIMigrationMapExistsheadMigrationMapsid200ResponseType)
	return m
}

// SCIMigrationMapFind200ResponseType
//
// Definition: MigrationMap_find200ResponseType
type SCIMigrationMapFind200ResponseType []*SCIModelsMigrationMap

func MakeSCIMigrationMapFind200ResponseType() SCIMigrationMapFind200ResponseType {
	m := make(SCIMigrationMapFind200ResponseType, 0)
	return m
}

// MigrationMapCount
//
// Operation ID: MigrationMap_count
//
// Count instances of the model matched by where from the data source.
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIMigrationMapService) MigrationMapCount(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIMigrationMapCount200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIMigrationMapCount200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIMigrationMapCount200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationMapCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIMigrationMapCount200ResponseTypeAPIResponse), err
}

// MigrationMapCreate
//
// Operation ID: MigrationMap_create
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigrationMap
func (s *SCIMigrationMapService) MigrationMapCreate(ctx context.Context, data *SCIModelsMigrationMap, mutators ...RequestMutator) (*SCIModelsMigrationMapAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsMigrationMapAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsMigrationMapAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIMigrationMapCreate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsMigrationMapAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsMigrationMapAPIResponse), err
}

// MigrationMapCreateChangeStreamGetMigrationMapsChangeStream
//
// Operation ID: MigrationMap_createChangeStream_get_MigrationMaps_change-stream
//
// Create a change stream.
//
// Optional Parameters:
// - options string
//		- nullable
func (s *SCIMigrationMapService) MigrationMapCreateChangeStreamGetMigrationMapsChangeStream(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationMapCreateChangeStreamGetMigrationMapsChangeStream, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["options"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("options", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// MigrationMapCreateChangeStreamPostMigrationMapsChangeStream
//
// Operation ID: MigrationMap_createChangeStream_post_MigrationMaps_change-stream
//
// Create a change stream.
//
// Form Data Parameters:
// - options string
//		- nullable
func (s *SCIMigrationMapService) MigrationMapCreateChangeStreamPostMigrationMapsChangeStream(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIMigrationMapCreateChangeStreamPostMigrationMapsChangeStream, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// MigrationMapDeleteById
//
// Operation ID: MigrationMap_deleteById
//
// Delete a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationMapService) MigrationMapDeleteById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSCIMigrationMapDeleteById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// MigrationMapExistsGetMigrationMapsIdExists
//
// Operation ID: MigrationMap_exists_get_MigrationMaps_{id}_exists
//
// Check whether a model instance exists in the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationMapService) MigrationMapExistsGetMigrationMapsIdExists(ctx context.Context, id string, mutators ...RequestMutator) (*SCIMigrationMapExistsgetMigrationMapsidexists200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIMigrationMapExistsgetMigrationMapsidexists200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIMigrationMapExistsgetMigrationMapsidexists200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationMapExistsGetMigrationMapsIdExists, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIMigrationMapExistsgetMigrationMapsidexists200ResponseTypeAPIResponse), err
}

// MigrationMapFind
//
// Operation ID: MigrationMap_find
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationMapService) MigrationMapFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationMapFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// MigrationMapFindById
//
// Operation ID: MigrationMap_findById
//
// Find a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationMapService) MigrationMapFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsMigrationMapAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsMigrationMapAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsMigrationMapAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationMapFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsMigrationMapAPIResponse), err
}

// MigrationMapFindOne
//
// Operation ID: MigrationMap_findOne
//
// Find first instance of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationMapService) MigrationMapFindOne(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsMigrationMapAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsMigrationMapAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsMigrationMapAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationMapFindOne, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsMigrationMapAPIResponse), err
}

// MigrationMapPrototypeUpdateAttributes
//
// Operation ID: MigrationMap_prototype_updateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigrationMap
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationMapService) MigrationMapPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsMigrationMap, id string, mutators ...RequestMutator) (*SCIModelsMigrationMapAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsMigrationMapAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsMigrationMapAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIMigrationMapPrototypeUpdateAttributes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsMigrationMapAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsMigrationMapAPIResponse), err
}

// MigrationMapUpdateAll
//
// Operation ID: MigrationMap_updateAll
//
// Update instances of the model matched by where from the data source.
//
// Request Body:
//	 - body *SCIModelsMigrationMap
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIMigrationMapService) MigrationMapUpdateAll(ctx context.Context, data *SCIModelsMigrationMap, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIMigrationMapUpdateAll, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*RawAPIResponse), err
	}
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// MigrationMapUpsert
//
// Operation ID: MigrationMap_upsert
//
// Update an existing model instance or insert a new one into the data source.
//
// Request Body:
//	 - body *SCIModelsMigrationMap
func (s *SCIMigrationMapService) MigrationMapUpsert(ctx context.Context, data *SCIModelsMigrationMap, mutators ...RequestMutator) (*SCIModelsMigrationMapAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsMigrationMapAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsMigrationMapAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIMigrationMapUpsert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsMigrationMapAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsMigrationMapAPIResponse), err
}
