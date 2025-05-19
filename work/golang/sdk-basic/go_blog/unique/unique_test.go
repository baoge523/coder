package unique

import (
	"fmt"
	"strings"
	"testing"
)

// https://golang.google.cn/blog/unique

var uniqueMapPool map[string]string

/**
point 0x14000102230
point 0x14000102230
a point 0x14000102240
b point 0x14000102250
*/
func TestSimpleString(t *testing.T) {
	uniqueMapPool = make(map[string]string)
	val := Internal("abcd")
	fmt.Printf("point %p \n",&val)
	val = Internal("abcd")
	fmt.Printf("point %p \n",&val)

	a := "qwer" // 0x14000102240
	b := "qwer" // 0x14000102250
	fmt.Printf("a point %p \n",&a)
	fmt.Printf("b point %p \n",&b)
}

// 这里存在几个问题：
// 1、it never removes strings from pool
// 2、it cannot be safely used by multiple goroutines concurrently
// 3、it only works with strings
func Internal(value string) string {
	val,ok := uniqueMapPool[value]
	if !ok {
		val = strings.Clone(value)
		uniqueMapPool[val] = val
	}
	return val
}

// ------------- 参考 net/netip package in the standard library ----------------

// 需要的版本是1.23

// Addr represents an IPv4 or IPv6 address (with or without a scoped
// addressing zone), similar to net.IP or net.IPAddr.
type Addr struct {
	// Other irrelevant unexported fields...

	// Details about the address, wrapped up together and canonicalized.
	// z unique.Handle[addrDetail]
}

// addrDetail indicates whether the address is IPv4 or IPv6, and if IPv6,
// specifies the zone name for the address.
type addrDetail struct {
	isV6   bool   // IPv4 is false, IPv6 is true.
	zoneV6 string // May be != "" if IsV6 is true.
}

// var z6noz = unique.Make(addrDetail{isV6: true})

// WithZone returns an IP that's the same as ip but with the provided
// zone. If zone is empty, the zone is removed. If ip is an IPv4
// address, WithZone is a no-op and returns ip unchanged.
func (ip Addr) WithZone(zone string) Addr {
	//if !ip.Is6() {
	//	return ip
	//}
	//if zone == "" {
	//	ip.z = z6noz
	//	return ip
	//}
	//ip.z = unique.Make(addrDetail{isV6: true, zoneV6: zone})
	return ip
}
