package abfvpp

import (
	"fmt"
)

type Abfpolicy struct {
    PolicyId int `json:"policyid"`
    AclId int `json:"aclid"`
    Via string `json:"via"`
}

type Abfattach struct {
    PolicyId int `json:"policyid"`
    Type string `json:"type"`
    Intervpp string `json:"intervpp"`
}

type Abf struct {
    Policy []Abfpolicy `json:"policy"`
    Attach []Abfattach `json:"attach"`
}

func (element Abf) CallBackFunc() {
    fmt.Printf("CallBackFunc abf......\n")
}