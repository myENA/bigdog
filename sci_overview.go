package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIOverviewService struct {
	apiClient *SCIClient
}

func NewSCIOverviewService(c *SCIClient) *SCIOverviewService {
	s := new(SCIOverviewService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIOverviewService() *SCIOverviewService {
	return NewSCIOverviewService(ss.apiClient)
}

// SCIOverview62overviewData
//
// Definition: Overview.Overview.62.overview.Data
type SCIOverview62overviewData []*SCIOverview62overviewDataType

func MakeSCIOverview62overviewData() SCIOverview62overviewData {
	m := make(SCIOverview62overviewData, 0)
	return m
}

// SCIOverview62overviewDataType
//
// Definition: Overview.Overview.62.overview.DataType
type SCIOverview62overviewDataType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ClientCount interface{} `json:"clientCount,omitempty"`

	RebootCount *float64 `json:"rebootCount,omitempty"`

	SessionCount interface{} `json:"sessionCount,omitempty"`

	ShortSessionRatio interface{} `json:"shortSessionRatio,omitempty"`

	Total *int `json:"total,omitempty"`

	TotalHistory *int `json:"totalHistory,omitempty"`

	UserTraffic interface{} `json:"userTraffic,omitempty"`
}

func NewSCIOverview62overviewDataType() *SCIOverview62overviewDataType {
	m := new(SCIOverview62overviewDataType)
	return m
}

// SCIOverview62overviewMetaData
//
// Definition: Overview.Overview.62.overview.MetaData
type SCIOverview62overviewMetaData struct {
	LastAnomalyUpdatedTime *string `json:"lastAnomalyUpdatedTime,omitempty"`
}

func NewSCIOverview62overviewMetaData() *SCIOverview62overviewMetaData {
	m := new(SCIOverview62overviewMetaData)
	return m
}

// SCIOverview63controllerData
//
// Definition: Overview.Overview.63.controller.Data
type SCIOverview63controllerData []*SCIOverview63controllerDataType

func MakeSCIOverview63controllerData() SCIOverview63controllerData {
	m := make(SCIOverview63controllerData, 0)
	return m
}

// SCIOverview63controllerDataType
//
// Definition: Overview.Overview.63.controller.DataType
type SCIOverview63controllerDataType struct {
	Offline *int `json:"offline,omitempty"`

	Online *int `json:"online,omitempty"`

	SzCount *int `json:"szCount,omitempty"`

	Total *int `json:"total,omitempty"`

	ZdCount *int `json:"zdCount,omitempty"`
}

func NewSCIOverview63controllerDataType() *SCIOverview63controllerDataType {
	m := new(SCIOverview63controllerDataType)
	return m
}

// SCIOverview64apOverviewData
//
// Definition: Overview.Overview.64.apOverview.Data
type SCIOverview64apOverviewData []*SCIOverview64apOverviewDataType

func MakeSCIOverview64apOverviewData() SCIOverview64apOverviewData {
	m := make(SCIOverview64apOverviewData, 0)
	return m
}

// SCIOverview64apOverviewDataType
//
// Definition: Overview.Overview.64.apOverview.DataType
type SCIOverview64apOverviewDataType struct {
	Offline *int `json:"offline,omitempty"`

	Online *int `json:"online,omitempty"`

	Others *int `json:"others,omitempty"`

	Reboots *int `json:"reboots,omitempty"`

	Total *int `json:"total,omitempty"`

	TotalApsWAlarm *int `json:"totalApsWAlarm,omitempty"`

	TotalApsWReboot *int `json:"totalApsWReboot,omitempty"`
}

func NewSCIOverview64apOverviewDataType() *SCIOverview64apOverviewDataType {
	m := new(SCIOverview64apOverviewDataType)
	return m
}

// SCIOverview66apAlarmOverviewData
//
// Definition: Overview.Overview.66.apAlarmOverview.Data
type SCIOverview66apAlarmOverviewData []*SCIOverview66apAlarmOverviewDataType

func MakeSCIOverview66apAlarmOverviewData() SCIOverview66apAlarmOverviewData {
	m := make(SCIOverview66apAlarmOverviewData, 0)
	return m
}

// SCIOverview66apAlarmOverviewDataType
//
// Definition: Overview.Overview.66.apAlarmOverview.DataType
type SCIOverview66apAlarmOverviewDataType struct {
	AlarmType *string `json:"alarmType,omitempty"`

	Count *int `json:"count,omitempty"`
}

func NewSCIOverview66apAlarmOverviewDataType() *SCIOverview66apAlarmOverviewDataType {
	m := new(SCIOverview66apAlarmOverviewDataType)
	return m
}

// SCIOverview66apAlarmOverviewMetaData
//
// Definition: Overview.Overview.66.apAlarmOverview.MetaData
type SCIOverview66apAlarmOverviewMetaData struct {
	AlarmStat *SCIOverview66apAlarmOverviewMetaDataAlarmStatType `json:"alarmStat,omitempty"`

	AlarmTotalCount *int `json:"alarmTotalCount,omitempty"`
}

func NewSCIOverview66apAlarmOverviewMetaData() *SCIOverview66apAlarmOverviewMetaData {
	m := new(SCIOverview66apAlarmOverviewMetaData)
	return m
}

// SCIOverview66apAlarmOverviewMetaDataAlarmStatType
//
// Definition: Overview.Overview.66.apAlarmOverview.MetaDataAlarmStatType
type SCIOverview66apAlarmOverviewMetaDataAlarmStatType struct {
	ByApGroup *SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByApGroupType `json:"byApGroup,omitempty"`

	ByZone *SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType `json:"byZone,omitempty"`
}

func NewSCIOverview66apAlarmOverviewMetaDataAlarmStatType() *SCIOverview66apAlarmOverviewMetaDataAlarmStatType {
	m := new(SCIOverview66apAlarmOverviewMetaDataAlarmStatType)
	return m
}

// SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByApGroupType
//
// Definition: Overview.Overview.66.apAlarmOverview.MetaDataAlarmStatTypeByApGroupType
type SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByApGroupType struct {
	Alarm *int `json:"alarm,omitempty"`

	Total *int `json:"total,omitempty"`
}

func NewSCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByApGroupType() *SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByApGroupType {
	m := new(SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByApGroupType)
	return m
}

// SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType
//
// Definition: Overview.Overview.66.apAlarmOverview.MetaDataAlarmStatTypeByZoneType
type SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType struct {
	Alarm *int `json:"alarm,omitempty"`

	Total *int `json:"total,omitempty"`
}

func NewSCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType() *SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType {
	m := new(SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType)
	return m
}

// SCIOverview67switchOverviewData
//
// Definition: Overview.Overview.67.switchOverview.Data
type SCIOverview67switchOverviewData []*SCIOverview67switchOverviewDataType

func MakeSCIOverview67switchOverviewData() SCIOverview67switchOverviewData {
	m := make(SCIOverview67switchOverviewData, 0)
	return m
}

// SCIOverview67switchOverviewDataType
//
// Definition: Overview.Overview.67.switchOverview.DataType
type SCIOverview67switchOverviewDataType struct {
	Offline *int `json:"offline,omitempty"`

	Online *int `json:"online,omitempty"`

	Others *int `json:"others,omitempty"`

	Total *int `json:"total,omitempty"`
}

func NewSCIOverview67switchOverviewDataType() *SCIOverview67switchOverviewDataType {
	m := new(SCIOverview67switchOverviewDataType)
	return m
}

// SCIOverview68apClientCountOverviewData
//
// Definition: Overview.Overview.68.apClientCountOverview.Data
type SCIOverview68apClientCountOverviewData []*SCIOverview68apClientCountOverviewDataType

func MakeSCIOverview68apClientCountOverviewData() SCIOverview68apClientCountOverviewData {
	m := make(SCIOverview68apClientCountOverviewData, 0)
	return m
}

// SCIOverview68apClientCountOverviewDataType
//
// Definition: Overview.Overview.68.apClientCountOverview.DataType
type SCIOverview68apClientCountOverviewDataType struct {
	Key *string `json:"key,omitempty"`

	Label *string `json:"label,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func NewSCIOverview68apClientCountOverviewDataType() *SCIOverview68apClientCountOverviewDataType {
	m := new(SCIOverview68apClientCountOverviewDataType)
	return m
}

// SCIOverview68apClientCountOverviewMetaData
//
// Definition: Overview.Overview.68.apClientCountOverview.MetaData
type SCIOverview68apClientCountOverviewMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCIOverview68apClientCountOverviewMetaData() *SCIOverview68apClientCountOverviewMetaData {
	m := new(SCIOverview68apClientCountOverviewMetaData)
	return m
}

// SCIOverview69totalTrafficMinMaxRateData
//
// Definition: Overview.Overview.69.totalTrafficMinMaxRate.Data
type SCIOverview69totalTrafficMinMaxRateData []*SCIOverview69totalTrafficMinMaxRateDataType

func MakeSCIOverview69totalTrafficMinMaxRateData() SCIOverview69totalTrafficMinMaxRateData {
	m := make(SCIOverview69totalTrafficMinMaxRateData, 0)
	return m
}

// SCIOverview69totalTrafficMinMaxRateDataType
//
// Definition: Overview.Overview.69.totalTrafficMinMaxRate.DataType
type SCIOverview69totalTrafficMinMaxRateDataType struct {
	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`
}

func NewSCIOverview69totalTrafficMinMaxRateDataType() *SCIOverview69totalTrafficMinMaxRateDataType {
	m := new(SCIOverview69totalTrafficMinMaxRateDataType)
	return m
}

// SCIOverview70sessionsOverviewData
//
// Definition: Overview.Overview.70.sessionsOverview.Data
type SCIOverview70sessionsOverviewData []*SCIOverview70sessionsOverviewDataType

func MakeSCIOverview70sessionsOverviewData() SCIOverview70sessionsOverviewData {
	m := make(SCIOverview70sessionsOverviewData, 0)
	return m
}

// SCIOverview70sessionsOverviewDataType
//
// Definition: Overview.Overview.70.sessionsOverview.DataType
type SCIOverview70sessionsOverviewDataType struct {
	Current *float64 `json:"current,omitempty"`
}

func NewSCIOverview70sessionsOverviewDataType() *SCIOverview70sessionsOverviewDataType {
	m := new(SCIOverview70sessionsOverviewDataType)
	return m
}

// SCIOverview71ssidOverviewData
//
// Definition: Overview.Overview.71.ssidOverview.Data
type SCIOverview71ssidOverviewData []*SCIOverview71ssidOverviewDataType

func MakeSCIOverview71ssidOverviewData() SCIOverview71ssidOverviewData {
	m := make(SCIOverview71ssidOverviewData, 0)
	return m
}

// SCIOverview71ssidOverviewDataType
//
// Definition: Overview.Overview.71.ssidOverview.DataType
type SCIOverview71ssidOverviewDataType struct {
	Current *int `json:"current,omitempty"`

	Previous *int `json:"previous,omitempty"`
}

func NewSCIOverview71ssidOverviewDataType() *SCIOverview71ssidOverviewDataType {
	m := new(SCIOverview71ssidOverviewDataType)
	return m
}

// SCIOverview71ssidOverviewMetaData
//
// Definition: Overview.Overview.71.ssidOverview.MetaData
type SCIOverview71ssidOverviewMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalMetric *int `json:"totalMetric,omitempty"`
}

