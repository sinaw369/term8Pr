package Tpage

import (
	"encoding/json"
	"fmt"
	"strings"
)

func FixDataForShowInGoAdminSearchLog(Data []string, JustCount int64) ([]map[string]interface{}, int) {

	var sliceOfMaps []map[string]interface{}
	var DataMaps map[string]interface{}
	var Idx int = 1
	if Data != nil {
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
				level = "None"
			}
			time, ok := m["time"].(string)
			if !ok {
				time = ""
			}
			delete(m, "time")
			delete(m, "level")
			newMsg, _ := json.Marshal(m)
			DataMaps = map[string]interface{}{
				"id":       Idx,
				"filename": msg[0],
				"level":    level,
				"message":  string(newMsg),
				"time":     time,
			}
			// Write data to Excel
			sliceOfMaps = append(sliceOfMaps, DataMaps)
			Idx++
		}

	} else {
		DataMaps = map[string]interface{}{
			"id":             Idx,
			"numberOfResult": int(JustCount),
		}
		sliceOfMaps = append(sliceOfMaps, DataMaps)
	}
	return sliceOfMaps, len(sliceOfMaps)

}
