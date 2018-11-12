import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'
import isInteger from 'lodash/isInteger'

// 获取币种维护列表
export function list (params) {
  return request({
    url: requestUrl('/transaction/bzwh/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取币种信息
export function info (id) {
  return request({
    url: requestUrl('/transaction/bzwh/info' + (isInteger(id) ? `/${id}` : '')),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 添加币种数据
export function add (params) {
  return request({
    url: requestUrl('/transaction/bzwh/save'),
    method: 'post',
    data: requestParam(params)
  })
}

// 修改币种数据
export function update (params) {
  return request({
    url: requestUrl('/transaction/bzwh/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除币种数据
export function del (params) {
  return request({
    url: requestUrl('/transaction/bzwh/delete'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}
