package Softwaremodel2

type SoftwareModel struct {
	SoftwareModel Software `json:"softwareModel" validate:"required"`
}

type Software struct {
	SoftwareProperties 	[]SoftwareProperty 	`json:"software_properties"`
	Environment  		Environment        	`json:"environment"`  	// can be on-premise or cloud.
}

// Software defines the software being migrated
type SoftwareProperty struct {
	SoftwareName 	string      `json:"software_name"`  // any software, e.g., "nginx", "nfs", "mysql"
	Version     	string      `json:"version"`
	InstallPath 	string      `json:"installPath"`
	DataPath    	string      `json:"dataPath"`
	Config      	ConfigData  `json:"config"`
	Dependencies 	[]Dependency `json:"dependencies"`

	Methods      	[]MigrationMethod 	`json:"methods"` // supported migration methods
}

type Dependency struct { // software dependency
	Name     string `json:"name"`
	Version  string `json:"version"`
	Required bool   `json:"required"`
}

type Environment struct {
	Type     string 	`json:"type"` // on-premise or cloud
	Provider string 	`json:"provider,omitempty"` // (if cloud) aws, azure, gcp, etc. 
	Region   string 	`json:"region,omitempty"`
	Zone     string 	`json:"zone,omitempty"`
	OS       OperatingSystem `json:"os"`	
}

type OperatingSystem struct {
	Type         string `json:"type"` 	 // Ubuntu, Debian, Windows
	Version      string `json:"version"` // OS version
	Architecture string `json:"architecture,omitempty"` // x86_64, arm64
}

type MigrationMethod struct {
	Type string `json:"type"` // binary, os-package, container, kubernetes, managed-service
	
	// One of the following will be populated based on the Type
	BinaryDetails        *BinaryMigration        `json:"binaryDetails,omitempty"`
	OSPackageDetails     *OSPackageMigration     `json:"osPackageDetails,omitempty"`
	ContainerDetails     *ContainerMigration     `json:"containerDetails,omitempty"`
	KubernetesDetails    *KubernetesMigration    `json:"kubernetesDetails,omitempty"`
	ManagedServiceDetails *ManagedServiceMigration `json:"managedServiceDetails,omitempty"`
}

type BinaryMigration struct { // details for binary executable migration
	BinaryPaths      []string `json:"binaryPaths"`
	LibraryPaths     []string `json:"libraryPaths"`
	DataPaths        []string `json:"dataPaths"`
	PreMigrationScript  string `json:"preMigrationScript,omitempty"`
	PostMigrationScript string `json:"postMigrationScript,omitempty"`
}

type OSPackageMigration struct { // details for OS package installation
	PackageName      string   `json:"packageName"`
	PackageVersion   string   `json:"packageVersion,omitempty"`
	Repository       string   `json:"repository,omitempty"`
	RepositoryKey    string   `json:"repositoryKey,omitempty"`
	InstallCommand   string   `json:"installCommand"`
	Dependencies     []string `json:"dependencies,omitempty"`
	PreInstallScript string   `json:"preInstallScript,omitempty"`
	PostInstallScript string  `json:"postInstallScript,omitempty"`
}

type ContainerMigration struct { // details for container deployment
	ImageName         string   `json:"imageName"`
	ImageTag          string   `json:"imageTag"`
	Registry          string   `json:"registry,omitempty"`
	DockerfilePath    string   `json:"dockerfilePath,omitempty"`
	ComposeFilePath   string   `json:"composeFilePath,omitempty"`
	NetworkMode       string   `json:"networkMode,omitempty"`
	Volumes           []Volume `json:"volumes,omitempty"`
	Ports             []Port   `json:"ports,omitempty"`
	Environment       []EnvVar `json:"environment,omitempty"`
	RestartPolicy     string   `json:"restartPolicy,omitempty"`
}

type KubernetesMigration struct { // details for Kubernetes deployment
	Namespace       string `json:"namespace"`
	DeploymentYAML  string `json:"deploymentYaml,omitempty"`
	ServiceYAML     string `json:"serviceYaml,omitempty"`
	ConfigMapYAML   string `json:"configMapYaml,omitempty"`
	SecretYAML      string `json:"secretYaml,omitempty"`
	PersistentVolumeYAML string `json:"persistentVolumeYaml,omitempty"`
	HelmChartPath   string `json:"helmChartPath,omitempty"`
	HelmValuesYAML  string `json:"helmValuesYaml,omitempty"`
	KubeConfigPath  string `json:"kubeConfigPath"`
}

type ManagedServiceMigration struct { // details for managed service migration
	SourceServiceName string `json:"sourceServiceName"`
	SourceServiceID   string `json:"sourceServiceId,omitempty"`
	TargetServiceName string `json:"targetServiceName"`
	TargetServiceTier string `json:"targetServiceTier,omitempty"`
	DataExportMethod  string `json:"dataExportMethod"`
	DataImportMethod  string `json:"dataImportMethod"`
	BackupPath        string `json:"backupPath,omitempty"`
	APIParameters     map[string]string `json:"apiParameters,omitempty"`
}

