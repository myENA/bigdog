package bigdog

// API Version: v9_1

import (
	"errors"
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

func newWSGAPRoutineConfigIntervalRspAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPRoutineConfigIntervalRspAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPRoutineConfigIntervalRspAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAPRoutineConfigIntervalRsp)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAPRoutineConfigIntervalRsp() *WSGAPRoutineConfigIntervalRsp {
	m := new(WSGAPRoutineConfigIntervalRsp)
	return m
}
