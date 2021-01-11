package bigdog

// API Version: v9_1

import (
	"errors"
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

func (r *WSGAPRoutineStatusIntervalRspAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAPRoutineStatusIntervalRsp)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAPRoutineStatusIntervalRsp() *WSGAPRoutineStatusIntervalRsp {
	m := new(WSGAPRoutineStatusIntervalRsp)
	return m
}
