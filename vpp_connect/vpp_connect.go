package vpp_connect

import (
    "log"
    "sync"
    "git.fd.io/govpp.git"
    "git.fd.io/govpp.git/core"
    "git.fd.io/govpp.git/api"
)

type Callback func(ch api.Channel) interface{}

type vppConnect struct {
    conn *core.Connection
    ch api.Channel
}

var instance *vppConnect
var mu sync.Mutex

func (vc *vppConnect) connectVpp() {
    var err error
    conev := make(chan core.ConnectionEvent)

    // connect to VPP
    vc.conn, conev, err = govpp.AsyncConnect("")
    if err != nil {
        log.Fatalln("ERROR:", err)
    }

    select {
    case e := <-conev:
        if e.State != core.Connected {
            log.Fatalf("failed to connect: %v", e.Error)
        }
    }

    vc.ch, err = vc.conn.NewAPIChannel()
    if err != nil {
        log.Fatalln("ERROR:", err)
    }
}

func GetVppConnect() *vppConnect {
    mu.Lock()
    defer mu.Unlock()

    if instance == nil {
        instance = &vppConnect{}
        instance.connectVpp()
    }
    return instance
}

func (vc *vppConnect) DisConnectVpp() {
    vc.conn.Disconnect()
    vc.ch.Close()
}

func (vc *vppConnect) CallVpp(cb Callback) interface{} {
    return cb(vc.ch)
}