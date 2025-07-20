import service from '@/utils/request'
// @Tags Player
// @Summary 创建选手
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Player true "创建选手"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /player/createPlayer [post]
export const createPlayer = (data) => {
  return service({
    url: '/player/createPlayer',
    method: 'post',
    data
  })
}

// @Tags Player
// @Summary 删除选手
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Player true "删除选手"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /player/deletePlayer [delete]
export const deletePlayer = (params) => {
  return service({
    url: '/player/deletePlayer',
    method: 'delete',
    params
  })
}

// @Tags Player
// @Summary 批量删除选手
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除选手"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /player/deletePlayer [delete]
export const deletePlayerByIds = (params) => {
  return service({
    url: '/player/deletePlayerByIds',
    method: 'delete',
    params
  })
}

// @Tags Player
// @Summary 更新选手
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Player true "更新选手"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /player/updatePlayer [put]
export const updatePlayer = (data) => {
  return service({
    url: '/player/updatePlayer',
    method: 'put',
    data
  })
}

// @Tags Player
// @Summary 用id查询选手
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Player true "用id查询选手"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /player/findPlayer [get]
export const findPlayer = (params) => {
  return service({
    url: '/player/findPlayer',
    method: 'get',
    params
  })
}

// @Tags Player
// @Summary 分页获取选手列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取选手列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /player/getPlayerList [get]
export const getPlayerList = (params) => {
  return service({
    url: '/player/getPlayerList',
    method: 'get',
    params
  })
}
// @Tags Player
// @Summary 不需要鉴权的选手接口
// @Accept application/json
// @Produce application/json
// @Param data query request.PlayerSearch true "分页获取选手列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /player/getPlayerPublic [get]
export const getPlayerPublic = () => {
  return service({
    url: '/player/getPlayerPublic',
    method: 'get',
  })
}
