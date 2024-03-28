package config
import (

"fmt"
"github.com/JeffLiang2018/go-submodules/pcm/cfgMsg"
)

func OutputString() string {
	fmt.Println("In pcm/cfgMsg, " + cfgMsg.OutputCfgMsg())
	return "This is package pcm/config."
}
