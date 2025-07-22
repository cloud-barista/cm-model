package cloudmodel

// type MigratedVmInfraModel struct {
// 	MigratedVmInfraModel MigratedVmInfraInfo `json:"migratedVmInfraModel" validate:"required"`
// }

type VmInfraInfo struct {
	TbMciInfo
}

type MciInfoList struct {
	Mci []TbMciInfo `json:"mci"`
}

type IdList struct {
	IdList []string `json:"idList"`
}
