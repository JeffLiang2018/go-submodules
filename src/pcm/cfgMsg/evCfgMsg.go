package cfgMsg
import (
	"fmt"
	"github.com/JeffLiang2018/go-submodules/events"
)

func OutputCfgMsg() string {
	fmt.Println("In pcm/cfgMsg, " + events.SearchAllEvents())
	return "This is package pcm/cfgMsg."
}