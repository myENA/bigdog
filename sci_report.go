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

type SCIReportService struct {
	apiClient *SCIClient
}

func NewSCIReportService(c *SCIClient) *SCIReportService {
	s := new(SCIReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIReportService() *SCIReportService {
	return NewSCIReportService(ss.apiClient)
}

// SCIReportAirtimeUtilizationReport1overview200ResponseType
//
// Definition: report_AirtimeUtilizationReport_1_overview200ResponseType
type SCIReportAirtimeUtilizationReport1overview200ResponseType struct {
	Data SCIAirtimeUtilizationReport1overviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAirtimeUtilizationReport1overview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAirtimeUtilizationReport1overview200ResponseType
}

func newSCIReportAirtimeUtilizationReport1overview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAirtimeUtilizationReport1overview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAirtimeUtilizationReport1overview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAirtimeUtilizationReport1overview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAirtimeUtilizationReport1overview200ResponseType() *SCIReportAirtimeUtilizationReport1overview200ResponseType {
	m := new(SCIReportAirtimeUtilizationReport1overview200ResponseType)
	return m
}

// SCIReportAirtimeUtilizationReport2topChart200ResponseType
//
// Definition: report_AirtimeUtilizationReport_2_topChart200ResponseType
type SCIReportAirtimeUtilizationReport2topChart200ResponseType struct {
	Data SCIAirtimeUtilizationReport2topChartData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAirtimeUtilizationReport2topChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAirtimeUtilizationReport2topChart200ResponseType
}

func newSCIReportAirtimeUtilizationReport2topChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAirtimeUtilizationReport2topChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAirtimeUtilizationReport2topChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAirtimeUtilizationReport2topChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAirtimeUtilizationReport2topChart200ResponseType() *SCIReportAirtimeUtilizationReport2topChart200ResponseType {
	m := new(SCIReportAirtimeUtilizationReport2topChart200ResponseType)
	return m
}

// SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType
//
// Definition: report_AirtimeUtilizationReport_3_topAPsByAirtime24Table200ResponseType
type SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType struct {
	Data SCIAirtimeUtilizationReport3topAPsByAirtime24TableData `json:"data,omitempty"`

	Metadata *SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData `json:"metadata,omitempty"`
}

type SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType
}

func newSCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType() *SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType {
	m := new(SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType)
	return m
}

// SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType
//
// Definition: report_AirtimeUtilizationReport_4_topAPsByAirtime5Table200ResponseType
type SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType struct {
	Data SCIAirtimeUtilizationReport4topAPsByAirtime5TableData `json:"data,omitempty"`

	Metadata *SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData `json:"metadata,omitempty"`
}

type SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType
}

func newSCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType() *SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType {
	m := new(SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType)
	return m
}

// SCIReportAirtimeUtilizationReport5trendChart200ResponseType
//
// Definition: report_AirtimeUtilizationReport_5_trendChart200ResponseType
type SCIReportAirtimeUtilizationReport5trendChart200ResponseType struct {
	Data SCIAirtimeUtilizationReport5trendChartData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAirtimeUtilizationReport5trendChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAirtimeUtilizationReport5trendChart200ResponseType
}

func newSCIReportAirtimeUtilizationReport5trendChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAirtimeUtilizationReport5trendChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAirtimeUtilizationReport5trendChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAirtimeUtilizationReport5trendChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAirtimeUtilizationReport5trendChart200ResponseType() *SCIReportAirtimeUtilizationReport5trendChart200ResponseType {
	m := new(SCIReportAirtimeUtilizationReport5trendChart200ResponseType)
	return m
}

// SCIReportAirtimeUtilizationReport6trendTable200ResponseType
//
// Definition: report_AirtimeUtilizationReport_6_trendTable200ResponseType
type SCIReportAirtimeUtilizationReport6trendTable200ResponseType struct {
	Data SCIAirtimeUtilizationReport6trendTableData `json:"data,omitempty"`

	Metadata *SCIAirtimeUtilizationReport6trendTableMetaData `json:"metadata,omitempty"`
}

type SCIReportAirtimeUtilizationReport6trendTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAirtimeUtilizationReport6trendTable200ResponseType
}

func newSCIReportAirtimeUtilizationReport6trendTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAirtimeUtilizationReport6trendTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAirtimeUtilizationReport6trendTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAirtimeUtilizationReport6trendTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAirtimeUtilizationReport6trendTable200ResponseType() *SCIReportAirtimeUtilizationReport6trendTable200ResponseType {
	m := new(SCIReportAirtimeUtilizationReport6trendTable200ResponseType)
	return m
}

// SCIReportAPDetailsReport5trendChart200ResponseType
//
// Definition: report_APDetailsReport_5_trendChart200ResponseType
type SCIReportAPDetailsReport5trendChart200ResponseType struct {
	Data SCIAPDetailsReport5trendChartData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport5trendChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport5trendChart200ResponseType
}

func newSCIReportAPDetailsReport5trendChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport5trendChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport5trendChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport5trendChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport5trendChart200ResponseType() *SCIReportAPDetailsReport5trendChart200ResponseType {
	m := new(SCIReportAPDetailsReport5trendChart200ResponseType)
	return m
}

// SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType
//
// Definition: report_APDetailsReport_7_top10ApplicationsByTrafficVolume200ResponseType
type SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType
}

func newSCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType() *SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType {
	m := new(SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType)
	return m
}

// SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType
//
// Definition: report_APDetailsReport_8_topAppsByTrafficTable200ResponseType
type SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType struct {
	Data SCIAPDetailsReport8topAppsByTrafficTableData `json:"data,omitempty"`

	Metadata *SCIAPDetailsReport8topAppsByTrafficTableMetaData `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType
}

func newSCIReportAPDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType() *SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType {
	m := new(SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType)
	return m
}

// SCIReportAPDetailsReport14topTable200ResponseType
//
// Definition: report_APDetailsReport_14_topTable200ResponseType
type SCIReportAPDetailsReport14topTable200ResponseType struct {
	Data SCIAPDetailsReport14topTableData `json:"data,omitempty"`

	Metadata *SCIAPDetailsReport14topTableMetaData `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport14topTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport14topTable200ResponseType
}

func newSCIReportAPDetailsReport14topTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport14topTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport14topTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport14topTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport14topTable200ResponseType() *SCIReportAPDetailsReport14topTable200ResponseType {
	m := new(SCIReportAPDetailsReport14topTable200ResponseType)
	return m
}

// SCIReportAPDetailsReport15trendChart200ResponseType
//
// Definition: report_APDetailsReport_15_trendChart200ResponseType
type SCIReportAPDetailsReport15trendChart200ResponseType struct {
	Data SCIAPDetailsReport15trendChartData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport15trendChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport15trendChart200ResponseType
}

func newSCIReportAPDetailsReport15trendChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport15trendChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport15trendChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport15trendChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport15trendChart200ResponseType() *SCIReportAPDetailsReport15trendChart200ResponseType {
	m := new(SCIReportAPDetailsReport15trendChart200ResponseType)
	return m
}

// SCIReportAPDetailsReport22trafficTrend200ResponseType
//
// Definition: report_APDetailsReport_22_trafficTrend200ResponseType
type SCIReportAPDetailsReport22trafficTrend200ResponseType struct {
	Data SCIAPDetailsReport22trafficTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport22trafficTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport22trafficTrend200ResponseType
}

func newSCIReportAPDetailsReport22trafficTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport22trafficTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport22trafficTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport22trafficTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport22trafficTrend200ResponseType() *SCIReportAPDetailsReport22trafficTrend200ResponseType {
	m := new(SCIReportAPDetailsReport22trafficTrend200ResponseType)
	return m
}

// SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType
//
// Definition: report_APDetailsReport_40_topSsidsByTrafficTable200ResponseType
type SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType struct {
	Data SCIAPDetailsReport40topSsidsByTrafficTableData `json:"data,omitempty"`

	Metadata *SCIAPDetailsReport40topSsidsByTrafficTableMetaData `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType
}

func newSCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType() *SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType {
	m := new(SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType)
	return m
}

// SCIReportAPDetailsReport75apSummary200ResponseType
//
// Definition: report_APDetailsReport_75_apSummary200ResponseType
type SCIReportAPDetailsReport75apSummary200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport75apSummary200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport75apSummary200ResponseType
}

func newSCIReportAPDetailsReport75apSummary200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport75apSummary200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport75apSummary200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport75apSummary200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport75apSummary200ResponseType() *SCIReportAPDetailsReport75apSummary200ResponseType {
	m := new(SCIReportAPDetailsReport75apSummary200ResponseType)
	return m
}

// SCIReportAPDetailsReport76apPerformance200ResponseType
//
// Definition: report_APDetailsReport_76_apPerformance200ResponseType
type SCIReportAPDetailsReport76apPerformance200ResponseType struct {
	Data SCIAPDetailsReport76apPerformanceData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport76apPerformance200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport76apPerformance200ResponseType
}

func newSCIReportAPDetailsReport76apPerformance200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport76apPerformance200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport76apPerformance200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport76apPerformance200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport76apPerformance200ResponseType() *SCIReportAPDetailsReport76apPerformance200ResponseType {
	m := new(SCIReportAPDetailsReport76apPerformance200ResponseType)
	return m
}

// SCIReportAPDetailsReport77apDetails200ResponseType
//
// Definition: report_APDetailsReport_77_apDetails200ResponseType
type SCIReportAPDetailsReport77apDetails200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport77apDetails200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport77apDetails200ResponseType
}

func newSCIReportAPDetailsReport77apDetails200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport77apDetails200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport77apDetails200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport77apDetails200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport77apDetails200ResponseType() *SCIReportAPDetailsReport77apDetails200ResponseType {
	m := new(SCIReportAPDetailsReport77apDetails200ResponseType)
	return m
}

// SCIReportAPDetailsReport78apStatsOverview200ResponseType
//
// Definition: report_APDetailsReport_78_apStatsOverview200ResponseType
type SCIReportAPDetailsReport78apStatsOverview200ResponseType struct {
	Data SCIAPDetailsReport78apStatsOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport78apStatsOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport78apStatsOverview200ResponseType
}

func newSCIReportAPDetailsReport78apStatsOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport78apStatsOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport78apStatsOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport78apStatsOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport78apStatsOverview200ResponseType() *SCIReportAPDetailsReport78apStatsOverview200ResponseType {
	m := new(SCIReportAPDetailsReport78apStatsOverview200ResponseType)
	return m
}

// SCIReportAPDetailsReport79apUptimeHistory200ResponseType
//
// Definition: report_APDetailsReport_79_apUptimeHistory200ResponseType
type SCIReportAPDetailsReport79apUptimeHistory200ResponseType struct {
	Data SCIAPDetailsReport79apUptimeHistoryData `json:"data,omitempty"`

	Metadata *SCIAPDetailsReport79apUptimeHistoryMetaData `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport79apUptimeHistory200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport79apUptimeHistory200ResponseType
}

func newSCIReportAPDetailsReport79apUptimeHistory200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport79apUptimeHistory200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport79apUptimeHistory200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport79apUptimeHistory200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport79apUptimeHistory200ResponseType() *SCIReportAPDetailsReport79apUptimeHistory200ResponseType {
	m := new(SCIReportAPDetailsReport79apUptimeHistory200ResponseType)
	return m
}

// SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType
//
// Definition: report_APDetailsReport_80_top10ClientsByTrafficVolume200ResponseType
type SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType
}

func newSCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType() *SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType {
	m := new(SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType)
	return m
}

