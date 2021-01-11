package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGAVCAppCategory
//
// Definition: avc_appCategory
type WSGAVCAppCategory struct {
	// Id
	// Identifier of the Application Category
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the Application Category
	Name *string `json:"name,omitempty"`
}

type WSGAVCAppCategoryAPIResponse struct {
	*RawAPIResponse
	Data *WSGAVCAppCategory
}

func newWSGAVCAppCategoryAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAVCAppCategoryAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAVCAppCategoryAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAVCAppCategory)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAVCAppCategory() *WSGAVCAppCategory {
	m := new(WSGAVCAppCategory)
	return m
}

// WSGAVCAppCategoryList
//
// Definition: avc_appCategoryList
type WSGAVCAppCategoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAVCAppCategory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAVCAppCategoryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAVCAppCategoryList
}

func newWSGAVCAppCategoryListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAVCAppCategoryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAVCAppCategoryListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAVCAppCategoryList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAVCAppCategoryList() *WSGAVCAppCategoryList {
	m := new(WSGAVCAppCategoryList)
	return m
}

// WSGAVCApplication
//
// Definition: avc_application
type WSGAVCApplication struct {
	// AppId
	// Identifier of the Application
	AppId *string `json:"appId,omitempty"`

	// CatId
	// Identifier of the Application Category
	CatId *string `json:"catId,omitempty"`

	// Name
	// Name of the Application
	Name *string `json:"name,omitempty"`
}

type WSGAVCApplicationAPIResponse struct {
	*RawAPIResponse
	Data *WSGAVCApplication
}

func newWSGAVCApplicationAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAVCApplicationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAVCApplicationAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAVCApplication)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAVCApplication() *WSGAVCApplication {
	m := new(WSGAVCApplication)
	return m
}