func NewSCIOverview71ssidOverviewMetaData() *SCIOverview71ssidOverviewMetaData {
	m := new(SCIOverview71ssidOverviewMetaData)
	return m
}

// SCIOverview72radioOverviewData
//
// Definition: Overview.Overview.72.radioOverview.Data
type SCIOverview72radioOverviewData []*SCIOverview72radioOverviewDataType

func MakeSCIOverview72radioOverviewData() SCIOverview72radioOverviewData {
	m := make(SCIOverview72radioOverviewData, 0)
	return m
}

// SCIOverview72radioOverviewDataType
//
// Definition: Overview.Overview.72.radioOverview.DataType
type SCIOverview72radioOverviewDataType struct {
	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	Radio *string `json:"radio,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`
}

func NewSCIOverview72radioOverviewDataType() *SCIOverview72radioOverviewDataType {
	m := new(SCIOverview72radioOverviewDataType)
	return m
}

// SCIOverview73applicationsOverviewData
//
// Definition: Overview.Overview.73.applicationsOverview.Data
type SCIOverview73applicationsOverviewData []*SCIOverview73applicationsOverviewDataType

func MakeSCIOverview73applicationsOverviewData() SCIOverview73applicationsOverviewData {
	m := make(SCIOverview73applicationsOverviewData, 0)
	return m
}

