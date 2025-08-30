package adminService

import (
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/params"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/search"
)

type SearchLogs interface {
	GetFilesInFolder(SLR params.SearchLogRequest) ([]string, int64, error)
}

type Service struct {
	SearchLogs SearchLogs
}

func New(SL *search.SearchLogs) *Service {
	return &Service{
		SearchLogs: SL,
	}
}
