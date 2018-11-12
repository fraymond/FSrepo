import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'
//  import isInteger from 'lodash/isInteger'

// 获取主页所有对标代币列表
export function getCurrRelation () {
  return request({
    url: requestUrl('/web/trade/getCurrRelation'),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 获取可交易币种列表
export function getTradeCurrency () {
  return request({
    url: requestUrl('/web/trade/getTradeCurrency'),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 获取主页页面查询委托订单列表
export function getUserAssets (params) {
  return request({
    url: requestUrl('/web/trade/getUserAssets'),
    method: 'post',
    data: requestParam(params)
  })
}

// 获取交易页面查询委托订单列表
export function getTransBill (params) {
  return request({
    url: requestUrl('/web/trade/getTransBill'),
    method: 'post',
    data: requestParam(params)
  })
}

// 买卖单提交
export function addTransBill (params) {
  return request({
    url: requestUrl('/web/trade/addTransBill'),
    method: 'post',
    data: requestParam(params)
  })
}

// 获取主页币种的交易数据
export function getAllJYCoin (params) {
  return request({
    url: requestUrl('/web/trade/getAllJYCoin'),
    method: 'post',
    data: requestParam(params)
  })
}

// 撤销当前委托
export function transBillRevoke (params) {
  return request({
    url: requestUrl('/web/trade/transBillRevoke'),
    method: 'post',
    data: requestParam(params)
  })
}

// 全部撤销
export function transBillRevokeAll (params) {
  return request({
    url: requestUrl('/web/trade/transBillRevokeAll'),
    method: 'post',
    data: requestParam(params)
  })
}

// 获取K线图数据
// export function getKline (params) {
//   return request({
//     url: requestUrl('/web/trade/getKline'),
//     method: 'post',
//     data: requestParam(params)
//   })
// }
export function getKline (params) {
  return request({
    url: requestUrl('/web/trade/getKline'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取实时最新成交价
export function getNewTradePrice (params) {
  return request({
    url: requestUrl('/web/trade/getNewTradePrice'),
    method: 'post',
    data: requestParam(params)
  })
}
// 获取前5条委托买/卖单数据
export function getTradeBill (params) {
  return request({
    url: requestUrl('/web/trade/getTradeBill'),
    method: 'post',
    data: requestParam(params)
  })
}
// 获取实时成交记录数据（最新20条）
export function getTradeBillRecord (params) {
  return request({
    url: requestUrl('/web/trade/getTradeBillRecord'),
    method: 'post',
    data: requestParam(params)
  })
}
