package Softwaremodel

type SoftwareModel struct {
	SoftwareModel Software `json:"softwareModel" validate:"required"`
}

type Software struct {
	SoftwareProperties 	[]SoftwareProperty 	`json:"software_properties"`
	Environment  		Environment        	`json:"environment"`  	// can be on-premise or cloud.
}

type SoftwareProperty struct {
	SoftwareName    	string            	`json:"software_name"`  // any software, e.g., "nginx", "nfs", "mysql"
	Version 			string            	`json:"version"` 		// version of the software
	Config          	SoftwareConfig    	`json:"config"`         // configuration settings for the software
	Methods       		MigrationMethod 	`json:"methods"` // supported migration methods
}

type Environment struct { // Computing environment
	Type   		string            	`json:"type"` 	// e.g., "on-premise" or "cloud"
	OS       	OperatingSystem 	`json:"os"`
	Config 		map[string]string 	`json:"config"`	// can include additional details such as network settings, credentials, etc.
}

type OperatingSystem struct {
	Type         string `json:"type"` 	 // Ubuntu, Debian, Windows
	Version      string `json:"version"` // OS version
	Architecture string `json:"architecture,omitempty"` // x86_64, arm64
}

type MigrationMethod string

const (
	MethodOSPackage  MigrationMethod = "os_package"  // Installing via OS package manager.
	MethodContainer  MigrationMethod = "container"   // Installing as a container package.
	MethodKubernetes MigrationMethod = "kubernetes"  // Installing as a Kubernetes package.
	MethodBinary     MigrationMethod = "binary"      // Moving the software as a binary executable.
)

type SoftwareConfig struct {
	ConfigFiles 	[]string          `json:"config_files"`  // List of configuration file paths
	ContentFiles 	[]string          `json:"content_files"` // List of additional content file paths
    LogFiles     	[]string          `json:"log_files"`     // List of log file paths
	ImageFiles     	[]string          `json:"image_files"`   // List of image file paths or binary file paths
	K8sFiles     	[]string          `json:"k8s_files"`     // List of files related to the kubernetes environment	
	Settings 		map[string]string `json:"settings"` // Can include environment variables, config file paths, ports, etc. (ex. Key/Value type)
}
