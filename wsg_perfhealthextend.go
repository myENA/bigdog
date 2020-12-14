package bigdog

// API Version: v9_1

// WSGPerformanceAndHealthExtensionsGroupBarEntry
//
// Definition: perfHealthExtend_groupBarEntry
type WSGPerformanceAndHealthExtensionsGroupBarEntry struct {
	Id interface{} `json:"id,omitempty"`

	Key *string `json:"key,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func NewWSGPerformanceAndHealthExtensionsGroupBarEntry() *WSGPerformanceAndHealthExtensionsGroupBarEntry {
	m := new(WSGPerformanceAndHealthExtensionsGroupBarEntry)
	return m
}

// WSGPerformanceAndHealthExtensionsGroupBarList
//
// Definition: perfHealthExtend_groupBarList
type WSGPerformanceAndHealthExtensionsGroupBarList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGPerformanceAndHealthExtensionsGroupBarEntry `json:"list,omitempty"`

	// RawDataTotalCount
	// Not sure what this is for, seems to always be 0
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGPerformanceAndHealthExtensionsGroupBarList() *WSGPerformanceAndHealthExtensionsGroupBarList {
	m := new(WSGPerformanceAndHealthExtensionsGroupBarList)
	return m
}

// WSGPerformanceAndHealthExtensionsLineAvgEntry
//
// Definition: perfHealthExtend_lineAvgEntry
type WSGPerformanceAndHealthExtensionsLineAvgEntry struct {
	// Key
	// Timestamp of entry
	Key *int `json:"key,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func NewWSGPerformanceAndHealthExtensionsLineAvgEntry() *WSGPerformanceAndHealthExtensionsLineAvgEntry {
	m := new(WSGPerformanceAndHealthExtensionsLineAvgEntry)
	return m
}

// WSGPerformanceAndHealthExtensionsLineList
//
// Definition: perfHealthExtend_lineList
type WSGPerformanceAndHealthExtensionsLineList struct {
	Avg []*WSGPerformanceAndHealthExtensionsLineAvgEntry `json:"avg,omitempty"`

	Id interface{} `json:"id,omitempty"`

	LikelyRange []*WSGPerformanceAndHealthExtensionsLineRangeEntry `json:"likelyRange,omitempty"`

	MaxMinRange []*WSGPerformanceAndHealthExtensionsLineRangeEntry `json:"maxMinRange,omitempty"`

	Median []*WSGPerformanceAndHealthExtensionsLineAvgEntry `json:"median,omitempty"`
}

func NewWSGPerformanceAndHealthExtensionsLineList() *WSGPerformanceAndHealthExtensionsLineList {
	m := new(WSGPerformanceAndHealthExtensionsLineList)
	return m
}

// WSGPerformanceAndHealthExtensionsLineRangeEntry
//
// Definition: perfHealthExtend_lineRangeEntry
type WSGPerformanceAndHealthExtensionsLineRangeEntry struct {
	High *float64 `json:"high,omitempty"`

	// Key
	// Timestamp of entry
	Key *int `json:"key,omitempty"`

	Low *float64 `json:"low,omitempty"`
}

func NewWSGPerformanceAndHealthExtensionsLineRangeEntry() *WSGPerformanceAndHealthExtensionsLineRangeEntry {
	m := new(WSGPerformanceAndHealthExtensionsLineRangeEntry)
	return m
}
