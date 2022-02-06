package main

import (
	"autopaster/OpenAndRead"
	"autopaster/Paste"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-vgo/robotgo"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	loadEnv()

	// Retrieve env variables
	filename := os.Getenv("FILENAME")
	sheetname := os.Getenv("SHEETNAME")
	colIndex, _ := strconv.Atoi(os.Getenv("COL_INDEX"))
	rowStartIndex, _ := strconv.Atoi(os.Getenv("ROW_START_INDEX"))

	if filename == "" || sheetname == "" || colIndex < 0 || rowStartIndex < 0 {
		log.Fatal("Please set env variables")
	}

	// Get current mouse position
	x, y := getPosition()
	if x == 0 && y == 0 {
		log.Fatal("Lỗi chọn vị trí")
	}
	fmt.Printf("Sẽ nhập tại vị trí: {X: %d, Y: %d}\n", x, y)

	// Channel to receive data
	data := make(chan string)

	// Open excel file
	go OpenAndRead.ReadFile(filename, sheetname, colIndex, rowStartIndex, data)

	// Loop to get data from channel
	for d := range data {
		fmt.Println("Đã nhập: " + d)
		Paste.ClickAndPaste(x, y, d)
	}
}

// Load .env file
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Lỗi khi đọc file .env")
	}
}

// Get current mouse position
func getPosition() (int, int) {
	fmt.Println("-- Nhấn ctrl + i để chọn vị trí nhập dữ liệu --")

	ev := robotgo.AddEvents("i", "ctrl")

	if ev {
		x, y := robotgo.GetMousePos()
		return x, y
	}

	return 0, 0
}
