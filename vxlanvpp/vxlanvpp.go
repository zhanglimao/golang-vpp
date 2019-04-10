package vxlanvpp

import (
	"fmt"
)

type Vxlan struct {
    Src string
    Dst string
    Vni int
}

func (element Vxlan) CallBackFunc() {
    fmt.Printf("CallBackFunc vxlan......\n")
}
