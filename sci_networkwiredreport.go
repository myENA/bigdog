package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCINetworkWiredReportService struct {
	apiClient *SCIClient
}

func NewSCINetworkWiredReportService(c *SCIClient) *SCINetworkWiredReportService {
	s := new(SCINetworkWiredReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCINetworkWiredReportService() *SCINetworkWiredReportService {
	return NewSCINetworkWiredReportService(ss.apiClient)
}

// SCINetworkWiredReport123topSwitchPOEUtilChartData
//
// Definition: NetworkWiredReport_NetworkWiredReport_123_topSwitchPOEUtilChart_Data
type SCINetworkWiredReport123topSwitchPOEUtilChartData [][]*SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType

func MakeSCINetworkWiredReport123topSwitchPOEUtilChartData() SCINetworkWiredReport123topSwitchPOEUtilChartData {
	m := make(SCINetworkWiredReport123topSwitchPOEUtilChartData, 0)
	return m
}

// SCINetworkWiredReport123topSwitchPOEUtilChartDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_123_topSwitchPOEUtilChart_DataType
type SCINetworkWiredReport123topSwitchPOEUtilChartDataType []*SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType

func MakeSCINetworkWiredReport123topSwitchPOEUtilChartDataType() SCINetworkWiredReport123topSwitchPOEUtilChartDataType {
	m := make(SCINetworkWiredReport123topSwitchPOEUtilChartDataType, 0)
	return m
}

// SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType
//
// Definition: NetworkWiredReport_NetworkWiredReport_123_topSwitchPOEUtilChart_DataTypeType
type SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType struct {
	End *string `json:"end,omitempty"`

	PoeUtilization *float64 `json:"poeUtilization,omitempty"`

	Start *string `json:"start,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType
	tmpType := new(_SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "poeUtilization")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "switchId")
	delete(tmpType.XAdditionalProperties, "switchName")
	*t = SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.PoeUtilization != nil {
		tmp["poeUtilization"] = t.PoeUtilization
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.SwitchId != nil {
		tmp["switchId"] = t.SwitchId
	}
	if t.SwitchName != nil {
		tmp["switchName"] = t.SwitchName
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType() *SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType {
	m := new(SCINetworkWiredReport123topSwitchPOEUtilChartDataTypeType)
	return m
}

// SCINetworkWiredReport123topSwitchPOEUtilChartMetaData
//
// Definition: NetworkWiredReport_NetworkWiredReport_123_topSwitchPOEUtilChart_MetaData
type SCINetworkWiredReport123topSwitchPOEUtilChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport123topSwitchPOEUtilChartMetaData) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport123topSwitchPOEUtilChartMetaData SCINetworkWiredReport123topSwitchPOEUtilChartMetaData
	tmpType := new(_SCINetworkWiredReport123topSwitchPOEUtilChartMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	*t = SCINetworkWiredReport123topSwitchPOEUtilChartMetaData(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport123topSwitchPOEUtilChartMetaData) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport123topSwitchPOEUtilChartMetaData() *SCINetworkWiredReport123topSwitchPOEUtilChartMetaData {
	m := new(SCINetworkWiredReport123topSwitchPOEUtilChartMetaData)
	return m
}

// SCINetworkWiredReport124topSwitchPOEUtilsData
//
// Definition: NetworkWiredReport_NetworkWiredReport_124_topSwitchPOEUtils_Data
type SCINetworkWiredReport124topSwitchPOEUtilsData []*SCINetworkWiredReport124topSwitchPOEUtilsDataType

func MakeSCINetworkWiredReport124topSwitchPOEUtilsData() SCINetworkWiredReport124topSwitchPOEUtilsData {
	m := make(SCINetworkWiredReport124topSwitchPOEUtilsData, 0)
	return m
}

// SCINetworkWiredReport124topSwitchPOEUtilsDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_124_topSwitchPOEUtils_DataType
type SCINetworkWiredReport124topSwitchPOEUtilsDataType struct {
	Index *float64 `json:"index,omitempty"`

	PoeTotal *float64 `json:"poeTotal,omitempty"`

	PoeUtilization *float64 `json:"poeUtilization,omitempty"`

	PoeUtilPercent *float64 `json:"poeUtilPercent,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport124topSwitchPOEUtilsDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport124topSwitchPOEUtilsDataType SCINetworkWiredReport124topSwitchPOEUtilsDataType
	tmpType := new(_SCINetworkWiredReport124topSwitchPOEUtilsDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "poeTotal")
	delete(tmpType.XAdditionalProperties, "poeUtilization")
	delete(tmpType.XAdditionalProperties, "poeUtilPercent")
	delete(tmpType.XAdditionalProperties, "switchId")
	delete(tmpType.XAdditionalProperties, "switchName")
	*t = SCINetworkWiredReport124topSwitchPOEUtilsDataType(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport124topSwitchPOEUtilsDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.PoeTotal != nil {
		tmp["poeTotal"] = t.PoeTotal
	}
	if t.PoeUtilization != nil {
		tmp["poeUtilization"] = t.PoeUtilization
	}
	if t.PoeUtilPercent != nil {
		tmp["poeUtilPercent"] = t.PoeUtilPercent
	}
	if t.SwitchId != nil {
		tmp["switchId"] = t.SwitchId
	}
	if t.SwitchName != nil {
		tmp["switchName"] = t.SwitchName
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport124topSwitchPOEUtilsDataType() *SCINetworkWiredReport124topSwitchPOEUtilsDataType {
	m := new(SCINetworkWiredReport124topSwitchPOEUtilsDataType)
	return m
}

// SCINetworkWiredReport128topSwitchesByTrafficTableData
//
// Definition: NetworkWiredReport_NetworkWiredReport_128_topSwitchesByTrafficTable_Data
type SCINetworkWiredReport128topSwitchesByTrafficTableData []*SCINetworkWiredReport128topSwitchesByTrafficTableDataType

func MakeSCINetworkWiredReport128topSwitchesByTrafficTableData() SCINetworkWiredReport128topSwitchesByTrafficTableData {
	m := make(SCINetworkWiredReport128topSwitchesByTrafficTableData, 0)
	return m
}

// SCINetworkWiredReport128topSwitchesByTrafficTableDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_128_topSwitchesByTrafficTable_DataType
type SCINetworkWiredReport128topSwitchesByTrafficTableDataType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	Index *float64 `json:"index,omitempty"`

	PortCount *float64 `json:"portCount,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`

	SwitchUnitCount *float64 `json:"switchUnitCount,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport128topSwitchesByTrafficTableDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport128topSwitchesByTrafficTableDataType SCINetworkWiredReport128topSwitchesByTrafficTableDataType
	tmpType := new(_SCINetworkWiredReport128topSwitchesByTrafficTableDataType)
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
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "portCount")
	delete(tmpType.XAdditionalProperties, "switchId")
	delete(tmpType.XAdditionalProperties, "switchName")
	delete(tmpType.XAdditionalProperties, "switchUnitCount")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	*t = SCINetworkWiredReport128topSwitchesByTrafficTableDataType(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport128topSwitchesByTrafficTableDataType) MarshalJSON() ([]byte, error) {
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
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.PortCount != nil {
		tmp["portCount"] = t.PortCount
	}
	if t.SwitchId != nil {
		tmp["switchId"] = t.SwitchId
	}
	if t.SwitchName != nil {
		tmp["switchName"] = t.SwitchName
	}
	if t.SwitchUnitCount != nil {
		tmp["switchUnitCount"] = t.SwitchUnitCount
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
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport128topSwitchesByTrafficTableDataType() *SCINetworkWiredReport128topSwitchesByTrafficTableDataType {
	m := new(SCINetworkWiredReport128topSwitchesByTrafficTableDataType)
	return m
}

// SCINetworkWiredReport128topSwitchesByTrafficTableMetaData
//
// Definition: NetworkWiredReport_NetworkWiredReport_128_topSwitchesByTrafficTable_MetaData
type SCINetworkWiredReport128topSwitchesByTrafficTableMetaData struct {
	MaxValues *SCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport128topSwitchesByTrafficTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport128topSwitchesByTrafficTableMetaData SCINetworkWiredReport128topSwitchesByTrafficTableMetaData
	tmpType := new(_SCINetworkWiredReport128topSwitchesByTrafficTableMetaData)
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
	*t = SCINetworkWiredReport128topSwitchesByTrafficTableMetaData(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport128topSwitchesByTrafficTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCINetworkWiredReport128topSwitchesByTrafficTableMetaData() *SCINetworkWiredReport128topSwitchesByTrafficTableMetaData {
	m := new(SCINetworkWiredReport128topSwitchesByTrafficTableMetaData)
	return m
}

// SCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType
//
// Definition: NetworkWiredReport_NetworkWiredReport_128_topSwitchesByTrafficTable_MetaDataMaxValuesType
type SCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`
}

func NewSCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType() *SCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType {
	m := new(SCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCINetworkWiredReport134wiredOverviewData
//
// Definition: NetworkWiredReport_NetworkWiredReport_134_wiredOverview_Data
type SCINetworkWiredReport134wiredOverviewData []*SCINetworkWiredReport134wiredOverviewDataType

func MakeSCINetworkWiredReport134wiredOverviewData() SCINetworkWiredReport134wiredOverviewData {
	m := make(SCINetworkWiredReport134wiredOverviewData, 0)
	return m
}

// SCINetworkWiredReport134wiredOverviewDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_134_wiredOverview_DataType
type SCINetworkWiredReport134wiredOverviewDataType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalSwitches *float64 `json:"totalSwitches,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport134wiredOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport134wiredOverviewDataType SCINetworkWiredReport134wiredOverviewDataType
	tmpType := new(_SCINetworkWiredReport134wiredOverviewDataType)
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
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalSwitches")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	*t = SCINetworkWiredReport134wiredOverviewDataType(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport134wiredOverviewDataType) MarshalJSON() ([]byte, error) {
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
	if t.TotalRxTraffic != nil {
		tmp["totalRxTraffic"] = t.TotalRxTraffic
	}
	if t.TotalSwitches != nil {
		tmp["totalSwitches"] = t.TotalSwitches
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.TotalTxTraffic != nil {
		tmp["totalTxTraffic"] = t.TotalTxTraffic
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport134wiredOverviewDataType() *SCINetworkWiredReport134wiredOverviewDataType {
	m := new(SCINetworkWiredReport134wiredOverviewDataType)
	return m
}

// SCINetworkWiredReport135wiredTrafficDistributionData
//
// Definition: NetworkWiredReport_NetworkWiredReport_135_wiredTrafficDistribution_Data
type SCINetworkWiredReport135wiredTrafficDistributionData []*SCINetworkWiredReport135wiredTrafficDistributionDataType

func MakeSCINetworkWiredReport135wiredTrafficDistributionData() SCINetworkWiredReport135wiredTrafficDistributionData {
	m := make(SCINetworkWiredReport135wiredTrafficDistributionData, 0)
	return m
}

// SCINetworkWiredReport135wiredTrafficDistributionDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_135_wiredTrafficDistribution_DataType
type SCINetworkWiredReport135wiredTrafficDistributionDataType struct {
	Categories []string `json:"categories,omitempty"`

	Series []*SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType `json:"series,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport135wiredTrafficDistributionDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport135wiredTrafficDistributionDataType SCINetworkWiredReport135wiredTrafficDistributionDataType
	tmpType := new(_SCINetworkWiredReport135wiredTrafficDistributionDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "categories")
	delete(tmpType.XAdditionalProperties, "series")
	*t = SCINetworkWiredReport135wiredTrafficDistributionDataType(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport135wiredTrafficDistributionDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Categories != nil {
		tmp["categories"] = t.Categories
	}
	if t.Series != nil {
		tmp["series"] = t.Series
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport135wiredTrafficDistributionDataType() *SCINetworkWiredReport135wiredTrafficDistributionDataType {
	m := new(SCINetworkWiredReport135wiredTrafficDistributionDataType)
	return m
}

// SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType
//
// Definition: NetworkWiredReport_NetworkWiredReport_135_wiredTrafficDistribution_DataTypeSeriesType
type SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType struct {
	Data []float64 `json:"data,omitempty"`

	Name *string `json:"name,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType
	tmpType := new(_SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "data")
	delete(tmpType.XAdditionalProperties, "name")
	*t = SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Data != nil {
		tmp["data"] = t.Data
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType() *SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType {
	m := new(SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType)
	return m
}

// SCINetworkWiredReport135wiredTrafficDistributionMetaData
//
// Definition: NetworkWiredReport_NetworkWiredReport_135_wiredTrafficDistribution_MetaData
type SCINetworkWiredReport135wiredTrafficDistributionMetaData struct {
	ChartType *string `json:"chartType,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport135wiredTrafficDistributionMetaData) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport135wiredTrafficDistributionMetaData SCINetworkWiredReport135wiredTrafficDistributionMetaData
	tmpType := new(_SCINetworkWiredReport135wiredTrafficDistributionMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "chartType")
	*t = SCINetworkWiredReport135wiredTrafficDistributionMetaData(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport135wiredTrafficDistributionMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ChartType != nil {
		tmp["chartType"] = t.ChartType
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport135wiredTrafficDistributionMetaData() *SCINetworkWiredReport135wiredTrafficDistributionMetaData {
	m := new(SCINetworkWiredReport135wiredTrafficDistributionMetaData)
	return m
}

// SCINetworkWiredReport136switchTrafficTrendData
//
// Definition: NetworkWiredReport_NetworkWiredReport_136_switchTrafficTrend_Data
type SCINetworkWiredReport136switchTrafficTrendData [][]*SCINetworkWiredReport136switchTrafficTrendDataTypeType

func MakeSCINetworkWiredReport136switchTrafficTrendData() SCINetworkWiredReport136switchTrafficTrendData {
	m := make(SCINetworkWiredReport136switchTrafficTrendData, 0)
	return m
}

// SCINetworkWiredReport136switchTrafficTrendDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_136_switchTrafficTrend_DataType
type SCINetworkWiredReport136switchTrafficTrendDataType []*SCINetworkWiredReport136switchTrafficTrendDataTypeType

func MakeSCINetworkWiredReport136switchTrafficTrendDataType() SCINetworkWiredReport136switchTrafficTrendDataType {
	m := make(SCINetworkWiredReport136switchTrafficTrendDataType, 0)
	return m
}

// SCINetworkWiredReport136switchTrafficTrendDataTypeType
//
// Definition: NetworkWiredReport_NetworkWiredReport_136_switchTrafficTrend_DataTypeType
type SCINetworkWiredReport136switchTrafficTrendDataTypeType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport136switchTrafficTrendDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport136switchTrafficTrendDataTypeType SCINetworkWiredReport136switchTrafficTrendDataTypeType
	tmpType := new(_SCINetworkWiredReport136switchTrafficTrendDataTypeType)
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
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	*t = SCINetworkWiredReport136switchTrafficTrendDataTypeType(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport136switchTrafficTrendDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.Start != nil {
		tmp["start"] = t.Start
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
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport136switchTrafficTrendDataTypeType() *SCINetworkWiredReport136switchTrafficTrendDataTypeType {
	m := new(SCINetworkWiredReport136switchTrafficTrendDataTypeType)
	return m
}

// SCINetworkWiredReport141switchErrorTrendData
//
// Definition: NetworkWiredReport_NetworkWiredReport_141_switchErrorTrend_Data
type SCINetworkWiredReport141switchErrorTrendData []*SCINetworkWiredReport141switchErrorTrendDataType

func MakeSCINetworkWiredReport141switchErrorTrendData() SCINetworkWiredReport141switchErrorTrendData {
	m := make(SCINetworkWiredReport141switchErrorTrendData, 0)
	return m
}

// SCINetworkWiredReport141switchErrorTrendDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_141_switchErrorTrend_DataType
type SCINetworkWiredReport141switchErrorTrendDataType struct {
	CrcErrors *float64 `json:"crcErrors,omitempty"`

	End *string `json:"end,omitempty"`

	InDiscards *float64 `json:"inDiscards,omitempty"`

	InErrors *float64 `json:"inErrors,omitempty"`

	OutErrors *float64 `json:"outErrors,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport141switchErrorTrendDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport141switchErrorTrendDataType SCINetworkWiredReport141switchErrorTrendDataType
	tmpType := new(_SCINetworkWiredReport141switchErrorTrendDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "crcErrors")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "inDiscards")
	delete(tmpType.XAdditionalProperties, "inErrors")
	delete(tmpType.XAdditionalProperties, "outErrors")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCINetworkWiredReport141switchErrorTrendDataType(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport141switchErrorTrendDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.CrcErrors != nil {
		tmp["crcErrors"] = t.CrcErrors
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.InDiscards != nil {
		tmp["inDiscards"] = t.InDiscards
	}
	if t.InErrors != nil {
		tmp["inErrors"] = t.InErrors
	}
	if t.OutErrors != nil {
		tmp["outErrors"] = t.OutErrors
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport141switchErrorTrendDataType() *SCINetworkWiredReport141switchErrorTrendDataType {
	m := new(SCINetworkWiredReport141switchErrorTrendDataType)
	return m
}

// SCINetworkWiredReport142topSwitchesByErrorsChartData
//
// Definition: NetworkWiredReport_NetworkWiredReport_142_topSwitchesByErrorsChart_Data
type SCINetworkWiredReport142topSwitchesByErrorsChartData [][]*SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType

func MakeSCINetworkWiredReport142topSwitchesByErrorsChartData() SCINetworkWiredReport142topSwitchesByErrorsChartData {
	m := make(SCINetworkWiredReport142topSwitchesByErrorsChartData, 0)
	return m
}

// SCINetworkWiredReport142topSwitchesByErrorsChartDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_142_topSwitchesByErrorsChart_DataType
type SCINetworkWiredReport142topSwitchesByErrorsChartDataType []*SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType

func MakeSCINetworkWiredReport142topSwitchesByErrorsChartDataType() SCINetworkWiredReport142topSwitchesByErrorsChartDataType {
	m := make(SCINetworkWiredReport142topSwitchesByErrorsChartDataType, 0)
	return m
}

// SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType
//
// Definition: NetworkWiredReport_NetworkWiredReport_142_topSwitchesByErrorsChart_DataTypeType
type SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType struct {
	End *string `json:"end,omitempty"`

	Error *float64 `json:"error,omitempty"`

	Start *string `json:"start,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType
	tmpType := new(_SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "error")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "switchId")
	delete(tmpType.XAdditionalProperties, "switchName")
	*t = SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.Error != nil {
		tmp["error"] = t.Error
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.SwitchId != nil {
		tmp["switchId"] = t.SwitchId
	}
	if t.SwitchName != nil {
		tmp["switchName"] = t.SwitchName
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType() *SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType {
	m := new(SCINetworkWiredReport142topSwitchesByErrorsChartDataTypeType)
	return m
}

// SCINetworkWiredReport142topSwitchesByErrorsChartMetaData
//
// Definition: NetworkWiredReport_NetworkWiredReport_142_topSwitchesByErrorsChart_MetaData
type SCINetworkWiredReport142topSwitchesByErrorsChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport142topSwitchesByErrorsChartMetaData) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport142topSwitchesByErrorsChartMetaData SCINetworkWiredReport142topSwitchesByErrorsChartMetaData
	tmpType := new(_SCINetworkWiredReport142topSwitchesByErrorsChartMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	*t = SCINetworkWiredReport142topSwitchesByErrorsChartMetaData(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport142topSwitchesByErrorsChartMetaData) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport142topSwitchesByErrorsChartMetaData() *SCINetworkWiredReport142topSwitchesByErrorsChartMetaData {
	m := new(SCINetworkWiredReport142topSwitchesByErrorsChartMetaData)
	return m
}

// SCINetworkWiredReport143topSwitchesByErrorsTableData
//
// Definition: NetworkWiredReport_NetworkWiredReport_143_topSwitchesByErrorsTable_Data
type SCINetworkWiredReport143topSwitchesByErrorsTableData []*SCINetworkWiredReport143topSwitchesByErrorsTableDataType

func MakeSCINetworkWiredReport143topSwitchesByErrorsTableData() SCINetworkWiredReport143topSwitchesByErrorsTableData {
	m := make(SCINetworkWiredReport143topSwitchesByErrorsTableData, 0)
	return m
}

// SCINetworkWiredReport143topSwitchesByErrorsTableDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_143_topSwitchesByErrorsTable_DataType
type SCINetworkWiredReport143topSwitchesByErrorsTableDataType struct {
	Error *float64 `json:"error,omitempty"`

	Index *float64 `json:"index,omitempty"`

	InErr *float64 `json:"inErr,omitempty"`

	OutErr *float64 `json:"outErr,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCINetworkWiredReport143topSwitchesByErrorsTableDataType) UnmarshalJSON(b []byte) error {
	type _SCINetworkWiredReport143topSwitchesByErrorsTableDataType SCINetworkWiredReport143topSwitchesByErrorsTableDataType
	tmpType := new(_SCINetworkWiredReport143topSwitchesByErrorsTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "error")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "inErr")
	delete(tmpType.XAdditionalProperties, "outErr")
	delete(tmpType.XAdditionalProperties, "switchId")
	delete(tmpType.XAdditionalProperties, "switchName")
	*t = SCINetworkWiredReport143topSwitchesByErrorsTableDataType(*tmpType)
	return nil
}

func (t *SCINetworkWiredReport143topSwitchesByErrorsTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Error != nil {
		tmp["error"] = t.Error
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.InErr != nil {
		tmp["inErr"] = t.InErr
	}
	if t.OutErr != nil {
		tmp["outErr"] = t.OutErr
	}
	if t.SwitchId != nil {
		tmp["switchId"] = t.SwitchId
	}
	if t.SwitchName != nil {
		tmp["switchName"] = t.SwitchName
	}
	return json.Marshal(tmp)
}

func NewSCINetworkWiredReport143topSwitchesByErrorsTableDataType() *SCINetworkWiredReport143topSwitchesByErrorsTableDataType {
	m := new(SCINetworkWiredReport143topSwitchesByErrorsTableDataType)
	return m
}

// ReportNetworkWiredReport123TopSwitchPOEUtilChart
//
// Network - Wired Report - Top Switches by PoE Usage
//
// Operation ID: report_NetworkWiredReport_123_topSwitchPOEUtilChart
// Operation path: /reports/17/sections/123/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport123TopSwitchPOEUtilChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportNetworkWiredReport123TopSwitchPOEUtilChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseTypeAPIResponse), err
}

// ReportNetworkWiredReport124TopSwitchPOEUtils
//
// Network - Wired Report - Top Switches by PoE Usage
//
// Operation ID: report_NetworkWiredReport_124_topSwitchPOEUtils
// Operation path: /reports/17/sections/124/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport124TopSwitchPOEUtils(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportNetworkWiredReport124TopSwitchPOEUtils, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseTypeAPIResponse), err
}

// ReportNetworkWiredReport127Top10SwitchesByTrafficVolume
//
// Network - Wired Report - Top Switches by Traffic
//
// Operation ID: report_NetworkWiredReport_127_top10SwitchesByTrafficVolume
// Operation path: /reports/17/sections/127/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport127Top10SwitchesByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportNetworkWiredReport127Top10SwitchesByTrafficVolume, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseTypeAPIResponse), err
}

// ReportNetworkWiredReport128TopSwitchesByTrafficTable
//
// Network - Wired Report - Top Switches by Traffic
//
// Operation ID: report_NetworkWiredReport_128_topSwitchesByTrafficTable
// Operation path: /reports/17/sections/128/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport128TopSwitchesByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportNetworkWiredReport128TopSwitchesByTrafficTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseTypeAPIResponse), err
}

// ReportNetworkWiredReport134WiredOverview
//
// Network - Wired Report - Overview
//
// Operation ID: report_NetworkWiredReport_134_wiredOverview
// Operation path: /reports/17/sections/134/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport134WiredOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport134wiredOverview200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWiredReport134wiredOverview200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWiredReport134wiredOverview200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportNetworkWiredReport134WiredOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWiredReport134wiredOverview200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWiredReport134wiredOverview200ResponseTypeAPIResponse), err
}

// ReportNetworkWiredReport135WiredTrafficDistribution
//
// Network - Wired Report - Traffic Distribution by Switch Model and Port Speed
//
// Operation ID: report_NetworkWiredReport_135_wiredTrafficDistribution
// Operation path: /reports/17/sections/135/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport135WiredTrafficDistribution(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportNetworkWiredReport135WiredTrafficDistribution, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseTypeAPIResponse), err
}

// ReportNetworkWiredReport136SwitchTrafficTrend
//
// Network - Wired Report - Traffic Trend
//
// Operation ID: report_NetworkWiredReport_136_switchTrafficTrend
// Operation path: /reports/17/sections/136/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport136SwitchTrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport136switchTrafficTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWiredReport136switchTrafficTrend200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWiredReport136switchTrafficTrend200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportNetworkWiredReport136SwitchTrafficTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWiredReport136switchTrafficTrend200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWiredReport136switchTrafficTrend200ResponseTypeAPIResponse), err
}

// ReportNetworkWiredReport141SwitchErrorTrend
//
// Network - Wired Report - Error Trend
//
// Operation ID: report_NetworkWiredReport_141_switchErrorTrend
// Operation path: /reports/17/sections/141/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport141SwitchErrorTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport141switchErrorTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWiredReport141switchErrorTrend200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWiredReport141switchErrorTrend200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportNetworkWiredReport141SwitchErrorTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWiredReport141switchErrorTrend200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWiredReport141switchErrorTrend200ResponseTypeAPIResponse), err
}

// ReportNetworkWiredReport142TopSwitchesByErrorsChart
//
// Network - Wired Report - Top Switches by Errors
//
// Operation ID: report_NetworkWiredReport_142_topSwitchesByErrorsChart
// Operation path: /reports/17/sections/142/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport142TopSwitchesByErrorsChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportNetworkWiredReport142TopSwitchesByErrorsChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseTypeAPIResponse), err
}

// ReportNetworkWiredReport143TopSwitchesByErrorsTable
//
// Network - Wired Report - Top Switches by Errors
//
// Operation ID: report_NetworkWiredReport_143_topSwitchesByErrorsTable
// Operation path: /reports/17/sections/143/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport143TopSwitchesByErrorsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportNetworkWiredReport143TopSwitchesByErrorsTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseTypeAPIResponse), err
}
