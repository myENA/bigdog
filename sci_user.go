package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
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

type SCIUserBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIUserBatchDelete200ResponseType() *SCIUserBatchDelete200ResponseType {
	m := new(SCIUserBatchDelete200ResponseType)
	return m
}

type SCIUserGetResourceGroupsForUpsert200ResponseType []interface{}

func MakeSCIUserGetResourceGroupsForUpsert200ResponseType() SCIUserGetResourceGroupsForUpsert200ResponseType {
	m := make(SCIUserGetResourceGroupsForUpsert200ResponseType, 0)
	return m
}

type SCIUserGetUsers200ResponseType []*SCIModelsUser

func MakeSCIUserGetUsers200ResponseType() SCIUserGetUsers200ResponseType {
	m := make(SCIUserGetUsers200ResponseType, 0)
	return m
}

// SCIUserLogin200ResponseType
//
// The response body contains properties of the AccessToken created on login.
// Depending on the value of `include` parameter, the body may contain additional properties:
//
//   - `user` - `{User}` - Data of the currently logged in user. (`include=user`)
type SCIUserLogin200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIUserLogin200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIUserLogin200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIUserLogin200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIUserLogin200ResponseType() *SCIUserLogin200ResponseType {
	m := new(SCIUserLogin200ResponseType)
	return m
}

type SCIUsergetfilters200ResponseType []*SCIModelsFilter

func MakeSCIUsergetfilters200ResponseType() SCIUsergetfilters200ResponseType {
	m := make(SCIUsergetfilters200ResponseType, 0)
	return m
}

type SCIUsergetschedules200ResponseType []*SCIModelsSchedule

func MakeSCIUsergetschedules200ResponseType() SCIUsergetschedules200ResponseType {
	m := make(SCIUsergetschedules200ResponseType, 0)
	return m
}

// UserBatchDelete
//
// Delete users and remove them from their related models.
//
// Request Body:
//	 - body string
func (s *SCIUserService) UserBatchDelete(ctx context.Context, body string) (*SCIUserBatchDelete200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIUserBatchDelete, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIUserBatchDelete200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UserCreateWithRelations
//
// Create user and its related models.
//
// Request Body:
//	 - body string
func (s *SCIUserService) UserCreateWithRelations(ctx context.Context, body string) (*SCIModelsUser, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIUserCreateWithRelations, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UserFindById
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
func (s *SCIUserService) UserFindById(ctx context.Context, id string, optionalParams map[string][]string) (*SCIModelsUser, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIUserFindById, false)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UserGetResourceGroupsForUpsert
//
// Get resource groups that current user can manage.
func (s *SCIUserService) UserGetResourceGroupsForUpsert(ctx context.Context) (SCIUserGetResourceGroupsForUpsert200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIUserGetResourceGroupsForUpsert, false)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIUserGetResourceGroupsForUpsert200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UserGetUsers
//
// Get users that current user can manage.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIUserService) UserGetUsers(ctx context.Context, optionalParams map[string][]string) (SCIUserGetUsers200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIUserGetUsers, false)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIUserGetUsers200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UserLogin
//
// Login a user with username/email and password.
//
// Request Body:
//	 - body *SCIModelsUserLogin
//
// Optional Parameters:
// - include string
//		- nullable
func (s *SCIUserService) UserLogin(ctx context.Context, body *SCIModelsUserLogin, optionalParams map[string][]string) (*SCIUserLogin200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIUserLogin200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIUserLogin, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	if v, ok := optionalParams["include"]; ok && len(v) > 0 {
		req.SetQueryParameter("include", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIUserLogin200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UserLogout
//
// Logout a user with access token.
func (s *SCIUserService) UserLogout(ctx context.Context) (*APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UserPrototypeCreateFilters
//
// Creates a new instance in filters of this model.
//
// Request Body:
//	 - body *SCIModelsFilter
//
// Required Parameters:
// - id string
//		- required
func (s *SCIUserService) UserPrototypeCreateFilters(ctx context.Context, body *SCIModelsFilter, id string) (*SCIModelsFilter, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIUserPrototypeCreateFilters, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsFilter()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UserPrototypeDestroyByIdFilters
//
// Delete a related item by id for filters.
//
// Required Parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIUserService) UserPrototypeDestroyByIdFilters(ctx context.Context, fk string, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCIUserPrototypeDestroyByIdFilters, false)
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UserPrototypeFindByIdFilters
//
// Find a related item by id for filters.
//
// Required Parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIUserService) UserPrototypeFindByIdFilters(ctx context.Context, fk string, id string) (*SCIModelsFilter, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIUserPrototypeFindByIdFilters, false)
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsFilter()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UserPrototypeGetFilters
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
func (s *SCIUserService) UserPrototypeGetFilters(ctx context.Context, id string, optionalParams map[string][]string) (SCIUsergetfilters200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIUsergetfilters200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIUserPrototypeGetFilters, false)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIUsergetfilters200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UserPrototypeGetSchedules
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
func (s *SCIUserService) UserPrototypeGetSchedules(ctx context.Context, id string, optionalParams map[string][]string) (SCIUsergetschedules200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIUsergetschedules200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIUserPrototypeGetSchedules, false)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIUsergetschedules200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UserPrototypeUpdateByIdFilters
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
func (s *SCIUserService) UserPrototypeUpdateByIdFilters(ctx context.Context, body *SCIModelsFilter, fk string, id string) (*SCIModelsFilter, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIUserPrototypeUpdateByIdFilters, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsFilter()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UserPrototypeUpdateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsUser
//
// Required Parameters:
// - id string
//		- required
func (s *SCIUserService) UserPrototypeUpdateAttributes(ctx context.Context, body *SCIModelsUser, id string) (*SCIModelsUser, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIUserPrototypeUpdateAttributes, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UserUpdateWithRelations
//
// Update a user and its related models.
//
// Request Body:
//	 - body string
//
// Required Parameters:
// - id string
//		- required
func (s *SCIUserService) UserUpdateWithRelations(ctx context.Context, body string, id string) (*SCIModelsUser, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIUserUpdateWithRelations, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UserValidateCurrentPassword
//
// check if current password entered is valid
//
// Request Body:
//	 - body string
func (s *SCIUserService) UserValidateCurrentPassword(ctx context.Context, body string) (*SCIModelsUser, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIUserValidateCurrentPassword, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
