package aclsvpp

import (
	"fmt"
	"strconv"
//	"encoding/binary"
	//"encoding/json"
	//"reflect"
	"git.fd.io/govpp.git/api"
	"git.fd.io/govpp.git/bin_api/acl"
)

type aclRule struct {
	IsPermit               uint8
	IsIPv6                 uint8
	SrcIPAddr              string
	SrcIPPrefixLen         uint8
	DstIPAddr              string
	DstIPPrefixLen         uint8
	Proto                  uint8
	SrcportOrIcmptypeFirst uint16
	SrcportOrIcmptypeLast  uint16
	DstportOrIcmpcodeFirst uint16
	DstportOrIcmpcodeLast  uint16
	TCPFlagsMask           uint8
	TCPFlagsValue          uint8
}

type Acls struct {
    Index uint32 `json:"index"`
    Tag int `json:"tag"`
    Rule []aclRule `json:"rule"`
}

func (element *Acls) GetConfigureFromVpp(ch api.Channel) interface{} {
	rets := []*acl.ACLDetails{}
	reqCtx := ch.SendMultiRequest(&acl.ACLDump{ACLIndex:-1})
  
	for {
		msg := &acl.ACLDetails{}
		stop, err := reqCtx.ReceiveReply(msg)
		if stop {
			break
		}
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		rets=append(rets, msg)
	}
	return rets
}

func UINT8ToIP(bs []uint8) string {
	var st []string

	for _, b := range bs {
		//fmt.Printf("\nstrconv.Itoa(int(b)): %+v\n", strconv.Itoa(int(b)))
		//fmt.Printf("\nreflect.TypeOf(strconv.Itoa(int(b))): %+v\n", reflect.TypeOf(strconv.Itoa(int(b))))
		st = append(st, strconv.Itoa(int(b)))
	}

	ret := fmt.Sprintf("%s.%s.%s.%s",st[0],st[1],st[2],st[3])

	//fmt.Printf("%v", ret)

	return ret
}

func UINT8ToInt(bs []uint8) int {
	var st []string

	/*
	for _, b := range bs {
		//fmt.Printf("\nstrconv.Itoa(int(b)): %+v\n", strconv.Itoa(int(b)))
		//fmt.Printf("\nreflect.TypeOf(strconv.Itoa(int(b))): %+v\n", reflect.TypeOf(strconv.Itoa(int(b))))
		st = append(st, strconv.Itoa(int(b)))
	}
	*/
	st = append(st, st[12])
	st = append(st, st[13])
	st = append(st, st[14])
	st = append(st, st[15])

	//str := fmt.Sprintf("%s%s%s%s",st[0],st[1],st[2],st[3])



	fmt.Printf("string(bs[12:15]): %v", string(bs[12:15]))

	//ret, _ := strconv.Atoi(st)

	return 1
	//return ret
}

func (element *Acls) FormatConfigToSt(config interface {}) interface {}{
	cf := config.([]*acl.ACLDetails)
	rets := []Acls{}

	for _, e := range cf {
		fmt.Printf("\ne: %+v\n", e)
	}
	
	fmt.Printf("length: %v\n", len(cf))

    for _, ele := range cf {
		msg := Acls{}
		msg.Index = ele.ACLIndex
		//msg.Tag = UINT8ToInt(ele.Tag)
		
		for _, e := range ele.R {
			tmp := aclRule{}
			tmp.IsPermit = e.IsPermit
			tmp.IsIPv6 = e.IsIPv6
			//UINT8ToIP(e.SrcIPAddr)
			tmp.SrcIPAddr = UINT8ToIP(e.SrcIPAddr)
			//fmt.Printf("\ntmp.SrcIPAddr: %+v\n", tmp.SrcIPAddr)
			tmp.SrcIPPrefixLen = e.SrcIPPrefixLen
			tmp.DstIPAddr = UINT8ToIP(e.DstIPAddr)
			tmp.DstIPPrefixLen = e.DstIPPrefixLen
			tmp.Proto = e.Proto
			tmp.SrcportOrIcmptypeFirst = e.SrcportOrIcmptypeFirst
			tmp.SrcportOrIcmptypeLast = e.SrcportOrIcmptypeLast
			tmp.DstportOrIcmpcodeFirst = e.DstportOrIcmpcodeFirst
			tmp.DstportOrIcmpcodeLast = e.DstportOrIcmpcodeLast
			tmp.TCPFlagsMask = e.TCPFlagsMask
			tmp.TCPFlagsValue = e.TCPFlagsValue

			fmt.Printf("\ntmp: %+v\n", tmp)

			msg.Rule = append(msg.Rule, tmp)
		}
		
		//msg.Rule = ele.R.([]Acls)
		rets=append(rets, msg)
	}
	fmt.Printf("rets: %+v\n", rets)
	return rets
}

func (element Acls) CallBackFunc() {
    fmt.Printf("CallBackFunc acls......\n")
}