// SCIOverview73applicationsOverviewDataType
//
// Definition: Overview.Overview.73.applicationsOverview.DataType
type SCIOverview73applicationsOverviewDataType struct {
	Key *string `json:"key,omitempty"`

	Label interface{} `json:"label,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func NewSCIOverview73applicationsOverviewDataType() *SCIOverview73applicationsOverviewDataType {
	m := new(SCIOverview73applicationsOverviewDataType)
	return m
}

// SCIOverview74apEventOverviewData
//
// Definition: Overview.Overview.74.apEventOverview.Data
type SCIOverview74apEventOverviewData []*SCIOverview74apEventOverviewDataType

func MakeSCIOverview74apEventOverviewData() SCIOverview74apEventOverviewData {
	m := make(SCIOverview74apEventOverviewData, 0)
	return m
}

// SCIOverview74apEventOverviewDataType
//
// Definition: Overview.Overview.74.apEventOverview.DataType
type SCIOverview74apEventOverviewDataType struct {
	Count *int `json:"count,omitempty"`

	EventType *string `json:"eventType,omitempty"`
}

func NewSCIOverview74apEventOverviewDataType() *SCIOverview74apEventOverviewDataType {
	m := new(SCIOverview74apEventOverviewDataType)
	return m
}

// SCIOverview74apEventOverviewMetaData
//
// Definition: Overview.Overview.74.apEventOverview.MetaData
type SCIOverview74apEventOverviewMetaData struct {
	EventTotalCount *int `json:"eventTotalCount,omitempty"`
}

func NewSCIOverview74apEventOverviewMetaData() *SCIOverview74apEventOverviewMetaData {
	m := new(SCIOverview74apEventOverviewMetaData)
	return m
}

// SCIOverview97factOverviewData
//
// Definition: Overview.Overview.97.factOverview.Data
type SCIOverview97factOverviewData []*SCIOverview97factOverviewDataType

func MakeSCIOverview97factOverviewData() SCIOverview97factOverviewData {
	m := make(SCIOverview97factOverviewData, 0)
	return m
}

// SCIOverview97factOverviewDataType
//
// Definition: Overview.Overview.97.factOverview.DataType
type SCIOverview97factOverviewDataType struct {
	Key *string `json:"key,omitempty"`

	Label *string `json:"label,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func NewSCIOverview97factOverviewDataType() *SCIOverview97factOverviewDataType {
	m := new(SCIOverview97factOverviewDataType)
	return m
}

