import service from '@/utils/request'
// @Tags Bplist
// @Summary 创建bp记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Bplist true "创建bp记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bplist/createBplist [post]
export const createBplist = (data) => {
  return service({
    url: '/bplist/createBplist',
    method: 'post',
    data
  })
}

// @Tags Bplist
// @Summary 删除bp记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Bplist true "删除bp记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bplist/deleteBplist [delete]
export const deleteBplist = (params) => {
  return service({
    url: '/bplist/deleteBplist',
    method: 'delete',
    params
  })
}

// @Tags Bplist
// @Summary 批量删除bp记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除bp记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bplist/deleteBplist [delete]
export const deleteBplistByIds = (params) => {
  return service({
    url: '/bplist/deleteBplistByIds',
    method: 'delete',
    params
  })
}

// @Tags Bplist
// @Summary 更新bp记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Bplist true "更新bp记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bplist/updateBplist [put]
export const updateBplist = (data) => {
  return service({
    url: '/bplist/updateBplist',
    method: 'put',
    data
  })
}

// @Tags Bplist
// @Summary 用id查询bp记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Bplist true "用id查询bp记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bplist/findBplist [get]
export const findBplist = (params) => {
  return service({
    url: '/bplist/findBplist',
    method: 'get',
    params
  })
}

// @Tags Bplist
// @Summary 分页获取bp记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取bp记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bplist/getBplistList [get]
export const getBplistList = (params) => {
  return service({
    url: '/bplist/getBplistList',
    method: 'get',
    params
  })
}
// @Tags Bplist
// @Summary 不需要鉴权的bp记录接口
// @Accept application/json
// @Produce application/json
// @Param data query request.BplistSearch true "分页获取bp记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bplist/getBplistPublic [get]
export const getBplistPublic = () => {
  return service({
    url: '/bplist/getBplistPublic',
    method: 'get',
  })
}
