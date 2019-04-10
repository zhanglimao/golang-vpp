package main

import (
    "fmt"
    "reflect"
    "strings"
    "git.fd.io/govpp.git/api"
    vc "vpp-demo/vpp_connect"
    aclsvpp "vpp-demo/aclsvpp"
    intervpp "vpp-demo/intervpp"
    abfvpp "vpp-demo/abfvpp"
    vxlanvpp "vpp-demo/vxlanvpp"
    bridgevpp "vpp-demo/bridgevpp"
)

type VppStructInter interface {
    GetConfigureFromVpp(ch api.Channel) interface{}
    FormatConfigToSt(config interface {}) interface{}
}

type VppStruct struct {
    Name string
    Intervpp []intervpp.Intervpp
    Bridge []bridgevpp.Bridge
    Vxlan []vxlanvpp.Vxlan
    Acl []aclsvpp.Acls
    Abf abfvpp.Abf
}

func (vs VppStruct) CallBackFunc() {
    fmt.Println("VppStruct CallBackFunc......")
    fmt.Println("Connect Vpp......")
}

func (cv *VppStruct) DumpVppToSt() {
  fmt.Printf("DumpVppToSt......\n")

}

func (cv VppStruct) DumpStToJson() {
  fmt.Printf("DumpStToJson......\n")

}

// ---------------------function callback----------------------

func main(){

    ins := VppStruct{}
    vppConn := vc.GetVppConnect()

    intervppst := &intervpp.Intervpp{}
    aclvppst := &aclsvpp.Acls{}

    interrfacesVPP := []VppStructInter{intervppst, aclvppst}

    for _, ele := range interrfacesVPP {
      ret := vppConn.CallVpp(ele.GetConfigureFromVpp)
      retSt := ele.FormatConfigToSt(ret)
      stType := reflect.TypeOf(retSt).String()
      fmt.Printf("%v\n",strings.Split(stType,"."))

      switch strings.Split(stType,".")[1] {
        case "Acls":
          ins.Acl = retSt.([]aclsvpp.Acls)
        case "Intervpp":
          ins.Intervpp = retSt.([]intervpp.Intervpp)
      }

      fmt.Println("-------------------------------")
      fmt.Printf("%+v\n", ins)
    }

    vppConn.DisConnectVpp()

}