package cloudinframodel

import (
	software    "github.com/cloud-barista/cm-model/sw"
)

type CloudInfraModel struct {
	OnpremiseInfraModel CloudInfra `json:"cloudInfraModel" validate:"required"`
}

type CloudInfra struct {
	CloudEnv 			CloudEnvironment  `json:"cloudenvironment" validate:"required"`
	Network 			NetworkProperty   `json:"network" validate:"required"`
	VMs 				[]VMProperty 	  `json:"vms" validate:"required"`
	Clusters    		[]ClusterProperty `json:"clusters,omitempty"`
	SoftwareProperties 	[]software.SoftwareProperty 	`json:"software_properties"` // For CSP's managed software service.
}

type CSPType string

const (
	AWS       CSPType = "AWS"
	GCP       CSPType = "GCP"
	Azure     CSPType = "Azure"
	Alibaba   CSPType = "Alibaba"
	Tencent   CSPType = "Tencent"
	OpenStack CSPType = "OpenStack"
	IBM       CSPType = "IBM"
	NCP    	  CSPType = "NCP"
	NHN    	  CSPType = "NHN"
	KTCloud   CSPType = "KTCloud"
)

type CloudEnvironment struct {
	CSPType    CSPType `json:"cspType"` // AWS, GCP, Azure, ...
	RegionName string  `json:"regionName"`
	ZoneName   string  `json:"zoneName,omitempty"`
}

type NetworkProperty struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	CIDR        string   `json:"cidr,omitempty"`
	VPCs        []VPC    `json:"vpcs,omitempty"`
	Label       map[string]string `json:"lavel,omitempty"`
}

type VPC struct {
	Name           string    		`json:"name"`
	Description    string    		`json:"description,omitempty"`
	CIDR           string    		`json:"cidr"`
	Subnets        []Subnet  		`json:"subnets,omitempty"`
	SecurityGroups []SecurityGroup 	`json:"securityGroups,omitempty"`
	RouteTable     RouteInfo 		`json:"routeTable,omitempty"`
	Label          map[string]string `json:"lavel,omitempty"`
}

type Subnet struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	CIDR        string   `json:"cidr"`
	Zone        string   `json:"zone,omitempty"`
	IsPublic    bool     `json:"isPublic"`
	Label       map[string]string `json:"lavel,omitempty"`
}

type RouteInfo struct {
	Name         string       `json:"name"`
	Description  string       `json:"description,omitempty"`
	Routes       []RouteEntry `json:"routes,omitempty"`
	AssociatedTo []string     `json:"associatedTo,omitempty"` // IDs of associated resources
}

type RouteEntry struct {
	DestinationCIDR string `json:"destinationCidr"`
	Target          string `json:"target"` // Gateway ID, Instance ID, etc.
	TargetType      string `json:"targetType"`
}

type VMProperty struct {
	Name            	string       	`json:"name"`
	Description     	string       	`json:"description,omitempty"`
	Status          	string       	`json:"status,omitempty"` // Running or Susppended
	VPC             	VPC       		`json:"vpc"`              // VPC Info
	Subnet          	Subnet       	`json:"subnet"`           // Subnet Info
	SecurityGroups  	[]SecurityGroup `json:"securityGroups"`   // Security Group Info.s
	VMImage       		VMImage       	`json:"vmImage"`        // VM Image Info
	VMSpec        		VMSpec       	`json:"vmSpec"`         // VM Spec Info
	KeyPair       		KeyPair       	`json:"keyPair"`        // Key Pair Info
	PrivateIP       	string       	`json:"privateIp,omitempty"`
	RootDisk			Disk         	`json:"rootDisk,omitempty"`
	AdditionalDisks 	[]Disk       	`json:"additionalDisks,omitempty"`
	NetworkInterfaces 	[]NIC        	`json:"networkInterfaces,omitempty"`
	UserData        	string       	`json:"userData,omitempty"` // Boot script
	AuthInfo			AuthInfo 	 	`json:"authinfo,omitempty"`
	InstallMonAgent		bool 	 	 	`json:"monAgent"`	// "installMonAgent": true or false
	Label       		map[string]string `json:"lavel,omitempty"`
}

type SecurityGroup struct {
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Rules       []SGRule  	  `json:"rules,omitempty"`
	Label       map[string]string `json:"lavel,omitempty"`
}

type SGRule struct {
	Name          string `json:"name,omitempty"`
	Direction     string `json:"direction"`     // "inbound" or "outbound"
	IPProtocol    string `json:"ipProtocol"`    // e.g., "tcp", "udp", "icmp"
	FromPort      int    `json:"fromPort"`      // Start port range
	ToPort        int    `json:"toPort"`        // End port range
	SourceCIDR    string `json:"sourceCidr,omitempty"`
	SourceSecGrp  string `json:"sourceSecGrp,omitempty"`
	DestCIDR      string `json:"destCidr,omitempty"`
	DestSecGrp    string `json:"destSecGrp,omitempty"`
	Description   string `json:"description,omitempty"`
}

