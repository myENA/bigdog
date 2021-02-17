package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGClientQueryList
//
// Definition: clientQuery_clientQueryList
type WSGClientQueryList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGClientQueryCreateClientQuery `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGClientQueryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGClientQueryList
}

func newWSGClientQueryListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGClientQueryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGClientQueryListAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(WSGClientQueryList)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewWSGClientQueryList() *WSGClientQueryList {
	m := new(WSGClientQueryList)
	return m
}

// WSGClientQueryCreateClientQuery
//
// Definition: clientQuery_createClientQuery
type WSGClientQueryCreateClientQuery struct {
	Alerts *int `json:"alerts,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	AuthMethod *string `json:"authMethod,omitempty"`

	AuthStatus *string `json:"authStatus,omitempty"`

	Bssid *string `json:"bssid,omitempty"`

	Channel *int `json:"channel,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	ControlPlaneName *string `json:"controlPlaneName,omitempty"`

	CpeMac *string `json:"cpeMac,omitempty"`

	DataPlaneName *string `json:"dataPlaneName,omitempty"`

	DeviceType *string `json:"deviceType,omitempty"`

	Downlink *int `json:"downlink,omitempty"`

	DownlinkRate *int `json:"downlinkRate,omitempty"`

	EncryptionMethod *string `json:"encryptionMethod,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	IpAddress *string `json:"ipAddress,omitempty"`

	Ipv6Address *string `json:"ipv6Address,omitempty"`

	MedianRxMCSRate *int `json:"medianRxMCSRate,omitempty"`

	MedianTxMCSRate *int `json:"medianTxMCSRate,omitempty"`

	ModelName *string `json:"modelName,omitempty"`

	OsType *string `json:"osType,omitempty"`

	OsVendorType *string `json:"osVendorType,omitempty"`

	RadioType *string `json:"radioType,omitempty"`

	Role *string `json:"role,omitempty"`

	Rssi *int `json:"rssi,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	RxFrames *int `json:"rxFrames,omitempty"`

	SessionStartTime *int `json:"sessionStartTime,omitempty"`

	Snr *int `json:"snr,omitempty"`

	Speedflex *int `json:"speedflex,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Status *string `json:"status,omitempty"`

	TcWithQuotaList []*WSGClientQueryTcWithQuota `json:"tcWithQuotaList,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`

	TxDropDataFrames *int `json:"txDropDataFrames,omitempty"`

	TxFrames *int `json:"txFrames,omitempty"`

	TxRatebps *float64 `json:"txRatebps,omitempty"`

	TxRxBytes *int `json:"txRxBytes,omitempty"`

	Uplink *int `json:"uplink,omitempty"`

	UplinkRate *int `json:"uplinkRate,omitempty"`

	UserName *string `json:"userName,omitempty"`

	Vlan *int `json:"vlan,omitempty"`

	WlanType *string `json:"wlanType,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`

	ZoneVersion *string `json:"zoneVersion,omitempty"`
}

func NewWSGClientQueryCreateClientQuery() *WSGClientQueryCreateClientQuery {
	m := new(WSGClientQueryCreateClientQuery)
	return m
}

// WSGClientQueryQueryCriteria
//
// Definition: clientQuery_queryCriteria
type WSGClientQueryQueryCriteria struct {
	// Attributes
	// Get specific columns only
	Attributes []string `json:"attributes,omitempty"`

	// Criteria
	// Add backward compatibility for UI framework
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	ExtraFilters []*WSGCommonQueryCriteriaSuperSetExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*WSGCommonQueryCriteriaSuperSetExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *WSGCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*WSGCommonQueryCriteriaSuperSetFiltersType `json:"filters,omitempty"`

	FullTextSearch *WSGCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required information
	Options *WSGCommonQueryCriteriaSuperSetOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - min:1
	Page *int `json:"page,omitempty"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *WSGCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewWSGClientQueryQueryCriteria() *WSGClientQueryQueryCriteria {
	m := new(WSGClientQueryQueryCriteria)
	return m
}

// WSGClientQueryTcWithQuota
//
// Definition: clientQuery_tcWithQuota
type WSGClientQueryTcWithQuota struct {
	TcMaxQuota *string `json:"tcMaxQuota,omitempty"`

	TcName *string `json:"tcName,omitempty"`

	TcRemainingQuota *string `json:"tcRemainingQuota,omitempty"`
}

func NewWSGClientQueryTcWithQuota() *WSGClientQueryTcWithQuota {
	m := new(WSGClientQueryTcWithQuota)
	return m
}
