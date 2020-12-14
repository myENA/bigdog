package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
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
// Definition: Overview_Overview_62_overview_Data
type SCIOverview62overviewData []*SCIOverview62overviewDataType

func MakeSCIOverview62overviewData() SCIOverview62overviewData {
	m := make(SCIOverview62overviewData, 0)
	return m
}

// SCIOverview62overviewDataType
//
// Definition: Overview_Overview_62_overview_DataType
type SCIOverview62overviewDataType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ClientCount interface{} `json:"clientCount,omitempty"`

	RebootCount *float64 `json:"rebootCount,omitempty"`

	SessionCount interface{} `json:"sessionCount,omitempty"`

	ShortSessionRatio interface{} `json:"shortSessionRatio,omitempty"`

	Total *float64 `json:"total,omitempty"`

	TotalHistory *float64 `json:"totalHistory,omitempty"`

	UserTraffic interface{} `json:"userTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview62overviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIOverview62overviewDataType SCIOverview62overviewDataType
	tmpType := new(_SCIOverview62overviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "rebootCount")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "shortSessionRatio")
	delete(tmpType.XAdditionalProperties, "total")
	delete(tmpType.XAdditionalProperties, "totalHistory")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	*t = SCIOverview62overviewDataType(*tmpType)
	return nil
}

func (t *SCIOverview62overviewDataType) MarshalJSON() ([]byte, error) {
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
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.RebootCount != nil {
		tmp["rebootCount"] = t.RebootCount
	}
	if t.SessionCount != nil {
		tmp["sessionCount"] = t.SessionCount
	}
	if t.ShortSessionRatio != nil {
		tmp["shortSessionRatio"] = t.ShortSessionRatio
	}
	if t.Total != nil {
		tmp["total"] = t.Total
	}
	if t.TotalHistory != nil {
		tmp["totalHistory"] = t.TotalHistory
	}
	if t.UserTraffic != nil {
		tmp["userTraffic"] = t.UserTraffic
	}
	return json.Marshal(tmp)
}

func NewSCIOverview62overviewDataType() *SCIOverview62overviewDataType {
	m := new(SCIOverview62overviewDataType)
	return m
}

// SCIOverview62overviewMetaData
//
// Definition: Overview_Overview_62_overview_MetaData
type SCIOverview62overviewMetaData struct {
	LastAnomalyUpdatedTime *string `json:"lastAnomalyUpdatedTime,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview62overviewMetaData) UnmarshalJSON(b []byte) error {
	type _SCIOverview62overviewMetaData SCIOverview62overviewMetaData
	tmpType := new(_SCIOverview62overviewMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "lastAnomalyUpdatedTime")
	*t = SCIOverview62overviewMetaData(*tmpType)
	return nil
}

func (t *SCIOverview62overviewMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.LastAnomalyUpdatedTime != nil {
		tmp["lastAnomalyUpdatedTime"] = t.LastAnomalyUpdatedTime
	}
	return json.Marshal(tmp)
}

func NewSCIOverview62overviewMetaData() *SCIOverview62overviewMetaData {
	m := new(SCIOverview62overviewMetaData)
	return m
}

// SCIOverview63controllerData
//
// Definition: Overview_Overview_63_controller_Data
type SCIOverview63controllerData []*SCIOverview63controllerDataType

func MakeSCIOverview63controllerData() SCIOverview63controllerData {
	m := make(SCIOverview63controllerData, 0)
	return m
}

// SCIOverview63controllerDataType
//
// Definition: Overview_Overview_63_controller_DataType
type SCIOverview63controllerDataType struct {
	Offline *float64 `json:"offline,omitempty"`

	Online *float64 `json:"online,omitempty"`

	SzCount *float64 `json:"szCount,omitempty"`

	Total *float64 `json:"total,omitempty"`

	ZdCount *float64 `json:"zdCount,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview63controllerDataType) UnmarshalJSON(b []byte) error {
	type _SCIOverview63controllerDataType SCIOverview63controllerDataType
	tmpType := new(_SCIOverview63controllerDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "offline")
	delete(tmpType.XAdditionalProperties, "online")
	delete(tmpType.XAdditionalProperties, "szCount")
	delete(tmpType.XAdditionalProperties, "total")
	delete(tmpType.XAdditionalProperties, "zdCount")
	*t = SCIOverview63controllerDataType(*tmpType)
	return nil
}

func (t *SCIOverview63controllerDataType) MarshalJSON() ([]byte, error) {
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
	if t.SzCount != nil {
		tmp["szCount"] = t.SzCount
	}
	if t.Total != nil {
		tmp["total"] = t.Total
	}
	if t.ZdCount != nil {
		tmp["zdCount"] = t.ZdCount
	}
	return json.Marshal(tmp)
}

func NewSCIOverview63controllerDataType() *SCIOverview63controllerDataType {
	m := new(SCIOverview63controllerDataType)
	return m
}

// SCIOverview64apOverviewData
//
// Definition: Overview_Overview_64_apOverview_Data
type SCIOverview64apOverviewData []*SCIOverview64apOverviewDataType

func MakeSCIOverview64apOverviewData() SCIOverview64apOverviewData {
	m := make(SCIOverview64apOverviewData, 0)
	return m
}

// SCIOverview64apOverviewDataType
//
// Definition: Overview_Overview_64_apOverview_DataType
type SCIOverview64apOverviewDataType struct {
	Offline *float64 `json:"offline,omitempty"`

	Online *float64 `json:"online,omitempty"`

	Others *float64 `json:"others,omitempty"`

	Reboots *float64 `json:"reboots,omitempty"`

	Total *float64 `json:"total,omitempty"`

	TotalApsWAlarm *float64 `json:"totalApsWAlarm,omitempty"`

	TotalApsWReboot *float64 `json:"totalApsWReboot,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview64apOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIOverview64apOverviewDataType SCIOverview64apOverviewDataType
	tmpType := new(_SCIOverview64apOverviewDataType)
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
	*t = SCIOverview64apOverviewDataType(*tmpType)
	return nil
}

