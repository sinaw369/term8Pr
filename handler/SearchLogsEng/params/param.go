package params

type SearchLogRequest struct {
	Year                string `json:"year"`
	Month               string `json:"month"`
	Day                 string `json:"day"`
	LogType             string `json:"log_type"`
	Limit               int    `json:"limit"`
	SearchKey           string `json:"search_key"`
	FileName            string `json:"file_name"`
	NotIncludeFileName  string `json:"not_include_file_name"`
	NotIncludeSearchKey string `json:"not_include_search_key"`
	CheckBox            bool   `json:"check_box"`
}
type SearchLogResponse struct {
	Message []string `json:"message"`
	Count   string   `json:"count"`
}

// LogEntry represents a log entry.
type LogEntry struct {
	Time string `json:"time"`
}

const (
	EmptyString    = ""
	DefultReadLine = 2000
)
