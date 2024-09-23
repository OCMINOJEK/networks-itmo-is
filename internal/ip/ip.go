package ip

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddress struct {
	octet1 int8
	octet2 int8
	octet3 int8
	octet4 int8
}

func FromString(str string) (IPAddress, error) {
	octets := strings.Split(str, ".")
	if len(octets) != 4 {
		return IPAddress{}, fmt.Errorf("invalid IP address format")
	}

	var ip IPAddress
	for i, octetStr := range octets {
		octet, err := strconv.ParseInt(octetStr, 10, 8)
		if err != nil || octet < 0 || octet > 255 {
			return IPAddress{}, fmt.Errorf("invalid octet value: %s", octetStr)
		}

		switch i {
		case 0:
			ip.octet1 = int8(octet)
		case 1:
			ip.octet2 = int8(octet)
		case 2:
			ip.octet3 = int8(octet)
		case 3:
			ip.octet4 = int8(octet)
		}
	}

	return ip, nil
}
