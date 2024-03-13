package events
import (
"fmt"
"github.com/JeffLiang2018/go-submodules/ipcMsg"
)

const EVENTS_VERSION = "v0.0.3"


func SearchAllEvents() string {
	fmt.Println("This is events/searchEngine!!!!!!")
	fmt.Println(ipcMsg.OutputString())
	return "From events/SearchEngine.go "+ EVENTS_VERSION + " SearchAllEvents returns " + ipcMsg.OutputString()
}
