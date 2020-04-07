package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGSCIService struct {
	apiClient *APIClient
}

func NewWSGSCIService(c *APIClient) *WSGSCIService {
	s := new(WSGSCIService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSCIService() *WSGSCIService {
	return NewWSGSCIService(ss.apiClient)
}

type WSGSCICreateSciProfile struct {
	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPassword *string `json:"sciPassword" validate:"required"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile" validate:"required"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerHost *string `json:"sciServerHost" validate:"required"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerPort *string `json:"sciServerPort" validate:"required"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciSystemId *string `json:"sciSystemId" validate:"required"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciUser *string `json:"sciUser" validate:"required"`
}

func NewWSGSCICreateSciProfile() *WSGSCICreateSciProfile {
	m := new(WSGSCICreateSciProfile)
	return m
}

type WSGSCIDeleteSciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	Id *string `json:"id" validate:"required"`
}

func NewWSGSCIDeleteSciProfile() *WSGSCIDeleteSciProfile {
	m := new(WSGSCIDeleteSciProfile)
	return m
}

type WSGSCIDeleteSciProfileList struct {
	// List
	// Constraints:
	//    - nullable
	List []*WSGSCIDeleteSciProfile `json:"list,omitempty" validate:"omitempty,dive"`
}

func NewWSGSCIDeleteSciProfileList() *WSGSCIDeleteSciProfileList {
	m := new(WSGSCIDeleteSciProfileList)
	return m
}

type WSGSCIModifyEventCode struct {
	// SciAcceptedEventCodes
	// Constraints:
	//    - required
	SciAcceptedEventCodes []int `json:"sciAcceptedEventCodes" validate:"required,dive"`
}

func NewWSGSCIModifyEventCode() *WSGSCIModifyEventCode {
	m := new(WSGSCIModifyEventCode)
	return m
}

type WSGSCIModifySciEnabled struct {
	// SciEnabled
	// Is SZ/SCI interface enabled or disabled
	// Constraints:
	//    - required
	SciEnabled *bool `json:"sciEnabled" validate:"required"`
}

func NewWSGSCIModifySciEnabled() *WSGSCIModifySciEnabled {
	m := new(WSGSCIModifySciEnabled)
	return m
}

type WSGSCIModifySciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPassword *string `json:"sciPassword" validate:"required"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile" validate:"required"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerHost *string `json:"sciServerHost" validate:"required"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerPort *string `json:"sciServerPort" validate:"required"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciSystemId *string `json:"sciSystemId" validate:"required"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciUser *string `json:"sciUser" validate:"required"`
}

func NewWSGSCIModifySciProfile() *WSGSCIModifySciProfile {
	m := new(WSGSCIModifySciProfile)
	return m
}

type WSGSCIEventCode struct {
	// FirstIndex
	// Index of the first event code returned from the complete event code set
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more SCI accepted event codes after the currently displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGSCIEventCodeListType `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Total SCI accepted event code count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSCIEventCode() *WSGSCIEventCode {
	m := new(WSGSCIEventCode)
	return m
}

type WSGSCIEventCodeListType struct {
	// Code
	// SCI accepted event code
	// Constraints:
	//    - nullable
	Code *int `json:"code,omitempty"`

	// Type
	// SCI accepted event type
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`
}

func NewWSGSCIEventCodeListType() *WSGSCIEventCodeListType {
	m := new(WSGSCIEventCodeListType)
	return m
}

type WSGSCIProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - nullable
	SciPassword *string `json:"sciPassword,omitempty"`

	// SciPriority
	// Priority of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - nullable
	SciPriority *int `json:"sciPriority,omitempty"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - nullable
	SciProfile *string `json:"sciProfile,omitempty"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - nullable
	SciServerHost *string `json:"sciServerHost,omitempty"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - nullable
	SciServerPort *string `json:"sciServerPort,omitempty"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - nullable
	SciSystemId *string `json:"sciSystemId,omitempty"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - nullable
	SciUser *string `json:"sciUser,omitempty"`
}

func NewWSGSCIProfile() *WSGSCIProfile {
	m := new(WSGSCIProfile)
	return m
}

type WSGSCIProfileList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGSCIProfileListExtraType `json:"extra,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGSCIProfile `json:"list,omitempty" validate:"omitempty,dive"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGSCIProfileList) UnmarshalJSON(b []byte) error {
	tmpt := new(WSGSCIProfileList)
	if err := json.Unmarshal(b, tmpt); err != nil {
		return err
	}
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	delete(tmp, "extra")
	delete(tmp, "list")
	tmpt.XAdditionalProperties = tmp
	*t = *tmpt
	return nil
}

func (t *WSGSCIProfileList) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Extra != nil {
		tmp["extra"] = t.Extra
	}
	if t.List != nil {
		tmp["list"] = t.List
	}
	return json.Marshal(tmp)
}

func NewWSGSCIProfileList() *WSGSCIProfileList {
	m := new(WSGSCIProfileList)
	return m
}

type WSGSCIProfileListExtraType struct {
	// SciEnabled
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - nullable
	SciEnabled *bool `json:"sciEnabled,omitempty"`
}

func NewWSGSCIProfileListExtraType() *WSGSCIProfileListExtraType {
	m := new(WSGSCIProfileListExtraType)
	return m
}

// AddSciSciEventCode
//
// Use this API command to modify SciAcceptedEventCodes.
//
// Request Body:
//	 - body *WSGSCIModifyEventCode
func (s *WSGSCIService) AddSciSciEventCode(ctx context.Context, body *WSGSCIModifyEventCode) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSciSciEventCode, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSciSciProfile
//
// Use this API command to create sciProfile.
//
// Request Body:
//	 - body *WSGSCICreateSciProfile
func (s *WSGSCIService) AddSciSciProfile(ctx context.Context, body *WSGSCICreateSciProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSciSciProfile, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteSciSciProfile
//
// Use this API command to delete sciProfile list.
//
// Request Body:
//	 - body *WSGSCIDeleteSciProfileList
func (s *WSGSCIService) DeleteSciSciProfile(ctx context.Context, body *WSGSCIDeleteSciProfileList) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteSciSciProfile, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteSciSciProfileById
//
// Use this API command to delete sciProfile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSCIService) DeleteSciSciProfileById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteSciSciProfileById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}

// FindSciSciEventCode
//
// Use this API command to retrieve SciAcceptedEventCodes.
func (s *WSGSCIService) FindSciSciEventCode(ctx context.Context) (*WSGSCIEventCode, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCIEventCode
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSciSciEventCode, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCIEventCode()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSciSciProfile
//
// Use this API command to retrieve sciProfile list.
func (s *WSGSCIService) FindSciSciProfile(ctx context.Context) (*WSGSCIProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCIProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSciSciProfile, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCIProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSciSciProfileById
//
// Use this API command to retrieve sciProfile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSCIService) FindSciSciProfileById(ctx context.Context, id string) (*WSGSCIProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCIProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSciSciProfileById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCIProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateSciSciEnabled
//
// Use this API command to modify SCI settings is enabled or not.
//
// Request Body:
//	 - body *WSGSCIModifySciEnabled
func (s *WSGSCIService) PartialUpdateSciSciEnabled(ctx context.Context, body *WSGSCIModifySciEnabled) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSciSciEnabled, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateSciSciProfileById
//
// Use this API command to modify sciProfile.
//
// Request Body:
//	 - body *WSGSCIModifySciProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSCIService) PartialUpdateSciSciProfileById(ctx context.Context, body *WSGSCIModifySciProfile, id string) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSciSciProfileById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
