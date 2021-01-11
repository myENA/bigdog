package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCINetworkWirelessReportService struct {
	apiClient *SCIClient
}

func NewSCINetworkWirelessReportService(c *SCIClient) *SCINetworkWirelessReportService {
	s := new(SCINetworkWirelessReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCINetworkWirelessReportService() *SCINetworkWirelessReportService {
	return NewSCINetworkWirelessReportService(ss.apiClient)
}

// SCINetworkWirelessReport20overviewData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_20_overview_Data
type SCINetworkWirelessReport20overviewData []*SCINetworkWirelessReport20overviewDataType

func MakeSCINetworkWirelessReport20overviewData() SCINetworkWirelessReport20overviewData {
	m := make(SCINetworkWirelessReport20overviewData, 0)
	return m
}

// SCINetworkWirelessReport20overviewDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_20_overview_DataType
type SCINetworkWirelessReport20overviewDataType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	TotalAps *float64 `json:"totalAps,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UniqueUsers *float64 `json:"uniqueUsers,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport20overviewDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport20overviewDataType SCINetworkWirelessReport20overviewDataType
	tmpType := new(_SCINetworkWirelessReport20overviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avgRateTotalRxTraffic")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTraffic")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTxTraffic")
	delete(tmpType.XAdditionalProperties, "totalAps")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "uniqueUsers")
	*t = SCINetworkWirelessReport20overviewDataType(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport20overviewDataType) MarshalJSON() ([]byte, error) {
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
	if t.AvgRateTotalTxTraffic != nil {
		tmp["avgRateTotalTxTraffic"] = t.AvgRateTotalTxTraffic
	}
	if t.TotalAps != nil {
		tmp["totalAps"] = t.TotalAps
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

func NewSCINetworkWirelessReport20overviewDataType() *SCINetworkWirelessReport20overviewDataType {
	m := new(SCINetworkWirelessReport20overviewDataType)
	return m
}

// SCINetworkWirelessReport21trafficDistributionData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_21_trafficDistribution_Data
type SCINetworkWirelessReport21trafficDistributionData []*SCINetworkWirelessReport21trafficDistributionDataType

func MakeSCINetworkWirelessReport21trafficDistributionData() SCINetworkWirelessReport21trafficDistributionData {
	m := make(SCINetworkWirelessReport21trafficDistributionData, 0)
	return m
}

// SCINetworkWirelessReport21trafficDistributionDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_21_trafficDistribution_DataType
type SCINetworkWirelessReport21trafficDistributionDataType struct {
	Children []*SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType `json:"children,omitempty"`

	Key *string `json:"key,omitempty"`

	Name *string `json:"name,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport21trafficDistributionDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport21trafficDistributionDataType SCINetworkWirelessReport21trafficDistributionDataType
	tmpType := new(_SCINetworkWirelessReport21trafficDistributionDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "children")
	delete(tmpType.XAdditionalProperties, "key")
	delete(tmpType.XAdditionalProperties, "name")
	*t = SCINetworkWirelessReport21trafficDistributionDataType(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport21trafficDistributionDataType) MarshalJSON() ([]byte, error) {
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
	if t.Key != nil {
		tmp["key"] = t.Key
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWirelessReport21trafficDistributionDataType() *SCINetworkWirelessReport21trafficDistributionDataType {
	m := new(SCINetworkWirelessReport21trafficDistributionDataType)
	return m
}

// SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_21_trafficDistribution_DataTypeChildrenType
type SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType struct {
	Children []*SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType `json:"children,omitempty"`

	Display *string `json:"display,omitempty"`

	Key *string `json:"key,omitempty"`

	Name *string `json:"name,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType
	tmpType := new(_SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "children")
	delete(tmpType.XAdditionalProperties, "display")
	delete(tmpType.XAdditionalProperties, "key")
	delete(tmpType.XAdditionalProperties, "name")
	*t = SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType) MarshalJSON() ([]byte, error) {
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
	if t.Display != nil {
		tmp["display"] = t.Display
	}
	if t.Key != nil {
		tmp["key"] = t.Key
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWirelessReport21trafficDistributionDataTypeChildrenType() *SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType {
	m := new(SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType)
	return m
}

// SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_21_trafficDistribution_DataTypeChildrenTypeChildrenType
type SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType struct {
	Display *string `json:"display,omitempty"`

	Key *string `json:"key,omitempty"`

	Name *string `json:"name,omitempty"`

	Size *float64 `json:"size,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType
	tmpType := new(_SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "display")
	delete(tmpType.XAdditionalProperties, "key")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "size")
	*t = SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Display != nil {
		tmp["display"] = t.Display
	}
	if t.Key != nil {
		tmp["key"] = t.Key
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.Size != nil {
		tmp["size"] = t.Size
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType() *SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType {
	m := new(SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType)
	return m
}

// SCINetworkWirelessReport22trafficTrendData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_22_trafficTrend_Data
type SCINetworkWirelessReport22trafficTrendData [][]*SCINetworkWirelessReport22trafficTrendDataTypeType

func MakeSCINetworkWirelessReport22trafficTrendData() SCINetworkWirelessReport22trafficTrendData {
	m := make(SCINetworkWirelessReport22trafficTrendData, 0)
	return m
}

// SCINetworkWirelessReport22trafficTrendDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_22_trafficTrend_DataType
type SCINetworkWirelessReport22trafficTrendDataType []*SCINetworkWirelessReport22trafficTrendDataTypeType

func MakeSCINetworkWirelessReport22trafficTrendDataType() SCINetworkWirelessReport22trafficTrendDataType {
	m := make(SCINetworkWirelessReport22trafficTrendDataType, 0)
	return m
}

// SCINetworkWirelessReport22trafficTrendDataTypeType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_22_trafficTrend_DataTypeType
type SCINetworkWirelessReport22trafficTrendDataTypeType struct {
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

func (t *SCINetworkWirelessReport22trafficTrendDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport22trafficTrendDataTypeType SCINetworkWirelessReport22trafficTrendDataTypeType
	tmpType := new(_SCINetworkWirelessReport22trafficTrendDataTypeType)
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
	*t = SCINetworkWirelessReport22trafficTrendDataTypeType(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport22trafficTrendDataTypeType) MarshalJSON() ([]byte, error) {
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

func NewSCINetworkWirelessReport22trafficTrendDataTypeType() *SCINetworkWirelessReport22trafficTrendDataTypeType {
	m := new(SCINetworkWirelessReport22trafficTrendDataTypeType)
	return m
}

// SCINetworkWirelessReport23trafficOverTimeTableData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_23_trafficOverTimeTable_Data
type SCINetworkWirelessReport23trafficOverTimeTableData []*SCINetworkWirelessReport23trafficOverTimeTableDataType

func MakeSCINetworkWirelessReport23trafficOverTimeTableData() SCINetworkWirelessReport23trafficOverTimeTableData {
	m := make(SCINetworkWirelessReport23trafficOverTimeTableData, 0)
	return m
}

// SCINetworkWirelessReport23trafficOverTimeTableDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_23_trafficOverTimeTable_DataType
type SCINetworkWirelessReport23trafficOverTimeTableDataType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	End *string `json:"end,omitempty"`

	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	SessionCount interface{} `json:"sessionCount,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTraffic24 *float64 `json:"totalTraffic_2-4,omitempty"`

	TotalTraffic5 *float64 `json:"totalTraffic_5,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UniqueUsers interface{} `json:"uniqueUsers,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport23trafficOverTimeTableDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport23trafficOverTimeTableDataType SCINetworkWirelessReport23trafficOverTimeTableDataType
	tmpType := new(_SCINetworkWirelessReport23trafficOverTimeTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avgRateTotalRxTraffic")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTraffic")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTxTraffic")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "mgmtRxBytes")
	delete(tmpType.XAdditionalProperties, "mgmtTraffic")
	delete(tmpType.XAdditionalProperties, "mgmtTxBytes")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic_2-4")
	delete(tmpType.XAdditionalProperties, "totalTraffic_5")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "uniqueUsers")
	delete(tmpType.XAdditionalProperties, "userRxBytes")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	delete(tmpType.XAdditionalProperties, "userTxBytes")
	*t = SCINetworkWirelessReport23trafficOverTimeTableDataType(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport23trafficOverTimeTableDataType) MarshalJSON() ([]byte, error) {
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
	if t.AvgRateTotalTxTraffic != nil {
		tmp["avgRateTotalTxTraffic"] = t.AvgRateTotalTxTraffic
	}
	if t.End != nil {
		tmp["end"] = t.End
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
	if t.UniqueUsers != nil {
		tmp["uniqueUsers"] = t.UniqueUsers
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

func NewSCINetworkWirelessReport23trafficOverTimeTableDataType() *SCINetworkWirelessReport23trafficOverTimeTableDataType {
	m := new(SCINetworkWirelessReport23trafficOverTimeTableDataType)
	return m
}

// SCINetworkWirelessReport23trafficOverTimeTableMetaData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_23_trafficOverTimeTable_MetaData
type SCINetworkWirelessReport23trafficOverTimeTableMetaData struct {
	MaxValues *SCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport23trafficOverTimeTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport23trafficOverTimeTableMetaData SCINetworkWirelessReport23trafficOverTimeTableMetaData
	tmpType := new(_SCINetworkWirelessReport23trafficOverTimeTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	*t = SCINetworkWirelessReport23trafficOverTimeTableMetaData(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport23trafficOverTimeTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCINetworkWirelessReport23trafficOverTimeTableMetaData() *SCINetworkWirelessReport23trafficOverTimeTableMetaData {
	m := new(SCINetworkWirelessReport23trafficOverTimeTableMetaData)
	return m
}

// SCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_23_trafficOverTimeTable_MetaDataMaxValuesType
type SCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTraffic24 *float64 `json:"totalTraffic_2-4,omitempty"`

	TotalTraffic5 *float64 `json:"totalTraffic_5,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UniqueUsers interface{} `json:"uniqueUsers,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`
}

func NewSCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType() *SCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType {
	m := new(SCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType)
	return m
}

// SCINetworkWirelessReport24topAPsByTrafficTableData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_24_topAPsByTrafficTable_Data
type SCINetworkWirelessReport24topAPsByTrafficTableData []*SCINetworkWirelessReport24topAPsByTrafficTableDataType

func MakeSCINetworkWirelessReport24topAPsByTrafficTableData() SCINetworkWirelessReport24topAPsByTrafficTableData {
	m := make(SCINetworkWirelessReport24topAPsByTrafficTableData, 0)
	return m
}

// SCINetworkWirelessReport24topAPsByTrafficTableDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_24_topAPsByTrafficTable_DataType
type SCINetworkWirelessReport24topAPsByTrafficTableDataType struct {
	ApIp *string `json:"apIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	CtrlMac *string `json:"ctrlMac,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	CtrlSerial *string `json:"ctrlSerial,omitempty"`

	Index *float64 `json:"index,omitempty"`

	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport24topAPsByTrafficTableDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport24topAPsByTrafficTableDataType SCINetworkWirelessReport24topAPsByTrafficTableDataType
	tmpType := new(_SCINetworkWirelessReport24topAPsByTrafficTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apIp")
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "avgRateTotalRxTraffic")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTraffic")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTxTraffic")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "ctrlMac")
	delete(tmpType.XAdditionalProperties, "ctrlName")
	delete(tmpType.XAdditionalProperties, "ctrlSerial")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "mgmtRxBytes")
	delete(tmpType.XAdditionalProperties, "mgmtTraffic")
	delete(tmpType.XAdditionalProperties, "mgmtTxBytes")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "userRxBytes")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	delete(tmpType.XAdditionalProperties, "userTxBytes")
	*t = SCINetworkWirelessReport24topAPsByTrafficTableDataType(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport24topAPsByTrafficTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ApIp != nil {
		tmp["apIp"] = t.ApIp
	}
	if t.ApMac != nil {
		tmp["apMac"] = t.ApMac
	}
	if t.ApName != nil {
		tmp["apName"] = t.ApName
	}
	if t.AvgRateTotalRxTraffic != nil {
		tmp["avgRateTotalRxTraffic"] = t.AvgRateTotalRxTraffic
	}
	if t.AvgRateTotalTraffic != nil {
		tmp["avgRateTotalTraffic"] = t.AvgRateTotalTraffic
	}
	if t.AvgRateTotalTxTraffic != nil {
		tmp["avgRateTotalTxTraffic"] = t.AvgRateTotalTxTraffic
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.CtrlMac != nil {
		tmp["ctrlMac"] = t.CtrlMac
	}
	if t.CtrlName != nil {
		tmp["ctrlName"] = t.CtrlName
	}
	if t.CtrlSerial != nil {
		tmp["ctrlSerial"] = t.CtrlSerial
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

func NewSCINetworkWirelessReport24topAPsByTrafficTableDataType() *SCINetworkWirelessReport24topAPsByTrafficTableDataType {
	m := new(SCINetworkWirelessReport24topAPsByTrafficTableDataType)
	return m
}

// SCINetworkWirelessReport24topAPsByTrafficTableMetaData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_24_topAPsByTrafficTable_MetaData
type SCINetworkWirelessReport24topAPsByTrafficTableMetaData struct {
	MaxValues *SCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport24topAPsByTrafficTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport24topAPsByTrafficTableMetaData SCINetworkWirelessReport24topAPsByTrafficTableMetaData
	tmpType := new(_SCINetworkWirelessReport24topAPsByTrafficTableMetaData)
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
	*t = SCINetworkWirelessReport24topAPsByTrafficTableMetaData(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport24topAPsByTrafficTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCINetworkWirelessReport24topAPsByTrafficTableMetaData() *SCINetworkWirelessReport24topAPsByTrafficTableMetaData {
	m := new(SCINetworkWirelessReport24topAPsByTrafficTableMetaData)
	return m
}

// SCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_24_topAPsByTrafficTable_MetaDataMaxValuesType
type SCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

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

func NewSCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType() *SCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType {
	m := new(SCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCINetworkWirelessReport25topAPsByClientsTableData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_25_topAPsByClientsTable_Data
type SCINetworkWirelessReport25topAPsByClientsTableData []*SCINetworkWirelessReport25topAPsByClientsTableDataType

func MakeSCINetworkWirelessReport25topAPsByClientsTableData() SCINetworkWirelessReport25topAPsByClientsTableData {
	m := make(SCINetworkWirelessReport25topAPsByClientsTableData, 0)
	return m
}

// SCINetworkWirelessReport25topAPsByClientsTableDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_25_topAPsByClientsTable_DataType
type SCINetworkWirelessReport25topAPsByClientsTableDataType struct {
	ApIp *string `json:"apIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	CtrlMac *string `json:"ctrlMac,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	CtrlSerial *string `json:"ctrlSerial,omitempty"`

	Index *float64 `json:"index,omitempty"`

	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport25topAPsByClientsTableDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport25topAPsByClientsTableDataType SCINetworkWirelessReport25topAPsByClientsTableDataType
	tmpType := new(_SCINetworkWirelessReport25topAPsByClientsTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apIp")
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "avgRateTotalRxTraffic")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTraffic")
	delete(tmpType.XAdditionalProperties, "avgRateTotalTxTraffic")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "ctrlMac")
	delete(tmpType.XAdditionalProperties, "ctrlName")
	delete(tmpType.XAdditionalProperties, "ctrlSerial")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "mgmtRxBytes")
	delete(tmpType.XAdditionalProperties, "mgmtTraffic")
	delete(tmpType.XAdditionalProperties, "mgmtTxBytes")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "userRxBytes")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	delete(tmpType.XAdditionalProperties, "userTxBytes")
	*t = SCINetworkWirelessReport25topAPsByClientsTableDataType(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport25topAPsByClientsTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ApIp != nil {
		tmp["apIp"] = t.ApIp
	}
	if t.ApMac != nil {
		tmp["apMac"] = t.ApMac
	}
	if t.ApName != nil {
		tmp["apName"] = t.ApName
	}
	if t.AvgRateTotalRxTraffic != nil {
		tmp["avgRateTotalRxTraffic"] = t.AvgRateTotalRxTraffic
	}
	if t.AvgRateTotalTraffic != nil {
		tmp["avgRateTotalTraffic"] = t.AvgRateTotalTraffic
	}
	if t.AvgRateTotalTxTraffic != nil {
		tmp["avgRateTotalTxTraffic"] = t.AvgRateTotalTxTraffic
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.CtrlMac != nil {
		tmp["ctrlMac"] = t.CtrlMac
	}
	if t.CtrlName != nil {
		tmp["ctrlName"] = t.CtrlName
	}
	if t.CtrlSerial != nil {
		tmp["ctrlSerial"] = t.CtrlSerial
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

func NewSCINetworkWirelessReport25topAPsByClientsTableDataType() *SCINetworkWirelessReport25topAPsByClientsTableDataType {
	m := new(SCINetworkWirelessReport25topAPsByClientsTableDataType)
	return m
}

// SCINetworkWirelessReport25topAPsByClientsTableMetaData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_25_topAPsByClientsTable_MetaData
type SCINetworkWirelessReport25topAPsByClientsTableMetaData struct {
	MaxValues *SCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport25topAPsByClientsTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport25topAPsByClientsTableMetaData SCINetworkWirelessReport25topAPsByClientsTableMetaData
	tmpType := new(_SCINetworkWirelessReport25topAPsByClientsTableMetaData)
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
	*t = SCINetworkWirelessReport25topAPsByClientsTableMetaData(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport25topAPsByClientsTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCINetworkWirelessReport25topAPsByClientsTableMetaData() *SCINetworkWirelessReport25topAPsByClientsTableMetaData {
	m := new(SCINetworkWirelessReport25topAPsByClientsTableMetaData)
	return m
}

// SCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_25_topAPsByClientsTable_MetaDataMaxValuesType
type SCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

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

func NewSCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType() *SCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType {
	m := new(SCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType)
	return m
}

// SCINetworkWirelessReport27top10ApByClientCountData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_27_top10ApByClientCount_Data
type SCINetworkWirelessReport27top10ApByClientCountData [][]*SCINetworkWirelessReport27top10ApByClientCountDataTypeType

func MakeSCINetworkWirelessReport27top10ApByClientCountData() SCINetworkWirelessReport27top10ApByClientCountData {
	m := make(SCINetworkWirelessReport27top10ApByClientCountData, 0)
	return m
}

// SCINetworkWirelessReport27top10ApByClientCountDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_27_top10ApByClientCount_DataType
type SCINetworkWirelessReport27top10ApByClientCountDataType []*SCINetworkWirelessReport27top10ApByClientCountDataTypeType

func MakeSCINetworkWirelessReport27top10ApByClientCountDataType() SCINetworkWirelessReport27top10ApByClientCountDataType {
	m := make(SCINetworkWirelessReport27top10ApByClientCountDataType, 0)
	return m
}

// SCINetworkWirelessReport27top10ApByClientCountDataTypeType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_27_top10ApByClientCount_DataTypeType
type SCINetworkWirelessReport27top10ApByClientCountDataTypeType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	UniqueUsers *float64 `json:"uniqueUsers,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport27top10ApByClientCountDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport27top10ApByClientCountDataTypeType SCINetworkWirelessReport27top10ApByClientCountDataTypeType
	tmpType := new(_SCINetworkWirelessReport27top10ApByClientCountDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "uniqueUsers")
	*t = SCINetworkWirelessReport27top10ApByClientCountDataTypeType(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport27top10ApByClientCountDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.UniqueUsers != nil {
		tmp["uniqueUsers"] = t.UniqueUsers
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWirelessReport27top10ApByClientCountDataTypeType() *SCINetworkWirelessReport27top10ApByClientCountDataTypeType {
	m := new(SCINetworkWirelessReport27top10ApByClientCountDataTypeType)
	return m
}

// SCINetworkWirelessReport27top10ApByClientCountMetaData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_27_top10ApByClientCount_MetaData
type SCINetworkWirelessReport27top10ApByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWirelessReport27top10ApByClientCountMetaData) UnmarshalJSON(b []byte) error {
	type _SCINetworkWirelessReport27top10ApByClientCountMetaData SCINetworkWirelessReport27top10ApByClientCountMetaData
	tmpType := new(_SCINetworkWirelessReport27top10ApByClientCountMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	delete(tmpType.XAdditionalProperties, "totalClients")
	*t = SCINetworkWirelessReport27top10ApByClientCountMetaData(*tmpType)
	return nil
}

func (t *SCINetworkWirelessReport27top10ApByClientCountMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCINetworkWirelessReport27top10ApByClientCountMetaData() *SCINetworkWirelessReport27top10ApByClientCountMetaData {
	m := new(SCINetworkWirelessReport27top10ApByClientCountMetaData)
	return m
}

// ReportNetworkWirelessReport20Overview
//
// Network - Wireless Report - Overview
//
// Operation ID: report_NetworkWirelessReport_20_overview
// Operation path: /reports/2/sections/20/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport20Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport20overview200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWirelessReport20overview200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWirelessReport20overview200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportNetworkWirelessReport20Overview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWirelessReport20overview200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWirelessReport20overview200ResponseTypeAPIResponse), err
}

// ReportNetworkWirelessReport21TrafficDistribution
//
// Network - Wireless Report - Traffic Distribution
//
// Operation ID: report_NetworkWirelessReport_21_trafficDistribution
// Operation path: /reports/2/sections/21/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport21TrafficDistribution(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport21trafficDistribution200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWirelessReport21trafficDistribution200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWirelessReport21trafficDistribution200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportNetworkWirelessReport21TrafficDistribution, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWirelessReport21trafficDistribution200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWirelessReport21trafficDistribution200ResponseTypeAPIResponse), err
}

// ReportNetworkWirelessReport22TrafficTrend
//
// Network - Wireless Report - Traffic Trend
//
// Operation ID: report_NetworkWirelessReport_22_trafficTrend
// Operation path: /reports/2/sections/22/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport22TrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport22trafficTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWirelessReport22trafficTrend200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWirelessReport22trafficTrend200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportNetworkWirelessReport22TrafficTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWirelessReport22trafficTrend200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWirelessReport22trafficTrend200ResponseTypeAPIResponse), err
}

// ReportNetworkWirelessReport23TrafficOverTimeTable
//
// Network - Wireless Report - Traffic Over Time
//
// Operation ID: report_NetworkWirelessReport_23_trafficOverTimeTable
// Operation path: /reports/2/sections/23/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport23TrafficOverTimeTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportNetworkWirelessReport23TrafficOverTimeTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseTypeAPIResponse), err
}

// ReportNetworkWirelessReport24TopAPsByTrafficTable
//
// Network - Wireless Report - Top APs by Traffic
//
// Operation ID: report_NetworkWirelessReport_24_topAPsByTrafficTable
// Operation path: /reports/2/sections/24/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport24TopAPsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportNetworkWirelessReport24TopAPsByTrafficTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseTypeAPIResponse), err
}

// ReportNetworkWirelessReport25TopAPsByClientsTable
//
// Network - Wireless Report - Top APs by Client Count
//
// Operation ID: report_NetworkWirelessReport_25_topAPsByClientsTable
// Operation path: /reports/2/sections/25/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport25TopAPsByClientsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportNetworkWirelessReport25TopAPsByClientsTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseTypeAPIResponse), err
}

// ReportNetworkWirelessReport26Top10APsByTrafficVolume
//
// Network - Wireless Report - Top APs by Traffic
//
// Operation ID: report_NetworkWirelessReport_26_top10APsByTrafficVolume
// Operation path: /reports/2/sections/26/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport26Top10APsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportNetworkWirelessReport26Top10APsByTrafficVolume, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseTypeAPIResponse), err
}

// ReportNetworkWirelessReport27Top10ApByClientCount
//
// Network - Wireless Report - Top APs by Client Count
//
// Operation ID: report_NetworkWirelessReport_27_top10ApByClientCount
// Operation path: /reports/2/sections/27/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport27Top10ApByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWirelessReport27top10ApByClientCount200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportNetworkWirelessReport27Top10ApByClientCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseTypeAPIResponse), err
}
