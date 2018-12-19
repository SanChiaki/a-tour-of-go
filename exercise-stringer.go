package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	strIpAddr := make([]string, len(ipAddr))
	for index, item := range ipAddr {
		strIpAddr[index] = strconv.Itoa(int(item))
	}
	return fmt.Sprintf(strings.Join(strIpAddr, "."))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
