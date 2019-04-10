package intervpp

import (
    "fmt"
    "strings"
    "regexp"
	"git.fd.io/govpp.git/api"
    "git.fd.io/govpp.git/bin_api/interfaces"
    "vpp-demo/vhostuservpp"
    vc "vpp-demo/vpp_connect"
)

type Intervpp struct {
    Name string
    Type string
    State string
    IpAddr string
    Mac string
    SockPath string
}

func (element *Intervpp) GetConfigureFromVpp(ch api.Channel) interface{} {
    rets := []*interfaces.SwInterfaceDetails{}
    reqCtx := ch.SendMultiRequest(&interfaces.SwInterfaceDump{})

    for {
        msg := &interfaces.SwInterfaceDetails{}
        stop, err := reqCtx.ReceiveReply(msg)
        if stop {
            break
        }
        if err != nil {
            fmt.Println("ERROR:", err)
        }
        // ifaceName := strings.TrimFunc(string(msg.InterfaceName), func(r rune) bool {
        //     return r == 0x00
        // })
        // fmt.Printf("\nInterface %q: %+v\n", ifaceName, msg)
        rets=append(rets, msg)
    }
    //fmt.Printf("\nInterface rets: %+v\n", rets)
    return rets
}

func (element *Intervpp) FormatInterType(interName string) string {
    regPolicy := regexp.MustCompile(`(.*)Ethernet.*`)
    ret := (regPolicy.FindSubmatch([]byte(interName)))
    if len(ret) > 1 {
    inter := string(regPolicy.FindSubmatch([]byte(interName))[1])
        switch inter {
        case "Gigabit":
            return "Gigabit"
        case "TenGigabit":
            return "TenGigabit"
        case "Virtual":
            return "Virtual"
        }
    } else {
        regPolicy = regexp.MustCompile(`(loop)(.*)`)
        ret := regPolicy.FindSubmatch([]byte(interName))
        if len(ret) > 1 {
            inter := string(regPolicy.FindSubmatch([]byte(interName))[1])
            if string(inter) == "loop" {
                return "LoopBack"
            }
        }
    }
    return "UnKnowType"
 }

func (element *Intervpp) FormatConfigToSt(config interface {}) interface {} {
    cf := config.([]*interfaces.SwInterfaceDetails)
    rets := []Intervpp{}

    vppConn := vc.GetVppConnect()

    vhostuser := new(vhostuservpp.VhostUserVpp)
    vhostrets := vppConn.CallVpp(vhostuser.GetConfigureFromVpp)
    vhostusers := vhostuser.FormatConfigToSt(vhostrets).([]vhostuservpp.VhostUserVpp)

    fmt.Printf("vhostusers: %+v", vhostusers)

    for _, ele := range cf {
        inter := Intervpp{}
        ifaceName := strings.TrimFunc(string(ele.InterfaceName), func(r rune) bool {
            return r == 0x00
        })
        //fmt.Printf("\nInterface %q: %+v\n", ifaceName, ele)
        inter.Name = ifaceName
        intertype := element.FormatInterType(ifaceName)
        switch intertype {
            case "Virtual":
                for _, vhost := range vhostusers {
                    if ele.SwIfIndex == vhost.Index {
                        inter.SockPath = vhost.SockPath
                        if vhost.Server>0 {
                            inter.Type = "VhostUser-server"
                        } else {
                            inter.Type = "VhostUser-client"
                        }
                    }
                }
            case "Gigabit":
                inter.Type = "TP"
            case "TenGigabit":
                inter.Type = "Fiber"
            case "LoopBack":
                inter.Type = "LoopBack"
        }
        
        if ele.AdminUpDown > 0 {
            inter.State = "UP"
        } else {
            inter.State = "DOWN"
        }

        rets=append(rets, inter)
    }
    fmt.Printf("\n interface: %+v\n", rets)

    return rets
}

// ---------------------function callback----------------------
func (element Intervpp) CallBackFunc(ch api.Channel) {
    fmt.Printf("CallBackFunc intervpp......\n")
    fmt.Printf("intervpp.Name:%s\n", element.Name)
    fmt.Printf("intervpp.Type:%s\n", element.Type)
    fmt.Printf("intervpp.State:%s\n", element.State)
    fmt.Printf("intervpp.IpAddr:%s\n", element.IpAddr)
    fmt.Printf("intervpp.Mac:%s\n", element.Mac)
    fmt.Printf("intervpp.SockPath:%s\n", element.SockPath)
}