package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMTrafficService struct {
	apiClient *APIClient
}

func NewSwitchMTrafficService(c *APIClient) *SwitchMTrafficService {
	s := new(SwitchMTrafficService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTrafficService() *SwitchMTrafficService {
	return NewSwitchMTrafficService(ss.apiClient)
}

type SwitchMTrafficTopPortErrorQueryResultList struct {
	// Extra
	// Extra information for top port error
	Extra *SwitchMTrafficTopPortErrorQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top port error returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top port error after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top port error count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top port error count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMTrafficTopPortErrorQueryResultList() *SwitchMTrafficTopPortErrorQueryResultList {
	m := new(SwitchMTrafficTopPortErrorQueryResultList)
	return m
}

// SwitchMTrafficTopPortErrorQueryResultListExtraType
//
// Extra information for top port error
type SwitchMTrafficTopPortErrorQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTrafficTopPortErrorQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTrafficTopPortErrorQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTrafficTopPortErrorQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTrafficTopPortErrorQueryResultListExtraType() *SwitchMTrafficTopPortErrorQueryResultListExtraType {
	m := new(SwitchMTrafficTopPortErrorQueryResultListExtraType)
	return m
}

type SwitchMTrafficTopPortTrafficUsageQueryResultList struct {
	// Extra
	// Extra information for top port traffic usage
	Extra *SwitchMTrafficTopPortTrafficUsageQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top port traffic usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top port traffic usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top port traffic usage count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top port traffic usage count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMTrafficTopPortTrafficUsageQueryResultList() *SwitchMTrafficTopPortTrafficUsageQueryResultList {
	m := new(SwitchMTrafficTopPortTrafficUsageQueryResultList)
	return m
}

// SwitchMTrafficTopPortTrafficUsageQueryResultListExtraType
//
// Extra information for top port traffic usage
type SwitchMTrafficTopPortTrafficUsageQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTrafficTopPortTrafficUsageQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTrafficTopPortTrafficUsageQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTrafficTopPortTrafficUsageQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTrafficTopPortTrafficUsageQueryResultListExtraType() *SwitchMTrafficTopPortTrafficUsageQueryResultListExtraType {
	m := new(SwitchMTrafficTopPortTrafficUsageQueryResultListExtraType)
	return m
}

type SwitchMTrafficTopSwitchPoEUtilizationQueryResultList struct {
	// Extra
	// Extra information for top PoE utilization
	Extra *SwitchMTrafficTopSwitchPoEUtilizationQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top PoE usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top PoE usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// PoE utilization count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total PoE utilization count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMTrafficTopSwitchPoEUtilizationQueryResultList() *SwitchMTrafficTopSwitchPoEUtilizationQueryResultList {
	m := new(SwitchMTrafficTopSwitchPoEUtilizationQueryResultList)
	return m
}

// SwitchMTrafficTopSwitchPoEUtilizationQueryResultListExtraType
//
// Extra information for top PoE utilization
type SwitchMTrafficTopSwitchPoEUtilizationQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTrafficTopSwitchPoEUtilizationQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTrafficTopSwitchPoEUtilizationQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTrafficTopSwitchPoEUtilizationQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTrafficTopSwitchPoEUtilizationQueryResultListExtraType() *SwitchMTrafficTopSwitchPoEUtilizationQueryResultListExtraType {
	m := new(SwitchMTrafficTopSwitchPoEUtilizationQueryResultListExtraType)
	return m
}

type SwitchMTrafficTopTrafficUsageQueryResultList struct {
	// Extra
	// Extra information for top traffic usage
	Extra *SwitchMTrafficTopTrafficUsageQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top traffic usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top traffic usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top traffic usage count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top traffic usage count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMTrafficTopTrafficUsageQueryResultList() *SwitchMTrafficTopTrafficUsageQueryResultList {
	m := new(SwitchMTrafficTopTrafficUsageQueryResultList)
	return m
}

// SwitchMTrafficTopTrafficUsageQueryResultListExtraType
//
// Extra information for top traffic usage
type SwitchMTrafficTopTrafficUsageQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTrafficTopTrafficUsageQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTrafficTopTrafficUsageQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTrafficTopTrafficUsageQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTrafficTopTrafficUsageQueryResultListExtraType() *SwitchMTrafficTopTrafficUsageQueryResultListExtraType {
	m := new(SwitchMTrafficTopTrafficUsageQueryResultListExtraType)
	return m
}

type SwitchMTraffic struct {
	// Rx
	// RX traffic of the switch
	Rx *string `json:"rx,omitempty"`

	// Timestamp
	// Timestamp of the switch traffic
	Timestamp *string `json:"timestamp,omitempty"`

	// Total
	// Total traffic of the switch
	Total *int `json:"total,omitempty"`

	// Tx
	// TX traffic of the switch
	Tx *string `json:"tx,omitempty"`
}

func NewSwitchMTraffic() *SwitchMTraffic {
	m := new(SwitchMTraffic)
	return m
}

type SwitchMTrafficQueryResultList struct {
	// Extra
	// Extra information for traffic list
	Extra *SwitchMTrafficQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first traffic list returned out of the complete traffic list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more traffic list after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTraffic `json:"list,omitempty"`

	// RawDataTotalCount
	// Total traffic count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of traffic list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMTrafficQueryResultList() *SwitchMTrafficQueryResultList {
	m := new(SwitchMTrafficQueryResultList)
	return m
}

// SwitchMTrafficQueryResultListExtraType
//
// Extra information for traffic list
type SwitchMTrafficQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTrafficQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTrafficQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTrafficQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTrafficQueryResultListExtraType() *SwitchMTrafficQueryResultListExtraType {
	m := new(SwitchMTrafficQueryResultListExtraType)
	return m
}

type SwitchMTrafficUsage struct {
	// Id
	// Identifier of the Traffic Usage
	Id *string `json:"id,omitempty"`

	// Key
	// Interface of the Traffic Usage
	Key *string `json:"key,omitempty"`

	// Value
	// Total Tx and Rx of Traffic Usage
	Value *float64 `json:"value,omitempty"`
}

func NewSwitchMTrafficUsage() *SwitchMTrafficUsage {
	m := new(SwitchMTrafficUsage)
	return m
}

// AddTrafficTopPoeutilization
//
// Use this API command retrieve the top 10 switches by the PoE utilization.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopPoeutilization(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMTrafficTopSwitchPoEUtilizationQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTrafficTopSwitchPoEUtilizationQueryResultList
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopPoeutilization, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMTrafficTopSwitchPoEUtilizationQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddTrafficTopPorterror
//
// Use this API command to get the top 10 switches by the porterror.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopPorterror(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMTrafficTopPortErrorQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTrafficTopPortErrorQueryResultList
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopPorterror, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMTrafficTopPortErrorQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddTrafficTopPortusage
//
// Use this API command to get the top 10 ports by the traffic.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopPortusage(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMTrafficTopPortTrafficUsageQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTrafficTopPortTrafficUsageQueryResultList
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopPortusage, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMTrafficTopPortTrafficUsageQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddTrafficTopUsage
//
// Use this API command to retrieve Top Swich/Port usage data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopUsage(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMTrafficTopTrafficUsageQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTrafficTopTrafficUsageQueryResultList
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopUsage, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMTrafficTopTrafficUsageQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddTrafficTotalTrend
//
// Use this API command to retrieve Swich/Port trend data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTotalTrend(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMTrafficQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTrafficQueryResultList
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTotalTrend, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMTrafficQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
