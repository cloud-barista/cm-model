package softwaremodel

type SoftwareModel struct {
	SoftwareModel Software `json:"softwareModel" validate:"required"`
}

type Software struct {
	HostEnvs                []Env                    `json:"host_envs"`
	BinaryInfos             []BinaryInfo             `json:"binaryInfo"`
	PackageInfos            []PackageInfo            `json:"package_infos"`
	ContainerInfos          []ContainerInfo          `json:"container_infos"`
	KubernetesInfo          KubernetesInfo           `json:"kubernetes_infos"`
	KubernetesResourceInfos []KubernetesResourceInfo `json:"kubernetes_resource_infos"`
	KubernetesHelmInfos     []KubernetesHelmInfo     `json:"kubernetes_helm_infos"`
}

type BinaryInfo struct {
	Index             SoftwareIndex        `json:"index,omitempty" validate:"required"`
	IndexDepends      []SoftwareIndex      `json:"index_depends"` // Migration dependencies (Migrations will be processed first in this list.)
	BinaryPath        string               `json:"binary_path,omitempty" validate:"required"`
	Version           string               `json:"version"`
	Architecture      SoftwareArchitecture `json:"architecture,omitempty" validate:"required"`
	IsStatic          bool                 `json:"is_static" validate:"required"`
	LibraryPaths      []string             `json:"library_paths"`
	CustomConfigPaths []string             `json:"custom_config_paths"`
	ConnectionInfo    ConnectionInfo       `json:"connection_info"` // Connection information if needed for software migration
	ServiceInfo       Service              `json:"service_info"`    // Provide service information if the binary run by the service
}

type PackageInfo struct {
	Index                SoftwareIndex        `json:"index,omitempty" validate:"required"`
	IndexDepends         []SoftwareIndex      `json:"index_depends"`                             // Migration dependencies (Migrations will be processed first in this list.)
	SoftwarePackageType  SoftwarePackageType  `json:"software_package_type" validate:"required"` // deb / rpm
	Name                 string               `json:"name,omitempty" validate:"required"`
	Version              string               `json:"version,omitempty" validate:"required"`
	OS                   string               `json:"os,omitempty"`
	OSVersion            string               `json:"os_version,omitempty"`
	Architecture         SoftwareArchitecture `json:"architecture,omitempty"`
	NeededPackages       []string             `json:"needed_packages"`         // Packages need to install with this package
	NeedToDeletePackages []string             `json:"need_to_delete_packages"` // Packages need to delete before installation
	CustomConfigPaths    []string             `json:"custom_config_paths"`     // Need to copy config paths (ex: /etc/exports for NFS Server)
	RepoURL              string               `json:"repo_url"`
	GPGKeyURL            string               `json:"gpg_key_url"`
	RepoUseOSVersionCode bool                 `json:"repo_use_os_version_code" default:"false"`
	ConnectionInfo       ConnectionInfo       `json:"connection_info"` // Connection information if needed for software migration
	ServiceInfo          Service              `json:"service_info"`    // Provide service information if the package uses service
}

type ContainerInfo struct {
	Index             SoftwareIndex   `json:"index,omitempty" validate:"required"`
	IndexDepends      []SoftwareIndex `json:"index_depends"`                         // Migration dependencies (Migrations will be processed first in this list.)
	Runtime           string          `json:"runtime,omitempty" validate:"required"` // Which runtime uses for the container (Docker, Podman, ...)
	ContainerName     string          `json:"container_name,omitempty" validate:"required"`
	ContainerImage    ContainerImage  `json:"container_image,omitempty" validate:"required"`
	ContainerPorts    []ContainerPort `json:"container_ports"`
	ContainerStatus   string          `json:"container_status" validate:"required"`
	DockerComposePath string          `json:"docker_compose_path"`
	MountPaths        []string        `json:"mount_paths"`
	Envs              []Env           `json:"envs"`
	NetworkMode       string          `json:"network_mode,omitempty" validate:"required"`
	RestartPolicy     string          `json:"restart_policy,omitempty" validate:"required"`
	ConnectionInfo    ConnectionInfo  `json:"connection_info"` // Connection information if needed for software migration
	ServiceInfo       Service         `json:"service_info"`    // Provide service information if the container run by the service
}

