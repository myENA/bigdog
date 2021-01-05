package bigdog

// API Version: 1.0.0

const (
	RouteSCIFacetGetApmacFacet                                                      = "/api/facets/apmac"
	RouteSCIFacetGetFacet                                                           = "/api/facets/{name}"
	RouteSCIFacetGetSsidFacet                                                       = "/api/facets/ssid"
	RouteSCIFacetGetSwitchesFacet                                                   = "/api/facets/switches"
	RouteSCIFacetGetSwitchHierarchyFacet                                            = "/api/facets/switchHierarchy"
	RouteSCIFacetGetSystemFacet                                                     = "/api/facets/system"
	RouteSCIMigrationCount                                                          = "/api/Migrations/count"
	RouteSCIMigrationCreate                                                         = "/api/Migrations"
	RouteSCIMigrationCreateChangeStreamGetMigrationsChangeStream                    = "/api/Migrations/change-stream"
	RouteSCIMigrationCreateChangeStreamPostMigrationsChangeStream                   = "/api/Migrations/change-stream"
	RouteSCIMigrationDeleteById                                                     = "/api/Migrations/{id}"
	RouteSCIMigrationExistsGetMigrationsIdExists                                    = "/api/Migrations/{id}/exists"
	RouteSCIMigrationFind                                                           = "/api/Migrations"
	RouteSCIMigrationFindById                                                       = "/api/Migrations/{id}"
	RouteSCIMigrationFindOne                                                        = "/api/Migrations/findOne"
	RouteSCIMigrationMigrateByName                                                  = "/api/Migrations/migrateByName"
	RouteSCIMigrationMigrateTo                                                      = "/api/Migrations/migrate"
	RouteSCIMigrationPrototypeUpdateAttributes                                      = "/api/Migrations/{id}"
	RouteSCIMigrationRollbackTo                                                     = "/api/Migrations/rollback"
	RouteSCIMigrationUpdateAll                                                      = "/api/Migrations/update"
	RouteSCIMigrationUpsert                                                         = "/api/Migrations"
	RouteSCIMigrationMapCount                                                       = "/api/MigrationMaps/count"
	RouteSCIMigrationMapCreate                                                      = "/api/MigrationMaps"
	RouteSCIMigrationMapCreateChangeStreamGetMigrationMapsChangeStream              = "/api/MigrationMaps/change-stream"
	RouteSCIMigrationMapCreateChangeStreamPostMigrationMapsChangeStream             = "/api/MigrationMaps/change-stream"
	RouteSCIMigrationMapDeleteById                                                  = "/api/MigrationMaps/{id}"
	RouteSCIMigrationMapExistsGetMigrationMapsIdExists                              = "/api/MigrationMaps/{id}/exists"
	RouteSCIMigrationMapFind                                                        = "/api/MigrationMaps"
	RouteSCIMigrationMapFindById                                                    = "/api/MigrationMaps/{id}"
	RouteSCIMigrationMapFindOne                                                     = "/api/MigrationMaps/findOne"
	RouteSCIMigrationMapPrototypeUpdateAttributes                                   = "/api/MigrationMaps/{id}"
	RouteSCIMigrationMapUpdateAll                                                   = "/api/MigrationMaps/update"
	RouteSCIMigrationMapUpsert                                                      = "/api/MigrationMaps"
	RouteSCIPciProfileBatchDelete                                                   = "/api/pciProfiles/batchDelete"
	RouteSCIPciProfileCreateWithRelations                                           = "/api/pciProfiles/createWithRelations"
	RouteSCIPciProfileFind                                                          = "/api/pciProfiles"
	RouteSCIPciProfilePrototypeCountReports                                         = "/api/pciProfiles/{id}/reports/count"
	RouteSCIPciProfilePrototypeCreateReports                                        = "/api/pciProfiles/{id}/reports"
	RouteSCIPciProfilePrototypeDeleteReports                                        = "/api/pciProfiles/{id}/reports"
	RouteSCIPciProfilePrototypeDestroyByIdReports                                   = "/api/pciProfiles/{id}/reports/{fk}"
	RouteSCIPciProfilePrototypeFindByIdReports                                      = "/api/pciProfiles/{id}/reports/{fk}"
	RouteSCIPciProfilePrototypeGetReports                                           = "/api/pciProfiles/{id}/reports"
	RouteSCIPciProfilePrototypeUpdateByIdReports                                    = "/api/pciProfiles/{id}/reports/{fk}"
	RouteSCIPciProfileUpdateWithRelations                                           = "/api/pciProfiles/{id}/updateWithRelations"
	RouteSCIPciReportDownloadReport                                                 = "/api/pciReports/download"
	RouteSCIReportAirtimeUtilizationReport1Overview                                 = "/api/reports/5/sections/1/data"
	RouteSCIReportAirtimeUtilizationReport2TopChart                                 = "/api/reports/5/sections/2/data"
	RouteSCIReportAirtimeUtilizationReport3TopAPsByAirtime24Table                   = "/api/reports/5/sections/3/data"
	RouteSCIReportAirtimeUtilizationReport4TopAPsByAirtime5Table                    = "/api/reports/5/sections/4/data"
	RouteSCIReportAirtimeUtilizationReport5TrendChart                               = "/api/reports/5/sections/5/data"
	RouteSCIReportAirtimeUtilizationReport6TrendTable                               = "/api/reports/5/sections/6/data"
	RouteSCIReportAPDetailsReport5TrendChart                                        = "/api/reports/11/sections/5/data"
	RouteSCIReportAPDetailsReport7Top10ApplicationsByTrafficVolume                  = "/api/reports/11/sections/7/data"
	RouteSCIReportAPDetailsReport8TopAppsByTrafficTable                             = "/api/reports/11/sections/8/data"
	RouteSCIReportAPDetailsReport14TopTable                                         = "/api/reports/11/sections/14/data"
	RouteSCIReportAPDetailsReport15TrendChart                                       = "/api/reports/11/sections/15/data"
	RouteSCIReportAPDetailsReport22TrafficTrend                                     = "/api/reports/11/sections/22/data"
	RouteSCIReportAPDetailsReport40TopSsidsByTrafficTable                           = "/api/reports/11/sections/40/data"
	RouteSCIReportAPDetailsReport75ApSummary                                        = "/api/reports/11/sections/75/data"
	RouteSCIReportAPDetailsReport76ApPerformance                                    = "/api/reports/11/sections/76/data"
	RouteSCIReportAPDetailsReport77ApDetails                                        = "/api/reports/11/sections/77/data"
	RouteSCIReportAPDetailsReport78ApStatsOverview                                  = "/api/reports/11/sections/78/data"
	RouteSCIReportAPDetailsReport79ApUptimeHistory                                  = "/api/reports/11/sections/79/data"
	RouteSCIReportAPDetailsReport80Top10ClientsByTrafficVolume                      = "/api/reports/11/sections/80/data"
	RouteSCIReportAPDetailsReport81SessionsTable                                    = "/api/reports/11/sections/81/data"
	RouteSCIReportAPDetailsReport82RssTrend                                         = "/api/reports/11/sections/82/data"
	RouteSCIReportAPDetailsReport83SnrTrend                                         = "/api/reports/11/sections/83/data"
	RouteSCIReportAPDetailsReport84AlarmsTable                                      = "/api/reports/11/sections/84/data"
	RouteSCIReportAPDetailsReport85EventsTable                                      = "/api/reports/11/sections/85/data"
	RouteSCIReportAPDetailsReport95Anomalies                                        = "/api/reports/11/sections/95/data"
	RouteSCIReportAPDetailsReport110ApAnomaly                                       = "/api/reports/11/sections/110/data"
	RouteSCIReportAPsRebootReport43TotalReboots                                     = "/api/reports/8/sections/43/data"
	RouteSCIReportAPsRebootReport44TopApRebootsTable                                = "/api/reports/8/sections/44/data"
	RouteSCIReportAPsRebootReport45TopApRebootsOverTime                             = "/api/reports/8/sections/45/data"
	RouteSCIReportClientDetailsReport7Top10ApplicationsByTrafficVolume              = "/api/reports/12/sections/7/data"
	RouteSCIReportClientDetailsReport8TopAppsByTrafficTable                         = "/api/reports/12/sections/8/data"
	RouteSCIReportClientDetailsReport82RssTrend                                     = "/api/reports/12/sections/82/data"
	RouteSCIReportClientDetailsReport83SnrTrend                                     = "/api/reports/12/sections/83/data"
	RouteSCIReportClientDetailsReport86Summary                                      = "/api/reports/12/sections/86/data"
	RouteSCIReportClientDetailsReport87ClientStats                                  = "/api/reports/12/sections/87/data"
	RouteSCIReportClientDetailsReport89TrafficTrend                                 = "/api/reports/12/sections/89/data"
	RouteSCIReportClientDetailsReport92SessionsTable                                = "/api/reports/12/sections/92/data"
	RouteSCIReportClientHealthReport144ClientHealthSummary                          = "/api/reports/20/sections/144/data"
	RouteSCIReportClientHealthReport148ClientConnectionHealth                       = "/api/reports/20/sections/148/data"
	RouteSCIReportClientHealthReport149ClientHealthMetricTrends                     = "/api/reports/20/sections/149/data"
	RouteSCIReportClientHealthReport150TopClientHealthScoreByGroup                  = "/api/reports/20/sections/150/data"
	RouteSCIReportComparisonReport140ComparisionOverview                            = "/api/reports/19/sections/140/data"
	RouteSCIReportComparisonReport145ComparisionMetric1                             = "/api/reports/19/sections/145/data"
	RouteSCIReportComparisonReport146ComparisionMetric2                             = "/api/reports/19/sections/146/data"
	RouteSCIReportComparisonReport147ComparisionTable                               = "/api/reports/19/sections/147/data"
	RouteSCIReportDownloadReport                                                    = "/api/reports/{id}/download/{format}"
	RouteSCIReportFind                                                              = "/api/reports"
	RouteSCIReportFindById                                                          = "/api/reports/{id}"
	RouteSCIReportGetData                                                           = "/api/reports/{id}/sections/{sectionId}/data"
	RouteSCIReportInventoryAPsReport46ApInventoryOverview                           = "/api/reports/9/sections/46/data"
	RouteSCIReportInventoryAPsReport47TopApsDisconnection                           = "/api/reports/9/sections/47/data"
	RouteSCIReportInventoryAPsReport48ApCountTrend                                  = "/api/reports/9/sections/48/data"
	RouteSCIReportInventoryAPsReport49ApStatusTrend                                 = "/api/reports/9/sections/49/data"
	RouteSCIReportInventoryAPsReport50TopApsModelsChart                             = "/api/reports/9/sections/50/data"
	RouteSCIReportInventoryAPsReport51Top10ApVersionsChart                          = "/api/reports/9/sections/51/data"
	RouteSCIReportInventoryAPsReport52TopApsRebootReasons                           = "/api/reports/9/sections/52/data"
	RouteSCIReportInventoryAPsReport53Top10ApsRebootCounts                          = "/api/reports/9/sections/53/data"
	RouteSCIReportInventoryAPsReport54TopApAlarmTypes                               = "/api/reports/9/sections/54/data"
	RouteSCIReportInventoryAPsReport55TopAPModels                                   = "/api/reports/9/sections/55/data"
	RouteSCIReportInventoryAPsReport56TopAPVersions                                 = "/api/reports/9/sections/56/data"
	RouteSCIReportInventoryAPsReport57TopAPsOffline                                 = "/api/reports/9/sections/57/data"
	RouteSCIReportInventoryAPsReport58TopAPsByReboots                               = "/api/reports/9/sections/58/data"
	RouteSCIReportInventoryAPsReport59ApsConfiguredInMultiCtrl                      = "/api/reports/9/sections/59/data"
	RouteSCIReportInventoryAPsReport60ApDetailsOnOfflineStatus                      = "/api/reports/9/sections/60/data"
	RouteSCIReportInventoryAPsReport61ApDetailsOtherStatus                          = "/api/reports/9/sections/61/data"
	RouteSCIReportInventoryControllersReport96Krack                                 = "/api/reports/15/sections/96/data"
	RouteSCIReportInventoryControllersReport98ResourceUtilization                   = "/api/reports/15/sections/98/data"
	RouteSCIReportInventoryControllersReport99LicenseUtilization                    = "/api/reports/15/sections/99/data"
	RouteSCIReportInventoryControllersReport114ControllerInventoryOverview          = "/api/reports/15/sections/114/data"
	RouteSCIReportInventoryControllersReport151LicenseUtilizationOverTimeChart      = "/api/reports/15/sections/151/data"
	RouteSCIReportInventorySwitchesReport113Overview                                = "/api/reports/16/sections/113/data"
	RouteSCIReportInventorySwitchesReport116SwitchCountTrend                        = "/api/reports/16/sections/116/data"
	RouteSCIReportInventorySwitchesReport117Top10SwitchVersionChart                 = "/api/reports/16/sections/117/data"
	RouteSCIReportInventorySwitchesReport118TopSwitchVersions                       = "/api/reports/16/sections/118/data"
	RouteSCIReportInventorySwitchesReport121TopSwitchModelsChart                    = "/api/reports/16/sections/121/data"
	RouteSCIReportInventorySwitchesReport122TopSwitchModels                         = "/api/reports/16/sections/122/data"
	RouteSCIReportInventorySwitchesReport132PortStatusTrend                         = "/api/reports/16/sections/132/data"
	RouteSCIReportLatestIngestedTime                                                = "/api/reports/latestIngestedTime"
	RouteSCIReportNetworkWiredReport123TopSwitchPOEUtilChart                        = "/api/reports/17/sections/123/data"
	RouteSCIReportNetworkWiredReport124TopSwitchPOEUtils                            = "/api/reports/17/sections/124/data"
	RouteSCIReportNetworkWiredReport127Top10SwitchesByTrafficVolume                 = "/api/reports/17/sections/127/data"
	RouteSCIReportNetworkWiredReport128TopSwitchesByTrafficTable                    = "/api/reports/17/sections/128/data"
	RouteSCIReportNetworkWiredReport134WiredOverview                                = "/api/reports/17/sections/134/data"
	RouteSCIReportNetworkWiredReport135WiredTrafficDistribution                     = "/api/reports/17/sections/135/data"
	RouteSCIReportNetworkWiredReport136SwitchTrafficTrend                           = "/api/reports/17/sections/136/data"
	RouteSCIReportNetworkWiredReport141SwitchErrorTrend                             = "/api/reports/17/sections/141/data"
	RouteSCIReportNetworkWiredReport142TopSwitchesByErrorsChart                     = "/api/reports/17/sections/142/data"
	RouteSCIReportNetworkWiredReport143TopSwitchesByErrorsTable                     = "/api/reports/17/sections/143/data"
	RouteSCIReportNetworkWirelessReport20Overview                                   = "/api/reports/2/sections/20/data"
	RouteSCIReportNetworkWirelessReport21TrafficDistribution                        = "/api/reports/2/sections/21/data"
	RouteSCIReportNetworkWirelessReport22TrafficTrend                               = "/api/reports/2/sections/22/data"
	RouteSCIReportNetworkWirelessReport23TrafficOverTimeTable                       = "/api/reports/2/sections/23/data"
	RouteSCIReportNetworkWirelessReport24TopAPsByTrafficTable                       = "/api/reports/2/sections/24/data"
	RouteSCIReportNetworkWirelessReport25TopAPsByClientsTable                       = "/api/reports/2/sections/25/data"
	RouteSCIReportNetworkWirelessReport26Top10APsByTrafficVolume                    = "/api/reports/2/sections/26/data"
	RouteSCIReportNetworkWirelessReport27Top10ApByClientCount                       = "/api/reports/2/sections/27/data"
	RouteSCIReportOverview62Overview                                                = "/api/reports/10/sections/62/data"
	RouteSCIReportOverview63Controller                                              = "/api/reports/10/sections/63/data"
	RouteSCIReportOverview64ApOverview                                              = "/api/reports/10/sections/64/data"
	RouteSCIReportOverview66ApAlarmOverview                                         = "/api/reports/10/sections/66/data"
	RouteSCIReportOverview67SwitchOverview                                          = "/api/reports/10/sections/67/data"
	RouteSCIReportOverview68ApClientCountOverview                                   = "/api/reports/10/sections/68/data"
	RouteSCIReportOverview69TotalTrafficMinMaxRate                                  = "/api/reports/10/sections/69/data"
	RouteSCIReportOverview70SessionsOverview                                        = "/api/reports/10/sections/70/data"
	RouteSCIReportOverview71SsidOverview                                            = "/api/reports/10/sections/71/data"
	RouteSCIReportOverview72RadioOverview                                           = "/api/reports/10/sections/72/data"
	RouteSCIReportOverview73ApplicationsOverview                                    = "/api/reports/10/sections/73/data"
	RouteSCIReportOverview74ApEventOverview                                         = "/api/reports/10/sections/74/data"
	RouteSCIReportOverview97FactOverview                                            = "/api/reports/10/sections/97/data"
	RouteSCIReportOverview115NetworkUsageOverview                                   = "/api/reports/10/sections/115/data"
	RouteSCIReportPrototypeGetSections                                              = "/api/reports/{id}/sections"
	RouteSCIReportSCNetworkTrafficReport93ScNetworkTraffic                          = "/api/reports/13/sections/93/data"
	RouteSCIReportSCNetworkTrafficReport94ScNetworkTrend                            = "/api/reports/13/sections/94/data"
	RouteSCIReportSCNetworkTrafficReport100DroppedCallRate                          = "/api/reports/13/sections/100/data"
	RouteSCIReportSCNetworkTrafficReport101ConnectionSetupSuccessRate               = "/api/reports/13/sections/101/data"
	RouteSCIReportSCNetworkTrafficReport102HandoverSuccessRate                      = "/api/reports/13/sections/102/data"
	RouteSCIReportSCNetworkTrafficReport103AvgThroughput                            = "/api/reports/13/sections/103/data"
	RouteSCIReportSCNetworkTrafficReport104ScAvailability                           = "/api/reports/13/sections/104/data"
	RouteSCIReportSCNetworkTrafficReport105RscConnectionStats                       = "/api/reports/13/sections/105/data"
	RouteSCIReportSCNetworkTrafficReport106RscGpsStats                              = "/api/reports/13/sections/106/data"
	RouteSCIReportSCNetworkTrafficReport107TrafficVolume                            = "/api/reports/13/sections/107/data"
	RouteSCIReportSCNetworkTrafficReport108PhaseSyncLoss                            = "/api/reports/13/sections/108/data"
	RouteSCIReportSCNetworkTrafficReport109FrequencySyncLoss                        = "/api/reports/13/sections/109/data"
	RouteSCIReportSCNetworkTrafficReport111RscTrafficStats                          = "/api/reports/13/sections/111/data"
	RouteSCIReportSessionsSummaryReport33TopTable                                   = "/api/reports/6/sections/33/data"
	RouteSCIReportSessionsSummaryReport34Overview                                   = "/api/reports/6/sections/34/data"
	RouteSCIReportSessionsSummaryReport42DurationPercentile                         = "/api/reports/6/sections/42/data"
	RouteSCIReportSwitchDetailsReport125SwitchSummary                               = "/api/reports/18/sections/125/data"
	RouteSCIReportSwitchDetailsReport126SwitchResourceUtilization                   = "/api/reports/18/sections/126/data"
	RouteSCIReportSwitchDetailsReport129TopSwitchPortsByTrafficChart                = "/api/reports/18/sections/129/data"
	RouteSCIReportSwitchDetailsReport130TopSwitchPortsByTrafficTable                = "/api/reports/18/sections/130/data"
	RouteSCIReportSwitchDetailsReport131SwitchTrafficTrend                          = "/api/reports/18/sections/131/data"
	RouteSCIReportSwitchDetailsReport137LldpNeighborTable                           = "/api/reports/18/sections/137/data"
	RouteSCIReportSwitchDetailsReport138SwitchUptimeHistory                         = "/api/reports/18/sections/138/data"
	RouteSCIReportSwitchDetailsReport139SwitchDetails                               = "/api/reports/18/sections/139/data"
	RouteSCIReportSwitchDetailsReport152PerSwitchErrorTrend                         = "/api/reports/18/sections/152/data"
	RouteSCIReportWirelessApplicationsReport7Top10ApplicationsByTrafficVolume       = "/api/reports/3/sections/7/data"
	RouteSCIReportWirelessApplicationsReport8TopAppsByTrafficTable                  = "/api/reports/3/sections/8/data"
	RouteSCIReportWirelessApplicationsReport9TopAppsByClientsTable                  = "/api/reports/3/sections/9/data"
	RouteSCIReportWirelessApplicationsReport10Overview                              = "/api/reports/3/sections/10/data"
	RouteSCIReportWirelessApplicationsReport11Top10ApplicationsByClientCount        = "/api/reports/3/sections/11/data"
	RouteSCIReportWirelessClientsReport12Overview                                   = "/api/reports/1/sections/12/data"
	RouteSCIReportWirelessClientsReport13TopChart                                   = "/api/reports/1/sections/13/data"
	RouteSCIReportWirelessClientsReport14TopTable                                   = "/api/reports/1/sections/14/data"
	RouteSCIReportWirelessClientsReport15TrendChart                                 = "/api/reports/1/sections/15/data"
	RouteSCIReportWirelessClientsReport16TrendTable                                 = "/api/reports/1/sections/16/data"
	RouteSCIReportWirelessClientsReport17TopPercentile                              = "/api/reports/1/sections/17/data"
	RouteSCIReportWirelessClientsReport18TopNOSByClientCount                        = "/api/reports/1/sections/18/data"
	RouteSCIReportWirelessClientsReport19Top10ManufacturersByClientCount            = "/api/reports/1/sections/19/data"
	RouteSCIReportWirelessClientsReport112Top10AuthenticationMechanismByClientCount = "/api/reports/1/sections/112/data"
	RouteSCIReportWithRelations                                                     = "/api/reports/withRelations"
	RouteSCIReportWLANsReport35Overview                                             = "/api/reports/4/sections/35/data"
	RouteSCIReportWLANsReport36Top10SsidsByTraffic                                  = "/api/reports/4/sections/36/data"
	RouteSCIReportWLANsReport37ActiveSsidsTrend                                     = "/api/reports/4/sections/37/data"
	RouteSCIReportWLANsReport38Top10SsidsByClientCount                              = "/api/reports/4/sections/38/data"
	RouteSCIReportWLANsReport39SsidChangesOverTime                                  = "/api/reports/4/sections/39/data"
	RouteSCIReportWLANsReport40TopSsidsByTrafficTable                               = "/api/reports/4/sections/40/data"
	RouteSCIReportWLANsReport41TopSsidsByClientsTable                               = "/api/reports/4/sections/41/data"
	RouteSCIResourceGroupBatchDelete                                                = "/api/resourceGroups/batchDelete"
	RouteSCIResourceGroupCreate                                                     = "/api/resourceGroups"
	RouteSCIResourceGroupFind                                                       = "/api/resourceGroups"
	RouteSCIResourceGroupFindById                                                   = "/api/resourceGroups/{id}"
	RouteSCIResourceGroupPrototypeUpdateAttributes                                  = "/api/resourceGroups/{id}"
	RouteSCIScheduleBatchDelete                                                     = "/api/schedules/batchDelete"
	RouteSCIScheduleCreateWithRelations                                             = "/api/schedules/createWithRelations"
	RouteSCIScheduleExecuteJob                                                      = "/api/schedules/executeJob"
	RouteSCIScheduleUpdateWithRelations                                             = "/api/schedules/{id}/updateWithRelations"
	RouteSCISettingDeleteById                                                       = "/api/settings/{id}"
	RouteSCISettingFindById                                                         = "/api/settings/{id}"
	RouteSCISettingSendTestEmail                                                    = "/api/settings/sendTestEmail"
	RouteSCISettingUpsert                                                           = "/api/settings"
	RouteSCISystemCreate                                                            = "/api/systems"
	RouteSCISystemDeleteById                                                        = "/api/systems/{id}"
	RouteSCISystemFind                                                              = "/api/systems"
	RouteSCISystemFindById                                                          = "/api/systems/{id}"
	RouteSCISystemGetSsids                                                          = "/api/systems/getSsids"
	RouteSCISystemPrototypeUpdateAttributes                                         = "/api/systems/{id}"
	RouteSCIUserBatchDelete                                                         = "/api/users/batchDelete"
	RouteSCIUserCreateWithRelations                                                 = "/api/users/createWithRelations"
	RouteSCIUserFindById                                                            = "/api/users/{id}"
	RouteSCIUserGetResourceGroupsForUpsert                                          = "/api/users/getResourceGroupsForUpsert"
	RouteSCIUserGetUsers                                                            = "/api/users"
	RouteSCIUserLogin                                                               = "/api/users/login"
	RouteSCIUserLogout                                                              = "/api/users/logout"
	RouteSCIUserPrototypeCreateFilters                                              = "/api/users/{id}/filters"
	RouteSCIUserPrototypeDestroyByIdFilters                                         = "/api/users/{id}/filters/{fk}"
	RouteSCIUserPrototypeFindByIdFilters                                            = "/api/users/{id}/filters/{fk}"
	RouteSCIUserPrototypeGetFilters                                                 = "/api/users/{id}/filters"
	RouteSCIUserPrototypeGetSchedules                                               = "/api/users/{id}/schedules"
	RouteSCIUserPrototypeUpdateAttributes                                           = "/api/users/{id}"
	RouteSCIUserPrototypeUpdateByIdFilters                                          = "/api/users/{id}/filters/{fk}"
	RouteSCIUserUpdateWithRelations                                                 = "/api/users/{id}/updateWithRelations"
	RouteSCIUserValidateCurrentPassword                                             = "/api/users/validateCurrentPassword"
	RouteSCIZdXmlGetAjaxRequest                                                     = "/api/zdXmls/getAjaxRequest"
	RouteSCIZdXmlUpload                                                             = "/api/zdXmls/{container}/upload"
)
