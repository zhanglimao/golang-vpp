package vxlanvpp

import (
	"fmt"
)

type Vxlan struct {
    Src string `json:"src,omitempty"`
    Dst string `json:"dst,omitempty"`
    Vni int `json:"vni,omitempty"`
}

func (element Vxlan) CallBackFunc() {
    fmt.Printf("CallBackFunc vxlan......\n")
}
