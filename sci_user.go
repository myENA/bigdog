package bigdog

// API Version: 1.0.0

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type SCIUserService struct {
	apiClient *SCIClient
}

func NewSCIUserService(c *SCIClient) *SCIUserService {
	s := new(SCIUserService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIUserService() *SCIUserService {
	return NewSCIUserService(ss.apiClient)
}

// SCIUserBatchDelete200ResponseType
//
// Definition: user_batchDelete200ResponseType
type SCIUserBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

type SCIUserBatchDelete200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIUserBatchDelete200ResponseType
}

func newSCIUserBatchDelete200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIUserBatchDelete200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIUserBatchDelete200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SCIUserBatchDelete200ResponseType)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSCIUserBatchDelete200ResponseType() *SCIUserBatchDelete200ResponseType {
	m := new(SCIUserBatchDelete200ResponseType)
	return m
}

// SCIUserGetResourceGroupsForUpsert200ResponseType
//
// Definition: user_getResourceGroupsForUpsert200ResponseType
type SCIUserGetResourceGroupsForUpsert200ResponseType []interface{}

type SCIUserGetResourceGroupsForUpsert200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIUserGetResourceGroupsForUpsert200ResponseType
}

func newSCIUserGetResourceGroupsForUpsert200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIUserGetResourceGroupsForUpsert200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIUserGetResourceGroupsForUpsert200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIUserGetResourceGroupsForUpsert200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIUserGetResourceGroupsForUpsert200ResponseType() SCIUserGetResourceGroupsForUpsert200ResponseType {
	m := make(SCIUserGetResourceGroupsForUpsert200ResponseType, 0)
	return m
}

// SCIUserGetUsers200ResponseType
//
// Definition: user_getUsers200ResponseType
type SCIUserGetUsers200ResponseType []*SCIModelsUser

type SCIUserGetUsers200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIUserGetUsers200ResponseType
}

func newSCIUserGetUsers200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIUserGetUsers200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIUserGetUsers200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIUserGetUsers200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIUserGetUsers200ResponseType() SCIUserGetUsers200ResponseType {
	m := make(SCIUserGetUsers200ResponseType, 0)
	return m
}

// SCIUserPrototypegetfilters200ResponseType
//
// Definition: user_prototype_get_filters200ResponseType
type SCIUserPrototypegetfilters200ResponseType []*SCIModelsFilter

type SCIUserPrototypegetfilters200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIUserPrototypegetfilters200ResponseType
}

func newSCIUserPrototypegetfilters200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIUserPrototypegetfilters200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIUserPrototypegetfilters200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIUserPrototypegetfilters200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIUserPrototypegetfilters200ResponseType() SCIUserPrototypegetfilters200ResponseType {
	m := make(SCIUserPrototypegetfilters200ResponseType, 0)
	return m
}

// SCIUserPrototypegetschedules200ResponseType
//
// Definition: user_prototype_get_schedules200ResponseType
type SCIUserPrototypegetschedules200ResponseType []*SCIModelsSchedule

type SCIUserPrototypegetschedules200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIUserPrototypegetschedules200ResponseType
}

func newSCIUserPrototypegetschedules200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIUserPrototypegetschedules200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIUserPrototypegetschedules200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCIUserPrototypegetschedules200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCIUserPrototypegetschedules200ResponseType() SCIUserPrototypegetschedules200ResponseType {
	m := make(SCIUserPrototypegetschedules200ResponseType, 0)
	return m
}

// SCIUserLoginRequest
//
// Definition: user_userLoginRequest
//
// Credentials used to log a user in
type SCIUserLoginRequest struct {
	// Password
	// Constraints:
	//    - required
	Password *string `json:"password"`

	// Username
	// Constraints:
	//    - required
	Username *string `json:"username"`
}

func NewSCIUserLoginRequest() *SCIUserLoginRequest {
	m := new(SCIUserLoginRequest)
	return m
}

