package vxlanvpp

import (
	"fmt"
)

type Vxlan struct {
    Src string `json:"src"`
    Dst string `json:"dst"`
    Vni int `json:"vni"`
}

func (element Vxlan) CallBackFunc() {
    fmt.Printf("CallBackFunc vxlan......\n")
}
