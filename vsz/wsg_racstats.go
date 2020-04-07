package vsz

// API Version: v9_0

type WSGRacStatsGgsnGtp struct {
	// CbladeId
	// Control Blade ID
	// Constraints:
	//    - nullable
	CbladeId *string `json:"cbladeId,omitempty"`

	// CbladeName
	// Control Blade Name
	// Constraints:
	//    - nullable
	CbladeName *string `json:"cbladeName,omitempty"`

	// FailPdpCrt
	// PDP (Failed)
	// Constraints:
	//    - nullable
	FailPdpCrt *int `json:"failPdpCrt,omitempty"`

	// FailPdpDelInitAdm
	// Number of data admin initiated delete (Failed)
	// Constraints:
	//    - nullable
	FailPdpDelInitAdm *int `json:"failPdpDelInitAdm,omitempty"`

	// FailPdpDelInitAP
	// Number of AP initiated delete due to timeout at AP (Failed)
	// Constraints:
	//    - nullable
	FailPdpDelInitAP *int `json:"failPdpDelInitAP,omitempty"`

	// FailPdpDelInitClnt
	// Number of client initiated delete (Failed)
	// Constraints:
	//    - nullable
	FailPdpDelInitClnt *int `json:"failPdpDelInitClnt,omitempty"`

	// FailPdpDelInitD
	// Number of data plane initiated delete due to timeout at data plane (Failed)
	// Constraints:
	//    - nullable
	FailPdpDelInitD *int `json:"failPdpDelInitD,omitempty"`

	// FailPdpDelInitDM
	// Number of controller initiated delete due to DM from AAA (Failed)
	// Constraints:
	//    - nullable
	FailPdpDelInitDM *int `json:"failPdpDelInitDM,omitempty"`

	// FailPdpDelInitErr
	// Number of controller initiated delete due to critical error (Failed)
	// Constraints:
	//    - nullable
	FailPdpDelInitErr *int `json:"failPdpDelInitErr,omitempty"`

	// FailPdpDelInitSCG
	// Number of controller initiated delete due to timeout at controller (Failed)
	// Constraints:
	//    - nullable
	FailPdpDelInitSCG *int `json:"failPdpDelInitSCG,omitempty"`

	// FailPdpDelRcvd
	// Number of successful GGSN initiated delete (Failed)
	// Constraints:
	//    - nullable
	FailPdpDelRcvd *int `json:"failPdpDelRcvd,omitempty"`

	// FailPdpUpdInitAAA
	// Number of controller initiated update - CoA from AAA (Failed)
	// Constraints:
	//    - nullable
	FailPdpUpdInitAAA *int `json:"failPdpUpdInitAAA,omitempty"`

	// FailPdpUpdInitRM
	// PDP Update Initiated (Failed)
	// Constraints:
	//    - nullable
	FailPdpUpdInitRM *int `json:"failPdpUpdInitRM,omitempty"`

	// FailPdpUpdRcvd
	// PDP Update Received (Failed)
	// Constraints:
	//    - nullable
	FailPdpUpdRcvd *int `json:"failPdpUpdRcvd,omitempty"`

	// GgsnIp
	// GGSN IP
	// Constraints:
	//    - nullable
	GgsnIp *string `json:"ggsnIp,omitempty"`

	// MvnoId
	// MVNO ID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// MvnoName
	// MVNO Name
	// Constraints:
	//    - nullable
	MvnoName *string `json:"mvnoName,omitempty"`

	// NumActPdp
	// PDP (Active)
	// Constraints:
	//    - nullable
	NumActPdp *int `json:"numActPdp,omitempty"`

	// RecCreateTime
	// Time this entry was created
	// Constraints:
	//    - nullable
	RecCreateTime *int `json:"recCreateTime,omitempty"`

	// RecUpdateTime
	// Last time this entry was updated
	// Constraints:
	//    - nullable
	RecUpdateTime *int `json:"recUpdateTime,omitempty"`

	// SuccPdpCrt
	// PDP (Successful)
	// Constraints:
	//    - nullable
	SuccPdpCrt *int `json:"succPdpCrt,omitempty"`

	// SuccPdpDelInitAdm
	// Number of admin initiated delete (Successful)
	// Constraints:
	//    - nullable
	SuccPdpDelInitAdm *int `json:"succPdpDelInitAdm,omitempty"`

	// SuccPdpDelInitAP
	// Number of APinitiated delete due to timeout at AP (Successful)
	// Constraints:
	//    - nullable
	SuccPdpDelInitAP *int `json:"succPdpDelInitAP,omitempty"`

	// SuccPdpDelInitClnt
	// Number of clientinitiated delete (Successful)
	// Constraints:
	//    - nullable
	SuccPdpDelInitClnt *int `json:"succPdpDelInitClnt,omitempty"`

	// SuccPdpDelInitD
	// Number of data plane initiated delete due to timeout at data plane (Successful)
	// Constraints:
	//    - nullable
	SuccPdpDelInitD *int `json:"succPdpDelInitD,omitempty"`

	// SuccPdpDelInitDM
	// Number of controller initiated delete due to DM from AAA (Successful)
	// Constraints:
	//    - nullable
	SuccPdpDelInitDM *int `json:"succPdpDelInitDM,omitempty"`

	// SuccPdpDelInitErr
	// Number of controller initiated delete due to critical error (Successful)
	// Constraints:
	//    - nullable
	SuccPdpDelInitErr *int `json:"succPdpDelInitErr,omitempty"`

	// SuccPdpDelInitSCG
	// Number of controller initiated delete due to timeout at controller (Successful)
	// Constraints:
	//    - nullable
	SuccPdpDelInitSCG *int `json:"succPdpDelInitSCG,omitempty"`

	// SuccPdpDelRcvd
	// Number of successful GGSN initiated delete (Successful)
	// Constraints:
	//    - nullable
	SuccPdpDelRcvd *int `json:"succPdpDelRcvd,omitempty"`

	// SuccPdpUpdInitAAA
	// Number of controller initiated update - CoA from AAA (Successful)
	// Constraints:
	//    - nullable
	SuccPdpUpdInitAAA *int `json:"succPdpUpdInitAAA,omitempty"`

	// SuccPdpUpdInitRM
	// PDP Update Initiated (Successful)
	// Constraints:
	//    - nullable
	SuccPdpUpdInitRM *int `json:"succPdpUpdInitRM,omitempty"`

	// SuccPdpUpdRcvd
	// PDP Update Received (Successful)
	// Constraints:
	//    - nullable
	SuccPdpUpdRcvd *int `json:"succPdpUpdRcvd,omitempty"`
}

