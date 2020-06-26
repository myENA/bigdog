package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
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
// Definition: Migration.count200ResponseType
type SCIMigrationCount200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIMigrationCount200ResponseType() *SCIMigrationCount200ResponseType {
	m := new(SCIMigrationCount200ResponseType)
	return m
}

// SCIMigrationgetMigrationsidexists200ResponseType
//
// Definition: Migration.exists__get_Migrations_{id}_exists200ResponseType
type SCIMigrationgetMigrationsidexists200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationgetMigrationsidexists200ResponseType() *SCIMigrationgetMigrationsidexists200ResponseType {
	m := new(SCIMigrationgetMigrationsidexists200ResponseType)
	return m
}

// SCIMigrationheadMigrationsid200ResponseType
//
// Definition: Migration.exists__head_Migrations_{id}200ResponseType
type SCIMigrationheadMigrationsid200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationheadMigrationsid200ResponseType() *SCIMigrationheadMigrationsid200ResponseType {
	m := new(SCIMigrationheadMigrationsid200ResponseType)
	return m
}

// SCIMigrationFind200ResponseType
//
// Definition: Migration.find200ResponseType
type SCIMigrationFind200ResponseType []*SCIModelsMigration

func MakeSCIMigrationFind200ResponseType() SCIMigrationFind200ResponseType {
	m := make(SCIMigrationFind200ResponseType, 0)
	return m
}

// MigrationCount
//
// Operation ID: Migration.count
//
// Count instances of the model matched by where from the data source.
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIMigrationService) MigrationCount(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIMigrationCount200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationCount, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.SetQueryParameter("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIMigrationCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationCreate
//
// Operation ID: Migration.create
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
func (s *SCIMigrationService) MigrationCreate(ctx context.Context, data *SCIModelsMigration, mutators ...RequestMutator) (*SCIModelsMigration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationCreate, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationCreateChangeStreamGetMigrationsChangeStream
//
// Operation ID: Migration.createChangeStream__get_Migrations_change-stream
//
// Create a change stream.
//
// Optional Parameters:
// - options string
//		- nullable
func (s *SCIMigrationService) MigrationCreateChangeStreamGetMigrationsChangeStream(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*FileResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *FileResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationCreateChangeStreamGetMigrationsChangeStream, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["options"]; ok && len(v) > 0 {
		req.SetQueryParameter("options", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(FileResponse)
	rm, err = handleFileResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationCreateChangeStreamPostMigrationsChangeStream
//
// Operation ID: Migration.createChangeStream__post_Migrations_change-stream
//
// Create a change stream.
//
// Form Data Parameters:
// - options string
//		- nullable
func (s *SCIMigrationService) MigrationCreateChangeStreamPostMigrationsChangeStream(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*FileResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *FileResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationCreateChangeStreamPostMigrationsChangeStream, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(FileResponse)
	rm, err = handleFileResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationDeleteById
//
// Operation ID: Migration.deleteById
//
// Delete a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationDeleteById(ctx context.Context, id string, mutators ...RequestMutator) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCIMigrationDeleteById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationExistsGetMigrationsIdExists
//
// Operation ID: Migration.exists__get_Migrations_{id}_exists
//
// Check whether a model instance exists in the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationExistsGetMigrationsIdExists(ctx context.Context, id string, mutators ...RequestMutator) (*SCIMigrationgetMigrationsidexists200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationgetMigrationsidexists200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationExistsGetMigrationsIdExists, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIMigrationgetMigrationsidexists200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationFind
//
// Operation ID: Migration.find
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationService) MigrationFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCIMigrationFind200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationFind, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIMigrationFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// MigrationFindById
//
// Operation ID: Migration.findById
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
func (s *SCIMigrationService) MigrationFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsMigration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationFindById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationFindOne
//
// Operation ID: Migration.findOne
//
// Find first instance of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationService) MigrationFindOne(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsMigration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationFindOne, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMigrateByName
//
// Operation ID: Migration.migrateByName
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
func (s *SCIMigrationService) MigrationMigrateByName(ctx context.Context, name string, optionalParams map[string][]string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMigrateByName, true)
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetQueryParameter("name", []string{name})
	if v, ok := optionalParams["record"]; ok && len(v) > 0 {
		req.SetQueryParameter("record", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// MigrationMigrateTo
//
// Operation ID: Migration.migrateTo
//
// Run all pending migrations
//
// Optional Parameters:
// - to string
//		- nullable
func (s *SCIMigrationService) MigrationMigrateTo(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMigrateTo, true)
	req.SetHeader(headerKeyAccept, "*/*")
	if v, ok := optionalParams["to"]; ok && len(v) > 0 {
		req.SetQueryParameter("to", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// MigrationPrototypeUpdateAttributes
//
// Operation ID: Migration.prototype.updateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsMigration, id string, mutators ...RequestMutator) (*SCIModelsMigration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIMigrationPrototypeUpdateAttributes, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationRollbackTo
//
// Operation ID: Migration.rollbackTo
//
// Rollback migrations
//
// Required Parameters:
// - to string
//		- required
func (s *SCIMigrationService) MigrationRollbackTo(ctx context.Context, to string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationRollbackTo, true)
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetQueryParameter("to", []string{to})
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// MigrationUpdateAll
//
// Operation ID: Migration.updateAll
//
// Update instances of the model matched by where from the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIMigrationService) MigrationUpdateAll(ctx context.Context, data *SCIModelsMigration, optionalParams map[string][]string, mutators ...RequestMutator) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationUpdateAll, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.SetQueryParameter("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationUpsert
//
// Operation ID: Migration.upsert
//
// Update an existing model instance or insert a new one into the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
func (s *SCIMigrationService) MigrationUpsert(ctx context.Context, data *SCIModelsMigration, mutators ...RequestMutator) (*SCIModelsMigration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIMigrationUpsert, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsMigration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
