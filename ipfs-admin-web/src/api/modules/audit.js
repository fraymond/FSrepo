import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'

// 获取锁仓记录列表
export function list (params) {
  return request({
    url: requestUrl('/transactionRecords/sc/listAll'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 解锁
export function unLock (params) {
  return request({
    url: requestUrl('/transactionRecords/sc/unLock'),
    method: 'post',
    data: requestParam(params)
  })
}
