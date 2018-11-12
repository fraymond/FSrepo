import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'
import isInteger from 'lodash/isInteger'

// 获取轮播图列表
export function list (params) {
  return request({
    url: requestUrl('/sysManagement/banner/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取轮播图信息
export function info (id) {
  return request({
    url: requestUrl('/sysManagement/banner/info' + (isInteger(id) ? `/${id}` : '')),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 添加轮播图
export function add (params) {
  return request({
    url: requestUrl('/sysManagement/banner/save'),
    method: 'post',
    data: requestParam(params)
  })
}

// 修改轮播图
export function update (params) {
  return request({
    url: requestUrl('/sysManagement/banner/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除轮播图
export function del (params) {
  return request({
    url: requestUrl('/sysManagement/banner/delete'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}
