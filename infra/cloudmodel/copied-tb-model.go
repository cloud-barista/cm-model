package cloudmodel

// * To avoid circular dependencies, the following structs are copied from the cb-tumblebug framework.
// TODO: When the cb-tumblebug framework is updated, we should synchronize these structs.
// * Version: CB-Tumblebug v0.10.8

// * Path: src/core/model/mci.go, Line: 89-109
// TbMciReq is struct for requirements to create MCI
type TbMciReq struct {
	Name string `json:"name" validate:"required" example:"mci01"`

	// InstallMonAgent Option for CB-Dragonfly agent installation ([yes/no] default:yes)
	InstallMonAgent string `json:"installMonAgent" example:"no" default:"no" enums:"yes,no"` // yes or no

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	// SystemLabel is for describing the mci in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel" example:"" default:""`

	PlacementAlgo string `json:"placementAlgo,omitempty"`
	Description   string `json:"description" example:"Made in CB-TB"`

	Vm []TbVmReq `json:"vm" validate:"required"`

	// PostCommand is for the command to bootstrap the VMs
	PostCommand MciCmdReq `json:"postCommand" validate:"omitempty"`
}

// * Path: src/core/model/mci.go, Line: 165-194
// TbVmReq is struct to get requirements to create a new server instance
type TbVmReq struct {
	// VM name or subGroup name if is (not empty) && (> 0). If it is a group, actual VM name will be generated with -N postfix.
	Name string `json:"name" validate:"required" example:"g1-1"`

	// CspResourceId is resource identifier managed by CSP (required for option=register)
	CspResourceId string `json:"cspResourceId,omitempty" example:"i-014fa6ede6ada0b2c"`

	// if subGroupSize is (not empty) && (> 0), subGroup will be generated. VMs will be created accordingly.
	SubGroupSize string `json:"subGroupSize" example:"3" default:""`

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	Description string `json:"description" example:"Description"`

	ConnectionName string `json:"connectionName" validate:"required" example:"testcloud01-seoul"`
	SpecId         string `json:"specId" validate:"required"`
	// ImageType        string   `json:"imageType"`
	ImageId          string   `json:"imageId" validate:"required"`
	VNetId           string   `json:"vNetId" validate:"required"`
	SubnetId         string   `json:"subnetId" validate:"required"`
	SecurityGroupIds []string `json:"securityGroupIds" validate:"required"`
	SshKeyId         string   `json:"sshKeyId" validate:"required"`
	VmUserName       string   `json:"vmUserName,omitempty"`
	VmUserPassword   string   `json:"vmUserPassword,omitempty"`
	RootDiskType     string   `json:"rootDiskType,omitempty" example:"default, TYPE1, ..."`  // "", "default", "TYPE1", AWS: ["standard", "gp2", "gp3"], Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"], GCP: ["pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"], ALIBABA: ["cloud_efficiency", "cloud", "cloud_ssd"], TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
	RootDiskSize     string   `json:"rootDiskSize,omitempty" example:"default, 30, 42, ..."` // "default", Integer (GB): ["50", ..., "1000"]
	DataDiskIds      []string `json:"dataDiskIds"`
}

// * Path: src/core/model/mci.go, Line: 204-223
// TbMciDynamicReq is struct for requirements to create MCI dynamically (with default resource option)
type TbMciDynamicReq struct {
	Name string `json:"name" validate:"required" example:"mci01"`

	// InstallMonAgent Option for CB-Dragonfly agent installation ([yes/no] default:no)
	InstallMonAgent string `json:"installMonAgent" example:"no" default:"no" enums:"yes,no"` // yes or no

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	// SystemLabel is for describing the mci in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel" example:"" default:""`

	Description string `json:"description" example:"Made in CB-TB"`

	Vm []TbVmDynamicReq `json:"vm" validate:"required"`

	// PostCommand is for the command to bootstrap the VMs
	PostCommand MciCmdReq `json:"postCommand"`
}

