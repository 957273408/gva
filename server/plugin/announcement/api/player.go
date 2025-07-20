package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var Player = new(player)

type player struct {}

// CreatePlayer 创建选手
// @Tags Player
// @Summary 创建选手
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Player true "创建选手"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /player/createPlayer [post]
func (a *player) CreatePlayer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Player
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = servicePlayer.CreatePlayer(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePlayer 删除选手
// @Tags Player
// @Summary 删除选手
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Player true "删除选手"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /player/deletePlayer [delete]
func (a *player) DeletePlayer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := servicePlayer.DeletePlayer(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeletePlayerByIds 批量删除选手
// @Tags Player
// @Summary 批量删除选手
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /player/deletePlayerByIds [delete]
func (a *player) DeletePlayerByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := servicePlayer.DeletePlayerByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdatePlayer 更新选手
// @Tags Player
// @Summary 更新选手
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Player true "更新选手"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /player/updatePlayer [put]
func (a *player) UpdatePlayer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Player
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = servicePlayer.UpdatePlayer(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindPlayer 用id查询选手
// @Tags Player
// @Summary 用id查询选手
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询选手"
// @Success 200 {object} response.Response{data=model.Player,msg=string} "查询成功"
// @Router /player/findPlayer [get]
func (a *player) FindPlayer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	replayer, err := servicePlayer.GetPlayer(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(replayer, c)
}
// GetPlayerList 分页获取选手列表
// @Tags Player
// @Summary 分页获取选手列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PlayerSearch true "分页获取选手列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /player/getPlayerList [get]
func (a *player) GetPlayerList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.PlayerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := servicePlayer.GetPlayerInfoList(ctx,pageInfo)
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
// GetPlayerPublic 不需要鉴权的选手接口
// @Tags Player
// @Summary 不需要鉴权的选手接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /player/getPlayerPublic [get]
func (a *player) GetPlayerPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    servicePlayer.GetPlayerPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的选手接口信息"}, "获取成功", c)
}
