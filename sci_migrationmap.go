package ruckus

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIMigrationmapService struct {
	apiClient *SCIClient
}

func NewSCIMigrationmapService(c *SCIClient) *SCIMigrationmapService {
	s := new(SCIMigrationmapService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIMigrationmapService() *SCIMigrationmapService {
	return NewSCIMigrationmapService(ss.apiClient)
}

type SCIMigrationmapCount200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIMigrationmapCount200ResponseType() *SCIMigrationmapCount200ResponseType {
	m := new(SCIMigrationmapCount200ResponseType)
	return m
}

type SCIMigrationmapDeleteById200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIMigrationmapDeleteById200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIMigrationmapDeleteById200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIMigrationmapDeleteById200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIMigrationmapDeleteById200ResponseType() *SCIMigrationmapDeleteById200ResponseType {
	m := new(SCIMigrationmapDeleteById200ResponseType)
	return m
}

type SCIMigrationmapFind200ResponseType []*SCIModelsMigrationMap

func MakeSCIMigrationmapFind200ResponseType() SCIMigrationmapFind200ResponseType {
	m := make(SCIMigrationmapFind200ResponseType, 0)
	return m
}

// SCIMigrationmapUpdateAll200ResponseType
//
// The number of instances updated
type SCIMigrationmapUpdateAll200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIMigrationmapUpdateAll200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIMigrationmapUpdateAll200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIMigrationmapUpdateAll200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIMigrationmapUpdateAll200ResponseType() *SCIMigrationmapUpdateAll200ResponseType {
	m := new(SCIMigrationmapUpdateAll200ResponseType)
	return m
}

// MigrationMapCount
//
// Count instances of the model matched by where from the data source.
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIMigrationmapService) MigrationMapCount(ctx context.Context, optionalParams map[string][]string) (*SCIMigrationmapCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationmapCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapCount, false)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.SetQueryParameter("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIMigrationmapCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapCreate
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigrationMap
func (s *SCIMigrationmapService) MigrationMapCreate(ctx context.Context, body *SCIModelsMigrationMap) (*SCIModelsMigrationMap, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationMapCreate, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsMigrationMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapCreateChangeStreamGetMigrationMapsChangeStream
//
// Create a change stream.
//
// Optional Parameters:
// - options string
//		- nullable
func (s *SCIMigrationmapService) MigrationMapCreateChangeStreamGetMigrationMapsChangeStream(ctx context.Context, optionalParams map[string][]string) ([]byte, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapCreateChangeStreamGetMigrationMapsChangeStream, false)
	if v, ok := optionalParams["options"]; ok && len(v) > 0 {
		req.SetQueryParameter("options", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// MigrationMapCreateChangeStreamPostMigrationMapsChangeStream
//
// Create a change stream.
//
// Request Body:
//	 - body string
func (s *SCIMigrationmapService) MigrationMapCreateChangeStreamPostMigrationMapsChangeStream(ctx context.Context, body string) ([]byte, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationMapCreateChangeStreamPostMigrationMapsChangeStream, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// MigrationMapDeleteById
//
// Delete a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationmapService) MigrationMapDeleteById(ctx context.Context, id string) (*SCIMigrationmapDeleteById200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationmapDeleteById200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCIMigrationMapDeleteById, false)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIMigrationmapDeleteById200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapExistsGetMigrationMapsIdExists
//
// Check whether a model instance exists in the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationmapService) MigrationMapExistsGetMigrationMapsIdExists(ctx context.Context, id string) (*SCIMigrationmapexistsgetMigrationMapsidexists200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationmapexistsgetMigrationMapsidexists200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapExistsGetMigrationMapsIdExists, false)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIMigrationmapexistsgetMigrationMapsidexists200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapFind
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationmapService) MigrationMapFind(ctx context.Context, optionalParams map[string][]string) (SCIMigrationmapFind200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIMigrationmapFind200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapFind, false)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIMigrationmapFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// MigrationMapFindById
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
func (s *SCIMigrationmapService) MigrationMapFindById(ctx context.Context, id string, optionalParams map[string][]string) (*SCIModelsMigrationMap, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapFindById, false)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsMigrationMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapFindOne
//
// Find first instance of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIMigrationmapService) MigrationMapFindOne(ctx context.Context, optionalParams map[string][]string) (*SCIModelsMigrationMap, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIMigrationMapFindOne, false)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsMigrationMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapPrototypeUpdateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsMigrationMap
//
// Required Parameters:
// - id string
//		- required
func (s *SCIMigrationmapService) MigrationMapPrototypeUpdateAttributes(ctx context.Context, body *SCIModelsMigrationMap, id string) (*SCIModelsMigrationMap, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIMigrationMapPrototypeUpdateAttributes, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsMigrationMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapUpdateAll
//
// Update instances of the model matched by where from the data source.
//
// Request Body:
//	 - body *SCIModelsMigrationMap
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIMigrationmapService) MigrationMapUpdateAll(ctx context.Context, body *SCIModelsMigrationMap, optionalParams map[string][]string) (*SCIMigrationmapUpdateAll200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIMigrationmapUpdateAll200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIMigrationMapUpdateAll, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.SetQueryParameter("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIMigrationmapUpdateAll200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// MigrationMapUpsert
//
// Update an existing model instance or insert a new one into the data source.
//
// Request Body:
//	 - body *SCIModelsMigrationMap
func (s *SCIMigrationmapService) MigrationMapUpsert(ctx context.Context, body *SCIModelsMigrationMap) (*SCIModelsMigrationMap, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIMigrationMapUpsert, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsMigrationMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
