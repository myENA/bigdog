package vsz

// API Version: v9_0

type WSGToolSpeedFlex struct {
	// ClientIp
	// Constraints:
	//    - nullable
	ClientIp *WSGCommonIpAddress `json:"clientIp,omitempty"`

	// ClientMac
	// Constraints:
	//    - nullable
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

	// ServerIp
	// Constraints:
	//    - nullable
	ServerIp *WSGCommonIpAddress `json:"serverIp,omitempty"`

	// ServerMac
	// Constraints:
	//    - nullable
	ServerMac *WSGCommonMac `json:"serverMac,omitempty"`

	// Syspmtu
	// Default: 1500
	// Constraints:
	//    - nullable
	Syspmtu *int `json:"syspmtu,omitempty"`

	// Tool
	// SpeedFlex tool
	// Constraints:
	//    - required
	//    - oneof:[ZAP_DOWN,ZAP_UP]
	Tool *string `json:"tool" validate:"required,oneof=ZAP_DOWN ZAP_UP"`
}

func NewWSGToolSpeedFlex() *WSGToolSpeedFlex {
	m := new(WSGToolSpeedFlex)
	return m
}

type WSGToolTestResult struct {
	// Downlink
	// Downlink
	// Constraints:
	//    - nullable
	Downlink *int `json:"downlink,omitempty"`

	// Etf
	// ETF
	// Constraints:
	//    - nullable
	Etf *int `json:"etf,omitempty"`

	// Latency
	// Latency
	// Constraints:
	//    - nullable
	Latency *int `json:"latency,omitempty"`

	// PacketLoss
	// Packet loss
	// Constraints:
	//    - nullable
	PacketLoss *int `json:"packetLoss,omitempty"`

	// ResultId
	// Result ID
	// Constraints:
	//    - nullable
	ResultId *int `json:"resultId,omitempty"`

	// Uplink
	// Uplink
	// Constraints:
	//    - nullable
	Uplink *int `json:"uplink,omitempty"`

	// Wcid
	// Constraints:
	//    - nullable
	Wcid *string `json:"wcid,omitempty"`
}

func NewWSGToolTestResult() *WSGToolTestResult {
	m := new(WSGToolTestResult)
	return m
}
