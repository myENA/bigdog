package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
)

// WSGDeviceCapacityDevicesSummary
//
// Definition: deviceCapacity_devicesSummary
type WSGDeviceCapacityDevicesSummary struct {
	// ApCapacity
	// maximum ap capacity of this node.
	ApCapacity *int `json:"apCapacity,omitempty"`

	// Aps
	// connected aps in this node.
	Aps *int `json:"aps,omitempty"`

	// MaxApOfCluster
	// maximum ap capacity per cluster.
	MaxApOfCluster *int `json:"maxApOfCluster,omitempty"`

	// MaxSwitchOfCluster
	// maximum switch capacity per cluster.
	MaxSwitchOfCluster *int `json:"maxSwitchOfCluster,omitempty"`

	// SwitchCapacity
	// maximum switch capacity of this node.
	SwitchCapacity *int `json:"switchCapacity,omitempty"`

	// Switches
	// connected switches in this node.
	Switches *int `json:"switches,omitempty"`

	// TotalApCapacity
	// maximum total ap capacity of this node.
	TotalApCapacity *int `json:"totalApCapacity,omitempty"`

	// TotalAps
	// total connected switches in the cluster.
	TotalAps *int `json:"totalAps,omitempty"`

	// TotalRemainingApCapacity
	// total remaining ap capacity of this node.
	TotalRemainingApCapacity *int `json:"totalRemainingApCapacity,omitempty"`

	// TotalRemainingSwitchCapacity
	// total remaining switch capacity of this node.
	TotalRemainingSwitchCapacity *int `json:"totalRemainingSwitchCapacity,omitempty"`

	// TotalSwitchCapacity
	// maximum total switch capacity of this node.
	TotalSwitchCapacity *int `json:"totalSwitchCapacity,omitempty"`

	// TotalSwitches
	// total connected switches in the cluster.
	TotalSwitches *int `json:"totalSwitches,omitempty"`
}

type WSGDeviceCapacityDevicesSummaryAPIResponse struct {
	*RawAPIResponse
	Data *WSGDeviceCapacityDevicesSummary
}

func newWSGDeviceCapacityDevicesSummaryAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGDeviceCapacityDevicesSummaryAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGDeviceCapacityDevicesSummaryAPIResponse) Hydrate() error {
	r.Data = new(WSGDeviceCapacityDevicesSummary)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGDeviceCapacityDevicesSummary() *WSGDeviceCapacityDevicesSummary {
	m := new(WSGDeviceCapacityDevicesSummary)
	return m
}