// * Path: src/core/model/mci.go, Line: 225-250
// TbVmDynamicReq is struct to get requirements to create a new server instance dynamically (with default resource option)
type TbVmDynamicReq struct {
	// VM name or subGroup name if is (not empty) && (> 0). If it is a group, actual VM name will be generated with -N postfix.
	Name string `json:"name" example:"g1-1"`

	// if subGroupSize is (not empty) && (> 0), subGroup will be generated. VMs will be created accordingly.
	SubGroupSize string `json:"subGroupSize" example:"3" default:"1"`

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	Description string `json:"description" example:"Description"`

	// CommonSpec is field for id of a spec in common namespace
	CommonSpec string `json:"commonSpec" validate:"required" example:"aws+ap-northeast-2+t2.small"`
	// CommonImage is field for id of a image in common namespace
	CommonImage string `json:"commonImage" validate:"required" example:"ubuntu18.04"`

	RootDiskType string `json:"rootDiskType,omitempty" example:"default, TYPE1, ..." default:"default"`  // "", "default", "TYPE1", AWS: ["standard", "gp2", "gp3"], Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"], GCP: ["pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"], ALIBABA: ["cloud_efficiency", "cloud", "cloud_essd"], TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
	RootDiskSize string `json:"rootDiskSize,omitempty" example:"default, 30, 42, ..." default:"default"` // "default", Integer (GB): ["50", ..., "1000"]

	VmUserPassword string `json:"vmUserPassword,omitempty" default:""`
	// if ConnectionName is given, the VM tries to use associtated credential.
	// if not, it will use predefined ConnectionName in Spec objects
	ConnectionName string `json:"connectionName,omitempty" default:""`
}

// * Path: src/core/model/mci.go, Line: 703-707
// MciCmdReq is struct for remote command
type MciCmdReq struct {
	UserName string   `json:"userName" example:"cb-user" default:""`
	Command  []string `json:"command" validate:"required" example:"client_ip=$(echo $SSH_CLIENT | awk '{print $1}'); echo SSH client IP is: $client_ip"`
}

// * Path: src/core/model/vnet.go, Line: 17-26
// TbVNetReq is a struct to handle 'Create vNet' request toward CB-Tumblebug.
type TbVNetReq struct { // Tumblebug
	Name           string        `json:"name" validate:"required" example:"vnet00"`
	ConnectionName string        `json:"connectionName" validate:"required" example:"aws-ap-northeast-2"`
	CidrBlock      string        `json:"cidrBlock" example:"10.0.0.0/16"`
	SubnetInfoList []TbSubnetReq `json:"subnetInfoList"`
	Description    string        `json:"description" example:"vnet00 managed by CB-Tumblebug"`
	// todo: restore the tag list later
	// TagList        []KeyValue    `json:"tagList,omitempty"`
}

// * Path: src/core/model/subnet.go, Line: 17-25
// TbSubnetReq is a struct that represents TB subnet object.
type TbSubnetReq struct { // Tumblebug
	Name        string `json:"name" validate:"required" example:"subnet00"`
	IPv4_CIDR   string `json:"ipv4_CIDR" validate:"required" example:"10.0.1.0/24"`
	Zone        string `json:"zone,omitempty" default:""`
	Description string `json:"description,omitempty" example:"subnet00 managed by CB-Tumblebug"`
	// todo: restore the tag list later
	// TagList     []KeyValue `json:"tagList,omitempty"`
}

// * Path: src/core/model/sshkey.go, Line: 38-52
// TbSshKeyReq is a struct to handle 'Create SSH key' request toward CB-Tumblebug.
type TbSshKeyReq struct {
	Name           string `json:"name" validate:"required"`
	ConnectionName string `json:"connectionName" validate:"required"`
	Description    string `json:"description"`

	// Fields for "Register existing SSH keys" feature
	// CspResourceId is required to register object from CSP (option=register)
	CspResourceId    string `json:"cspResourceId"`
	Fingerprint      string `json:"fingerprint"`
	Username         string `json:"username"`
	VerifiedUsername string `json:"verifiedUsername"`
	PublicKey        string `json:"publicKey"`
	PrivateKey       string `json:"privateKey"`
}

