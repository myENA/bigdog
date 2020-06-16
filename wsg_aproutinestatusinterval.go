package bigdog

// API Version: v9_0

// WSGAPRoutineStatusIntervalRsp
//
// Definition: apRoutineStatusInterval_apRoutineStatusIntervalRsp
type WSGAPRoutineStatusIntervalRsp struct {
	ApRoutineStatusInterval *int `json:"apRoutineStatusInterval,omitempty"`
}

func NewWSGAPRoutineStatusIntervalRsp() *WSGAPRoutineStatusIntervalRsp {
	m := new(WSGAPRoutineStatusIntervalRsp)
	return m
}
