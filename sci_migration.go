package ruckus

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
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

type SCIMigrationCount200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIMigrationCount200ResponseType() *SCIMigrationCount200ResponseType {
	m := new(SCIMigrationCount200ResponseType)
	return m
}

type SCIMigrationDeleteById200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIMigrationDeleteById200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIMigrationDeleteById200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIMigrationDeleteById200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIMigrationDeleteById200ResponseType() *SCIMigrationDeleteById200ResponseType {
	m := new(SCIMigrationDeleteById200ResponseType)
	return m
}

type SCIMigrationgetMigrationsidexists200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationgetMigrationsidexists200ResponseType() *SCIMigrationgetMigrationsidexists200ResponseType {
	m := new(SCIMigrationgetMigrationsidexists200ResponseType)
	return m
}

type SCIMigrationheadMigrationsid200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationheadMigrationsid200ResponseType() *SCIMigrationheadMigrationsid200ResponseType {
	m := new(SCIMigrationheadMigrationsid200ResponseType)
	return m
}

type SCIMigrationFind200ResponseType []*SCIModelsMigration

func MakeSCIMigrationFind200ResponseType() SCIMigrationFind200ResponseType {
	m := make(SCIMigrationFind200ResponseType, 0)
	return m
}

// SCIMigrationUpdateAll200ResponseType
//
// The number of instances updated
type SCIMigrationUpdateAll200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIMigrationUpdateAll200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIMigrationUpdateAll200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIMigrationUpdateAll200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIMigrationUpdateAll200ResponseType() *SCIMigrationUpdateAll200ResponseType {
	m := new(SCIMigrationUpdateAll200ResponseType)
	return m
}

// MigrationCount
//
// Count instances of the model matched by where from the data source.
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIMigrationService) MigrationCount(ctx context.Context, optionalParams map[string][]string) (*SCIMigrationCount200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationCount, false)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.SetQueryParameter("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIMigrationCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationCreate
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
func (s *SCIMigrationService) MigrationCreate(ctx context.Context, body *SCIModelsMigration) (*SCIModelsMigration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationCreate, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsMigration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationCreateChangeStreamGetMigrationsChangeStream
//
// Create a change stream.
//
// Optional Parameters:
// - options string
//		- nullable
func (s *SCIMigrationService) MigrationCreateChangeStreamGetMigrationsChangeStream(ctx context.Context, optionalParams map[string][]string) ([]byte, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationCreateChangeStreamGetMigrationsChangeStream, false)
	if v, ok := optionalParams["options"]; ok && len(v) > 0 {
		req.SetQueryParameter("options", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// MigrationCreateChangeStreamPostMigrationsChangeStream
//
// Create a change stream.
//
// Request Body:
//	 - body string
func (s *SCIMigrationService) MigrationCreateChangeStreamPostMigrationsChangeStream(ctx context.Context, body string) ([]byte, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationCreateChangeStreamPostMigrationsChangeStream, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// MigrationDeleteById
//
// Delete a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationDeleteById(ctx context.Context, id string) (*SCIMigrationDeleteById200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationDeleteById200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCIMigrationDeleteById, false)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIMigrationDeleteById200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationExistsGetMigrationsIdExists
//
// Check whether a model instance exists in the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationExistsGetMigrationsIdExists(ctx context.Context, id string) (*SCIMigrationgetMigrationsidexists200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationExistsGetMigrationsIdExists, false)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIMigrationgetMigrationsidexists200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationFind
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationService) MigrationFind(ctx context.Context, optionalParams map[string][]string) (SCIMigrationFind200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationFind, false)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIMigrationFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// MigrationFindById
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
func (s *SCIMigrationService) MigrationFindById(ctx context.Context, id string, optionalParams map[string][]string) (*SCIModelsMigration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationFindById, false)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsMigration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationFindOne
//
// Find first instance of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationService) MigrationFindOne(ctx context.Context, optionalParams map[string][]string) (*SCIModelsMigration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationFindOne, false)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsMigration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMigrateByName
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
func (s *SCIMigrationService) MigrationMigrateByName(ctx context.Context, name string, optionalParams map[string][]string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMigrateByName, false)
	req.SetQueryParameter("name", []string{name})
	if v, ok := optionalParams["record"]; ok && len(v) > 0 {
		req.SetQueryParameter("record", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// MigrationMigrateTo
//
// Run all pending migrations
//
// Optional Parameters:
// - to string
//		- nullable
func (s *SCIMigrationService) MigrationMigrateTo(ctx context.Context, optionalParams map[string][]string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMigrateTo, false)
	if v, ok := optionalParams["to"]; ok && len(v) > 0 {
		req.SetQueryParameter("to", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// MigrationPrototypeUpdateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationService) MigrationPrototypeUpdateAttributes(ctx context.Context, body *SCIModelsMigration, id string) (*SCIModelsMigration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIMigrationPrototypeUpdateAttributes, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsMigration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationRollbackTo
//
// Rollback migrations
//
// Required Parameters:
// - to string
//		- required
func (s *SCIMigrationService) MigrationRollbackTo(ctx context.Context, to string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationRollbackTo, false)
	req.SetQueryParameter("to", []string{to})
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// MigrationUpdateAll
//
// Update instances of the model matched by where from the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIMigrationService) MigrationUpdateAll(ctx context.Context, body *SCIModelsMigration, optionalParams map[string][]string) (*SCIMigrationUpdateAll200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationUpdateAll200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationUpdateAll, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.SetQueryParameter("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIMigrationUpdateAll200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationUpsert
//
// Update an existing model instance or insert a new one into the data source.
//
// Request Body:
//	 - body *SCIModelsMigration
func (s *SCIMigrationService) MigrationUpsert(ctx context.Context, body *SCIModelsMigration) (*SCIModelsMigration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIMigrationUpsert, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsMigration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
