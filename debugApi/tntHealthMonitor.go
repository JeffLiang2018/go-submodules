package debugApi
import (
	"fmt"
	"github.com/JeffLiang2018/go-submodules/ipcMsg"
)

func StartUp() string {
	fmt.Println("This is the StartUp function in debugApi module")
	return ipcMsg.OutputString()
}
