package abfvpp

import (
	"fmt"
)

type Abfpolicy struct {
    PolicyId int
    AclId int
    Via string
}

type Abfattach struct {
    PolicyId int
    Type string
    Intervpp string
}

type Abf struct {
    Policy []Abfpolicy
    Attach []Abfattach
}

func (element Abf) CallBackFunc() {
    fmt.Printf("CallBackFunc abf......\n")
}