func NewWSGRacStatsGgsnGtp() *WSGRacStatsGgsnGtp {
	m := new(WSGRacStatsGgsnGtp)
	return m
}

type WSGRacStatsGgsnGtpcCon struct {
	// CbladeId
	// Control Blade ID
	// Constraints:
	//    - nullable
	CbladeId *string `json:"cbladeId,omitempty"`

	// CbladeName
	// Control Blade Name
	// Constraints:
	//    - nullable
	CbladeName *string `json:"cbladeName,omitempty"`

	// EchoReqRcvd
	// Echo Rsp Sent
	// Constraints:
	//    - nullable
	EchoReqRcvd *int `json:"echoReqRcvd,omitempty"`

	// EchoReqSent
	// Echo Req Sent
	// Constraints:
	//    - nullable
	EchoReqSent *string `json:"echoReqSent,omitempty"`

	// EchoRespRcvd
	// Echo Rsp Rcvd
	// Constraints:
	//    - nullable
	EchoRespRcvd *int `json:"echoRespRcvd,omitempty"`

	// EchoRespSent
	// Echo Req Sent
	// Constraints:
	//    - nullable
	EchoRespSent *int `json:"echoRespSent,omitempty"`

	// GgsnIp
	// IP of GGSN
	// Constraints:
	//    - nullable
	GgsnIp *string `json:"ggsnIp,omitempty"`

	// NumPathFailure
	// PathFailure
	// Constraints:
	//    - nullable
	NumPathFailure *int `json:"numPathFailure,omitempty"`

	// RecCrtTime
	// Time this entry was created
	// Constraints:
	//    - nullable
	RecCrtTime *int `json:"recCrtTime,omitempty"`

	// RecUpdTime
	// Last time this entry was updated
	// Constraints:
	//    - nullable
	RecUpdTime *int `json:"recUpdTime,omitempty"`
}

func NewWSGRacStatsGgsnGtpcCon() *WSGRacStatsGgsnGtpcCon {
	m := new(WSGRacStatsGgsnGtpcCon)
	return m
}

type WSGRacStatsGgsnGtpcConList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGRacStatsGgsnGtpcCon `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGRacStatsGgsnGtpcConList() *WSGRacStatsGgsnGtpcConList {
	m := new(WSGRacStatsGgsnGtpcConList)
	return m
}

type WSGRacStatsGgsnGtpList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGRacStatsGgsnGtp `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGRacStatsGgsnGtpList() *WSGRacStatsGgsnGtpList {
	m := new(WSGRacStatsGgsnGtpList)
	return m
}

