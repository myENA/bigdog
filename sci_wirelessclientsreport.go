package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIWirelessClientsReportService struct {
	apiClient *SCIClient
}

func NewSCIWirelessClientsReportService(c *SCIClient) *SCIWirelessClientsReportService {
	s := new(SCIWirelessClientsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIWirelessClientsReportService() *SCIWirelessClientsReportService {
	return NewSCIWirelessClientsReportService(ss.apiClient)
}

// SCIWirelessClientsReport12overviewData
//
// Definition: WirelessClientsReport_WirelessClientsReport_12_overview_Data
type SCIWirelessClientsReport12overviewData []*SCIWirelessClientsReport12overviewDataType

func MakeSCIWirelessClientsReport12overviewData() SCIWirelessClientsReport12overviewData {
	m := make(SCIWirelessClientsReport12overviewData, 0)
	return m
}

// SCIWirelessClientsReport12overviewDataType
//
// Definition: WirelessClientsReport_WirelessClientsReport_12_overview_DataType
type SCIWirelessClientsReport12overviewDataType struct {
	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UniqueUsers *float64 `json:"uniqueUsers,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessClientsReport12overviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport12overviewDataType SCIWirelessClientsReport12overviewDataType
	tmpType := new(_SCIWirelessClientsReport12overviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "uniqueUsers")
	*t = SCIWirelessClientsReport12overviewDataType(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport12overviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
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
	if t.UniqueUsers != nil {
		tmp["uniqueUsers"] = t.UniqueUsers
	}
	return json.Marshal(tmp)
}

func NewSCIWirelessClientsReport12overviewDataType() *SCIWirelessClientsReport12overviewDataType {
	m := new(SCIWirelessClientsReport12overviewDataType)
	return m
}

// SCIWirelessClientsReport14topTableData
//
// Definition: WirelessClientsReport_WirelessClientsReport_14_topTable_Data
type SCIWirelessClientsReport14topTableData []*SCIWirelessClientsReport14topTableDataType

func MakeSCIWirelessClientsReport14topTableData() SCIWirelessClientsReport14topTableData {
	m := make(SCIWirelessClientsReport14topTableData, 0)
	return m
}

// SCIWirelessClientsReport14topTableDataType
//
// Definition: WirelessClientsReport_WirelessClientsReport_14_topTable_DataType
type SCIWirelessClientsReport14topTableDataType struct {
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

func (t *SCIWirelessClientsReport14topTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport14topTableDataType SCIWirelessClientsReport14topTableDataType
	tmpType := new(_SCIWirelessClientsReport14topTableDataType)
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
	*t = SCIWirelessClientsReport14topTableDataType(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport14topTableDataType) MarshalJSON() ([]byte, error) {
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

func NewSCIWirelessClientsReport14topTableDataType() *SCIWirelessClientsReport14topTableDataType {
	m := new(SCIWirelessClientsReport14topTableDataType)
	return m
}

// SCIWirelessClientsReport14topTableMetaData
//
// Definition: WirelessClientsReport_WirelessClientsReport_14_topTable_MetaData
type SCIWirelessClientsReport14topTableMetaData struct {
	MaxValues *SCIWirelessClientsReport14topTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessClientsReport14topTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport14topTableMetaData SCIWirelessClientsReport14topTableMetaData
	tmpType := new(_SCIWirelessClientsReport14topTableMetaData)
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
	*t = SCIWirelessClientsReport14topTableMetaData(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport14topTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIWirelessClientsReport14topTableMetaData() *SCIWirelessClientsReport14topTableMetaData {
	m := new(SCIWirelessClientsReport14topTableMetaData)
	return m
}

// SCIWirelessClientsReport14topTableMetaDataMaxValuesType
//
// Definition: WirelessClientsReport_WirelessClientsReport_14_topTable_MetaDataMaxValuesType
type SCIWirelessClientsReport14topTableMetaDataMaxValuesType struct {
	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`
}

func NewSCIWirelessClientsReport14topTableMetaDataMaxValuesType() *SCIWirelessClientsReport14topTableMetaDataMaxValuesType {
	m := new(SCIWirelessClientsReport14topTableMetaDataMaxValuesType)
	return m
}

// SCIWirelessClientsReport15trendChartData
//
// Definition: WirelessClientsReport_WirelessClientsReport_15_trendChart_Data
type SCIWirelessClientsReport15trendChartData [][]*SCIWirelessClientsReport15trendChartDataTypeType

func MakeSCIWirelessClientsReport15trendChartData() SCIWirelessClientsReport15trendChartData {
	m := make(SCIWirelessClientsReport15trendChartData, 0)
	return m
}

// SCIWirelessClientsReport15trendChartDataType
//
// Definition: WirelessClientsReport_WirelessClientsReport_15_trendChart_DataType
type SCIWirelessClientsReport15trendChartDataType []*SCIWirelessClientsReport15trendChartDataTypeType

func MakeSCIWirelessClientsReport15trendChartDataType() SCIWirelessClientsReport15trendChartDataType {
	m := make(SCIWirelessClientsReport15trendChartDataType, 0)
	return m
}

// SCIWirelessClientsReport15trendChartDataTypeType
//
// Definition: WirelessClientsReport_WirelessClientsReport_15_trendChart_DataTypeType
type SCIWirelessClientsReport15trendChartDataTypeType struct {
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

func (t *SCIWirelessClientsReport15trendChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport15trendChartDataTypeType SCIWirelessClientsReport15trendChartDataTypeType
	tmpType := new(_SCIWirelessClientsReport15trendChartDataTypeType)
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
	*t = SCIWirelessClientsReport15trendChartDataTypeType(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport15trendChartDataTypeType) MarshalJSON() ([]byte, error) {
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

func NewSCIWirelessClientsReport15trendChartDataTypeType() *SCIWirelessClientsReport15trendChartDataTypeType {
	m := new(SCIWirelessClientsReport15trendChartDataTypeType)
	return m
}

// SCIWirelessClientsReport16trendTableData
//
// Definition: WirelessClientsReport_WirelessClientsReport_16_trendTable_Data
type SCIWirelessClientsReport16trendTableData []*SCIWirelessClientsReport16trendTableDataType

func MakeSCIWirelessClientsReport16trendTableData() SCIWirelessClientsReport16trendTableData {
	m := make(SCIWirelessClientsReport16trendTableData, 0)
	return m
}

// SCIWirelessClientsReport16trendTableDataType
//
// Definition: WirelessClientsReport_WirelessClientsReport_16_trendTable_DataType
type SCIWirelessClientsReport16trendTableDataType struct {
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

func (t *SCIWirelessClientsReport16trendTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport16trendTableDataType SCIWirelessClientsReport16trendTableDataType
	tmpType := new(_SCIWirelessClientsReport16trendTableDataType)
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
	*t = SCIWirelessClientsReport16trendTableDataType(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport16trendTableDataType) MarshalJSON() ([]byte, error) {
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

func NewSCIWirelessClientsReport16trendTableDataType() *SCIWirelessClientsReport16trendTableDataType {
	m := new(SCIWirelessClientsReport16trendTableDataType)
	return m
}

// SCIWirelessClientsReport16trendTableMetaData
//
// Definition: WirelessClientsReport_WirelessClientsReport_16_trendTable_MetaData
type SCIWirelessClientsReport16trendTableMetaData struct {
	MaxValues *SCIWirelessClientsReport16trendTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessClientsReport16trendTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport16trendTableMetaData SCIWirelessClientsReport16trendTableMetaData
	tmpType := new(_SCIWirelessClientsReport16trendTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	*t = SCIWirelessClientsReport16trendTableMetaData(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport16trendTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIWirelessClientsReport16trendTableMetaData() *SCIWirelessClientsReport16trendTableMetaData {
	m := new(SCIWirelessClientsReport16trendTableMetaData)
	return m
}

// SCIWirelessClientsReport16trendTableMetaDataMaxValuesType
//
// Definition: WirelessClientsReport_WirelessClientsReport_16_trendTable_MetaDataMaxValuesType
type SCIWirelessClientsReport16trendTableMetaDataMaxValuesType struct {
	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	Uniqueusers *float64 `json:"unique_users,omitempty"`

	Uniqueusers24 *float64 `json:"unique_users_2-4,omitempty"`

	Uniqueusers5 *float64 `json:"unique_users_5,omitempty"`
}

func NewSCIWirelessClientsReport16trendTableMetaDataMaxValuesType() *SCIWirelessClientsReport16trendTableMetaDataMaxValuesType {
	m := new(SCIWirelessClientsReport16trendTableMetaDataMaxValuesType)
	return m
}

// SCIWirelessClientsReport17topPercentileData
//
// Definition: WirelessClientsReport_WirelessClientsReport_17_topPercentile_Data
type SCIWirelessClientsReport17topPercentileData [][]*SCIWirelessClientsReport17topPercentileDataTypeType

func MakeSCIWirelessClientsReport17topPercentileData() SCIWirelessClientsReport17topPercentileData {
	m := make(SCIWirelessClientsReport17topPercentileData, 0)
	return m
}

// SCIWirelessClientsReport17topPercentileDataType
//
// Definition: WirelessClientsReport_WirelessClientsReport_17_topPercentile_DataType
type SCIWirelessClientsReport17topPercentileDataType []*SCIWirelessClientsReport17topPercentileDataTypeType

func MakeSCIWirelessClientsReport17topPercentileDataType() SCIWirelessClientsReport17topPercentileDataType {
	m := make(SCIWirelessClientsReport17topPercentileDataType, 0)
	return m
}

// SCIWirelessClientsReport17topPercentileDataTypeType
//
// Definition: WirelessClientsReport_WirelessClientsReport_17_topPercentile_DataTypeType
type SCIWirelessClientsReport17topPercentileDataTypeType struct {
	ClientCount *float64 `json:"clientCount,omitempty"`

	MaxTraffic *float64 `json:"maxTraffic,omitempty"`

	Percent *string `json:"percent,omitempty"`

	PercentRank *float64 `json:"percentRank,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessClientsReport17topPercentileDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport17topPercentileDataTypeType SCIWirelessClientsReport17topPercentileDataTypeType
	tmpType := new(_SCIWirelessClientsReport17topPercentileDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "maxTraffic")
	delete(tmpType.XAdditionalProperties, "percent")
	delete(tmpType.XAdditionalProperties, "percentRank")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	*t = SCIWirelessClientsReport17topPercentileDataTypeType(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport17topPercentileDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.MaxTraffic != nil {
		tmp["maxTraffic"] = t.MaxTraffic
	}
	if t.Percent != nil {
		tmp["percent"] = t.Percent
	}
	if t.PercentRank != nil {
		tmp["percentRank"] = t.PercentRank
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	return json.Marshal(tmp)
}

func NewSCIWirelessClientsReport17topPercentileDataTypeType() *SCIWirelessClientsReport17topPercentileDataTypeType {
	m := new(SCIWirelessClientsReport17topPercentileDataTypeType)
	return m
}

// SCIWirelessClientsReport18topNOSByClientCountData
//
// Definition: WirelessClientsReport_WirelessClientsReport_18_topNOSByClientCount_Data
type SCIWirelessClientsReport18topNOSByClientCountData [][]*SCIWirelessClientsReport18topNOSByClientCountDataTypeType

func MakeSCIWirelessClientsReport18topNOSByClientCountData() SCIWirelessClientsReport18topNOSByClientCountData {
	m := make(SCIWirelessClientsReport18topNOSByClientCountData, 0)
	return m
}

// SCIWirelessClientsReport18topNOSByClientCountDataType
//
// Definition: WirelessClientsReport_WirelessClientsReport_18_topNOSByClientCount_DataType
type SCIWirelessClientsReport18topNOSByClientCountDataType []*SCIWirelessClientsReport18topNOSByClientCountDataTypeType

func MakeSCIWirelessClientsReport18topNOSByClientCountDataType() SCIWirelessClientsReport18topNOSByClientCountDataType {
	m := make(SCIWirelessClientsReport18topNOSByClientCountDataType, 0)
	return m
}

// SCIWirelessClientsReport18topNOSByClientCountDataTypeType
//
// Definition: WirelessClientsReport_WirelessClientsReport_18_topNOSByClientCount_DataTypeType
type SCIWirelessClientsReport18topNOSByClientCountDataTypeType struct {
	ClientCount *float64 `json:"clientCount,omitempty"`

	End *string `json:"end,omitempty"`

	OsType *string `json:"osType,omitempty"`

	Start *string `json:"start,omitempty"`

	Unknown *float64 `json:"Unknown,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessClientsReport18topNOSByClientCountDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport18topNOSByClientCountDataTypeType SCIWirelessClientsReport18topNOSByClientCountDataTypeType
	tmpType := new(_SCIWirelessClientsReport18topNOSByClientCountDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "osType")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "Unknown")
	*t = SCIWirelessClientsReport18topNOSByClientCountDataTypeType(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport18topNOSByClientCountDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.OsType != nil {
		tmp["osType"] = t.OsType
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.Unknown != nil {
		tmp["Unknown"] = t.Unknown
	}
	return json.Marshal(tmp)
}

func NewSCIWirelessClientsReport18topNOSByClientCountDataTypeType() *SCIWirelessClientsReport18topNOSByClientCountDataTypeType {
	m := new(SCIWirelessClientsReport18topNOSByClientCountDataTypeType)
	return m
}

// SCIWirelessClientsReport18topNOSByClientCountMetaData
//
// Definition: WirelessClientsReport_WirelessClientsReport_18_topNOSByClientCount_MetaData
type SCIWirelessClientsReport18topNOSByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessClientsReport18topNOSByClientCountMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport18topNOSByClientCountMetaData SCIWirelessClientsReport18topNOSByClientCountMetaData
	tmpType := new(_SCIWirelessClientsReport18topNOSByClientCountMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	delete(tmpType.XAdditionalProperties, "totalClients")
	*t = SCIWirelessClientsReport18topNOSByClientCountMetaData(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport18topNOSByClientCountMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ColorKeys != nil {
		tmp["colorKeys"] = t.ColorKeys
	}
	if t.TotalClients != nil {
		tmp["totalClients"] = t.TotalClients
	}
	return json.Marshal(tmp)
}

func NewSCIWirelessClientsReport18topNOSByClientCountMetaData() *SCIWirelessClientsReport18topNOSByClientCountMetaData {
	m := new(SCIWirelessClientsReport18topNOSByClientCountMetaData)
	return m
}

// SCIWirelessClientsReport19top10ManufacturersByClientCountData
//
// Definition: WirelessClientsReport_WirelessClientsReport_19_top10ManufacturersByClientCount_Data
type SCIWirelessClientsReport19top10ManufacturersByClientCountData [][]*SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType

func MakeSCIWirelessClientsReport19top10ManufacturersByClientCountData() SCIWirelessClientsReport19top10ManufacturersByClientCountData {
	m := make(SCIWirelessClientsReport19top10ManufacturersByClientCountData, 0)
	return m
}

// SCIWirelessClientsReport19top10ManufacturersByClientCountDataType
//
// Definition: WirelessClientsReport_WirelessClientsReport_19_top10ManufacturersByClientCount_DataType
type SCIWirelessClientsReport19top10ManufacturersByClientCountDataType []*SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType

func MakeSCIWirelessClientsReport19top10ManufacturersByClientCountDataType() SCIWirelessClientsReport19top10ManufacturersByClientCountDataType {
	m := make(SCIWirelessClientsReport19top10ManufacturersByClientCountDataType, 0)
	return m
}

// SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType
//
// Definition: WirelessClientsReport_WirelessClientsReport_19_top10ManufacturersByClientCount_DataTypeType
type SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType struct {
	ClientCount *float64 `json:"clientCount,omitempty"`

	End *string `json:"end,omitempty"`

	Manufacturer *string `json:"manufacturer,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType
	tmpType := new(_SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "manufacturer")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Manufacturer != nil {
		tmp["manufacturer"] = t.Manufacturer
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType() *SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType {
	m := new(SCIWirelessClientsReport19top10ManufacturersByClientCountDataTypeType)
	return m
}

// SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData
//
// Definition: WirelessClientsReport_WirelessClientsReport_19_top10ManufacturersByClientCount_MetaData
type SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData
	tmpType := new(_SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	delete(tmpType.XAdditionalProperties, "totalClients")
	*t = SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ColorKeys != nil {
		tmp["colorKeys"] = t.ColorKeys
	}
	if t.TotalClients != nil {
		tmp["totalClients"] = t.TotalClients
	}
	return json.Marshal(tmp)
}

func NewSCIWirelessClientsReport19top10ManufacturersByClientCountMetaData() *SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData {
	m := new(SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData)
	return m
}

// SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData
//
// Definition: WirelessClientsReport_WirelessClientsReport_112_top10AuthenticationMechanismByClientCount_Data
type SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData [][]*SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType

func MakeSCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData() SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData {
	m := make(SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData, 0)
	return m
}

// SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType
//
// Definition: WirelessClientsReport_WirelessClientsReport_112_top10AuthenticationMechanismByClientCount_DataType
type SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType []*SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType

func MakeSCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType() SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType {
	m := make(SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType, 0)
	return m
}

// SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType
//
// Definition: WirelessClientsReport_WirelessClientsReport_112_top10AuthenticationMechanismByClientCount_DataTypeType
type SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType struct {
	AuthMethod *string `json:"authMethod,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	End *string `json:"end,omitempty"`

	Open *float64 `json:"Open,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType
	tmpType := new(_SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "authMethod")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "Open")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Open != nil {
		tmp["Open"] = t.Open
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType() *SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType {
	m := new(SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataTypeType)
	return m
}

// SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData
//
// Definition: WirelessClientsReport_WirelessClientsReport_112_top10AuthenticationMechanismByClientCount_MetaData
type SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData
	tmpType := new(_SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	delete(tmpType.XAdditionalProperties, "totalClients")
	*t = SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData(*tmpType)
	return nil
}

func (t *SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ColorKeys != nil {
		tmp["colorKeys"] = t.ColorKeys
	}
	if t.TotalClients != nil {
		tmp["totalClients"] = t.TotalClients
	}
	return json.Marshal(tmp)
}

func NewSCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData() *SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData {
	m := new(SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData)
	return m
}

// ReportWirelessClientsReport12Overview
//
// Operation ID: report_WirelessClientsReport_12_overview
//
// Wireless - Clients Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport12Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport12overview200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport12overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessClientsReport12Overview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport12overview200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport13TopChart
//
// Operation ID: report_WirelessClientsReport_13_topChart
//
// Wireless - Clients Report - Top 10 Unique Clients by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport13TopChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport13topChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport13topChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessClientsReport13TopChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport13topChart200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport14TopTable
//
// Operation ID: report_WirelessClientsReport_14_topTable
//
// Wireless - Clients Report - Clients Details
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport14TopTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport14topTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport14topTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessClientsReport14TopTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport14topTable200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport15TrendChart
//
// Operation ID: report_WirelessClientsReport_15_trendChart
//
// Wireless - Clients Report - Unique Clients Trend Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport15TrendChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport15trendChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport15trendChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessClientsReport15TrendChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport15trendChart200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport16TrendTable
//
// Operation ID: report_WirelessClientsReport_16_trendTable
//
// Wireless - Clients Report - Unique Clients Trend Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport16TrendTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport16trendTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport16trendTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessClientsReport16TrendTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport16trendTable200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport17TopPercentile
//
// Operation ID: report_WirelessClientsReport_17_topPercentile
//
// Wireless - Clients Report - Top Clients by Traffic Percentile
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport17TopPercentile(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport17topPercentile200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport17topPercentile200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessClientsReport17TopPercentile, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport17topPercentile200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport18TopNOSByClientCount
//
// Operation ID: report_WirelessClientsReport_18_topNOSByClientCount
//
// Wireless - Clients Report - Top 10 OS by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport18TopNOSByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport18topNOSByClientCount200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessClientsReport18TopNOSByClientCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport18topNOSByClientCount200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport19Top10ManufacturersByClientCount
//
// Operation ID: report_WirelessClientsReport_19_top10ManufacturersByClientCount
//
// Wireless - Clients Report - Top 10 Manufacturers by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport19Top10ManufacturersByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessClientsReport19Top10ManufacturersByClientCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport112Top10AuthenticationMechanismByClientCount
//
// Operation ID: report_WirelessClientsReport_112_top10AuthenticationMechanismByClientCount
//
// Wireless - Clients Report - Top 10 Authentication Methods by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport112Top10AuthenticationMechanismByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessClientsReport112Top10AuthenticationMechanismByClientCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
