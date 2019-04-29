package northbounddatastreaming

// API Version: v8_0

type CreateNorthboundDataStreamingProfile struct {
	Name       *string `json:"name,omitempty"`
	Password   *string `json:"password,omitempty"`
	ServerHost *string `json:"serverHost,omitempty"`
	ServerPort *string `json:"serverPort,omitempty"`
	SystemID   *string `json:"systemId,omitempty"`
	User       *string `json:"user,omitempty"`
}

type EmptyResult struct {
	NorthboundDataStreamingEmptyResult *string `json:"northboundDataStreaming_emptyResult,omitempty"`
}

type ModifyNorthboundDataStreamingEventCodes struct {
	NorthboundDataStreamingAcceptedEventCodes []int `json:"northboundDataStreamingAcceptedEventCodes,omitempty"`
}

type ModifyNorthboundDataStreamingProfile struct {
	Name       *string `json:"name,omitempty"`
	Password   *string `json:"password,omitempty"`
	ServerHost *string `json:"serverHost,omitempty"`
	ServerPort *string `json:"serverPort,omitempty"`
	SystemID   *string `json:"systemId,omitempty"`
	User       *string `json:"user,omitempty"`
}

type NorthboundDataStreamingEventCodes struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type NorthboundDataStreamingEventCodesListType struct {
	Code *int    `json:"code,omitempty"`
	Type *string `json:"type,omitempty"`
}

type NorthboundDataStreamingProfile struct {
	ID         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Password   *string `json:"password,omitempty"`
	ServerHost *string `json:"serverHost,omitempty"`
	ServerPort *string `json:"serverPort,omitempty"`
	SystemID   *string `json:"systemId,omitempty"`
	User       *string `json:"user,omitempty"`
}

type NorthboundDataStreamingProfileListExtraType struct {
	NorthboundDataStreamingEnabled *bool    `json:"northboundDataStreamingEnabled,omitempty"`
	StreamingByDomainZoneEnabled   *bool    `json:"streamingByDomainZoneEnabled,omitempty"`
	StreamingDomainIds             []string `json:"streamingDomainIds,omitempty"`
	StreamingZoneIds               []string `json:"streamingZoneIds,omitempty"`
}

type NorthboundDataStreamingSettings struct {
	NorthboundDataStreamingEnabled *bool    `json:"northboundDataStreamingEnabled,omitempty"`
	StreamingByDomainZoneEnabled   *bool    `json:"streamingByDomainZoneEnabled,omitempty"`
	StreamingDomainIds             []string `json:"streamingDomainIds,omitempty"`
	StreamingZoneIds               []string `json:"streamingZoneIds,omitempty"`
}
