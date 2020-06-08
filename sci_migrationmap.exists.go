package ruckus

// API Version: 1.0.0

type SCIMigrationmapexistsgetMigrationMapsidexists200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationmapexistsgetMigrationMapsidexists200ResponseType() *SCIMigrationmapexistsgetMigrationMapsidexists200ResponseType {
	m := new(SCIMigrationmapexistsgetMigrationMapsidexists200ResponseType)
	return m
}

type SCIMigrationmapexistsheadMigrationMapsid200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationmapexistsheadMigrationMapsid200ResponseType() *SCIMigrationmapexistsheadMigrationMapsid200ResponseType {
	m := new(SCIMigrationmapexistsheadMigrationMapsid200ResponseType)
	return m
}