func (t *SCIOverview64apOverviewDataType) MarshalJSON() ([]byte, error) {
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

func NewSCIOverview64apOverviewDataType() *SCIOverview64apOverviewDataType {
	m := new(SCIOverview64apOverviewDataType)
	return m
}

// SCIOverview66apAlarmOverviewData
//
// Definition: Overview_Overview_66_apAlarmOverview_Data
type SCIOverview66apAlarmOverviewData []*SCIOverview66apAlarmOverviewDataType

func MakeSCIOverview66apAlarmOverviewData() SCIOverview66apAlarmOverviewData {
	m := make(SCIOverview66apAlarmOverviewData, 0)
	return m
}

// SCIOverview66apAlarmOverviewDataType
//
// Definition: Overview_Overview_66_apAlarmOverview_DataType
type SCIOverview66apAlarmOverviewDataType struct {
	AlarmType *string `json:"alarmType,omitempty"`

	Count *float64 `json:"count,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview66apAlarmOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIOverview66apAlarmOverviewDataType SCIOverview66apAlarmOverviewDataType
	tmpType := new(_SCIOverview66apAlarmOverviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "alarmType")
	delete(tmpType.XAdditionalProperties, "count")
	*t = SCIOverview66apAlarmOverviewDataType(*tmpType)
	return nil
}

func (t *SCIOverview66apAlarmOverviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AlarmType != nil {
		tmp["alarmType"] = t.AlarmType
	}
	if t.Count != nil {
		tmp["count"] = t.Count
	}
	return json.Marshal(tmp)
}

func NewSCIOverview66apAlarmOverviewDataType() *SCIOverview66apAlarmOverviewDataType {
	m := new(SCIOverview66apAlarmOverviewDataType)
	return m
}

// SCIOverview66apAlarmOverviewMetaData
//
// Definition: Overview_Overview_66_apAlarmOverview_MetaData
type SCIOverview66apAlarmOverviewMetaData struct {
	AlarmStat *SCIOverview66apAlarmOverviewMetaDataAlarmStatType `json:"alarmStat,omitempty"`

	AlarmTotalCount *float64 `json:"alarmTotalCount,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview66apAlarmOverviewMetaData) UnmarshalJSON(b []byte) error {
	type _SCIOverview66apAlarmOverviewMetaData SCIOverview66apAlarmOverviewMetaData
	tmpType := new(_SCIOverview66apAlarmOverviewMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "alarmStat")
	delete(tmpType.XAdditionalProperties, "alarmTotalCount")
	*t = SCIOverview66apAlarmOverviewMetaData(*tmpType)
	return nil
}

func (t *SCIOverview66apAlarmOverviewMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AlarmStat != nil {
		tmp["alarmStat"] = t.AlarmStat
	}
	if t.AlarmTotalCount != nil {
		tmp["alarmTotalCount"] = t.AlarmTotalCount
	}
	return json.Marshal(tmp)
}

func NewSCIOverview66apAlarmOverviewMetaData() *SCIOverview66apAlarmOverviewMetaData {
	m := new(SCIOverview66apAlarmOverviewMetaData)
	return m
}

// SCIOverview66apAlarmOverviewMetaDataAlarmStatType
//
// Definition: Overview_Overview_66_apAlarmOverview_MetaDataAlarmStatType
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
// Definition: Overview_Overview_66_apAlarmOverview_MetaDataAlarmStatTypeByApGroupType
type SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByApGroupType struct {
	Alarm *float64 `json:"alarm,omitempty"`

	Total *float64 `json:"total,omitempty"`
}

func NewSCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByApGroupType() *SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByApGroupType {
	m := new(SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByApGroupType)
	return m
}

// SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType
//
// Definition: Overview_Overview_66_apAlarmOverview_MetaDataAlarmStatTypeByZoneType
type SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType struct {
	Alarm *float64 `json:"alarm,omitempty"`

	Total *float64 `json:"total,omitempty"`
}

func NewSCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType() *SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType {
	m := new(SCIOverview66apAlarmOverviewMetaDataAlarmStatTypeByZoneType)
	return m
}

// SCIOverview67switchOverviewData
//
// Definition: Overview_Overview_67_switchOverview_Data
type SCIOverview67switchOverviewData []*SCIOverview67switchOverviewDataType

func MakeSCIOverview67switchOverviewData() SCIOverview67switchOverviewData {
	m := make(SCIOverview67switchOverviewData, 0)
	return m
}

// SCIOverview67switchOverviewDataType
//
// Definition: Overview_Overview_67_switchOverview_DataType
type SCIOverview67switchOverviewDataType struct {
	Offline *float64 `json:"offline,omitempty"`

	Online *float64 `json:"online,omitempty"`

	Others *float64 `json:"others,omitempty"`

	Total *float64 `json:"total,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview67switchOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIOverview67switchOverviewDataType SCIOverview67switchOverviewDataType
	tmpType := new(_SCIOverview67switchOverviewDataType)
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
	delete(tmpType.XAdditionalProperties, "total")
	*t = SCIOverview67switchOverviewDataType(*tmpType)
	return nil
}

func (t *SCIOverview67switchOverviewDataType) MarshalJSON() ([]byte, error) {
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
	if t.Total != nil {
		tmp["total"] = t.Total
	}
	return json.Marshal(tmp)
}

func NewSCIOverview67switchOverviewDataType() *SCIOverview67switchOverviewDataType {
	m := new(SCIOverview67switchOverviewDataType)
	return m
}

// SCIOverview68apClientCountOverviewData
//
// Definition: Overview_Overview_68_apClientCountOverview_Data
type SCIOverview68apClientCountOverviewData []*SCIOverview68apClientCountOverviewDataType

func MakeSCIOverview68apClientCountOverviewData() SCIOverview68apClientCountOverviewData {
	m := make(SCIOverview68apClientCountOverviewData, 0)
	return m
}

// SCIOverview68apClientCountOverviewDataType
//
// Definition: Overview_Overview_68_apClientCountOverview_DataType
type SCIOverview68apClientCountOverviewDataType struct {
	Key *string `json:"key,omitempty"`

	Label *string `json:"label,omitempty"`

	Value *float64 `json:"value,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview68apClientCountOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIOverview68apClientCountOverviewDataType SCIOverview68apClientCountOverviewDataType
	tmpType := new(_SCIOverview68apClientCountOverviewDataType)
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
	*t = SCIOverview68apClientCountOverviewDataType(*tmpType)
	return nil
}

func (t *SCIOverview68apClientCountOverviewDataType) MarshalJSON() ([]byte, error) {
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

func NewSCIOverview68apClientCountOverviewDataType() *SCIOverview68apClientCountOverviewDataType {
	m := new(SCIOverview68apClientCountOverviewDataType)
	return m
}

// SCIOverview68apClientCountOverviewMetaData
//
// Definition: Overview_Overview_68_apClientCountOverview_MetaData
type SCIOverview68apClientCountOverviewMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview68apClientCountOverviewMetaData) UnmarshalJSON(b []byte) error {
	type _SCIOverview68apClientCountOverviewMetaData SCIOverview68apClientCountOverviewMetaData
	tmpType := new(_SCIOverview68apClientCountOverviewMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	*t = SCIOverview68apClientCountOverviewMetaData(*tmpType)
	return nil
}

func (t *SCIOverview68apClientCountOverviewMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIOverview68apClientCountOverviewMetaData() *SCIOverview68apClientCountOverviewMetaData {
	m := new(SCIOverview68apClientCountOverviewMetaData)
	return m
}

// SCIOverview69totalTrafficMinMaxRateData
//
// Definition: Overview_Overview_69_totalTrafficMinMaxRate_Data
type SCIOverview69totalTrafficMinMaxRateData [][]*SCIOverview69totalTrafficMinMaxRateDataTypeType

func MakeSCIOverview69totalTrafficMinMaxRateData() SCIOverview69totalTrafficMinMaxRateData {
	m := make(SCIOverview69totalTrafficMinMaxRateData, 0)
	return m
}

// SCIOverview69totalTrafficMinMaxRateDataType
//
// Definition: Overview_Overview_69_totalTrafficMinMaxRate_DataType
type SCIOverview69totalTrafficMinMaxRateDataType []*SCIOverview69totalTrafficMinMaxRateDataTypeType

func MakeSCIOverview69totalTrafficMinMaxRateDataType() SCIOverview69totalTrafficMinMaxRateDataType {
	m := make(SCIOverview69totalTrafficMinMaxRateDataType, 0)
	return m
}

// SCIOverview69totalTrafficMinMaxRateDataTypeType
//
// Definition: Overview_Overview_69_totalTrafficMinMaxRate_DataTypeType
type SCIOverview69totalTrafficMinMaxRateDataTypeType struct {
	Avg *float64 `json:"avg,omitempty"`

	Max *float64 `json:"max,omitempty"`

	MaxLabel *string `json:"maxLabel,omitempty"`

	Min *float64 `json:"min,omitempty"`

	MinLabel *string `json:"minLabel,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview69totalTrafficMinMaxRateDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIOverview69totalTrafficMinMaxRateDataTypeType SCIOverview69totalTrafficMinMaxRateDataTypeType
	tmpType := new(_SCIOverview69totalTrafficMinMaxRateDataTypeType)
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
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	*t = SCIOverview69totalTrafficMinMaxRateDataTypeType(*tmpType)
	return nil
}

func (t *SCIOverview69totalTrafficMinMaxRateDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.TotalRxTraffic != nil {
		tmp["totalRxTraffic"] = t.TotalRxTraffic
	}
	if t.TotalTxTraffic != nil {
		tmp["totalTxTraffic"] = t.TotalTxTraffic
	}
	return json.Marshal(tmp)
}

func NewSCIOverview69totalTrafficMinMaxRateDataTypeType() *SCIOverview69totalTrafficMinMaxRateDataTypeType {
	m := new(SCIOverview69totalTrafficMinMaxRateDataTypeType)
	return m
}

// SCIOverview70sessionsOverviewData
//
// Definition: Overview_Overview_70_sessionsOverview_Data
type SCIOverview70sessionsOverviewData [][]*SCIOverview70sessionsOverviewDataTypeType

func MakeSCIOverview70sessionsOverviewData() SCIOverview70sessionsOverviewData {
	m := make(SCIOverview70sessionsOverviewData, 0)
	return m
}

// SCIOverview70sessionsOverviewDataType
//
// Definition: Overview_Overview_70_sessionsOverview_DataType
type SCIOverview70sessionsOverviewDataType []*SCIOverview70sessionsOverviewDataTypeType

func MakeSCIOverview70sessionsOverviewDataType() SCIOverview70sessionsOverviewDataType {
	m := make(SCIOverview70sessionsOverviewDataType, 0)
	return m
}

// SCIOverview70sessionsOverviewDataTypeType
//
// Definition: Overview_Overview_70_sessionsOverview_DataTypeType
type SCIOverview70sessionsOverviewDataTypeType struct {
	Avg *float64 `json:"avg,omitempty"`

	Current *float64 `json:"current,omitempty"`

	Max *float64 `json:"max,omitempty"`

	MaxLabel *string `json:"maxLabel,omitempty"`

	Min *float64 `json:"min,omitempty"`

	MinLabel *string `json:"minLabel,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview70sessionsOverviewDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIOverview70sessionsOverviewDataTypeType SCIOverview70sessionsOverviewDataTypeType
	tmpType := new(_SCIOverview70sessionsOverviewDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avg")
	delete(tmpType.XAdditionalProperties, "current")
	delete(tmpType.XAdditionalProperties, "max")
	delete(tmpType.XAdditionalProperties, "maxLabel")
	delete(tmpType.XAdditionalProperties, "min")
	delete(tmpType.XAdditionalProperties, "minLabel")
	*t = SCIOverview70sessionsOverviewDataTypeType(*tmpType)
	return nil
}

func (t *SCIOverview70sessionsOverviewDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.Current != nil {
		tmp["current"] = t.Current
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

func NewSCIOverview70sessionsOverviewDataTypeType() *SCIOverview70sessionsOverviewDataTypeType {
	m := new(SCIOverview70sessionsOverviewDataTypeType)
	return m
}

// SCIOverview71ssidOverviewData
//
// Definition: Overview_Overview_71_ssidOverview_Data
type SCIOverview71ssidOverviewData [][]*SCIOverview71ssidOverviewDataTypeType

func MakeSCIOverview71ssidOverviewData() SCIOverview71ssidOverviewData {
	m := make(SCIOverview71ssidOverviewData, 0)
	return m
}

// SCIOverview71ssidOverviewDataType
//
// Definition: Overview_Overview_71_ssidOverview_DataType
type SCIOverview71ssidOverviewDataType []*SCIOverview71ssidOverviewDataTypeType

func MakeSCIOverview71ssidOverviewDataType() SCIOverview71ssidOverviewDataType {
	m := make(SCIOverview71ssidOverviewDataType, 0)
	return m
}

// SCIOverview71ssidOverviewDataTypeType
//
// Definition: Overview_Overview_71_ssidOverview_DataTypeType
type SCIOverview71ssidOverviewDataTypeType struct {
	Current *float64 `json:"current,omitempty"`

	Previous *float64 `json:"previous,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	TotalMetric *float64 `json:"totalMetric,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview71ssidOverviewDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIOverview71ssidOverviewDataTypeType SCIOverview71ssidOverviewDataTypeType
	tmpType := new(_SCIOverview71ssidOverviewDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "current")
	delete(tmpType.XAdditionalProperties, "previous")
	delete(tmpType.XAdditionalProperties, "ssid")
	delete(tmpType.XAdditionalProperties, "totalMetric")
	*t = SCIOverview71ssidOverviewDataTypeType(*tmpType)
	return nil
}

func (t *SCIOverview71ssidOverviewDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Current != nil {
		tmp["current"] = t.Current
	}
	if t.Previous != nil {
		tmp["previous"] = t.Previous
	}
	if t.Ssid != nil {
		tmp["ssid"] = t.Ssid
	}
	if t.TotalMetric != nil {
		tmp["totalMetric"] = t.TotalMetric
	}
	return json.Marshal(tmp)
}

func NewSCIOverview71ssidOverviewDataTypeType() *SCIOverview71ssidOverviewDataTypeType {
	m := new(SCIOverview71ssidOverviewDataTypeType)
	return m
}

// SCIOverview71ssidOverviewMetaData
//
// Definition: Overview_Overview_71_ssidOverview_MetaData
type SCIOverview71ssidOverviewMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalMetric *float64 `json:"totalMetric,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview71ssidOverviewMetaData) UnmarshalJSON(b []byte) error {
	type _SCIOverview71ssidOverviewMetaData SCIOverview71ssidOverviewMetaData
	tmpType := new(_SCIOverview71ssidOverviewMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	delete(tmpType.XAdditionalProperties, "totalMetric")
	*t = SCIOverview71ssidOverviewMetaData(*tmpType)
	return nil
}

func (t *SCIOverview71ssidOverviewMetaData) MarshalJSON() ([]byte, error) {
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
	if t.TotalMetric != nil {
		tmp["totalMetric"] = t.TotalMetric
	}
	return json.Marshal(tmp)
}

func NewSCIOverview71ssidOverviewMetaData() *SCIOverview71ssidOverviewMetaData {
	m := new(SCIOverview71ssidOverviewMetaData)
	return m
}

// SCIOverview72radioOverviewData
//
// Definition: Overview_Overview_72_radioOverview_Data
type SCIOverview72radioOverviewData []*SCIOverview72radioOverviewDataType

func MakeSCIOverview72radioOverviewData() SCIOverview72radioOverviewData {
	m := make(SCIOverview72radioOverviewData, 0)
	return m
}

// SCIOverview72radioOverviewDataType
//
// Definition: Overview_Overview_72_radioOverview_DataType
type SCIOverview72radioOverviewDataType struct {
	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	Radio *string `json:"radio,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview72radioOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIOverview72radioOverviewDataType SCIOverview72radioOverviewDataType
	tmpType := new(_SCIOverview72radioOverviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "airtimeUtilizationAvg")
	delete(tmpType.XAdditionalProperties, "radio")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	*t = SCIOverview72radioOverviewDataType(*tmpType)
	return nil
}

func (t *SCIOverview72radioOverviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AirtimeUtilizationAvg != nil {
		tmp["airtimeUtilizationAvg"] = t.AirtimeUtilizationAvg
	}
	if t.Radio != nil {
		tmp["radio"] = t.Radio
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	return json.Marshal(tmp)
}

func NewSCIOverview72radioOverviewDataType() *SCIOverview72radioOverviewDataType {
	m := new(SCIOverview72radioOverviewDataType)
	return m
}

// SCIOverview73applicationsOverviewData
//
// Definition: Overview_Overview_73_applicationsOverview_Data
type SCIOverview73applicationsOverviewData []*SCIOverview73applicationsOverviewDataType

func MakeSCIOverview73applicationsOverviewData() SCIOverview73applicationsOverviewData {
	m := make(SCIOverview73applicationsOverviewData, 0)
	return m
}

// SCIOverview73applicationsOverviewDataType
//
// Definition: Overview_Overview_73_applicationsOverview_DataType
type SCIOverview73applicationsOverviewDataType struct {
	Key *float64 `json:"key,omitempty"`

	Label *string `json:"label,omitempty"`

	Value *float64 `json:"value,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview73applicationsOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIOverview73applicationsOverviewDataType SCIOverview73applicationsOverviewDataType
	tmpType := new(_SCIOverview73applicationsOverviewDataType)
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
	*t = SCIOverview73applicationsOverviewDataType(*tmpType)
	return nil
}

func (t *SCIOverview73applicationsOverviewDataType) MarshalJSON() ([]byte, error) {
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

func NewSCIOverview73applicationsOverviewDataType() *SCIOverview73applicationsOverviewDataType {
	m := new(SCIOverview73applicationsOverviewDataType)
	return m
}

// SCIOverview74apEventOverviewData
//
// Definition: Overview_Overview_74_apEventOverview_Data
type SCIOverview74apEventOverviewData []*SCIOverview74apEventOverviewDataType

func MakeSCIOverview74apEventOverviewData() SCIOverview74apEventOverviewData {
	m := make(SCIOverview74apEventOverviewData, 0)
	return m
}

// SCIOverview74apEventOverviewDataType
//
// Definition: Overview_Overview_74_apEventOverview_DataType
type SCIOverview74apEventOverviewDataType struct {
	Count *float64 `json:"count,omitempty"`

	EventType *string `json:"eventType,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview74apEventOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIOverview74apEventOverviewDataType SCIOverview74apEventOverviewDataType
	tmpType := new(_SCIOverview74apEventOverviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "count")
	delete(tmpType.XAdditionalProperties, "eventType")
	*t = SCIOverview74apEventOverviewDataType(*tmpType)
	return nil
}

func (t *SCIOverview74apEventOverviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Count != nil {
		tmp["count"] = t.Count
	}
	if t.EventType != nil {
		tmp["eventType"] = t.EventType
	}
	return json.Marshal(tmp)
}

func NewSCIOverview74apEventOverviewDataType() *SCIOverview74apEventOverviewDataType {
	m := new(SCIOverview74apEventOverviewDataType)
	return m
}

// SCIOverview74apEventOverviewMetaData
//
// Definition: Overview_Overview_74_apEventOverview_MetaData
type SCIOverview74apEventOverviewMetaData struct {
	EventTotalCount *float64 `json:"eventTotalCount,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview74apEventOverviewMetaData) UnmarshalJSON(b []byte) error {
	type _SCIOverview74apEventOverviewMetaData SCIOverview74apEventOverviewMetaData
	tmpType := new(_SCIOverview74apEventOverviewMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "eventTotalCount")
	*t = SCIOverview74apEventOverviewMetaData(*tmpType)
	return nil
}

func (t *SCIOverview74apEventOverviewMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.EventTotalCount != nil {
		tmp["eventTotalCount"] = t.EventTotalCount
	}
	return json.Marshal(tmp)
}

func NewSCIOverview74apEventOverviewMetaData() *SCIOverview74apEventOverviewMetaData {
	m := new(SCIOverview74apEventOverviewMetaData)
	return m
}

// SCIOverview97factOverviewData
//
// Definition: Overview_Overview_97_factOverview_Data
type SCIOverview97factOverviewData []*SCIOverview97factOverviewDataType

func MakeSCIOverview97factOverviewData() SCIOverview97factOverviewData {
	m := make(SCIOverview97factOverviewData, 0)
	return m
}

// SCIOverview97factOverviewDataType
//
// Definition: Overview_Overview_97_factOverview_DataType
type SCIOverview97factOverviewDataType struct {
	Key *string `json:"key,omitempty"`

	Label *string `json:"label,omitempty"`

	Value *float64 `json:"value,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview97factOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIOverview97factOverviewDataType SCIOverview97factOverviewDataType
	tmpType := new(_SCIOverview97factOverviewDataType)
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
	*t = SCIOverview97factOverviewDataType(*tmpType)
	return nil
}

func (t *SCIOverview97factOverviewDataType) MarshalJSON() ([]byte, error) {
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

func NewSCIOverview97factOverviewDataType() *SCIOverview97factOverviewDataType {
	m := new(SCIOverview97factOverviewDataType)
	return m
}

// SCIOverview115networkUsageOverviewData
//
// Definition: Overview_Overview_115_networkUsageOverview_Data
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

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview115networkUsageOverviewData) UnmarshalJSON(b []byte) error {
	type _SCIOverview115networkUsageOverviewData SCIOverview115networkUsageOverviewData
	tmpType := new(_SCIOverview115networkUsageOverviewData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "Ap")
	delete(tmpType.XAdditionalProperties, "ApGroup")
	delete(tmpType.XAdditionalProperties, "App")
	delete(tmpType.XAdditionalProperties, "Domain")
	delete(tmpType.XAdditionalProperties, "OS")
	delete(tmpType.XAdditionalProperties, "SSID")
	delete(tmpType.XAdditionalProperties, "Switch")
	delete(tmpType.XAdditionalProperties, "System")
	delete(tmpType.XAdditionalProperties, "Zone")
	*t = SCIOverview115networkUsageOverviewData(*tmpType)
	return nil
}

func (t *SCIOverview115networkUsageOverviewData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Ap != nil {
		tmp["Ap"] = t.Ap
	}
	if t.ApGroup != nil {
		tmp["ApGroup"] = t.ApGroup
	}
	if t.App != nil {
		tmp["App"] = t.App
	}
	if t.Domain != nil {
		tmp["Domain"] = t.Domain
	}
	if t.OS != nil {
		tmp["OS"] = t.OS
	}
	if t.SSID != nil {
		tmp["SSID"] = t.SSID
	}
	if t.Switch != nil {
		tmp["Switch"] = t.Switch
	}
	if t.System != nil {
		tmp["System"] = t.System
	}
	if t.Zone != nil {
		tmp["Zone"] = t.Zone
	}
	return json.Marshal(tmp)
}

func NewSCIOverview115networkUsageOverviewData() *SCIOverview115networkUsageOverviewData {
	m := new(SCIOverview115networkUsageOverviewData)
	return m
}

// SCIOverview115networkUsageOverviewDataApGroupType
//
// Definition: Overview_Overview_115_networkUsageOverview_DataApGroupType
type SCIOverview115networkUsageOverviewDataApGroupType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *float64 `json:"x,omitempty"`

	Y *float64 `json:"y,omitempty"`

	Z *float64 `json:"z,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview115networkUsageOverviewDataApGroupType) UnmarshalJSON(b []byte) error {
	type _SCIOverview115networkUsageOverviewDataApGroupType SCIOverview115networkUsageOverviewDataApGroupType
	tmpType := new(_SCIOverview115networkUsageOverviewDataApGroupType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "macId")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "x")
	delete(tmpType.XAdditionalProperties, "y")
	delete(tmpType.XAdditionalProperties, "z")
	*t = SCIOverview115networkUsageOverviewDataApGroupType(*tmpType)
	return nil
}

func (t *SCIOverview115networkUsageOverviewDataApGroupType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MacId != nil {
		tmp["macId"] = t.MacId
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.X != nil {
		tmp["x"] = t.X
	}
	if t.Y != nil {
		tmp["y"] = t.Y
	}
	if t.Z != nil {
		tmp["z"] = t.Z
	}
	return json.Marshal(tmp)
}

func NewSCIOverview115networkUsageOverviewDataApGroupType() *SCIOverview115networkUsageOverviewDataApGroupType {
	m := new(SCIOverview115networkUsageOverviewDataApGroupType)
	return m
}

// SCIOverview115networkUsageOverviewDataAppType
//
// Definition: Overview_Overview_115_networkUsageOverview_DataAppType
type SCIOverview115networkUsageOverviewDataAppType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *float64 `json:"x,omitempty"`

	Y *float64 `json:"y,omitempty"`

	Z *float64 `json:"z,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview115networkUsageOverviewDataAppType) UnmarshalJSON(b []byte) error {
	type _SCIOverview115networkUsageOverviewDataAppType SCIOverview115networkUsageOverviewDataAppType
	tmpType := new(_SCIOverview115networkUsageOverviewDataAppType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "macId")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "x")
	delete(tmpType.XAdditionalProperties, "y")
	delete(tmpType.XAdditionalProperties, "z")
	*t = SCIOverview115networkUsageOverviewDataAppType(*tmpType)
	return nil
}

func (t *SCIOverview115networkUsageOverviewDataAppType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MacId != nil {
		tmp["macId"] = t.MacId
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.X != nil {
		tmp["x"] = t.X
	}
	if t.Y != nil {
		tmp["y"] = t.Y
	}
	if t.Z != nil {
		tmp["z"] = t.Z
	}
	return json.Marshal(tmp)
}

func NewSCIOverview115networkUsageOverviewDataAppType() *SCIOverview115networkUsageOverviewDataAppType {
	m := new(SCIOverview115networkUsageOverviewDataAppType)
	return m
}

// SCIOverview115networkUsageOverviewDataApType
//
// Definition: Overview_Overview_115_networkUsageOverview_DataApType
type SCIOverview115networkUsageOverviewDataApType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *float64 `json:"x,omitempty"`

	Y *float64 `json:"y,omitempty"`

	Z *float64 `json:"z,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview115networkUsageOverviewDataApType) UnmarshalJSON(b []byte) error {
	type _SCIOverview115networkUsageOverviewDataApType SCIOverview115networkUsageOverviewDataApType
	tmpType := new(_SCIOverview115networkUsageOverviewDataApType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "macId")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "x")
	delete(tmpType.XAdditionalProperties, "y")
	delete(tmpType.XAdditionalProperties, "z")
	*t = SCIOverview115networkUsageOverviewDataApType(*tmpType)
	return nil
}

func (t *SCIOverview115networkUsageOverviewDataApType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MacId != nil {
		tmp["macId"] = t.MacId
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.X != nil {
		tmp["x"] = t.X
	}
	if t.Y != nil {
		tmp["y"] = t.Y
	}
	if t.Z != nil {
		tmp["z"] = t.Z
	}
	return json.Marshal(tmp)
}

func NewSCIOverview115networkUsageOverviewDataApType() *SCIOverview115networkUsageOverviewDataApType {
	m := new(SCIOverview115networkUsageOverviewDataApType)
	return m
}

// SCIOverview115networkUsageOverviewDataDomainType
//
// Definition: Overview_Overview_115_networkUsageOverview_DataDomainType
type SCIOverview115networkUsageOverviewDataDomainType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *float64 `json:"x,omitempty"`

	Y *float64 `json:"y,omitempty"`

	Z *float64 `json:"z,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview115networkUsageOverviewDataDomainType) UnmarshalJSON(b []byte) error {
	type _SCIOverview115networkUsageOverviewDataDomainType SCIOverview115networkUsageOverviewDataDomainType
	tmpType := new(_SCIOverview115networkUsageOverviewDataDomainType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "macId")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "x")
	delete(tmpType.XAdditionalProperties, "y")
	delete(tmpType.XAdditionalProperties, "z")
	*t = SCIOverview115networkUsageOverviewDataDomainType(*tmpType)
	return nil
}

func (t *SCIOverview115networkUsageOverviewDataDomainType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MacId != nil {
		tmp["macId"] = t.MacId
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.X != nil {
		tmp["x"] = t.X
	}
	if t.Y != nil {
		tmp["y"] = t.Y
	}
	if t.Z != nil {
		tmp["z"] = t.Z
	}
	return json.Marshal(tmp)
}

func NewSCIOverview115networkUsageOverviewDataDomainType() *SCIOverview115networkUsageOverviewDataDomainType {
	m := new(SCIOverview115networkUsageOverviewDataDomainType)
	return m
}

// SCIOverview115networkUsageOverviewDataOSType
//
// Definition: Overview_Overview_115_networkUsageOverview_DataOSType
type SCIOverview115networkUsageOverviewDataOSType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *float64 `json:"x,omitempty"`

	Y *float64 `json:"y,omitempty"`

	Z *float64 `json:"z,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview115networkUsageOverviewDataOSType) UnmarshalJSON(b []byte) error {
	type _SCIOverview115networkUsageOverviewDataOSType SCIOverview115networkUsageOverviewDataOSType
	tmpType := new(_SCIOverview115networkUsageOverviewDataOSType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "macId")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "x")
	delete(tmpType.XAdditionalProperties, "y")
	delete(tmpType.XAdditionalProperties, "z")
	*t = SCIOverview115networkUsageOverviewDataOSType(*tmpType)
	return nil
}

func (t *SCIOverview115networkUsageOverviewDataOSType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MacId != nil {
		tmp["macId"] = t.MacId
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.X != nil {
		tmp["x"] = t.X
	}
	if t.Y != nil {
		tmp["y"] = t.Y
	}
	if t.Z != nil {
		tmp["z"] = t.Z
	}
	return json.Marshal(tmp)
}

func NewSCIOverview115networkUsageOverviewDataOSType() *SCIOverview115networkUsageOverviewDataOSType {
	m := new(SCIOverview115networkUsageOverviewDataOSType)
	return m
}

// SCIOverview115networkUsageOverviewDataSSIDType
//
// Definition: Overview_Overview_115_networkUsageOverview_DataSSIDType
type SCIOverview115networkUsageOverviewDataSSIDType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *float64 `json:"x,omitempty"`

	Y *float64 `json:"y,omitempty"`

	Z *float64 `json:"z,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview115networkUsageOverviewDataSSIDType) UnmarshalJSON(b []byte) error {
	type _SCIOverview115networkUsageOverviewDataSSIDType SCIOverview115networkUsageOverviewDataSSIDType
	tmpType := new(_SCIOverview115networkUsageOverviewDataSSIDType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "macId")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "x")
	delete(tmpType.XAdditionalProperties, "y")
	delete(tmpType.XAdditionalProperties, "z")
	*t = SCIOverview115networkUsageOverviewDataSSIDType(*tmpType)
	return nil
}

func (t *SCIOverview115networkUsageOverviewDataSSIDType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MacId != nil {
		tmp["macId"] = t.MacId
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.X != nil {
		tmp["x"] = t.X
	}
	if t.Y != nil {
		tmp["y"] = t.Y
	}
	if t.Z != nil {
		tmp["z"] = t.Z
	}
	return json.Marshal(tmp)
}

func NewSCIOverview115networkUsageOverviewDataSSIDType() *SCIOverview115networkUsageOverviewDataSSIDType {
	m := new(SCIOverview115networkUsageOverviewDataSSIDType)
	return m
}

// SCIOverview115networkUsageOverviewDataSwitchType
//
// Definition: Overview_Overview_115_networkUsageOverview_DataSwitchType
type SCIOverview115networkUsageOverviewDataSwitchType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *float64 `json:"x,omitempty"`

	Y *float64 `json:"y,omitempty"`

	Z *float64 `json:"z,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview115networkUsageOverviewDataSwitchType) UnmarshalJSON(b []byte) error {
	type _SCIOverview115networkUsageOverviewDataSwitchType SCIOverview115networkUsageOverviewDataSwitchType
	tmpType := new(_SCIOverview115networkUsageOverviewDataSwitchType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "macId")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "x")
	delete(tmpType.XAdditionalProperties, "y")
	delete(tmpType.XAdditionalProperties, "z")
	*t = SCIOverview115networkUsageOverviewDataSwitchType(*tmpType)
	return nil
}

func (t *SCIOverview115networkUsageOverviewDataSwitchType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MacId != nil {
		tmp["macId"] = t.MacId
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.X != nil {
		tmp["x"] = t.X
	}
	if t.Y != nil {
		tmp["y"] = t.Y
	}
	if t.Z != nil {
		tmp["z"] = t.Z
	}
	return json.Marshal(tmp)
}

func NewSCIOverview115networkUsageOverviewDataSwitchType() *SCIOverview115networkUsageOverviewDataSwitchType {
	m := new(SCIOverview115networkUsageOverviewDataSwitchType)
	return m
}

// SCIOverview115networkUsageOverviewDataSystemType
//
// Definition: Overview_Overview_115_networkUsageOverview_DataSystemType
type SCIOverview115networkUsageOverviewDataSystemType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *float64 `json:"x,omitempty"`

	Y *float64 `json:"y,omitempty"`

	Z *float64 `json:"z,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview115networkUsageOverviewDataSystemType) UnmarshalJSON(b []byte) error {
	type _SCIOverview115networkUsageOverviewDataSystemType SCIOverview115networkUsageOverviewDataSystemType
	tmpType := new(_SCIOverview115networkUsageOverviewDataSystemType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "macId")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "x")
	delete(tmpType.XAdditionalProperties, "y")
	delete(tmpType.XAdditionalProperties, "z")
	*t = SCIOverview115networkUsageOverviewDataSystemType(*tmpType)
	return nil
}

func (t *SCIOverview115networkUsageOverviewDataSystemType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MacId != nil {
		tmp["macId"] = t.MacId
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.X != nil {
		tmp["x"] = t.X
	}
	if t.Y != nil {
		tmp["y"] = t.Y
	}
	if t.Z != nil {
		tmp["z"] = t.Z
	}
	return json.Marshal(tmp)
}

func NewSCIOverview115networkUsageOverviewDataSystemType() *SCIOverview115networkUsageOverviewDataSystemType {
	m := new(SCIOverview115networkUsageOverviewDataSystemType)
	return m
}

// SCIOverview115networkUsageOverviewDataZoneType
//
// Definition: Overview_Overview_115_networkUsageOverview_DataZoneType
type SCIOverview115networkUsageOverviewDataZoneType struct {
	MacId *string `json:"macId,omitempty"`

	Name *string `json:"name,omitempty"`

	X *float64 `json:"x,omitempty"`

	Y *float64 `json:"y,omitempty"`

	Z *float64 `json:"z,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIOverview115networkUsageOverviewDataZoneType) UnmarshalJSON(b []byte) error {
	type _SCIOverview115networkUsageOverviewDataZoneType SCIOverview115networkUsageOverviewDataZoneType
	tmpType := new(_SCIOverview115networkUsageOverviewDataZoneType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "macId")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "x")
	delete(tmpType.XAdditionalProperties, "y")
	delete(tmpType.XAdditionalProperties, "z")
	*t = SCIOverview115networkUsageOverviewDataZoneType(*tmpType)
	return nil
}

func (t *SCIOverview115networkUsageOverviewDataZoneType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MacId != nil {
		tmp["macId"] = t.MacId
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.X != nil {
		tmp["x"] = t.X
	}
	if t.Y != nil {
		tmp["y"] = t.Y
	}
	if t.Z != nil {
		tmp["z"] = t.Z
	}
	return json.Marshal(tmp)
}

func NewSCIOverview115networkUsageOverviewDataZoneType() *SCIOverview115networkUsageOverviewDataZoneType {
	m := new(SCIOverview115networkUsageOverviewDataZoneType)
	return m
}

// ReportOverview62Overview
//
// Operation ID: report_Overview_62_overview
//
// Overview - Ruckus SmartAnalytics
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview62Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview62overview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview62Overview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_63_controller
//
// Overview - Controllers
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview63Controller(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview63controller200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview63Controller, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_64_apOverview
//
// Overview - Access Points
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview64ApOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview64apOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview64ApOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_66_apAlarmOverview
//
// Overview - Alarms
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview66ApAlarmOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview66apAlarmOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview66ApAlarmOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_67_switchOverview
//
// Overview - Switches
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview67SwitchOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview67switchOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview67SwitchOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_68_apClientCountOverview
//
// Overview - Top APs by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview68ApClientCountOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview68apClientCountOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview68ApClientCountOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_69_totalTrafficMinMaxRate
//
// Overview - Total Wireless Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview69TotalTrafficMinMaxRate(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview69totalTrafficMinMaxRate200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview69TotalTrafficMinMaxRate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_70_sessionsOverview
//
// Overview - Unique Wireless Sessions
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview70SessionsOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview70sessionsOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview70SessionsOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_71_ssidOverview
//
// Overview - WLANs
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview71SsidOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview71ssidOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview71SsidOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_72_radioOverview
//
// Overview - Radios
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview72RadioOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview72radioOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview72RadioOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_73_applicationsOverview
//
// Overview - Applications (Wireless)
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview73ApplicationsOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview73applicationsOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview73ApplicationsOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_74_apEventOverview
//
// Overview - Events
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview74ApEventOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview74apEventOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview74ApEventOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_97_factOverview
//
// Overview - Did you know?
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview97FactOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview97factOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview97FactOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_Overview_115_networkUsageOverview
//
// Overview - Network Usage Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIOverviewService) ReportOverview115NetworkUsageOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportOverview115networkUsageOverview200ResponseType, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportOverview115NetworkUsageOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportOverview115networkUsageOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
