import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'
import isInteger from 'lodash/isInteger'

// 获取对标代币列表
export function list (params) {
  return request({
    url: requestUrl('/transaction/currencyRelation/currencyRelationList'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取未/已分配币种列表信息
export function info (cmcAutoid) {
  return request({
    url: requestUrl('/transaction/currencyRelation/distributionCurrencyList' + (isInteger(cmcAutoid) ? `/${cmcAutoid}` : '')),
    method: 'get',
    params: requestParam(cmcAutoid, 'get')
  })
}

// 改变未/已分配币种列表信息
export function changeInfo (params) {
  return request({
    url: requestUrl('/transaction/currencyRelation/addOrDeleteDistribution'),
    method: 'post',
    data: requestParam(params)
  })
}
