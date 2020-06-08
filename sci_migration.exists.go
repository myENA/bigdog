package ruckus

// API Version: 1.0.0

type SCIMigrationexistsgetMigrationsidexists200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationexistsgetMigrationsidexists200ResponseType() *SCIMigrationexistsgetMigrationsidexists200ResponseType {
	m := new(SCIMigrationexistsgetMigrationsidexists200ResponseType)
	return m
}

type SCIMigrationexistsheadMigrationsid200ResponseType struct {
	Exists *bool `json:"exists,omitempty"`
}

func NewSCIMigrationexistsheadMigrationsid200ResponseType() *SCIMigrationexistsheadMigrationsid200ResponseType {
	m := new(SCIMigrationexistsheadMigrationsid200ResponseType)
	return m
}
