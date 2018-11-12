import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'
import isInteger from 'lodash/isInteger'

// 获取用户列表
export function list (params) {
  return request({
    url: requestUrl('/sys/user/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取用户信息
export function info (id) {
  return request({
    url: requestUrl('/sys/user/info' + (isInteger(id) ? `/${id}` : '')),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 修改密码
export function updatePassword (params) {
  return request({
    url: requestUrl('/sys/user/password'),
    method: 'post',
    data: requestParam(params)
  })
}

// 添加用户
export function add (params) {
  return request({
    url: requestUrl('/sys/user/save'),
    method: 'post',
    data: requestParam(params)
  })
}

// 修改用户
export function update (params) {
  return request({
    url: requestUrl('/sys/user/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除用户
export function del (params) {
  return request({
    url: requestUrl('/sys/user/delete'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 获取子链管理列表
export function list1 (params) {
  return request({
    url: requestUrl('/back/subchain/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取子链详细信息
export function info1 (id) {
  return request({
    url: requestUrl('/back/subchain/info' + (isInteger(id) ? `/${id}` : '')),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 新增存储字链
export function add1 (params) {
  return request({
    url: requestUrl('/back/subchain/add'),
    method: 'post',
    data: requestParam(params)
  })
}

// 修改存储子链
export function update1 (params) {
  return request({
    url: requestUrl('/back/subchain/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除存储子链
export function del1 (params) {
  return request({
    url: requestUrl('/back/subchain/delete'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除存储子链
export function logList (params) {
  return request({
    url: requestUrl('/back/subchainlog/list'),
    method: 'get',
    data: requestParam(params, 'get')
  })
}
