package aproutinestatusinterval

// API Version: v8_1

type ApRoutineStatusIntervalRsp struct {
	ApRoutineStatusInterval *int `json:"apRoutineStatusInterval,omitempty"`
}

func NewApRoutineStatusIntervalRsp() *ApRoutineStatusIntervalRsp {
	apRoutineStatusIntervalRspType := new(ApRoutineStatusIntervalRsp)
	return apRoutineStatusIntervalRspType
}

func NewDefaultApRoutineStatusIntervalRsp() *ApRoutineStatusIntervalRsp {
	apRoutineStatusIntervalRspType := new(ApRoutineStatusIntervalRsp)
	return apRoutineStatusIntervalRspType
}