type VMImage struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description,omitempty"`
	OSType        string   `json:"osType"`        // e.g., "Linux", "Windows"
	OSDistro      string   `json:"osDistro"`      // e.g., "Ubuntu", "Windows Server"
	OSVersion     string   `json:"osVersion"`     // e.g., "22.04", "2019"
	Architecture  string   `json:"architecture,omitempty"`  // e.g., "x86_64", "arm64"
	SizeGB        int      `json:"sizeGB,omitempty"`
	SourceURL     string   `json:"sourceUrl,omitempty"`
	SourceImageID string   `json:"sourceImageId,omitempty"`
	Label         map[string]string `json:"lavel,omitempty"`
}

type VMSpec struct {
	ID           	string    `json:"id"`
	Name         	string    `json:"name"`
	Description  	string    `json:"description,omitempty"`
	VCPU         	int       `json:"vCpu"`
	MemoryGB     	float64   `json:"memoryGb"`
	GPUType     	string    `json:"gpuType,omitempty"`
	GPUCount     	int       `json:"gpuCount,omitempty"`
	DiskSizeGB   	int       `json:"diskSizeGb,omitempty"`
	DiskType     	string    `json:"diskType,omitempty"`   // e.g., "SSD", "HDD"
	NetworkBandwidth int    `json:"networkBandwidth,omitempty"` // In Mbps
	CSPInstanceType string  `json:"cspInstanceType,omitempty"`  // e.g., "t2.micro", "n1-standard-1"
	Label       	map[string]string `json:"lavel,omitempty"`
}

type KeyPair struct {
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	PublicKey      string `json:"publicKey,omitempty"`
	PrivateKey     string `json:"privateKey,omitempty"` // Note: Should be handled securely
	Fingerprint    string `json:"fingerprint,omitempty"`
	KeyFormat      string `json:"keyFormat,omitempty"` // e.g., "ssh-rsa", "ecdsa"
}

type Disk struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	SizeGB      int      `json:"sizeGb"`
	DiskType    string   `json:"diskType"`   // e.g., "SSD", "HDD"
	IOPS        int      `json:"iops,omitempty"`
	Throughput  int      `json:"throughput,omitempty"` // In MBps
	Zone        string   `json:"zone,omitempty"`
	Snapshot    string   `json:"snapshot,omitempty"`   // Source snapshot ID if applicable
	AttachedTo  string   `json:"attachedTo,omitempty"` // VM ID if attached
	DeviceName  string   `json:"deviceName,omitempty"` // e.g., "/dev/sda1"
	IsBootDisk  bool     `json:"isBootDisk"`
	IsEncrypted bool     `json:"isEncrypted"`
	Label       map[string]string `json:"lavel,omitempty"`
}

type NIC struct { // network interface card
	Name          string   `json:"name"`
	Description   string   `json:"description,omitempty"`
	SubnetID      string   `json:"subnetId"`
	PrivateIP     string   `json:"privateIp,omitempty"`
	PublicIP      string   `json:"publicIp,omitempty"`
	SecurityGroups []string `json:"securityGroups,omitempty"`
	IsPrimary     bool     `json:"isPrimary"`
	MacAddress    string   `json:"macAddress,omitempty"`
}

type AuthInfo struct {
	Type          string `json:"type"` // ssh-key, password, api-key
	Username      string `json:"username"`
	Password      string `json:"password,omitempty"`
	KeyPath       string `json:"keyPath,omitempty"`
	APICredential string `json:"apiCredential,omitempty"`
}

type ClusterProperty struct { // container orchestration cluster (e.g., Kubernetes)
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Description      string       `json:"description,omitempty"`
	Type             string       `json:"type"`              // e.g., "Kubernetes", "ECS", "GKE"
	Version          string       `json:"version,omitempty"` // e.g., "1.23"
	VPC              string       `json:"vpc"`               // VPC ID
	Subnets          []string     `json:"subnets"`           // Subnet IDs
	NodePools        []NodePool   `json:"nodePools,omitempty"`
	EndpointURL      string       `json:"endpointUrl,omitempty"`
	CAData           string       `json:"caData,omitempty"`  // Certificate Authority data
	KubeConfig       string       `json:"kubeConfig,omitempty"`
	Label       	 map[string]string `json:"lavel,omitempty"`
	AddOns           []AddOn      `json:"addOns,omitempty"`
}

type NodePool struct { // group of worker nodes in a cluster
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description,omitempty"`
	VMSpec       string   `json:"vmSpec"`       // VM Spec ID
	VMImage      string   `json:"vmImage"`      // VM Image ID
	AutoScaling  bool     `json:"autoScaling"`
	MinNodes     int      `json:"minNodes,omitempty"`
	MaxNodes     int      `json:"maxNodes,omitempty"`
	DesiredNodes int      `json:"desiredNodes"`
	NodeLabels   []KeyValue `json:"nodeLabels,omitempty"`
	Taints       []Taint  `json:"taints,omitempty"`
	Subnets      []string `json:"subnets,omitempty"` // Subnet IDs
	DiskSizeGB   int      `json:"diskSizeGb,omitempty"`
	DiskType     string   `json:"diskType,omitempty"`
}

type AddOn struct { // add-on component for a cluster
	Name        string    `json:"name"`
	Version     string    `json:"version,omitempty"`
	Enabled     bool      `json:"enabled"`
	Config      []KeyValue `json:"config,omitempty"`
}

type Taint struct { // Kubernetes node taint
	Key    string `json:"key"`
	Value  string `json:"value,omitempty"`
	Effect string `json:"effect"` // e.g., "NoSchedule", "PreferNoSchedule", "NoExecute"
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
