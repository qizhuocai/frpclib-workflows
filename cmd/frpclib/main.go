package main

import (
	"strings"

	_ "github.com/fatedier/frp/assets/frpc"
	"github.com/fatedier/frp/cmd/frpclib/sub"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/fatedier/frp/pkg/util/version"
)

func main() {
	system.EnableCompatibilityMode()
	sub.Execute()
}

func GetVersion() string {
	return version.Full()
}

func RunFile(uid string, cfgFilePath string) (errString string) {
	err, _ := sub.RunClientWithUid(uid, cfgFilePath)
	if err != nil {
		return err.Error()
	}
	return ""
}

func RunContent(uid string, cfgContent string) (errString string) {
	err, _ := sub.RunClientByContent(uid, cfgContent)
	if err != nil {
		return err.Error()
	}
	return ""
}

func Close(uid string) (ret bool) {
	return sub.Close(uid)
}

func GetUids() string {
	uids := sub.GetUids()
	return strings.Join(uids, ",")
}

func IsRunning(uid string) (running bool) {
	return sub.IsRunning(uid)
}