// SCIUserLoginResponse
//
// Definition: user_userLoginResponse
//
// User Login details
type SCIUserLoginResponse struct {
	Created *string `json:"created,omitempty"`

	Id *string `json:"id,omitempty"`

	ResourceGroup *string `json:"resourceGroup,omitempty"`

	Role *string `json:"role,omitempty"`

	Ttl *int `json:"ttl,omitempty"`

	UserId *int `json:"userId,omitempty"`

	Username *string `json:"username,omitempty"`
}

type SCIUserLoginResponseAPIResponse struct {
	*RawAPIResponse
	Data *SCIUserLoginResponse
}

func newSCIUserLoginResponseAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIUserLoginResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIUserLoginResponseAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SCIUserLoginResponse)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSCIUserLoginResponse() *SCIUserLoginResponse {
	m := new(SCIUserLoginResponse)
	return m
}

// UserBatchDelete
//
// Delete users and remove them from their related models.
//
// Operation ID: user_batchDelete
// Operation path: /users/batchDelete
// Success code: 200 (OK)
//
// Form data parameters:
// - ids string
//		- required
func (s *SCIUserService) UserBatchDelete(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIUserBatchDelete200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIUserBatchDelete200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIUserBatchDelete, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(formValues)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIUserBatchDelete200ResponseTypeAPIResponse), err
}

// UserCreateWithRelations
//
// Create user and its related models.
//
// Operation ID: user_createWithRelations
// Operation path: /users/createWithRelations
// Success code: 200 (OK)
//
// Form data parameters:
// - email string
//		- required
// - firstName string
//		- required
// - isExternal bool
//		- nullable
//		- default:false
// - lastName string
//		- required
// - password string
//		- required
// - resourceGroupId float64
//		- required
// - role string
//		- required
// - username string
//		- required
func (s *SCIUserService) UserCreateWithRelations(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIModelsUserAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsUserAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIUserCreateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(formValues)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsUserAPIResponse), err
}

// UserFindById
//
// Find a model instance by id from the data source.
//
// Operation ID: user_findById
// Operation path: /users/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIUserService) UserFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsUserAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsUserAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIUserFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsUserAPIResponse), err
}

// UserGetResourceGroupsForUpsert
//
// Get resource groups that current user can manage.
//
// Operation ID: user_getResourceGroupsForUpsert
// Operation path: /users/getResourceGroupsForUpsert
// Success code: 200 (OK)
func (s *SCIUserService) UserGetResourceGroupsForUpsert(ctx context.Context, mutators ...RequestMutator) (*SCIUserGetResourceGroupsForUpsert200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIUserGetResourceGroupsForUpsert200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIUserGetResourceGroupsForUpsert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIUserGetResourceGroupsForUpsert200ResponseTypeAPIResponse), err
}

// UserGetUsers
//
// Get users that current user can manage.
//
// Operation ID: user_getUsers
// Operation path: /users
// Success code: 200 (OK)
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIUserService) UserGetUsers(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIUserGetUsers200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIUserGetUsers200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIUserGetUsers, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIUserGetUsers200ResponseTypeAPIResponse), err
}

// UserLogin
//
// Login a user with username/email and password.
//
// Operation ID: user_login
// Operation path: /users/login
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIUserLoginRequest
//
// Optional parameters:
// - include string
//		- nullable
func (s *SCIUserService) UserLogin(ctx context.Context, credentials *SCIUserLoginRequest, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIUserLoginResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIUserLoginResponseAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIUserLogin, false)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(credentials)
	if v, ok := optionalParams["include"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("include", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIUserLoginResponseAPIResponse), err
}

// UserLogout
//
// Logout a user with access token.
//
// Operation ID: user_logout
// Operation path: /users/logout
// Success code: 204 (No Content)
//
// Required parameters:
// - accesstoken string
//		- required
func (s *SCIUserService) UserLogout(ctx context.Context, accesstoken string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIUserLogout, false)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("access_token", accesstoken)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UserPrototypeCreateFilters
//
// Creates a new instance in filters of this model.
//
// Operation ID: user_prototype_create_filters
// Operation path: /users/{id}/filters
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsFilter
//
// Required parameters:
// - id string
//		- required
func (s *SCIUserService) UserPrototypeCreateFilters(ctx context.Context, data *SCIModelsFilter, id string, mutators ...RequestMutator) (*SCIModelsFilterAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsFilterAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIUserPrototypeCreateFilters, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(data)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsFilterAPIResponse), err
}

