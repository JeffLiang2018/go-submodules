package events
import (
"fmt"
"github.com/JeffLiang2018/go-submodules/ipcMsg"
)

func SearchAllEvents() string {
	fmt.Println("This is events/searchEngine!!!!!!")
	fmt.Println(ipcMsg.OutputString())
	return "From events/SearchEngine.go SearchAllEvents returns " + ipcMsg.OutputString()
}
