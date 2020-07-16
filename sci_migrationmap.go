package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
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

func NewSCIMigrationMapCount200ResponseType() *SCIMigrationMapCount200ResponseType {
	m := new(SCIMigrationMapCount200ResponseType)
	return m
}

// SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType
//
// Definition: MigrationMap_exists__get_MigrationMaps_{id}_exists200ResponseType
type SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationMapExistsgetMigrationMapsidexists200ResponseType() *SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType {
	m := new(SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType)
	return m
}

// SCIMigrationMapExistsheadMigrationMapsid200ResponseType
//
// Definition: MigrationMap_exists__head_MigrationMaps_{id}200ResponseType
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
func (s *SCIMigrationMapService) MigrationMapCount(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIMigrationMapCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationMapCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapCount, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.SetQueryParameter("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIMigrationMapCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapCreate
//
// Operation ID: MigrationMap_create
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigrationMap
func (s *SCIMigrationMapService) MigrationMapCreate(ctx context.Context, data *SCIModelsMigrationMap, mutators ...RequestMutator) (*SCIModelsMigrationMap, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsMigrationMap
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationMapCreate, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigrationMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapCreateChangeStreamGetMigrationMapsChangeStream
//
// Operation ID: MigrationMap_createChangeStream__get_MigrationMaps_change-stream
//
// Create a change stream.
//
// Optional Parameters:
// - options string
//		- nullable
func (s *SCIMigrationMapService) MigrationMapCreateChangeStreamGetMigrationMapsChangeStream(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapCreateChangeStreamGetMigrationMapsChangeStream, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["options"]; ok && len(v) > 0 {
		req.SetQueryParameter("options", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapCreateChangeStreamPostMigrationMapsChangeStream
//
// Operation ID: MigrationMap_createChangeStream__post_MigrationMaps_change-stream
//
// Create a change stream.
//
// Form Data Parameters:
// - options string
//		- nullable
func (s *SCIMigrationMapService) MigrationMapCreateChangeStreamPostMigrationMapsChangeStream(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationMapCreateChangeStreamPostMigrationMapsChangeStream, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIMigrationMapService) MigrationMapDeleteById(ctx context.Context, id string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCIMigrationMapDeleteById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapExistsGetMigrationMapsIdExists
//
// Operation ID: MigrationMap_exists__get_MigrationMaps_{id}_exists
//
// Check whether a model instance exists in the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationMapService) MigrationMapExistsGetMigrationMapsIdExists(ctx context.Context, id string, mutators ...RequestMutator) (*SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationMapExistsgetMigrationMapsidexists200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapExistsGetMigrationMapsIdExists, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIMigrationMapExistsgetMigrationMapsidexists200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIMigrationMapService) MigrationMapFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCIMigrationMapFind200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIMigrationMapFind200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapFind, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIMigrationMapFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *SCIMigrationMapService) MigrationMapFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsMigrationMap, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsMigrationMap
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapFindById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigrationMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIMigrationMapService) MigrationMapFindOne(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsMigrationMap, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsMigrationMap
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapFindOne, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigrationMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIMigrationMapService) MigrationMapPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsMigrationMap, id string, mutators ...RequestMutator) (*SCIModelsMigrationMap, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsMigrationMap
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSCIMigrationMapPrototypeUpdateAttributes, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigrationMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIMigrationMapService) MigrationMapUpdateAll(ctx context.Context, data *SCIModelsMigrationMap, optionalParams map[string][]string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationMapUpdateAll, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.SetQueryParameter("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapUpsert
//
// Operation ID: MigrationMap_upsert
//
// Update an existing model instance or insert a new one into the data source.
//
// Request Body:
//	 - body *SCIModelsMigrationMap
func (s *SCIMigrationMapService) MigrationMapUpsert(ctx context.Context, data *SCIModelsMigrationMap, mutators ...RequestMutator) (*SCIModelsMigrationMap, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsMigrationMap
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSCIMigrationMapUpsert, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigrationMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
