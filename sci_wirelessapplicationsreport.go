package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIWirelessApplicationsReportService struct {
	apiClient *SCIClient
}

func NewSCIWirelessApplicationsReportService(c *SCIClient) *SCIWirelessApplicationsReportService {
	s := new(SCIWirelessApplicationsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIWirelessApplicationsReportService() *SCIWirelessApplicationsReportService {
	return NewSCIWirelessApplicationsReportService(ss.apiClient)
}

// SCIWirelessApplicationsReport8topAppsByTrafficTableData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_8_topAppsByTrafficTable_Data
type SCIWirelessApplicationsReport8topAppsByTrafficTableData []*SCIWirelessApplicationsReport8topAppsByTrafficTableDataType

func MakeSCIWirelessApplicationsReport8topAppsByTrafficTableData() SCIWirelessApplicationsReport8topAppsByTrafficTableData {
	m := make(SCIWirelessApplicationsReport8topAppsByTrafficTableData, 0)
	return m
}

// SCIWirelessApplicationsReport8topAppsByTrafficTableDataType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_8_topAppsByTrafficTable_DataType
type SCIWirelessApplicationsReport8topAppsByTrafficTableDataType struct {
	App *string `json:"app,omitempty"`

	Index *float64 `json:"index,omitempty"`

	Port *string `json:"port,omitempty"`

	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessApplicationsReport8topAppsByTrafficTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessApplicationsReport8topAppsByTrafficTableDataType SCIWirelessApplicationsReport8topAppsByTrafficTableDataType
	tmpType := new(_SCIWirelessApplicationsReport8topAppsByTrafficTableDataType)
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
	*t = SCIWirelessApplicationsReport8topAppsByTrafficTableDataType(*tmpType)
	return nil
}

func (t *SCIWirelessApplicationsReport8topAppsByTrafficTableDataType) MarshalJSON() ([]byte, error) {
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

func NewSCIWirelessApplicationsReport8topAppsByTrafficTableDataType() *SCIWirelessApplicationsReport8topAppsByTrafficTableDataType {
	m := new(SCIWirelessApplicationsReport8topAppsByTrafficTableDataType)
	return m
}

// SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_8_topAppsByTrafficTable_MetaData
type SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData struct {
	MaxValues *SCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData
	tmpType := new(_SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData)
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
	*t = SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData(*tmpType)
	return nil
}

func (t *SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIWirelessApplicationsReport8topAppsByTrafficTableMetaData() *SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData {
	m := new(SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData)
	return m
}

// SCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_8_topAppsByTrafficTable_MetaDataMaxValuesType
type SCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType struct {
	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`
}

func NewSCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType() *SCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType {
	m := new(SCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCIWirelessApplicationsReport9topAppsByClientsTableData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_9_topAppsByClientsTable_Data
type SCIWirelessApplicationsReport9topAppsByClientsTableData []*SCIWirelessApplicationsReport9topAppsByClientsTableDataType

func MakeSCIWirelessApplicationsReport9topAppsByClientsTableData() SCIWirelessApplicationsReport9topAppsByClientsTableData {
	m := make(SCIWirelessApplicationsReport9topAppsByClientsTableData, 0)
	return m
}

// SCIWirelessApplicationsReport9topAppsByClientsTableDataType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_9_topAppsByClientsTable_DataType
type SCIWirelessApplicationsReport9topAppsByClientsTableDataType struct {
	App *string `json:"app,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	Index *float64 `json:"index,omitempty"`

	Port *string `json:"port,omitempty"`

	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessApplicationsReport9topAppsByClientsTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessApplicationsReport9topAppsByClientsTableDataType SCIWirelessApplicationsReport9topAppsByClientsTableDataType
	tmpType := new(_SCIWirelessApplicationsReport9topAppsByClientsTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "app")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "port")
	delete(tmpType.XAdditionalProperties, "rxBytes")
	delete(tmpType.XAdditionalProperties, "traffic")
	delete(tmpType.XAdditionalProperties, "txBytes")
	*t = SCIWirelessApplicationsReport9topAppsByClientsTableDataType(*tmpType)
	return nil
}

func (t *SCIWirelessApplicationsReport9topAppsByClientsTableDataType) MarshalJSON() ([]byte, error) {
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
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
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

func NewSCIWirelessApplicationsReport9topAppsByClientsTableDataType() *SCIWirelessApplicationsReport9topAppsByClientsTableDataType {
	m := new(SCIWirelessApplicationsReport9topAppsByClientsTableDataType)
	return m
}

// SCIWirelessApplicationsReport9topAppsByClientsTableMetaData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_9_topAppsByClientsTable_MetaData
type SCIWirelessApplicationsReport9topAppsByClientsTableMetaData struct {
	MaxValues *SCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessApplicationsReport9topAppsByClientsTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWirelessApplicationsReport9topAppsByClientsTableMetaData SCIWirelessApplicationsReport9topAppsByClientsTableMetaData
	tmpType := new(_SCIWirelessApplicationsReport9topAppsByClientsTableMetaData)
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
	*t = SCIWirelessApplicationsReport9topAppsByClientsTableMetaData(*tmpType)
	return nil
}

func (t *SCIWirelessApplicationsReport9topAppsByClientsTableMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIWirelessApplicationsReport9topAppsByClientsTableMetaData() *SCIWirelessApplicationsReport9topAppsByClientsTableMetaData {
	m := new(SCIWirelessApplicationsReport9topAppsByClientsTableMetaData)
	return m
}

// SCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_9_topAppsByClientsTable_MetaDataMaxValuesType
type SCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType struct {
	RxBytes *float64 `json:"rxBytes,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`
}

func NewSCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType() *SCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType {
	m := new(SCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType)
	return m
}

// SCIWirelessApplicationsReport10overviewData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_10_overview_Data
type SCIWirelessApplicationsReport10overviewData [][]*SCIWirelessApplicationsReport10overviewDataTypeType

func MakeSCIWirelessApplicationsReport10overviewData() SCIWirelessApplicationsReport10overviewData {
	m := make(SCIWirelessApplicationsReport10overviewData, 0)
	return m
}

// SCIWirelessApplicationsReport10overviewDataType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_10_overview_DataType
type SCIWirelessApplicationsReport10overviewDataType []*SCIWirelessApplicationsReport10overviewDataTypeType

func MakeSCIWirelessApplicationsReport10overviewDataType() SCIWirelessApplicationsReport10overviewDataType {
	m := make(SCIWirelessApplicationsReport10overviewDataType, 0)
	return m
}

// SCIWirelessApplicationsReport10overviewDataTypeType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_10_overview_DataTypeType
type SCIWirelessApplicationsReport10overviewDataTypeType struct {
	Applications *float64 `json:"applications,omitempty"`

	TotalRxBytes *float64 `json:"totalRxBytes,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxBytes *float64 `json:"totalTxBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessApplicationsReport10overviewDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessApplicationsReport10overviewDataTypeType SCIWirelessApplicationsReport10overviewDataTypeType
	tmpType := new(_SCIWirelessApplicationsReport10overviewDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "applications")
	delete(tmpType.XAdditionalProperties, "totalRxBytes")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxBytes")
	*t = SCIWirelessApplicationsReport10overviewDataTypeType(*tmpType)
	return nil
}

func (t *SCIWirelessApplicationsReport10overviewDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Applications != nil {
		tmp["applications"] = t.Applications
	}
	if t.TotalRxBytes != nil {
		tmp["totalRxBytes"] = t.TotalRxBytes
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.TotalTxBytes != nil {
		tmp["totalTxBytes"] = t.TotalTxBytes
	}
	return json.Marshal(tmp)
}

func NewSCIWirelessApplicationsReport10overviewDataTypeType() *SCIWirelessApplicationsReport10overviewDataTypeType {
	m := new(SCIWirelessApplicationsReport10overviewDataTypeType)
	return m
}

// SCIWirelessApplicationsReport10overviewMetaData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_10_overview_MetaData
type SCIWirelessApplicationsReport10overviewMetaData struct {
	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessApplicationsReport10overviewMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWirelessApplicationsReport10overviewMetaData SCIWirelessApplicationsReport10overviewMetaData
	tmpType := new(_SCIWirelessApplicationsReport10overviewMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "percentage")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "traffic")
	*t = SCIWirelessApplicationsReport10overviewMetaData(*tmpType)
	return nil
}

func (t *SCIWirelessApplicationsReport10overviewMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
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

func NewSCIWirelessApplicationsReport10overviewMetaData() *SCIWirelessApplicationsReport10overviewMetaData {
	m := new(SCIWirelessApplicationsReport10overviewMetaData)
	return m
}

// SCIWirelessApplicationsReport11top10ApplicationsByClientCountData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_11_top10ApplicationsByClientCount_Data
type SCIWirelessApplicationsReport11top10ApplicationsByClientCountData [][]*SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType

func MakeSCIWirelessApplicationsReport11top10ApplicationsByClientCountData() SCIWirelessApplicationsReport11top10ApplicationsByClientCountData {
	m := make(SCIWirelessApplicationsReport11top10ApplicationsByClientCountData, 0)
	return m
}

// SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_11_top10ApplicationsByClientCount_DataType
type SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataType []*SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType

func MakeSCIWirelessApplicationsReport11top10ApplicationsByClientCountDataType() SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataType {
	m := make(SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataType, 0)
	return m
}

// SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_11_top10ApplicationsByClientCount_DataTypeType
type SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType struct {
	App *string `json:"app,omitempty"`

	Arp *float64 `json:"arp,omitempty"`

	Bacnetapp *float64 `json:"bacnet_app,omitempty"`

	Bjnp *float64 `json:"bjnp,omitempty"`

	Dhcp *float64 `json:"dhcp,omitempty"`

	Dns *float64 `json:"dns,omitempty"`

	End *string `json:"end,omitempty"`

	Lb01devenanet *float64 `json:"lb01.dev.ena.net,omitempty"`

	Lb02devenanet *float64 `json:"lb02.dev.ena.net,omitempty"`

	Nbns *float64 `json:"nbns,omitempty"`

	Rx *float64 `json:"rx,omitempty"`

	Snmp *float64 `json:"snmp,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Tx *float64 `json:"tx,omitempty"`

	UniqueUsers *float64 `json:"uniqueUsers,omitempty"`

	Unknown *float64 `json:"Unknown,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType
	tmpType := new(_SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "app")
	delete(tmpType.XAdditionalProperties, "arp")
	delete(tmpType.XAdditionalProperties, "bacnet_app")
	delete(tmpType.XAdditionalProperties, "bjnp")
	delete(tmpType.XAdditionalProperties, "dhcp")
	delete(tmpType.XAdditionalProperties, "dns")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "lb01.dev.ena.net")
	delete(tmpType.XAdditionalProperties, "lb02.dev.ena.net")
	delete(tmpType.XAdditionalProperties, "nbns")
	delete(tmpType.XAdditionalProperties, "rx")
	delete(tmpType.XAdditionalProperties, "snmp")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "tx")
	delete(tmpType.XAdditionalProperties, "uniqueUsers")
	delete(tmpType.XAdditionalProperties, "Unknown")
	*t = SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType(*tmpType)
	return nil
}

func (t *SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.Arp != nil {
		tmp["arp"] = t.Arp
	}
	if t.Bacnetapp != nil {
		tmp["bacnet_app"] = t.Bacnetapp
	}
	if t.Bjnp != nil {
		tmp["bjnp"] = t.Bjnp
	}
	if t.Dhcp != nil {
		tmp["dhcp"] = t.Dhcp
	}
	if t.Dns != nil {
		tmp["dns"] = t.Dns
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Lb01devenanet != nil {
		tmp["lb01.dev.ena.net"] = t.Lb01devenanet
	}
	if t.Lb02devenanet != nil {
		tmp["lb02.dev.ena.net"] = t.Lb02devenanet
	}
	if t.Nbns != nil {
		tmp["nbns"] = t.Nbns
	}
	if t.Rx != nil {
		tmp["rx"] = t.Rx
	}
	if t.Snmp != nil {
		tmp["snmp"] = t.Snmp
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.Tx != nil {
		tmp["tx"] = t.Tx
	}
	if t.UniqueUsers != nil {
		tmp["uniqueUsers"] = t.UniqueUsers
	}
	if t.Unknown != nil {
		tmp["Unknown"] = t.Unknown
	}
	return json.Marshal(tmp)
}

func NewSCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType() *SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType {
	m := new(SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType)
	return m
}

// SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_11_top10ApplicationsByClientCount_MetaData
type SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData
	tmpType := new(_SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	delete(tmpType.XAdditionalProperties, "totalClients")
	*t = SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData(*tmpType)
	return nil
}

func (t *SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData() *SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData {
	m := new(SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData)
	return m
}

// ReportWirelessApplicationsReport7Top10ApplicationsByTrafficVolume
//
// Operation ID: report_WirelessApplicationsReport_7_top10ApplicationsByTrafficVolume
//
// Wireless - Applications Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport7Top10ApplicationsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessApplicationsReport7Top10ApplicationsByTrafficVolume, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport8TopAppsByTrafficTable
//
// Operation ID: report_WirelessApplicationsReport_8_topAppsByTrafficTable
//
// Wireless - Applications Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport8TopAppsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessApplicationsReport8TopAppsByTrafficTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport9TopAppsByClientsTable
//
// Operation ID: report_WirelessApplicationsReport_9_topAppsByClientsTable
//
// Wireless - Applications Report - Top Applications by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport9TopAppsByClientsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessApplicationsReport9TopAppsByClientsTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport10Overview
//
// Operation ID: report_WirelessApplicationsReport_10_overview
//
// Wireless - Applications Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport10Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport10overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport10overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessApplicationsReport10Overview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport10overview200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport11Top10ApplicationsByClientCount
//
// Operation ID: report_WirelessApplicationsReport_11_top10ApplicationsByClientCount
//
// Wireless - Applications Report - Top Applications by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport11Top10ApplicationsByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWirelessApplicationsReport11Top10ApplicationsByClientCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