// SCIOverview115networkUsageOverviewData
//
// Definition: Overview.Overview.115.networkUsageOverview.Data
type SCIOverview115networkUsageOverviewData struct {
	Ap []*SCIOverview115networkUsageOverviewDataApType `json:"Ap,omitempty"`

	ApGroup []*SCIOverview115networkUsageOverviewDataApGroupType `json:"ApGroup,omitempty"`

	App []*SCIOverview115networkUsageOverviewDataAppType `json:"App,omitempty"`

	Domain []*SCIOverview115networkUsageOverviewDataDomainType `json:"Domain,omitempty"`

	OS []*SCIOverview115networkUsageOverviewDataOSType `json:"OS,omitempty"`

	SSID []*SCIOverview115networkUsageOverviewDataSSIDType `json:"SSID,omitempty"`

	Switch []*SCIOverview115networkUsageOverviewDataSwitchType `json:"Switch,omitempty"`

	System []*SCIOverview115networkUsageOverviewDataSystemType `json:"System,omitempty"`

	Zone []*SCIOverview115networkUsageOverviewDataZoneType `json:"Zone,omitempty"`
}

func NewSCIOverview115networkUsageOverviewData() *SCIOverview115networkUsageOverviewData {
	m := new(SCIOverview115networkUsageOverviewData)
	return m
}

