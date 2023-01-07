package mac

import "net"

// MacInfo
type MacInfo struct {
	Mac  string `json:"mac"`
	Ipv4 string `json:"ipv4"`
	Ipv6 string `json:"ipv6"`
}

func Fetch() ([]MacInfo, error) {

	// list system network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	mac := []MacInfo{}
	//loop system network interfaces
	for _, inter := range interfaces {

		if inter.Flags & net.FlagLoopback > 0 {
			continue
		}
		if inter.Flags & net.FlagUp == 0 {
			continue
		}
		if inter.HardwareAddr == nil {
		    continue
		}

		data := new(MacInfo)
		data.Mac = inter.HardwareAddr.String()

		// get ipv4 and ipv6
		addrs, _ := inter.Addrs()
		for _, addr := range addrs {
			ipnet, _ := addr.(*net.IPNet)
			if ipnet.IP.To4() != nil {
				data.Ipv4 = ipnet.IP.String()
			} else {
				data.Ipv6 = ipnet.IP.String()
			}
		}
		mac = append(mac, *data)

	}
	return mac, nil

}
