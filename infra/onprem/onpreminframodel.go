package Onpreminframodel

import (
	cloudinfra  "github.com/cloud-barista/cm-model/infra/cloud"
	software    "github.com/cloud-barista/cm-model/sw"
)

type OnpremiseInfraModel struct {
	OnpremiseInfraModel OnpremInfra `json:"onpremiseInfraModel" validate:"required"`
}

type OnpremInfra struct {
	Network 			NetworkProperty  				`json:"network,omitempty"`
	Servers 			[]ServerProperty 				`json:"servers" validate:"required"`
	Clusters    		[]cloudinfra.ClusterProperty 	`json:"clusters,omitempty"`
	SoftwareProperties 	[]software.SoftwareProperty 	`json:"software_properties"` // For CSP's managed software service.
}

type NetworkProperty struct { // note: referrence command `ip route`
	IPv4Networks  []string 		 `json:"ipv4Networks,omitempty" example:"172.26.240.0/20"`
	IPv6Networks  []string 		 `json:"ipv6Networks,omitempty"` // TBD
}

type ServerProperty struct {
	Hostname      string                      `json:"hostname"`
	CPU           CpuProperty                 `json:"cpu"`
	Memory        MemoryProperty              `json:"memory"`
	RootDisk      DiskProperty                `json:"rootDisk"`
	DataDisks     []DiskProperty              `json:"dataDisks,omitempty"`
	Interfaces    []NetworkInterfaceProperty  `json:"interfaces"`
	RoutingTable  []RouteProperty             `json:"routingTable"`
	FirewallRules []FirewallRuleProperty 	  `json:"firewallRules,omitempty"`
	OS            OsProperty                  `json:"os"`
}

type CpuProperty struct {
	Architecture string  `json:"architecture" example:"x86_64"`
	Cpus         uint32  `json:"cpus" validate:"required" example:"2"`     // Number of physical CPUs (sockets)
	Cores        uint32  `json:"cores" validate:"required" example:"18"`   // Number of physical cores per CPU
	Threads      uint32  `json:"threads" validate:"required" example:"36"` // Number of logical CPUs (threads) per CPU with hyper-threading enabled
	MaxSpeed     float32 `json:"maxSpeed,omitempty" example:"3.6"`         // Maximum speed in GHz
	Vendor       string  `json:"vendor,omitempty" example:"GenuineIntel"`
	Model        string  `json:"model,omitempty" example:"Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"`
}

type MemoryProperty struct {
	Type      string `json:"type" validate:"required" example:"DDR4"`
	TotalSize uint64 `json:"totalSize" validate:"required" example:"128"` // Unit GiB
	Available uint64 `json:"available,omitempty"`                         // Unit GiB
	Used      uint64 `json:"used,omitempty"`                              // Unit GiB
}

type DiskProperty struct { // note: reference command `df -h`
	Label     string `json:"label" validate:"required"`
	Type      string `json:"type" validate:"required" example:"SSD"`       // SSD, HDD
	TotalSize uint64 `json:"totalSize" validate:"required" example:"1024"` // Unit GiB
	Available uint64 `json:"available,omitempty"`                          // Unit GiB
	Used      uint64 `json:"used,omitempty"`                               // Unit GiB
}

type NetworkInterfaceProperty struct { // note: reference command `ifconfig`
	Name           string   `json:"name,omitempty" validate:"required"` // Interface name (e.g., eth0, ens01, enp0s3)
	MacAddress     string   `json:"macAddress,omitempty"`               // MAC address
	IPv4CidrBlocks []string `json:"ipv4CidrBlocks,omitempty"`           // IPv4 address with prefix length (e.g., 192.168.0.21/24), instead of inet addr, Bcast, and Mask
	IPv6CidrBlocks []string `json:"ipv6CidrBlocks,omitempty"`           // IPv6 address with prefix length (e.g., "2001:db8::1/64")
	Mtu            int      `json:"mtu,omitempty"`                      // Maximum Transmission Unit (MTU) in bytes
	State          string   `json:"state,omitempty"`                    // Interface state (e.g., UP, DOWN)
}

type RouteProperty struct { // note: reference command `ip route`
	Destination string `json:"destination,omitempty"` // Destination network, expressed in CIDR format
	Gateway     string `json:"gateway,omitempty"`     // Gateway address to which packets are forwarded
	Interface   string `json:"interface,omitempty"`   // Network interface associated with the route
	Metric      int    `json:"metric,omitempty"`      // Metric value indicating the priority of the route
	Protocol    string `json:"protocol,omitempty"`    // Protocol used to set the route (e.g., kernel, static)
	Scope       string `json:"scope,omitempty"`       // Scope of the route (e.g., global, link, host)
	Source      string `json:"source,omitempty"`      // Optionally stores the source address (used for policy-based routing)
	LinkState   string `json:"linkState,omitempty"`   // Link state of the route (e.g., UP, DOWN)
}

type FirewallRuleProperty struct { // note: reference command `sudo ufw status verbose`
	SrcCIDR   string `json:"srcCIDR,omitempty"`
	DstCIDR   string `json:"dstCIDR,omitempty"`
	SrcPorts  string `json:"srcPorts,omitempty"`
	DstPorts  string `json:"dstPorts,omitempty"`
	Protocol  string `json:"protocol,omitempty"`  // TCP, UDP, ICMP
	Direction string `json:"direction,omitempty"` // inbound, outbound
	Action    string `json:"action,omitempty"`    // allow, deny
}

type OsProperty struct { // note: reference command `cat /etc/os-release`
	PrettyName      string `json:"prettyName" validate:"required" example:"Ubuntu 22.04.3 LTS"` // Pretty name
	Version         string `json:"version,omitempty" example:"22.04.3 LTS (Jammy Jellyfish)"`   // Full version string
	Name            string `json:"name,omitempty" example:"Ubuntu"`
	VersionID       string `json:"versionId,omitempty" example:"22.04"`
	VersionCodename string `json:"versionCodename,omitempty" example:"jammy"`
	ID              string `json:"id,omitempty" example:"ubuntu"`
	IDLike          string `json:"idLike,omitempty" example:"debian"`
}
