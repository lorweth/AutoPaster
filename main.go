package main

import (
	OpenAndRead "autopaster/OpenAndRead"
	"autopaster/Paste"
	"fmt"
)

func main() {

	var filename string
	var sheetName string
	var colIndex int
	var startRowIndex int

	fmt.Println("------------------------ Ctrl + C để thoát ------------------------")
	fmt.Println("Hãy để file excel của bạn ở cùng thư mục với app này")
	fmt.Println("Nhập tên file excel:")
	fmt.Scanln(&filename)
	fmt.Println("Nhập tên sheet:")
	fmt.Scanln(&sheetName)
	fmt.Println("Nhập số thứ tự cột:")
	fmt.Scanln(&colIndex)
	fmt.Println("Nhập số thứ tự dòng bắt đầu:")
	fmt.Scanln(&startRowIndex)
	fmt.Println("Bạn có 5s để chọn vị trí chuột")

	data := make(chan string)

	x, y := Paste.GetMousePosition(5)

	go OpenAndRead.ReadFile(filename, sheetName, colIndex, startRowIndex, data)

	for i := range data {
		fmt.Println("Đã đọc: " + i)
		Paste.ClickAndPaste(x, y, i)
	}
}
