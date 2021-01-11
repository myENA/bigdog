package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type SCIMigrationService struct {
	apiClient *SCIClient
}

func NewSCIMigrationService(c *SCIClient) *SCIMigrationService {
	s := new(SCIMigrationService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIMigrationService() *SCIMigrationService {
	return NewSCIMigrationService(ss.apiClient)
}

// SCIMigrationCount200ResponseType
//
// Definition: Migration_count200ResponseType
type SCIMigrationCount200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

type SCIMigrationCount200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIMigrationCount200ResponseType
}

func newSCIMigrationCount200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIMigrationCount200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIMigrationCount200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SCIMigrationCount200ResponseType)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSCIMigrationCount200ResponseType() *SCIMigrationCount200ResponseType {
	m := new(SCIMigrationCount200ResponseType)
	return m
}

// SCIMigrationExistsgetMigrationsidexists200ResponseType
//
// Definition: Migration_exists_get_Migrations_{id}_exists200ResponseType
type SCIMigrationExistsgetMigrationsidexists200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

type SCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIMigrationExistsgetMigrationsidexists200ResponseType
}

func newSCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SCIMigrationExistsgetMigrationsidexists200ResponseType)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSCIMigrationExistsgetMigrationsidexists200ResponseType() *SCIMigrationExistsgetMigrationsidexists200ResponseType {
	m := new(SCIMigrationExistsgetMigrationsidexists200ResponseType)
	return m
}

// SCIMigrationExistsheadMigrationsid200ResponseType
//
// Definition: Migration_exists_head_Migrations_{id}200ResponseType
type SCIMigrationExistsheadMigrationsid200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationExistsheadMigrationsid200ResponseType() *SCIMigrationExistsheadMigrationsid200ResponseType {
	m := new(SCIMigrationExistsheadMigrationsid200ResponseType)
	return m
}

// SCIMigrationFind200ResponseType
//
// Definition: Migration_find200ResponseType
type SCIMigrationFind200ResponseType []*SCIModelsMigration

type SCIMigrationFind200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIMigrationFind200ResponseType
}

func newSCIMigrationFind200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIMigrationFind200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIMigrationFind200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIMigrationFind200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIMigrationFind200ResponseType() SCIMigrationFind200ResponseType {
	m := make(SCIMigrationFind200ResponseType, 0)
	return m
}

// MigrationCount
//
// Count instances of the model matched by where from the data source.
//
// Operation ID: Migration_count
// Operation path: /Migrations/count
// Success code: 200 (OK)
//
// Optional parameters:
// - where string
//		- nullable
func (s *SCIMigrationService) MigrationCount(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIMigrationCount200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIMigrationCount200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIMigrationCount200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIMigrationCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIMigrationCount200ResponseTypeAPIResponse), err
}

// MigrationCreate
//
// Create a new instance of the model and persist it into the data source.
//
// Operation ID: Migration_create
// Operation path: /Migrations
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsMigration
func (s *SCIMigrationService) MigrationCreate(ctx context.Context, data *SCIModelsMigration, mutators ...RequestMutator) (*SCIModelsMigrationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsMigrationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsMigrationAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIMigrationCreate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsMigrationAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsMigrationAPIResponse), err
}

// MigrationCreateChangeStreamGetMigrationsChangeStream
//
// Create a change stream.
//
// Operation ID: Migration_createChangeStream_get_Migrations_change-stream
// Operation path: /Migrations/change-stream
// Success code: 200 (OK)
//
// Optional parameters:
// - options string
//		- nullable
func (s *SCIMigrationService) MigrationCreateChangeStreamGetMigrationsChangeStream(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*FileAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIMigrationCreateChangeStreamGetMigrationsChangeStream, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["options"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("options", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*FileAPIResponse), err
}

// MigrationCreateChangeStreamPostMigrationsChangeStream
//
// Create a change stream.
//
// Operation ID: Migration_createChangeStream_post_Migrations_change-stream
// Operation path: /Migrations/change-stream
// Success code: 200 (OK)
//
// Form data parameters:
// - options string
//		- nullable
func (s *SCIMigrationService) MigrationCreateChangeStreamPostMigrationsChangeStream(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*FileAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIMigrationCreateChangeStreamPostMigrationsChangeStream, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*FileAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*FileAPIResponse), err
}

