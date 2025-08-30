package adminHandler

import (
	"github.com/labstack/echo/v4"
	"os"
)

func (Ha *Handler) DownloadFile(c echo.Context) error {
	// Set response header to indicate that it's an attachment and specify the file name
	c.Response().Header().Set("Content-Disposition", "attachment; filename=Search_result.xlsx")
	// Set content type for Excel files
	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	// Serve the file
	err := c.File("Search_result.xlsx")
	if err != nil {
		return err
	}

	// Delete the file after it has been served
	if err := os.Remove("Search_result.xlsx"); err != nil {
		return err
	}

	return nil
}
