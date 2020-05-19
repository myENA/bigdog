package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGClientTrafficAnalysisService struct {
	apiClient *APIClient
}

func NewWSGClientTrafficAnalysisService(c *APIClient) *WSGClientTrafficAnalysisService {
	s := new(WSGClientTrafficAnalysisService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGClientTrafficAnalysisService() *WSGClientTrafficAnalysisService {
	return NewWSGClientTrafficAnalysisService(ss.apiClient)
}

type WSGClientTrafficAnalysisSummary struct {
	Extra *WSGClientTrafficAnalysisSummaryExtraType `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGClientTrafficAnalysisSummaryListType `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGClientTrafficAnalysisSummary() *WSGClientTrafficAnalysisSummary {
	m := new(WSGClientTrafficAnalysisSummary)
	return m
}

type WSGClientTrafficAnalysisSummaryExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGClientTrafficAnalysisSummaryExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGClientTrafficAnalysisSummaryExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGClientTrafficAnalysisSummaryExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGClientTrafficAnalysisSummaryExtraType() *WSGClientTrafficAnalysisSummaryExtraType {
	m := new(WSGClientTrafficAnalysisSummaryExtraType)
	return m
}

type WSGClientTrafficAnalysisSummaryListType struct {
	ExtraValues24 *WSGClientTrafficAnalysisSummaryListTypeExtraValues24Type `json:"extraValues24,omitempty"`

	ExtraValues50 *WSGClientTrafficAnalysisSummaryListTypeExtraValues50Type `json:"extraValues50,omitempty"`

	ExtraValuesTotal *WSGClientTrafficAnalysisSummaryListTypeExtraValuesTotalType `json:"extraValuesTotal,omitempty"`

	Key *string `json:"key,omitempty"`

	Total *float64 `json:"total,omitempty"`

	Value *float64 `json:"value,omitempty"`

	Value24 *float64 `json:"value24,omitempty"`

	Value50 *float64 `json:"value50,omitempty"`
}

func NewWSGClientTrafficAnalysisSummaryListType() *WSGClientTrafficAnalysisSummaryListType {
	m := new(WSGClientTrafficAnalysisSummaryListType)
	return m
}

type WSGClientTrafficAnalysisSummaryListTypeExtraValues24Type struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGClientTrafficAnalysisSummaryListTypeExtraValues24Type) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGClientTrafficAnalysisSummaryListTypeExtraValues24Type{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGClientTrafficAnalysisSummaryListTypeExtraValues24Type) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGClientTrafficAnalysisSummaryListTypeExtraValues24Type() *WSGClientTrafficAnalysisSummaryListTypeExtraValues24Type {
	m := new(WSGClientTrafficAnalysisSummaryListTypeExtraValues24Type)
	return m
}

type WSGClientTrafficAnalysisSummaryListTypeExtraValues50Type struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGClientTrafficAnalysisSummaryListTypeExtraValues50Type) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGClientTrafficAnalysisSummaryListTypeExtraValues50Type{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGClientTrafficAnalysisSummaryListTypeExtraValues50Type) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGClientTrafficAnalysisSummaryListTypeExtraValues50Type() *WSGClientTrafficAnalysisSummaryListTypeExtraValues50Type {
	m := new(WSGClientTrafficAnalysisSummaryListTypeExtraValues50Type)
	return m
}

type WSGClientTrafficAnalysisSummaryListTypeExtraValuesTotalType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGClientTrafficAnalysisSummaryListTypeExtraValuesTotalType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGClientTrafficAnalysisSummaryListTypeExtraValuesTotalType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGClientTrafficAnalysisSummaryListTypeExtraValuesTotalType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGClientTrafficAnalysisSummaryListTypeExtraValuesTotalType() *WSGClientTrafficAnalysisSummaryListTypeExtraValuesTotalType {
	m := new(WSGClientTrafficAnalysisSummaryListTypeExtraValuesTotalType)
	return m
}

// FindClientTrafficAnalysisByQueryCriteria
//
// View wireless client traffic analysis data
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGClientTrafficAnalysisService) FindClientTrafficAnalysisByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGClientTrafficAnalysisSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClientTrafficAnalysisSummary
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindClientTrafficAnalysisByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClientTrafficAnalysisSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