// MigrationDeleteById
//
// Delete a model instance by id from the data source.
//
// Operation ID: Migration_deleteById
// Operation path: /Migrations/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationDeleteById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceSCI, http.MethodDelete, RouteSCIMigrationDeleteById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// MigrationExistsGetMigrationsIdExists
//
// Check whether a model instance exists in the data source.
//
// Operation ID: Migration_exists_get_Migrations_{id}_exists
// Operation path: /Migrations/{id}/exists
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationExistsGetMigrationsIdExists(ctx context.Context, id string, mutators ...RequestMutator) (*SCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIMigrationExistsGetMigrationsIdExists, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse), err
}

// MigrationFind
//
// Find all instances of the model matched by filter from the data source.
//
// Operation ID: Migration_find
// Operation path: /Migrations
// Success code: 200 (OK)
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIMigrationService) MigrationFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIMigrationFind200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIMigrationFind200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIMigrationFind200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIMigrationFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIMigrationFind200ResponseTypeAPIResponse), err
}

// MigrationFindById
//
// Find a model instance by id from the data source.
//
// Operation ID: Migration_findById
// Operation path: /Migrations/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIMigrationService) MigrationFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsMigrationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsMigrationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsMigrationAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIMigrationFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsMigrationAPIResponse), err
}

// MigrationFindOne
//
// Find first instance of the model matched by filter from the data source.
//
// Operation ID: Migration_findOne
// Operation path: /Migrations/findOne
// Success code: 200 (OK)
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIMigrationService) MigrationFindOne(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsMigrationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsMigrationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsMigrationAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIMigrationFindOne, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsMigrationAPIResponse), err
}

// MigrationMigrateByName
//
// Run specific migration by name
//
// Operation ID: Migration_migrateByName
// Operation path: /Migrations/migrateByName
// Success code: 204 (No Content)
//
// Required parameters:
// - name string
//		- required
//
// Optional parameters:
// - record bool
//		- nullable
func (s *SCIMigrationService) MigrationMigrateByName(ctx context.Context, name string, optionalParams map[string][]string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIMigrationMigrateByName, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("name", name)
	if v, ok := optionalParams["record"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("record", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// MigrationMigrateTo
//
// Run all pending migrations
//
// Operation ID: Migration_migrateTo
// Operation path: /Migrations/migrate
// Success code: 204 (No Content)
//
// Optional parameters:
// - to string
//		- nullable
func (s *SCIMigrationService) MigrationMigrateTo(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIMigrationMigrateTo, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["to"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("to", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// MigrationPrototypeUpdateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Operation ID: Migration_prototype_updateAttributes
// Operation path: /Migrations/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsMigration
//
// Required parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsMigration, id string, mutators ...RequestMutator) (*SCIModelsMigrationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsMigrationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsMigrationAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPut, RouteSCIMigrationPrototypeUpdateAttributes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsMigrationAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsMigrationAPIResponse), err
}

// MigrationRollbackTo
//
// Rollback migrations
//
// Operation ID: Migration_rollbackTo
// Operation path: /Migrations/rollback
// Success code: 204 (No Content)
//
// Required parameters:
// - to string
//		- required
func (s *SCIMigrationService) MigrationRollbackTo(ctx context.Context, to string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIMigrationRollbackTo, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("to", to)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// MigrationUpdateAll
//
// Update instances of the model matched by where from the data source.
//
// Operation ID: Migration_updateAll
// Operation path: /Migrations/update
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsMigration
//
// Optional parameters:
// - where string
//		- nullable
func (s *SCIMigrationService) MigrationUpdateAll(ctx context.Context, data *SCIModelsMigration, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIMigrationUpdateAll, true)
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

// MigrationUpsert
//
// Update an existing model instance or insert a new one into the data source.
//
// Operation ID: Migration_upsert
// Operation path: /Migrations
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsMigration
func (s *SCIMigrationService) MigrationUpsert(ctx context.Context, data *SCIModelsMigration, mutators ...RequestMutator) (*SCIModelsMigrationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsMigrationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsMigrationAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPut, RouteSCIMigrationUpsert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsMigrationAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsMigrationAPIResponse), err
}
