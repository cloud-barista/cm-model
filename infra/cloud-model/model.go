package cloudmodel

// RecommendedVmInfraModel represents the recommended virtual machine infrastructure model.
type RecommendedVmInfraModel struct {
	RecommendedVmInfraModel RecommendedVmInfra `json:"recommendedVmInfraModel" validate:"required"`
}

// RecommendedVmInfra represents the recommended virtual machine infrastructure information.
type RecommendedVmInfra struct {
	Status                  string               `json:"status"`
	Description             string               `json:"description"`
	TargetVmInfra           TbMciReq             `json:"targetVmInfra"`
	TargetVNet              TbVNetReq            `json:"targetVNet"`
	TargetSshKey            TbSshKeyReq          `json:"targetSshKey"`
	TargetVmSpecList        []TbSpecInfo         `json:"targetVmSpecList"`
	TargetVmOsImageList     []TbImageInfo        `json:"targetVmOsImageList"`
	TargetSecurityGroupList []TbSecurityGroupReq `json:"targetSecurityGroupList"`
}

// RecommendedVmInfraDynamic represents the recommended virtual machine infrastructure information.
type RecommendedVmInfraDynamic struct {
	Status        string          `json:"status"`
	Description   string          `json:"description"`
	TargetVmInfra TbMciDynamicReq `json:"targetVmInfra"`
}

// RecommendedVmInfraDynamicList represents a list of recommended virtual machine infrastructure information.
type RecommendedVmInfraDynamicList struct {
	Description       string                      `json:"description"`
	Count             int                         `json:"count"`
	TargetVmInfraList []RecommendedVmInfraDynamic `json:"targetVmInfraList"`
}

// RecommendedVNet represents the recommended virtual network information.
// * May be mainly used this object
type RecommendedVNet struct {
	Status      string    `json:"status"`
	Description string    `json:"description"`
	TargetVNet  TbVNetReq `json:"targetVNet"`
}

// RecommendedVNetList represents a list of recommended virtual network information.
type RecommendedVNetList struct {
	Description    string            `json:"description"`
	Count          int               `json:"count"`
	TargetVNetList []RecommendedVNet `json:"targetVNetList"`
}

// RecommendedSecurityGroup represents the recommended security group information.
type RecommendedSecurityGroup struct {
	Status              string             `json:"status"`
	SourceServers       []string           `json:"sourceServers"`
	Description         string             `json:"description"`
	TargetSecurityGroup TbSecurityGroupReq `json:"targetSecurityGroup"`
}

// RecommendedSecurityGroupList represents a list of recommended security group information.
type RecommendedSecurityGroupList struct {
	Status                  string                     `json:"status"`
	Description             string                     `json:"description"`
	Count                   int                        `json:"count"`
	TargetSecurityGroupList []RecommendedSecurityGroup `json:"targetSecurityGroupList"`
}

// RecommendedVmSpec represents the recommended virtual machine specification information for a single server.
type RecommendedVmSpec struct {
	Status        string     `json:"status"`
	SourceServers []string   `json:"sourceServers"`
	Description   string     `json:"description"`
	TargetVmSpec  TbSpecInfo `json:"targetVmSpec"`
}

// RecommendedVmSpecList represents a collection of recommended VM specifications across multiple source servers.
type RecommendedVmSpecList struct {
	Status                string              `json:"status"`
	Description           string              `json:"description"`
	Count                 int                 `json:"count"`
	RecommendedVmSpecList []RecommendedVmSpec `json:"recommendedVmSpecList"`
}

// RecommendedVmOsImage represents the recommended virtual machine OS image information for a single server.
type RecommendedVmOsImage struct {
	Status          string      `json:"status"`
	SourceServers   []string    `json:"sourceServers"`
	Description     string      `json:"description"`
	TargetVmOsImage TbImageInfo `json:"targetVmOsImage"`
	// Count            int                   `json:"count"`
	// TargetVmOsImages []tbmodel.TbImageInfo `json:"targetVmOsImages"`
}

// RecommendedVmOsImageList represents a collection of recommended VM OS images across multiple source servers.
type RecommendedVmOsImageList struct {
	Status                   string                 `json:"status"`
	Description              string                 `json:"description"`
	Count                    int                    `json:"count"`
	RecommendedVmOsImageList []RecommendedVmOsImage `json:"recommendedVmOsImageList"`
}
