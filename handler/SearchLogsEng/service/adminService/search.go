package adminService

import (
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/params"
)

func (s *Service) GetFilesInFolder(SLR params.SearchLogRequest) ([]string, int64, error) {
	Data, Counter, err := s.SearchLogs.GetFilesInFolder(SLR)
	if err != nil {
		return nil, 0, err
	}
	return Data, Counter, nil
}
func (s *Service) GetFilesInFolderInternal(SLR params.SearchLogRequest) ([]string, int64, error) {
	Data, Counter, err := s.SearchLogs.GetFilesInFolder(SLR)
	if err != nil {
		return nil, 0, err
	}
	return Data, Counter, nil
}