// WSGAVCApplicationList
//
// Definition: avc_applicationList
type WSGAVCApplicationList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAVCApplication `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAVCApplicationListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAVCApplicationList
}

func newWSGAVCApplicationListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAVCApplicationListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAVCApplicationListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAVCApplicationList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAVCApplicationList() *WSGAVCApplicationList {
	m := new(WSGAVCApplicationList)
	return m
}

// WSGAVCApplicationPolicyProfile
//
// Definition: avc_applicationPolicyProfile
type WSGAVCApplicationPolicyProfile struct {
	ApplicationRules []*WSGAVCApplicationRule `json:"applicationRules,omitempty"`

	// AvcEventEnable
	// Send ARC logs from AP to SmartZone
	AvcEventEnable *bool `json:"avcEventEnable,omitempty"`

	// AvcLogEnable
	// Send ARC logs from AP to external syslog server
	AvcLogEnable *bool `json:"avcLogEnable,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Application Policy Profile belongs
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the Application Policy Profile
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`
}

type WSGAVCApplicationPolicyProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGAVCApplicationPolicyProfile
}

func newWSGAVCApplicationPolicyProfileAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAVCApplicationPolicyProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAVCApplicationPolicyProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAVCApplicationPolicyProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAVCApplicationPolicyProfile() *WSGAVCApplicationPolicyProfile {
	m := new(WSGAVCApplicationPolicyProfile)
	return m
}

// WSGAVCApplicationPolicyProfileList
//
// Definition: avc_applicationPolicyProfileList
type WSGAVCApplicationPolicyProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAVCApplicationPolicyProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAVCApplicationPolicyProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAVCApplicationPolicyProfileList
}

func newWSGAVCApplicationPolicyProfileListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAVCApplicationPolicyProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAVCApplicationPolicyProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAVCApplicationPolicyProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAVCApplicationPolicyProfileList() *WSGAVCApplicationPolicyProfileList {
	m := new(WSGAVCApplicationPolicyProfileList)
	return m
}

// WSGAVCApplicationRule
//
// Definition: avc_applicationRule
type WSGAVCApplicationRule struct {
	// AppId
	// Identifier of the Application from Signature Package
	AppId *string `json:"appId,omitempty"`

	// ApplicationType
	// Type of the application when ruleType
	// Constraints:
	//    - oneof:[SIGNATURE,USER_DEFINED]
	ApplicationType *string `json:"applicationType,omitempty"`

	// AppName
	// Name of the Application from Signature Package
	AppName *string `json:"appName,omitempty"`

	// CatId
	// Identifier of the Application Category from Signature Package (If applicationType is UserDefind, the catId is 32768)
	CatId *string `json:"catId,omitempty"`

	// CatName
	// Name of the Application Category from Signature Package
	CatName *string `json:"catName,omitempty"`

	// ClassificationType
	// QoS downlink classification type
	// Constraints:
	//    - oneof:[VOICE,VIDEO,BEST_EFFORT,BACKGROUND]
	ClassificationType *string `json:"classificationType,omitempty"`

	// Downlink
	// Downlink rate limiting (unit: Kbps)
	Downlink *int `json:"downlink,omitempty"`

	// MarkingPriority
	// QoS uplink marking priority
	// Constraints:
	//    - oneof:[IEEE802_1p,DSCP,BOTH]
	MarkingPriority *string `json:"markingPriority,omitempty"`

	// MarkingType
	// QoS uplink marking type
	// Constraints:
	//    - oneof:[VOICE,VIDEO,BEST_EFFORT,BACKGROUND]
	MarkingType *string `json:"markingType,omitempty"`

	Priority *int `json:"priority,omitempty"`

	// RuleType
	// Type of the application rule
	// Constraints:
	//    - oneof:[DENY,QOS,RATE_LIMITING]
	RuleType *string `json:"ruleType,omitempty"`

	// Uplink
	// Uplink rate limiting (unit: Kbps)
	Uplink *int `json:"uplink,omitempty"`
}

func NewWSGAVCApplicationRule() *WSGAVCApplicationRule {
	m := new(WSGAVCApplicationRule)
	return m
}

// WSGAVCCreateApplicationPolicyProfile
//
// Definition: avc_createApplicationPolicyProfile
type WSGAVCCreateApplicationPolicyProfile struct {
	// ApplicationRules
	// Constraints:
	//    - required
	ApplicationRules []*WSGAVCApplicationRule `json:"applicationRules"`

	// AvcEventEnable
	// Send ARC logs from AP to SmartZone
	AvcEventEnable *bool `json:"avcEventEnable,omitempty"`

	// AvcLogEnable
	// Send ARC logs from AP to external syslog server
	AvcLogEnable *bool `json:"avcLogEnable,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Application Policy Profile belongs
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGAVCCreateApplicationPolicyProfile() *WSGAVCCreateApplicationPolicyProfile {
	m := new(WSGAVCCreateApplicationPolicyProfile)
	return m
}

// WSGAVCCreateUserDefinedProfile
//
// Definition: avc_createUserDefinedProfile
type WSGAVCCreateUserDefinedProfile struct {
	DestIp *WSGCommonIpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	DestPort *int `json:"destPort"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the User Defined Profile belongs
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	Netmask *WSGCommonSubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	// Constraints:
	//    - required
	//    - oneof:[TCP,UDP]
	Protocol *string `json:"protocol"`

	// Type
	// Type of the User Defined Profile
	// Constraints:
	//    - required
	//    - oneof:[IP_WITH_PORT,PORT_ONLY]
	Type *string `json:"type"`
}

func NewWSGAVCCreateUserDefinedProfile() *WSGAVCCreateUserDefinedProfile {
	m := new(WSGAVCCreateUserDefinedProfile)
	return m
}

// WSGAVCDeleteBulk
//
// Definition: avc_deleteBulk
type WSGAVCDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGAVCDeleteBulk() *WSGAVCDeleteBulk {
	m := new(WSGAVCDeleteBulk)
	return m
}

// WSGAVCModifyApplicationPolicyProfile
//
// Definition: avc_modifyApplicationPolicyProfile
type WSGAVCModifyApplicationPolicyProfile struct {
	ApplicationRules []*WSGAVCApplicationRule `json:"applicationRules,omitempty"`

	// AvcEventEnable
	// Send ARC logs from AP to SmartZone
	AvcEventEnable *bool `json:"avcEventEnable,omitempty"`

	// AvcLogEnable
	// Send ARC logs from AP to external syslog server
	AvcLogEnable *bool `json:"avcLogEnable,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGAVCModifyApplicationPolicyProfile() *WSGAVCModifyApplicationPolicyProfile {
	m := new(WSGAVCModifyApplicationPolicyProfile)
	return m
}

