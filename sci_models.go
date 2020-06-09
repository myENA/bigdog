package bigdog

// API Version: 1.0.0

import (
	"encoding/json"
)

type SCIModelsFilter struct {
	// Filter
	// {"type":"and","fields": [{"type":"or","fields": [{"type":"selector","dimension":"radio","value":"2.4"}]}]}}
	// Constraints:
	//    - required
	//    - max:1000000
	Filter *string `json:"filter"`

	Id *float64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	ReportId *float64 `json:"reportId,omitempty"`

	UrlSegmentName *string `json:"urlSegmentName,omitempty"`

	UserId *float64 `json:"userId,omitempty"`
}

func NewSCIModelsFilter() *SCIModelsFilter {
	m := new(SCIModelsFilter)
	return m
}

// SCIModelsMigration
//
// View and run pending migrations.
type SCIModelsMigration struct {
	Id *float64 `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// RunDtTm
	// Constraints:
	//    - required
	RunDtTm *string `json:"runDtTm"`
}

func NewSCIModelsMigration() *SCIModelsMigration {
	m := new(SCIModelsMigration)
	return m
}

// SCIModelsMigrationMap
//
// Migration Mappings.
type SCIModelsMigrationMap struct {
	Data *SCIModelsMigrationMapDataType `json:"data,omitempty"`

	// From
	// Constraints:
	//    - required
	From *string `json:"from"`

	Id *float64 `json:"id,omitempty"`

	// To
	// Constraints:
	//    - required
	To *string `json:"to"`

	// Type
	// Constraints:
	//    - required
	Type *string `json:"type"`
}

func NewSCIModelsMigrationMap() *SCIModelsMigrationMap {
	m := new(SCIModelsMigrationMap)
	return m
}

type SCIModelsMigrationMapDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIModelsMigrationMapDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIModelsMigrationMapDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIModelsMigrationMapDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIModelsMigrationMapDataType() *SCIModelsMigrationMapDataType {
	m := new(SCIModelsMigrationMapDataType)
	return m
}

type SCIModelsPciProfile struct {
	// Answers
	// list of selected ssids
	// Constraints:
	//    - required
	Answers *SCIModelsPciProfileAnswersType `json:"answers"`

	Id *float64 `json:"id,omitempty"`

	// Name
	// Name of the profile
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// Ssids
	// list of selected ssids
	Ssids []SCIModelsXany `json:"ssids,omitempty"`
}

func NewSCIModelsPciProfile() *SCIModelsPciProfile {
	m := new(SCIModelsPciProfile)
	return m
}

// SCIModelsPciProfileAnswersType
//
// list of selected ssids
type SCIModelsPciProfileAnswersType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIModelsPciProfileAnswersType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIModelsPciProfileAnswersType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIModelsPciProfileAnswersType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIModelsPciProfileAnswersType() *SCIModelsPciProfileAnswersType {
	m := new(SCIModelsPciProfileAnswersType)
	return m
}

type SCIModelsPciReport struct {
	// Date
	// Date of report generation
	// Constraints:
	//    - required
	//    - default:'$now'
	Date *string `json:"date"`

	Id *float64 `json:"id,omitempty"`

	ProfileId *float64 `json:"profileId,omitempty"`

	// Statuses
	// Status of the controllers in this report
	// Constraints:
	//    - required
	Statuses []*SCIModelsPciReportStatusesType `json:"statuses"`
}

func NewSCIModelsPciReport() *SCIModelsPciReport {
	m := new(SCIModelsPciReport)
	return m
}

type SCIModelsPciReportStatusesType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIModelsPciReportStatusesType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIModelsPciReportStatusesType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIModelsPciReportStatusesType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIModelsPciReportStatusesType() *SCIModelsPciReportStatusesType {
	m := new(SCIModelsPciReportStatusesType)
	return m
}

type SCIModelsReport struct {
	Component *string `json:"component,omitempty"`

	DatasourcesUsed []SCIModelsXany `json:"datasourcesUsed,omitempty"`

	ExcludedFilters []SCIModelsXany `json:"excludedFilters,omitempty"`

	FilterDataSource *string `json:"filterDataSource,omitempty"`

	Headers []SCIModelsXany `json:"headers,omitempty"`

	Id *float64 `json:"id,omitempty"`

	Layout []SCIModelsXany `json:"layout,omitempty"`

	RouteParameters *SCIModelsReportRouteParametersType `json:"routeParameters,omitempty"`

	// Title
	// Constraints:
	//    - required
	Title *string `json:"title"`

	// UrlSegmentName
	// Constraints:
	//    - required
	UrlSegmentName *string `json:"urlSegmentName"`
}

func NewSCIModelsReport() *SCIModelsReport {
	m := new(SCIModelsReport)
	return m
}

type SCIModelsReportRouteParametersType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIModelsReportRouteParametersType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIModelsReportRouteParametersType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIModelsReportRouteParametersType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIModelsReportRouteParametersType() *SCIModelsReportRouteParametersType {
	m := new(SCIModelsReportRouteParametersType)
	return m
}

type SCIModelsResourceGroup struct {
	Description *string `json:"description,omitempty"`

	// Filter
	// {"type":"and","fields": [{"type":"or","fields": [{"type":"selector","dimension":"radio","value":"2.4"}]}]}}
	// Constraints:
	//    - required
	//    - max:1000000
	Filter *string `json:"filter"`

	Id *float64 `json:"id,omitempty"`

	// IsDefault
	// Constraints:
	//    - required
	//    - default:false
	IsDefault *bool `json:"isDefault"`

	LastUpdated *string `json:"lastUpdated,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *string `json:"name"`
}

