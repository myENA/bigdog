package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"io"
)

// WSGAPRoutineStatusIntervalRsp
//
// Definition: apRoutineStatusInterval_apRoutineStatusIntervalRsp
type WSGAPRoutineStatusIntervalRsp struct {
	ApRoutineStatusInterval *int `json:"apRoutineStatusInterval,omitempty"`
}

type WSGAPRoutineStatusIntervalRspAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPRoutineStatusIntervalRsp
}

func newWSGAPRoutineStatusIntervalRspAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPRoutineStatusIntervalRspAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPRoutineStatusIntervalRspAPIResponse) Hydrate() error {
	r.Data = new(WSGAPRoutineStatusIntervalRsp)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPRoutineStatusIntervalRsp() *WSGAPRoutineStatusIntervalRsp {
	m := new(WSGAPRoutineStatusIntervalRsp)
	return m
}