// SCIReportAPDetailsReport81sessionsTable200ResponseType
//
// Definition: report_APDetailsReport_81_sessionsTable200ResponseType
type SCIReportAPDetailsReport81sessionsTable200ResponseType struct {
	Data SCIAPDetailsReport81sessionsTableData `json:"data,omitempty"`

	Metadata *SCIAPDetailsReport81sessionsTableMetaData `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport81sessionsTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport81sessionsTable200ResponseType
}

func newSCIReportAPDetailsReport81sessionsTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport81sessionsTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport81sessionsTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport81sessionsTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport81sessionsTable200ResponseType() *SCIReportAPDetailsReport81sessionsTable200ResponseType {
	m := new(SCIReportAPDetailsReport81sessionsTable200ResponseType)
	return m
}

// SCIReportAPDetailsReport82rssTrend200ResponseType
//
// Definition: report_APDetailsReport_82_rssTrend200ResponseType
type SCIReportAPDetailsReport82rssTrend200ResponseType struct {
	Data SCIAPDetailsReport82rssTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport82rssTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport82rssTrend200ResponseType
}

func newSCIReportAPDetailsReport82rssTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport82rssTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport82rssTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport82rssTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport82rssTrend200ResponseType() *SCIReportAPDetailsReport82rssTrend200ResponseType {
	m := new(SCIReportAPDetailsReport82rssTrend200ResponseType)
	return m
}

// SCIReportAPDetailsReport83snrTrend200ResponseType
//
// Definition: report_APDetailsReport_83_snrTrend200ResponseType
type SCIReportAPDetailsReport83snrTrend200ResponseType struct {
	Data SCIAPDetailsReport83snrTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport83snrTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport83snrTrend200ResponseType
}

func newSCIReportAPDetailsReport83snrTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport83snrTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport83snrTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport83snrTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport83snrTrend200ResponseType() *SCIReportAPDetailsReport83snrTrend200ResponseType {
	m := new(SCIReportAPDetailsReport83snrTrend200ResponseType)
	return m
}

// SCIReportAPDetailsReport84alarmsTable200ResponseType
//
// Definition: report_APDetailsReport_84_alarmsTable200ResponseType
type SCIReportAPDetailsReport84alarmsTable200ResponseType struct {
	Data SCIAPDetailsReport84alarmsTableData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport84alarmsTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport84alarmsTable200ResponseType
}

func newSCIReportAPDetailsReport84alarmsTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport84alarmsTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport84alarmsTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport84alarmsTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport84alarmsTable200ResponseType() *SCIReportAPDetailsReport84alarmsTable200ResponseType {
	m := new(SCIReportAPDetailsReport84alarmsTable200ResponseType)
	return m
}

// SCIReportAPDetailsReport85eventsTable200ResponseType
//
// Definition: report_APDetailsReport_85_eventsTable200ResponseType
type SCIReportAPDetailsReport85eventsTable200ResponseType struct {
	Data SCIAPDetailsReport85eventsTableData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport85eventsTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport85eventsTable200ResponseType
}

func newSCIReportAPDetailsReport85eventsTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport85eventsTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport85eventsTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport85eventsTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport85eventsTable200ResponseType() *SCIReportAPDetailsReport85eventsTable200ResponseType {
	m := new(SCIReportAPDetailsReport85eventsTable200ResponseType)
	return m
}

// SCIReportAPDetailsReport95anomalies200ResponseType
//
// Definition: report_APDetailsReport_95_anomalies200ResponseType
type SCIReportAPDetailsReport95anomalies200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport95anomalies200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport95anomalies200ResponseType
}

func newSCIReportAPDetailsReport95anomalies200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport95anomalies200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport95anomalies200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport95anomalies200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport95anomalies200ResponseType() *SCIReportAPDetailsReport95anomalies200ResponseType {
	m := new(SCIReportAPDetailsReport95anomalies200ResponseType)
	return m
}

// SCIReportAPDetailsReport110apAnomaly200ResponseType
//
// Definition: report_APDetailsReport_110_apAnomaly200ResponseType
type SCIReportAPDetailsReport110apAnomaly200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPDetailsReport110apAnomaly200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPDetailsReport110apAnomaly200ResponseType
}

func newSCIReportAPDetailsReport110apAnomaly200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPDetailsReport110apAnomaly200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPDetailsReport110apAnomaly200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPDetailsReport110apAnomaly200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPDetailsReport110apAnomaly200ResponseType() *SCIReportAPDetailsReport110apAnomaly200ResponseType {
	m := new(SCIReportAPDetailsReport110apAnomaly200ResponseType)
	return m
}

// SCIReportAPsRebootReport43totalReboots200ResponseType
//
// Definition: report_APsRebootReport_43_totalReboots200ResponseType
type SCIReportAPsRebootReport43totalReboots200ResponseType struct {
	Data SCIAPsRebootReport43totalRebootsData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPsRebootReport43totalReboots200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPsRebootReport43totalReboots200ResponseType
}

func newSCIReportAPsRebootReport43totalReboots200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPsRebootReport43totalReboots200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPsRebootReport43totalReboots200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPsRebootReport43totalReboots200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPsRebootReport43totalReboots200ResponseType() *SCIReportAPsRebootReport43totalReboots200ResponseType {
	m := new(SCIReportAPsRebootReport43totalReboots200ResponseType)
	return m
}

// SCIReportAPsRebootReport44topApRebootsTable200ResponseType
//
// Definition: report_APsRebootReport_44_topApRebootsTable200ResponseType
type SCIReportAPsRebootReport44topApRebootsTable200ResponseType struct {
	Data SCIAPsRebootReport44topApRebootsTableData `json:"data,omitempty"`

	Metadata *SCIAPsRebootReport44topApRebootsTableMetaData `json:"metadata,omitempty"`
}

type SCIReportAPsRebootReport44topApRebootsTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPsRebootReport44topApRebootsTable200ResponseType
}

func newSCIReportAPsRebootReport44topApRebootsTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPsRebootReport44topApRebootsTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPsRebootReport44topApRebootsTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPsRebootReport44topApRebootsTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPsRebootReport44topApRebootsTable200ResponseType() *SCIReportAPsRebootReport44topApRebootsTable200ResponseType {
	m := new(SCIReportAPsRebootReport44topApRebootsTable200ResponseType)
	return m
}

// SCIReportAPsRebootReport45topApRebootsOverTime200ResponseType
//
// Definition: report_APsRebootReport_45_topApRebootsOverTime200ResponseType
type SCIReportAPsRebootReport45topApRebootsOverTime200ResponseType struct {
	Data SCIAPsRebootReport45topApRebootsOverTimeData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportAPsRebootReport45topApRebootsOverTime200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportAPsRebootReport45topApRebootsOverTime200ResponseType
}

func newSCIReportAPsRebootReport45topApRebootsOverTime200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportAPsRebootReport45topApRebootsOverTime200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportAPsRebootReport45topApRebootsOverTime200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportAPsRebootReport45topApRebootsOverTime200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportAPsRebootReport45topApRebootsOverTime200ResponseType() *SCIReportAPsRebootReport45topApRebootsOverTime200ResponseType {
	m := new(SCIReportAPsRebootReport45topApRebootsOverTime200ResponseType)
	return m
}

// SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseType
//
// Definition: report_ClientDetailsReport_7_top10ApplicationsByTrafficVolume200ResponseType
type SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseType
}

func newSCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseType() *SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseType {
	m := new(SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseType)
	return m
}

// SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseType
//
// Definition: report_ClientDetailsReport_8_topAppsByTrafficTable200ResponseType
type SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseType struct {
	Data SCIClientDetailsReport8topAppsByTrafficTableData `json:"data,omitempty"`

	Metadata *SCIClientDetailsReport8topAppsByTrafficTableMetaData `json:"metadata,omitempty"`
}

type SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseType
}

func newSCIReportClientDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientDetailsReport8topAppsByTrafficTable200ResponseType() *SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseType {
	m := new(SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseType)
	return m
}

// SCIReportClientDetailsReport82rssTrend200ResponseType
//
// Definition: report_ClientDetailsReport_82_rssTrend200ResponseType
type SCIReportClientDetailsReport82rssTrend200ResponseType struct {
	Data SCIClientDetailsReport82rssTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportClientDetailsReport82rssTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientDetailsReport82rssTrend200ResponseType
}

func newSCIReportClientDetailsReport82rssTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientDetailsReport82rssTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientDetailsReport82rssTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientDetailsReport82rssTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientDetailsReport82rssTrend200ResponseType() *SCIReportClientDetailsReport82rssTrend200ResponseType {
	m := new(SCIReportClientDetailsReport82rssTrend200ResponseType)
	return m
}

// SCIReportClientDetailsReport83snrTrend200ResponseType
//
// Definition: report_ClientDetailsReport_83_snrTrend200ResponseType
type SCIReportClientDetailsReport83snrTrend200ResponseType struct {
	Data SCIClientDetailsReport83snrTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportClientDetailsReport83snrTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientDetailsReport83snrTrend200ResponseType
}

func newSCIReportClientDetailsReport83snrTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientDetailsReport83snrTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientDetailsReport83snrTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientDetailsReport83snrTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientDetailsReport83snrTrend200ResponseType() *SCIReportClientDetailsReport83snrTrend200ResponseType {
	m := new(SCIReportClientDetailsReport83snrTrend200ResponseType)
	return m
}

// SCIReportClientDetailsReport86summary200ResponseType
//
// Definition: report_ClientDetailsReport_86_summary200ResponseType
type SCIReportClientDetailsReport86summary200ResponseType struct {
	Data SCIClientDetailsReport86summaryData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportClientDetailsReport86summary200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientDetailsReport86summary200ResponseType
}

func newSCIReportClientDetailsReport86summary200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientDetailsReport86summary200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientDetailsReport86summary200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientDetailsReport86summary200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientDetailsReport86summary200ResponseType() *SCIReportClientDetailsReport86summary200ResponseType {
	m := new(SCIReportClientDetailsReport86summary200ResponseType)
	return m
}

// SCIReportClientDetailsReport87clientStats200ResponseType
//
// Definition: report_ClientDetailsReport_87_clientStats200ResponseType
type SCIReportClientDetailsReport87clientStats200ResponseType struct {
	Data SCIClientDetailsReport87clientStatsData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportClientDetailsReport87clientStats200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientDetailsReport87clientStats200ResponseType
}

func newSCIReportClientDetailsReport87clientStats200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientDetailsReport87clientStats200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientDetailsReport87clientStats200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientDetailsReport87clientStats200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientDetailsReport87clientStats200ResponseType() *SCIReportClientDetailsReport87clientStats200ResponseType {
	m := new(SCIReportClientDetailsReport87clientStats200ResponseType)
	return m
}

// SCIReportClientDetailsReport89trafficTrend200ResponseType
//
// Definition: report_ClientDetailsReport_89_trafficTrend200ResponseType
type SCIReportClientDetailsReport89trafficTrend200ResponseType struct {
	Data SCIClientDetailsReport89trafficTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportClientDetailsReport89trafficTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientDetailsReport89trafficTrend200ResponseType
}

func newSCIReportClientDetailsReport89trafficTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientDetailsReport89trafficTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientDetailsReport89trafficTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientDetailsReport89trafficTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientDetailsReport89trafficTrend200ResponseType() *SCIReportClientDetailsReport89trafficTrend200ResponseType {
	m := new(SCIReportClientDetailsReport89trafficTrend200ResponseType)
	return m
}

// SCIReportClientDetailsReport92sessionsTable200ResponseType
//
// Definition: report_ClientDetailsReport_92_sessionsTable200ResponseType
type SCIReportClientDetailsReport92sessionsTable200ResponseType struct {
	Data SCIClientDetailsReport92sessionsTableData `json:"data,omitempty"`

	Metadata *SCIClientDetailsReport92sessionsTableMetaData `json:"metadata,omitempty"`
}

type SCIReportClientDetailsReport92sessionsTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientDetailsReport92sessionsTable200ResponseType
}

func newSCIReportClientDetailsReport92sessionsTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientDetailsReport92sessionsTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientDetailsReport92sessionsTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientDetailsReport92sessionsTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientDetailsReport92sessionsTable200ResponseType() *SCIReportClientDetailsReport92sessionsTable200ResponseType {
	m := new(SCIReportClientDetailsReport92sessionsTable200ResponseType)
	return m
}

// SCIReportClientHealthReport144clientHealthSummary200ResponseType
//
// Definition: report_ClientHealthReport_144_clientHealthSummary200ResponseType
type SCIReportClientHealthReport144clientHealthSummary200ResponseType struct {
	Data SCIClientHealthReport144clientHealthSummaryData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportClientHealthReport144clientHealthSummary200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientHealthReport144clientHealthSummary200ResponseType
}

func newSCIReportClientHealthReport144clientHealthSummary200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientHealthReport144clientHealthSummary200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientHealthReport144clientHealthSummary200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientHealthReport144clientHealthSummary200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientHealthReport144clientHealthSummary200ResponseType() *SCIReportClientHealthReport144clientHealthSummary200ResponseType {
	m := new(SCIReportClientHealthReport144clientHealthSummary200ResponseType)
	return m
}

// SCIReportClientHealthReport148clientConnectionHealth200ResponseType
//
// Definition: report_ClientHealthReport_148_clientConnectionHealth200ResponseType
type SCIReportClientHealthReport148clientConnectionHealth200ResponseType struct {
	Data SCIClientHealthReport148clientConnectionHealthData `json:"data,omitempty"`

	Metadata *SCIClientHealthReport148clientConnectionHealthMetaData `json:"metadata,omitempty"`
}

type SCIReportClientHealthReport148clientConnectionHealth200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientHealthReport148clientConnectionHealth200ResponseType
}

func newSCIReportClientHealthReport148clientConnectionHealth200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientHealthReport148clientConnectionHealth200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientHealthReport148clientConnectionHealth200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientHealthReport148clientConnectionHealth200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientHealthReport148clientConnectionHealth200ResponseType() *SCIReportClientHealthReport148clientConnectionHealth200ResponseType {
	m := new(SCIReportClientHealthReport148clientConnectionHealth200ResponseType)
	return m
}

// SCIReportClientHealthReport149clientHealthMetricTrends200ResponseType
//
// Definition: report_ClientHealthReport_149_clientHealthMetricTrends200ResponseType
type SCIReportClientHealthReport149clientHealthMetricTrends200ResponseType struct {
	Data SCIClientHealthReport149clientHealthMetricTrendsData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportClientHealthReport149clientHealthMetricTrends200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientHealthReport149clientHealthMetricTrends200ResponseType
}

func newSCIReportClientHealthReport149clientHealthMetricTrends200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientHealthReport149clientHealthMetricTrends200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientHealthReport149clientHealthMetricTrends200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientHealthReport149clientHealthMetricTrends200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientHealthReport149clientHealthMetricTrends200ResponseType() *SCIReportClientHealthReport149clientHealthMetricTrends200ResponseType {
	m := new(SCIReportClientHealthReport149clientHealthMetricTrends200ResponseType)
	return m
}

// SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType
//
// Definition: report_ClientHealthReport_150_topClientHealthScoreByGroup200ResponseType
type SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType
}

func newSCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType() *SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType {
	m := new(SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType)
	return m
}

// SCIReportComparisonReport140comparisionOverview200ResponseType
//
// Definition: report_ComparisonReport_140_comparisionOverview200ResponseType
type SCIReportComparisonReport140comparisionOverview200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportComparisonReport140comparisionOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportComparisonReport140comparisionOverview200ResponseType
}

func newSCIReportComparisonReport140comparisionOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportComparisonReport140comparisionOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportComparisonReport140comparisionOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportComparisonReport140comparisionOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportComparisonReport140comparisionOverview200ResponseType() *SCIReportComparisonReport140comparisionOverview200ResponseType {
	m := new(SCIReportComparisonReport140comparisionOverview200ResponseType)
	return m
}

// SCIReportComparisonReport145comparisionMetric1200ResponseType
//
// Definition: report_ComparisonReport_145_comparisionMetric1200ResponseType
type SCIReportComparisonReport145comparisionMetric1200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportComparisonReport145comparisionMetric1200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportComparisonReport145comparisionMetric1200ResponseType
}

func newSCIReportComparisonReport145comparisionMetric1200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportComparisonReport145comparisionMetric1200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportComparisonReport145comparisionMetric1200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportComparisonReport145comparisionMetric1200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportComparisonReport145comparisionMetric1200ResponseType() *SCIReportComparisonReport145comparisionMetric1200ResponseType {
	m := new(SCIReportComparisonReport145comparisionMetric1200ResponseType)
	return m
}

// SCIReportComparisonReport146comparisionMetric2200ResponseType
//
// Definition: report_ComparisonReport_146_comparisionMetric2200ResponseType
type SCIReportComparisonReport146comparisionMetric2200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportComparisonReport146comparisionMetric2200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportComparisonReport146comparisionMetric2200ResponseType
}

func newSCIReportComparisonReport146comparisionMetric2200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportComparisonReport146comparisionMetric2200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportComparisonReport146comparisionMetric2200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportComparisonReport146comparisionMetric2200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportComparisonReport146comparisionMetric2200ResponseType() *SCIReportComparisonReport146comparisionMetric2200ResponseType {
	m := new(SCIReportComparisonReport146comparisionMetric2200ResponseType)
	return m
}

// SCIReportComparisonReport147comparisionTable200ResponseType
//
// Definition: report_ComparisonReport_147_comparisionTable200ResponseType
type SCIReportComparisonReport147comparisionTable200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportComparisonReport147comparisionTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportComparisonReport147comparisionTable200ResponseType
}

func newSCIReportComparisonReport147comparisionTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportComparisonReport147comparisionTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportComparisonReport147comparisionTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportComparisonReport147comparisionTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportComparisonReport147comparisionTable200ResponseType() *SCIReportComparisonReport147comparisionTable200ResponseType {
	m := new(SCIReportComparisonReport147comparisionTable200ResponseType)
	return m
}

// SCIReportFind200ResponseType
//
// Definition: report_find200ResponseType
type SCIReportFind200ResponseType []*SCIModelsReport

type SCIReportFind200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIReportFind200ResponseType
}

func newSCIReportFind200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportFind200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportFind200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = make(SCIReportFind200ResponseType, 0)
	return json.NewDecoder(r).Decode(&r.Data)
}
func MakeSCIReportFind200ResponseType() SCIReportFind200ResponseType {
	m := make(SCIReportFind200ResponseType, 0)
	return m
}

// SCIReportGetData200ResponseType
//
// Definition: report_getData200ResponseType
type SCIReportGetData200ResponseType struct {
	Data []interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportGetData200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportGetData200ResponseType
}

func newSCIReportGetData200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportGetData200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportGetData200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportGetData200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportGetData200ResponseType() *SCIReportGetData200ResponseType {
	m := new(SCIReportGetData200ResponseType)
	return m
}

// SCIReportInventoryAPsReport46apInventoryOverview200ResponseType
//
// Definition: report_InventoryAPsReport_46_apInventoryOverview200ResponseType
type SCIReportInventoryAPsReport46apInventoryOverview200ResponseType struct {
	Data SCIInventoryAPsReport46apInventoryOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport46apInventoryOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport46apInventoryOverview200ResponseType
}

func newSCIReportInventoryAPsReport46apInventoryOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport46apInventoryOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport46apInventoryOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport46apInventoryOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport46apInventoryOverview200ResponseType() *SCIReportInventoryAPsReport46apInventoryOverview200ResponseType {
	m := new(SCIReportInventoryAPsReport46apInventoryOverview200ResponseType)
	return m
}

// SCIReportInventoryAPsReport47topApsDisconnection200ResponseType
//
// Definition: report_InventoryAPsReport_47_topApsDisconnection200ResponseType
type SCIReportInventoryAPsReport47topApsDisconnection200ResponseType struct {
	Data SCIInventoryAPsReport47topApsDisconnectionData `json:"data,omitempty"`

	Metadata *SCIInventoryAPsReport47topApsDisconnectionMetaData `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport47topApsDisconnection200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport47topApsDisconnection200ResponseType
}