// UserPrototypeDestroyByIdFilters
//
// Delete a related item by id for filters.
//
// Operation ID: user_prototype_destroyById_filters
// Operation path: /users/{id}/filters/{fk}
// Success code: 204 (No Content)
//
// Required parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIUserService) UserPrototypeDestroyByIdFilters(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodDelete, RouteSCIUserPrototypeDestroyByIdFilters, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UserPrototypeFindByIdFilters
//
// Find a related item by id for filters.
//
// Operation ID: user_prototype_findById_filters
// Operation path: /users/{id}/filters/{fk}
// Success code: 200 (OK)
//
// Required parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIUserService) UserPrototypeFindByIdFilters(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*SCIModelsFilterAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsFilterAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIUserPrototypeFindByIdFilters, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsFilterAPIResponse), err
}

// UserPrototypeGetFilters
//
// Queries filters of user.
//
// Operation ID: user_prototype_get_filters
// Operation path: /users/{id}/filters
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIUserService) UserPrototypeGetFilters(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIUserPrototypegetfilters200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIUserPrototypegetfilters200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIUserPrototypeGetFilters, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIUserPrototypegetfilters200ResponseTypeAPIResponse), err
}

// UserPrototypeGetSchedules
//
// Queries schedules of user.
//
// Operation ID: user_prototype_get_schedules
// Operation path: /users/{id}/schedules
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIUserService) UserPrototypeGetSchedules(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIUserPrototypegetschedules200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIUserPrototypegetschedules200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIUserPrototypeGetSchedules, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIUserPrototypegetschedules200ResponseTypeAPIResponse), err
}

// UserPrototypeUpdateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Operation ID: user_prototype_updateAttributes
// Operation path: /users/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsUser
//
// Required parameters:
// - id string
//		- required
func (s *SCIUserService) UserPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsUser, id string, mutators ...RequestMutator) (*SCIModelsUserAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsUserAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPut, RouteSCIUserPrototypeUpdateAttributes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(data)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsUserAPIResponse), err
}

// UserPrototypeUpdateByIdFilters
//
// Update a related item by id for filters.
//
// Operation ID: user_prototype_updateById_filters
// Operation path: /users/{id}/filters/{fk}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsFilter
//
// Required parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIUserService) UserPrototypeUpdateByIdFilters(ctx context.Context, data *SCIModelsFilter, fk string, id string, mutators ...RequestMutator) (*SCIModelsFilterAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsFilterAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPut, RouteSCIUserPrototypeUpdateByIdFilters, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(data)
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsFilterAPIResponse), err
}

// UserUpdateWithRelations
//
// Update a user and its related models.
//
// Operation ID: user_updateWithRelations
// Operation path: /users/{id}/updateWithRelations
// Success code: 200 (OK)
//
// Form data parameters:
// - email string
//		- required
// - firstName string
//		- required
// - lastName string
//		- required
// - password string
//		- nullable
// - resourceGroupId float64
//		- required
// - role string
//		- required
// - username string
//		- required
//
// Required parameters:
// - id string
//		- required
func (s *SCIUserService) UserUpdateWithRelations(ctx context.Context, formValues url.Values, id string, mutators ...RequestMutator) (*SCIModelsUserAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsUserAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPut, RouteSCIUserUpdateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(formValues)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsUserAPIResponse), err
}

// UserValidateCurrentPassword
//
// check if current password entered is valid
//
// Operation ID: user_validateCurrentPassword
// Operation path: /users/validateCurrentPassword
// Success code: 200 (OK)
//
// Form data parameters:
// - currentPassword string
//		- required
func (s *SCIUserService) UserValidateCurrentPassword(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIModelsUserAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsUserAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIUserValidateCurrentPassword, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(formValues)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsUserAPIResponse), err
}