// * Path: src/core/model/spec.go, Line: 106-157
// TbSpecInfo is a struct that represents TB spec object.
type TbSpecInfo struct { // Tumblebug
	// Id is unique identifier for the object
	Id string `json:"id" example:"aws-ap-southeast-1" gorm:"primaryKey"`
	// Uid is universally unique identifier for the object, used for labelSelector
	Uid string `json:"uid,omitempty" example:"wef12awefadf1221edcf"`

	// CspSpecName is name of the spec given by CSP
	CspSpecName string `json:"cspSpecName,omitempty" example:"csp-06eb41e14121c550a"`

	// Name is human-readable string to represent the object
	Name           string `json:"name" example:"aws-ap-southeast-1"`
	Namespace      string `json:"namespace,omitempty" example:"default" gorm:"primaryKey"`
	ConnectionName string `json:"connectionName,omitempty"`
	ProviderName   string `json:"providerName,omitempty"`
	RegionName     string `json:"regionName,omitempty"`
	// InfraType can be one of vm|k8s|kubernetes|container, etc.
	InfraType             string   `json:"infraType,omitempty"`
	Architecture          string   `json:"architecture,omitempty" example:"x86_64"`
	OsType                string   `json:"osType,omitempty"`
	VCPU                  uint16   `json:"vCPU,omitempty"`
	MemoryGiB             float32  `json:"memoryGiB,omitempty"`
	DiskSizeGB            float32  `json:"diskSizeGB,omitempty"`
	MaxTotalStorageTiB    uint16   `json:"maxTotalStorageTiB,omitempty"`
	NetBwGbps             uint16   `json:"netBwGbps,omitempty"`
	AcceleratorModel      string   `json:"acceleratorModel,omitempty"`
	AcceleratorCount      uint8    `json:"acceleratorCount,omitempty"`
	AcceleratorMemoryGB   float32  `json:"acceleratorMemoryGB,omitempty"`
	AcceleratorType       string   `json:"acceleratorType,omitempty"`
	CostPerHour           float32  `json:"costPerHour,omitempty"`
	Description           string   `json:"description,omitempty"`
	OrderInFilteredResult uint16   `json:"orderInFilteredResult,omitempty"`
	EvaluationStatus      string   `json:"evaluationStatus,omitempty"`
	EvaluationScore01     float32  `json:"evaluationScore01"`
	EvaluationScore02     float32  `json:"evaluationScore02"`
	EvaluationScore03     float32  `json:"evaluationScore03"`
	EvaluationScore04     float32  `json:"evaluationScore04"`
	EvaluationScore05     float32  `json:"evaluationScore05"`
	EvaluationScore06     float32  `json:"evaluationScore06"`
	EvaluationScore07     float32  `json:"evaluationScore07"`
	EvaluationScore08     float32  `json:"evaluationScore08"`
	EvaluationScore09     float32  `json:"evaluationScore09"`
	EvaluationScore10     float32  `json:"evaluationScore10"`
	RootDiskType          string   `json:"rootDiskType"`
	RootDiskSize          string   `json:"rootDiskSize"`
	AssociatedObjectList  []string `json:"associatedObjectList,omitempty" gorm:"type:text;serializer:json"`
	IsAutoGenerated       bool     `json:"isAutoGenerated,omitempty"`

	// SystemLabel is for describing the Resource in a keyword (any string can be used) for special System purpose
	SystemLabel string     `json:"systemLabel,omitempty" example:"Managed by CB-Tumblebug" default:""`
	Details     []KeyValue `json:"details" gorm:"type:text;serializer:json"`
}

