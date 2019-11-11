package vsz

// API Version: v8_1

type WSGToolSpeedFlex struct {
	ClientIp *WSGCommonIpAddress `json:"clientIp,omitempty"`

	ClientMac *WSGCommonMac `json:"clientMac,omitempty"`

	// Model
	// Test model
	// Constraints:
	//    - nullable
	//    - oneof:[AP,CLIENT,TRACE,HOP,NULL]
	Model *string `json:"model,omitempty" validate:"omitempty,oneof=AP CLIENT TRACE HOP NULL"`

	// Protocol
	// Protocol used in the SpeedFlex test
	// Constraints:
	//    - required
	//    - oneof:[UDP,TCP]
	Protocol *string `json:"protocol" validate:"required,oneof=UDP TCP"`

	ServerIp *WSGCommonIpAddress `json:"serverIp,omitempty"`

	ServerMac *WSGCommonMac `json:"serverMac,omitempty"`

	// Syspmtu
	// Default: 1500
	Syspmtu *int `json:"syspmtu,omitempty"`

	// Tool
	// SpeedFlex tool
	// Constraints:
	//    - required
	//    - oneof:[ZAP_DOWN,ZAP_UP]
	Tool *string `json:"tool" validate:"required,oneof=ZAP_DOWN ZAP_UP"`
}

type WSGToolTestResult struct {
	// Downlink
	// Downlink
	Downlink *int `json:"downlink,omitempty"`

	// Etf
	// ETF
	Etf *int `json:"etf,omitempty"`

	// Latency
	// Latency
	Latency *int `json:"latency,omitempty"`

	// PacketLoss
	// Packet loss
	PacketLoss *int `json:"packetLoss,omitempty"`

	// ResultId
	// Result ID
	ResultId *int `json:"resultId,omitempty"`

	// Uplink
	// Uplink
	Uplink *int `json:"uplink,omitempty"`

	Wcid *string `json:"wcid,omitempty"`
}
