package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
)

// WSGDomainDevicePolicyCreateDomainDevicePolicy
//
// Definition: domainDevicePolicy_createDomainDevicePolicy
type WSGDomainDevicePolicyCreateDomainDevicePolicy struct {
	// DefaultAction
	// defaultAction of the device policy config
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// domainId of the device policy config
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Rule
	// rule of the device policy config
	// Constraints:
	//    - required
	Rule []*WSGDomainDevicePolicyRule `json:"rule"`
}

func NewWSGDomainDevicePolicyCreateDomainDevicePolicy() *WSGDomainDevicePolicyCreateDomainDevicePolicy {
	m := new(WSGDomainDevicePolicyCreateDomainDevicePolicy)
	return m
}

// WSGDomainDevicePolicyProfile
//
// Definition: domainDevicePolicy_domainDevicePolicyProfile
type WSGDomainDevicePolicyProfile struct {
	// DefaultAction
	// defaultAction of the device policy config
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// identifier of the device policy config
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rule
	// rule of the device policy config
	Rule []*WSGDomainDevicePolicyRule `json:"rule,omitempty"`
}

type WSGDomainDevicePolicyProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGDomainDevicePolicyProfile
}

func newWSGDomainDevicePolicyProfileAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGDomainDevicePolicyProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGDomainDevicePolicyProfileAPIResponse) Hydrate() error {
	r.Data = new(WSGDomainDevicePolicyProfile)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGDomainDevicePolicyProfile() *WSGDomainDevicePolicyProfile {
	m := new(WSGDomainDevicePolicyProfile)
	return m
}

// WSGDomainDevicePolicyProfileByQueryCriteria
//
// Definition: domainDevicePolicy_domainDevicePolicyProfileByQueryCriteria
type WSGDomainDevicePolicyProfileByQueryCriteria struct {
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first device policy returned out of the complete device policy list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more device policy after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDomainDevicePolicyProfile `json:"list,omitempty"`

	// RawDataTotalCount
	// Total device policy count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Current page device policy count
	TotalCount *int `json:"totalCount,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGDomainDevicePolicyProfileByQueryCriteria) UnmarshalJSON(b []byte) error {
	type _WSGDomainDevicePolicyProfileByQueryCriteria WSGDomainDevicePolicyProfileByQueryCriteria
	tmpType := new(_WSGDomainDevicePolicyProfileByQueryCriteria)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "extra")
	delete(tmpType.XAdditionalProperties, "firstIndex")
	delete(tmpType.XAdditionalProperties, "hasMore")
	delete(tmpType.XAdditionalProperties, "list")
	delete(tmpType.XAdditionalProperties, "rawDataTotalCount")
	delete(tmpType.XAdditionalProperties, "totalCount")
	*t = WSGDomainDevicePolicyProfileByQueryCriteria(*tmpType)
	return nil
}

func (t *WSGDomainDevicePolicyProfileByQueryCriteria) MarshalJSON() ([]byte, error) {
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
	if t.FirstIndex != nil {
		tmp["firstIndex"] = t.FirstIndex
	}
	if t.HasMore != nil {
		tmp["hasMore"] = t.HasMore
	}
	if t.List != nil {
		tmp["list"] = t.List
	}
	if t.RawDataTotalCount != nil {
		tmp["rawDataTotalCount"] = t.RawDataTotalCount
	}
	if t.TotalCount != nil {
		tmp["totalCount"] = t.TotalCount
	}
	return json.Marshal(tmp)
}

type WSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse struct {
	*RawAPIResponse
	Data *WSGDomainDevicePolicyProfileByQueryCriteria
}

func newWSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse) Hydrate() error {
	r.Data = new(WSGDomainDevicePolicyProfileByQueryCriteria)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGDomainDevicePolicyProfileByQueryCriteria() *WSGDomainDevicePolicyProfileByQueryCriteria {
	m := new(WSGDomainDevicePolicyProfileByQueryCriteria)
	return m
}

// WSGDomainDevicePolicyRule
//
// Definition: domainDevicePolicy_domainDevicePolicyRule
type WSGDomainDevicePolicyRule struct {
	// Action
	// defaultAction of the device policy config
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DeviceType
	// deviceType of the device policy rule
	DeviceType *string `json:"deviceType,omitempty"`

	// Downlink
	// downlink rate Limiting of the device policy rule, rate unit is mbps, value must include be 0, 0.10, 0.25, 0.50, 0.75, 1.00, 1.25,1.50, 1.75, 2.00, 2.25, 2.50, 2.75, 3.00, 3.20, 3.50, 3.75, 4.00, 4.25, 4.50, 4.75, 5.00, 5.25, 5.50, 5.75, 6.00, 6.25, 6.50, 6.75,  7.00, 7.25, 7.50, 7.75, 8.00, 8.25, 8.50, 8.75,  9.00, 9.25, 9.50, 9.75,  10.00, 10.25, 10.50, 10.75, 11.00, 11.25, 11.50, 11.75, 12.00, 12.25, 12.50, 12.75, 13.00, 13.25, 13.50, 13.75, 14.00, 14.25, 14.50, 14.75, 15.00, 15.25, 15.50, 15.75, 16.00, 16.25, 16.50, 16.75, 17.00, 17.25, 17.50, 17.75, 18.00, 18.25, 18.50, 18.75, 19.00, 19.25, 19.50, 19.75, 20.00
	// Constraints:
	//    - min:0.000000
	//    - max:200.000000
	Downlink *float64 `json:"downlink,omitempty"`

	// OsVendor
	// osVendor of the device policy rule
	OsVendor *string `json:"osVendor,omitempty"`

	// Uplink
	// uplink rate Limiting of the device policy rule, rate unit is mbps, value must include be 0, 0.10, 0.25, 0.50, 0.75, 1.00, 1.25,1.50, 1.75, 2.00, 2.25, 2.50, 2.75, 3.00, 3.20, 3.50, 3.75, 4.00, 4.25, 4.50, 4.75, 5.00, 5.25, 5.50, 5.75, 6.00, 6.25, 6.50, 6.75,  7.00, 7.25, 7.50, 7.75, 8.00, 8.25, 8.50, 8.75,  9.00, 9.25, 9.50, 9.75,  10.00, 10.25, 10.50, 10.75, 11.00, 11.25, 11.50, 11.75, 12.00, 12.25, 12.50, 12.75, 13.00, 13.25, 13.50, 13.75, 14.00, 14.25, 14.50, 14.75, 15.00, 15.25, 15.50, 15.75, 16.00, 16.25, 16.50, 16.75, 17.00, 17.25, 17.50, 17.75, 18.00, 18.25, 18.50, 18.75, 19.00, 19.25, 19.50, 19.75, 20.00
	// Constraints:
	//    - min:0.000000
	//    - max:200.000000
	Uplink *float64 `json:"uplink,omitempty"`

	// Vlan
	// VLAN Members of the ethernet port profile
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	Vlan *int `json:"vlan,omitempty"`
}

func NewWSGDomainDevicePolicyRule() *WSGDomainDevicePolicyRule {
	m := new(WSGDomainDevicePolicyRule)
	return m
}

// WSGDomainDevicePolicyModifyDomainDevicePolicy
//
// Definition: domainDevicePolicy_modifyDomainDevicePolicy
type WSGDomainDevicePolicyModifyDomainDevicePolicy struct {
	// DefaultAction
	// defaultAction of the device policy config
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rule
	// rule of the device policy config
	Rule []*WSGDomainDevicePolicyRule `json:"rule,omitempty"`
}

func NewWSGDomainDevicePolicyModifyDomainDevicePolicy() *WSGDomainDevicePolicyModifyDomainDevicePolicy {
	m := new(WSGDomainDevicePolicyModifyDomainDevicePolicy)
	return m
}

// WSGDomainDevicePolicyProfileList
//
// Definition: domainDevicePolicy_profileList
type WSGDomainDevicePolicyProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDomainDevicePolicyProfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDomainDevicePolicyProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDomainDevicePolicyProfileList
}

func newWSGDomainDevicePolicyProfileListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGDomainDevicePolicyProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGDomainDevicePolicyProfileListAPIResponse) Hydrate() error {
	r.Data = new(WSGDomainDevicePolicyProfileList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGDomainDevicePolicyProfileList() *WSGDomainDevicePolicyProfileList {
	m := new(WSGDomainDevicePolicyProfileList)
	return m
}

// WSGDomainDevicePolicyProfileListType
//
// Definition: domainDevicePolicy_profileListType
type WSGDomainDevicePolicyProfileListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGDomainDevicePolicyProfileListType() *WSGDomainDevicePolicyProfileListType {
	m := new(WSGDomainDevicePolicyProfileListType)
	return m
}
