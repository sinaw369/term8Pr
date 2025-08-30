package DownloadExecl

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"strings"
)

func CreateExcel(Data []string) {
	// Create a new Excel file
	file := excelize.NewFile()

	// Create a new sheet
	index := file.NewSheet("Sheet1")

	// Set headers
	headers := map[string]string{
		"A1": "fileName",
		"B1": "level",
		"C1": "message",
		"D1": "time",
	}
	style, err := file.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	if err != nil {
		fmt.Println(err)
	}
	file.SetCellStyle("Sheet1", "A1", "D1", style)
	for k, v := range headers {
		file.SetCellValue("Sheet1", k, v)
		// Center align headers
	}

	// Write data to Excel
	row := 2
	for _, value := range Data {
		msg := strings.SplitN(value, "\n", 2)

		// Parse JSON from the message
		var m map[string]interface{}
		err := json.Unmarshal([]byte(msg[1]), &m)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		// Extract values from JSON
		level, ok := m["level"].(string)
		if !ok {
			level = ""
		}
		time, ok := m["time"].(string)
		if !ok {
			time = ""
		}
		delete(m, "time")
		delete(m, "level")
		newMsg, _ := json.Marshal(m)

		// Write data to Excel
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), msg[0])
		file.SetColWidth("Sheet1", "A", "A", float64(len(msg[0])))

		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), level)
		file.SetColWidth("Sheet1", "B", "B", float64(len(level)))

		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), string(newMsg))
		file.SetColWidth("Sheet1", "C", "C", float64(len(string(newMsg))))

		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), time)
		file.SetColWidth("Sheet1", "D", "D", float64(len(time)))

		row++
	}

	// Set active sheet of the workbook
	file.SetActiveSheet(index)

	// Save the Excel file
	err = file.SaveAs("Search_result.xlsx")
	if err != nil {
		log.Fatal("Cannot save Excel file:", err)
	}

}
