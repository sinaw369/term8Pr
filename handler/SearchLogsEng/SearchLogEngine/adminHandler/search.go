package adminHandler

import (
	"encoding/json"
	"fmt"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/params"
	"net/http"
	"time"

	//"encoding/json"
	//	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"strconv"
	"strings"
)

func (Ha *Handler) SearchLogHandlerExternal(c *context.Context) {
	var limit int
	// Convert the limit value to an integer
	if strings.TrimSpace(c.FormValue("limit")) != "" {
		limit, _ = strconv.Atoi(c.FormValue("limit"))
	} else {
		limit = 0
	}
	FormReq := params.SearchLogRequest{
		Year:                c.FormValue("year"),
		Month:               c.FormValue("month"),
		Day:                 c.FormValue("day"),
		LogType:             c.FormValue("logType"),
		Limit:               limit,
		SearchKey:           strings.TrimSpace(c.FormValue("searchKey")),
		FileName:            strings.TrimSpace(c.FormValue("fileName")),
		NotIncludeFileName:  strings.TrimSpace(c.FormValue("notIncludeFileName")),
		NotIncludeSearchKey: strings.TrimSpace(c.FormValue("notIncludeSearchKey")),
		CheckBox:            false,
	}
	if c.FormValue("countOnly") == "on" {
		FormReq.CheckBox = true
	}
	Data, responseCounter, _ := Ha.adminSvc.GetFilesInFolder(FormReq)

	searchResponse := params.SearchLogResponse{
		Message: Data, // You can customize the response message
		Count:   fmt.Sprintf(" تعداد نتایج یافت شده:%d", responseCounter),
	}

	// Marshal the SearchResponse struct into JSON format
	response, err := json.Marshal(&searchResponse)
	if err != nil {
		//return c.JSON(http.StatusInternalServerError, map[string]interface{"error": "Internal Server Error"})
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Internal Server Error"})

		return
	}
	// Set Content-Type header
	c.Response.Header.Set("Content-Type", "application/json")

	//go func() {
	//	DownloadExecl.CreateExcel(Data)
	//}()
	c.HTML(http.StatusOK, string(response))
}
func (Ha *Handler) SearchLogHandlerInternal(c *context.Context) ([]string, int64, error) {
	var limit int
	// Convert the limit value to an integer
	if strings.TrimSpace(c.FormValue("Limit")) != "" {
		limit, _ = strconv.Atoi(c.FormValue("Limit"))
	} else {
		limit = 100
	}
	dateString := c.FormValue("Date")

	layout := "2006-01-02"
	// Parse the date string into a time.Time value
	t, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, 0, err
	}

	// Extract the year, month, and day
	year := strconv.Itoa(t.Year())
	month := strconv.Itoa(int(t.Month()))
	if int(t.Month()) < 10 {
		month = "0" + month
	}
	day := strconv.Itoa(t.Day())
	if t.Day() < 10 {
		day = "0" + day
	}
	FormReq := params.SearchLogRequest{
		Year:                year,
		Month:               month,
		Day:                 day,
		LogType:             c.FormValue("LogType"),
		Limit:               limit,
		SearchKey:           strings.TrimSpace(c.FormValue("searchKey")),
		FileName:            strings.TrimSpace(c.FormValue("FileName")),
		NotIncludeFileName:  strings.TrimSpace(c.FormValue("notIncludeFileName")),
		NotIncludeSearchKey: strings.TrimSpace(c.FormValue("notIncludeSearchKey")),
		CheckBox:            false,
	}
	if c.FormValue("justCount") == "1" {
		FormReq.CheckBox = true
	}
	//fmt.Printf("%+v", FormReq)
	//FormReq := params.SearchLogRequest{
	//	Year:                "2024",
	//	Month:               "01",
	//	Day:                 "01",
	//	LogType:             "",
	//	Limit:               100,
	//	SearchKey:           "",
	//	FileName:            "",
	//	NotIncludeFileName:  "",
	//	NotIncludeSearchKey: "",
	//	CheckBox:            false,
	//}
	Data, counter, _ := Ha.adminSvc.GetFilesInFolder(FormReq)
	//searchResponse := params.SearchLogResponse{
	//	Message: Data, // You can customize the response message
	//	Count:   fmt.Sprintf(" تعداد نتایج یافت شده:%d", responseCounter),
	//}

	// Marshal the SearchResponse struct into JSON format
	//response, err := json.Marshal(&searchResponse)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	//}
	// Set Content-Type header
	return Data, counter, nil
}