type WSGRacStatsRadiusProxy struct {
	// AaASerIp
	// AAA IP
	// Constraints:
	//    - nullable
	AaASerIp *string `json:"aaASerIp,omitempty"`

	// AaaServiceName
	// AAA Service Name
	// Constraints:
	//    - nullable
	AaaServiceName *string `json:"aaaServiceName,omitempty"`

	// CbladeId
	// Control Balde ID
	// Constraints:
	//    - nullable
	CbladeId *string `json:"cbladeId,omitempty"`

	// CbladeName
	// Control Balde Name
	// Constraints:
	//    - nullable
	CbladeName *string `json:"cbladeName,omitempty"`

	// MvnoId
	// MVNO ID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// MvnoName
	// MVNO Name
	// Constraints:
	//    - nullable
	MvnoName *string `json:"mvnoName,omitempty"`

	// NasType
	// NAS Type. 1: Ruckus AP 2: 3rd Party AP
	// Constraints:
	//    - nullable
	NasType *int `json:"nasType,omitempty"`

	// NumAccRqRcvdNas
	// Number of RADIUS accounting requests received from NAS
	// Constraints:
	//    - nullable
	NumAccRqRcvdNas *int `json:"numAccRqRcvdNas,omitempty"`

	// NumAccRspSntNas
	// Number of RADIUS accounting responses sent to NAS
	// Constraints:
	//    - nullable
	NumAccRspSntNas *int `json:"numAccRspSntNas,omitempty"`

	// NumAcsAcpSntNas
	// Number of RADIUS access accepts sent to NAS
	// Constraints:
	//    - nullable
	NumAcsAcpSntNas *int `json:"numAcsAcpSntNas,omitempty"`

	// NumAcsChSntNas
	// Number of RADIUS access challenges sent to NAS
	// Constraints:
	//    - nullable
	NumAcsChSntNas *int `json:"numAcsChSntNas,omitempty"`

	// NumAcsRejSntNas
	// Number of RADIUS access rejects sent to the NAS
	// Constraints:
	//    - nullable
	NumAcsRejSntNas *int `json:"numAcsRejSntNas,omitempty"`

	// NumAcsRqRcvdNas
	// Number of RADIUS access requests received from NAS
	// Constraints:
	//    - nullable
	NumAcsRqRcvdNas *int `json:"numAcsRqRcvdNas,omitempty"`

	// NumAPAcctReq
	// Number of AP accounting requests
	// Constraints:
	//    - nullable
	NumAPAcctReq *int `json:"numAPAcctReq,omitempty"`

	// NumAPAcctRsp
	// Number of AP accounting responses
	// Constraints:
	//    - nullable
	NumAPAcctRsp *int `json:"numAPAcctRsp,omitempty"`

	// NumAuthOnlyAAA
	// Number of CoA authorize only processed
	// Constraints:
	//    - nullable
	NumAuthOnlyAAA *int `json:"numAuthOnlyAAA,omitempty"`

	// NumAuthOnlyFailedAAA
	// Number of failed CoA authorize only procedures
	// Constraints:
	//    - nullable
	NumAuthOnlyFailedAAA *int `json:"numAuthOnlyFailedAAA,omitempty"`

	// NumAuthOnlySuccAAA
	// Number of successful CoA authorize only procedures
	// Constraints:
	//    - nullable
	NumAuthOnlySuccAAA *int `json:"numAuthOnlySuccAAA,omitempty"`

	// NumCoaFailRcdNas
	// Number of Failed CoA response from NAS
	// Constraints:
	//    - nullable
	NumCoaFailRcdNas *int `json:"numCoaFailRcdNas,omitempty"`

	// NumCoaSntNas
	// Number of CoA sent to NAS
	// Constraints:
	//    - nullable
	NumCoaSntNas *int `json:"numCoaSntNas,omitempty"`

	// NumCoaSucRcdNas
	// Number of successful CoA responses from NAS
	// Constraints:
	//    - nullable
	NumCoaSucRcdNas *int `json:"numCoaSucRcdNas,omitempty"`

	// NumDmFailRcdNas
	// Number of RADIUS DM responses received from NAS (Failed)
	// Constraints:
	//    - nullable
	NumDmFailRcdNas *int `json:"numDmFailRcdNas,omitempty"`

	// NumDmSntNas
	// Number of RADIUS DM requests sent to NAS
	// Constraints:
	//    - nullable
	NumDmSntNas *int `json:"numDmSntNas,omitempty"`

	// NumDmSucRcdNas
	// Number of RADIUS DM responses received from NAS (Successful)
	// Constraints:
	//    - nullable
	NumDmSucRcdNas *int `json:"numDmSucRcdNas,omitempty"`

	// NumDroppedRateLimitAcct
	// Number of dropped accounting requests
	// Constraints:
	//    - nullable
	NumDroppedRateLimitAcct *int `json:"numDroppedRateLimitAcct,omitempty"`

	// NumDroppedRateLimitAuth
	// Number of dropped authentication requests
	// Constraints:
	//    - nullable
	NumDroppedRateLimitAuth *int `json:"numDroppedRateLimitAuth,omitempty"`

	// NumFailAcct
	// Number of accounting sessions established (Failed)
	// Constraints:
	//    - nullable
	NumFailAcct *int `json:"numFailAcct,omitempty"`

	// NumFailAuth
	// Number of authentications (Failed)
	// Constraints:
	//    - nullable
	NumFailAuth *int `json:"numFailAuth,omitempty"`

	// NumInCompAuth
	// Number of authentications (Incomplete)
	// Constraints:
	//    - nullable
	NumInCompAuth *int `json:"numInCompAuth,omitempty"`

	// NumOfAccAcceptAaa
	// Number of RADIUS access accepts received from AAA server
	// Constraints:
	//    - nullable
	NumOfAccAcceptAaa *int `json:"numOfAccAcceptAaa,omitempty"`

	// NumOfAccChallAaa
	// Number of RADIUS access challenges received from AAA server
	// Constraints:
	//    - nullable
	NumOfAccChallAaa *int `json:"numOfAccChallAaa,omitempty"`

	// NumOfAccRejAaa
	// Number of RADIUS access rejects received from AAA server
	// Constraints:
	//    - nullable
	NumOfAccRejAaa *int `json:"numOfAccRejAaa,omitempty"`

	// NumOfAccReqAaa
	// Number of RADIUS access requests sent to AAA server
	// Constraints:
	//    - nullable
	NumOfAccReqAaa *int `json:"numOfAccReqAaa,omitempty"`

	// NumOfAcctReqAaa
	// Number of RADIUS accounting requests sent to AAA server
	// Constraints:
	//    - nullable
	NumOfAcctReqAaa *int `json:"numOfAcctReqAaa,omitempty"`

	// NumOfAcctRspAaa
	// Number of RADIUS accounting responses received from AAA server
	// Constraints:
	//    - nullable
	NumOfAcctRspAaa *int `json:"numOfAcctRspAaa,omitempty"`

	// NumOfCoAFailAaa
	// Number of RADIUS CoA responses sent to AAA server (Failed)
	// Constraints:
	//    - nullable
	NumOfCoAFailAaa *int `json:"numOfCoAFailAaa,omitempty"`

	// NumOfCoAReqAaa
	// Number of RADIUS CoA responses sent to AAA server (Successful)
	// Constraints:
	//    - nullable
	NumOfCoAReqAaa *int `json:"numOfCoAReqAaa,omitempty"`

	// NumOfCoAResAaa
	// Number of RADIUS CoA requests received from AAA server
	// Constraints:
	//    - nullable
	NumOfCoAResAaa *int `json:"numOfCoAResAaa,omitempty"`

	// NumOfDmFailAaa
	// Number of RADIUS DM responses sent to AAA server (Failed)
	// Constraints:
	//    - nullable
	NumOfDmFailAaa *int `json:"numOfDmFailAaa,omitempty"`

	// NumOfDmReqAaa
	// Number of RADIUS DM responses sent to AAA server (Successful)
	// Constraints:
	//    - nullable
	NumOfDmReqAaa *int `json:"numOfDmReqAaa,omitempty"`

	// NumOfDmResAaa
	// Number of RADIUS DM requests received from AAA server
	// Constraints:
	//    - nullable
	NumOfDmResAaa *int `json:"numOfDmResAaa,omitempty"`

	// NumRadFailAPAcc
	// Number of AP accounting sessions established (Failed)
	// Constraints:
	//    - nullable
	NumRadFailAPAcc *int `json:"numRadFailAPAcc,omitempty"`

	// NumRadSuccAPAcc
	// Number of AP accounting sessions established (Successful)
	// Constraints:
	//    - nullable
	NumRadSuccAPAcc *int `json:"numRadSuccAPAcc,omitempty"`

	// NumSuccAcct
	// Number of accounting sessions established (Successful)
	// Constraints:
	//    - nullable
	NumSuccAcct *int `json:"numSuccAcct,omitempty"`

	// NumSuccAuth
	// Number of authentications (Successful)
	// Constraints:
	//    - nullable
	NumSuccAuth *int `json:"numSuccAuth,omitempty"`

	// RecCreateTime
	// Time this entry was created
	// Constraints:
	//    - nullable
	RecCreateTime *int `json:"recCreateTime,omitempty"`

	// RecUpdateTime
	// Last time this entry was updated
	// Constraints:
	//    - nullable
	RecUpdateTime *int `json:"recUpdateTime,omitempty"`
}

func NewWSGRacStatsRadiusProxy() *WSGRacStatsRadiusProxy {
	m := new(WSGRacStatsRadiusProxy)
	return m
}

type WSGRacStatsRadiusProxyList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGRacStatsRadiusProxy `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGRacStatsRadiusProxyList() *WSGRacStatsRadiusProxyList {
	m := new(WSGRacStatsRadiusProxyList)
	return m
}