func newSCIReportInventoryAPsReport47topApsDisconnection200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport47topApsDisconnection200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport47topApsDisconnection200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport47topApsDisconnection200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport47topApsDisconnection200ResponseType() *SCIReportInventoryAPsReport47topApsDisconnection200ResponseType {
	m := new(SCIReportInventoryAPsReport47topApsDisconnection200ResponseType)
	return m
}

// SCIReportInventoryAPsReport48apCountTrend200ResponseType
//
// Definition: report_InventoryAPsReport_48_apCountTrend200ResponseType
type SCIReportInventoryAPsReport48apCountTrend200ResponseType struct {
	Data SCIInventoryAPsReport48apCountTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport48apCountTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport48apCountTrend200ResponseType
}

func newSCIReportInventoryAPsReport48apCountTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport48apCountTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport48apCountTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport48apCountTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport48apCountTrend200ResponseType() *SCIReportInventoryAPsReport48apCountTrend200ResponseType {
	m := new(SCIReportInventoryAPsReport48apCountTrend200ResponseType)
	return m
}

// SCIReportInventoryAPsReport49apStatusTrend200ResponseType
//
// Definition: report_InventoryAPsReport_49_apStatusTrend200ResponseType
type SCIReportInventoryAPsReport49apStatusTrend200ResponseType struct {
	Data SCIInventoryAPsReport49apStatusTrendData `json:"data,omitempty"`

	Metadata *SCIInventoryAPsReport49apStatusTrendMetaData `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport49apStatusTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport49apStatusTrend200ResponseType
}

func newSCIReportInventoryAPsReport49apStatusTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport49apStatusTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport49apStatusTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport49apStatusTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport49apStatusTrend200ResponseType() *SCIReportInventoryAPsReport49apStatusTrend200ResponseType {
	m := new(SCIReportInventoryAPsReport49apStatusTrend200ResponseType)
	return m
}

// SCIReportInventoryAPsReport50topApsModelsChart200ResponseType
//
// Definition: report_InventoryAPsReport_50_topApsModelsChart200ResponseType
type SCIReportInventoryAPsReport50topApsModelsChart200ResponseType struct {
	Data SCIInventoryAPsReport50topApsModelsChartData `json:"data,omitempty"`

	Metadata *SCIInventoryAPsReport50topApsModelsChartMetaData `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport50topApsModelsChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport50topApsModelsChart200ResponseType
}

func newSCIReportInventoryAPsReport50topApsModelsChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport50topApsModelsChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport50topApsModelsChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport50topApsModelsChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport50topApsModelsChart200ResponseType() *SCIReportInventoryAPsReport50topApsModelsChart200ResponseType {
	m := new(SCIReportInventoryAPsReport50topApsModelsChart200ResponseType)
	return m
}

// SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseType
//
// Definition: report_InventoryAPsReport_51_top10ApVersionsChart200ResponseType
type SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseType struct {
	Data SCIInventoryAPsReport51top10ApVersionsChartData `json:"data,omitempty"`

	Metadata *SCIInventoryAPsReport51top10ApVersionsChartMetaData `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseType
}

func newSCIReportInventoryAPsReport51top10ApVersionsChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport51top10ApVersionsChart200ResponseType() *SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseType {
	m := new(SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseType)
	return m
}

// SCIReportInventoryAPsReport52topApsRebootReasons200ResponseType
//
// Definition: report_InventoryAPsReport_52_topApsRebootReasons200ResponseType
type SCIReportInventoryAPsReport52topApsRebootReasons200ResponseType struct {
	Data SCIInventoryAPsReport52topApsRebootReasonsData `json:"data,omitempty"`

	Metadata *SCIInventoryAPsReport52topApsRebootReasonsMetaData `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport52topApsRebootReasons200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport52topApsRebootReasons200ResponseType
}

func newSCIReportInventoryAPsReport52topApsRebootReasons200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport52topApsRebootReasons200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport52topApsRebootReasons200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport52topApsRebootReasons200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport52topApsRebootReasons200ResponseType() *SCIReportInventoryAPsReport52topApsRebootReasons200ResponseType {
	m := new(SCIReportInventoryAPsReport52topApsRebootReasons200ResponseType)
	return m
}

// SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseType
//
// Definition: report_InventoryAPsReport_53_top10ApsRebootCounts200ResponseType
type SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseType struct {
	Data SCIInventoryAPsReport53top10ApsRebootCountsData `json:"data,omitempty"`

	Metadata *SCIInventoryAPsReport53top10ApsRebootCountsMetaData `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseType
}

func newSCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseType() *SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseType {
	m := new(SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseType)
	return m
}

// SCIReportInventoryAPsReport54topApAlarmTypes200ResponseType
//
// Definition: report_InventoryAPsReport_54_topApAlarmTypes200ResponseType
type SCIReportInventoryAPsReport54topApAlarmTypes200ResponseType struct {
	Data SCIInventoryAPsReport54topApAlarmTypesData `json:"data,omitempty"`

	Metadata *SCIInventoryAPsReport54topApAlarmTypesMetaData `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport54topApAlarmTypes200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport54topApAlarmTypes200ResponseType
}

func newSCIReportInventoryAPsReport54topApAlarmTypes200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport54topApAlarmTypes200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport54topApAlarmTypes200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport54topApAlarmTypes200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport54topApAlarmTypes200ResponseType() *SCIReportInventoryAPsReport54topApAlarmTypes200ResponseType {
	m := new(SCIReportInventoryAPsReport54topApAlarmTypes200ResponseType)
	return m
}

// SCIReportInventoryAPsReport55topAPModels200ResponseType
//
// Definition: report_InventoryAPsReport_55_topAPModels200ResponseType
type SCIReportInventoryAPsReport55topAPModels200ResponseType struct {
	Data SCIInventoryAPsReport55topAPModelsData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport55topAPModels200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport55topAPModels200ResponseType
}

func newSCIReportInventoryAPsReport55topAPModels200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport55topAPModels200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport55topAPModels200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport55topAPModels200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport55topAPModels200ResponseType() *SCIReportInventoryAPsReport55topAPModels200ResponseType {
	m := new(SCIReportInventoryAPsReport55topAPModels200ResponseType)
	return m
}

// SCIReportInventoryAPsReport56topAPVersions200ResponseType
//
// Definition: report_InventoryAPsReport_56_topAPVersions200ResponseType
type SCIReportInventoryAPsReport56topAPVersions200ResponseType struct {
	Data SCIInventoryAPsReport56topAPVersionsData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport56topAPVersions200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport56topAPVersions200ResponseType
}

func newSCIReportInventoryAPsReport56topAPVersions200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport56topAPVersions200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport56topAPVersions200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport56topAPVersions200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport56topAPVersions200ResponseType() *SCIReportInventoryAPsReport56topAPVersions200ResponseType {
	m := new(SCIReportInventoryAPsReport56topAPVersions200ResponseType)
	return m
}

