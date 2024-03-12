package ipcMsg

import (
	"fmt"
)

const (
	ipcMsgVersion = "v0.0.2"
	PrintString = "ipcMsg " + ipcMsgVersion
)

func OutputString() string {
	fmt.Println("ipcMsg module is called!!!")
	return PrintString
}