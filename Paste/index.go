package Paste

import (
	"github.com/go-vgo/robotgo"
)

func ClickAndPaste(x, y int, str string) {
	robotgo.Move(x, y)
	robotgo.Click()
	robotgo.PasteStr(str)
	robotgo.Sleep(1)
	robotgo.KeyTap("enter")
}
