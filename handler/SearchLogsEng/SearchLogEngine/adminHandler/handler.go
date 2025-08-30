package adminHandler

import (
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/service/adminService"
)

type Handler struct {
	adminSvc adminService.Service
}

func New(ASvc adminService.Service) *Handler {
	return &Handler{
		adminSvc: ASvc,
	}
}
