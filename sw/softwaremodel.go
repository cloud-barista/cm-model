package Softwaremodel

type SoftwareModel struct {
	SoftwareModel Software `json:"softwareModel" validate:"required"`
}

type Software struct {
	SoftwareProperties 	[]SoftwareProperty 	`json:"software_properties"`
	Environment  		Environment        	`json:"environment"`  	// can be on-premise or cloud.
}

type SoftwareProperty struct {
	SoftwareName    	string            	`json:"software_name"`    // any software, e.g., "nginx", "nfs", "mysql"
	Version 			string            	`json:"version"` // version of the software
	Config          	SoftwareConfig    	`json:"config"`           // configuration settings for the software

	// Optional parameters for container and Kubernetes migration methods.
	ContainerInfo  		*ContainerConfig  	`json:"container_info,omitempty"`
	KubernetesInfo 		*KubernetesConfig 	`json:"kubernetes_info,omitempty"`
	
	Methods       		[]MigrationMethod 	`json:"methods"` // supported migration methods
}

type Environment struct { // Computing environment
	Type   		string            	`json:"type"` 		// e.g., "on-premise" or "cloud"
	Provider 	string 			  	`json:"provider,omitempty"` // (if cloud) aws, azure, gcp, etc. 
	Region   	string 				`json:"region,omitempty"`
	Zone     	string 				`json:"zone,omitempty"`
	OS       	OperatingSystem 	`json:"os"`
	Config 		map[string]string 	`json:"config"`		// can include additional details such as network settings, credentials, etc.
}

type OperatingSystem struct {
	Type         string `json:"type"` 	 // Ubuntu, Debian, Windows
	Version      string `json:"version"` // OS version
	Architecture string `json:"architecture,omitempty"` // x86_64, arm64
}

type MigrationMethod string

const (
	MethodBinary     MigrationMethod = "binary"      // Moving the software as a binary executable.
	MethodOSPackage  MigrationMethod = "os_package"  // Installing via OS package manager.
	MethodContainer  MigrationMethod = "container"   // Installing as a container package.
	MethodKubernetes MigrationMethod = "kubernetes"  // Installing as a Kubernetes package.
	MethodManaged    MigrationMethod = "managed_csp" // Using the CSP's managed software service.
)

type SoftwareConfig struct {
	// Settings can include environment variables, config file paths, ports, etc.
	Settings 	map[string]string `json:"settings"`
}

type ContainerConfig struct {
	Network     map[string]string            `json:"network"` // e.g., bridge, host, overlay, etc.
	Volume      map[string]string            `json:"volume"`  // volume details or mount point info.
	Image       map[string]string            `json:"image"`   // container image info.
	ExtraConfig map[string]string 			 `json:"extra_config,omitempty"` // Additional container configuration.
}

type KubernetesConfig struct {
	Cluster     map[string]string            `json:"cluster"`  // cluster name or identifier.
	Network     map[string]string            `json:"network"`  // network settings or CNI configurations.
	Volume      map[string]string            `json:"volume"`   // volume settings.
	Policies    map[string]string            `json:"policies"` // any relevant policies.
	ExtraConfig map[string]string 			 `json:"extra_config,omitempty"` // Additional Kubernetes configuration.
}

// // MigrationTask represents a complete migration process.
// // It includes both the source and target environments and a set of software properties.
// type MigrationTask struct {
// 	TaskID             string             `json:"task_id"`
// 	SourceEnvironment  Environment        `json:"source"`  // can be on-premise or cloud.
// 	TargetEnvironment  Environment        `json:"target"`  // always a cloud environment.
// 	SoftwareProperties []SoftwareProperty `json:"software_properties"`
// }
