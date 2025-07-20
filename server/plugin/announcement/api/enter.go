package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/service"

var (
	Api           = new(api)
	serviceInfo   = service.Service.Info
	serviceHero   = service.Service.Hero
	serviceTeam   = service.Service.Team
	servicePlayer = service.Service.Player
	serviceBplist = service.Service.Bplist
)

type api struct {
	Info   info
	Hero   hero
	Team   team
	Player player
	Bplist bplist
}
