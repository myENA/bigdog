package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
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

func newWSGAPRoutineStatusIntervalRspAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAPRoutineStatusIntervalRspAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
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
