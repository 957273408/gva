package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Team = new(team)

type team struct {}

// Init 初始化 战队 路由信息
func (r *team) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("team").Use(middleware.OperationRecord())
		group.POST("createTeam", apiTeam.CreateTeam)   // 新建战队
		group.DELETE("deleteTeam", apiTeam.DeleteTeam) // 删除战队
		group.DELETE("deleteTeamByIds", apiTeam.DeleteTeamByIds) // 批量删除战队
		group.PUT("updateTeam", apiTeam.UpdateTeam)    // 更新战队
	}
	{
	    group := private.Group("team")
		group.GET("findTeam", apiTeam.FindTeam)        // 根据ID获取战队
		group.GET("getTeamList", apiTeam.GetTeamList)  // 获取战队列表
	}
	{
	    group := public.Group("team")
	    group.GET("getTeamPublic", apiTeam.GetTeamPublic)  // 战队开放接口
	}
}
