package northbounddatastreaming

// API Version: v8_0

type CreateNorthboundDataStreamingProfile struct {
	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Name *string `json:"name,omitempty"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Password *string `json:"password,omitempty"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	ServerHost *string `json:"serverHost,omitempty"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	ServerPort *string `json:"serverPort,omitempty"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	SystemId *string `json:"systemId,omitempty"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	User *string `json:"user,omitempty"`
}

type EmptyResult struct {
	NorthboundDataStreamingEmptyResult *string `json:"northboundDataStreaming_emptyResult,omitempty"`
}

type ModifyNorthboundDataStreamingEventCodes struct {
	NorthboundDataStreamingAcceptedEventCodes []int `json:"northboundDataStreamingAcceptedEventCodes,omitempty"`
}

type ModifyNorthboundDataStreamingProfile struct {
	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Name *string `json:"name,omitempty"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Password *string `json:"password,omitempty"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	ServerHost *string `json:"serverHost,omitempty"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	ServerPort *string `json:"serverPort,omitempty"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	SystemId *string `json:"systemId,omitempty"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	User *string `json:"user,omitempty"`
}

type NorthboundDataStreamingEventCodes struct {
	// FirstIndex
	// Index of the first event code returned from the complete event code set
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more Northbound Data Streaming accepted event codes after the currently
	// displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*NorthboundDataStreamingEventCodesListType `json:"list,omitempty"`

	// TotalCount
	// Total Northbound Data Streaming accepted event code count
	TotalCount *int `json:"totalCount,omitempty"`
}

type NorthboundDataStreamingEventCodesListType struct {
	// Code
	// Northbound Data Streaming accepted event code
	Code *int `json:"code,omitempty"`

	// Type
	// Northbound Data Streaming accepted event type
	Type *string `json:"type,omitempty"`
}

type NorthboundDataStreamingProfile struct {
	// Id
	// UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Id *string `json:"id,omitempty"`

	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Name *string `json:"name,omitempty"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Password *string `json:"password,omitempty"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	ServerHost *string `json:"serverHost,omitempty"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	ServerPort *string `json:"serverPort,omitempty"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	SystemId *string `json:"systemId,omitempty"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	User *string `json:"user,omitempty"`
}

type NorthboundDataStreamingProfileListExtraType struct {
	// NorthboundDataStreamingEnabled
	// Is Northbound Data Streaming enabled or disabled
	NorthboundDataStreamingEnabled *bool `json:"northboundDataStreamingEnabled,omitempty"`

	// StreamingByDomainZoneEnabled
	// Is Northbound Data Streaming enabled by domain/zone settings
	StreamingByDomainZoneEnabled *bool `json:"streamingByDomainZoneEnabled,omitempty"`

	// StreamingDomainIds
	// Domain Ids for 'streamingByDomainZoneEnabled' settings
	StreamingDomainIds []string `json:"streamingDomainIds,omitempty"`

	// StreamingZoneIds
	// Zone Ids for 'streamingByDomainZoneEnabled' settings
	StreamingZoneIds []string `json:"streamingZoneIds,omitempty"`
}

type NorthboundDataStreamingSettings struct {
	// NorthboundDataStreamingEnabled
	// Is Northbound Data Streaming enabled or disabled
	NorthboundDataStreamingEnabled *bool `json:"northboundDataStreamingEnabled,omitempty"`

	// StreamingByDomainZoneEnabled
	// Is Northbound Data Streaming enabled by domain/zone settings
	StreamingByDomainZoneEnabled *bool `json:"streamingByDomainZoneEnabled,omitempty"`

	// StreamingDomainIds
	// Domain Ids for 'streamingByDomainZoneEnabled' settings
	StreamingDomainIds []string `json:"streamingDomainIds,omitempty"`

	// StreamingZoneIds
	// Zone Ids for 'streamingByDomainZoneEnabled' settings
	StreamingZoneIds []string `json:"streamingZoneIds,omitempty"`
}