// SCIReportInventoryAPsReport57topAPsOffline200ResponseType
//
// Definition: report_InventoryAPsReport_57_topAPsOffline200ResponseType
type SCIReportInventoryAPsReport57topAPsOffline200ResponseType struct {
	Data SCIInventoryAPsReport57topAPsOfflineData `json:"data,omitempty"`

	Metadata *SCIInventoryAPsReport57topAPsOfflineMetaData `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport57topAPsOffline200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport57topAPsOffline200ResponseType
}

func newSCIReportInventoryAPsReport57topAPsOffline200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport57topAPsOffline200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport57topAPsOffline200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport57topAPsOffline200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport57topAPsOffline200ResponseType() *SCIReportInventoryAPsReport57topAPsOffline200ResponseType {
	m := new(SCIReportInventoryAPsReport57topAPsOffline200ResponseType)
	return m
}

// SCIReportInventoryAPsReport58topAPsByReboots200ResponseType
//
// Definition: report_InventoryAPsReport_58_topAPsByReboots200ResponseType
type SCIReportInventoryAPsReport58topAPsByReboots200ResponseType struct {
	Data SCIInventoryAPsReport58topAPsByRebootsData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport58topAPsByReboots200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport58topAPsByReboots200ResponseType
}

func newSCIReportInventoryAPsReport58topAPsByReboots200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport58topAPsByReboots200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport58topAPsByReboots200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport58topAPsByReboots200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport58topAPsByReboots200ResponseType() *SCIReportInventoryAPsReport58topAPsByReboots200ResponseType {
	m := new(SCIReportInventoryAPsReport58topAPsByReboots200ResponseType)
	return m
}

// SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseType
//
// Definition: report_InventoryAPsReport_59_apsConfiguredInMultiCtrl200ResponseType
type SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseType
}

func newSCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseType() *SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseType {
	m := new(SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseType)
	return m
}

// SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseType
//
// Definition: report_InventoryAPsReport_60_apDetailsOnOfflineStatus200ResponseType
type SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseType struct {
	Data SCIInventoryAPsReport60apDetailsOnOfflineStatusData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseType
}

func newSCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseType() *SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseType {
	m := new(SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseType)
	return m
}

// SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseType
//
// Definition: report_InventoryAPsReport_61_apDetailsOtherStatus200ResponseType
type SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseType
}

func newSCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseType() *SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseType {
	m := new(SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseType)
	return m
}

// SCIReportInventoryControllersReport96krack200ResponseType
//
// Definition: report_InventoryControllersReport_96_krack200ResponseType
type SCIReportInventoryControllersReport96krack200ResponseType struct {
	Data SCIInventoryControllersReport96krackData `json:"data,omitempty"`

	Metadata *SCIInventoryControllersReport96krackMetaData `json:"metadata,omitempty"`
}

type SCIReportInventoryControllersReport96krack200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryControllersReport96krack200ResponseType
}

func newSCIReportInventoryControllersReport96krack200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryControllersReport96krack200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryControllersReport96krack200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryControllersReport96krack200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryControllersReport96krack200ResponseType() *SCIReportInventoryControllersReport96krack200ResponseType {
	m := new(SCIReportInventoryControllersReport96krack200ResponseType)
	return m
}

// SCIReportInventoryControllersReport98resourceUtilization200ResponseType
//
// Definition: report_InventoryControllersReport_98_resourceUtilization200ResponseType
type SCIReportInventoryControllersReport98resourceUtilization200ResponseType struct {
	Data SCIInventoryControllersReport98resourceUtilizationData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryControllersReport98resourceUtilization200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryControllersReport98resourceUtilization200ResponseType
}

func newSCIReportInventoryControllersReport98resourceUtilization200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryControllersReport98resourceUtilization200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryControllersReport98resourceUtilization200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryControllersReport98resourceUtilization200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryControllersReport98resourceUtilization200ResponseType() *SCIReportInventoryControllersReport98resourceUtilization200ResponseType {
	m := new(SCIReportInventoryControllersReport98resourceUtilization200ResponseType)
	return m
}

// SCIReportInventoryControllersReport99licenseUtilization200ResponseType
//
// Definition: report_InventoryControllersReport_99_licenseUtilization200ResponseType
type SCIReportInventoryControllersReport99licenseUtilization200ResponseType struct {
	Data SCIInventoryControllersReport99licenseUtilizationData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryControllersReport99licenseUtilization200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryControllersReport99licenseUtilization200ResponseType
}

func newSCIReportInventoryControllersReport99licenseUtilization200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryControllersReport99licenseUtilization200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryControllersReport99licenseUtilization200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryControllersReport99licenseUtilization200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryControllersReport99licenseUtilization200ResponseType() *SCIReportInventoryControllersReport99licenseUtilization200ResponseType {
	m := new(SCIReportInventoryControllersReport99licenseUtilization200ResponseType)
	return m
}

// SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType
//
// Definition: report_InventoryControllersReport_114_controllerInventoryOverview200ResponseType
type SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType struct {
	Data SCIInventoryControllersReport114controllerInventoryOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType
}

func newSCIReportInventoryControllersReport114controllerInventoryOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType() *SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType {
	m := new(SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType)
	return m
}

// SCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseType
//
// Definition: report_InventoryControllersReport_151_licenseUtilizationOverTimeChart200ResponseType
type SCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseType struct {
	Data SCIInventoryControllersReport151licenseUtilizationOverTimeChartData `json:"data,omitempty"`

	Metadata *SCIInventoryControllersReport151licenseUtilizationOverTimeChartMetaData `json:"metadata,omitempty"`
}

type SCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseType
}

func newSCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseType() *SCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseType {
	m := new(SCIReportInventoryControllersReport151licenseUtilizationOverTimeChart200ResponseType)
	return m
}

// SCIReportInventorySwitchesReport113overview200ResponseType
//
// Definition: report_InventorySwitchesReport_113_overview200ResponseType
type SCIReportInventorySwitchesReport113overview200ResponseType struct {
	Data SCIInventorySwitchesReport113overviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventorySwitchesReport113overview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventorySwitchesReport113overview200ResponseType
}

func newSCIReportInventorySwitchesReport113overview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventorySwitchesReport113overview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventorySwitchesReport113overview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventorySwitchesReport113overview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventorySwitchesReport113overview200ResponseType() *SCIReportInventorySwitchesReport113overview200ResponseType {
	m := new(SCIReportInventorySwitchesReport113overview200ResponseType)
	return m
}

// SCIReportInventorySwitchesReport116switchCountTrend200ResponseType
//
// Definition: report_InventorySwitchesReport_116_switchCountTrend200ResponseType
type SCIReportInventorySwitchesReport116switchCountTrend200ResponseType struct {
	Data SCIInventorySwitchesReport116switchCountTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventorySwitchesReport116switchCountTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventorySwitchesReport116switchCountTrend200ResponseType
}

func newSCIReportInventorySwitchesReport116switchCountTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventorySwitchesReport116switchCountTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventorySwitchesReport116switchCountTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventorySwitchesReport116switchCountTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventorySwitchesReport116switchCountTrend200ResponseType() *SCIReportInventorySwitchesReport116switchCountTrend200ResponseType {
	m := new(SCIReportInventorySwitchesReport116switchCountTrend200ResponseType)
	return m
}

// SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseType
//
// Definition: report_InventorySwitchesReport_117_top10SwitchVersionChart200ResponseType
type SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseType struct {
	Data SCIInventorySwitchesReport117top10SwitchVersionChartData `json:"data,omitempty"`

	Metadata *SCIInventorySwitchesReport117top10SwitchVersionChartMetaData `json:"metadata,omitempty"`
}

type SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseType
}

func newSCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseType() *SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseType {
	m := new(SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseType)
	return m
}

// SCIReportInventorySwitchesReport118topSwitchVersions200ResponseType
//
// Definition: report_InventorySwitchesReport_118_topSwitchVersions200ResponseType
type SCIReportInventorySwitchesReport118topSwitchVersions200ResponseType struct {
	Data SCIInventorySwitchesReport118topSwitchVersionsData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventorySwitchesReport118topSwitchVersions200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventorySwitchesReport118topSwitchVersions200ResponseType
}

func newSCIReportInventorySwitchesReport118topSwitchVersions200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventorySwitchesReport118topSwitchVersions200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventorySwitchesReport118topSwitchVersions200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventorySwitchesReport118topSwitchVersions200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventorySwitchesReport118topSwitchVersions200ResponseType() *SCIReportInventorySwitchesReport118topSwitchVersions200ResponseType {
	m := new(SCIReportInventorySwitchesReport118topSwitchVersions200ResponseType)
	return m
}

// SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseType
//
// Definition: report_InventorySwitchesReport_121_topSwitchModelsChart200ResponseType
type SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseType struct {
	Data SCIInventorySwitchesReport121topSwitchModelsChartData `json:"data,omitempty"`

	Metadata *SCIInventorySwitchesReport121topSwitchModelsChartMetaData `json:"metadata,omitempty"`
}

type SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseType
}

func newSCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseType() *SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseType {
	m := new(SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseType)
	return m
}

// SCIReportInventorySwitchesReport122topSwitchModels200ResponseType
//
// Definition: report_InventorySwitchesReport_122_topSwitchModels200ResponseType
type SCIReportInventorySwitchesReport122topSwitchModels200ResponseType struct {
	Data SCIInventorySwitchesReport122topSwitchModelsData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportInventorySwitchesReport122topSwitchModels200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventorySwitchesReport122topSwitchModels200ResponseType
}

func newSCIReportInventorySwitchesReport122topSwitchModels200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventorySwitchesReport122topSwitchModels200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventorySwitchesReport122topSwitchModels200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventorySwitchesReport122topSwitchModels200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventorySwitchesReport122topSwitchModels200ResponseType() *SCIReportInventorySwitchesReport122topSwitchModels200ResponseType {
	m := new(SCIReportInventorySwitchesReport122topSwitchModels200ResponseType)
	return m
}

// SCIReportInventorySwitchesReport132portStatusTrend200ResponseType
//
// Definition: report_InventorySwitchesReport_132_portStatusTrend200ResponseType
type SCIReportInventorySwitchesReport132portStatusTrend200ResponseType struct {
	Data SCIInventorySwitchesReport132portStatusTrendData `json:"data,omitempty"`

	Metadata *SCIInventorySwitchesReport132portStatusTrendMetaData `json:"metadata,omitempty"`
}

type SCIReportInventorySwitchesReport132portStatusTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportInventorySwitchesReport132portStatusTrend200ResponseType
}

func newSCIReportInventorySwitchesReport132portStatusTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportInventorySwitchesReport132portStatusTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportInventorySwitchesReport132portStatusTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportInventorySwitchesReport132portStatusTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportInventorySwitchesReport132portStatusTrend200ResponseType() *SCIReportInventorySwitchesReport132portStatusTrend200ResponseType {
	m := new(SCIReportInventorySwitchesReport132portStatusTrend200ResponseType)
	return m
}

// SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseType
//
// Definition: report_NetworkWiredReport_123_topSwitchPOEUtilChart200ResponseType
type SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseType struct {
	Data SCINetworkWiredReport123topSwitchPOEUtilChartData `json:"data,omitempty"`

	Metadata *SCINetworkWiredReport123topSwitchPOEUtilChartMetaData `json:"metadata,omitempty"`
}

type SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseType
}

func newSCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseType() *SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseType {
	m := new(SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseType)
	return m
}

// SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseType
//
// Definition: report_NetworkWiredReport_124_topSwitchPOEUtils200ResponseType
type SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseType struct {
	Data SCINetworkWiredReport124topSwitchPOEUtilsData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseType
}

func newSCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseType() *SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseType {
	m := new(SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseType)
	return m
}

// SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseType
//
// Definition: report_NetworkWiredReport_127_top10SwitchesByTrafficVolume200ResponseType
type SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseType
}

func newSCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseType() *SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseType {
	m := new(SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseType)
	return m
}

// SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseType
//
// Definition: report_NetworkWiredReport_128_topSwitchesByTrafficTable200ResponseType
type SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseType struct {
	Data SCINetworkWiredReport128topSwitchesByTrafficTableData `json:"data,omitempty"`

	Metadata *SCINetworkWiredReport128topSwitchesByTrafficTableMetaData `json:"metadata,omitempty"`
}

type SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseType
}

func newSCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseType() *SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseType {
	m := new(SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseType)
	return m
}

// SCIReportNetworkWiredReport134wiredOverview200ResponseType
//
// Definition: report_NetworkWiredReport_134_wiredOverview200ResponseType
type SCIReportNetworkWiredReport134wiredOverview200ResponseType struct {
	Data SCINetworkWiredReport134wiredOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportNetworkWiredReport134wiredOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWiredReport134wiredOverview200ResponseType
}

func newSCIReportNetworkWiredReport134wiredOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWiredReport134wiredOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWiredReport134wiredOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWiredReport134wiredOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWiredReport134wiredOverview200ResponseType() *SCIReportNetworkWiredReport134wiredOverview200ResponseType {
	m := new(SCIReportNetworkWiredReport134wiredOverview200ResponseType)
	return m
}

// SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseType
//
// Definition: report_NetworkWiredReport_135_wiredTrafficDistribution200ResponseType
type SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseType struct {
	Data SCINetworkWiredReport135wiredTrafficDistributionData `json:"data,omitempty"`

	Metadata *SCINetworkWiredReport135wiredTrafficDistributionMetaData `json:"metadata,omitempty"`
}

type SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseType
}

func newSCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseType() *SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseType {
	m := new(SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseType)
	return m
}

// SCIReportNetworkWiredReport136switchTrafficTrend200ResponseType
//
// Definition: report_NetworkWiredReport_136_switchTrafficTrend200ResponseType
type SCIReportNetworkWiredReport136switchTrafficTrend200ResponseType struct {
	Data SCINetworkWiredReport136switchTrafficTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportNetworkWiredReport136switchTrafficTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWiredReport136switchTrafficTrend200ResponseType
}

func newSCIReportNetworkWiredReport136switchTrafficTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWiredReport136switchTrafficTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWiredReport136switchTrafficTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWiredReport136switchTrafficTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWiredReport136switchTrafficTrend200ResponseType() *SCIReportNetworkWiredReport136switchTrafficTrend200ResponseType {
	m := new(SCIReportNetworkWiredReport136switchTrafficTrend200ResponseType)
	return m
}

// SCIReportNetworkWiredReport141switchErrorTrend200ResponseType
//
// Definition: report_NetworkWiredReport_141_switchErrorTrend200ResponseType
type SCIReportNetworkWiredReport141switchErrorTrend200ResponseType struct {
	Data SCINetworkWiredReport141switchErrorTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportNetworkWiredReport141switchErrorTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWiredReport141switchErrorTrend200ResponseType
}

func newSCIReportNetworkWiredReport141switchErrorTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWiredReport141switchErrorTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWiredReport141switchErrorTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWiredReport141switchErrorTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWiredReport141switchErrorTrend200ResponseType() *SCIReportNetworkWiredReport141switchErrorTrend200ResponseType {
	m := new(SCIReportNetworkWiredReport141switchErrorTrend200ResponseType)
	return m
}

// SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseType
//
// Definition: report_NetworkWiredReport_142_topSwitchesByErrorsChart200ResponseType
type SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseType struct {
	Data SCINetworkWiredReport142topSwitchesByErrorsChartData `json:"data,omitempty"`

	Metadata *SCINetworkWiredReport142topSwitchesByErrorsChartMetaData `json:"metadata,omitempty"`
}

type SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseType
}

func newSCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseType() *SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseType {
	m := new(SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseType)
	return m
}

// SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseType
//
// Definition: report_NetworkWiredReport_143_topSwitchesByErrorsTable200ResponseType
type SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseType struct {
	Data SCINetworkWiredReport143topSwitchesByErrorsTableData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseType
}

func newSCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseType() *SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseType {
	m := new(SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseType)
	return m
}

// SCIReportNetworkWirelessReport20overview200ResponseType
//
// Definition: report_NetworkWirelessReport_20_overview200ResponseType
type SCIReportNetworkWirelessReport20overview200ResponseType struct {
	Data SCINetworkWirelessReport20overviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportNetworkWirelessReport20overview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWirelessReport20overview200ResponseType
}

func newSCIReportNetworkWirelessReport20overview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWirelessReport20overview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWirelessReport20overview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWirelessReport20overview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWirelessReport20overview200ResponseType() *SCIReportNetworkWirelessReport20overview200ResponseType {
	m := new(SCIReportNetworkWirelessReport20overview200ResponseType)
	return m
}

// SCIReportNetworkWirelessReport21trafficDistribution200ResponseType
//
// Definition: report_NetworkWirelessReport_21_trafficDistribution200ResponseType
type SCIReportNetworkWirelessReport21trafficDistribution200ResponseType struct {
	Data SCINetworkWirelessReport21trafficDistributionData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportNetworkWirelessReport21trafficDistribution200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWirelessReport21trafficDistribution200ResponseType
}

func newSCIReportNetworkWirelessReport21trafficDistribution200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWirelessReport21trafficDistribution200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWirelessReport21trafficDistribution200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWirelessReport21trafficDistribution200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWirelessReport21trafficDistribution200ResponseType() *SCIReportNetworkWirelessReport21trafficDistribution200ResponseType {
	m := new(SCIReportNetworkWirelessReport21trafficDistribution200ResponseType)
	return m
}

// SCIReportNetworkWirelessReport22trafficTrend200ResponseType
//
// Definition: report_NetworkWirelessReport_22_trafficTrend200ResponseType
type SCIReportNetworkWirelessReport22trafficTrend200ResponseType struct {
	Data SCINetworkWirelessReport22trafficTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportNetworkWirelessReport22trafficTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWirelessReport22trafficTrend200ResponseType
}

func newSCIReportNetworkWirelessReport22trafficTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWirelessReport22trafficTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWirelessReport22trafficTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWirelessReport22trafficTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWirelessReport22trafficTrend200ResponseType() *SCIReportNetworkWirelessReport22trafficTrend200ResponseType {
	m := new(SCIReportNetworkWirelessReport22trafficTrend200ResponseType)
	return m
}

// SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseType
//
// Definition: report_NetworkWirelessReport_23_trafficOverTimeTable200ResponseType
type SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseType struct {
	Data SCINetworkWirelessReport23trafficOverTimeTableData `json:"data,omitempty"`

	Metadata *SCINetworkWirelessReport23trafficOverTimeTableMetaData `json:"metadata,omitempty"`
}

type SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseType
}

func newSCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseType() *SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseType {
	m := new(SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseType)
	return m
}

// SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseType
//
// Definition: report_NetworkWirelessReport_24_topAPsByTrafficTable200ResponseType
type SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseType struct {
	Data SCINetworkWirelessReport24topAPsByTrafficTableData `json:"data,omitempty"`

	Metadata *SCINetworkWirelessReport24topAPsByTrafficTableMetaData `json:"metadata,omitempty"`
}

type SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseType
}

func newSCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseType() *SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseType {
	m := new(SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseType)
	return m
}

// SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseType
//
// Definition: report_NetworkWirelessReport_25_topAPsByClientsTable200ResponseType
type SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseType struct {
	Data SCINetworkWirelessReport25topAPsByClientsTableData `json:"data,omitempty"`

	Metadata *SCINetworkWirelessReport25topAPsByClientsTableMetaData `json:"metadata,omitempty"`
}

type SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseType
}

func newSCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseType() *SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseType {
	m := new(SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseType)
	return m
}

// SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseType
//
// Definition: report_NetworkWirelessReport_26_top10APsByTrafficVolume200ResponseType
type SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseType
}

func newSCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseType() *SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseType {
	m := new(SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseType)
	return m
}

// SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType
//
// Definition: report_NetworkWirelessReport_27_top10ApByClientCount200ResponseType
type SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType struct {
	Data SCINetworkWirelessReport27top10ApByClientCountData `json:"data,omitempty"`

	Metadata *SCINetworkWirelessReport27top10ApByClientCountMetaData `json:"metadata,omitempty"`
}

type SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType
}

func newSCIReportNetworkWirelessReport27top10ApByClientCount200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType() *SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType {
	m := new(SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType)
	return m
}

// SCIReportOverview62overview200ResponseType
//
// Definition: report_Overview_62_overview200ResponseType
type SCIReportOverview62overview200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata *SCIOverview62overviewMetaData `json:"metadata,omitempty"`
}

type SCIReportOverview62overview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview62overview200ResponseType
}

func newSCIReportOverview62overview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview62overview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview62overview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview62overview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview62overview200ResponseType() *SCIReportOverview62overview200ResponseType {
	m := new(SCIReportOverview62overview200ResponseType)
	return m
}

// SCIReportOverview63controller200ResponseType
//
// Definition: report_Overview_63_controller200ResponseType
type SCIReportOverview63controller200ResponseType struct {
	Data SCIOverview63controllerData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportOverview63controller200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview63controller200ResponseType
}

func newSCIReportOverview63controller200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview63controller200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview63controller200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview63controller200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview63controller200ResponseType() *SCIReportOverview63controller200ResponseType {
	m := new(SCIReportOverview63controller200ResponseType)
	return m
}

// SCIReportOverview64apOverview200ResponseType
//
// Definition: report_Overview_64_apOverview200ResponseType
type SCIReportOverview64apOverview200ResponseType struct {
	Data SCIOverview64apOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportOverview64apOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview64apOverview200ResponseType
}

func newSCIReportOverview64apOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview64apOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview64apOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview64apOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview64apOverview200ResponseType() *SCIReportOverview64apOverview200ResponseType {
	m := new(SCIReportOverview64apOverview200ResponseType)
	return m
}

// SCIReportOverview66apAlarmOverview200ResponseType
//
// Definition: report_Overview_66_apAlarmOverview200ResponseType
type SCIReportOverview66apAlarmOverview200ResponseType struct {
	Data SCIOverview66apAlarmOverviewData `json:"data,omitempty"`

	Metadata *SCIOverview66apAlarmOverviewMetaData `json:"metadata,omitempty"`
}

type SCIReportOverview66apAlarmOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview66apAlarmOverview200ResponseType
}

func newSCIReportOverview66apAlarmOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview66apAlarmOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview66apAlarmOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview66apAlarmOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview66apAlarmOverview200ResponseType() *SCIReportOverview66apAlarmOverview200ResponseType {
	m := new(SCIReportOverview66apAlarmOverview200ResponseType)
	return m
}

// SCIReportOverview67switchOverview200ResponseType
//
// Definition: report_Overview_67_switchOverview200ResponseType
type SCIReportOverview67switchOverview200ResponseType struct {
	Data SCIOverview67switchOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportOverview67switchOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview67switchOverview200ResponseType
}

func newSCIReportOverview67switchOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview67switchOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview67switchOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview67switchOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview67switchOverview200ResponseType() *SCIReportOverview67switchOverview200ResponseType {
	m := new(SCIReportOverview67switchOverview200ResponseType)
	return m
}

// SCIReportOverview68apClientCountOverview200ResponseType
//
// Definition: report_Overview_68_apClientCountOverview200ResponseType
type SCIReportOverview68apClientCountOverview200ResponseType struct {
	Data SCIOverview68apClientCountOverviewData `json:"data,omitempty"`

	Metadata *SCIOverview68apClientCountOverviewMetaData `json:"metadata,omitempty"`
}

type SCIReportOverview68apClientCountOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview68apClientCountOverview200ResponseType
}

func newSCIReportOverview68apClientCountOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview68apClientCountOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview68apClientCountOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview68apClientCountOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview68apClientCountOverview200ResponseType() *SCIReportOverview68apClientCountOverview200ResponseType {
	m := new(SCIReportOverview68apClientCountOverview200ResponseType)
	return m
}

// SCIReportOverview69totalTrafficMinMaxRate200ResponseType
//
// Definition: report_Overview_69_totalTrafficMinMaxRate200ResponseType
type SCIReportOverview69totalTrafficMinMaxRate200ResponseType struct {
	Data SCIOverview69totalTrafficMinMaxRateData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportOverview69totalTrafficMinMaxRate200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview69totalTrafficMinMaxRate200ResponseType
}

func newSCIReportOverview69totalTrafficMinMaxRate200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview69totalTrafficMinMaxRate200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview69totalTrafficMinMaxRate200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview69totalTrafficMinMaxRate200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview69totalTrafficMinMaxRate200ResponseType() *SCIReportOverview69totalTrafficMinMaxRate200ResponseType {
	m := new(SCIReportOverview69totalTrafficMinMaxRate200ResponseType)
	return m
}

// SCIReportOverview70sessionsOverview200ResponseType
//
// Definition: report_Overview_70_sessionsOverview200ResponseType
type SCIReportOverview70sessionsOverview200ResponseType struct {
	Data SCIOverview70sessionsOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportOverview70sessionsOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview70sessionsOverview200ResponseType
}

func newSCIReportOverview70sessionsOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview70sessionsOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview70sessionsOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview70sessionsOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview70sessionsOverview200ResponseType() *SCIReportOverview70sessionsOverview200ResponseType {
	m := new(SCIReportOverview70sessionsOverview200ResponseType)
	return m
}

// SCIReportOverview71ssidOverview200ResponseType
//
// Definition: report_Overview_71_ssidOverview200ResponseType
type SCIReportOverview71ssidOverview200ResponseType struct {
	Data SCIOverview71ssidOverviewData `json:"data,omitempty"`

	Metadata *SCIOverview71ssidOverviewMetaData `json:"metadata,omitempty"`
}

type SCIReportOverview71ssidOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview71ssidOverview200ResponseType
}

func newSCIReportOverview71ssidOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview71ssidOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview71ssidOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview71ssidOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview71ssidOverview200ResponseType() *SCIReportOverview71ssidOverview200ResponseType {
	m := new(SCIReportOverview71ssidOverview200ResponseType)
	return m
}

// SCIReportOverview72radioOverview200ResponseType
//
// Definition: report_Overview_72_radioOverview200ResponseType
type SCIReportOverview72radioOverview200ResponseType struct {
	Data SCIOverview72radioOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportOverview72radioOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview72radioOverview200ResponseType
}

func newSCIReportOverview72radioOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview72radioOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview72radioOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview72radioOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview72radioOverview200ResponseType() *SCIReportOverview72radioOverview200ResponseType {
	m := new(SCIReportOverview72radioOverview200ResponseType)
	return m
}

// SCIReportOverview73applicationsOverview200ResponseType
//
// Definition: report_Overview_73_applicationsOverview200ResponseType
type SCIReportOverview73applicationsOverview200ResponseType struct {
	Data SCIOverview73applicationsOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportOverview73applicationsOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview73applicationsOverview200ResponseType
}

func newSCIReportOverview73applicationsOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview73applicationsOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview73applicationsOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview73applicationsOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview73applicationsOverview200ResponseType() *SCIReportOverview73applicationsOverview200ResponseType {
	m := new(SCIReportOverview73applicationsOverview200ResponseType)
	return m
}

// SCIReportOverview74apEventOverview200ResponseType
//
// Definition: report_Overview_74_apEventOverview200ResponseType
type SCIReportOverview74apEventOverview200ResponseType struct {
	Data SCIOverview74apEventOverviewData `json:"data,omitempty"`

	Metadata *SCIOverview74apEventOverviewMetaData `json:"metadata,omitempty"`
}

type SCIReportOverview74apEventOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview74apEventOverview200ResponseType
}

func newSCIReportOverview74apEventOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview74apEventOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview74apEventOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview74apEventOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview74apEventOverview200ResponseType() *SCIReportOverview74apEventOverview200ResponseType {
	m := new(SCIReportOverview74apEventOverview200ResponseType)
	return m
}

// SCIReportOverview97factOverview200ResponseType
//
// Definition: report_Overview_97_factOverview200ResponseType
type SCIReportOverview97factOverview200ResponseType struct {
	Data SCIOverview97factOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportOverview97factOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview97factOverview200ResponseType
}

func newSCIReportOverview97factOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview97factOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview97factOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview97factOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview97factOverview200ResponseType() *SCIReportOverview97factOverview200ResponseType {
	m := new(SCIReportOverview97factOverview200ResponseType)
	return m
}

// SCIReportOverview115networkUsageOverview200ResponseType
//
// Definition: report_Overview_115_networkUsageOverview200ResponseType
type SCIReportOverview115networkUsageOverview200ResponseType struct {
	Data *SCIOverview115networkUsageOverviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportOverview115networkUsageOverview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportOverview115networkUsageOverview200ResponseType
}

func newSCIReportOverview115networkUsageOverview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportOverview115networkUsageOverview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportOverview115networkUsageOverview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportOverview115networkUsageOverview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportOverview115networkUsageOverview200ResponseType() *SCIReportOverview115networkUsageOverview200ResponseType {
	m := new(SCIReportOverview115networkUsageOverview200ResponseType)
	return m
}

// SCIReportPrototypegetsections200ResponseType
//
// Definition: report_prototype_get_sections200ResponseType
type SCIReportPrototypegetsections200ResponseType []*SCIModelsSection

type SCIReportPrototypegetsections200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIReportPrototypegetsections200ResponseType
}

func newSCIReportPrototypegetsections200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportPrototypegetsections200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportPrototypegetsections200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = make(SCIReportPrototypegetsections200ResponseType, 0)
	return json.NewDecoder(r).Decode(&r.Data)
}
func MakeSCIReportPrototypegetsections200ResponseType() SCIReportPrototypegetsections200ResponseType {
	m := make(SCIReportPrototypegetsections200ResponseType, 0)
	return m
}

// SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType
//
// Definition: report_SCNetworkTrafficReport_93_scNetworkTraffic200ResponseType
type SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType
}

func newSCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType() *SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType
//
// Definition: report_SCNetworkTrafficReport_94_scNetworkTrend200ResponseType
type SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType
}

func newSCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType() *SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType
//
// Definition: report_SCNetworkTrafficReport_100_droppedCallRate200ResponseType
type SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType
}

func newSCIReportSCNetworkTrafficReport100droppedCallRate200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType() *SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType
//
// Definition: report_SCNetworkTrafficReport_101_connectionSetupSuccessRate200ResponseType
type SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType
}

func newSCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType() *SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType
//
// Definition: report_SCNetworkTrafficReport_102_handoverSuccessRate200ResponseType
type SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType
}

func newSCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType() *SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport103avgThroughput200ResponseType
//
// Definition: report_SCNetworkTrafficReport_103_avgThroughput200ResponseType
type SCIReportSCNetworkTrafficReport103avgThroughput200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport103avgThroughput200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport103avgThroughput200ResponseType
}

func newSCIReportSCNetworkTrafficReport103avgThroughput200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport103avgThroughput200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport103avgThroughput200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport103avgThroughput200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport103avgThroughput200ResponseType() *SCIReportSCNetworkTrafficReport103avgThroughput200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport103avgThroughput200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport104scAvailability200ResponseType
//
// Definition: report_SCNetworkTrafficReport_104_scAvailability200ResponseType
type SCIReportSCNetworkTrafficReport104scAvailability200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport104scAvailability200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport104scAvailability200ResponseType
}

func newSCIReportSCNetworkTrafficReport104scAvailability200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport104scAvailability200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport104scAvailability200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport104scAvailability200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport104scAvailability200ResponseType() *SCIReportSCNetworkTrafficReport104scAvailability200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport104scAvailability200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType
//
// Definition: report_SCNetworkTrafficReport_105_rscConnectionStats200ResponseType
type SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType
}

func newSCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType() *SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType
//
// Definition: report_SCNetworkTrafficReport_106_rscGpsStats200ResponseType
type SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType
}

func newSCIReportSCNetworkTrafficReport106rscGpsStats200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType() *SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport107trafficVolume200ResponseType
//
// Definition: report_SCNetworkTrafficReport_107_trafficVolume200ResponseType
type SCIReportSCNetworkTrafficReport107trafficVolume200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport107trafficVolume200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport107trafficVolume200ResponseType
}

func newSCIReportSCNetworkTrafficReport107trafficVolume200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport107trafficVolume200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport107trafficVolume200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport107trafficVolume200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport107trafficVolume200ResponseType() *SCIReportSCNetworkTrafficReport107trafficVolume200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport107trafficVolume200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType
//
// Definition: report_SCNetworkTrafficReport_108_phaseSyncLoss200ResponseType
type SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType
}

func newSCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType() *SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType
//
// Definition: report_SCNetworkTrafficReport_109_frequencySyncLoss200ResponseType
type SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType
}

func newSCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType() *SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType)
	return m
}

// SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType
//
// Definition: report_SCNetworkTrafficReport_111_rscTrafficStats200ResponseType
type SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType
}

func newSCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType() *SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType {
	m := new(SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType)
	return m
}

// SCIReportSearchQuery
//
// Definition: report_searchQuery
//
// Search query
type SCIReportSearchQuery struct {
	CaseSensitive *bool `json:"caseSensitive,omitempty"`

	// Type
	// Constraints:
	//    - required
	Type *string `json:"type"`

	Value *string `json:"value,omitempty"`

	Values []string `json:"values,omitempty"`
}

func NewSCIReportSearchQuery() *SCIReportSearchQuery {
	m := new(SCIReportSearchQuery)
	return m
}

// SCIReportSessionsSummaryReport33topTable200ResponseType
//
// Definition: report_SessionsSummaryReport_33_topTable200ResponseType
type SCIReportSessionsSummaryReport33topTable200ResponseType struct {
	Data SCISessionsSummaryReport33topTableData `json:"data,omitempty"`

	Metadata *SCISessionsSummaryReport33topTableMetaData `json:"metadata,omitempty"`
}

type SCIReportSessionsSummaryReport33topTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSessionsSummaryReport33topTable200ResponseType
}

func newSCIReportSessionsSummaryReport33topTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSessionsSummaryReport33topTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSessionsSummaryReport33topTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSessionsSummaryReport33topTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSessionsSummaryReport33topTable200ResponseType() *SCIReportSessionsSummaryReport33topTable200ResponseType {
	m := new(SCIReportSessionsSummaryReport33topTable200ResponseType)
	return m
}

// SCIReportSessionsSummaryReport34overview200ResponseType
//
// Definition: report_SessionsSummaryReport_34_overview200ResponseType
type SCIReportSessionsSummaryReport34overview200ResponseType struct {
	Data SCISessionsSummaryReport34overviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSessionsSummaryReport34overview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSessionsSummaryReport34overview200ResponseType
}

func newSCIReportSessionsSummaryReport34overview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSessionsSummaryReport34overview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSessionsSummaryReport34overview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSessionsSummaryReport34overview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSessionsSummaryReport34overview200ResponseType() *SCIReportSessionsSummaryReport34overview200ResponseType {
	m := new(SCIReportSessionsSummaryReport34overview200ResponseType)
	return m
}

// SCIReportSessionsSummaryReport42durationPercentile200ResponseType
//
// Definition: report_SessionsSummaryReport_42_durationPercentile200ResponseType
type SCIReportSessionsSummaryReport42durationPercentile200ResponseType struct {
	Data SCISessionsSummaryReport42durationPercentileData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSessionsSummaryReport42durationPercentile200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSessionsSummaryReport42durationPercentile200ResponseType
}

func newSCIReportSessionsSummaryReport42durationPercentile200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSessionsSummaryReport42durationPercentile200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSessionsSummaryReport42durationPercentile200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSessionsSummaryReport42durationPercentile200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSessionsSummaryReport42durationPercentile200ResponseType() *SCIReportSessionsSummaryReport42durationPercentile200ResponseType {
	m := new(SCIReportSessionsSummaryReport42durationPercentile200ResponseType)
	return m
}

// SCIReportSwitchDetailsReport125switchSummary200ResponseType
//
// Definition: report_SwitchDetailsReport_125_switchSummary200ResponseType
type SCIReportSwitchDetailsReport125switchSummary200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSwitchDetailsReport125switchSummary200ResponseType
}

func newSCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSwitchDetailsReport125switchSummary200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSwitchDetailsReport125switchSummary200ResponseType() *SCIReportSwitchDetailsReport125switchSummary200ResponseType {
	m := new(SCIReportSwitchDetailsReport125switchSummary200ResponseType)
	return m
}

// SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseType
//
// Definition: report_SwitchDetailsReport_126_switchResourceUtilization200ResponseType
type SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseType struct {
	Data SCISwitchDetailsReport126switchResourceUtilizationData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseType
}

func newSCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSwitchDetailsReport126switchResourceUtilization200ResponseType() *SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseType {
	m := new(SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseType)
	return m
}

// SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseType
//
// Definition: report_SwitchDetailsReport_129_topSwitchPortsByTrafficChart200ResponseType
type SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseType struct {
	Data SCISwitchDetailsReport129topSwitchPortsByTrafficChartData `json:"data,omitempty"`

	Metadata *SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData `json:"metadata,omitempty"`
}

type SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseType
}

func newSCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseType() *SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseType {
	m := new(SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseType)
	return m
}

// SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseType
//
// Definition: report_SwitchDetailsReport_130_topSwitchPortsByTrafficTable200ResponseType
type SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseType struct {
	Data SCISwitchDetailsReport130topSwitchPortsByTrafficTableData `json:"data,omitempty"`

	Metadata *SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData `json:"metadata,omitempty"`
}

type SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseType
}

func newSCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseType() *SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseType {
	m := new(SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseType)
	return m
}

// SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseType
//
// Definition: report_SwitchDetailsReport_131_switchTrafficTrend200ResponseType
type SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseType
}

func newSCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSwitchDetailsReport131switchTrafficTrend200ResponseType() *SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseType {
	m := new(SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseType)
	return m
}

// SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseType
//
// Definition: report_SwitchDetailsReport_137_lldpNeighborTable200ResponseType
type SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseType
}

func newSCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSwitchDetailsReport137lldpNeighborTable200ResponseType() *SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseType {
	m := new(SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseType)
	return m
}

// SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseType
//
// Definition: report_SwitchDetailsReport_138_switchUptimeHistory200ResponseType
type SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseType struct {
	Data SCISwitchDetailsReport138switchUptimeHistoryData `json:"data,omitempty"`

	Metadata *SCISwitchDetailsReport138switchUptimeHistoryMetaData `json:"metadata,omitempty"`
}

type SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseType
}

func newSCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSwitchDetailsReport138switchUptimeHistory200ResponseType() *SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseType {
	m := new(SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseType)
	return m
}

// SCIReportSwitchDetailsReport139switchDetails200ResponseType
//
// Definition: report_SwitchDetailsReport_139_switchDetails200ResponseType
type SCIReportSwitchDetailsReport139switchDetails200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSwitchDetailsReport139switchDetails200ResponseType
}

func newSCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSwitchDetailsReport139switchDetails200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSwitchDetailsReport139switchDetails200ResponseType() *SCIReportSwitchDetailsReport139switchDetails200ResponseType {
	m := new(SCIReportSwitchDetailsReport139switchDetails200ResponseType)
	return m
}

// SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseType
//
// Definition: report_SwitchDetailsReport_152_perSwitchErrorTrend200ResponseType
type SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseType struct {
	Data SCISwitchDetailsReport152perSwitchErrorTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseType
}

func newSCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseType() *SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseType {
	m := new(SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseType)
	return m
}

// SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType
//
// Definition: report_WirelessApplicationsReport_7_top10ApplicationsByTrafficVolume200ResponseType
type SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType
}

func newSCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType() *SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType {
	m := new(SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType)
	return m
}

// SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType
//
// Definition: report_WirelessApplicationsReport_8_topAppsByTrafficTable200ResponseType
type SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType struct {
	Data SCIWirelessApplicationsReport8topAppsByTrafficTableData `json:"data,omitempty"`

	Metadata *SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData `json:"metadata,omitempty"`
}

type SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType
}

func newSCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType() *SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType {
	m := new(SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType)
	return m
}

// SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType
//
// Definition: report_WirelessApplicationsReport_9_topAppsByClientsTable200ResponseType
type SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType struct {
	Data SCIWirelessApplicationsReport9topAppsByClientsTableData `json:"data,omitempty"`

	Metadata *SCIWirelessApplicationsReport9topAppsByClientsTableMetaData `json:"metadata,omitempty"`
}

type SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType
}

func newSCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType() *SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType {
	m := new(SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType)
	return m
}

// SCIReportWirelessApplicationsReport10overview200ResponseType
//
// Definition: report_WirelessApplicationsReport_10_overview200ResponseType
type SCIReportWirelessApplicationsReport10overview200ResponseType struct {
	Data SCIWirelessApplicationsReport10overviewData `json:"data,omitempty"`

	Metadata *SCIWirelessApplicationsReport10overviewMetaData `json:"metadata,omitempty"`
}

type SCIReportWirelessApplicationsReport10overview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessApplicationsReport10overview200ResponseType
}

func newSCIReportWirelessApplicationsReport10overview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessApplicationsReport10overview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessApplicationsReport10overview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessApplicationsReport10overview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessApplicationsReport10overview200ResponseType() *SCIReportWirelessApplicationsReport10overview200ResponseType {
	m := new(SCIReportWirelessApplicationsReport10overview200ResponseType)
	return m
}

// SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType
//
// Definition: report_WirelessApplicationsReport_11_top10ApplicationsByClientCount200ResponseType
type SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType struct {
	Data SCIWirelessApplicationsReport11top10ApplicationsByClientCountData `json:"data,omitempty"`

	Metadata *SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData `json:"metadata,omitempty"`
}

type SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType
}

func newSCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType() *SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType {
	m := new(SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType)
	return m
}

// SCIReportWirelessClientsReport12overview200ResponseType
//
// Definition: report_WirelessClientsReport_12_overview200ResponseType
type SCIReportWirelessClientsReport12overview200ResponseType struct {
	Data SCIWirelessClientsReport12overviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportWirelessClientsReport12overview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessClientsReport12overview200ResponseType
}

func newSCIReportWirelessClientsReport12overview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessClientsReport12overview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessClientsReport12overview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessClientsReport12overview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessClientsReport12overview200ResponseType() *SCIReportWirelessClientsReport12overview200ResponseType {
	m := new(SCIReportWirelessClientsReport12overview200ResponseType)
	return m
}

// SCIReportWirelessClientsReport13topChart200ResponseType
//
// Definition: report_WirelessClientsReport_13_topChart200ResponseType
type SCIReportWirelessClientsReport13topChart200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportWirelessClientsReport13topChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessClientsReport13topChart200ResponseType
}

func newSCIReportWirelessClientsReport13topChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessClientsReport13topChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessClientsReport13topChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessClientsReport13topChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessClientsReport13topChart200ResponseType() *SCIReportWirelessClientsReport13topChart200ResponseType {
	m := new(SCIReportWirelessClientsReport13topChart200ResponseType)
	return m
}

// SCIReportWirelessClientsReport14topTable200ResponseType
//
// Definition: report_WirelessClientsReport_14_topTable200ResponseType
type SCIReportWirelessClientsReport14topTable200ResponseType struct {
	Data SCIWirelessClientsReport14topTableData `json:"data,omitempty"`

	Metadata *SCIWirelessClientsReport14topTableMetaData `json:"metadata,omitempty"`
}

type SCIReportWirelessClientsReport14topTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessClientsReport14topTable200ResponseType
}

func newSCIReportWirelessClientsReport14topTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessClientsReport14topTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessClientsReport14topTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessClientsReport14topTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessClientsReport14topTable200ResponseType() *SCIReportWirelessClientsReport14topTable200ResponseType {
	m := new(SCIReportWirelessClientsReport14topTable200ResponseType)
	return m
}

// SCIReportWirelessClientsReport15trendChart200ResponseType
//
// Definition: report_WirelessClientsReport_15_trendChart200ResponseType
type SCIReportWirelessClientsReport15trendChart200ResponseType struct {
	Data SCIWirelessClientsReport15trendChartData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportWirelessClientsReport15trendChart200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessClientsReport15trendChart200ResponseType
}

func newSCIReportWirelessClientsReport15trendChart200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessClientsReport15trendChart200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessClientsReport15trendChart200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessClientsReport15trendChart200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessClientsReport15trendChart200ResponseType() *SCIReportWirelessClientsReport15trendChart200ResponseType {
	m := new(SCIReportWirelessClientsReport15trendChart200ResponseType)
	return m
}

// SCIReportWirelessClientsReport16trendTable200ResponseType
//
// Definition: report_WirelessClientsReport_16_trendTable200ResponseType
type SCIReportWirelessClientsReport16trendTable200ResponseType struct {
	Data SCIWirelessClientsReport16trendTableData `json:"data,omitempty"`

	Metadata *SCIWirelessClientsReport16trendTableMetaData `json:"metadata,omitempty"`
}

type SCIReportWirelessClientsReport16trendTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessClientsReport16trendTable200ResponseType
}

func newSCIReportWirelessClientsReport16trendTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessClientsReport16trendTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessClientsReport16trendTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessClientsReport16trendTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessClientsReport16trendTable200ResponseType() *SCIReportWirelessClientsReport16trendTable200ResponseType {
	m := new(SCIReportWirelessClientsReport16trendTable200ResponseType)
	return m
}

// SCIReportWirelessClientsReport17topPercentile200ResponseType
//
// Definition: report_WirelessClientsReport_17_topPercentile200ResponseType
type SCIReportWirelessClientsReport17topPercentile200ResponseType struct {
	Data SCIWirelessClientsReport17topPercentileData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportWirelessClientsReport17topPercentile200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessClientsReport17topPercentile200ResponseType
}

func newSCIReportWirelessClientsReport17topPercentile200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessClientsReport17topPercentile200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessClientsReport17topPercentile200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessClientsReport17topPercentile200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessClientsReport17topPercentile200ResponseType() *SCIReportWirelessClientsReport17topPercentile200ResponseType {
	m := new(SCIReportWirelessClientsReport17topPercentile200ResponseType)
	return m
}

// SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType
//
// Definition: report_WirelessClientsReport_18_topNOSByClientCount200ResponseType
type SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType struct {
	Data SCIWirelessClientsReport18topNOSByClientCountData `json:"data,omitempty"`

	Metadata *SCIWirelessClientsReport18topNOSByClientCountMetaData `json:"metadata,omitempty"`
}

type SCIReportWirelessClientsReport18topNOSByClientCount200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType
}

func newSCIReportWirelessClientsReport18topNOSByClientCount200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessClientsReport18topNOSByClientCount200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessClientsReport18topNOSByClientCount200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessClientsReport18topNOSByClientCount200ResponseType() *SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType {
	m := new(SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType)
	return m
}

// SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType
//
// Definition: report_WirelessClientsReport_19_top10ManufacturersByClientCount200ResponseType
type SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType struct {
	Data SCIWirelessClientsReport19top10ManufacturersByClientCountData `json:"data,omitempty"`

	Metadata *SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData `json:"metadata,omitempty"`
}

type SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType
}

func newSCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType() *SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType {
	m := new(SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType)
	return m
}

// SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType
//
// Definition: report_WirelessClientsReport_112_top10AuthenticationMechanismByClientCount200ResponseType
type SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType struct {
	Data SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData `json:"data,omitempty"`

	Metadata *SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData `json:"metadata,omitempty"`
}

type SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType
}

func newSCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType() *SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType {
	m := new(SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType)
	return m
}

// SCIReportWithRelations
//
// Definition: report_withRelations
type SCIReportWithRelations struct {
	Component *string `json:"component,omitempty"`

	DatasourcesUsed []string `json:"datasourcesUsed,omitempty"`

	ExcludedFilters interface{} `json:"excludedFilters,omitempty"`

	FilterDataSource interface{} `json:"filterDataSource,omitempty"`

	Filters []interface{} `json:"filters,omitempty"`

	Headers []string `json:"headers,omitempty"`

	Id *float64 `json:"id,omitempty"`

	Layout []*SCIReportWithRelationsLayoutType `json:"layout,omitempty"`

	RouteParameters *SCIReportWithRelationsRouteParametersType `json:"routeParameters,omitempty"`

	Schedules []interface{} `json:"schedules,omitempty"`

	Sections []*SCIReportWithRelationsSectionsType `json:"sections,omitempty"`

	Title *string `json:"title,omitempty"`

	UrlSegmentName *string `json:"urlSegmentName,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelations) UnmarshalJSON(b []byte) error {
	type _SCIReportWithRelations SCIReportWithRelations
	tmpType := new(_SCIReportWithRelations)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "component")
	delete(tmpType.XAdditionalProperties, "datasourcesUsed")
	delete(tmpType.XAdditionalProperties, "excludedFilters")
	delete(tmpType.XAdditionalProperties, "filterDataSource")
	delete(tmpType.XAdditionalProperties, "filters")
	delete(tmpType.XAdditionalProperties, "headers")
	delete(tmpType.XAdditionalProperties, "id")
	delete(tmpType.XAdditionalProperties, "layout")
	delete(tmpType.XAdditionalProperties, "routeParameters")
	delete(tmpType.XAdditionalProperties, "schedules")
	delete(tmpType.XAdditionalProperties, "sections")
	delete(tmpType.XAdditionalProperties, "title")
	delete(tmpType.XAdditionalProperties, "urlSegmentName")
	*t = SCIReportWithRelations(*tmpType)
	return nil
}

