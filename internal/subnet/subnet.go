package subnet

import (
	"net"
)

func CalculateNetwork(cidr string) (net.IP, net.IP, error) {
	ip, ipMask, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, nil, err
	}

	network := ip.Mask(ipMask.Mask)
	broadcast := make(net.IP, len(network))
	for i := range network {
		broadcast[i] = network[i] | ^ipMask.Mask[i]
	}

	return network, broadcast, nil
}

func IPToUint32(ip net.IP) uint32 {
	ip4 := ip.To4()
	return uint32(ip4[0])<<24 | uint32(ip4[1])<<16 | uint32(ip4[2])<<8 | uint32(ip4[3])
}

func Uint32ToIP(ip uint32) net.IP {
	return net.IPv4(byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func GetIPRange(network net.IP, broadcast net.IP) []string {

	networkInt := IPToUint32(network)
	broadcastInt := IPToUint32(broadcast)

	// Make slice large enough to hold all hosts including network and broadcast.
	ips := make([]string, (broadcastInt-networkInt)+1)

	for i := networkInt; i <= broadcastInt; i++ {
		ips[i-networkInt] = (Uint32ToIP(i)).String()
	}

	return ips
}
