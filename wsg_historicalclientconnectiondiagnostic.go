package bigdog

// API Version: v9_1

// WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry
//
// Definition: historicalClientConnectionDiagnostic_clientConnectionFailureTypeCountEntry
type WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry struct {
	DisplayName *string `json:"displayName,omitempty"`

	Key *string `json:"key,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func NewWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry() *WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry {
	m := new(WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry)
	return m
}
