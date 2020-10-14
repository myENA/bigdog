package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGWLANGroupService struct {
	apiClient *VSZClient
}

func NewWSGWLANGroupService(c *VSZClient) *WSGWLANGroupService {
	s := new(WSGWLANGroupService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWLANGroupService() *WSGWLANGroupService {
	return NewWSGWLANGroupService(ss.apiClient)
}

// WSGWLANGroupCreateWlanGroup
//
// Definition: wlangroup_createWlanGroup
type WSGWLANGroupCreateWlanGroup struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGWLANGroupCreateWlanGroup() *WSGWLANGroupCreateWlanGroup {
	m := new(WSGWLANGroupCreateWlanGroup)
	return m
}

// WSGWLANGroupModifyWlanGroup
//
// Definition: wlangroup_modifyWlanGroup
type WSGWLANGroupModifyWlanGroup struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGWLANGroupModifyWlanGroup() *WSGWLANGroupModifyWlanGroup {
	m := new(WSGWLANGroupModifyWlanGroup)
	return m
}

// WSGWLANGroupModifyWlanGroupMember
//
// Definition: wlangroup_modifyWlanGroupMember
type WSGWLANGroupModifyWlanGroupMember struct {
	// AccessVlan
	// Access VLAN
	// Constraints:
	//    - min:1
	//    - max:4094
	AccessVlan *int `json:"accessVlan,omitempty"`

	// NasId
	// NAS-ID
	// Constraints:
	//    - max:63
	NasId *string `json:"nasId,omitempty"`

	VlanPooling *WSGCommonGenericRef `json:"vlanPooling,omitempty"`
}

func NewWSGWLANGroupModifyWlanGroupMember() *WSGWLANGroupModifyWlanGroupMember {
	m := new(WSGWLANGroupModifyWlanGroupMember)
	return m
}

// WSGWLANGroup
//
// Definition: wlangroup_wlanGroup
type WSGWLANGroup struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the WLAN group
	Id *string `json:"id,omitempty"`

	// Members
	// Members of the WLAN group
	Members []*WSGWLANGroupWlanMember `json:"members,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// ZoneId
	// Identifier of the zone to which the WLAN group belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGWLANGroup() *WSGWLANGroup {
	m := new(WSGWLANGroup)
	return m
}

// WSGWLANGroupList
//
// Definition: wlangroup_wlanGroupList
type WSGWLANGroupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANGroup `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWLANGroupList() *WSGWLANGroupList {
	m := new(WSGWLANGroupList)
	return m
}

// WSGWLANGroupWlanMember
//
// Definition: wlangroup_wlanMember
type WSGWLANGroupWlanMember struct {
	// AccessVlan
	// Access VLAN
	// Constraints:
	//    - min:1
	//    - max:4094
	AccessVlan *int `json:"accessVlan,omitempty"`

	// Id
	// Identifier of the WLAN
	// Constraints:
	//    - required
	Id *string `json:"id"`

	// Name
	// Name of the WLAN
	Name *string `json:"name,omitempty"`

	// NasId
	// NAS-ID
	// Constraints:
	//    - max:63
	NasId *string `json:"nasId,omitempty"`

	VlanPooling *WSGCommonGenericRef `json:"vlanPooling,omitempty"`
}

func NewWSGWLANGroupWlanMember() *WSGWLANGroupWlanMember {
	m := new(WSGWLANGroupWlanMember)
	return m
}

// AddRkszonesWlangroupsByZoneId
//
// Operation ID: addRkszonesWlangroupsByZoneId
//
// Use this API command to create a new WLAN group.
//
// Request Body:
//	 - body *WSGWLANGroupCreateWlanGroup
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANGroupService) AddRkszonesWlangroupsByZoneId(ctx context.Context, body *WSGWLANGroupCreateWlanGroup, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlangroupsByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlangroupsMembersById
//
// Operation ID: addRkszonesWlangroupsMembersById
//
// Use this API command to add a member to a WLAN group.
//
// Request Body:
//	 - body *WSGWLANGroupWlanMember
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANGroupService) AddRkszonesWlangroupsMembersById(ctx context.Context, body *WSGWLANGroupWlanMember, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlangroupsMembersById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusCreated, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesWlangroupsById
//
// Operation ID: deleteRkszonesWlangroupsById
//
// Use this API command to delete a WLAN group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlangroupsById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesWlangroupsMembersByMemberId
//
// Operation ID: deleteRkszonesWlangroupsMembersByMemberId
//
// Use this API command to remove a member from a WLAN group.
//
// Required Parameters:
// - id string
//		- required
// - memberId string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersByMemberId(ctx context.Context, id string, memberId string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlangroupsMembersByMemberId, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	req.SetPathParameter("memberId", memberId)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesWlangroupsMembersNasIdByMemberId
//
// Operation ID: deleteRkszonesWlangroupsMembersNasIdByMemberId
//
// Use this API command to disable a member NAS-ID override of a WLAN group.
//
// Required Parameters:
// - id string
//		- required
// - memberId string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersNasIdByMemberId(ctx context.Context, id string, memberId string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlangroupsMembersNasIdByMemberId, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	req.SetPathParameter("memberId", memberId)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesWlangroupsMembersVlanOverrideByMemberId
//
// Operation ID: deleteRkszonesWlangroupsMembersVlanOverrideByMemberId
//
// Use this API command to disable a member VLAN override of a WLAN group.
//
// Required Parameters:
// - id string
//		- required
// - memberId string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersVlanOverrideByMemberId(ctx context.Context, id string, memberId string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlangroupsMembersVlanOverrideByMemberId, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	req.SetPathParameter("memberId", memberId)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesWlangroupsById
//
// Operation ID: findRkszonesWlangroupsById
//
// Use this API command to retrieve the WLAN group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANGroupService) FindRkszonesWlangroupsById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGWLANGroup, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWLANGroup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesWlangroupsById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGWLANGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesWlangroupsByZoneId
//
// Operation ID: findRkszonesWlangroupsByZoneId
//
// Use this API command to retrieve the list of WLAN groups within a zone.
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGWLANGroupService) FindRkszonesWlangroupsByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGWLANGroupList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWLANGroupList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesWlangroupsByZoneId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGWLANGroupList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesWlangroupsById
//
// Operation ID: partialUpdateRkszonesWlangroupsById
//
// Use this API command to modify the configuration of a WLAN group.
//
// Request Body:
//	 - body *WSGWLANGroupModifyWlanGroup
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANGroupService) PartialUpdateRkszonesWlangroupsById(ctx context.Context, body *WSGWLANGroupModifyWlanGroup, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesWlangroupsById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateRkszonesWlangroupsMembersByMemberId
//
// Operation ID: partialUpdateRkszonesWlangroupsMembersByMemberId
//
// Use this API command to modify a member of a WLAN group.
//
// Request Body:
//	 - body *WSGWLANGroupModifyWlanGroupMember
//
// Required Parameters:
// - id string
//		- required
// - memberId string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANGroupService) PartialUpdateRkszonesWlangroupsMembersByMemberId(ctx context.Context, body *WSGWLANGroupModifyWlanGroupMember, id string, memberId string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesWlangroupsMembersByMemberId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("memberId", memberId)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
