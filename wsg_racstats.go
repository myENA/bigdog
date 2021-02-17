package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGRACStatsGgsnGtp
//
// Definition: racStats_ggsnGtp
type WSGRACStatsGgsnGtp struct {
	// CbladeId
	// Control Blade ID
	CbladeId *string `json:"cbladeId,omitempty"`

	// CbladeName
	// Control Blade Name
	CbladeName *string `json:"cbladeName,omitempty"`

	// FailPdpCrt
	// PDP (Failed)
	FailPdpCrt *int `json:"failPdpCrt,omitempty"`

	// FailPdpDelInitAdm
	// Number of data admin initiated delete (Failed)
	FailPdpDelInitAdm *int `json:"failPdpDelInitAdm,omitempty"`

	// FailPdpDelInitAP
	// Number of AP initiated delete due to timeout at AP (Failed)
	FailPdpDelInitAP *int `json:"failPdpDelInitAP,omitempty"`

	// FailPdpDelInitClnt
	// Number of client initiated delete (Failed)
	FailPdpDelInitClnt *int `json:"failPdpDelInitClnt,omitempty"`

	// FailPdpDelInitD
	// Number of data plane initiated delete due to timeout at data plane (Failed)
	FailPdpDelInitD *int `json:"failPdpDelInitD,omitempty"`

	// FailPdpDelInitDM
	// Number of controller initiated delete due to DM from AAA (Failed)
	FailPdpDelInitDM *int `json:"failPdpDelInitDM,omitempty"`

	// FailPdpDelInitErr
	// Number of controller initiated delete due to critical error (Failed)
	FailPdpDelInitErr *int `json:"failPdpDelInitErr,omitempty"`

	// FailPdpDelInitSCG
	// Number of controller initiated delete due to timeout at controller (Failed)
	FailPdpDelInitSCG *int `json:"failPdpDelInitSCG,omitempty"`

	// FailPdpDelRcvd
	// Number of successful GGSN initiated delete (Failed)
	FailPdpDelRcvd *int `json:"failPdpDelRcvd,omitempty"`

	// FailPdpUpdInitAAA
	// Number of controller initiated update - CoA from AAA (Failed)
	FailPdpUpdInitAAA *int `json:"failPdpUpdInitAAA,omitempty"`

	// FailPdpUpdInitRM
	// PDP Update Initiated (Failed)
	FailPdpUpdInitRM *int `json:"failPdpUpdInitRM,omitempty"`

	// FailPdpUpdRcvd
	// PDP Update Received (Failed)
	FailPdpUpdRcvd *int `json:"failPdpUpdRcvd,omitempty"`

	// GgsnIp
	// GGSN IP
	GgsnIp *string `json:"ggsnIp,omitempty"`

	// MvnoId
	// MVNO ID
	MvnoId *string `json:"mvnoId,omitempty"`

	// MvnoName
	// MVNO Name
	MvnoName *string `json:"mvnoName,omitempty"`

	// NumActPdp
	// PDP (Active)
	NumActPdp *int `json:"numActPdp,omitempty"`

	// RecCreateTime
	// Time this entry was created
	RecCreateTime *int `json:"recCreateTime,omitempty"`

	// RecUpdateTime
	// Last time this entry was updated
	RecUpdateTime *int `json:"recUpdateTime,omitempty"`

	// SuccPdpCrt
	// PDP (Successful)
	SuccPdpCrt *int `json:"succPdpCrt,omitempty"`

	// SuccPdpDelInitAdm
	// Number of admin initiated delete (Successful)
	SuccPdpDelInitAdm *int `json:"succPdpDelInitAdm,omitempty"`

	// SuccPdpDelInitAP
	// Number of APinitiated delete due to timeout at AP (Successful)
	SuccPdpDelInitAP *int `json:"succPdpDelInitAP,omitempty"`

	// SuccPdpDelInitClnt
	// Number of clientinitiated delete (Successful)
	SuccPdpDelInitClnt *int `json:"succPdpDelInitClnt,omitempty"`

	// SuccPdpDelInitD
	// Number of data plane initiated delete due to timeout at data plane (Successful)
	SuccPdpDelInitD *int `json:"succPdpDelInitD,omitempty"`

	// SuccPdpDelInitDM
	// Number of controller initiated delete due to DM from AAA (Successful)
	SuccPdpDelInitDM *int `json:"succPdpDelInitDM,omitempty"`

	// SuccPdpDelInitErr
	// Number of controller initiated delete due to critical error (Successful)
	SuccPdpDelInitErr *int `json:"succPdpDelInitErr,omitempty"`

	// SuccPdpDelInitSCG
	// Number of controller initiated delete due to timeout at controller (Successful)
	SuccPdpDelInitSCG *int `json:"succPdpDelInitSCG,omitempty"`

	// SuccPdpDelRcvd
	// Number of successful GGSN initiated delete (Successful)
	SuccPdpDelRcvd *int `json:"succPdpDelRcvd,omitempty"`

	// SuccPdpUpdInitAAA
	// Number of controller initiated update - CoA from AAA (Successful)
	SuccPdpUpdInitAAA *int `json:"succPdpUpdInitAAA,omitempty"`

	// SuccPdpUpdInitRM
	// PDP Update Initiated (Successful)
	SuccPdpUpdInitRM *int `json:"succPdpUpdInitRM,omitempty"`

	// SuccPdpUpdRcvd
	// PDP Update Received (Successful)
	SuccPdpUpdRcvd *int `json:"succPdpUpdRcvd,omitempty"`
}

