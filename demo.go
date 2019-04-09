package main

import (
    "fmt"
//    "encoding/json"
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
    Name string `json:"name"`
    Intervpp []intervpp.Intervpp `json:"intervpp"`
    Bridge []bridgevpp.Bridge `json:"bridge,omitempty"`
    Vxlan []vxlanvpp.Vxlan `json:"vxlan,omitempty"`
    Acl []aclsvpp.Acls `json:"acl,omitempty"`
    Abf abfvpp.Abf `json:"abf,omitempty"`
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
/*
    js := testdata
    var xm ClusterVpp
    err := json.Unmarshal([]byte(js), &xm)
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Printf("xm: %+v\n", xm)
*/   
    //xm.CallBackFuncMaster()

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
        case "Acl":
          ins.Acl = retSt.([]aclsvpp.Acls)
        case "Intervpp":
          ins.Intervpp = retSt.([]intervpp.Intervpp)
      }

      fmt.Println("-------------------------------")
      fmt.Printf("%+v", ins)
    }

    //vppConn.CallVpp(interrfacesVPP[0].GetConfigureFromVpp)

    //iv.FormatConfigToSt(rets)

 // rets := vc.CallVpp(av.GetConfigureFromVpp)

 // av.FormatConfigToSt(rets)


    vppConn.DisConnectVpp()

}

var testdata = `{
  "master": {
    "name": "master",
    "intervpp": [
      {
        "name": "TenGigabitEthernet81/0/0",
        "type": "Normal",
        "state": "up",
        "ipaddr": "1.1.1.2/24"
      },
      {
        "name": "TenGigabitEthernet81/0/1",
        "type": "Normal",
        "state": "up",
        "ipaddr": "4.4.4.2/24"
      },
      {
        "name": "GigabitEthernet83/0/2",
        "type": "Normal",
        "state": "up",
        "ipaddr": "8.8.8.1/24"
      },
      {
        "name": "loop0",
        "type": "LoopBack",
        "state": "up",
        "ipaddr": "2.2.2.2/24, 3.3.3.3/24",
        "mac": "de:ad:00:00:00:11"
      },
      {
        "name": "VirtualEthernet0/0/0",
        "type": "VhostUser-server",
        "state": "up",
        "sockpath": "/tmp/sock1.sock"
      },
      {
        "name": "VirtualEthernet0/0/1",
        "type": "VhostUser-server",
        "state": "up",
        "sockpath": "/tmp/sock2.sock"
      }
    ],
    "bridge": [
      {
        "intervpp": "loop0",
        "type": "bvi"
      },
      {
        "intervpp": "VirtualEthernet0/0/0",
        "type": "port"
      },
      {
        "intervpp": "VirtualEthernet0/0/1",
        "type": "port"
      },
      {
        "intervpp": "GigabitEthernet83/0/2",
        "type": "port"
      }
    ],
    "vxlan": [
      {
        "src": "8.8.8.1",
        "dst": "8.8.8.2",
        "vni": 10
      },
      {
        "src": "8.8.8.1",
        "dst": "8.8.8.2",
        "vni": 10
      }
    ],
    "acl": [
      {
        "index": 0,
        "tag": 1,
        "rule": [
          {
            "isPermit": 1,
            "srcipaddr": "1.1.1.1",
            "dstipaddr": "2.2.2.2",
            "proto": 6
          },
          {
            "isPermit": 1,
            "srcipaddr": "3.3.3.3",
            "dstipaddr": "4.4.4.4",
            "proto": 7
          }
        ]
      },
      {
        "index": 1,
        "tag": 2,
        "rule": [
          {
            "isPermit": 0,
            "srcipaddr": "1.1.1.1",
            "dstipaddr": "2.2.2.2",
            "proto": 8
          },
          {
            "isPermit": 0,
            "srcipaddr": "3.3.3.3",
            "dstipaddr": "4.4.4.4",
            "proto": 9
          }
        ]
      }
    ],
    "abf": {
      "policy": [
        {
          "policyid": 1,
          "aclid": 0,
          "via": "2.2.2.1"
        },
        {
          "policyid": 2,
          "aclid": 0,
          "via": "2.2.2.1"
        },
        {
          "policyid": 3,
          "aclid": 0,
          "via": "2.2.2.1"
        },
        {
          "policyid": 4,
          "aclid": 0,
          "via": "2.2.2.1"
        }
      ],
      "attach": [
        {
          "policyid": 1,
          "type": "ip4",
          "intervpp": "TenGigabitEthernet81/0/0"
        },
        {
          "policyid": 2,
          "type": "ip4",
          "intervpp": "TenGigabitEthernet81/0/0"
        },
        {
          "policyid": 3,
          "type": "ip4",
          "intervpp": "TenGigabitEthernet81/0/1"
        },
        {
          "policyid": 4,
          "type": "ip4",
          "intervpp": "TenGigabitEthernet81/0/1"
        }
      ]
    }
  },
  "nodes": [
    {
      "name": "node1",
      "intervpp": [
        {
          "name": "TenGigabitEthernet81/0/0",
          "type": "Normal",
          "state": "up",
          "ipaddr": "1.1.1.2/24"
        }
      ],
      "bridge": [
        {
          "intervpp": "loop0",
          "type": "bvi"
        },
        {
          "intervpp": "VirtualEthernet0/0/0",
          "type": "port"
        }
      ],
      "vxlan": [
        {
          "src": "8.8.8.1",
          "dst": "8.8.8.2",
          "vni": 10
        }
      ],
      "acls": [
        {
          "id": 0,
          "type": "ipv4",
          "action": "permit",
          "proto": 17
        }
      ],
      "abf": {
        "policy": [
          {
            "policyid": 1,
            "aclid": 0,
            "via": "2.2.2.1"
          },
          {
            "policyid": 2,
            "aclid": 0,
            "via": "2.2.2.1"
          }
        ]
      },
      "attach": [
        {
          "policyid": 1,
          "type": "ip4",
          "intervpp": "TenGigabitEthernet81/0/0"
        },
        {
          "policyid": 2,
          "type": "ip4",
          "intervpp": "TenGigabitEthernet81/0/0"
        }
      ]
    },
    {
      "name": "node2",
      "intervpp": [
        {
          "name": "TenGigabitEthernet81/0/0",
          "type": "Normal",
          "state": "up",
          "ipaddr": "1.1.1.2/24"
        }
      ],
      "bridge": [
        {
          "intervpp": "loop0",
          "type": "bvi"
        },
        {
          "intervpp": "VirtualEthernet0/0/0",
          "type": "port"
        }
      ],
      "vxlan": [
        {
          "src": "8.8.8.1",
          "dst": "8.8.8.2",
          "vni": 10
        }
      ],
      "acls": [
        {
          "id": 0,
          "type": "ipv4",
          "action": "permit",
          "proto": 17
        }
      ],
      "abf": {
        "policy": [
          {
            "policyid": 1,
            "aclid": 0,
            "via": "2.2.2.1"
          },
          {
            "policyid": 2,
            "aclid": 0,
            "via": "2.2.2.1"
          }
        ]
      },
      "attach": [
        {
          "policyid": 1,
          "type": "ip4",
          "intervpp": "TenGigabitEthernet81/0/0"
        },
        {
          "policyid": 2,
          "type": "ip4",
          "intervpp": "TenGigabitEthernet81/0/0"
        }
      ]
    }
  ]
}`