func (t *SCIReportWithRelations) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Component != nil {
		tmp["component"] = t.Component
	}
	if t.DatasourcesUsed != nil {
		tmp["datasourcesUsed"] = t.DatasourcesUsed
	}
	if t.ExcludedFilters != nil {
		tmp["excludedFilters"] = t.ExcludedFilters
	}
	if t.FilterDataSource != nil {
		tmp["filterDataSource"] = t.FilterDataSource
	}
	if t.Filters != nil {
		tmp["filters"] = t.Filters
	}
	if t.Headers != nil {
		tmp["headers"] = t.Headers
	}
	if t.Id != nil {
		tmp["id"] = t.Id
	}
	if t.Layout != nil {
		tmp["layout"] = t.Layout
	}
	if t.RouteParameters != nil {
		tmp["routeParameters"] = t.RouteParameters
	}
	if t.Schedules != nil {
		tmp["schedules"] = t.Schedules
	}
	if t.Sections != nil {
		tmp["sections"] = t.Sections
	}
	if t.Title != nil {
		tmp["title"] = t.Title
	}
	if t.UrlSegmentName != nil {
		tmp["urlSegmentName"] = t.UrlSegmentName
	}
	return json.Marshal(tmp)
}

type SCIReportWithRelationsAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWithRelations
}

func newSCIReportWithRelationsAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWithRelationsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWithRelationsAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWithRelations)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWithRelations() *SCIReportWithRelations {
	m := new(SCIReportWithRelations)
	return m
}

// SCIReportWithRelationsLayoutType
//
// Definition: report_withRelationsLayoutType
type SCIReportWithRelationsLayoutType struct {
	DesiredWidth *string `json:"desiredWidth,omitempty"`

	Layout []*SCIReportWithRelationsLayoutTypeLayoutType `json:"layout,omitempty"`

	Section *float64 `json:"section,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelationsLayoutType) UnmarshalJSON(b []byte) error {
	type _SCIReportWithRelationsLayoutType SCIReportWithRelationsLayoutType
	tmpType := new(_SCIReportWithRelationsLayoutType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "desiredWidth")
	delete(tmpType.XAdditionalProperties, "layout")
	delete(tmpType.XAdditionalProperties, "section")
	*t = SCIReportWithRelationsLayoutType(*tmpType)
	return nil
}

func (t *SCIReportWithRelationsLayoutType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.DesiredWidth != nil {
		tmp["desiredWidth"] = t.DesiredWidth
	}
	if t.Layout != nil {
		tmp["layout"] = t.Layout
	}
	if t.Section != nil {
		tmp["section"] = t.Section
	}
	return json.Marshal(tmp)
}

func NewSCIReportWithRelationsLayoutType() *SCIReportWithRelationsLayoutType {
	m := new(SCIReportWithRelationsLayoutType)
	return m
}

// SCIReportWithRelationsLayoutTypeLayoutType
//
// Definition: report_withRelationsLayoutTypeLayoutType
type SCIReportWithRelationsLayoutTypeLayoutType struct {
	DesiredWidth *string `json:"desiredWidth,omitempty"`

	Section *float64 `json:"section,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelationsLayoutTypeLayoutType) UnmarshalJSON(b []byte) error {
	type _SCIReportWithRelationsLayoutTypeLayoutType SCIReportWithRelationsLayoutTypeLayoutType
	tmpType := new(_SCIReportWithRelationsLayoutTypeLayoutType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "desiredWidth")
	delete(tmpType.XAdditionalProperties, "section")
	*t = SCIReportWithRelationsLayoutTypeLayoutType(*tmpType)
	return nil
}

