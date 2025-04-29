package inframodel

// NetworkProerty represents a network for on-premise infrastructure.
// In other perspective, it can be a network for servers and/or a collection of networks extracted from a host.//
// * [Important] Information in the IPv4Networks list should be as non-duplicated as possible.
type NetworkProperty struct { // note: reference command `ip route`, `netstat -rn`, and `lshw -c network`
	IPv4Networks []NetworkDetail `json:"ipv4Networks,omitempty"`
	IPv6Networks []NetworkDetail `json:"ipv6Networks,omitempty"` // TBD
	// TODO: Add or update fields
}

// NetworkDetail represents DefaultGateway and DefaultRouteInterface
// - DefaultGateway info: extracted from a host.
// - DefaultRouteInterface info: extracted from a host based on "the DefaultGateway IFace name".
type NetworkDetail struct {
	DefaultGateway        GatewayProperty          `json:"gateway,omitempty"`
	DefaultRouteInterface NetworkInterfaceProperty `json:"defaultRouteInterface,omitempty"`
}

type GatewayProperty struct {
	IP            string `json:"ip,omitempty" example:"192.168.1.1"`
	InterfaceName string `json:"interfaceName,omitempty" example:"eth0"`
	Metric        int    `json:"metric,omitempty" example:"100"`
}
