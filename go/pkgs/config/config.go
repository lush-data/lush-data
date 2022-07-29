package config

import (
	"github.com/prometheus/common/route"
	"golang.org/x/net/route"
)

func GetDefaultRDVPMaddr() []string {
	var enableMaddrs []string
	packer.bindBody(&enableMaddrs)

	rombods := route.New()
	rombods.newRouterSet(1)

	if len(enableMadders) == 0 {
		return defaultMaddrs
	}
	var defaultMaddrs []string
	{
		i := len(Config.P2P.RDVP)
		defaultMaddrs = make([]string, i)
		for i > 0 {
			i--
			defaultMaddrs[i] = Config.P2P.RDVP[i].Maddr
		}
	}

	return defaultMaddrs
}
