package subnet

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateNetwork(t *testing.T) {
	ipNetwork := "172.20.33.0/23"
	networkWant := net.ParseIP("172.20.32.0")
	broadcastWant := net.ParseIP("172.20.33.255")
	network, broadcast, err := CalculateNetwork(ipNetwork)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	assert.Equal(t, networkWant.To4(), network)
	assert.Equal(t, broadcastWant.To4(), broadcast)

}

func TestIpToUint32(t *testing.T) {
	ip := "172.20.30.33"
	var ipInt uint32 = 2886999585

	uint := IPToUint32(net.ParseIP(ip))

	if uint != ipInt {
		t.Errorf("got %d, expected %d", uint, ipInt)
	}
}

func TestUint32ToIP(t *testing.T) {
	ip := "172.20.30.33"
	var ipInt uint32 = 2886999585

	ipGot := Uint32ToIP(ipInt)

	assert.Equal(t, net.ParseIP(ip), ipGot)

}

func TestGetIPRange(t *testing.T) {
	ipNetwork := "192.168.30.0/30"
	ipsWant := []string{
		"192.168.30.0",
		"192.168.30.1",
		"192.168.30.2",
		"192.168.30.3",
	}

	network, broadcast, _ := CalculateNetwork(ipNetwork)

	ips := GetIPRange(network, broadcast)

	assert.Equal(t, ipsWant, ips)

}