func NewSCIModelsResourceGroup() *SCIModelsResourceGroup {
	m := new(SCIModelsResourceGroup)
	return m
}

type SCIModelsSchedule struct {
	// Day
	// 1 for 1st or monday
	Day *float64 `json:"day,omitempty"`

	// Enabled
	// 0 for disabled, 1 for enabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled"`

	FilterId *float64 `json:"filterId,omitempty"`

	// Format
	// pdf, csv
	// Constraints:
	//    - required
	Format *string `json:"format"`

	// Frequency
	// daily, weekly, monthly, quarterly, yearly
	// Constraints:
	//    - required
	Frequency *string `json:"frequency"`

	// Hour
	// 0 for midnight, 13 for 1pm, UTC
	Hour *float64 `json:"hour,omitempty"`

	Id *float64 `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// Recipients
	// one@domain.com
	// two@domain.com
	// Constraints:
	//    - required
	Recipients []SCIModelsXany `json:"recipients"`

	ReportId *float64 `json:"reportId,omitempty"`

	// Timezone
	// Timezone
	// Constraints:
	//    - required
	Timezone *string `json:"timezone"`

	UserId *float64 `json:"userId,omitempty"`
}

func NewSCIModelsSchedule() *SCIModelsSchedule {
	m := new(SCIModelsSchedule)
	return m
}

type SCIModelsSection struct {
	Component *string `json:"component,omitempty"`

	DefaultParameters *SCIModelsSectionDefaultParametersType `json:"defaultParameters,omitempty"`

	Id *float64 `json:"id,omitempty"`

	Layout *SCIModelsSectionLayoutType `json:"layout,omitempty"`

	// QueryName
	// Constraints:
	//    - required
	QueryName *string `json:"queryName"`

	SystemOwnerOnly *bool `json:"systemOwnerOnly,omitempty"`

	// Title
	// Constraints:
	//    - required
	Title *string `json:"title"`

	Url *string `json:"url,omitempty"`
}

func NewSCIModelsSection() *SCIModelsSection {
	m := new(SCIModelsSection)
	return m
}

type SCIModelsSectionDefaultParametersType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIModelsSectionDefaultParametersType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIModelsSectionDefaultParametersType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIModelsSectionDefaultParametersType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIModelsSectionDefaultParametersType() *SCIModelsSectionDefaultParametersType {
	m := new(SCIModelsSectionDefaultParametersType)
	return m
}

type SCIModelsSectionLayoutType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIModelsSectionLayoutType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIModelsSectionLayoutType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIModelsSectionLayoutType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIModelsSectionLayoutType() *SCIModelsSectionLayoutType {
	m := new(SCIModelsSectionLayoutType)
	return m
}

type SCIModelsSetting struct {
	// Key
	// Constraints:
	//    - required
	Key *string `json:"key"`

	// Values
	// Constraints:
	//    - required
	Values *SCIModelsSettingValuesType `json:"values"`
}

func NewSCIModelsSetting() *SCIModelsSetting {
	m := new(SCIModelsSetting)
	return m
}

type SCIModelsSettingValuesType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIModelsSettingValuesType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIModelsSettingValuesType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIModelsSettingValuesType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIModelsSettingValuesType() *SCIModelsSettingValuesType {
	m := new(SCIModelsSettingValuesType)
	return m
}

type SCIModelsSystem struct {
	BackupLocation *string `json:"backupLocation,omitempty"`

	// Format
	// Constraints:
	//    - required
	Format *string `json:"format"`

	// Id
	// Constraints:
	//    - required
	Id *string `json:"id"`

	// LastContact
	// last contacted/aggregated time
	LastContact *string `json:"lastContact,omitempty"`

	// Location
	// Constraints:
	//    - required
	Location *string `json:"location"`

	PciData *SCIModelsSystemPciDataType `json:"pciData,omitempty"`

	// Type
	// Constraints:
	//    - required
	Type *string `json:"type"`

	// User
	// Constraints:
	//    - required
	User *string `json:"user"`

	Version *string `json:"version,omitempty"`
}

func NewSCIModelsSystem() *SCIModelsSystem {
	m := new(SCIModelsSystem)
	return m
}

type SCIModelsSystemPciDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIModelsSystemPciDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIModelsSystemPciDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIModelsSystemPciDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIModelsSystemPciDataType() *SCIModelsSystemPciDataType {
	m := new(SCIModelsSystemPciDataType)
	return m
}

type SCIModelsUser struct {
	// Email
	// Constraints:
	//    - required
	Email *string `json:"email"`

	// FirstName
	// Constraints:
	//    - required
	FirstName *string `json:"firstName"`

	Id *float64 `json:"id,omitempty"`

	IsExternal *bool `json:"isExternal,omitempty"`

	LastLogin *string `json:"lastLogin,omitempty"`

	// LastName
	// Constraints:
	//    - required
	LastName *string `json:"lastName"`

	Username *string `json:"username,omitempty"`
}

func NewSCIModelsUser() *SCIModelsUser {
	m := new(SCIModelsUser)
	return m
}

// SCIModelsUserLogin
//
// Credentials used to log a user in
type SCIModelsUserLogin struct {
	// Password
	// Constraints:
	//    - required
	Password *string `json:"password"`

	// Username
	// Constraints:
	//    - required
	Username *string `json:"username"`
}

func NewSCIModelsUserLogin() *SCIModelsUserLogin {
	m := new(SCIModelsUserLogin)
	return m
}

type SCIModelsXany interface{}

func MakeSCIModelsXany() SCIModelsXany {
	m := new(SCIModelsXany)
	return m
}
