package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIFacetService struct {
	apiClient *SCIClient
}

func NewSCIFacetService(c *SCIClient) *SCIFacetService {
	s := new(SCIFacetService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIFacetService() *SCIFacetService {
	return NewSCIFacetService(ss.apiClient)
}

// SCIFacetGetApmacFacet200ResponseType
//
// Definition: facet_getApmacFacet200ResponseType
type SCIFacetGetApmacFacet200ResponseType []string

func MakeSCIFacetGetApmacFacet200ResponseType() SCIFacetGetApmacFacet200ResponseType {
	m := make(SCIFacetGetApmacFacet200ResponseType, 0)
	return m
}

// SCIFacetGetFacet200ResponseType
//
// Definition: facet_getFacet200ResponseType
type SCIFacetGetFacet200ResponseType []interface{}

func MakeSCIFacetGetFacet200ResponseType() SCIFacetGetFacet200ResponseType {
	m := make(SCIFacetGetFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSsidFacet200ResponseType
//
// Definition: facet_getSsidFacet200ResponseType
type SCIFacetGetSsidFacet200ResponseType []string

func MakeSCIFacetGetSsidFacet200ResponseType() SCIFacetGetSsidFacet200ResponseType {
	m := make(SCIFacetGetSsidFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSwitchesFacet200ResponseType
//
// Definition: facet_getSwitchesFacet200ResponseType
type SCIFacetGetSwitchesFacet200ResponseType []string

func MakeSCIFacetGetSwitchesFacet200ResponseType() SCIFacetGetSwitchesFacet200ResponseType {
	m := make(SCIFacetGetSwitchesFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSwitchHierarchyFacet200ResponseType
//
// Definition: facet_getSwitchHierarchyFacet200ResponseType
type SCIFacetGetSwitchHierarchyFacet200ResponseType []*SCIFacetSwitchHierarchyFacet

func MakeSCIFacetGetSwitchHierarchyFacet200ResponseType() SCIFacetGetSwitchHierarchyFacet200ResponseType {
	m := make(SCIFacetGetSwitchHierarchyFacet200ResponseType, 0)
	return m
}

// SCIFacetGetSystemFacet200ResponseType
//
// Definition: facet_getSystemFacet200ResponseType
type SCIFacetGetSystemFacet200ResponseType []*SCIFacetSystemFacet

func MakeSCIFacetGetSystemFacet200ResponseType() SCIFacetGetSystemFacet200ResponseType {
	m := make(SCIFacetGetSystemFacet200ResponseType, 0)
	return m
}

// SCIFacetSwitchHierarchyFacet
//
// Definition: facet_switchHierarchyFacet
type SCIFacetSwitchHierarchyFacet struct {
	Children []*SCIFacetSwitchHierarchyFacetChildrenType `json:"children,omitempty"`

	Id *string `json:"id,omitempty"`

	State *SCIFacetSwitchHierarchyFacetStateType `json:"state,omitempty"`

	Text *string `json:"text,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIFacetSwitchHierarchyFacet) UnmarshalJSON(b []byte) error {
	type _SCIFacetSwitchHierarchyFacet SCIFacetSwitchHierarchyFacet
	tmpType := new(_SCIFacetSwitchHierarchyFacet)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "children")
	delete(tmpType.XAdditionalProperties, "id")
	delete(tmpType.XAdditionalProperties, "state")
	delete(tmpType.XAdditionalProperties, "text")
	*t = SCIFacetSwitchHierarchyFacet(*tmpType)
	return nil
}

func (t *SCIFacetSwitchHierarchyFacet) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Children != nil {
		tmp["children"] = t.Children
	}
	if t.Id != nil {
		tmp["id"] = t.Id
	}
	if t.State != nil {
		tmp["state"] = t.State
	}
	if t.Text != nil {
		tmp["text"] = t.Text
	}
	return json.Marshal(tmp)
}

func NewSCIFacetSwitchHierarchyFacet() *SCIFacetSwitchHierarchyFacet {
	m := new(SCIFacetSwitchHierarchyFacet)
	return m
}

// SCIFacetSwitchHierarchyFacetChildrenType
//
// Definition: facet_switchHierarchyFacetChildrenType
type SCIFacetSwitchHierarchyFacetChildrenType struct {
	Children *bool `json:"children,omitempty"`

	Data *string `json:"data,omitempty"`

	FilterText *string `json:"filterText,omitempty"`

	Id *string `json:"id,omitempty"`

	Text *string `json:"text,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIFacetSwitchHierarchyFacetChildrenType) UnmarshalJSON(b []byte) error {
	type _SCIFacetSwitchHierarchyFacetChildrenType SCIFacetSwitchHierarchyFacetChildrenType
	tmpType := new(_SCIFacetSwitchHierarchyFacetChildrenType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "children")
	delete(tmpType.XAdditionalProperties, "data")
	delete(tmpType.XAdditionalProperties, "filterText")
	delete(tmpType.XAdditionalProperties, "id")
	delete(tmpType.XAdditionalProperties, "text")
	*t = SCIFacetSwitchHierarchyFacetChildrenType(*tmpType)
	return nil
}

func (t *SCIFacetSwitchHierarchyFacetChildrenType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Children != nil {
		tmp["children"] = t.Children
	}
	if t.Data != nil {
		tmp["data"] = t.Data
	}
	if t.FilterText != nil {
		tmp["filterText"] = t.FilterText
	}
	if t.Id != nil {
		tmp["id"] = t.Id
	}
	if t.Text != nil {
		tmp["text"] = t.Text
	}
	return json.Marshal(tmp)
}

func NewSCIFacetSwitchHierarchyFacetChildrenType() *SCIFacetSwitchHierarchyFacetChildrenType {
	m := new(SCIFacetSwitchHierarchyFacetChildrenType)
	return m
}

// SCIFacetSwitchHierarchyFacetStateType
//
// Definition: facet_switchHierarchyFacetStateType
type SCIFacetSwitchHierarchyFacetStateType struct {
	Opened *bool `json:"opened,omitempty"`
}

func NewSCIFacetSwitchHierarchyFacetStateType() *SCIFacetSwitchHierarchyFacetStateType {
	m := new(SCIFacetSwitchHierarchyFacetStateType)
	return m
}

// SCIFacetSystemFacet
//
// Definition: facet_systemFacet
type SCIFacetSystemFacet struct {
	Children []*SCIFacetSystemFacetChildrenType `json:"children,omitempty"`

	Id *string `json:"id,omitempty"`

	State *SCIFacetSystemFacetStateType `json:"state,omitempty"`

	Text *string `json:"text,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIFacetSystemFacet) UnmarshalJSON(b []byte) error {
	type _SCIFacetSystemFacet SCIFacetSystemFacet
	tmpType := new(_SCIFacetSystemFacet)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "children")
	delete(tmpType.XAdditionalProperties, "id")
	delete(tmpType.XAdditionalProperties, "state")
	delete(tmpType.XAdditionalProperties, "text")
	*t = SCIFacetSystemFacet(*tmpType)
	return nil
}

func (t *SCIFacetSystemFacet) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Children != nil {
		tmp["children"] = t.Children
	}
	if t.Id != nil {
		tmp["id"] = t.Id
	}
	if t.State != nil {
		tmp["state"] = t.State
	}
	if t.Text != nil {
		tmp["text"] = t.Text
	}
	return json.Marshal(tmp)
}

func NewSCIFacetSystemFacet() *SCIFacetSystemFacet {
	m := new(SCIFacetSystemFacet)
	return m
}

// SCIFacetSystemFacetChildrenType
//
// Definition: facet_systemFacetChildrenType
type SCIFacetSystemFacetChildrenType struct {
	Children *bool `json:"children,omitempty"`

	Data *string `json:"data,omitempty"`

	FilterText *string `json:"filterText,omitempty"`

	Id *string `json:"id,omitempty"`

	Text *string `json:"text,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIFacetSystemFacetChildrenType) UnmarshalJSON(b []byte) error {
	type _SCIFacetSystemFacetChildrenType SCIFacetSystemFacetChildrenType
	tmpType := new(_SCIFacetSystemFacetChildrenType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "children")
	delete(tmpType.XAdditionalProperties, "data")
	delete(tmpType.XAdditionalProperties, "filterText")
	delete(tmpType.XAdditionalProperties, "id")
	delete(tmpType.XAdditionalProperties, "text")
	*t = SCIFacetSystemFacetChildrenType(*tmpType)
	return nil
}

func (t *SCIFacetSystemFacetChildrenType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Children != nil {
		tmp["children"] = t.Children
	}
	if t.Data != nil {
		tmp["data"] = t.Data
	}
	if t.FilterText != nil {
		tmp["filterText"] = t.FilterText
	}
	if t.Id != nil {
		tmp["id"] = t.Id
	}
	if t.Text != nil {
		tmp["text"] = t.Text
	}
	return json.Marshal(tmp)
}

func NewSCIFacetSystemFacetChildrenType() *SCIFacetSystemFacetChildrenType {
	m := new(SCIFacetSystemFacetChildrenType)
	return m
}

// SCIFacetSystemFacetStateType
//
// Definition: facet_systemFacetStateType
type SCIFacetSystemFacetStateType struct {
	Opened *bool `json:"opened,omitempty"`
}

func NewSCIFacetSystemFacetStateType() *SCIFacetSystemFacetStateType {
	m := new(SCIFacetSystemFacetStateType)
	return m
}

// FacetGetApmacFacet
//
// Operation ID: facet_getApmacFacet
//
// Retrieve list of apmac facets based on query
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIFacetService) FacetGetApmacFacet(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (SCIFacetGetApmacFacet200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIFacetGetApmacFacet200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIFacetGetApmacFacet, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIFacetGetApmacFacet200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FacetGetFacet
//
// Operation ID: facet_getFacet
//
// For the <b><code>filter</code></b> field below, an example would be
// <pre>
//   <code class="json">
//     { "type": "and", "fields": [{ "type": "selector", "dimension": "system", "value": "sys1" }]}
//   </code>
// </pre>
//
// Request Body:
//	 - body *SCICommonQueryBody
//
// Required Parameters:
// - name string
//		- required
//		- oneof:[system,switchHierarchy,apmac,ssid,switches]
func (s *SCIFacetService) FacetGetFacet(ctx context.Context, body *SCICommonQueryBody, name string, mutators ...RequestMutator) (SCIFacetGetFacet200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIFacetGetFacet200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIFacetGetFacet, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("name", name)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIFacetGetFacet200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FacetGetSsidFacet
//
// Operation ID: facet_getSsidFacet
//
// Retrieve list of ssid facets based on query
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIFacetService) FacetGetSsidFacet(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (SCIFacetGetSsidFacet200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIFacetGetSsidFacet200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIFacetGetSsidFacet, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIFacetGetSsidFacet200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FacetGetSwitchesFacet
//
// Operation ID: facet_getSwitchesFacet
//
// Retrieve list of switches facets based on query
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIFacetService) FacetGetSwitchesFacet(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (SCIFacetGetSwitchesFacet200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIFacetGetSwitchesFacet200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIFacetGetSwitchesFacet, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIFacetGetSwitchesFacet200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FacetGetSwitchHierarchyFacet
//
// Operation ID: facet_getSwitchHierarchyFacet
//
// Retrieve list of switchHierarchy facets based on query
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIFacetService) FacetGetSwitchHierarchyFacet(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (SCIFacetGetSwitchHierarchyFacet200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIFacetGetSwitchHierarchyFacet200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIFacetGetSwitchHierarchyFacet, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIFacetGetSwitchHierarchyFacet200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FacetGetSystemFacet
//
// Operation ID: facet_getSystemFacet
//
// Retrieve list of system facets based on query
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIFacetService) FacetGetSystemFacet(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (SCIFacetGetSystemFacet200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIFacetGetSystemFacet200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIFacetGetSystemFacet, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIFacetGetSystemFacet200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
