package abfvpp

import (
	"fmt"
)

type Abfpolicy struct {
    PolicyId int `json:"policyid,omitempty"`
    AclId int `json:"aclid,omitempty"`
    Via string `json:"via,omitempty"`
}

type Abfattach struct {
    PolicyId int `json:"policyid,omitempty"`
    Type string `json:"type,omitempty"`
    Intervpp string `json:"intervpp,omitempty"`
}

type Abf struct {
    Policy []Abfpolicy `json:"policy,omitempty"`
    Attach []Abfattach `json:"attach,omitempty"`
}

func (element Abf) CallBackFunc() {
    fmt.Printf("CallBackFunc abf......\n")
}