package main

import (
	"C" // 必须导入C包以使用 //export 注释
	"strings"

	_ "github.com/fatedier/frp/assets/frpc"
	"github.com/fatedier/frp/cmd/frpclib/sub"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/fatedier/frp/pkg/util/version"
)

//export GetVersion
func GetVersion() *C.char {
	return C.CString(version.Full())
}

//export RunFile
func RunFile(uid *C.char, cfgFilePath *C.char) *C.char {
	err, _ := sub.RunClientWithUid(C.GoString(uid), C.GoString(cfgFilePath))
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export RunContent
func RunContent(uid *C.char, cfgContent *C.char) *C.char {
	err, _ := sub.RunClientByContent(C.GoString(uid), C.GoString(cfgContent))
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export Close
func Close(uid *C.char) C.bool {
	return C.bool(sub.Close(C.GoString(uid)))
}

//export GetUids
func GetUids() *C.char {
	uids := sub.GetUids()
	return C.CString(strings.Join(uids, ","))
}

//export IsRunning
func IsRunning(uid *C.char) C.bool {
	return C.bool(sub.IsRunning(C.GoString(uid)))
}

func main() {
	system.EnableCompatibilityMode()
	sub.Execute()
}