type Volume struct { // container volume mount
	HostPath      string `json:"hostPath"`
	ContainerPath string `json:"containerPath"`
	ReadOnly      bool   `json:"readOnly,omitempty"`
}

type Port struct { // container port mapping
	HostPort      int    `json:"hostPort"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol,omitempty"` // tcp, udp
}

type EnvVar struct { // environment variable for container
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ConfigData struct { // software-specific configuration info
	CommonConfig  map[string]string `json:"commonConfig,omitempty"`
	ConfigFiles   []ConfigFile      `json:"configFiles,omitempty"`
	
	// Software-specific configurations (one will be populated based on software type)
	NginxConfig   *NginxConfig     `json:"nginxConfig,omitempty"`
	NFSConfig     *NFSConfig       `json:"nfsConfig,omitempty"`
	MySQLConfig   *MySQLConfig     `json:"mysqlConfig,omitempty"`
}

type ConfigFile struct {
	Path       string `json:"path"`
	Content    string `json:"content,omitempty"`
	SourcePath string `json:"sourcePath,omitempty"`
}

type NginxConfig struct {
	ServerBlocks    	[]NginxServerBlock 	`json:"serverBlocks"`
	WorkerProcesses 	int                	`json:"workerProcesses,omitempty"`
	WorkerConnections 	int              	`json:"workerConnections,omitempty"`
	KeepaliveTimeout 	int               	`json:"keepaliveTimeout,omitempty"`
	ClientMaxBodySize 	string           	`json:"clientMaxBodySize,omitempty"`
	SSLCertificates  	[]SSLCertificate  	`json:"sslCertificates,omitempty"`
	LoadBalancing    	*LoadBalancingConfig `json:"loadBalancing,omitempty"`
	MainConfigPath   	string            	`json:"mainConfigPath"`
}

type NginxServerBlock struct { // Nginx server configuration block
	ServerName    string     `json:"serverName"`
	Listen        int        `json:"listen"`
	SSL           bool       `json:"ssl,omitempty"`
	Root          string     `json:"root,omitempty"`
	Index         string     `json:"index,omitempty"`
	Locations     []Location `json:"locations"`
}

type Location struct {  // Nginx location block
	Path        string `json:"path"`
	ProxyPass   string `json:"proxyPass,omitempty"`
	Alias       string `json:"alias,omitempty"`
	TryFiles    string `json:"tryFiles,omitempty"`
	ExtraConfig string `json:"extraConfig,omitempty"`
}

type SSLCertificate struct { // SSL certificate configuration
	CertPath  string `json:"certPath"`
	KeyPath   string `json:"keyPath"`
	ChainPath string `json:"chainPath,omitempty"`
}

type LoadBalancingConfig struct { // Nginx load balancing configuration
	Method      string   `json:"method"` // round_robin, least_conn, ip_hash
	Upstream    string   `json:"upstream"`
	Servers     []string `json:"servers"`
}

type NFSConfig struct {
	Exports         []NFSExport `json:"exports"`
	DefaultOptions  string      `json:"defaultOptions,omitempty"`
	IDMapping       string      `json:"idMapping,omitempty"`
	NFSDThreads     int         `json:"nfsdThreads,omitempty"`
	ExportsFilePath string      `json:"exportsFilePath"`
}

type NFSExport struct { // NFS export configuration
	ExportPath     string   `json:"exportPath"`
	AllowedClients []string `json:"allowedClients"`
	Options        string   `json:"options,omitempty"`
}

type MySQLConfig struct {
	DataDir         string     `json:"dataDir"`
	Port            int        `json:"port"`
	BindAddress     string     `json:"bindAddress,omitempty"`
	ConfigPath      string     `json:"configPath"`
	Databases       []Database `json:"databases"`
	Users           []User     `json:"users"`
	MaxConnections  int        `json:"maxConnections,omitempty"`
	BufferPoolSize  string     `json:"bufferPoolSize,omitempty"`
	LogBin          bool       `json:"logBin,omitempty"`
	CharacterSet    string     `json:"characterSet,omitempty"`
	Collation       string     `json:"collation,omitempty"`
	ReplicationConfig *ReplicationConfig `json:"replicationConfig,omitempty"`
}

type Database struct { // MySQL database
	Name         string `json:"name"`
	CharacterSet string `json:"characterSet,omitempty"`
	Collation    string `json:"collation,omitempty"`
	DumpFile     string `json:"dumpFile,omitempty"`
}

type User struct { // MySQL user
	Username   string   `json:"username"`
	Password   string   `json:"password,omitempty"`
	Host       string   `json:"host,omitempty"`
	Privileges []string `json:"privileges"`
	Databases  []string `json:"databases"`
}

type ReplicationConfig struct { // MySQL replication configuration
	ReplicationType string `json:"replicationType"` // master, slave
	MasterHost      string `json:"masterHost,omitempty"`
	MasterPort      int    `json:"masterPort,omitempty"`
	ReplicaUser     string `json:"replicaUser,omitempty"`
	ReplicaPassword string `json:"replicaPassword,omitempty"`
	AutoPosition    bool   `json:"autoPosition,omitempty"`
}
