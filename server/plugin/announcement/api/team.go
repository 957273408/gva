package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var Team = new(team)

type team struct {}

// CreateTeam 创建战队
// @Tags Team
// @Summary 创建战队
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Team true "创建战队"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /team/createTeam [post]
func (a *team) CreateTeam(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Team
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceTeam.CreateTeam(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTeam 删除战队
// @Tags Team
// @Summary 删除战队
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Team true "删除战队"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /team/deleteTeam [delete]
func (a *team) DeleteTeam(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceTeam.DeleteTeam(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteTeamByIds 批量删除战队
// @Tags Team
// @Summary 批量删除战队
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /team/deleteTeamByIds [delete]
func (a *team) DeleteTeamByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceTeam.DeleteTeamByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateTeam 更新战队
// @Tags Team
// @Summary 更新战队
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Team true "更新战队"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /team/updateTeam [put]
func (a *team) UpdateTeam(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Team
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceTeam.UpdateTeam(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindTeam 用id查询战队
// @Tags Team
// @Summary 用id查询战队
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询战队"
// @Success 200 {object} response.Response{data=model.Team,msg=string} "查询成功"
// @Router /team/findTeam [get]
func (a *team) FindTeam(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reteam, err := serviceTeam.GetTeam(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reteam, c)
}
// GetTeamList 分页获取战队列表
// @Tags Team
// @Summary 分页获取战队列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.TeamSearch true "分页获取战队列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /team/getTeamList [get]
func (a *team) GetTeamList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.TeamSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceTeam.GetTeamInfoList(ctx,pageInfo)
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
// GetTeamPublic 不需要鉴权的战队接口
// @Tags Team
// @Summary 不需要鉴权的战队接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /team/getTeamPublic [get]
func (a *team) GetTeamPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceTeam.GetTeamPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的战队接口信息"}, "获取成功", c)
}
