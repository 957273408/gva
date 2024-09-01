package first

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type KhinfoRouter struct{}

// InitKhinfoRouter 初始化 客户信息 路由信息
func (s *KhinfoRouter) InitKhinfoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	khinfoRouter := Router.Group("khinfo").Use(middleware.OperationRecord())
	khinfoRouterWithoutRecord := Router.Group("khinfo")
	khinfoRouterWithoutAuth := PublicRouter.Group("khinfo")
	{
		khinfoRouter.POST("createKhinfo", khinfoApi.CreateKhinfo)             // 新建客户信息
		khinfoRouter.DELETE("deleteKhinfo", khinfoApi.DeleteKhinfo)           // 删除客户信息
		khinfoRouter.DELETE("deleteKhinfoByIds", khinfoApi.DeleteKhinfoByIds) // 批量删除客户信息
		khinfoRouter.PUT("updateKhinfo", khinfoApi.UpdateKhinfo)              // 更新客户信息
	}
	{
		khinfoRouterWithoutRecord.GET("findKhinfo", khinfoApi.FindKhinfo)       // 根据ID获取客户信息
		khinfoRouterWithoutRecord.GET("getKhinfoList", khinfoApi.GetKhinfoList) // 获取客户信息列表
	}
	{
		khinfoRouterWithoutAuth.POST("createKhinfoPublic", khinfoApi.CreateKhinfo) // 获取客户信息列表
	}
}