type KubernetesInfo struct {
	KubernetesVersion          string `json:"kubernetes_version,omitempty" validate:"required"`
	KubeConfigYAML             string `json:"kube_config_yaml,omitempty" validate:"required"`
	KubernetesContainerRuntime string `json:"kubernetes_container_runtime,omitempty" validate:"required"`
}

type KubernetesResourceInfo struct {
	Index        SoftwareIndex   `json:"index,omitempty" validate:"required"`
	IndexDepends []SoftwareIndex `json:"index_depends"` // Migration dependencies (Migrations will be processed first in this list.)
	Namespace    string          `json:"namespace" validate:"required"`
	Kind         string          `json:"kind" validate:"required"`
	Name         string          `json:"name" validate:"required"`
}

type KubernetesHelmInfo struct {
	Index          SoftwareIndex   `json:"index,omitempty" validate:"required"`
	IndexDepends   []SoftwareIndex `json:"index_depends"` // Migration dependencies (Migrations will be processed first in this list.)
	RepoURL        string          `json:"repo_url" validate:"required"`
	Release        string          `json:"release" validate:"required"`
	HelmChartPath  string          `json:"helm_chart_path" validate:"required"`
	HelmValuesYAML string          `json:"helm_values_yaml"`
}

type SoftwareType string

const (
	SoftwareTypePackage    SoftwareType = "package"    // Installing via OS package manager.
	SoftwareTypeContainer  SoftwareType = "container"  // Installing as a container package.
	SoftwareTypeKubernetes SoftwareType = "kubernetes" // Installing as a Kubernetes package.
	SoftwareTypeBinary     SoftwareType = "binary"     // Moving the software as a binary executable.
)

type SoftwareIndex uint64 // Each software has its own index

type SoftwareArchitecture string

const (
	SoftwareArchitectureCommon SoftwareArchitecture = "common"
	SoftwareArchitectureX8664  SoftwareArchitecture = "x86_64"
	SoftwareArchitectureX86    SoftwareArchitecture = "x86"
	SoftwareArchitectureARM    SoftwareArchitecture = "arm"
	SoftwareArchitectureARM64  SoftwareArchitecture = "arm64"
)

type SoftwarePackageType string

const (
	SoftwarePackageTypeDEB SoftwarePackageType = "deb"
	SoftwarePackageTypeRPM SoftwarePackageType = "rpm"
)

type ContainerImage struct {
	ImageName         string               `json:"image_name" validate:"required"`
	ImageVersion      string               `json:"image_version" validate:"required"`
	ImageArchitecture SoftwareArchitecture `json:"image_architecture" validate:"required"`
}

type ContainerPort struct {
	ContainerPort int    `json:"container_port" validate:"required"` // NetworkSettings.Ports.{Port}/{Protocol} -> {Port}
	Protocol      string `json:"protocol" validate:"required"`       // NetworkSettings.Ports.{Port}/{Protocol} -> {Protocol}
	HostIP        string `json:"host_ip" validate:"required"`        // NetworkSettings.Ports.{Port}/{Protocol}.HostIp
	HostPort      int    `json:"host_port" validate:"required"`      // NetworkSettings.Ports.{Port}/{Protocol}.HostPort
}

type Service struct {
	ServiceName         string    `json:"service_name,omitempty" validate:"required"`
	ServiceStatus       string    `json:"service_status,omitempty" validate:"required"`
	ServiceEnabled      bool      `json:"service_enabled,omitempty" validate:"required"`
	ServiceFilePath     string    `json:"service_file_path,omitempty" validate:"required"`
	ServiceUser         string    `json:"service_user"`  // User permission of the service
	ServiceGroup        string    `json:"service_group"` // Group permission of the service
	ServiceDependencies []Service `json:"service_dependencies"`
}

type ConnectionInfo struct {
	ListenPorts []Port `json:"listen_ports"`
	User        User   `json:"user"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Port struct {
	Port     int    `json:"container_port,omitempty" validate:"required"`
	Protocol string `json:"protocol,omitempty" validate:"required"`
	HostIP   string `json:"host_ip,omitempty" validate:"required"`
}

type Env struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Value string `json:"value,omitempty"`
}
