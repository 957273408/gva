package router

import (
	"github.com/gin-gonic/gin"
)

var Bplist = new(bplist)

type bplist struct{}

// Init 初始化 bp记录 路由信息
func (r *bplist) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := public.Group("bplist")
		group.POST("createBplist", apiBplist.CreateBplist)             // 新建bp记录
		group.DELETE("deleteBplist", apiBplist.DeleteBplist)           // 删除bp记录
		group.DELETE("deleteBplistByIds", apiBplist.DeleteBplistByIds) // 批量删除bp记录
		group.PUT("updateBplist", apiBplist.UpdateBplist)              // 更新bp记录
	}
	{
		group := public.Group("bplist")
		group.GET("findBplist", apiBplist.FindBplist)       // 根据ID获取bp记录
		group.GET("getBplistList", apiBplist.GetBplistList) // 获取bp记录列表
	}
	{
		group := public.Group("bplist")
		group.GET("getBplistPublic", apiBplist.GetBplistPublic) // bp记录开放接口
	}
}
