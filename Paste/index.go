package Paste

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func GetMousePosition(second int) (int, int) {
	time.Sleep(5 * time.Second)
	x, y := robotgo.GetMousePos()
	fmt.Printf("Dự liệu sẽ được nhập vào vị trí: {%d, %d}\n", x, y)
	return x, y
}

func ClickAndPaste(x, y int, str string) {
	robotgo.Move(x, y)
	robotgo.Click("left", true)
	robotgo.TypeStr(str, 0.5)
	robotgo.Sleep(1)
	robotgo.KeyTap("enter")
}