// * Path: src/core/model/common.go, Line: 29-33
// KeyValue is struct for key-value pair
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// * Path: src/core/model/image.go, Line: 79-114
// TbImageInfo is a struct that represents TB image object.
type TbImageInfo struct {
	// Composite primary key
	Namespace    string `json:"namespace" example:"default" gorm:"primaryKey"`
	ProviderName string `json:"providerName" gorm:"primaryKey"`
	CspImageName string `json:"cspImageName" example:"csp-06eb41e14121c550a" gorm:"primaryKey" description:"The name of the CSP image for querying image information."`

	// Array field for supporting multiple regions
	RegionList []string `json:"regionList" gorm:"type:text;serializer:json"`

	Id  string `json:"id" example:"aws-ap-southeast-1"`
	Uid string `json:"uid,omitempty" example:"wef12awefadf1221edcf"`

	Name           string `json:"name" example:"aws-ap-southeast-1"`
	ConnectionName string `json:"connectionName,omitempty"`
	InfraType      string `json:"infraType,omitempty"` // vm|k8s|kubernetes|container, etc.

	FetchedTime  string `json:"fetchedTime,omitempty"`
	CreationDate string `json:"creationDate,omitempty"`

	IsGPUImage        bool `json:"isGPUImage,omitempty" gorm:"column:is_gpu_image" enum:"true|false" default:"false" description:"Whether the image is GPU-enabled or not."`
	IsKubernetesImage bool `json:"isKubernetesImage,omitempty" gorm:"column:is_kubernetes_image" enum:"true|false" default:"false" description:"Whether the image is Kubernetes-enabled or not."`

	OSType string `json:"osType,omitempty" gorm:"column:os_type" example:"ubuntu 22.04" description:"Simplified OS name and version string"`

	OSArchitecture OSArchitecture `json:"osArchitecture" gorm:"column:os_architecture" example:"x86_64" description:"The architecture of the operating system of the image."`        // arm64, x86_64 etc.
	OSPlatform     OSPlatform     `json:"osPlatform" gorm:"column:os_platform" example:"Linux/UNIX" description:"The platform of the operating system of the image."`                // Linux/UNIX, Windows, NA
	OSDistribution string         `json:"osDistribution" gorm:"column:os_distribution" example:"Ubuntu 22.04~" description:"The distribution of the operating system of the image."` // Ubuntu 22.04~, CentOS 8 etc.
	OSDiskType     string         `json:"osDiskType" gorm:"column:os_disk_type" example:"HDD" description:"The type of the OS disk of for the VM being created."`                    // ebs, HDD, etc.
	OSDiskSizeGB   float64        `json:"osDiskSizeGB" gorm:"column:os_disk_size_gb" example:"50" description:"The (minimum) OS disk size in GB for the VM being created."`          // 10, 50, 100 etc.
	ImageStatus    ImageStatus    `json:"imageStatus" example:"Available" description:"The status of the image, e.g., Available, Deprecated, NA."`                                   // Available, Deprecated, NA

	Details     []KeyValue `json:"details" gorm:"type:text;serializer:json"`
	SystemLabel string     `json:"systemLabel,omitempty" example:"Managed by CB-Tumblebug" default:""`
	Description string     `json:"description,omitempty"`
}

// * Path: src/core/model/image.go, Line: 23
type OSArchitecture string

// * Path: src/core/model/image.go, Line: 38
type OSPlatform string

// * Path: src/core/model/image.go, Line: 46
type ImageStatus string

// * Path: src/core/model/securitygroup.go, Line: 65-75
// TbSecurityGroupReq is a struct to handle 'Create security group' request toward CB-Tumblebug.
type TbSecurityGroupReq struct { // Tumblebug
	Name           string                `json:"name" validate:"required"`
	ConnectionName string                `json:"connectionName" validate:"required"`
	VNetId         string                `json:"vNetId" validate:"required"`
	Description    string                `json:"description"`
	FirewallRules  *[]TbFirewallRuleInfo `json:"firewallRules"` // validate:"required"`

	// CspResourceId is required to register object from CSP (option=register)
	CspResourceId string `json:"cspResourceId"`
}

// * Path: src/core/model/securitygroup.go, Line: 77-84
// TbFirewallRuleInfo is a struct to handle firewall rule info of CB-Tumblebug.
type TbFirewallRuleInfo struct {
	FromPort   string `validate:"required"` //`json:"fromPort"`
	ToPort     string `validate:"required"` //`json:"toPort"`
	IPProtocol string `validate:"required"` //`json:"ipProtocol"`
	Direction  string `validate:"required"` //`json:"direction"`
	CIDR       string
}