func (t *SCIReportWithRelationsLayoutTypeLayoutType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.DesiredWidth != nil {
		tmp["desiredWidth"] = t.DesiredWidth
	}
	if t.Section != nil {
		tmp["section"] = t.Section
	}
	return json.Marshal(tmp)
}

func NewSCIReportWithRelationsLayoutTypeLayoutType() *SCIReportWithRelationsLayoutTypeLayoutType {
	m := new(SCIReportWithRelationsLayoutTypeLayoutType)
	return m
}

// SCIReportWithRelationsRouteParametersType
//
// Definition: report_withRelationsRouteParametersType
type SCIReportWithRelationsRouteParametersType struct {
	Id *string `json:"id,omitempty"`
}

func NewSCIReportWithRelationsRouteParametersType() *SCIReportWithRelationsRouteParametersType {
	m := new(SCIReportWithRelationsRouteParametersType)
	return m
}

// SCIReportWithRelationsSectionsType
//
// Definition: report_withRelationsSectionsType
type SCIReportWithRelationsSectionsType struct {
	Component *string `json:"component,omitempty"`

	DefaultParameters *SCIReportWithRelationsSectionsTypeDefaultParametersType `json:"defaultParameters,omitempty"`

	Id *float64 `json:"id,omitempty"`

	Layout *SCIReportWithRelationsSectionsTypeLayoutType `json:"layout,omitempty"`

	QueryName *string `json:"queryName,omitempty"`

	SystemOwnerOnly *bool `json:"systemOwnerOnly,omitempty"`

	Title *string `json:"title,omitempty"`

	Url interface{} `json:"url,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelationsSectionsType) UnmarshalJSON(b []byte) error {
	type _SCIReportWithRelationsSectionsType SCIReportWithRelationsSectionsType
	tmpType := new(_SCIReportWithRelationsSectionsType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "component")
	delete(tmpType.XAdditionalProperties, "defaultParameters")
	delete(tmpType.XAdditionalProperties, "id")
	delete(tmpType.XAdditionalProperties, "layout")
	delete(tmpType.XAdditionalProperties, "queryName")
	delete(tmpType.XAdditionalProperties, "systemOwnerOnly")
	delete(tmpType.XAdditionalProperties, "title")
	delete(tmpType.XAdditionalProperties, "url")
	*t = SCIReportWithRelationsSectionsType(*tmpType)
	return nil
}

func (t *SCIReportWithRelationsSectionsType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Component != nil {
		tmp["component"] = t.Component
	}
	if t.DefaultParameters != nil {
		tmp["defaultParameters"] = t.DefaultParameters
	}
	if t.Id != nil {
		tmp["id"] = t.Id
	}
	if t.Layout != nil {
		tmp["layout"] = t.Layout
	}
	if t.QueryName != nil {
		tmp["queryName"] = t.QueryName
	}
	if t.SystemOwnerOnly != nil {
		tmp["systemOwnerOnly"] = t.SystemOwnerOnly
	}
	if t.Title != nil {
		tmp["title"] = t.Title
	}
	if t.Url != nil {
		tmp["url"] = t.Url
	}
	return json.Marshal(tmp)
}

func NewSCIReportWithRelationsSectionsType() *SCIReportWithRelationsSectionsType {
	m := new(SCIReportWithRelationsSectionsType)
	return m
}

// SCIReportWithRelationsSectionsTypeDefaultParametersType
//
// Definition: report_withRelationsSectionsTypeDefaultParametersType
type SCIReportWithRelationsSectionsTypeDefaultParametersType struct {
	Granularity *string `json:"granularity,omitempty"`

	Hidden *bool `json:"hidden,omitempty"`

	Limit *float64 `json:"limit,omitempty"`

	Metric *string `json:"metric,omitempty"`

	ToggleSection *float64 `json:"toggleSection,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeDefaultParametersType() *SCIReportWithRelationsSectionsTypeDefaultParametersType {
	m := new(SCIReportWithRelationsSectionsTypeDefaultParametersType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutType
//
// Definition: report_withRelationsSectionsTypeLayoutType
type SCIReportWithRelationsSectionsTypeLayoutType struct {
	Columns []*SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType `json:"columns,omitempty"`

	Formats *SCIReportWithRelationsSectionsTypeLayoutTypeFormatsType `json:"formats,omitempty"`

	Headers []*SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType `json:"headers,omitempty"`

	Image *SCIReportWithRelationsSectionsTypeLayoutTypeImageType `json:"image,omitempty"`

	Rows []*SCIReportWithRelationsSectionsTypeLayoutTypeRowsType `json:"rows,omitempty"`

	Series []*SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType `json:"series,omitempty"`

	SubSections []*SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType `json:"subSections,omitempty"`

	WidgetTheme *string `json:"widgetTheme,omitempty"`

	Width *string `json:"width,omitempty"`

	XAxisType *string `json:"xAxisType,omitempty"`

	YAxisType *string `json:"yAxisType,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutType() *SCIReportWithRelationsSectionsTypeLayoutType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeColumnsType
type SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType struct {
	ColumnName *string `json:"columnName,omitempty"`

	CustomComponent *string `json:"customComponent,omitempty"`

	DisplayName *string `json:"displayName,omitempty"`

	DrillDownRoute *string `json:"drillDownRoute,omitempty"`

	Format *string `json:"format,omitempty"`

	Hidden *bool `json:"hidden,omitempty"`

	Suffix *string `json:"suffix,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType) UnmarshalJSON(b []byte) error {
	type _SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType
	tmpType := new(_SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "columnName")
	delete(tmpType.XAdditionalProperties, "customComponent")
	delete(tmpType.XAdditionalProperties, "displayName")
	delete(tmpType.XAdditionalProperties, "drillDownRoute")
	delete(tmpType.XAdditionalProperties, "format")
	delete(tmpType.XAdditionalProperties, "hidden")
	delete(tmpType.XAdditionalProperties, "suffix")
	*t = SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType(*tmpType)
	return nil
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ColumnName != nil {
		tmp["columnName"] = t.ColumnName
	}
	if t.CustomComponent != nil {
		tmp["customComponent"] = t.CustomComponent
	}
	if t.DisplayName != nil {
		tmp["displayName"] = t.DisplayName
	}
	if t.DrillDownRoute != nil {
		tmp["drillDownRoute"] = t.DrillDownRoute
	}
	if t.Format != nil {
		tmp["format"] = t.Format
	}
	if t.Hidden != nil {
		tmp["hidden"] = t.Hidden
	}
	if t.Suffix != nil {
		tmp["suffix"] = t.Suffix
	}
	return json.Marshal(tmp)
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeColumnsType() *SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeColumnsType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeFormatsType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeFormatsType
type SCIReportWithRelationsSectionsTypeLayoutTypeFormatsType struct {
	AvgRateTotalTraffic *string `json:"avgRateTotalTraffic,omitempty"`

	AvgSessionDuration *string `json:"avgSessionDuration,omitempty"`

	NumApps *string `json:"numApps,omitempty"`

	SessionCount *string `json:"sessionCount,omitempty"`

	TotalAps *string `json:"totalAps,omitempty"`

	TotalTraffic *string `json:"totalTraffic,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeFormatsType() *SCIReportWithRelationsSectionsTypeLayoutTypeFormatsType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeFormatsType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeHeadersType
type SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType struct {
	Component *string `json:"component,omitempty"`

	Content *SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentType `json:"content,omitempty"`

	Name *string `json:"name,omitempty"`

	Options *SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeOptionsType `json:"options,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType) UnmarshalJSON(b []byte) error {
	type _SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType
	tmpType := new(_SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "component")
	delete(tmpType.XAdditionalProperties, "content")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "options")
	*t = SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType(*tmpType)
	return nil
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Component != nil {
		tmp["component"] = t.Component
	}
	if t.Content != nil {
		tmp["content"] = t.Content
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.Options != nil {
		tmp["options"] = t.Options
	}
	return json.Marshal(tmp)
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeHeadersType() *SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeHeadersTypeContentType
type SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentType struct {
	Formats *SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentTypeFormatsType `json:"formats,omitempty"`

	Text *string `json:"text,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentType() *SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentTypeFormatsType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeHeadersTypeContentTypeFormatsType
type SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentTypeFormatsType struct {
	Percentage *string `json:"percentage,omitempty"`

	TotalTraffic *string `json:"totalTraffic,omitempty"`

	Traffic *string `json:"traffic,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentTypeFormatsType() *SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentTypeFormatsType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeContentTypeFormatsType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeOptionsType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeHeadersTypeOptionsType
type SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeOptionsType struct {
	RxBytes *string `json:"rxBytes,omitempty"`

	Traffic *string `json:"traffic,omitempty"`

	TxBytes *string `json:"txBytes,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeOptionsType() *SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeOptionsType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeHeadersTypeOptionsType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeImageType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeImageType
type SCIReportWithRelationsSectionsTypeLayoutTypeImageType struct {
	Name *string `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeImageType() *SCIReportWithRelationsSectionsTypeLayoutTypeImageType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeImageType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeRowsType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeRowsType
type SCIReportWithRelationsSectionsTypeLayoutTypeRowsType struct {
	DrillDownRoute *string `json:"drillDownRoute,omitempty"`

	DrillDownTitle *string `json:"drillDownTitle,omitempty"`

	Label *string `json:"label,omitempty"`

	Value *string `json:"value,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeRowsType) UnmarshalJSON(b []byte) error {
	type _SCIReportWithRelationsSectionsTypeLayoutTypeRowsType SCIReportWithRelationsSectionsTypeLayoutTypeRowsType
	tmpType := new(_SCIReportWithRelationsSectionsTypeLayoutTypeRowsType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "drillDownRoute")
	delete(tmpType.XAdditionalProperties, "drillDownTitle")
	delete(tmpType.XAdditionalProperties, "label")
	delete(tmpType.XAdditionalProperties, "value")
	*t = SCIReportWithRelationsSectionsTypeLayoutTypeRowsType(*tmpType)
	return nil
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeRowsType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.DrillDownRoute != nil {
		tmp["drillDownRoute"] = t.DrillDownRoute
	}
	if t.DrillDownTitle != nil {
		tmp["drillDownTitle"] = t.DrillDownTitle
	}
	if t.Label != nil {
		tmp["label"] = t.Label
	}
	if t.Value != nil {
		tmp["value"] = t.Value
	}
	return json.Marshal(tmp)
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeRowsType() *SCIReportWithRelationsSectionsTypeLayoutTypeRowsType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeRowsType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeSeriesType
type SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType struct {
	Area *bool `json:"area,omitempty"`

	Color *string `json:"color,omitempty"`

	Key *string `json:"key,omitempty"`

	Values *string `json:"values,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType) UnmarshalJSON(b []byte) error {
	type _SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType
	tmpType := new(_SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "area")
	delete(tmpType.XAdditionalProperties, "color")
	delete(tmpType.XAdditionalProperties, "key")
	delete(tmpType.XAdditionalProperties, "values")
	*t = SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType(*tmpType)
	return nil
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Area != nil {
		tmp["area"] = t.Area
	}
	if t.Color != nil {
		tmp["color"] = t.Color
	}
	if t.Key != nil {
		tmp["key"] = t.Key
	}
	if t.Values != nil {
		tmp["values"] = t.Values
	}
	return json.Marshal(tmp)
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeSeriesType() *SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeSeriesType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeSubSectionsType
type SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType struct {
	Component *string `json:"component,omitempty"`

	Layout *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType `json:"layout,omitempty"`

	Title *string `json:"title,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType) UnmarshalJSON(b []byte) error {
	type _SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType
	tmpType := new(_SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "component")
	delete(tmpType.XAdditionalProperties, "layout")
	delete(tmpType.XAdditionalProperties, "title")
	*t = SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType(*tmpType)
	return nil
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Component != nil {
		tmp["component"] = t.Component
	}
	if t.Layout != nil {
		tmp["layout"] = t.Layout
	}
	if t.Title != nil {
		tmp["title"] = t.Title
	}
	return json.Marshal(tmp)
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType() *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType
type SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType struct {
	Component *string `json:"component,omitempty"`

	FormatMetadata *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeFormatMetadataType `json:"formatMetadata,omitempty"`

	HideLegendLabels *bool `json:"hideLegendLabels,omitempty"`

	Label *string `json:"label,omitempty"`

	LabelFormat *string `json:"labelFormat,omitempty"`

	Series []*SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType `json:"series,omitempty"`

	Title *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeTitleType `json:"title,omitempty"`

	Value *string `json:"value,omitempty"`

	Width *string `json:"width,omitempty"`

	XAxisType *string `json:"xAxisType,omitempty"`

	YAxisType *string `json:"yAxisType,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType() *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeFormatMetadataType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeFormatMetadataType
type SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeFormatMetadataType struct {
	SubTitle *string `json:"subTitle,omitempty"`

	TotalTraffic *string `json:"totalTraffic,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeFormatMetadataType() *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeFormatMetadataType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeFormatMetadataType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType
type SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType struct {
	Color *string `json:"color,omitempty"`

	Disabled *bool `json:"disabled,omitempty"`

	Key *string `json:"key,omitempty"`

	Values *string `json:"values,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType) UnmarshalJSON(b []byte) error {
	type _SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType
	tmpType := new(_SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "color")
	delete(tmpType.XAdditionalProperties, "disabled")
	delete(tmpType.XAdditionalProperties, "key")
	delete(tmpType.XAdditionalProperties, "values")
	*t = SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType(*tmpType)
	return nil
}

func (t *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Color != nil {
		tmp["color"] = t.Color
	}
	if t.Disabled != nil {
		tmp["disabled"] = t.Disabled
	}
	if t.Key != nil {
		tmp["key"] = t.Key
	}
	if t.Values != nil {
		tmp["values"] = t.Values
	}
	return json.Marshal(tmp)
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType() *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeTitleType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeTitleType
type SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeTitleType struct {
	Default *string `json:"default,omitempty"`

	RxBytes *string `json:"rxBytes,omitempty"`

	TxBytes *string `json:"txBytes,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeTitleType() *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeTitleType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeTitleType)
	return m
}

// SCIReportWLANsReport35overview200ResponseType
//
// Definition: report_WLANsReport_35_overview200ResponseType
type SCIReportWLANsReport35overview200ResponseType struct {
	Data SCIWLANsReport35overviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportWLANsReport35overview200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWLANsReport35overview200ResponseType
}

func newSCIReportWLANsReport35overview200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWLANsReport35overview200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWLANsReport35overview200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWLANsReport35overview200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWLANsReport35overview200ResponseType() *SCIReportWLANsReport35overview200ResponseType {
	m := new(SCIReportWLANsReport35overview200ResponseType)
	return m
}

// SCIReportWLANsReport36top10SsidsByTraffic200ResponseType
//
// Definition: report_WLANsReport_36_top10SsidsByTraffic200ResponseType
type SCIReportWLANsReport36top10SsidsByTraffic200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportWLANsReport36top10SsidsByTraffic200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWLANsReport36top10SsidsByTraffic200ResponseType
}

func newSCIReportWLANsReport36top10SsidsByTraffic200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWLANsReport36top10SsidsByTraffic200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWLANsReport36top10SsidsByTraffic200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWLANsReport36top10SsidsByTraffic200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWLANsReport36top10SsidsByTraffic200ResponseType() *SCIReportWLANsReport36top10SsidsByTraffic200ResponseType {
	m := new(SCIReportWLANsReport36top10SsidsByTraffic200ResponseType)
	return m
}

// SCIReportWLANsReport37activeSsidsTrend200ResponseType
//
// Definition: report_WLANsReport_37_activeSsidsTrend200ResponseType
type SCIReportWLANsReport37activeSsidsTrend200ResponseType struct {
	Data SCIWLANsReport37activeSsidsTrendData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportWLANsReport37activeSsidsTrend200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWLANsReport37activeSsidsTrend200ResponseType
}

func newSCIReportWLANsReport37activeSsidsTrend200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWLANsReport37activeSsidsTrend200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWLANsReport37activeSsidsTrend200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWLANsReport37activeSsidsTrend200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWLANsReport37activeSsidsTrend200ResponseType() *SCIReportWLANsReport37activeSsidsTrend200ResponseType {
	m := new(SCIReportWLANsReport37activeSsidsTrend200ResponseType)
	return m
}

// SCIReportWLANsReport38top10SsidsByClientCount200ResponseType
//
// Definition: report_WLANsReport_38_top10SsidsByClientCount200ResponseType
type SCIReportWLANsReport38top10SsidsByClientCount200ResponseType struct {
	Data SCIWLANsReport38top10SsidsByClientCountData `json:"data,omitempty"`

	Metadata *SCIWLANsReport38top10SsidsByClientCountMetaData `json:"metadata,omitempty"`
}

type SCIReportWLANsReport38top10SsidsByClientCount200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWLANsReport38top10SsidsByClientCount200ResponseType
}

func newSCIReportWLANsReport38top10SsidsByClientCount200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWLANsReport38top10SsidsByClientCount200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWLANsReport38top10SsidsByClientCount200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWLANsReport38top10SsidsByClientCount200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWLANsReport38top10SsidsByClientCount200ResponseType() *SCIReportWLANsReport38top10SsidsByClientCount200ResponseType {
	m := new(SCIReportWLANsReport38top10SsidsByClientCount200ResponseType)
	return m
}

// SCIReportWLANsReport39ssidChangesOverTime200ResponseType
//
// Definition: report_WLANsReport_39_ssidChangesOverTime200ResponseType
type SCIReportWLANsReport39ssidChangesOverTime200ResponseType struct {
	Data SCIWLANsReport39ssidChangesOverTimeData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

type SCIReportWLANsReport39ssidChangesOverTime200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWLANsReport39ssidChangesOverTime200ResponseType
}

func newSCIReportWLANsReport39ssidChangesOverTime200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWLANsReport39ssidChangesOverTime200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWLANsReport39ssidChangesOverTime200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWLANsReport39ssidChangesOverTime200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWLANsReport39ssidChangesOverTime200ResponseType() *SCIReportWLANsReport39ssidChangesOverTime200ResponseType {
	m := new(SCIReportWLANsReport39ssidChangesOverTime200ResponseType)
	return m
}

// SCIReportWLANsReport40topSsidsByTrafficTable200ResponseType
//
// Definition: report_WLANsReport_40_topSsidsByTrafficTable200ResponseType
type SCIReportWLANsReport40topSsidsByTrafficTable200ResponseType struct {
	Data SCIWLANsReport40topSsidsByTrafficTableData `json:"data,omitempty"`

	Metadata *SCIWLANsReport40topSsidsByTrafficTableMetaData `json:"metadata,omitempty"`
}

type SCIReportWLANsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWLANsReport40topSsidsByTrafficTable200ResponseType
}

func newSCIReportWLANsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWLANsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWLANsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWLANsReport40topSsidsByTrafficTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWLANsReport40topSsidsByTrafficTable200ResponseType() *SCIReportWLANsReport40topSsidsByTrafficTable200ResponseType {
	m := new(SCIReportWLANsReport40topSsidsByTrafficTable200ResponseType)
	return m
}

// SCIReportWLANsReport41topSsidsByClientsTable200ResponseType
//
// Definition: report_WLANsReport_41_topSsidsByClientsTable200ResponseType
type SCIReportWLANsReport41topSsidsByClientsTable200ResponseType struct {
	Data SCIWLANsReport41topSsidsByClientsTableData `json:"data,omitempty"`

	Metadata *SCIWLANsReport41topSsidsByClientsTableMetaData `json:"metadata,omitempty"`
}

type SCIReportWLANsReport41topSsidsByClientsTable200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIReportWLANsReport41topSsidsByClientsTable200ResponseType
}

func newSCIReportWLANsReport41topSsidsByClientsTable200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIReportWLANsReport41topSsidsByClientsTable200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIReportWLANsReport41topSsidsByClientsTable200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIReportWLANsReport41topSsidsByClientsTable200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIReportWLANsReport41topSsidsByClientsTable200ResponseType() *SCIReportWLANsReport41topSsidsByClientsTable200ResponseType {
	m := new(SCIReportWLANsReport41topSsidsByClientsTable200ResponseType)
	return m
}

// ReportDownloadReport
//
// Operation ID: report_downloadReport
// Operation path: /reports/{id}/download/{format}
// Success code: 204 (No Content)
//
// Form data parameters:
// - state string
//		- required
// - timezone string
//		- required
//
// Required parameters:
// - format string
//		- required
// - id string
//		- required
func (s *SCIReportService) ReportDownloadReport(ctx context.Context, formValues url.Values, format string, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportDownloadReport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("format", format)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// ReportFind
//
// Find all instances of the model matched by filter from the data source.
//
// Operation ID: report_find
// Operation path: /reports
// Success code: 200 (OK)
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIReportService) ReportFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIReportFind200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportFind200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportFind200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIReportFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportFind200ResponseTypeAPIResponse), err
}

// ReportFindById
//
// Find a model instance by id from the data source.
//
// Operation ID: report_findById
// Operation path: /reports/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIReportService) ReportFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsReportAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsReportAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsReportAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIReportFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsReportAPIResponse), err
}

// ReportGetData
//
// For the <b><code>filter</code></b> field below, an example would be <pre><code class="json">{ "type": "or", "fields": [{ "type": "selector", "dimension": "apMac", "value": "000000000000" }]}</code></pre>
//
// Operation ID: report_getData
// Operation path: /reports/{id}/sections/{sectionId}/data
// Success code: 200 (OK)
//
// Form data parameters:
// - end string
//		- required
// - filter string
//		- nullable
// - granularity string
//		- nullable
// - limit float64
//		- nullable
// - metric string
//		- nullable
// - pagingIdentifiers string
//		- nullable
// - start string
//		- required
// - switchFilter string
//		- nullable
//
// Required parameters:
// - id string
//		- required
// - sectionId string
//		- required
func (s *SCIReportService) ReportGetData(ctx context.Context, formValues url.Values, id string, sectionId string, mutators ...RequestMutator) (*SCIReportGetData200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportGetData200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportGetData200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportGetData, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*SCIReportGetData200ResponseTypeAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("sectionId", sectionId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportGetData200ResponseTypeAPIResponse), err
}

// ReportLatestIngestedTime
//
// Operation ID: report_latestIngestedTime
// Operation path: /reports/latestIngestedTime
// Success code: 200 (OK)
func (s *SCIReportService) ReportLatestIngestedTime(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCIReportLatestIngestedTime, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// ReportPrototypeGetSections
//
// Queries sections of report.
//
// Operation ID: report_prototype_get_sections
// Operation path: /reports/{id}/sections
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIReportService) ReportPrototypeGetSections(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIReportPrototypegetsections200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportPrototypegetsections200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportPrototypegetsections200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIReportPrototypeGetSections, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportPrototypegetsections200ResponseTypeAPIResponse), err
}

// ReportWithRelations
//
// For the <b><code>urlSegmentName</code></b> field below, examples could be <code>overview</code>, <code>network</code>, <code>ap</code>, <code>clients</code>
//
// Operation ID: report_withRelations
// Operation path: /reports/withRelations
// Success code: 200 (OK)
//
// Required parameters:
// - urlSegmentName string
//		- required
//		- oneof:[ap,client,switch,wlans,airtime,applications,network,aps,overview,comparison,health]
func (s *SCIReportService) ReportWithRelations(ctx context.Context, urlSegmentName string, mutators ...RequestMutator) (*SCIReportWithRelationsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportWithRelationsAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportWithRelationsAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("urlSegmentName", urlSegmentName)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportWithRelationsAPIResponse), err
}