func NewWSGRACStatsGgsnGtp() *WSGRACStatsGgsnGtp {
	m := new(WSGRACStatsGgsnGtp)
	return m
}

// WSGRACStatsGgsnGtpcCon
//
// Definition: racStats_ggsnGtpcCon
type WSGRACStatsGgsnGtpcCon struct {
	// CbladeId
	// Control Blade ID
	CbladeId *string `json:"cbladeId,omitempty"`

	// CbladeName
	// Control Blade Name
	CbladeName *string `json:"cbladeName,omitempty"`

	// EchoReqRcvd
	// Echo Rsp Sent
	EchoReqRcvd *int `json:"echoReqRcvd,omitempty"`

	// EchoReqSent
	// Echo Req Sent
	EchoReqSent *string `json:"echoReqSent,omitempty"`

	// EchoRespRcvd
	// Echo Rsp Rcvd
	EchoRespRcvd *int `json:"echoRespRcvd,omitempty"`

	// EchoRespSent
	// Echo Req Sent
	EchoRespSent *int `json:"echoRespSent,omitempty"`

	// GgsnIp
	// IP of GGSN
	GgsnIp *string `json:"ggsnIp,omitempty"`

	// NumPathFailure
	// PathFailure
	NumPathFailure *int `json:"numPathFailure,omitempty"`

	// RecCrtTime
	// Time this entry was created
	RecCrtTime *int `json:"recCrtTime,omitempty"`

	// RecUpdTime
	// Last time this entry was updated
	RecUpdTime *int `json:"recUpdTime,omitempty"`
}

func NewWSGRACStatsGgsnGtpcCon() *WSGRACStatsGgsnGtpcCon {
	m := new(WSGRACStatsGgsnGtpcCon)
	return m
}

// WSGRACStatsGgsnGtpcConList
//
// Definition: racStats_ggsnGtpcConList
type WSGRACStatsGgsnGtpcConList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGRACStatsGgsnGtpcCon `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGRACStatsGgsnGtpcConListAPIResponse struct {
	*RawAPIResponse
	Data *WSGRACStatsGgsnGtpcConList
}

func newWSGRACStatsGgsnGtpcConListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGRACStatsGgsnGtpcConListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGRACStatsGgsnGtpcConListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGRACStatsGgsnGtpcConList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGRACStatsGgsnGtpcConList() *WSGRACStatsGgsnGtpcConList {
	m := new(WSGRACStatsGgsnGtpcConList)
	return m
}

// WSGRACStatsGgsnGtpList
//
// Definition: racStats_ggsnGtpList
type WSGRACStatsGgsnGtpList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGRACStatsGgsnGtp `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGRACStatsGgsnGtpListAPIResponse struct {
	*RawAPIResponse
	Data *WSGRACStatsGgsnGtpList
}

func newWSGRACStatsGgsnGtpListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGRACStatsGgsnGtpListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGRACStatsGgsnGtpListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGRACStatsGgsnGtpList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGRACStatsGgsnGtpList() *WSGRACStatsGgsnGtpList {
	m := new(WSGRACStatsGgsnGtpList)
	return m
}

