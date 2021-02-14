package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type SCIInventoryAPsReportService struct {
	apiClient *SCIClient
}

func NewSCIInventoryAPsReportService(c *SCIClient) *SCIInventoryAPsReportService {
	s := new(SCIInventoryAPsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIInventoryAPsReportService() *SCIInventoryAPsReportService {
	return NewSCIInventoryAPsReportService(ss.apiClient)
}

// SCIInventoryAPsReport46apInventoryOverviewData
//
// Definition: InventoryAPsReport_InventoryAPsReport_46_apInventoryOverview_Data
type SCIInventoryAPsReport46apInventoryOverviewData []*SCIInventoryAPsReport46apInventoryOverviewDataType

func MakeSCIInventoryAPsReport46apInventoryOverviewData() SCIInventoryAPsReport46apInventoryOverviewData {
	m := make(SCIInventoryAPsReport46apInventoryOverviewData, 0)
	return m
}

// SCIInventoryAPsReport46apInventoryOverviewDataType
//
// Definition: InventoryAPsReport_InventoryAPsReport_46_apInventoryOverview_DataType
type SCIInventoryAPsReport46apInventoryOverviewDataType struct {
	Offline *float64 `json:"offline,omitempty"`

	Online *float64 `json:"online,omitempty"`

	Others *float64 `json:"others,omitempty"`

	Reboots *float64 `json:"reboots,omitempty"`

	Total *float64 `json:"total,omitempty"`

	TotalApsWAlarm *float64 `json:"totalApsWAlarm,omitempty"`

	TotalApsWReboot *float64 `json:"totalApsWReboot,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport46apInventoryOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport46apInventoryOverviewDataType SCIInventoryAPsReport46apInventoryOverviewDataType
	tmpType := new(_SCIInventoryAPsReport46apInventoryOverviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "offline")
	delete(tmpType.XAdditionalProperties, "online")
	delete(tmpType.XAdditionalProperties, "others")
	delete(tmpType.XAdditionalProperties, "reboots")
	delete(tmpType.XAdditionalProperties, "total")
	delete(tmpType.XAdditionalProperties, "totalApsWAlarm")
	delete(tmpType.XAdditionalProperties, "totalApsWReboot")
	*t = SCIInventoryAPsReport46apInventoryOverviewDataType(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport46apInventoryOverviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Offline != nil {
		tmp["offline"] = t.Offline
	}
	if t.Online != nil {
		tmp["online"] = t.Online
	}
	if t.Others != nil {
		tmp["others"] = t.Others
	}
	if t.Reboots != nil {
		tmp["reboots"] = t.Reboots
	}
	if t.Total != nil {
		tmp["total"] = t.Total
	}
	if t.TotalApsWAlarm != nil {
		tmp["totalApsWAlarm"] = t.TotalApsWAlarm
	}
	if t.TotalApsWReboot != nil {
		tmp["totalApsWReboot"] = t.TotalApsWReboot
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport46apInventoryOverviewDataType() *SCIInventoryAPsReport46apInventoryOverviewDataType {
	m := new(SCIInventoryAPsReport46apInventoryOverviewDataType)
	return m
}

// SCIInventoryAPsReport47topApsDisconnectionData
//
// Definition: InventoryAPsReport_InventoryAPsReport_47_topApsDisconnection_Data
type SCIInventoryAPsReport47topApsDisconnectionData []*SCIInventoryAPsReport47topApsDisconnectionDataType

func MakeSCIInventoryAPsReport47topApsDisconnectionData() SCIInventoryAPsReport47topApsDisconnectionData {
	m := make(SCIInventoryAPsReport47topApsDisconnectionData, 0)
	return m
}

// SCIInventoryAPsReport47topApsDisconnectionDataType
//
// Definition: InventoryAPsReport_InventoryAPsReport_47_topApsDisconnection_DataType
type SCIInventoryAPsReport47topApsDisconnectionDataType struct {
	Key *string `json:"key,omitempty"`

	Label *string `json:"label,omitempty"`

	Value *float64 `json:"value,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport47topApsDisconnectionDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport47topApsDisconnectionDataType SCIInventoryAPsReport47topApsDisconnectionDataType
	tmpType := new(_SCIInventoryAPsReport47topApsDisconnectionDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "key")
	delete(tmpType.XAdditionalProperties, "label")
	delete(tmpType.XAdditionalProperties, "value")
	*t = SCIInventoryAPsReport47topApsDisconnectionDataType(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport47topApsDisconnectionDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Key != nil {
		tmp["key"] = t.Key
	}
	if t.Label != nil {
		tmp["label"] = t.Label
	}
	if t.Value != nil {
		tmp["value"] = t.Value
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport47topApsDisconnectionDataType() *SCIInventoryAPsReport47topApsDisconnectionDataType {
	m := new(SCIInventoryAPsReport47topApsDisconnectionDataType)
	return m
}

// SCIInventoryAPsReport47topApsDisconnectionMetaData
//
// Definition: InventoryAPsReport_InventoryAPsReport_47_topApsDisconnection_MetaData
type SCIInventoryAPsReport47topApsDisconnectionMetaData struct {
	Legend *SCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType `json:"legend,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport47topApsDisconnectionMetaData) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport47topApsDisconnectionMetaData SCIInventoryAPsReport47topApsDisconnectionMetaData
	tmpType := new(_SCIInventoryAPsReport47topApsDisconnectionMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "legend")
	*t = SCIInventoryAPsReport47topApsDisconnectionMetaData(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport47topApsDisconnectionMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Legend != nil {
		tmp["legend"] = t.Legend
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport47topApsDisconnectionMetaData() *SCIInventoryAPsReport47topApsDisconnectionMetaData {
	m := new(SCIInventoryAPsReport47topApsDisconnectionMetaData)
	return m
}

// SCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType
//
// Definition: InventoryAPsReport_InventoryAPsReport_47_topApsDisconnection_MetaDataLegendType
type SCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType struct {
	On *string `json:"on,omitempty"`
}

func NewSCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType() *SCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType {
	m := new(SCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType)
	return m
}

// SCIInventoryAPsReport48apCountTrendData
//
// Definition: InventoryAPsReport_InventoryAPsReport_48_apCountTrend_Data
type SCIInventoryAPsReport48apCountTrendData []*SCIInventoryAPsReport48apCountTrendDataType

func MakeSCIInventoryAPsReport48apCountTrendData() SCIInventoryAPsReport48apCountTrendData {
	m := make(SCIInventoryAPsReport48apCountTrendData, 0)
	return m
}

// SCIInventoryAPsReport48apCountTrendDataType
//
// Definition: InventoryAPsReport_InventoryAPsReport_48_apCountTrend_DataType
type SCIInventoryAPsReport48apCountTrendDataType struct {
	ConnectedAPs *float64 `json:"connectedAPs,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalAPs *float64 `json:"totalAPs,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport48apCountTrendDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport48apCountTrendDataType SCIInventoryAPsReport48apCountTrendDataType
	tmpType := new(_SCIInventoryAPsReport48apCountTrendDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "connectedAPs")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "totalAPs")
	*t = SCIInventoryAPsReport48apCountTrendDataType(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport48apCountTrendDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ConnectedAPs != nil {
		tmp["connectedAPs"] = t.ConnectedAPs
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.TotalAPs != nil {
		tmp["totalAPs"] = t.TotalAPs
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport48apCountTrendDataType() *SCIInventoryAPsReport48apCountTrendDataType {
	m := new(SCIInventoryAPsReport48apCountTrendDataType)
	return m
}

// SCIInventoryAPsReport49apStatusTrendData
//
// Definition: InventoryAPsReport_InventoryAPsReport_49_apStatusTrend_Data
type SCIInventoryAPsReport49apStatusTrendData [][]*SCIInventoryAPsReport49apStatusTrendDataTypeType

func MakeSCIInventoryAPsReport49apStatusTrendData() SCIInventoryAPsReport49apStatusTrendData {
	m := make(SCIInventoryAPsReport49apStatusTrendData, 0)
	return m
}

// SCIInventoryAPsReport49apStatusTrendDataType
//
// Definition: InventoryAPsReport_InventoryAPsReport_49_apStatusTrend_DataType
type SCIInventoryAPsReport49apStatusTrendDataType []*SCIInventoryAPsReport49apStatusTrendDataTypeType

func MakeSCIInventoryAPsReport49apStatusTrendDataType() SCIInventoryAPsReport49apStatusTrendDataType {
	m := make(SCIInventoryAPsReport49apStatusTrendDataType, 0)
	return m
}

// SCIInventoryAPsReport49apStatusTrendDataTypeType
//
// Definition: InventoryAPsReport_InventoryAPsReport_49_apStatusTrend_DataTypeType
type SCIInventoryAPsReport49apStatusTrendDataTypeType struct {
	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	Count *float64 `json:"count,omitempty"`

	End *string `json:"end,omitempty"`

	Offline *float64 `json:"Offline,omitempty"`

	Online *float64 `json:"Online,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport49apStatusTrendDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport49apStatusTrendDataTypeType SCIInventoryAPsReport49apStatusTrendDataTypeType
	tmpType := new(_SCIInventoryAPsReport49apStatusTrendDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "connectionStatus")
	delete(tmpType.XAdditionalProperties, "count")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "Offline")
	delete(tmpType.XAdditionalProperties, "Online")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCIInventoryAPsReport49apStatusTrendDataTypeType(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport49apStatusTrendDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ConnectionStatus != nil {
		tmp["connectionStatus"] = t.ConnectionStatus
	}
	if t.Count != nil {
		tmp["count"] = t.Count
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Offline != nil {
		tmp["Offline"] = t.Offline
	}
	if t.Online != nil {
		tmp["Online"] = t.Online
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport49apStatusTrendDataTypeType() *SCIInventoryAPsReport49apStatusTrendDataTypeType {
	m := new(SCIInventoryAPsReport49apStatusTrendDataTypeType)
	return m
}

// SCIInventoryAPsReport49apStatusTrendMetaData
//
// Definition: InventoryAPsReport_InventoryAPsReport_49_apStatusTrend_MetaData
type SCIInventoryAPsReport49apStatusTrendMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport49apStatusTrendMetaData) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport49apStatusTrendMetaData SCIInventoryAPsReport49apStatusTrendMetaData
	tmpType := new(_SCIInventoryAPsReport49apStatusTrendMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	*t = SCIInventoryAPsReport49apStatusTrendMetaData(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport49apStatusTrendMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIInventoryAPsReport49apStatusTrendMetaData() *SCIInventoryAPsReport49apStatusTrendMetaData {
	m := new(SCIInventoryAPsReport49apStatusTrendMetaData)
	return m
}

// SCIInventoryAPsReport50topApsModelsChartData
//
// Definition: InventoryAPsReport_InventoryAPsReport_50_topApsModelsChart_Data
type SCIInventoryAPsReport50topApsModelsChartData [][]*SCIInventoryAPsReport50topApsModelsChartDataTypeType

func MakeSCIInventoryAPsReport50topApsModelsChartData() SCIInventoryAPsReport50topApsModelsChartData {
	m := make(SCIInventoryAPsReport50topApsModelsChartData, 0)
	return m
}

// SCIInventoryAPsReport50topApsModelsChartDataType
//
// Definition: InventoryAPsReport_InventoryAPsReport_50_topApsModelsChart_DataType
type SCIInventoryAPsReport50topApsModelsChartDataType []*SCIInventoryAPsReport50topApsModelsChartDataTypeType

func MakeSCIInventoryAPsReport50topApsModelsChartDataType() SCIInventoryAPsReport50topApsModelsChartDataType {
	m := make(SCIInventoryAPsReport50topApsModelsChartDataType, 0)
	return m
}

// SCIInventoryAPsReport50topApsModelsChartDataTypeType
//
// Definition: InventoryAPsReport_InventoryAPsReport_50_topApsModelsChart_DataTypeType
type SCIInventoryAPsReport50topApsModelsChartDataTypeType struct {
	ApCount *float64 `json:"apCount,omitempty"`

	ApModel *string `json:"apModel,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport50topApsModelsChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport50topApsModelsChartDataTypeType SCIInventoryAPsReport50topApsModelsChartDataTypeType
	tmpType := new(_SCIInventoryAPsReport50topApsModelsChartDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apCount")
	delete(tmpType.XAdditionalProperties, "apModel")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCIInventoryAPsReport50topApsModelsChartDataTypeType(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport50topApsModelsChartDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.ApModel != nil {
		tmp["apModel"] = t.ApModel
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport50topApsModelsChartDataTypeType() *SCIInventoryAPsReport50topApsModelsChartDataTypeType {
	m := new(SCIInventoryAPsReport50topApsModelsChartDataTypeType)
	return m
}

// SCIInventoryAPsReport50topApsModelsChartMetaData
//
// Definition: InventoryAPsReport_InventoryAPsReport_50_topApsModelsChart_MetaData
type SCIInventoryAPsReport50topApsModelsChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport50topApsModelsChartMetaData) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport50topApsModelsChartMetaData SCIInventoryAPsReport50topApsModelsChartMetaData
	tmpType := new(_SCIInventoryAPsReport50topApsModelsChartMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	*t = SCIInventoryAPsReport50topApsModelsChartMetaData(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport50topApsModelsChartMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIInventoryAPsReport50topApsModelsChartMetaData() *SCIInventoryAPsReport50topApsModelsChartMetaData {
	m := new(SCIInventoryAPsReport50topApsModelsChartMetaData)
	return m
}

// SCIInventoryAPsReport51top10ApVersionsChartData
//
// Definition: InventoryAPsReport_InventoryAPsReport_51_top10ApVersionsChart_Data
type SCIInventoryAPsReport51top10ApVersionsChartData [][]*SCIInventoryAPsReport51top10ApVersionsChartDataTypeType

func MakeSCIInventoryAPsReport51top10ApVersionsChartData() SCIInventoryAPsReport51top10ApVersionsChartData {
	m := make(SCIInventoryAPsReport51top10ApVersionsChartData, 0)
	return m
}

// SCIInventoryAPsReport51top10ApVersionsChartDataType
//
// Definition: InventoryAPsReport_InventoryAPsReport_51_top10ApVersionsChart_DataType
type SCIInventoryAPsReport51top10ApVersionsChartDataType []*SCIInventoryAPsReport51top10ApVersionsChartDataTypeType

func MakeSCIInventoryAPsReport51top10ApVersionsChartDataType() SCIInventoryAPsReport51top10ApVersionsChartDataType {
	m := make(SCIInventoryAPsReport51top10ApVersionsChartDataType, 0)
	return m
}

// SCIInventoryAPsReport51top10ApVersionsChartDataTypeType
//
// Definition: InventoryAPsReport_InventoryAPsReport_51_top10ApVersionsChart_DataTypeType
type SCIInventoryAPsReport51top10ApVersionsChartDataTypeType struct {
	ApCount *float64 `json:"apCount,omitempty"`

	ApFwVersion *string `json:"apFwVersion,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	Unknown *float64 `json:"Unknown,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport51top10ApVersionsChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport51top10ApVersionsChartDataTypeType SCIInventoryAPsReport51top10ApVersionsChartDataTypeType
	tmpType := new(_SCIInventoryAPsReport51top10ApVersionsChartDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apCount")
	delete(tmpType.XAdditionalProperties, "apFwVersion")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "Unknown")
	*t = SCIInventoryAPsReport51top10ApVersionsChartDataTypeType(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport51top10ApVersionsChartDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.ApFwVersion != nil {
		tmp["apFwVersion"] = t.ApFwVersion
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.Unknown != nil {
		tmp["Unknown"] = t.Unknown
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport51top10ApVersionsChartDataTypeType() *SCIInventoryAPsReport51top10ApVersionsChartDataTypeType {
	m := new(SCIInventoryAPsReport51top10ApVersionsChartDataTypeType)
	return m
}

// SCIInventoryAPsReport51top10ApVersionsChartMetaData
//
// Definition: InventoryAPsReport_InventoryAPsReport_51_top10ApVersionsChart_MetaData
type SCIInventoryAPsReport51top10ApVersionsChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport51top10ApVersionsChartMetaData) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport51top10ApVersionsChartMetaData SCIInventoryAPsReport51top10ApVersionsChartMetaData
	tmpType := new(_SCIInventoryAPsReport51top10ApVersionsChartMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	*t = SCIInventoryAPsReport51top10ApVersionsChartMetaData(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport51top10ApVersionsChartMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIInventoryAPsReport51top10ApVersionsChartMetaData() *SCIInventoryAPsReport51top10ApVersionsChartMetaData {
	m := new(SCIInventoryAPsReport51top10ApVersionsChartMetaData)
	return m
}

// SCIInventoryAPsReport55topAPModelsData
//
// Definition: InventoryAPsReport_InventoryAPsReport_55_topAPModels_Data
type SCIInventoryAPsReport55topAPModelsData []*SCIInventoryAPsReport55topAPModelsDataType

func MakeSCIInventoryAPsReport55topAPModelsData() SCIInventoryAPsReport55topAPModelsData {
	m := make(SCIInventoryAPsReport55topAPModelsData, 0)
	return m
}

// SCIInventoryAPsReport55topAPModelsDataType
//
// Definition: InventoryAPsReport_InventoryAPsReport_55_topAPModels_DataType
type SCIInventoryAPsReport55topAPModelsDataType struct {
	ApCount *float64 `json:"apCount,omitempty"`

	ApModel *string `json:"apModel,omitempty"`

	ApPercent *float64 `json:"apPercent,omitempty"`

	Index *float64 `json:"index,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport55topAPModelsDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport55topAPModelsDataType SCIInventoryAPsReport55topAPModelsDataType
	tmpType := new(_SCIInventoryAPsReport55topAPModelsDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apCount")
	delete(tmpType.XAdditionalProperties, "apModel")
	delete(tmpType.XAdditionalProperties, "apPercent")
	delete(tmpType.XAdditionalProperties, "index")
	*t = SCIInventoryAPsReport55topAPModelsDataType(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport55topAPModelsDataType) MarshalJSON() ([]byte, error) {
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
	if t.ApModel != nil {
		tmp["apModel"] = t.ApModel
	}
	if t.ApPercent != nil {
		tmp["apPercent"] = t.ApPercent
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport55topAPModelsDataType() *SCIInventoryAPsReport55topAPModelsDataType {
	m := new(SCIInventoryAPsReport55topAPModelsDataType)
	return m
}

// SCIInventoryAPsReport56topAPVersionsData
//
// Definition: InventoryAPsReport_InventoryAPsReport_56_topAPVersions_Data
type SCIInventoryAPsReport56topAPVersionsData []*SCIInventoryAPsReport56topAPVersionsDataType

func MakeSCIInventoryAPsReport56topAPVersionsData() SCIInventoryAPsReport56topAPVersionsData {
	m := make(SCIInventoryAPsReport56topAPVersionsData, 0)
	return m
}

// SCIInventoryAPsReport56topAPVersionsDataType
//
// Definition: InventoryAPsReport_InventoryAPsReport_56_topAPVersions_DataType
type SCIInventoryAPsReport56topAPVersionsDataType struct {
	ApCount *float64 `json:"apCount,omitempty"`

	ApFwVersion *string `json:"apFwVersion,omitempty"`

	ApPercent *float64 `json:"apPercent,omitempty"`

	Index *float64 `json:"index,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport56topAPVersionsDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport56topAPVersionsDataType SCIInventoryAPsReport56topAPVersionsDataType
	tmpType := new(_SCIInventoryAPsReport56topAPVersionsDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apCount")
	delete(tmpType.XAdditionalProperties, "apFwVersion")
	delete(tmpType.XAdditionalProperties, "apPercent")
	delete(tmpType.XAdditionalProperties, "index")
	*t = SCIInventoryAPsReport56topAPVersionsDataType(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport56topAPVersionsDataType) MarshalJSON() ([]byte, error) {
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
	if t.ApFwVersion != nil {
		tmp["apFwVersion"] = t.ApFwVersion
	}
	if t.ApPercent != nil {
		tmp["apPercent"] = t.ApPercent
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport56topAPVersionsDataType() *SCIInventoryAPsReport56topAPVersionsDataType {
	m := new(SCIInventoryAPsReport56topAPVersionsDataType)
	return m
}

// SCIInventoryAPsReport57topAPsOfflineData
//
// Definition: InventoryAPsReport_InventoryAPsReport_57_topAPsOffline_Data
type SCIInventoryAPsReport57topAPsOfflineData []*SCIInventoryAPsReport57topAPsOfflineDataType

func MakeSCIInventoryAPsReport57topAPsOfflineData() SCIInventoryAPsReport57topAPsOfflineData {
	m := make(SCIInventoryAPsReport57topAPsOfflineData, 0)
	return m
}

// SCIInventoryAPsReport57topAPsOfflineDataType
//
// Definition: InventoryAPsReport_InventoryAPsReport_57_topAPsOffline_DataType
type SCIInventoryAPsReport57topAPsOfflineDataType struct {
	ApIp *string `json:"apIp,omitempty"`

	ApLocation *string `json:"apLocation,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApModel *string `json:"apModel,omitempty"`

	ApName *string `json:"apName,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	DisconnectDuration *float64 `json:"disconnectDuration,omitempty"`

	Index *float64 `json:"index,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport57topAPsOfflineDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport57topAPsOfflineDataType SCIInventoryAPsReport57topAPsOfflineDataType
	tmpType := new(_SCIInventoryAPsReport57topAPsOfflineDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apIp")
	delete(tmpType.XAdditionalProperties, "apLocation")
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apModel")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "ctrlName")
	delete(tmpType.XAdditionalProperties, "disconnectDuration")
	delete(tmpType.XAdditionalProperties, "index")
	*t = SCIInventoryAPsReport57topAPsOfflineDataType(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport57topAPsOfflineDataType) MarshalJSON() ([]byte, error) {
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
	if t.ApLocation != nil {
		tmp["apLocation"] = t.ApLocation
	}
	if t.ApMac != nil {
		tmp["apMac"] = t.ApMac
	}
	if t.ApModel != nil {
		tmp["apModel"] = t.ApModel
	}
	if t.ApName != nil {
		tmp["apName"] = t.ApName
	}
	if t.CtrlName != nil {
		tmp["ctrlName"] = t.CtrlName
	}
	if t.DisconnectDuration != nil {
		tmp["disconnectDuration"] = t.DisconnectDuration
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport57topAPsOfflineDataType() *SCIInventoryAPsReport57topAPsOfflineDataType {
	m := new(SCIInventoryAPsReport57topAPsOfflineDataType)
	return m
}

// SCIInventoryAPsReport57topAPsOfflineMetaData
//
// Definition: InventoryAPsReport_InventoryAPsReport_57_topAPsOffline_MetaData
type SCIInventoryAPsReport57topAPsOfflineMetaData struct {
	MaxValues *SCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType `json:"maxValues,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport57topAPsOfflineMetaData) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport57topAPsOfflineMetaData SCIInventoryAPsReport57topAPsOfflineMetaData
	tmpType := new(_SCIInventoryAPsReport57topAPsOfflineMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	*t = SCIInventoryAPsReport57topAPsOfflineMetaData(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport57topAPsOfflineMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIInventoryAPsReport57topAPsOfflineMetaData() *SCIInventoryAPsReport57topAPsOfflineMetaData {
	m := new(SCIInventoryAPsReport57topAPsOfflineMetaData)
	return m
}

// SCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType
//
// Definition: InventoryAPsReport_InventoryAPsReport_57_topAPsOffline_MetaDataMaxValuesType
type SCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType struct {
	DisconnectDuration *float64 `json:"disconnectDuration,omitempty"`
}

func NewSCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType() *SCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType {
	m := new(SCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType)
	return m
}

// SCIInventoryAPsReport60apDetailsOnOfflineStatusData
//
// Definition: InventoryAPsReport_InventoryAPsReport_60_apDetailsOnOfflineStatus_Data
type SCIInventoryAPsReport60apDetailsOnOfflineStatusData []*SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType

func MakeSCIInventoryAPsReport60apDetailsOnOfflineStatusData() SCIInventoryAPsReport60apDetailsOnOfflineStatusData {
	m := make(SCIInventoryAPsReport60apDetailsOnOfflineStatusData, 0)
	return m
}

// SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType
//
// Definition: InventoryAPsReport_InventoryAPsReport_60_apDetailsOnOfflineStatus_DataType
type SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType struct {
	ApIp *string `json:"apIp,omitempty"`

	ApLocation *string `json:"apLocation,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApModel *string `json:"apModel,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	Index *float64 `json:"index,omitempty"`

	LastStatusChangeTime *string `json:"lastStatusChangeTime,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType
	tmpType := new(_SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apIp")
	delete(tmpType.XAdditionalProperties, "apLocation")
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apModel")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "connectionStatus")
	delete(tmpType.XAdditionalProperties, "ctrlName")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "lastStatusChangeTime")
	*t = SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType(*tmpType)
	return nil
}

func (t *SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType) MarshalJSON() ([]byte, error) {
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
	if t.ApLocation != nil {
		tmp["apLocation"] = t.ApLocation
	}
	if t.ApMac != nil {
		tmp["apMac"] = t.ApMac
	}
	if t.ApModel != nil {
		tmp["apModel"] = t.ApModel
	}
	if t.ApName != nil {
		tmp["apName"] = t.ApName
	}
	if t.ConnectionStatus != nil {
		tmp["connectionStatus"] = t.ConnectionStatus
	}
	if t.CtrlName != nil {
		tmp["ctrlName"] = t.CtrlName
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.LastStatusChangeTime != nil {
		tmp["lastStatusChangeTime"] = t.LastStatusChangeTime
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryAPsReport60apDetailsOnOfflineStatusDataType() *SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType {
	m := new(SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType)
	return m
}

// ReportInventoryAPsReport46ApInventoryOverview
//
// Inventory - APs Report - Overview
//
// Operation ID: report_InventoryAPsReport_46_apInventoryOverview
// Operation path: /reports/9/sections/46/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport46ApInventoryOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport46apInventoryOverview200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport46apInventoryOverview200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport46ApInventoryOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport46apInventoryOverview200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport47TopApsDisconnection
//
// Inventory - APs Report - Top APs by Offline Duration
//
// Operation ID: report_InventoryAPsReport_47_topApsDisconnection
// Operation path: /reports/9/sections/47/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport47TopApsDisconnection(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport47topApsDisconnection200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport47topApsDisconnection200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport47TopApsDisconnection, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport47topApsDisconnection200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport48ApCountTrend
//
// Inventory - APs Report - AP Count Trend
//
// Operation ID: report_InventoryAPsReport_48_apCountTrend
// Operation path: /reports/9/sections/48/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport48ApCountTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport48apCountTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport48apCountTrend200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport48ApCountTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport48apCountTrend200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport49ApStatusTrend
//
// Inventory - APs Report - AP Status Trends
//
// Operation ID: report_InventoryAPsReport_49_apStatusTrend
// Operation path: /reports/9/sections/49/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport49ApStatusTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport49apStatusTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport49apStatusTrend200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport49ApStatusTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport49apStatusTrend200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport50TopApsModelsChart
//
// Inventory - APs Report - Top AP Models
//
// Operation ID: report_InventoryAPsReport_50_topApsModelsChart
// Operation path: /reports/9/sections/50/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport50TopApsModelsChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport50topApsModelsChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport50topApsModelsChart200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport50TopApsModelsChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport50topApsModelsChart200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport51Top10ApVersionsChart
//
// Inventory - APs Report - Top AP Software Versions
//
// Operation ID: report_InventoryAPsReport_51_top10ApVersionsChart
// Operation path: /reports/9/sections/51/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport51Top10ApVersionsChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport51top10ApVersionsChart200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport51Top10ApVersionsChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport52TopApsRebootReasons
//
// Inventory - APs Report - Top 10 AP Reboot Reasons
//
// Operation ID: report_InventoryAPsReport_52_topApsRebootReasons
// Operation path: /reports/9/sections/52/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport52TopApsRebootReasons(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport52topApsRebootReasons200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport52topApsRebootReasons200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport52TopApsRebootReasons, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport52topApsRebootReasons200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport53Top10ApsRebootCounts
//
// Inventory - APs Report - Top APs by Reboot Count
//
// Operation ID: report_InventoryAPsReport_53_top10ApsRebootCounts
// Operation path: /reports/9/sections/53/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport53Top10ApsRebootCounts(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport53Top10ApsRebootCounts, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport54TopApAlarmTypes
//
// Inventory - APs Report - Top 10 AP Alarm Types
//
// Operation ID: report_InventoryAPsReport_54_topApAlarmTypes
// Operation path: /reports/9/sections/54/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport54TopApAlarmTypes(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport54topApAlarmTypes200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport54topApAlarmTypes200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport54TopApAlarmTypes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport54topApAlarmTypes200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport55TopAPModels
//
// Inventory - APs Report - Top AP Models
//
// Operation ID: report_InventoryAPsReport_55_topAPModels
// Operation path: /reports/9/sections/55/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport55TopAPModels(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport55topAPModels200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport55topAPModels200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport55TopAPModels, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport55topAPModels200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport56TopAPVersions
//
// Inventory - APs Report - Top AP Software Versions
//
// Operation ID: report_InventoryAPsReport_56_topAPVersions
// Operation path: /reports/9/sections/56/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport56TopAPVersions(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport56topAPVersions200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport56topAPVersions200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport56TopAPVersions, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport56topAPVersions200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport57TopAPsOffline
//
// Inventory - APs Report - Top APs by Offline Duration
//
// Operation ID: report_InventoryAPsReport_57_topAPsOffline
// Operation path: /reports/9/sections/57/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport57TopAPsOffline(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport57topAPsOffline200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport57topAPsOffline200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport57TopAPsOffline, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport57topAPsOffline200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport58TopAPsByReboots
//
// Inventory - APs Report - Top APs by Reboot Count
//
// Operation ID: report_InventoryAPsReport_58_topAPsByReboots
// Operation path: /reports/9/sections/58/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport58TopAPsByReboots(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport58topAPsByReboots200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport58topAPsByReboots200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport58TopAPsByReboots, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport58topAPsByReboots200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport59ApsConfiguredInMultiCtrl
//
// Inventory - APs Report - APs Configured in Multiple Systems
//
// Operation ID: report_InventoryAPsReport_59_apsConfiguredInMultiCtrl
// Operation path: /reports/9/sections/59/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport59ApsConfiguredInMultiCtrl(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport59ApsConfiguredInMultiCtrl, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport60ApDetailsOnOfflineStatus
//
// Inventory - APs Report - AP Details for Online/Offline Status
//
// Operation ID: report_InventoryAPsReport_60_apDetailsOnOfflineStatus
// Operation path: /reports/9/sections/60/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport60ApDetailsOnOfflineStatus(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport60ApDetailsOnOfflineStatus, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseTypeAPIResponse), err
}

// ReportInventoryAPsReport61ApDetailsOtherStatus
//
// Inventory - APs Report - AP Details for Other Statuses
//
// Operation ID: report_InventoryAPsReport_61_apDetailsOtherStatus
// Operation path: /reports/9/sections/61/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport61ApDetailsOtherStatus(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventoryAPsReport61ApDetailsOtherStatus, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseTypeAPIResponse), err
}
