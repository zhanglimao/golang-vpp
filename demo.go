package main

import (
    "fmt"
    "encoding/json"
    vc "vpp-demo/vpp_connect"
    aclsvpp "vpp-demo/aclsvpp"
    intervpp "vpp-demo/intervpp"
    abfvpp "vpp-demo/abfvpp"
    vxlanvpp "vpp-demo/vxlanvpp"
    bridgevpp "vpp-demo/bridgevpp"
//    vhostuservpp "vpp-demo/vhostuservpp"
)

type VppStruct struct {
    Name string `json:"name"`
    Intervpp []intervpp.Intervpp `json:"intervpp"`
    Bridge []bridgevpp.Bridge `json:"bridge"`
    Vxlan []vxlanvpp.Vxlan `json:"vxlan"`
    Acl []aclsvpp.Acls `json:"acl"`
    Abf abfvpp.Abf `json:"abf"`
}

type ClusterVpp struct {
    Master VppStruct `json:"master"`
    Nodes []VppStruct `json:"nodes"`
}

func (vs VppStruct) CallBackFunc() {
    fmt.Println("VppStruct CallBackFunc......")
    fmt.Println("Connect Vpp......")
}

func (cp ClusterVpp) CallBackFuncMaster() {
  fmt.Printf("CallBackFuncMaster......\n")
  cp.Master.CallBackFunc()
}

// ---------------------function callback----------------------

func main(){
    js := testdata
    var xm ClusterVpp
    err := json.Unmarshal([]byte(js), &xm)
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Printf("xm: %+v\n", xm)
    
    //xm.CallBackFuncMaster()
    vppConn := vc.GetVppConnect()
    iv := new(intervpp.Intervpp)
//    av := new(aclsvpp.Acls)
//    vhostuser := new(vhostuservpp.VhostUserVpp)

    rets := vppConn.CallVpp(iv.GetConfigureFromVpp)
    iv.FormatConfigToSt(rets)

 //   rets := vc.CallVpp(av.GetConfigureFromVpp)

 //   av.FormatConfigToSt(rets)


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


