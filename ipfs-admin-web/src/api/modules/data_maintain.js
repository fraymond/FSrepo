import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'
import isInteger from 'lodash/isInteger'

// 获取用户列表
export function list (params) {
  return request({
    url: requestUrl('/transaction/currency/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取用户信息
export function info (id) {
  return request({
    url: requestUrl('/transaction/currency/info/' + (isInteger(id) ? `/${id}` : '')),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 添加用户
export function add (params) {
  return request({
    url: requestUrl('/transaction/currency/save'),
    method: 'post',
    data: requestParam(params)
  })
}

// 修改用户
export function update (params) {
  return request({
    url: requestUrl('/transaction/currency/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除用户
export function del (params) {
  return request({
    url: requestUrl('/transaction/currency/delete'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 修改交易货币
export function updateJy (params) {
  return request({
    url: requestUrl('/transaction/currency/updateJy'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 修改交易货币
export function updateJyZt (params) {
  return request({
    url: requestUrl('/transaction/currency/updateJyZt'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}
