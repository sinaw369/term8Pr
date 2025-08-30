package SearchLogEngine

import (
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/SearchLogEngine/adminHandler"
)

type SearchLogServer struct {
	AdminHandler adminHandler.Handler
}

func New(aHandler adminHandler.Handler) SearchLogServer {
	return SearchLogServer{
		AdminHandler: aHandler,
	}
}

//func (s Server) Serve() {
//
//	s.Router.GET("/health-check", s.healthCheck)
//	s.Router.POST("/search", s.adminHandler.SearchLogHandler)
//	s.Router.Static("/", "ui/.")
//	s.Router.GET("/download", s.adminHandler.DownloadFile) // New route for downloading the file
//	Start server
//	address := fmt.Sprintf(":%d", 3010)
//	fmt.Printf("start echo server on %s\n", address)
//	if err := s.Router.Start(address); err != nil {
//		fmt.Println("router start error", err)
//	}
//}
