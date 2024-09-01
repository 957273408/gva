import service from '@/utils/request'

// @Tags Khinfo
// @Summary 创建客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Khinfo true "创建客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /khinfo/createKhinfo [post]
export const createKhinfo = (data) => {
  return service({
    url: '/khinfo/createKhinfo',
    method: 'post',
    data
  })
}

// @Tags Khinfo
// @Summary 删除客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Khinfo true "删除客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /khinfo/deleteKhinfo [delete]
export const deleteKhinfo = (params) => {
  return service({
    url: '/khinfo/deleteKhinfo',
    method: 'delete',
    params
  })
}

// @Tags Khinfo
// @Summary 批量删除客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /khinfo/deleteKhinfo [delete]
export const deleteKhinfoByIds = (params) => {
  return service({
    url: '/khinfo/deleteKhinfoByIds',
    method: 'delete',
    params
  })
}

// @Tags Khinfo
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Khinfo true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /khinfo/updateKhinfo [put]
export const updateKhinfo = (data) => {
  return service({
    url: '/khinfo/updateKhinfo',
    method: 'put',
    data
  })
}

// @Tags Khinfo
// @Summary 用id查询客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Khinfo true "用id查询客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /khinfo/findKhinfo [get]
export const findKhinfo = (params) => {
  return service({
    url: '/khinfo/findKhinfo',
    method: 'get',
    params
  })
}

// @Tags Khinfo
// @Summary 分页获取客户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取客户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /khinfo/getKhinfoList [get]
export const getKhinfoList = (params) => {
  return service({
    url: '/khinfo/getKhinfoList',
    method: 'get',
    params
  })
}
