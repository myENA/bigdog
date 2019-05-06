package tool

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type SpeedFlex struct {
	ClientIp *common.IpAddress `json:"clientIp,omitempty"`

	ClientMac *common.Mac `json:"clientMac,omitempty"`

	// Model
	// Test model
	Model *string `json:"model,omitempty"`

	// Protocol
	// Protocol used in the SpeedFlex test
	Protocol *string `json:"protocol,omitempty"`

	ServerIp *common.IpAddress `json:"serverIp,omitempty"`

	ServerMac *common.Mac `json:"serverMac,omitempty"`

	// Syspmtu
	// Default: 1500
	Syspmtu *int `json:"syspmtu,omitempty"`

	// Tool
	// SpeedFlex tool
	Tool *string `json:"tool,omitempty"`
}

type TestResult struct {
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
