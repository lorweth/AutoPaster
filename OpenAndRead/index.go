package OpenAndRead

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ReadFile(filename string, sheetName string, colIndex int, startRowIndex int, data chan string) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(sheetName)
	if err != nil {
		panic(err)
	}
	for i, row := range rows {
		if i >= startRowIndex {
			for j, colCell := range row {
				if j == colIndex {
					data <- colCell
				}
			}
		}
	}
	close(data)
}
