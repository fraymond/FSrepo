import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'

// 获取用户列表
export function list (params) {
  return request({
    url: requestUrl('/transactionRecords/czjl/listAll'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取用户列表
export function currList (params) {
  return request({
    url: requestUrl('/transaction/currency/currList'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}
