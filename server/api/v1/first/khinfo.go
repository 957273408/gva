package first

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/first"
    firstReq "github.com/flipped-aurora/gin-vue-admin/server/model/first/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type KhinfoApi struct {}

// CreateKhinfo 创建客户信息
// @Tags Khinfo
// @Summary 创建客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body first.Khinfo true "创建客户信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /khinfo/createKhinfo [post]
func (khinfoApi *KhinfoApi) CreateKhinfo(c *gin.Context) {
	var khinfo first.Khinfo
	err := c.ShouldBindJSON(&khinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    khinfo.CreatedBy = utils.GetUserID(c)
	err = khinfoService.CreateKhinfo(&khinfo)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteKhinfo 删除客户信息
// @Tags Khinfo
// @Summary 删除客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body first.Khinfo true "删除客户信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /khinfo/deleteKhinfo [delete]
func (khinfoApi *KhinfoApi) DeleteKhinfo(c *gin.Context) {
	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := khinfoService.DeleteKhinfo(ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteKhinfoByIds 批量删除客户信息
// @Tags Khinfo
// @Summary 批量删除客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /khinfo/deleteKhinfoByIds [delete]
func (khinfoApi *KhinfoApi) DeleteKhinfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := khinfoService.DeleteKhinfoByIds(IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateKhinfo 更新客户信息
// @Tags Khinfo
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body first.Khinfo true "更新客户信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /khinfo/updateKhinfo [put]
func (khinfoApi *KhinfoApi) UpdateKhinfo(c *gin.Context) {
	var khinfo first.Khinfo
	err := c.ShouldBindJSON(&khinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    khinfo.UpdatedBy = utils.GetUserID(c)
	err = khinfoService.UpdateKhinfo(khinfo)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindKhinfo 用id查询客户信息
// @Tags Khinfo
// @Summary 用id查询客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query first.Khinfo true "用id查询客户信息"
// @Success 200 {object} response.Response{data=first.Khinfo,msg=string} "查询成功"
// @Router /khinfo/findKhinfo [get]
func (khinfoApi *KhinfoApi) FindKhinfo(c *gin.Context) {
	ID := c.Query("ID")
	rekhinfo, err := khinfoService.GetKhinfo(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rekhinfo, c)
}

// GetKhinfoList 分页获取客户信息列表
// @Tags Khinfo
// @Summary 分页获取客户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query firstReq.KhinfoSearch true "分页获取客户信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /khinfo/getKhinfoList [get]
func (khinfoApi *KhinfoApi) GetKhinfoList(c *gin.Context) {
	var pageInfo firstReq.KhinfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := khinfoService.GetKhinfoInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetKhinfoPublic 不需要鉴权的客户信息接口
// @Tags Khinfo
// @Summary 不需要鉴权的客户信息接口
// @accept application/json
// @Produce application/json
// @Param data query firstReq.KhinfoSearch true "分页获取客户信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /khinfo/getKhinfoPublic [get]
func (khinfoApi *KhinfoApi) GetKhinfoPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的客户信息接口信息",
    }, "获取成功", c)
}
