package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Player = new(player)

type player struct {}

// Init 初始化 选手 路由信息
func (r *player) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("player").Use(middleware.OperationRecord())
		group.POST("createPlayer", apiPlayer.CreatePlayer)   // 新建选手
		group.DELETE("deletePlayer", apiPlayer.DeletePlayer) // 删除选手
		group.DELETE("deletePlayerByIds", apiPlayer.DeletePlayerByIds) // 批量删除选手
		group.PUT("updatePlayer", apiPlayer.UpdatePlayer)    // 更新选手
	}
	{
	    group := private.Group("player")
		group.GET("findPlayer", apiPlayer.FindPlayer)        // 根据ID获取选手
		group.GET("getPlayerList", apiPlayer.GetPlayerList)  // 获取选手列表
	}
	{
	    group := public.Group("player")
	    group.GET("getPlayerPublic", apiPlayer.GetPlayerPublic)  // 选手开放接口
	}
}
