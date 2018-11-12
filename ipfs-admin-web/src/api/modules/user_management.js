import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'
import isInteger from 'lodash/isInteger'
import isString from 'lodash/isstring'

// 获取用户列表
export function list (params) {
  return request({
    url: requestUrl('/yhgl/yhgl/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 根据邮箱查询用户信息
export function getEmailInfo (cmcEmail) {
  return request({
    url: requestUrl('/yhgl/yhgl/infoByEmail' + (isString(cmcEmail) ? `/${cmcEmail}` : '')),
    method: 'get',
    params: requestParam(cmcEmail, 'get')
  })
}

// 获取用户信息
export function info (id) {
  return request({
    url: requestUrl('/yhgl/yhgl/info' + (isInteger(id) ? `/${id}` : '')),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 修改vip用户
export function update (params) {
  return request({
    url: requestUrl('/yhgl/yhgl/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 获取资产地址
export function getZcInfo (cmcYhglId) {
  return request({
    url: requestUrl('/yhgl/yhzc/listByYhgl' + (isInteger(cmcYhglId) ? `/${cmcYhglId}` : '')),
    method: 'get',
    params: requestParam(cmcYhglId, 'get')
  })
}

// 充值资产
export function updateZc (params) {
  return request({
    url: requestUrl('/yhgl/yhzc/updateByUser'),
    method: 'post',
    data: requestParam(params)
  })
}

// 获取钱包地址
export function getYhdzInfo (cmcYhglId) {
  return request({
    url: requestUrl('/yhgl/yhdz/listByYhgl' + (isInteger(cmcYhglId) ? `/${cmcYhglId}` : '')),
    method: 'get',
    params: requestParam(cmcYhglId, 'get')
  })
}

// 获取用户通讯录
export function getTxlInfo (cmcYhglId) {
  return request({
    url: requestUrl('/yhgl/yhglmail/listByYhgl' + (isInteger(cmcYhglId) ? `/${cmcYhglId}` : '')),
    method: 'get',
    params: requestParam(cmcYhglId, 'get')
  })
}
