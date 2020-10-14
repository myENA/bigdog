package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
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
func (s *SCIUserService) UserBatchDelete(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIUserBatchDelete200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIUserBatchDelete200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIUserBatchDelete, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIUserBatchDelete200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserCreateWithRelations(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIModelsUser, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsUser
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIUserCreateWithRelations, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsUser, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsUser
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIUserFindById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UserGetResourceGroupsForUpsert
//
// Operation ID: user_getResourceGroupsForUpsert
//
// Get resource groups that current user can manage.
func (s *SCIUserService) UserGetResourceGroupsForUpsert(ctx context.Context, mutators ...RequestMutator) (SCIUserGetResourceGroupsForUpsert200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIUserGetResourceGroupsForUpsert200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIUserGetResourceGroupsForUpsert, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIUserGetResourceGroupsForUpsert200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserGetUsers(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCIUserGetUsers200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIUserGetUsers200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIUserGetUsers, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIUserGetUsers200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserLogin(ctx context.Context, credentials *SCIUserLoginRequest, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIUserLoginResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIUserLoginResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIUserLogin, false)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(credentials); err != nil {
		return resp, rm, err
	}
	if v, ok := optionalParams["include"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("include", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIUserLoginResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserLogout(ctx context.Context, accesstoken string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIUserLogout, false)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetQueryParameter("access_token", accesstoken)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *SCIUserService) UserPrototypeCreateFilters(ctx context.Context, data *SCIModelsFilter, id string, mutators ...RequestMutator) (*SCIModelsFilter, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsFilter
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIUserPrototypeCreateFilters, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsFilter()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserPrototypeDestroyByIdFilters(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCIUserPrototypeDestroyByIdFilters, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *SCIUserService) UserPrototypeFindByIdFilters(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*SCIModelsFilter, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsFilter
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIUserPrototypeFindByIdFilters, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsFilter()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserPrototypeGetFilters(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (SCIUserPrototypegetfilters200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIUserPrototypegetfilters200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIUserPrototypeGetFilters, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIUserPrototypegetfilters200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserPrototypeGetSchedules(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (SCIUserPrototypegetschedules200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIUserPrototypegetschedules200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIUserPrototypeGetSchedules, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIUserPrototypegetschedules200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsUser, id string, mutators ...RequestMutator) (*SCIModelsUser, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsUser
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSCIUserPrototypeUpdateAttributes, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserPrototypeUpdateByIdFilters(ctx context.Context, data *SCIModelsFilter, fk string, id string, mutators ...RequestMutator) (*SCIModelsFilter, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsFilter
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSCIUserPrototypeUpdateByIdFilters, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsFilter()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserUpdateWithRelations(ctx context.Context, formValues url.Values, id string, mutators ...RequestMutator) (*SCIModelsUser, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsUser
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSCIUserUpdateWithRelations, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIUserService) UserValidateCurrentPassword(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIModelsUser, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsUser
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIUserValidateCurrentPassword, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
