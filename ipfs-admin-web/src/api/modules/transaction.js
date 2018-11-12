import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'
import requestParamArr from '../requestParamArr'

// 获取委托交易记录列表
export function wtList (params) {
  return request({
    url: requestUrl('/transactionRecords/transbill/listAll'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 撤销委托交易订单
export function wtRemove (params) {
  return request({
    url: requestUrl('/transactionRecords/transbill/revoke'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 获取提币&充币记录列表
export function list (params) {
  return request({
    url: requestUrl('/transactionRecords/czjl/listAll'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取提币&充币记录列表
export function currList (params) {
  return request({
    url: requestUrl('/transaction/currency/currList'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取锁仓记录列表
export function scList (params) {
  return request({
    url: requestUrl('/transactionRecords/sc/listAll'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 解锁锁仓记录
export function scRemove (params) {
  return request({
    url: requestUrl('/transactionRecords/sc/unLock'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 获取成交记录列表
export function cjList (params) {
  return request({
    url: requestUrl('/transactionRecords/transbillrecord/listAll'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取提币审核列表
export function tbAuditList (params) {
  return request({
    url: requestUrl('/transactionRecords/czjl/tbshList'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 提币审核 通过
export function checkPass (params) {
  return request({
    url: requestUrl('/transactionRecords/czjl/check'),
    method: 'post',
    data: requestParamArr(params, 'post', false)
  })
}

// 提币审核 驳回
export function checkNoPass (params) {
  return request({
    url: requestUrl('/transactionRecords/czjl/checkNoPass'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 导出充币/提币列表Excel
export function exportExcel (params) {
  return request({
    // headers: {
    //   'content-type': 'application/vnd.ms-excel;charset=UTF-8'
    // },
    url: requestUrl('transactionRecords/czjl/exportExcel'),
    method: 'get',
    params: requestParamArr(params, 'get')
  })
}