// WSGRACStatsRadiusProxy
//
// Definition: racStats_radiusProxy
type WSGRACStatsRadiusProxy struct {
	// AaASerIp
	// AAA IP
	AaASerIp *string `json:"aaASerIp,omitempty"`

	// AaaServiceName
	// AAA Service Name
	AaaServiceName *string `json:"aaaServiceName,omitempty"`

	// CbladeId
	// Control Balde ID
	CbladeId *string `json:"cbladeId,omitempty"`

	// CbladeName
	// Control Balde Name
	CbladeName *string `json:"cbladeName,omitempty"`

	// MvnoId
	// MVNO ID
	MvnoId *string `json:"mvnoId,omitempty"`

	// MvnoName
	// MVNO Name
	MvnoName *string `json:"mvnoName,omitempty"`

	// NasType
	// NAS Type. 1: Ruckus AP 2: 3rd Party AP
	NasType *int `json:"nasType,omitempty"`

	// NumAccRqRcvdNas
	// Number of RADIUS accounting requests received from NAS
	NumAccRqRcvdNas *int `json:"numAccRqRcvdNas,omitempty"`

	// NumAccRspSntNas
	// Number of RADIUS accounting responses sent to NAS
	NumAccRspSntNas *int `json:"numAccRspSntNas,omitempty"`

	// NumAcsAcpSntNas
	// Number of RADIUS access accepts sent to NAS
	NumAcsAcpSntNas *int `json:"numAcsAcpSntNas,omitempty"`

	// NumAcsChSntNas
	// Number of RADIUS access challenges sent to NAS
	NumAcsChSntNas *int `json:"numAcsChSntNas,omitempty"`

	// NumAcsRejSntNas
	// Number of RADIUS access rejects sent to the NAS
	NumAcsRejSntNas *int `json:"numAcsRejSntNas,omitempty"`

	// NumAcsRqRcvdNas
	// Number of RADIUS access requests received from NAS
	NumAcsRqRcvdNas *int `json:"numAcsRqRcvdNas,omitempty"`

	// NumAPAcctReq
	// Number of AP accounting requests
	NumAPAcctReq *int `json:"numAPAcctReq,omitempty"`

	// NumAPAcctRsp
	// Number of AP accounting responses
	NumAPAcctRsp *int `json:"numAPAcctRsp,omitempty"`

	// NumAuthOnlyAAA
	// Number of CoA authorize only processed
	NumAuthOnlyAAA *int `json:"numAuthOnlyAAA,omitempty"`

	// NumAuthOnlyFailedAAA
	// Number of failed CoA authorize only procedures
	NumAuthOnlyFailedAAA *int `json:"numAuthOnlyFailedAAA,omitempty"`

	// NumAuthOnlySuccAAA
	// Number of successful CoA authorize only procedures
	NumAuthOnlySuccAAA *int `json:"numAuthOnlySuccAAA,omitempty"`

	// NumCoaFailRcdNas
	// Number of Failed CoA response from NAS
	NumCoaFailRcdNas *int `json:"numCoaFailRcdNas,omitempty"`

	// NumCoaSntNas
	// Number of CoA sent to NAS
	NumCoaSntNas *int `json:"numCoaSntNas,omitempty"`

	// NumCoaSucRcdNas
	// Number of successful CoA responses from NAS
	NumCoaSucRcdNas *int `json:"numCoaSucRcdNas,omitempty"`

	// NumDmFailRcdNas
	// Number of RADIUS DM responses received from NAS (Failed)
	NumDmFailRcdNas *int `json:"numDmFailRcdNas,omitempty"`

	// NumDmSntNas
	// Number of RADIUS DM requests sent to NAS
	NumDmSntNas *int `json:"numDmSntNas,omitempty"`

	// NumDmSucRcdNas
	// Number of RADIUS DM responses received from NAS (Successful)
	NumDmSucRcdNas *int `json:"numDmSucRcdNas,omitempty"`

	// NumDroppedRateLimitAcct
	// Number of dropped accounting requests
	NumDroppedRateLimitAcct *int `json:"numDroppedRateLimitAcct,omitempty"`

	// NumDroppedRateLimitAuth
	// Number of dropped authentication requests
	NumDroppedRateLimitAuth *int `json:"numDroppedRateLimitAuth,omitempty"`

	// NumFailAcct
	// Number of accounting sessions established (Failed)
	NumFailAcct *int `json:"numFailAcct,omitempty"`

	// NumFailAuth
	// Number of authentications (Failed)
	NumFailAuth *int `json:"numFailAuth,omitempty"`

	// NumInCompAuth
	// Number of authentications (Incomplete)
	NumInCompAuth *int `json:"numInCompAuth,omitempty"`

	// NumOfAccAcceptAaa
	// Number of RADIUS access accepts received from AAA server
	NumOfAccAcceptAaa *int `json:"numOfAccAcceptAaa,omitempty"`

	// NumOfAccChallAaa
	// Number of RADIUS access challenges received from AAA server
	NumOfAccChallAaa *int `json:"numOfAccChallAaa,omitempty"`

	// NumOfAccRejAaa
	// Number of RADIUS access rejects received from AAA server
	NumOfAccRejAaa *int `json:"numOfAccRejAaa,omitempty"`

	// NumOfAccReqAaa
	// Number of RADIUS access requests sent to AAA server
	NumOfAccReqAaa *int `json:"numOfAccReqAaa,omitempty"`

	// NumOfAcctReqAaa
	// Number of RADIUS accounting requests sent to AAA server
	NumOfAcctReqAaa *int `json:"numOfAcctReqAaa,omitempty"`

	// NumOfAcctRspAaa
	// Number of RADIUS accounting responses received from AAA server
	NumOfAcctRspAaa *int `json:"numOfAcctRspAaa,omitempty"`

	// NumOfCoAFailAaa
	// Number of RADIUS CoA responses sent to AAA server (Failed)
	NumOfCoAFailAaa *int `json:"numOfCoAFailAaa,omitempty"`

	// NumOfCoAReqAaa
	// Number of RADIUS CoA responses sent to AAA server (Successful)
	NumOfCoAReqAaa *int `json:"numOfCoAReqAaa,omitempty"`

	// NumOfCoAResAaa
	// Number of RADIUS CoA requests received from AAA server
	NumOfCoAResAaa *int `json:"numOfCoAResAaa,omitempty"`

	// NumOfDmFailAaa
	// Number of RADIUS DM responses sent to AAA server (Failed)
	NumOfDmFailAaa *int `json:"numOfDmFailAaa,omitempty"`

	// NumOfDmReqAaa
	// Number of RADIUS DM responses sent to AAA server (Successful)
	NumOfDmReqAaa *int `json:"numOfDmReqAaa,omitempty"`

	// NumOfDmResAaa
	// Number of RADIUS DM requests received from AAA server
	NumOfDmResAaa *int `json:"numOfDmResAaa,omitempty"`

	// NumRadFailAPAcc
	// Number of AP accounting sessions established (Failed)
	NumRadFailAPAcc *int `json:"numRadFailAPAcc,omitempty"`

	// NumRadSuccAPAcc
	// Number of AP accounting sessions established (Successful)
	NumRadSuccAPAcc *int `json:"numRadSuccAPAcc,omitempty"`

	// NumSuccAcct
	// Number of accounting sessions established (Successful)
	NumSuccAcct *int `json:"numSuccAcct,omitempty"`

	// NumSuccAuth
	// Number of authentications (Successful)
	NumSuccAuth *int `json:"numSuccAuth,omitempty"`

	// RecCreateTime
	// Time this entry was created
	RecCreateTime *int `json:"recCreateTime,omitempty"`

	// RecUpdateTime
	// Last time this entry was updated
	RecUpdateTime *int `json:"recUpdateTime,omitempty"`
}

func NewWSGRACStatsRadiusProxy() *WSGRACStatsRadiusProxy {
	m := new(WSGRACStatsRadiusProxy)
	return m
}

// WSGRACStatsRadiusProxyList
//
// Definition: racStats_radiusProxyList
type WSGRACStatsRadiusProxyList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGRACStatsRadiusProxy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGRACStatsRadiusProxyListAPIResponse struct {
	*RawAPIResponse
	Data *WSGRACStatsRadiusProxyList
}

func newWSGRACStatsRadiusProxyListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGRACStatsRadiusProxyListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGRACStatsRadiusProxyListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGRACStatsRadiusProxyList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGRACStatsRadiusProxyList() *WSGRACStatsRadiusProxyList {
	m := new(WSGRACStatsRadiusProxyList)
	return m
}
