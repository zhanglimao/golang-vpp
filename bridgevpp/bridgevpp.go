package bridgevpp

import (
	"fmt"
)

type Bridge struct {
    Intervpp string `json:"intervpp"`
    Type string `json:"type"`
}

func (element Bridge) CallBackFunc() {
    fmt.Printf("CallBackFunc bridge......\n")
}