// SCIOverview115networkUsageOverviewDataApGroupType
//
// Definition: Overview.Overview.115.networkUsageOverview.DataApGroupType
type SCIOverview115networkUsageOverviewDataApGroupType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *int `json:"x,omitempty"`

	Y *int `json:"y,omitempty"`

	Z *int `json:"z,omitempty"`
}

func NewSCIOverview115networkUsageOverviewDataApGroupType() *SCIOverview115networkUsageOverviewDataApGroupType {
	m := new(SCIOverview115networkUsageOverviewDataApGroupType)
	return m
}

// SCIOverview115networkUsageOverviewDataAppType
//
// Definition: Overview.Overview.115.networkUsageOverview.DataAppType
type SCIOverview115networkUsageOverviewDataAppType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *int `json:"x,omitempty"`

	Y *int `json:"y,omitempty"`

	Z *int `json:"z,omitempty"`
}

func NewSCIOverview115networkUsageOverviewDataAppType() *SCIOverview115networkUsageOverviewDataAppType {
	m := new(SCIOverview115networkUsageOverviewDataAppType)
	return m
}

// SCIOverview115networkUsageOverviewDataApType
//
// Definition: Overview.Overview.115.networkUsageOverview.DataApType
type SCIOverview115networkUsageOverviewDataApType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *int `json:"x,omitempty"`

	Y *int `json:"y,omitempty"`

	Z *int `json:"z,omitempty"`
}

func NewSCIOverview115networkUsageOverviewDataApType() *SCIOverview115networkUsageOverviewDataApType {
	m := new(SCIOverview115networkUsageOverviewDataApType)
	return m
}

// SCIOverview115networkUsageOverviewDataDomainType
//
// Definition: Overview.Overview.115.networkUsageOverview.DataDomainType
type SCIOverview115networkUsageOverviewDataDomainType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *int `json:"x,omitempty"`

	Y *int `json:"y,omitempty"`

	Z *int `json:"z,omitempty"`
}

func NewSCIOverview115networkUsageOverviewDataDomainType() *SCIOverview115networkUsageOverviewDataDomainType {
	m := new(SCIOverview115networkUsageOverviewDataDomainType)
	return m
}

// SCIOverview115networkUsageOverviewDataOSType
//
// Definition: Overview.Overview.115.networkUsageOverview.DataOSType
type SCIOverview115networkUsageOverviewDataOSType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *int `json:"x,omitempty"`

	Y *int `json:"y,omitempty"`

	Z *int `json:"z,omitempty"`
}

func NewSCIOverview115networkUsageOverviewDataOSType() *SCIOverview115networkUsageOverviewDataOSType {
	m := new(SCIOverview115networkUsageOverviewDataOSType)
	return m
}

// SCIOverview115networkUsageOverviewDataSSIDType
//
// Definition: Overview.Overview.115.networkUsageOverview.DataSSIDType
type SCIOverview115networkUsageOverviewDataSSIDType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *int `json:"x,omitempty"`

	Y *int `json:"y,omitempty"`

	Z *int `json:"z,omitempty"`
}

func NewSCIOverview115networkUsageOverviewDataSSIDType() *SCIOverview115networkUsageOverviewDataSSIDType {
	m := new(SCIOverview115networkUsageOverviewDataSSIDType)
	return m
}

// SCIOverview115networkUsageOverviewDataSwitchType
//
// Definition: Overview.Overview.115.networkUsageOverview.DataSwitchType
type SCIOverview115networkUsageOverviewDataSwitchType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *int `json:"x,omitempty"`

	Y *int `json:"y,omitempty"`

	Z *int `json:"z,omitempty"`
}

func NewSCIOverview115networkUsageOverviewDataSwitchType() *SCIOverview115networkUsageOverviewDataSwitchType {
	m := new(SCIOverview115networkUsageOverviewDataSwitchType)
	return m
}

// SCIOverview115networkUsageOverviewDataSystemType
//
// Definition: Overview.Overview.115.networkUsageOverview.DataSystemType
type SCIOverview115networkUsageOverviewDataSystemType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *int `json:"x,omitempty"`

	Y *int `json:"y,omitempty"`

	Z *int `json:"z,omitempty"`
}

