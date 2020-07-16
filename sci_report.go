package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
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

func NewSCIReportComparisonReport147comparisionTable200ResponseType() *SCIReportComparisonReport147comparisionTable200ResponseType {
	m := new(SCIReportComparisonReport147comparisionTable200ResponseType)
	return m
}

// SCIReportFind200ResponseType
//
// Definition: report_find200ResponseType
type SCIReportFind200ResponseType []*SCIModelsReport

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

func NewSCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType() *SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType {
	m := new(SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType)
	return m
}

// SCIReportInventorySwitchesReport113overview200ResponseType
//
// Definition: report_InventorySwitchesReport_113_overview200ResponseType
type SCIReportInventorySwitchesReport113overview200ResponseType struct {
	Data SCIInventorySwitchesReport113overviewData `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
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

func NewSCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType() *SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType {
	m := new(SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType)
	return m
}

// SCIReportOverview62overview200ResponseType
//
// Definition: report_Overview_62_overview200ResponseType
type SCIReportOverview62overview200ResponseType struct {
	Data SCIOverview62overviewData `json:"data,omitempty"`

	Metadata *SCIOverview62overviewMetaData `json:"metadata,omitempty"`
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

func NewSCIReportOverview115networkUsageOverview200ResponseType() *SCIReportOverview115networkUsageOverview200ResponseType {
	m := new(SCIReportOverview115networkUsageOverview200ResponseType)
	return m
}

// SCIReportPrototypegetsections200ResponseType
//
// Definition: report_prototype___get__sections200ResponseType
type SCIReportPrototypegetsections200ResponseType []*SCIModelsSection

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

func NewSCIReportSwitchDetailsReport139switchDetails200ResponseType() *SCIReportSwitchDetailsReport139switchDetails200ResponseType {
	m := new(SCIReportSwitchDetailsReport139switchDetails200ResponseType)
	return m
}

// SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType
//
// Definition: report_WirelessApplicationsReport_7_top10ApplicationsByTrafficVolume200ResponseType
type SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType struct {
	Data interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
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

	Id *int `json:"id,omitempty"`

	Layout []*SCIReportWithRelationsLayoutType `json:"layout,omitempty"`

	RouteParameters *SCIReportWithRelationsRouteParametersType `json:"routeParameters,omitempty"`

	Schedules []interface{} `json:"schedules,omitempty"`

	Sections []*SCIReportWithRelationsSectionsType `json:"sections,omitempty"`

	Title *string `json:"title,omitempty"`

	UrlSegmentName *string `json:"urlSegmentName,omitempty"`
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

	Section *int `json:"section,omitempty"`
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

	Id *int `json:"id,omitempty"`

	Layout *SCIReportWithRelationsSectionsTypeLayoutType `json:"layout,omitempty"`

	QueryName *string `json:"queryName,omitempty"`

	SystemOwnerOnly *bool `json:"systemOwnerOnly,omitempty"`

	Title *string `json:"title,omitempty"`

	Url interface{} `json:"url,omitempty"`
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

	Metric *string `json:"metric,omitempty"`

	ToggleSection *int `json:"toggleSection,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeDefaultParametersType() *SCIReportWithRelationsSectionsTypeDefaultParametersType {
	m := new(SCIReportWithRelationsSectionsTypeDefaultParametersType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutType
//
// Definition: report_withRelationsSectionsTypeLayoutType
type SCIReportWithRelationsSectionsTypeLayoutType struct {
	Headers []*SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType `json:"headers,omitempty"`

	SubSections []*SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType `json:"subSections,omitempty"`

	Width *string `json:"width,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutType() *SCIReportWithRelationsSectionsTypeLayoutType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeHeadersType
type SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType struct {
	Component *string `json:"component,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeHeadersType() *SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeHeadersType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeSubSectionsType
type SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType struct {
	Component *string `json:"component,omitempty"`

	Layout *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType `json:"layout,omitempty"`
}

func NewSCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType() *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType {
	m := new(SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsType)
	return m
}

// SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType
//
// Definition: report_withRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType
type SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutType struct {
	FormatMetadata *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeFormatMetadataType `json:"formatMetadata,omitempty"`

	Label *string `json:"label,omitempty"`

	LabelFormat *string `json:"labelFormat,omitempty"`

	Series []*SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeSeriesType `json:"series,omitempty"`

	Title *SCIReportWithRelationsSectionsTypeLayoutTypeSubSectionsTypeLayoutTypeTitleType `json:"title,omitempty"`

	Value *string `json:"value,omitempty"`

	Width *string `json:"width,omitempty"`
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

func NewSCIReportWLANsReport41topSsidsByClientsTable200ResponseType() *SCIReportWLANsReport41topSsidsByClientsTable200ResponseType {
	m := new(SCIReportWLANsReport41topSsidsByClientsTable200ResponseType)
	return m
}

// ReportDownloadReport
//
// Operation ID: report_downloadReport
//
// Form Data Parameters:
// - state string
//		- required
// - timezone string
//		- required
//
// Required Parameters:
// - format string
//		- required
// - id string
//		- required
func (s *SCIReportService) ReportDownloadReport(ctx context.Context, formValues url.Values, format string, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportDownloadReport, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return rm, err
	}
	req.SetPathParameter("format", format)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// ReportFind
//
// Operation ID: report_find
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIReportService) ReportFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCIReportFind200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIReportFind200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIReportFind, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIReportFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// ReportFindById
//
// Operation ID: report_findById
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
func (s *SCIReportService) ReportFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsReport, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsReport
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIReportFindById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsReport()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportGetData
//
// Operation ID: report_getData
//
// For the <b><code>filter</code></b> field below, an example would be <pre><code class="json">{ "type": "or", "fields": [{ "type": "selector", "dimension": "apMac", "value": "000000000000" }]}</code></pre>
//
// Form Data Parameters:
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
// Required Parameters:
// - id string
//		- required
// - sectionId string
//		- required
func (s *SCIReportService) ReportGetData(ctx context.Context, formValues url.Values, id string, sectionId string, mutators ...RequestMutator) (*SCIReportGetData200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportGetData200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportGetData, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("sectionId", sectionId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportGetData200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportPrototypeGetSections
//
// Operation ID: report_prototype___get__sections
//
// Queries sections of report.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIReportService) ReportPrototypeGetSections(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (SCIReportPrototypegetsections200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIReportPrototypegetsections200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIReportPrototypeGetSections, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIReportPrototypegetsections200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// ReportWithRelations
//
// Operation ID: report_withRelations
//
// For the <b><code>urlSegmentName</code></b> field below, examples could be <code>overview</code>, <code>network</code>, <code>ap</code>, <code>clients</code>
//
// Required Parameters:
// - urlSegmentName string
//		- required
//		- oneof:[ap,client,switch,wlans,airtime,applications,network,aps,overview,comparison,health]
func (s *SCIReportService) ReportWithRelations(ctx context.Context, urlSegmentName string, mutators ...RequestMutator) (*SCIReportWithRelations, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWithRelations
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWithRelations, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetQueryParameter("urlSegmentName", []string{urlSegmentName})
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWithRelations()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
