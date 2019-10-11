package aproutineconfiginterval

// API Version: v8_1

type ApRoutineConfigIntervalReq struct {
	ApRoutineConfigInterval *int `json:"apRoutineConfigInterval,omitempty"`
}

func NewApRoutineConfigIntervalReq() *ApRoutineConfigIntervalReq {
	apRoutineConfigIntervalReqType := new(ApRoutineConfigIntervalReq)
	return apRoutineConfigIntervalReqType
}

func NewApRoutineConfigIntervalReqWithDefaults() *ApRoutineConfigIntervalReq {
	apRoutineConfigIntervalReqType := new(ApRoutineConfigIntervalReq)
	return apRoutineConfigIntervalReqType
}

type ApRoutineConfigIntervalRsp struct {
	ApRoutineConfigInterval *int `json:"apRoutineConfigInterval,omitempty"`

	LowerBound *int `json:"lowerBound,omitempty"`

	UpperBound *int `json:"upperBound,omitempty"`
}

func NewApRoutineConfigIntervalRsp() *ApRoutineConfigIntervalRsp {
	apRoutineConfigIntervalRspType := new(ApRoutineConfigIntervalRsp)
	return apRoutineConfigIntervalRspType
}

func NewApRoutineConfigIntervalRspWithDefaults() *ApRoutineConfigIntervalRsp {
	apRoutineConfigIntervalRspType := new(ApRoutineConfigIntervalRsp)
	return apRoutineConfigIntervalRspType
}
