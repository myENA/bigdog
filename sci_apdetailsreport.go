package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIAPDetailsReportService struct {
	apiClient *SCIClient
}

func NewSCIAPDetailsReportService(c *SCIClient) *SCIAPDetailsReportService {
	s := new(SCIAPDetailsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIAPDetailsReportService() *SCIAPDetailsReportService {
	return NewSCIAPDetailsReportService(ss.apiClient)
}

// SCIAPDetailsReport5trendChartData
//
// Definition: APDetailsReport_APDetailsReport_5_trendChart_Data
type SCIAPDetailsReport5trendChartData [][]*SCIAPDetailsReport5trendChartDataTypeType

func MakeSCIAPDetailsReport5trendChartData() SCIAPDetailsReport5trendChartData {
	m := make(SCIAPDetailsReport5trendChartData, 0)
	return m
}

// SCIAPDetailsReport5trendChartDataType
//
// Definition: APDetailsReport_APDetailsReport_5_trendChart_DataType
type SCIAPDetailsReport5trendChartDataType []*SCIAPDetailsReport5trendChartDataTypeType

func MakeSCIAPDetailsReport5trendChartDataType() SCIAPDetailsReport5trendChartDataType {
	m := make(SCIAPDetailsReport5trendChartDataType, 0)
	return m
}

// SCIAPDetailsReport5trendChartDataTypeType
//
// Definition: APDetailsReport_APDetailsReport_5_trendChart_DataTypeType
type SCIAPDetailsReport5trendChartDataTypeType struct {
	AirtimeBusyAvg *float64 `json:"airtimeBusyAvg,omitempty"`

	AirtimeIdleAvg *float64 `json:"airtimeIdleAvg,omitempty"`

	AirtimeRxAvg *float64 `json:"airtimeRxAvg,omitempty"`

	AirtimeTxAvg *float64 `json:"airtimeTxAvg,omitempty"`

	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport5trendChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport5trendChartDataTypeType SCIAPDetailsReport5trendChartDataTypeType
	tmpType := new(_SCIAPDetailsReport5trendChartDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "airtimeBusyAvg")
	delete(tmpType.XAdditionalProperties, "airtimeIdleAvg")
	delete(tmpType.XAdditionalProperties, "airtimeRxAvg")
	delete(tmpType.XAdditionalProperties, "airtimeTxAvg")
	delete(tmpType.XAdditionalProperties, "airtimeUtilizationAvg")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCIAPDetailsReport5trendChartDataTypeType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport5trendChartDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AirtimeBusyAvg != nil {
		tmp["airtimeBusyAvg"] = t.AirtimeBusyAvg
	}
	if t.AirtimeIdleAvg != nil {
		tmp["airtimeIdleAvg"] = t.AirtimeIdleAvg
	}
	if t.AirtimeRxAvg != nil {
		tmp["airtimeRxAvg"] = t.AirtimeRxAvg
	}
	if t.AirtimeTxAvg != nil {
		tmp["airtimeTxAvg"] = t.AirtimeTxAvg
	}
	if t.AirtimeUtilizationAvg != nil {
		tmp["airtimeUtilizationAvg"] = t.AirtimeUtilizationAvg
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport5trendChartDataTypeType() *SCIAPDetailsReport5trendChartDataTypeType {
	m := new(SCIAPDetailsReport5trendChartDataTypeType)
	return m
}

// SCIAPDetailsReport8topAppsByTrafficTableData
//
// Definition: APDetailsReport_APDetailsReport_8_topAppsByTrafficTable_Data
type SCIAPDetailsReport8topAppsByTrafficTableData []*SCIAPDetailsReport8topAppsByTrafficTableDataType

func MakeSCIAPDetailsReport8topAppsByTrafficTableData() SCIAPDetailsReport8topAppsByTrafficTableData {
	m := make(SCIAPDetailsReport8topAppsByTrafficTableData, 0)
	return m
}

// SCIAPDetailsReport8topAppsByTrafficTableDataType
//
// Definition: APDetailsReport_APDetailsReport_8_topAppsByTrafficTable_DataType
type SCIAPDetailsReport8topAppsByTrafficTableDataType struct {
	App *string `json:"app,omitempty"`

	Index *float64 `json:"index,omitempty"`

	Port *string `json:"port,omitempty"`

	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport8topAppsByTrafficTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport8topAppsByTrafficTableDataType SCIAPDetailsReport8topAppsByTrafficTableDataType
	tmpType := new(_SCIAPDetailsReport8topAppsByTrafficTableDataType)
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
	*t = SCIAPDetailsReport8topAppsByTrafficTableDataType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport8topAppsByTrafficTableDataType) MarshalJSON() ([]byte, error) {
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

func NewSCIAPDetailsReport8topAppsByTrafficTableDataType() *SCIAPDetailsReport8topAppsByTrafficTableDataType {
	m := new(SCIAPDetailsReport8topAppsByTrafficTableDataType)
	return m
}

// SCIAPDetailsReport8topAppsByTrafficTableMetaData
//
// Definition: APDetailsReport_APDetailsReport_8_topAppsByTrafficTable_MetaData
type SCIAPDetailsReport8topAppsByTrafficTableMetaData struct {
	MaxValues *SCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport8topAppsByTrafficTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport8topAppsByTrafficTableMetaData SCIAPDetailsReport8topAppsByTrafficTableMetaData
	tmpType := new(_SCIAPDetailsReport8topAppsByTrafficTableMetaData)
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
	*t = SCIAPDetailsReport8topAppsByTrafficTableMetaData(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport8topAppsByTrafficTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIAPDetailsReport8topAppsByTrafficTableMetaData() *SCIAPDetailsReport8topAppsByTrafficTableMetaData {
	m := new(SCIAPDetailsReport8topAppsByTrafficTableMetaData)
	return m
}

// SCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType
//
// Definition: APDetailsReport_APDetailsReport_8_topAppsByTrafficTable_MetaDataMaxValuesType
type SCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType struct {
	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`
}

func NewSCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType() *SCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType {
	m := new(SCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCIAPDetailsReport14topTableData
//
// Definition: APDetailsReport_APDetailsReport_14_topTable_Data
type SCIAPDetailsReport14topTableData []*SCIAPDetailsReport14topTableDataType

func MakeSCIAPDetailsReport14topTableData() SCIAPDetailsReport14topTableData {
	m := make(SCIAPDetailsReport14topTableData, 0)
	return m
}

// SCIAPDetailsReport14topTableDataType
//
// Definition: APDetailsReport_APDetailsReport_14_topTable_DataType
type SCIAPDetailsReport14topTableDataType struct {
	ClientIp *string `json:"clientIp,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	Index *float64 `json:"index,omitempty"`

	Manufacturer *string `json:"manufacturer,omitempty"`

	OsType *string `json:"osType,omitempty"`

	RxBytes *float64 `json:"rxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	Username *string `json:"username,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport14topTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport14topTableDataType SCIAPDetailsReport14topTableDataType
	tmpType := new(_SCIAPDetailsReport14topTableDataType)
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
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "manufacturer")
	delete(tmpType.XAdditionalProperties, "osType")
	delete(tmpType.XAdditionalProperties, "rxBytes")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "traffic")
	delete(tmpType.XAdditionalProperties, "txBytes")
	delete(tmpType.XAdditionalProperties, "username")
	*t = SCIAPDetailsReport14topTableDataType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport14topTableDataType) MarshalJSON() ([]byte, error) {
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
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.Manufacturer != nil {
		tmp["manufacturer"] = t.Manufacturer
	}
	if t.OsType != nil {
		tmp["osType"] = t.OsType
	}
	if t.RxBytes != nil {
		tmp["rxBytes"] = t.RxBytes
	}
	if t.SessionCount != nil {
		tmp["sessionCount"] = t.SessionCount
	}
	if t.Traffic != nil {
		tmp["traffic"] = t.Traffic
	}
	if t.TxBytes != nil {
		tmp["txBytes"] = t.TxBytes
	}
	if t.Username != nil {
		tmp["username"] = t.Username
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport14topTableDataType() *SCIAPDetailsReport14topTableDataType {
	m := new(SCIAPDetailsReport14topTableDataType)
	return m
}

// SCIAPDetailsReport14topTableMetaData
//
// Definition: APDetailsReport_APDetailsReport_14_topTable_MetaData
type SCIAPDetailsReport14topTableMetaData struct {
	MaxValues *SCIAPDetailsReport14topTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport14topTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport14topTableMetaData SCIAPDetailsReport14topTableMetaData
	tmpType := new(_SCIAPDetailsReport14topTableMetaData)
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
	*t = SCIAPDetailsReport14topTableMetaData(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport14topTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIAPDetailsReport14topTableMetaData() *SCIAPDetailsReport14topTableMetaData {
	m := new(SCIAPDetailsReport14topTableMetaData)
	return m
}

// SCIAPDetailsReport14topTableMetaDataMaxValuesType
//
// Definition: APDetailsReport_APDetailsReport_14_topTable_MetaDataMaxValuesType
type SCIAPDetailsReport14topTableMetaDataMaxValuesType struct {
	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`
}

func NewSCIAPDetailsReport14topTableMetaDataMaxValuesType() *SCIAPDetailsReport14topTableMetaDataMaxValuesType {
	m := new(SCIAPDetailsReport14topTableMetaDataMaxValuesType)
	return m
}

// SCIAPDetailsReport15trendChartData
//
// Definition: APDetailsReport_APDetailsReport_15_trendChart_Data
type SCIAPDetailsReport15trendChartData [][]*SCIAPDetailsReport15trendChartDataTypeType

func MakeSCIAPDetailsReport15trendChartData() SCIAPDetailsReport15trendChartData {
	m := make(SCIAPDetailsReport15trendChartData, 0)
	return m
}

// SCIAPDetailsReport15trendChartDataType
//
// Definition: APDetailsReport_APDetailsReport_15_trendChart_DataType
type SCIAPDetailsReport15trendChartDataType []*SCIAPDetailsReport15trendChartDataTypeType

func MakeSCIAPDetailsReport15trendChartDataType() SCIAPDetailsReport15trendChartDataType {
	m := make(SCIAPDetailsReport15trendChartDataType, 0)
	return m
}

// SCIAPDetailsReport15trendChartDataTypeType
//
// Definition: APDetailsReport_APDetailsReport_15_trendChart_DataTypeType
type SCIAPDetailsReport15trendChartDataTypeType struct {
	End *string `json:"end,omitempty"`

	RxBytes *float64 `json:"rxBytes,omitempty"`

	Start *string `json:"start,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	Uniqueusers *float64 `json:"unique_users,omitempty"`

	Uniqueusers24 *float64 `json:"unique_users_2-4,omitempty"`

	Uniqueusers5 *float64 `json:"unique_users_5,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport15trendChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport15trendChartDataTypeType SCIAPDetailsReport15trendChartDataTypeType
	tmpType := new(_SCIAPDetailsReport15trendChartDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "rxBytes")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "traffic")
	delete(tmpType.XAdditionalProperties, "txBytes")
	delete(tmpType.XAdditionalProperties, "unique_users")
	delete(tmpType.XAdditionalProperties, "unique_users_2-4")
	delete(tmpType.XAdditionalProperties, "unique_users_5")
	*t = SCIAPDetailsReport15trendChartDataTypeType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport15trendChartDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.RxBytes != nil {
		tmp["rxBytes"] = t.RxBytes
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.Traffic != nil {
		tmp["traffic"] = t.Traffic
	}
	if t.TxBytes != nil {
		tmp["txBytes"] = t.TxBytes
	}
	if t.Uniqueusers != nil {
		tmp["unique_users"] = t.Uniqueusers
	}
	if t.Uniqueusers24 != nil {
		tmp["unique_users_2-4"] = t.Uniqueusers24
	}
	if t.Uniqueusers5 != nil {
		tmp["unique_users_5"] = t.Uniqueusers5
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport15trendChartDataTypeType() *SCIAPDetailsReport15trendChartDataTypeType {
	m := new(SCIAPDetailsReport15trendChartDataTypeType)
	return m
}

// SCIAPDetailsReport22trafficTrendData
//
// Definition: APDetailsReport_APDetailsReport_22_trafficTrend_Data
type SCIAPDetailsReport22trafficTrendData [][]*SCIAPDetailsReport22trafficTrendDataTypeType

func MakeSCIAPDetailsReport22trafficTrendData() SCIAPDetailsReport22trafficTrendData {
	m := make(SCIAPDetailsReport22trafficTrendData, 0)
	return m
}

// SCIAPDetailsReport22trafficTrendDataType
//
// Definition: APDetailsReport_APDetailsReport_22_trafficTrend_DataType
type SCIAPDetailsReport22trafficTrendDataType []*SCIAPDetailsReport22trafficTrendDataTypeType

func MakeSCIAPDetailsReport22trafficTrendDataType() SCIAPDetailsReport22trafficTrendDataType {
	m := make(SCIAPDetailsReport22trafficTrendDataType, 0)
	return m
}

// SCIAPDetailsReport22trafficTrendDataTypeType
//
// Definition: APDetailsReport_APDetailsReport_22_trafficTrend_DataTypeType
type SCIAPDetailsReport22trafficTrendDataTypeType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTraffic24 *float64 `json:"avgRateTotalTraffic_2-4,omitempty"`

	AvgRateTotalTraffic5 *float64 `json:"avgRateTotalTraffic_5,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	End *string `json:"end,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTraffic24 *float64 `json:"totalTraffic_2-4,omitempty"`

	TotalTraffic5 *float64 `json:"totalTraffic_5,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport22trafficTrendDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport22trafficTrendDataTypeType SCIAPDetailsReport22trafficTrendDataTypeType
	tmpType := new(_SCIAPDetailsReport22trafficTrendDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avgRateTotalRxTraffic")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTraffic")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTraffic_2-4")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTraffic_5")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTxTraffic")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "mgmtTraffic")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic_2-4")
	delete(tmpType.XAdditionalProperties, "totalTraffic_5")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	*t = SCIAPDetailsReport22trafficTrendDataTypeType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport22trafficTrendDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AvgRateTotalRxTraffic != nil {
		tmp["avgRateTotalRxTraffic"] = t.AvgRateTotalRxTraffic
	}
	if t.AvgRateTotalTraffic != nil {
		tmp["avgRateTotalTraffic"] = t.AvgRateTotalTraffic
	}
	if t.AvgRateTotalTraffic24 != nil {
		tmp["avgRateTotalTraffic_2-4"] = t.AvgRateTotalTraffic24
	}
	if t.AvgRateTotalTraffic5 != nil {
		tmp["avgRateTotalTraffic_5"] = t.AvgRateTotalTraffic5
	}
	if t.AvgRateTotalTxTraffic != nil {
		tmp["avgRateTotalTxTraffic"] = t.AvgRateTotalTxTraffic
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.MgmtTraffic != nil {
		tmp["mgmtTraffic"] = t.MgmtTraffic
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.TotalRxTraffic != nil {
		tmp["totalRxTraffic"] = t.TotalRxTraffic
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
	if t.TotalTxTraffic != nil {
		tmp["totalTxTraffic"] = t.TotalTxTraffic
	}
	if t.UserTraffic != nil {
		tmp["userTraffic"] = t.UserTraffic
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport22trafficTrendDataTypeType() *SCIAPDetailsReport22trafficTrendDataTypeType {
	m := new(SCIAPDetailsReport22trafficTrendDataTypeType)
	return m
}

// SCIAPDetailsReport40topSsidsByTrafficTableData
//
// Definition: APDetailsReport_APDetailsReport_40_topSsidsByTrafficTable_Data
type SCIAPDetailsReport40topSsidsByTrafficTableData []*SCIAPDetailsReport40topSsidsByTrafficTableDataType

func MakeSCIAPDetailsReport40topSsidsByTrafficTableData() SCIAPDetailsReport40topSsidsByTrafficTableData {
	m := make(SCIAPDetailsReport40topSsidsByTrafficTableData, 0)
	return m
}

// SCIAPDetailsReport40topSsidsByTrafficTableDataType
//
// Definition: APDetailsReport_APDetailsReport_40_topSsidsByTrafficTable_DataType
type SCIAPDetailsReport40topSsidsByTrafficTableDataType struct {
	ApCount *float64 `json:"apCount,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	Index *float64 `json:"index,omitempty"`

	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport40topSsidsByTrafficTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport40topSsidsByTrafficTableDataType SCIAPDetailsReport40topSsidsByTrafficTableDataType
	tmpType := new(_SCIAPDetailsReport40topSsidsByTrafficTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apCount")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "mgmtRxBytes")
	delete(tmpType.XAdditionalProperties, "mgmtTraffic")
	delete(tmpType.XAdditionalProperties, "mgmtTxBytes")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "ssid")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "userRxBytes")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	delete(tmpType.XAdditionalProperties, "userTxBytes")
	*t = SCIAPDetailsReport40topSsidsByTrafficTableDataType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport40topSsidsByTrafficTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ApCount != nil {
		tmp["apCount"] = t.ApCount
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.MgmtRxBytes != nil {
		tmp["mgmtRxBytes"] = t.MgmtRxBytes
	}
	if t.MgmtTraffic != nil {
		tmp["mgmtTraffic"] = t.MgmtTraffic
	}
	if t.MgmtTxBytes != nil {
		tmp["mgmtTxBytes"] = t.MgmtTxBytes
	}
	if t.SessionCount != nil {
		tmp["sessionCount"] = t.SessionCount
	}
	if t.Ssid != nil {
		tmp["ssid"] = t.Ssid
	}
	if t.TotalRxTraffic != nil {
		tmp["totalRxTraffic"] = t.TotalRxTraffic
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.TotalTxTraffic != nil {
		tmp["totalTxTraffic"] = t.TotalTxTraffic
	}
	if t.UserRxBytes != nil {
		tmp["userRxBytes"] = t.UserRxBytes
	}
	if t.UserTraffic != nil {
		tmp["userTraffic"] = t.UserTraffic
	}
	if t.UserTxBytes != nil {
		tmp["userTxBytes"] = t.UserTxBytes
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport40topSsidsByTrafficTableDataType() *SCIAPDetailsReport40topSsidsByTrafficTableDataType {
	m := new(SCIAPDetailsReport40topSsidsByTrafficTableDataType)
	return m
}

// SCIAPDetailsReport40topSsidsByTrafficTableMetaData
//
// Definition: APDetailsReport_APDetailsReport_40_topSsidsByTrafficTable_MetaData
type SCIAPDetailsReport40topSsidsByTrafficTableMetaData struct {
	MaxValues *SCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport40topSsidsByTrafficTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport40topSsidsByTrafficTableMetaData SCIAPDetailsReport40topSsidsByTrafficTableMetaData
	tmpType := new(_SCIAPDetailsReport40topSsidsByTrafficTableMetaData)
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
	*t = SCIAPDetailsReport40topSsidsByTrafficTableMetaData(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport40topSsidsByTrafficTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIAPDetailsReport40topSsidsByTrafficTableMetaData() *SCIAPDetailsReport40topSsidsByTrafficTableMetaData {
	m := new(SCIAPDetailsReport40topSsidsByTrafficTableMetaData)
	return m
}

// SCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType
//
// Definition: APDetailsReport_APDetailsReport_40_topSsidsByTrafficTable_MetaDataMaxValuesType
type SCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType struct {
	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`
}

func NewSCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType() *SCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType {
	m := new(SCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCIAPDetailsReport76apPerformanceData
//
// Definition: APDetailsReport_APDetailsReport_76_apPerformance_Data
type SCIAPDetailsReport76apPerformanceData [][]*SCIAPDetailsReport76apPerformanceDataTypeType

func MakeSCIAPDetailsReport76apPerformanceData() SCIAPDetailsReport76apPerformanceData {
	m := make(SCIAPDetailsReport76apPerformanceData, 0)
	return m
}

// SCIAPDetailsReport76apPerformanceDataType
//
// Definition: APDetailsReport_APDetailsReport_76_apPerformance_DataType
type SCIAPDetailsReport76apPerformanceDataType []*SCIAPDetailsReport76apPerformanceDataTypeType

func MakeSCIAPDetailsReport76apPerformanceDataType() SCIAPDetailsReport76apPerformanceDataType {
	m := make(SCIAPDetailsReport76apPerformanceDataType, 0)
	return m
}

// SCIAPDetailsReport76apPerformanceDataTypeType
//
// Definition: APDetailsReport_APDetailsReport_76_apPerformance_DataTypeType
type SCIAPDetailsReport76apPerformanceDataTypeType struct {
	Avg *float64 `json:"avg,omitempty"`

	Max *float64 `json:"max,omitempty"`

	MaxLabel *string `json:"maxLabel,omitempty"`

	Min *float64 `json:"min,omitempty"`

	MinLabel *string `json:"minLabel,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport76apPerformanceDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport76apPerformanceDataTypeType SCIAPDetailsReport76apPerformanceDataTypeType
	tmpType := new(_SCIAPDetailsReport76apPerformanceDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avg")
	delete(tmpType.XAdditionalProperties, "max")
	delete(tmpType.XAdditionalProperties, "maxLabel")
	delete(tmpType.XAdditionalProperties, "min")
	delete(tmpType.XAdditionalProperties, "minLabel")
	*t = SCIAPDetailsReport76apPerformanceDataTypeType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport76apPerformanceDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Avg != nil {
		tmp["avg"] = t.Avg
	}
	if t.Max != nil {
		tmp["max"] = t.Max
	}
	if t.MaxLabel != nil {
		tmp["maxLabel"] = t.MaxLabel
	}
	if t.Min != nil {
		tmp["min"] = t.Min
	}
	if t.MinLabel != nil {
		tmp["minLabel"] = t.MinLabel
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport76apPerformanceDataTypeType() *SCIAPDetailsReport76apPerformanceDataTypeType {
	m := new(SCIAPDetailsReport76apPerformanceDataTypeType)
	return m
}

// SCIAPDetailsReport78apStatsOverviewData
//
// Definition: APDetailsReport_APDetailsReport_78_apStatsOverview_Data
type SCIAPDetailsReport78apStatsOverviewData []*SCIAPDetailsReport78apStatsOverviewDataType

func MakeSCIAPDetailsReport78apStatsOverviewData() SCIAPDetailsReport78apStatsOverviewData {
	m := make(SCIAPDetailsReport78apStatsOverviewData, 0)
	return m
}

// SCIAPDetailsReport78apStatsOverviewDataType
//
// Definition: APDetailsReport_APDetailsReport_78_apStatsOverview_DataType
type SCIAPDetailsReport78apStatsOverviewDataType struct {
	TotalClientCount *float64 `json:"totalClientCount,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalSsids *float64 `json:"totalSsids,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport78apStatsOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport78apStatsOverviewDataType SCIAPDetailsReport78apStatsOverviewDataType
	tmpType := new(_SCIAPDetailsReport78apStatsOverviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "totalClientCount")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalSsids")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	*t = SCIAPDetailsReport78apStatsOverviewDataType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport78apStatsOverviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.TotalClientCount != nil {
		tmp["totalClientCount"] = t.TotalClientCount
	}
	if t.TotalRxTraffic != nil {
		tmp["totalRxTraffic"] = t.TotalRxTraffic
	}
	if t.TotalSsids != nil {
		tmp["totalSsids"] = t.TotalSsids
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.TotalTxTraffic != nil {
		tmp["totalTxTraffic"] = t.TotalTxTraffic
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport78apStatsOverviewDataType() *SCIAPDetailsReport78apStatsOverviewDataType {
	m := new(SCIAPDetailsReport78apStatsOverviewDataType)
	return m
}

// SCIAPDetailsReport79apUptimeHistoryData
//
// Definition: APDetailsReport_APDetailsReport_79_apUptimeHistory_Data
type SCIAPDetailsReport79apUptimeHistoryData []*SCIAPDetailsReport79apUptimeHistoryDataType

func MakeSCIAPDetailsReport79apUptimeHistoryData() SCIAPDetailsReport79apUptimeHistoryData {
	m := make(SCIAPDetailsReport79apUptimeHistoryData, 0)
	return m
}

// SCIAPDetailsReport79apUptimeHistoryDataType
//
// Definition: APDetailsReport_APDetailsReport_79_apUptimeHistory_DataType
type SCIAPDetailsReport79apUptimeHistoryDataType struct {
	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	Status *float64 `json:"status,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport79apUptimeHistoryDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport79apUptimeHistoryDataType SCIAPDetailsReport79apUptimeHistoryDataType
	tmpType := new(_SCIAPDetailsReport79apUptimeHistoryDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "status")
	*t = SCIAPDetailsReport79apUptimeHistoryDataType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport79apUptimeHistoryDataType) MarshalJSON() ([]byte, error) {
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
	if t.Status != nil {
		tmp["status"] = t.Status
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport79apUptimeHistoryDataType() *SCIAPDetailsReport79apUptimeHistoryDataType {
	m := new(SCIAPDetailsReport79apUptimeHistoryDataType)
	return m
}

// SCIAPDetailsReport79apUptimeHistoryMetaData
//
// Definition: APDetailsReport_APDetailsReport_79_apUptimeHistory_MetaData
type SCIAPDetailsReport79apUptimeHistoryMetaData struct {
	TotalDowntime *float64 `json:"totalDowntime,omitempty"`

	TotalUptime *float64 `json:"totalUptime,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport79apUptimeHistoryMetaData) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport79apUptimeHistoryMetaData SCIAPDetailsReport79apUptimeHistoryMetaData
	tmpType := new(_SCIAPDetailsReport79apUptimeHistoryMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "totalDowntime")
	delete(tmpType.XAdditionalProperties, "totalUptime")
	*t = SCIAPDetailsReport79apUptimeHistoryMetaData(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport79apUptimeHistoryMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.TotalDowntime != nil {
		tmp["totalDowntime"] = t.TotalDowntime
	}
	if t.TotalUptime != nil {
		tmp["totalUptime"] = t.TotalUptime
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport79apUptimeHistoryMetaData() *SCIAPDetailsReport79apUptimeHistoryMetaData {
	m := new(SCIAPDetailsReport79apUptimeHistoryMetaData)
	return m
}

// SCIAPDetailsReport81sessionsTableData
//
// Definition: APDetailsReport_APDetailsReport_81_sessionsTable_Data
type SCIAPDetailsReport81sessionsTableData []*SCIAPDetailsReport81sessionsTableDataType

func MakeSCIAPDetailsReport81sessionsTableData() SCIAPDetailsReport81sessionsTableData {
	m := make(SCIAPDetailsReport81sessionsTableData, 0)
	return m
}

// SCIAPDetailsReport81sessionsTableDataType
//
// Definition: APDetailsReport_APDetailsReport_81_sessionsTable_DataType
type SCIAPDetailsReport81sessionsTableDataType struct {
	AuthMethod *string `json:"authMethod,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	DisconnectTime *string `json:"disconnectTime,omitempty"`

	FirstConnection *string `json:"firstConnection,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	Radio *string `json:"radio,omitempty"`

	RxBytes *float64 `json:"rxBytes,omitempty"`

	SessionDuration *float64 `json:"sessionDuration,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport81sessionsTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport81sessionsTableDataType SCIAPDetailsReport81sessionsTableDataType
	tmpType := new(_SCIAPDetailsReport81sessionsTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "authMethod")
	delete(tmpType.XAdditionalProperties, "clientMac")
	delete(tmpType.XAdditionalProperties, "disconnectTime")
	delete(tmpType.XAdditionalProperties, "firstConnection")
	delete(tmpType.XAdditionalProperties, "hostname")
	delete(tmpType.XAdditionalProperties, "radio")
	delete(tmpType.XAdditionalProperties, "rxBytes")
	delete(tmpType.XAdditionalProperties, "sessionDuration")
	delete(tmpType.XAdditionalProperties, "ssid")
	delete(tmpType.XAdditionalProperties, "traffic")
	delete(tmpType.XAdditionalProperties, "txBytes")
	*t = SCIAPDetailsReport81sessionsTableDataType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport81sessionsTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AuthMethod != nil {
		tmp["authMethod"] = t.AuthMethod
	}
	if t.ClientMac != nil {
		tmp["clientMac"] = t.ClientMac
	}
	if t.DisconnectTime != nil {
		tmp["disconnectTime"] = t.DisconnectTime
	}
	if t.FirstConnection != nil {
		tmp["firstConnection"] = t.FirstConnection
	}
	if t.Hostname != nil {
		tmp["hostname"] = t.Hostname
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

func NewSCIAPDetailsReport81sessionsTableDataType() *SCIAPDetailsReport81sessionsTableDataType {
	m := new(SCIAPDetailsReport81sessionsTableDataType)
	return m
}

// SCIAPDetailsReport81sessionsTableMetaData
//
// Definition: APDetailsReport_APDetailsReport_81_sessionsTable_MetaData
type SCIAPDetailsReport81sessionsTableMetaData struct {
	MaxValues *SCIAPDetailsReport81sessionsTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport81sessionsTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport81sessionsTableMetaData SCIAPDetailsReport81sessionsTableMetaData
	tmpType := new(_SCIAPDetailsReport81sessionsTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	*t = SCIAPDetailsReport81sessionsTableMetaData(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport81sessionsTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIAPDetailsReport81sessionsTableMetaData() *SCIAPDetailsReport81sessionsTableMetaData {
	m := new(SCIAPDetailsReport81sessionsTableMetaData)
	return m
}

// SCIAPDetailsReport81sessionsTableMetaDataMaxValuesType
//
// Definition: APDetailsReport_APDetailsReport_81_sessionsTable_MetaDataMaxValuesType
type SCIAPDetailsReport81sessionsTableMetaDataMaxValuesType struct {
	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`
}

func NewSCIAPDetailsReport81sessionsTableMetaDataMaxValuesType() *SCIAPDetailsReport81sessionsTableMetaDataMaxValuesType {
	m := new(SCIAPDetailsReport81sessionsTableMetaDataMaxValuesType)
	return m
}

// SCIAPDetailsReport82rssTrendData
//
// Definition: APDetailsReport_APDetailsReport_82_rssTrend_Data
type SCIAPDetailsReport82rssTrendData []*SCIAPDetailsReport82rssTrendDataType

func MakeSCIAPDetailsReport82rssTrendData() SCIAPDetailsReport82rssTrendData {
	m := make(SCIAPDetailsReport82rssTrendData, 0)
	return m
}

// SCIAPDetailsReport82rssTrendDataType
//
// Definition: APDetailsReport_APDetailsReport_82_rssTrend_DataType
type SCIAPDetailsReport82rssTrendDataType struct {
	AvgRss *float64 `json:"avgRss,omitempty"`

	End *string `json:"end,omitempty"`

	MaxRss *float64 `json:"maxRss,omitempty"`

	MinRss *float64 `json:"minRss,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport82rssTrendDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport82rssTrendDataType SCIAPDetailsReport82rssTrendDataType
	tmpType := new(_SCIAPDetailsReport82rssTrendDataType)
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
	*t = SCIAPDetailsReport82rssTrendDataType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport82rssTrendDataType) MarshalJSON() ([]byte, error) {
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

func NewSCIAPDetailsReport82rssTrendDataType() *SCIAPDetailsReport82rssTrendDataType {
	m := new(SCIAPDetailsReport82rssTrendDataType)
	return m
}

// SCIAPDetailsReport83snrTrendData
//
// Definition: APDetailsReport_APDetailsReport_83_snrTrend_Data
type SCIAPDetailsReport83snrTrendData []*SCIAPDetailsReport83snrTrendDataType

func MakeSCIAPDetailsReport83snrTrendData() SCIAPDetailsReport83snrTrendData {
	m := make(SCIAPDetailsReport83snrTrendData, 0)
	return m
}

// SCIAPDetailsReport83snrTrendDataType
//
// Definition: APDetailsReport_APDetailsReport_83_snrTrend_DataType
type SCIAPDetailsReport83snrTrendDataType struct {
	AvgSnr *float64 `json:"avgSnr,omitempty"`

	End *string `json:"end,omitempty"`

	MaxSnr *float64 `json:"maxSnr,omitempty"`

	MinSnr *float64 `json:"minSnr,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport83snrTrendDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport83snrTrendDataType SCIAPDetailsReport83snrTrendDataType
	tmpType := new(_SCIAPDetailsReport83snrTrendDataType)
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
	*t = SCIAPDetailsReport83snrTrendDataType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport83snrTrendDataType) MarshalJSON() ([]byte, error) {
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

func NewSCIAPDetailsReport83snrTrendDataType() *SCIAPDetailsReport83snrTrendDataType {
	m := new(SCIAPDetailsReport83snrTrendDataType)
	return m
}

// SCIAPDetailsReport84alarmsTableData
//
// Definition: APDetailsReport_APDetailsReport_84_alarmsTable_Data
type SCIAPDetailsReport84alarmsTableData []*SCIAPDetailsReport84alarmsTableDataType

func MakeSCIAPDetailsReport84alarmsTableData() SCIAPDetailsReport84alarmsTableData {
	m := make(SCIAPDetailsReport84alarmsTableData, 0)
	return m
}

// SCIAPDetailsReport84alarmsTableDataType
//
// Definition: APDetailsReport_APDetailsReport_84_alarmsTable_DataType
type SCIAPDetailsReport84alarmsTableDataType struct {
	AlarmCode *string `json:"alarmCode,omitempty"`

	AlarmState *string `json:"alarmState,omitempty"`

	AlarmType *string `json:"alarmType,omitempty"`

	AlarmUUID *string `json:"alarmUUID,omitempty"`

	Category *string `json:"category,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Severity *string `json:"severity,omitempty"`

	Xtime *string `json:"__time,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport84alarmsTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport84alarmsTableDataType SCIAPDetailsReport84alarmsTableDataType
	tmpType := new(_SCIAPDetailsReport84alarmsTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "alarmCode")
	delete(tmpType.XAdditionalProperties, "alarmState")
	delete(tmpType.XAdditionalProperties, "alarmType")
	delete(tmpType.XAdditionalProperties, "alarmUUID")
	delete(tmpType.XAdditionalProperties, "category")
	delete(tmpType.XAdditionalProperties, "reason")
	delete(tmpType.XAdditionalProperties, "severity")
	delete(tmpType.XAdditionalProperties, "__time")
	*t = SCIAPDetailsReport84alarmsTableDataType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport84alarmsTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AlarmCode != nil {
		tmp["alarmCode"] = t.AlarmCode
	}
	if t.AlarmState != nil {
		tmp["alarmState"] = t.AlarmState
	}
	if t.AlarmType != nil {
		tmp["alarmType"] = t.AlarmType
	}
	if t.AlarmUUID != nil {
		tmp["alarmUUID"] = t.AlarmUUID
	}
	if t.Category != nil {
		tmp["category"] = t.Category
	}
	if t.Reason != nil {
		tmp["reason"] = t.Reason
	}
	if t.Severity != nil {
		tmp["severity"] = t.Severity
	}
	if t.Xtime != nil {
		tmp["__time"] = t.Xtime
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport84alarmsTableDataType() *SCIAPDetailsReport84alarmsTableDataType {
	m := new(SCIAPDetailsReport84alarmsTableDataType)
	return m
}

// SCIAPDetailsReport85eventsTableData
//
// Definition: APDetailsReport_APDetailsReport_85_eventsTable_Data
type SCIAPDetailsReport85eventsTableData []*SCIAPDetailsReport85eventsTableDataType

func MakeSCIAPDetailsReport85eventsTableData() SCIAPDetailsReport85eventsTableData {
	m := make(SCIAPDetailsReport85eventsTableData, 0)
	return m
}

// SCIAPDetailsReport85eventsTableDataType
//
// Definition: APDetailsReport_APDetailsReport_85_eventsTable_DataType
type SCIAPDetailsReport85eventsTableDataType struct {
	Category *string `json:"category,omitempty"`

	EventCode *string `json:"eventCode,omitempty"`

	EventType *string `json:"eventType,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Severity *string `json:"severity,omitempty"`

	Xtime *string `json:"__time,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPDetailsReport85eventsTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPDetailsReport85eventsTableDataType SCIAPDetailsReport85eventsTableDataType
	tmpType := new(_SCIAPDetailsReport85eventsTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "category")
	delete(tmpType.XAdditionalProperties, "eventCode")
	delete(tmpType.XAdditionalProperties, "eventType")
	delete(tmpType.XAdditionalProperties, "reason")
	delete(tmpType.XAdditionalProperties, "severity")
	delete(tmpType.XAdditionalProperties, "__time")
	*t = SCIAPDetailsReport85eventsTableDataType(*tmpType)
	return nil
}

func (t *SCIAPDetailsReport85eventsTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Category != nil {
		tmp["category"] = t.Category
	}
	if t.EventCode != nil {
		tmp["eventCode"] = t.EventCode
	}
	if t.EventType != nil {
		tmp["eventType"] = t.EventType
	}
	if t.Reason != nil {
		tmp["reason"] = t.Reason
	}
	if t.Severity != nil {
		tmp["severity"] = t.Severity
	}
	if t.Xtime != nil {
		tmp["__time"] = t.Xtime
	}
	return json.Marshal(tmp)
}

func NewSCIAPDetailsReport85eventsTableDataType() *SCIAPDetailsReport85eventsTableDataType {
	m := new(SCIAPDetailsReport85eventsTableDataType)
	return m
}

// ReportAPDetailsReport5TrendChart
//
// Operation ID: report_APDetailsReport_5_trendChart
//
// AP Details Report - Airtime Utilization Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport5TrendChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport5trendChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport5trendChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport5TrendChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport5trendChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport7Top10ApplicationsByTrafficVolume
//
// Operation ID: report_APDetailsReport_7_top10ApplicationsByTrafficVolume
//
// AP Details Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport7Top10ApplicationsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport7Top10ApplicationsByTrafficVolume, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport8TopAppsByTrafficTable
//
// Operation ID: report_APDetailsReport_8_topAppsByTrafficTable
//
// AP Details Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport8TopAppsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport8TopAppsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport14TopTable
//
// Operation ID: report_APDetailsReport_14_topTable
//
// AP Details Report - Clients Details
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport14TopTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport14topTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport14topTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport14TopTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport14topTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport15TrendChart
//
// Operation ID: report_APDetailsReport_15_trendChart
//
// AP Details Report - Unique Clients Trend Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport15TrendChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport15trendChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport15trendChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport15TrendChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport15trendChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport22TrafficTrend
//
// Operation ID: report_APDetailsReport_22_trafficTrend
//
// AP Details Report - Traffic Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport22TrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport22trafficTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport22trafficTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport22TrafficTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport22trafficTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport40TopSsidsByTrafficTable
//
// Operation ID: report_APDetailsReport_40_topSsidsByTrafficTable
//
// AP Details Report - Top SSIDs by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport40TopSsidsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport40TopSsidsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport75ApSummary
//
// Operation ID: report_APDetailsReport_75_apSummary
//
// AP Details Report - Summary
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport75ApSummary(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport75apSummary200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport75apSummary200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport75ApSummary, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport75apSummary200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport76ApPerformance
//
// Operation ID: report_APDetailsReport_76_apPerformance
//
// AP Details Report - Performance
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport76ApPerformance(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport76apPerformance200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport76apPerformance200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport76ApPerformance, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport76apPerformance200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport77ApDetails
//
// Operation ID: report_APDetailsReport_77_apDetails
//
// AP Details Report - Details
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport77ApDetails(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport77apDetails200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport77apDetails200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport77ApDetails, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport77apDetails200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport78ApStatsOverview
//
// Operation ID: report_APDetailsReport_78_apStatsOverview
//
// AP Details Report - Stats
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport78ApStatsOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport78apStatsOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport78apStatsOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport78ApStatsOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport78apStatsOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport79ApUptimeHistory
//
// Operation ID: report_APDetailsReport_79_apUptimeHistory
//
// AP Details Report - Uptime History
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport79ApUptimeHistory(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport79apUptimeHistory200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport79apUptimeHistory200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport79ApUptimeHistory, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport79apUptimeHistory200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport80Top10ClientsByTrafficVolume
//
// Operation ID: report_APDetailsReport_80_top10ClientsByTrafficVolume
//
// AP Details Report - Top 10 Clients by Traffic Volume
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport80Top10ClientsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport80Top10ClientsByTrafficVolume, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport81SessionsTable
//
// Operation ID: report_APDetailsReport_81_sessionsTable
//
// AP Details Report - Sessions
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport81SessionsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport81sessionsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport81sessionsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport81SessionsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport81sessionsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport82RssTrend
//
// Operation ID: report_APDetailsReport_82_rssTrend
//
// AP Details Report - RSS Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport82RssTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport82rssTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport82rssTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport82RssTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport82rssTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport83SnrTrend
//
// Operation ID: report_APDetailsReport_83_snrTrend
//
// AP Details Report - SNR Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport83SnrTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport83snrTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport83snrTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport83SnrTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport83snrTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport84AlarmsTable
//
// Operation ID: report_APDetailsReport_84_alarmsTable
//
// AP Details Report - Alarms
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport84AlarmsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport84alarmsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport84alarmsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport84AlarmsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport84alarmsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport85EventsTable
//
// Operation ID: report_APDetailsReport_85_eventsTable
//
// AP Details Report - Events
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport85EventsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport85eventsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport85eventsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport85EventsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport85eventsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport95Anomalies
//
// Operation ID: report_APDetailsReport_95_anomalies
//
// AP Details Report - Anomalies
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport95Anomalies(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport95anomalies200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport95anomalies200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport95Anomalies, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport95anomalies200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport110ApAnomaly
//
// Operation ID: report_APDetailsReport_110_apAnomaly
//
// AP Details Report - Anomalies for the Past 30 Days
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport110ApAnomaly(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport110apAnomaly200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport110apAnomaly200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport110ApAnomaly, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport110apAnomaly200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
