package vhostuservpp

import (
    "fmt"
    "strings"
	"git.fd.io/govpp.git/api"
    "git.fd.io/govpp.git/bin_api/vhost_user"
)

type VhostUserVpp struct {
    Name string `json:"name"`
    Index uint32 `json:"index"`
    Server uint8 `json:"server"`
    SockPath string `json:"ipaddr“`
}

func (element *VhostUserVpp) GetConfigureFromVpp(ch api.Channel) interface{}{

    rets := []*vhost_user.SwInterfaceVhostUserDetails{}
    reqCtx := ch.SendMultiRequest(&vhost_user.SwInterfaceVhostUserDump{})

    for {
        msg := &vhost_user.SwInterfaceVhostUserDetails{}
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

    return rets
}

func (element *VhostUserVpp) FormatConfigToSt(config interface {}) interface {} {
    cf := config.([]*vhost_user.SwInterfaceVhostUserDetails)
    rets := []VhostUserVpp{}

    for _, ele := range cf {
		vhostuser := VhostUserVpp{}

        ifaceName := strings.TrimFunc(string(ele.InterfaceName), func(r rune) bool {
            return r == 0x00
        })
        socketpath := strings.TrimFunc(string(ele.SockFilename), func(r rune) bool {
            return r == 0x00
        })
        fmt.Printf("\nvhostuser %q: %+v\n", ifaceName, ele)
        fmt.Printf("\nsocketpath %q: %+v\n", socketpath, ele)
        vhostuser.Name = ifaceName
        vhostuser.Index = ele.SwIfIndex
        vhostuser.SockPath = socketpath
        vhostuser.Server = ele.IsServer
        //vhostuser.Server = element.FormatInterType(ifaceName)

        rets=append(rets, vhostuser)
    }

    return rets
}
