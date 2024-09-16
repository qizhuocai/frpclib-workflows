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
func Close(uid *C.char) C.int {
	if sub.Close(C.GoString(uid)) {
		return 1 // 返回1表示成功
	}
	return 0 // 返回0表示失败
}

//export GetUids
func GetUids() *C.char {
	uids := sub.GetUids()
	return C.CString(strings.Join(uids, ","))
}

//export IsRunning
func IsRunning(uid *C.char) C.int {
	if sub.IsRunning(C.GoString(uid)) {
		return 1 // 返回1表示运行中
	}
	return 0 // 返回0表示未运行
}

func main() {
	system.EnableCompatibilityMode()
	sub.Execute()
}
