import service from '@/utils/request'
// @Tags Team
// @Summary 创建战队
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Team true "创建战队"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /team/createTeam [post]
export const createTeam = (data) => {
  return service({
    url: '/team/createTeam',
    method: 'post',
    data
  })
}

// @Tags Team
// @Summary 删除战队
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Team true "删除战队"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /team/deleteTeam [delete]
export const deleteTeam = (params) => {
  return service({
    url: '/team/deleteTeam',
    method: 'delete',
    params
  })
}

// @Tags Team
// @Summary 批量删除战队
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除战队"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /team/deleteTeam [delete]
export const deleteTeamByIds = (params) => {
  return service({
    url: '/team/deleteTeamByIds',
    method: 'delete',
    params
  })
}

// @Tags Team
// @Summary 更新战队
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Team true "更新战队"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /team/updateTeam [put]
export const updateTeam = (data) => {
  return service({
    url: '/team/updateTeam',
    method: 'put',
    data
  })
}

// @Tags Team
// @Summary 用id查询战队
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Team true "用id查询战队"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /team/findTeam [get]
export const findTeam = (params) => {
  return service({
    url: '/team/findTeam',
    method: 'get',
    params
  })
}

// @Tags Team
// @Summary 分页获取战队列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取战队列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /team/getTeamList [get]
export const getTeamList = (params) => {
  return service({
    url: '/team/getTeamList',
    method: 'get',
    params
  })
}
// @Tags Team
// @Summary 不需要鉴权的战队接口
// @Accept application/json
// @Produce application/json
// @Param data query request.TeamSearch true "分页获取战队列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /team/getTeamPublic [get]
export const getTeamPublic = () => {
  return service({
    url: '/team/getTeamPublic',
    method: 'get',
  })
}
