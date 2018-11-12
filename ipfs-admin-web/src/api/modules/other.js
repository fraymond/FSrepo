import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'
import isInteger from 'lodash/isInteger'

// 获取公告栏列表
export function noticeList (params) {
  return request({
    url: requestUrl('/sysManagement/notice/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取公告信息
export function noticeInfo (id) {
  return request({
    url: requestUrl('/sysManagement/notice/info' + (isInteger(id) ? `/${id}` : '')),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 添加公告
export function noticeAdd (params) {
  return request({
    url: requestUrl('/sysManagement/notice/save'),
    method: 'post',
    data: requestParam(params)
  })
}

// 修改公告
export function noticeUpdate (params) {
  return request({
    url: requestUrl('/sysManagement/notice/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除公告
export function noticeDel (params) {
  return request({
    url: requestUrl('/sysManagement/notice/delete'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 获取留言板列表
export function messageList (params) {
  return request({
    url: requestUrl('/sysManagement/yhglmessage/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取留言信息
export function messageInfo (id) {
  return request({
    url: requestUrl('/sysManagement/yhglmessage/info' + (isInteger(id) ? `/${id}` : '')),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 添加留言
export function messageAdd (params) {
  return request({
    url: requestUrl('/sysManagement/yhglmessage/save'),
    method: 'post',
    data: requestParam(params)
  })
}

// 修改留言
export function messageUpdate (params) {
  return request({
    url: requestUrl('/sysManagement/yhglmessage/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除留言
export function messageDel (params) {
  return request({
    url: requestUrl('/sysManagement/yhglmessage/delete'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 获取轮播图列表
export function bannerList (params) {
  return request({
    url: requestUrl('/sysManagement/banner/list'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取轮播图信息
export function bannerInfo (id) {
  return request({
    url: requestUrl('/sysManagement/banner/info' + (isInteger(id) ? `/${id}` : '')),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 添加轮播图
export function bannerAdd (params) {
  return request({
    url: requestUrl('/sysManagement/banner/save'),
    method: 'post',
    data: requestParam(params)
  })
}

// 修改轮播图
export function bannerUpdate (params) {
  return request({
    url: requestUrl('/sysManagement/banner/update'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除轮播图
export function bannerDel (params) {
  return request({
    url: requestUrl('/sysManagement/banner/delete'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}
