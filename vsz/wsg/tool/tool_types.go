package tool

// API Version: v8_0

type SpeedFlex struct {
	ClientIP  *string `json:"clientIp,omitempty"`
	ClientMac *string `json:"clientMac,omitempty"`
	Model     *string `json:"model,omitempty"`
	Protocol  *string `json:"protocol,omitempty"`
	ServerIP  *string `json:"serverIp,omitempty"`
	ServerMac *string `json:"serverMac,omitempty"`
	Syspmtu   *int    `json:"syspmtu,omitempty"`
	Tool      *string `json:"tool,omitempty"`
}

type TestResult struct {
	Downlink   *int    `json:"downlink,omitempty"`
	Etf        *int    `json:"etf,omitempty"`
	Latency    *int    `json:"latency,omitempty"`
	PacketLoss *int    `json:"packetLoss,omitempty"`
	ResultID   *int    `json:"resultId,omitempty"`
	Uplink     *int    `json:"uplink,omitempty"`
	Wcid       *string `json:"wcid,omitempty"`
}
