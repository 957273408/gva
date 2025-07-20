package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/api"

var (
	Router    = new(router)
	apiInfo   = api.Api.Info
	apiHero   = api.Api.Hero
	apiTeam   = api.Api.Team
	apiPlayer = api.Api.Player
	apiBplist = api.Api.Bplist
)

type router struct {
	Info   info
	Hero   hero
	Team   team
	Player player
	Bplist bplist
}
