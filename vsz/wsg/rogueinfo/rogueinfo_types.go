package rogueinfo

// API Version: v8_0

type RogueInfo struct {
	Channel        *int             `json:"channel,omitempty"`
	Classification *string          `json:"classification,omitempty"`
	DetectedByAP   []*apinfo.ApInfo `json:"detectedByAP,omitempty"`
	Encryption     *string          `json:"encryption,omitempty"`
	LastDetected   *int             `json:"lastDetected,omitempty"`
	MatchResult    *string          `json:"matchResult,omitempty"`
	Radio          *string          `json:"radio,omitempty"`
	RogueAPMac     *string          `json:"rogueAPMac,omitempty"`
	RogueMac       *string          `json:"rogueMac,omitempty"`
	Ssid           *string          `json:"ssid,omitempty"`
	Type           *string          `json:"type,omitempty"`
}

type RogueInfoList struct {
	Extra             interface{}  `json:"extra,omitempty"`
	FirstIndex        *int         `json:"firstIndex,omitempty"`
	HasMore           *bool        `json:"hasMore,omitempty"`
	List              []*RogueInfo `json:"list,omitempty"`
	RawDataTotalCount *int         `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int         `json:"totalCount,omitempty"`
}