// WSGAVCModifyUserDefinedProfile
//
// Definition: avc_modifyUserDefinedProfile
type WSGAVCModifyUserDefinedProfile struct {
	DestIp *WSGCommonIpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	// Constraints:
	//    - min:1
	//    - max:65535
	DestPort *int `json:"destPort,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Netmask *WSGCommonSubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	// Constraints:
	//    - oneof:[TCP,UDP]
	Protocol *string `json:"protocol,omitempty"`

	// Type
	// Type of the User Defined Profile
	// Constraints:
	//    - oneof:[IP_WITH_PORT,PORT_ONLY]
	Type *string `json:"type,omitempty"`
}

func NewWSGAVCModifyUserDefinedProfile() *WSGAVCModifyUserDefinedProfile {
	m := new(WSGAVCModifyUserDefinedProfile)
	return m
}

// WSGAVCSignaturePackage
//
// Definition: avc_signaturePackage
type WSGAVCSignaturePackage struct {
	// FileName
	// Name of the Signature Package
	FileName *string `json:"fileName,omitempty"`

	// Id
	// Identifier of the Signature Package
	Id *string `json:"id,omitempty"`

	// Size
	// Size of the Signature Package
	Size *int `json:"size,omitempty"`

	// Version
	// Version of the Signature Package
	Version *string `json:"version,omitempty"`
}

type WSGAVCSignaturePackageAPIResponse struct {
	*RawAPIResponse
	Data *WSGAVCSignaturePackage
}

func newWSGAVCSignaturePackageAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAVCSignaturePackageAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAVCSignaturePackageAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAVCSignaturePackage)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAVCSignaturePackage() *WSGAVCSignaturePackage {
	m := new(WSGAVCSignaturePackage)
	return m
}

// WSGAVCUserDefinedProfile
//
// Definition: avc_userDefinedProfile
type WSGAVCUserDefinedProfile struct {
	// AppId
	// AppId for Application Policy's User defined rule type
	AppId *int `json:"appId,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	DestIp *WSGCommonIpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	// Constraints:
	//    - min:1
	//    - max:65535
	DestPort *int `json:"destPort,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the User Defined Profile belongs
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the User Defined Profile
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Netmask *WSGCommonSubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	// Constraints:
	//    - oneof:[TCP,UDP]
	Protocol *string `json:"protocol,omitempty"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// Type
	// Type of the User Defined Profile
	// Constraints:
	//    - oneof:[IP_WITH_PORT,PORT_ONLY]
	Type *string `json:"type,omitempty"`
}

type WSGAVCUserDefinedProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGAVCUserDefinedProfile
}

func newWSGAVCUserDefinedProfileAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAVCUserDefinedProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAVCUserDefinedProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAVCUserDefinedProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAVCUserDefinedProfile() *WSGAVCUserDefinedProfile {
	m := new(WSGAVCUserDefinedProfile)
	return m
}

// WSGAVCUserDefinedProfileList
//
// Definition: avc_userDefinedProfileList
type WSGAVCUserDefinedProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAVCUserDefinedProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAVCUserDefinedProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAVCUserDefinedProfileList
}

func newWSGAVCUserDefinedProfileListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAVCUserDefinedProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAVCUserDefinedProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAVCUserDefinedProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAVCUserDefinedProfileList() *WSGAVCUserDefinedProfileList {
	m := new(WSGAVCUserDefinedProfileList)
	return m
}
