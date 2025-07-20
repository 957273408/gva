package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Hero = new(hero)

type hero struct {}

// Init 初始化 英雄 路由信息
func (r *hero) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("hero").Use(middleware.OperationRecord())
		group.POST("createHero", apiHero.CreateHero)   // 新建英雄
		group.DELETE("deleteHero", apiHero.DeleteHero) // 删除英雄
		group.DELETE("deleteHeroByIds", apiHero.DeleteHeroByIds) // 批量删除英雄
		group.PUT("updateHero", apiHero.UpdateHero)    // 更新英雄
	}
	{
	    group := private.Group("hero")
		group.GET("findHero", apiHero.FindHero)        // 根据ID获取英雄
		group.GET("getHeroList", apiHero.GetHeroList)  // 获取英雄列表
	}
	{
	    group := public.Group("hero")
	    group.GET("getHeroPublic", apiHero.GetHeroPublic)  // 英雄开放接口
	}
}
