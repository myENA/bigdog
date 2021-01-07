package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"encoding/json"
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

func newSCIMigrationCount200ResponseTypeAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SCIMigrationCount200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SCIMigrationCount200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIMigrationCount200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
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

func newSCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIMigrationExistsgetMigrationsidexists200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
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

func MakeSCIMigrationFind200ResponseType() SCIMigrationFind200ResponseType {
	m := make(SCIMigrationFind200ResponseType, 0)
	return m
}

// MigrationCount
//
// Operation ID: Migration_count
//
// Count instances of the model matched by where from the data source.
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIMigrationService) MigrationCount(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIMigrationCount200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIMigrationCount200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationCreate
//
// Operation ID: Migration_create
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
func (s *SCIMigrationService) MigrationCreate(ctx context.Context, data *SCIModelsMigration, mutators ...RequestMutator) (*SCIModelsMigrationAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsMigration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIMigrationCreate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigration()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationCreateChangeStreamGetMigrationsChangeStream
//
// Operation ID: Migration_createChangeStream_get_Migrations_change-stream
//
// Create a change stream.
//
// Optional Parameters:
// - options string
//		- nullable
func (s *SCIMigrationService) MigrationCreateChangeStreamGetMigrationsChangeStream(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationCreateChangeStreamGetMigrationsChangeStream, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["options"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("options", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationCreateChangeStreamPostMigrationsChangeStream
//
// Operation ID: Migration_createChangeStream_post_Migrations_change-stream
//
// Create a change stream.
//
// Form Data Parameters:
// - options string
//		- nullable
func (s *SCIMigrationService) MigrationCreateChangeStreamPostMigrationsChangeStream(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIMigrationCreateChangeStreamPostMigrationsChangeStream, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationDeleteById
//
// Operation ID: Migration_deleteById
//
// Delete a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationDeleteById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSCIMigrationDeleteById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationExistsGetMigrationsIdExists
//
// Operation ID: Migration_exists_get_Migrations_{id}_exists
//
// Check whether a model instance exists in the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationExistsGetMigrationsIdExists(ctx context.Context, id string, mutators ...RequestMutator) (*SCIMigrationExistsgetMigrationsidexists200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationExistsgetMigrationsidexists200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationExistsGetMigrationsIdExists, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIMigrationExistsgetMigrationsidexists200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationFind
//
// Operation ID: Migration_find
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationService) MigrationFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCIMigrationFind200ResponseType, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIMigrationFind200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIMigrationFind200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, &resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationFindById
//
// Operation ID: Migration_findById
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
func (s *SCIMigrationService) MigrationFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsMigrationAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsMigration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigration()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationFindOne
//
// Operation ID: Migration_findOne
//
// Find first instance of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationService) MigrationFindOne(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsMigrationAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsMigration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationFindOne, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigration()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationMigrateByName
//
// Operation ID: Migration_migrateByName
//
// Run specific migration by name
//
// Required Parameters:
// - name string
//		- required
//
// Optional Parameters:
// - record bool
//		- nullable
func (s *SCIMigrationService) MigrationMigrateByName(ctx context.Context, name string, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationMigrateByName, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("name", name)
	if v, ok := optionalParams["record"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("record", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, s.apiClient.autoHydrate, err)
	return rm, err
}

// MigrationMigrateTo
//
// Operation ID: Migration_migrateTo
//
// Run all pending migrations
//
// Optional Parameters:
// - to string
//		- nullable
func (s *SCIMigrationService) MigrationMigrateTo(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationMigrateTo, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["to"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("to", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, s.apiClient.autoHydrate, err)
	return rm, err
}

// MigrationPrototypeUpdateAttributes
//
// Operation ID: Migration_prototype_updateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsMigration, id string, mutators ...RequestMutator) (*SCIModelsMigrationAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsMigration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIMigrationPrototypeUpdateAttributes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigration()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationRollbackTo
//
// Operation ID: Migration_rollbackTo
//
// Rollback migrations
//
// Required Parameters:
// - to string
//		- required
func (s *SCIMigrationService) MigrationRollbackTo(ctx context.Context, to string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIMigrationRollbackTo, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("to", to)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, s.apiClient.autoHydrate, err)
	return rm, err
}

// MigrationUpdateAll
//
// Operation ID: Migration_updateAll
//
// Update instances of the model matched by where from the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIMigrationService) MigrationUpdateAll(ctx context.Context, data *SCIModelsMigration, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIMigrationUpdateAll, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// MigrationUpsert
//
// Operation ID: Migration_upsert
//
// Update an existing model instance or insert a new one into the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
func (s *SCIMigrationService) MigrationUpsert(ctx context.Context, data *SCIModelsMigration, mutators ...RequestMutator) (*SCIModelsMigrationAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsMigration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIMigrationUpsert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigration()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}
