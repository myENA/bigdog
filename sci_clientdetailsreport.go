package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIClientDetailsReportService struct {
	apiClient *SCIClient
}

func NewSCIClientDetailsReportService(c *SCIClient) *SCIClientDetailsReportService {
	s := new(SCIClientDetailsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIClientDetailsReportService() *SCIClientDetailsReportService {
	return NewSCIClientDetailsReportService(ss.apiClient)
}

// SCIClientDetailsReport8topAppsByTrafficTableData
//
// Definition: ClientDetailsReport_ClientDetailsReport_8_topAppsByTrafficTable_Data
type SCIClientDetailsReport8topAppsByTrafficTableData []*SCIClientDetailsReport8topAppsByTrafficTableDataType

func MakeSCIClientDetailsReport8topAppsByTrafficTableData() SCIClientDetailsReport8topAppsByTrafficTableData {
	m := make(SCIClientDetailsReport8topAppsByTrafficTableData, 0)
	return m
}

// SCIClientDetailsReport8topAppsByTrafficTableDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_8_topAppsByTrafficTable_DataType
type SCIClientDetailsReport8topAppsByTrafficTableDataType struct {
	App *string `json:"app,omitempty"`

	Index *float64 `json:"index,omitempty"`

	Port *string `json:"port,omitempty"`

	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientDetailsReport8topAppsByTrafficTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIClientDetailsReport8topAppsByTrafficTableDataType SCIClientDetailsReport8topAppsByTrafficTableDataType
	tmpType := new(_SCIClientDetailsReport8topAppsByTrafficTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "app")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "port")
	delete(tmpType.XAdditionalProperties, "rxBytes")
	delete(tmpType.XAdditionalProperties, "traffic")
	delete(tmpType.XAdditionalProperties, "txBytes")
	*t = SCIClientDetailsReport8topAppsByTrafficTableDataType(*tmpType)
	return nil
}

func (t *SCIClientDetailsReport8topAppsByTrafficTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.App != nil {
		tmp["app"] = t.App
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.Port != nil {
		tmp["port"] = t.Port
	}
	if t.RxBytes != nil {
		tmp["rxBytes"] = t.RxBytes
	}
	if t.Traffic != nil {
		tmp["traffic"] = t.Traffic
	}
	if t.TxBytes != nil {
		tmp["txBytes"] = t.TxBytes
	}
	return json.Marshal(tmp)
}

func NewSCIClientDetailsReport8topAppsByTrafficTableDataType() *SCIClientDetailsReport8topAppsByTrafficTableDataType {
	m := new(SCIClientDetailsReport8topAppsByTrafficTableDataType)
	return m
}

// SCIClientDetailsReport8topAppsByTrafficTableMetaData
//
// Definition: ClientDetailsReport_ClientDetailsReport_8_topAppsByTrafficTable_MetaData
type SCIClientDetailsReport8topAppsByTrafficTableMetaData struct {
	MaxValues *SCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientDetailsReport8topAppsByTrafficTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIClientDetailsReport8topAppsByTrafficTableMetaData SCIClientDetailsReport8topAppsByTrafficTableMetaData
	tmpType := new(_SCIClientDetailsReport8topAppsByTrafficTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "percentage")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "traffic")
	*t = SCIClientDetailsReport8topAppsByTrafficTableMetaData(*tmpType)
	return nil
}

func (t *SCIClientDetailsReport8topAppsByTrafficTableMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MaxValues != nil {
		tmp["maxValues"] = t.MaxValues
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.Percentage != nil {
		tmp["percentage"] = t.Percentage
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.Traffic != nil {
		tmp["traffic"] = t.Traffic
	}
	return json.Marshal(tmp)
}

func NewSCIClientDetailsReport8topAppsByTrafficTableMetaData() *SCIClientDetailsReport8topAppsByTrafficTableMetaData {
	m := new(SCIClientDetailsReport8topAppsByTrafficTableMetaData)
	return m
}

// SCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType
//
// Definition: ClientDetailsReport_ClientDetailsReport_8_topAppsByTrafficTable_MetaDataMaxValuesType
type SCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType struct {
	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`
}

func NewSCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType() *SCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType {
	m := new(SCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCIClientDetailsReport82rssTrendData
//
// Definition: ClientDetailsReport_ClientDetailsReport_82_rssTrend_Data
type SCIClientDetailsReport82rssTrendData []*SCIClientDetailsReport82rssTrendDataType

func MakeSCIClientDetailsReport82rssTrendData() SCIClientDetailsReport82rssTrendData {
	m := make(SCIClientDetailsReport82rssTrendData, 0)
	return m
}

// SCIClientDetailsReport82rssTrendDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_82_rssTrend_DataType
type SCIClientDetailsReport82rssTrendDataType struct {
	AvgRss *float64 `json:"avgRss,omitempty"`

	End *string `json:"end,omitempty"`

	MaxRss *float64 `json:"maxRss,omitempty"`

	MinRss *float64 `json:"minRss,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientDetailsReport82rssTrendDataType) UnmarshalJSON(b []byte) error {
	type _SCIClientDetailsReport82rssTrendDataType SCIClientDetailsReport82rssTrendDataType
	tmpType := new(_SCIClientDetailsReport82rssTrendDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avgRss")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "maxRss")
	delete(tmpType.XAdditionalProperties, "minRss")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCIClientDetailsReport82rssTrendDataType(*tmpType)
	return nil
}

func (t *SCIClientDetailsReport82rssTrendDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AvgRss != nil {
		tmp["avgRss"] = t.AvgRss
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.MaxRss != nil {
		tmp["maxRss"] = t.MaxRss
	}
	if t.MinRss != nil {
		tmp["minRss"] = t.MinRss
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCIClientDetailsReport82rssTrendDataType() *SCIClientDetailsReport82rssTrendDataType {
	m := new(SCIClientDetailsReport82rssTrendDataType)
	return m
}

// SCIClientDetailsReport83snrTrendData
//
// Definition: ClientDetailsReport_ClientDetailsReport_83_snrTrend_Data
type SCIClientDetailsReport83snrTrendData []*SCIClientDetailsReport83snrTrendDataType

func MakeSCIClientDetailsReport83snrTrendData() SCIClientDetailsReport83snrTrendData {
	m := make(SCIClientDetailsReport83snrTrendData, 0)
	return m
}

// SCIClientDetailsReport83snrTrendDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_83_snrTrend_DataType
type SCIClientDetailsReport83snrTrendDataType struct {
	AvgSnr *float64 `json:"avgSnr,omitempty"`

	End *string `json:"end,omitempty"`

	MaxSnr *float64 `json:"maxSnr,omitempty"`

	MinSnr *float64 `json:"minSnr,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientDetailsReport83snrTrendDataType) UnmarshalJSON(b []byte) error {
	type _SCIClientDetailsReport83snrTrendDataType SCIClientDetailsReport83snrTrendDataType
	tmpType := new(_SCIClientDetailsReport83snrTrendDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avgSnr")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "maxSnr")
	delete(tmpType.XAdditionalProperties, "minSnr")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCIClientDetailsReport83snrTrendDataType(*tmpType)
	return nil
}

func (t *SCIClientDetailsReport83snrTrendDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AvgSnr != nil {
		tmp["avgSnr"] = t.AvgSnr
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.MaxSnr != nil {
		tmp["maxSnr"] = t.MaxSnr
	}
	if t.MinSnr != nil {
		tmp["minSnr"] = t.MinSnr
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCIClientDetailsReport83snrTrendDataType() *SCIClientDetailsReport83snrTrendDataType {
	m := new(SCIClientDetailsReport83snrTrendDataType)
	return m
}

// SCIClientDetailsReport86summaryData
//
// Definition: ClientDetailsReport_ClientDetailsReport_86_summary_Data
type SCIClientDetailsReport86summaryData []*SCIClientDetailsReport86summaryDataType

func MakeSCIClientDetailsReport86summaryData() SCIClientDetailsReport86summaryData {
	m := make(SCIClientDetailsReport86summaryData, 0)
	return m
}

// SCIClientDetailsReport86summaryDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_86_summary_DataType
type SCIClientDetailsReport86summaryDataType struct {
	ClientIp *string `json:"clientIp,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	LastApMac *string `json:"lastApMac,omitempty"`

	LastApName *string `json:"lastApName,omitempty"`

	LastStatus *string `json:"lastStatus,omitempty"`

	Manufacturer *string `json:"manufacturer,omitempty"`

	OsType *string `json:"osType,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientDetailsReport86summaryDataType) UnmarshalJSON(b []byte) error {
	type _SCIClientDetailsReport86summaryDataType SCIClientDetailsReport86summaryDataType
	tmpType := new(_SCIClientDetailsReport86summaryDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "clientIp")
	delete(tmpType.XAdditionalProperties, "clientMac")
	delete(tmpType.XAdditionalProperties, "hostname")
	delete(tmpType.XAdditionalProperties, "lastApMac")
	delete(tmpType.XAdditionalProperties, "lastApName")
	delete(tmpType.XAdditionalProperties, "lastStatus")
	delete(tmpType.XAdditionalProperties, "manufacturer")
	delete(tmpType.XAdditionalProperties, "osType")
	*t = SCIClientDetailsReport86summaryDataType(*tmpType)
	return nil
}

func (t *SCIClientDetailsReport86summaryDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ClientIp != nil {
		tmp["clientIp"] = t.ClientIp
	}
	if t.ClientMac != nil {
		tmp["clientMac"] = t.ClientMac
	}
	if t.Hostname != nil {
		tmp["hostname"] = t.Hostname
	}
	if t.LastApMac != nil {
		tmp["lastApMac"] = t.LastApMac
	}
	if t.LastApName != nil {
		tmp["lastApName"] = t.LastApName
	}
	if t.LastStatus != nil {
		tmp["lastStatus"] = t.LastStatus
	}
	if t.Manufacturer != nil {
		tmp["manufacturer"] = t.Manufacturer
	}
	if t.OsType != nil {
		tmp["osType"] = t.OsType
	}
	return json.Marshal(tmp)
}

func NewSCIClientDetailsReport86summaryDataType() *SCIClientDetailsReport86summaryDataType {
	m := new(SCIClientDetailsReport86summaryDataType)
	return m
}

// SCIClientDetailsReport87clientStatsData
//
// Definition: ClientDetailsReport_ClientDetailsReport_87_clientStats_Data
type SCIClientDetailsReport87clientStatsData []*SCIClientDetailsReport87clientStatsDataType

func MakeSCIClientDetailsReport87clientStatsData() SCIClientDetailsReport87clientStatsData {
	m := make(SCIClientDetailsReport87clientStatsData, 0)
	return m
}

// SCIClientDetailsReport87clientStatsDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_87_clientStats_DataType
type SCIClientDetailsReport87clientStatsDataType struct {
	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgSessionDuration *float64 `json:"avgSessionDuration,omitempty"`

	NumApps *float64 `json:"numApps,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	TotalAps *float64 `json:"totalAps,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTraffic24 *float64 `json:"totalTraffic_2-4,omitempty"`

	TotalTraffic5 *float64 `json:"totalTraffic_5,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientDetailsReport87clientStatsDataType) UnmarshalJSON(b []byte) error {
	type _SCIClientDetailsReport87clientStatsDataType SCIClientDetailsReport87clientStatsDataType
	tmpType := new(_SCIClientDetailsReport87clientStatsDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avgRateTotalTraffic")
	delete(tmpType.XAdditionalProperties, "avgSessionDuration")
	delete(tmpType.XAdditionalProperties, "numApps")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "totalAps")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic_2-4")
	delete(tmpType.XAdditionalProperties, "totalTraffic_5")
	*t = SCIClientDetailsReport87clientStatsDataType(*tmpType)
	return nil
}

func (t *SCIClientDetailsReport87clientStatsDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AvgRateTotalTraffic != nil {
		tmp["avgRateTotalTraffic"] = t.AvgRateTotalTraffic
	}
	if t.AvgSessionDuration != nil {
		tmp["avgSessionDuration"] = t.AvgSessionDuration
	}
	if t.NumApps != nil {
		tmp["numApps"] = t.NumApps
	}
	if t.SessionCount != nil {
		tmp["sessionCount"] = t.SessionCount
	}
	if t.TotalAps != nil {
		tmp["totalAps"] = t.TotalAps
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.TotalTraffic24 != nil {
		tmp["totalTraffic_2-4"] = t.TotalTraffic24
	}
	if t.TotalTraffic5 != nil {
		tmp["totalTraffic_5"] = t.TotalTraffic5
	}
	return json.Marshal(tmp)
}

func NewSCIClientDetailsReport87clientStatsDataType() *SCIClientDetailsReport87clientStatsDataType {
	m := new(SCIClientDetailsReport87clientStatsDataType)
	return m
}

// SCIClientDetailsReport89trafficTrendData
//
// Definition: ClientDetailsReport_ClientDetailsReport_89_trafficTrend_Data
type SCIClientDetailsReport89trafficTrendData [][]*SCIClientDetailsReport89trafficTrendDataTypeType

func MakeSCIClientDetailsReport89trafficTrendData() SCIClientDetailsReport89trafficTrendData {
	m := make(SCIClientDetailsReport89trafficTrendData, 0)
	return m
}

// SCIClientDetailsReport89trafficTrendDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_89_trafficTrend_DataType
type SCIClientDetailsReport89trafficTrendDataType []*SCIClientDetailsReport89trafficTrendDataTypeType

func MakeSCIClientDetailsReport89trafficTrendDataType() SCIClientDetailsReport89trafficTrendDataType {
	m := make(SCIClientDetailsReport89trafficTrendDataType, 0)
	return m
}

// SCIClientDetailsReport89trafficTrendDataTypeType
//
// Definition: ClientDetailsReport_ClientDetailsReport_89_trafficTrend_DataTypeType
type SCIClientDetailsReport89trafficTrendDataTypeType struct {
	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	UserRxTraffic *float64 `json:"userRxTraffic,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTraffic24 *float64 `json:"userTraffic_2-4,omitempty"`

	UserTraffic5 *float64 `json:"userTraffic_5,omitempty"`

	UserTxTraffic *float64 `json:"userTxTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientDetailsReport89trafficTrendDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIClientDetailsReport89trafficTrendDataTypeType SCIClientDetailsReport89trafficTrendDataTypeType
	tmpType := new(_SCIClientDetailsReport89trafficTrendDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "userRxTraffic")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	delete(tmpType.XAdditionalProperties, "userTraffic_2-4")
	delete(tmpType.XAdditionalProperties, "userTraffic_5")
	delete(tmpType.XAdditionalProperties, "userTxTraffic")
	*t = SCIClientDetailsReport89trafficTrendDataTypeType(*tmpType)
	return nil
}

func (t *SCIClientDetailsReport89trafficTrendDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.UserRxTraffic != nil {
		tmp["userRxTraffic"] = t.UserRxTraffic
	}
	if t.UserTraffic != nil {
		tmp["userTraffic"] = t.UserTraffic
	}
	if t.UserTraffic24 != nil {
		tmp["userTraffic_2-4"] = t.UserTraffic24
	}
	if t.UserTraffic5 != nil {
		tmp["userTraffic_5"] = t.UserTraffic5
	}
	if t.UserTxTraffic != nil {
		tmp["userTxTraffic"] = t.UserTxTraffic
	}
	return json.Marshal(tmp)
}

func NewSCIClientDetailsReport89trafficTrendDataTypeType() *SCIClientDetailsReport89trafficTrendDataTypeType {
	m := new(SCIClientDetailsReport89trafficTrendDataTypeType)
	return m
}

// SCIClientDetailsReport92sessionsTableData
//
// Definition: ClientDetailsReport_ClientDetailsReport_92_sessionsTable_Data
type SCIClientDetailsReport92sessionsTableData []*SCIClientDetailsReport92sessionsTableDataType

func MakeSCIClientDetailsReport92sessionsTableData() SCIClientDetailsReport92sessionsTableData {
	m := make(SCIClientDetailsReport92sessionsTableData, 0)
	return m
}

// SCIClientDetailsReport92sessionsTableDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_92_sessionsTable_DataType
type SCIClientDetailsReport92sessionsTableDataType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	AuthMethod *string `json:"authMethod,omitempty"`

	DisconnectTime *string `json:"disconnectTime,omitempty"`

	FirstConnection *string `json:"firstConnection,omitempty"`

	Radio *string `json:"radio,omitempty"`

	RxBytes *float64 `json:"rxBytes,omitempty"`

	SessionDuration *float64 `json:"sessionDuration,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientDetailsReport92sessionsTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIClientDetailsReport92sessionsTableDataType SCIClientDetailsReport92sessionsTableDataType
	tmpType := new(_SCIClientDetailsReport92sessionsTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "authMethod")
	delete(tmpType.XAdditionalProperties, "disconnectTime")
	delete(tmpType.XAdditionalProperties, "firstConnection")
	delete(tmpType.XAdditionalProperties, "radio")
	delete(tmpType.XAdditionalProperties, "rxBytes")
	delete(tmpType.XAdditionalProperties, "sessionDuration")
	delete(tmpType.XAdditionalProperties, "ssid")
	delete(tmpType.XAdditionalProperties, "traffic")
	delete(tmpType.XAdditionalProperties, "txBytes")
	*t = SCIClientDetailsReport92sessionsTableDataType(*tmpType)
	return nil
}

func (t *SCIClientDetailsReport92sessionsTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ApMac != nil {
		tmp["apMac"] = t.ApMac
	}
	if t.ApName != nil {
		tmp["apName"] = t.ApName
	}
	if t.AuthMethod != nil {
		tmp["authMethod"] = t.AuthMethod
	}
	if t.DisconnectTime != nil {
		tmp["disconnectTime"] = t.DisconnectTime
	}
	if t.FirstConnection != nil {
		tmp["firstConnection"] = t.FirstConnection
	}
	if t.Radio != nil {
		tmp["radio"] = t.Radio
	}
	if t.RxBytes != nil {
		tmp["rxBytes"] = t.RxBytes
	}
	if t.SessionDuration != nil {
		tmp["sessionDuration"] = t.SessionDuration
	}
	if t.Ssid != nil {
		tmp["ssid"] = t.Ssid
	}
	if t.Traffic != nil {
		tmp["traffic"] = t.Traffic
	}
	if t.TxBytes != nil {
		tmp["txBytes"] = t.TxBytes
	}
	return json.Marshal(tmp)
}

func NewSCIClientDetailsReport92sessionsTableDataType() *SCIClientDetailsReport92sessionsTableDataType {
	m := new(SCIClientDetailsReport92sessionsTableDataType)
	return m
}

// SCIClientDetailsReport92sessionsTableMetaData
//
// Definition: ClientDetailsReport_ClientDetailsReport_92_sessionsTable_MetaData
type SCIClientDetailsReport92sessionsTableMetaData struct {
	MaxValues *SCIClientDetailsReport92sessionsTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientDetailsReport92sessionsTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIClientDetailsReport92sessionsTableMetaData SCIClientDetailsReport92sessionsTableMetaData
	tmpType := new(_SCIClientDetailsReport92sessionsTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	*t = SCIClientDetailsReport92sessionsTableMetaData(*tmpType)
	return nil
}

func (t *SCIClientDetailsReport92sessionsTableMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MaxValues != nil {
		tmp["maxValues"] = t.MaxValues
	}
	return json.Marshal(tmp)
}

func NewSCIClientDetailsReport92sessionsTableMetaData() *SCIClientDetailsReport92sessionsTableMetaData {
	m := new(SCIClientDetailsReport92sessionsTableMetaData)
	return m
}

// SCIClientDetailsReport92sessionsTableMetaDataMaxValuesType
//
// Definition: ClientDetailsReport_ClientDetailsReport_92_sessionsTable_MetaDataMaxValuesType
type SCIClientDetailsReport92sessionsTableMetaDataMaxValuesType struct {
	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`
}

func NewSCIClientDetailsReport92sessionsTableMetaDataMaxValuesType() *SCIClientDetailsReport92sessionsTableMetaDataMaxValuesType {
	m := new(SCIClientDetailsReport92sessionsTableMetaDataMaxValuesType)
	return m
}

// ReportClientDetailsReport7Top10ApplicationsByTrafficVolume
//
// Client Details Report - Top Applications by Traffic
//
// Operation ID: report_ClientDetailsReport_7_top10ApplicationsByTrafficVolume
// Operation path: /reports/12/sections/7/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport7Top10ApplicationsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportClientDetailsReport7Top10ApplicationsByTrafficVolume, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseTypeAPIResponse), err
}

// ReportClientDetailsReport8TopAppsByTrafficTable
//
// Client Details Report - Top Applications by Traffic
//
// Operation ID: report_ClientDetailsReport_8_topAppsByTrafficTable
// Operation path: /reports/12/sections/8/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport8TopAppsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportClientDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportClientDetailsReport8TopAppsByTrafficTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseTypeAPIResponse), err
}

// ReportClientDetailsReport82RssTrend
//
// Client Details Report - RSS Trend
//
// Operation ID: report_ClientDetailsReport_82_rssTrend
// Operation path: /reports/12/sections/82/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport82RssTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport82rssTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportClientDetailsReport82rssTrend200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportClientDetailsReport82rssTrend200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportClientDetailsReport82RssTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportClientDetailsReport82rssTrend200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportClientDetailsReport82rssTrend200ResponseTypeAPIResponse), err
}

// ReportClientDetailsReport83SnrTrend
//
// Client Details Report - SNR Trend
//
// Operation ID: report_ClientDetailsReport_83_snrTrend
// Operation path: /reports/12/sections/83/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport83SnrTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport83snrTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportClientDetailsReport83snrTrend200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportClientDetailsReport83snrTrend200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportClientDetailsReport83SnrTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportClientDetailsReport83snrTrend200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportClientDetailsReport83snrTrend200ResponseTypeAPIResponse), err
}

// ReportClientDetailsReport86Summary
//
// Client Details Report - Summary
//
// Operation ID: report_ClientDetailsReport_86_summary
// Operation path: /reports/12/sections/86/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport86Summary(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport86summary200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportClientDetailsReport86summary200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportClientDetailsReport86summary200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportClientDetailsReport86Summary, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportClientDetailsReport86summary200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportClientDetailsReport86summary200ResponseTypeAPIResponse), err
}

// ReportClientDetailsReport87ClientStats
//
// Client Details Report - Stats
//
// Operation ID: report_ClientDetailsReport_87_clientStats
// Operation path: /reports/12/sections/87/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport87ClientStats(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport87clientStats200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportClientDetailsReport87clientStats200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportClientDetailsReport87clientStats200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportClientDetailsReport87ClientStats, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportClientDetailsReport87clientStats200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportClientDetailsReport87clientStats200ResponseTypeAPIResponse), err
}

// ReportClientDetailsReport89TrafficTrend
//
// Client Details Report - Traffic Trend
//
// Operation ID: report_ClientDetailsReport_89_trafficTrend
// Operation path: /reports/12/sections/89/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport89TrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport89trafficTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportClientDetailsReport89trafficTrend200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportClientDetailsReport89trafficTrend200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportClientDetailsReport89TrafficTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportClientDetailsReport89trafficTrend200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportClientDetailsReport89trafficTrend200ResponseTypeAPIResponse), err
}

// ReportClientDetailsReport92SessionsTable
//
// Client Details Report - Sessions
//
// Operation ID: report_ClientDetailsReport_92_sessionsTable
// Operation path: /reports/12/sections/92/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport92SessionsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport92sessionsTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportClientDetailsReport92sessionsTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportClientDetailsReport92sessionsTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportClientDetailsReport92SessionsTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportClientDetailsReport92sessionsTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportClientDetailsReport92sessionsTable200ResponseTypeAPIResponse), err
}
