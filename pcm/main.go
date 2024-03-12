package main

import (
	"fmt"
	"github.com/JeffLiang2018/go-submodules/events"
	"github.com/JeffLiang2018/go-submodules/ipcMsg"
)

func main() {
	fmt.Println("This is pcm submodules. It would use ipcMsg modules")
	fmt.Println(ipcMsg.OutputString())
	fmt.Println(events.SearchAllEvents())
}