func NewSCIOverview115networkUsageOverviewDataSystemType() *SCIOverview115networkUsageOverviewDataSystemType {
	m := new(SCIOverview115networkUsageOverviewDataSystemType)
	return m
}

// SCIOverview115networkUsageOverviewDataZoneType
//
// Definition: Overview.Overview.115.networkUsageOverview.DataZoneType
type SCIOverview115networkUsageOverviewDataZoneType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *int `json:"x,omitempty"`

	Y *int `json:"y,omitempty"`

	Z *int `json:"z,omitempty"`
}

func NewSCIOverview115networkUsageOverviewDataZoneType() *SCIOverview115networkUsageOverviewDataZoneType {
	m := new(SCIOverview115networkUsageOverviewDataZoneType)
	return m
}

// ReportOverview62Overview
//
// Operation ID: report.Overview.62.overview
//
// Overview - Ruckus SmartAnalytics
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview62Overview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview62overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview62overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview62Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview62overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview63Controller
//
// Operation ID: report.Overview.63.controller
//
// Overview - Controllers
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview63Controller(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview63controller200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview63controller200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview63Controller, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview63controller200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview64ApOverview
//
// Operation ID: report.Overview.64.apOverview
//
// Overview - Access Points
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview64ApOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview64apOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview64apOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview64ApOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview64apOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview66ApAlarmOverview
//
// Operation ID: report.Overview.66.apAlarmOverview
//
// Overview - Alarms
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview66ApAlarmOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview66apAlarmOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview66apAlarmOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview66ApAlarmOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview66apAlarmOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview67SwitchOverview
//
// Operation ID: report.Overview.67.switchOverview
//
// Overview - Switches
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview67SwitchOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview67switchOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview67switchOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview67SwitchOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview67switchOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview68ApClientCountOverview
//
// Operation ID: report.Overview.68.apClientCountOverview
//
// Overview - Top APs by Client Count
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview68ApClientCountOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview68apClientCountOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview68apClientCountOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview68ApClientCountOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview68apClientCountOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview69TotalTrafficMinMaxRate
//
// Operation ID: report.Overview.69.totalTrafficMinMaxRate
//
// Overview - Total Wireless Traffic
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview69TotalTrafficMinMaxRate(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview69totalTrafficMinMaxRate200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview69totalTrafficMinMaxRate200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview69TotalTrafficMinMaxRate, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview69totalTrafficMinMaxRate200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview70SessionsOverview
//
// Operation ID: report.Overview.70.sessionsOverview
//
// Overview - Unique Wireless Sessions
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview70SessionsOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview70sessionsOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview70sessionsOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview70SessionsOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview70sessionsOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview71SsidOverview
//
// Operation ID: report.Overview.71.ssidOverview
//
// Overview - WLANs
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview71SsidOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview71ssidOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview71ssidOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview71SsidOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview71ssidOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview72RadioOverview
//
// Operation ID: report.Overview.72.radioOverview
//
// Overview - Radios
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview72RadioOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview72radioOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview72radioOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview72RadioOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview72radioOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview73ApplicationsOverview
//
// Operation ID: report.Overview.73.applicationsOverview
//
// Overview - Applications (Wireless)
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview73ApplicationsOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview73applicationsOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview73applicationsOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview73ApplicationsOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview73applicationsOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview74ApEventOverview
//
// Operation ID: report.Overview.74.apEventOverview
//
// Overview - Events
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview74ApEventOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview74apEventOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview74apEventOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview74ApEventOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview74apEventOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview97FactOverview
//
// Operation ID: report.Overview.97.factOverview
//
// Overview - Did you know?
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview97FactOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview97factOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview97factOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview97FactOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview97factOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportOverview115NetworkUsageOverview
//
// Operation ID: report.Overview.115.networkUsageOverview
//
// Overview - Network Usage Overview
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIOverviewService) ReportOverview115NetworkUsageOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview115networkUsageOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportOverview115networkUsageOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportOverview115NetworkUsageOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview115networkUsageOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
