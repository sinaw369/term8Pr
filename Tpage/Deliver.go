package Tpage

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/SearchLogEngine"
)

var Del *Deliver

type Deliver struct {
	SLHa SearchLogEngine.SearchLogServer
}

func New(Sl SearchLogEngine.SearchLogServer) *Deliver {
	Del = &Deliver{SLHa: Sl}
	return Del
}
func (d Deliver) SearchLogHandler(c *context.Context) ([]string, int64, error) {
	Data, JustCount, err := d.SLHa.AdminHandler.SearchLogHandlerInternal(c)
	if err != nil {
		return nil, 0, err
	}
	return Data, JustCount, nil
}
func DeliverTpage() *Deliver {
	return Del

}
