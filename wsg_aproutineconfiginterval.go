package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"io"
)

// WSGAPRoutineConfigIntervalReq
//
// Definition: apRoutineConfigInterval_apRoutineConfigIntervalReq
type WSGAPRoutineConfigIntervalReq struct {
	ApRoutineConfigInterval *int `json:"apRoutineConfigInterval,omitempty"`
}

func NewWSGAPRoutineConfigIntervalReq() *WSGAPRoutineConfigIntervalReq {
	m := new(WSGAPRoutineConfigIntervalReq)
	return m
}

// WSGAPRoutineConfigIntervalRsp
//
// Definition: apRoutineConfigInterval_apRoutineConfigIntervalRsp
type WSGAPRoutineConfigIntervalRsp struct {
	ApRoutineConfigInterval *int `json:"apRoutineConfigInterval,omitempty"`

	LowerBound *int `json:"lowerBound,omitempty"`

	UpperBound *int `json:"upperBound,omitempty"`
}

type WSGAPRoutineConfigIntervalRspAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPRoutineConfigIntervalRsp
}

func newWSGAPRoutineConfigIntervalRspAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPRoutineConfigIntervalRspAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPRoutineConfigIntervalRspAPIResponse) Hydrate() error {
	r.Data = new(WSGAPRoutineConfigIntervalRsp)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPRoutineConfigIntervalRsp() *WSGAPRoutineConfigIntervalRsp {
	m := new(WSGAPRoutineConfigIntervalRsp)
	return m
}
