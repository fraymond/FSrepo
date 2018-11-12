import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'
// import isInteger from 'lodash/isInteger'

// 获取轮播图数据
export function getBanner (params) {
  return request({
    url: requestUrl('web/home/getBanner'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取对标代币列表
export function getCurrRelation () {
  return request({
    url: requestUrl('web/home/getCurrRelation'),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 获取币种介绍列表
export function getBzwh () {
  return request({
    url: requestUrl('web/home/getBzwh'),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 获取币种费率列表
export function getRate () {
  return request({
    url: requestUrl('web/home/getRate'),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 获取主页币种交易数据
export function getAllJYCoin (params) {
  return request({
    url: requestUrl('web/home/getAllJYCoin'),
    method: 'post',
    data: requestParam(params)
  })
}

// 获取公告
export function getNotice () {
  return request({
    url: requestUrl('web/home/getNotice'),
    method: 'get',
    params: requestParam({}, 'get')
  })
}
