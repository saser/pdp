package server

import (
	"github.com/Saser/pdp/aip/resource"
	"github.com/Saser/pdp/aip/resourcename"

	financepb "github.com/Saser/pdp/finance/finance_go_proto"
)

var (
	accountPattern *resourcename.Pattern
)

func init() {
	accountRD := resource.DescriptorOf(&financepb.Account{})
	accountPattern = resourcename.MustCompile(accountRD.GetPattern()[0])
}
