package inframodel

type NetworkProperty struct { // note: referrence command `ip route`
	IPv4Networks  []string 		 `json:"ipv4Networks,omitempty" example:"172.26.240.0/20"`
	IPv6Networks  []string 		 `json:"ipv6Networks,omitempty"` // TBD
	// TODO: Add or update fields
}
