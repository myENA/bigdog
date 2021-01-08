package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"encoding/json"
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

func newSCIUserBatchDelete200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIUserBatchDelete200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIUserBatchDelete200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIUserBatchDelete200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIUserBatchDelete200ResponseType() *SCIUserBatchDelete200ResponseType {
	m := new(SCIUserBatchDelete200ResponseType)
	return m
}

// SCIUserGetResourceGroupsForUpsert200ResponseType
//
// Definition: user_getResourceGroupsForUpsert200ResponseType
type SCIUserGetResourceGroupsForUpsert200ResponseType []interface{}

func MakeSCIUserGetResourceGroupsForUpsert200ResponseType() SCIUserGetResourceGroupsForUpsert200ResponseType {
	m := make(SCIUserGetResourceGroupsForUpsert200ResponseType, 0)
	return m
}

// SCIUserGetUsers200ResponseType
//
// Definition: user_getUsers200ResponseType
type SCIUserGetUsers200ResponseType []*SCIModelsUser

func MakeSCIUserGetUsers200ResponseType() SCIUserGetUsers200ResponseType {
	m := make(SCIUserGetUsers200ResponseType, 0)
	return m
}

// SCIUserPrototypegetfilters200ResponseType
//
// Definition: user_prototype_get_filters200ResponseType
type SCIUserPrototypegetfilters200ResponseType []*SCIModelsFilter

func MakeSCIUserPrototypegetfilters200ResponseType() SCIUserPrototypegetfilters200ResponseType {
	m := make(SCIUserPrototypegetfilters200ResponseType, 0)
	return m
}

// SCIUserPrototypegetschedules200ResponseType
//
// Definition: user_prototype_get_schedules200ResponseType
type SCIUserPrototypegetschedules200ResponseType []*SCIModelsSchedule

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

func newSCIUserLoginResponseAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIUserLoginResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIUserLoginResponseAPIResponse) Hydrate() error {
	r.Data = new(SCIUserLoginResponse)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIUserLoginResponse() *SCIUserLoginResponse {
	m := new(SCIUserLoginResponse)
	return m
}

// UserBatchDelete
//
// Operation ID: user_batchDelete
//
// Delete users and remove them from their related models.
//
// Form Data Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*SCIUserBatchDelete200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIUserBatchDelete, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*SCIUserBatchDelete200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIUserBatchDelete200ResponseTypeAPIResponse), err
}

// UserCreateWithRelations
//
// Operation ID: user_createWithRelations
//
// Create user and its related models.
//
// Form Data Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsUserAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIUserCreateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*SCIModelsUserAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsUserAPIResponse), err
}

// UserFindById
//
// Operation ID: user_findById
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
func (s *SCIUserService) UserFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsUserAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsUserAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsUserAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIUserFindById, true)
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
// Operation ID: user_getResourceGroupsForUpsert
//
// Get resource groups that current user can manage.
func (s *SCIUserService) UserGetResourceGroupsForUpsert(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCIUserGetResourceGroupsForUpsert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UserGetUsers
//
// Operation ID: user_getUsers
//
// Get users that current user can manage.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIUserService) UserGetUsers(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCIUserGetUsers, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UserLogin
//
// Operation ID: user_login
//
// Login a user with username/email and password.
//
// Request Body:
//	 - body *SCIUserLoginRequest
//
// Optional Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*SCIUserLoginResponseAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIUserLogin, false)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(credentials); err != nil {
		return resp.(*SCIUserLoginResponseAPIResponse), err
	}
	if v, ok := optionalParams["include"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("include", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIUserLoginResponseAPIResponse), err
}

// UserLogout
//
// Operation ID: user_logout
//
// Logout a user with access token.
//
// Required Parameters:
// - accesstoken string
//		- required
func (s *SCIUserService) UserLogout(ctx context.Context, accesstoken string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIUserLogout, false)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("access_token", accesstoken)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UserPrototypeCreateFilters
//
// Operation ID: user_prototype_create_filters
//
// Creates a new instance in filters of this model.
//
// Request Body:
//	 - body *SCIModelsFilter
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsFilterAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIUserPrototypeCreateFilters, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsFilterAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsFilterAPIResponse), err
}

// UserPrototypeDestroyByIdFilters
//
// Operation ID: user_prototype_destroyById_filters
//
// Delete a related item by id for filters.
//
// Required Parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIUserService) UserPrototypeDestroyByIdFilters(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSCIUserPrototypeDestroyByIdFilters, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UserPrototypeFindByIdFilters
//
// Operation ID: user_prototype_findById_filters
//
// Find a related item by id for filters.
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsFilterAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIUserPrototypeFindByIdFilters, true)
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
// Operation ID: user_prototype_get_filters
//
// Queries filters of user.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIUserService) UserPrototypeGetFilters(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCIUserPrototypeGetFilters, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UserPrototypeGetSchedules
//
// Operation ID: user_prototype_get_schedules
//
// Queries schedules of user.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIUserService) UserPrototypeGetSchedules(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCIUserPrototypeGetSchedules, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UserPrototypeUpdateAttributes
//
// Operation ID: user_prototype_updateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsUser
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsUserAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIUserPrototypeUpdateAttributes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsUserAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsUserAPIResponse), err
}

// UserPrototypeUpdateByIdFilters
//
// Operation ID: user_prototype_updateById_filters
//
// Update a related item by id for filters.
//
// Request Body:
//	 - body *SCIModelsFilter
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsFilterAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIUserPrototypeUpdateByIdFilters, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsFilterAPIResponse), err
	}
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsFilterAPIResponse), err
}

// UserUpdateWithRelations
//
// Operation ID: user_updateWithRelations
//
// Update a user and its related models.
//
// Form Data Parameters:
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
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsUserAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIUserUpdateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*SCIModelsUserAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsUserAPIResponse), err
}

// UserValidateCurrentPassword
//
// Operation ID: user_validateCurrentPassword
//
// check if current password entered is valid
//
// Form Data Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsUserAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIUserValidateCurrentPassword, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*SCIModelsUserAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsUserAPIResponse), err
}
