package onprem

type OnPremInfra struct {
	Network NetworkProperty  `json:"network,omitempty"`
	Servers []ServerProperty `json:"servers" validate:"required"`
	// TODO: Add other fields
}

type ServerProperty struct {
	Hostname   string                     `json:"hostname"`
	CPU        CpuProperty                `json:"cpu"`
	Memory     MemoryProperty             `json:"memory"`
	RootDisk   DiskProperty               `json:"rootDisk"`
	DataDisks  []DiskProperty             `json:"dataDisks,omitempty"`
	Interfaces []NetworkInterfaceProperty `json:"interfaces"`
	OS         OsProperty                 `json:"os"`
}

type CpuProperty struct {
	Architecture string `json:"architecture" example:"x86_64"`
	Cpus         uint32 `json:"cpus" validate:"required" example:"2"`     // Number of physical CPUs (sockets)
	Cores        uint32 `json:"cores" validate:"required" example:"18"`   // Number of physical cores per CPU
	Threads      uint32 `json:"threads" validate:"required" example:"36"` // Number of logical CPUs (threads) per CPU with hyper-threading enabled
	MaxSpeed     uint32 `json:"maxSpeed,omitempty" example:"3.6"`         // Maximum speed in GHz
	Vendor       string `json:"vendor,omitempty" example:"GenuineIntel"`
	Model        string `json:"model,omitempty" example:"Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"`
}

type MemoryProperty struct {
	Type      string `json:"type" validate:"required" example:"DDR4"`
	TotalSize uint64 `json:"totalSize" validate:"required" example:"128"` // Unit GiB
	Available uint64 `json:"available,omitempty"`                         // Unit GiB
	Used      uint64 `json:"used,omitempty"`                              // Unit GiB
	// TODO: Add or update fields
}

type DiskProperty struct { // note: reference command `df -h`
	Label     string `json:"label" validate:"required"`
	Type      string `json:"type" validate:"required" example:"SSD"`       // SSD, HDD
	TotalSize uint64 `json:"totalSize" validate:"required" example:"1024"` // Unit GiB
	Available uint64 `json:"available,omitempty"`                          // Unit GiB
	Used      uint64 `json:"used,omitempty"`                               // Unit GiB
}

type NetworkInterfaceProperty struct { // note: reference command `ifconfig`
	MacAddress string `json:"macAddress,omitempty"`
	IpAddress  string `json:"ipAddress" validate:"required"`
	Mtu        int    `json:"mtu,omitempty"` // Maximum Transmission Unit (MTU) in bytes
}

type OsProperty struct { // note: reference command `cat /etc/os-release`
	Version         string `json:"version" validate:"required" example:"22.04.3 LTS (Jammy Jellyfish)"` // Full version string
	PrettyName      string `json:"prettyName,omitempty" example:"Ubuntu 22.04.3 LTS"`                   // Pretty name
	Name            string `json:"name,omitempty" example:"Ubuntu"`
	VersionID       string `json:"versionId,omitempty" example:"22.04"`
	VersionCodename string `json:"versionCodename,omitempty" example:"jammy"`
	ID              string `json:"id,omitempty" example:"ubuntu"`
	IDLike          string `json:"idLike,omitempty" example:"debian"`
}

type NetworkProperty struct { // note: referrence command `ip route`
	Gateway string `json:"gateway,omitempty" example:"172.26.240.0/20"` // Gateway IP address
	// TODO: Add or update fields
}
