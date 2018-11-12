import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'

// 获取收益记录列表
export function list (params) {
  return request({
    url: requestUrl('/dataAnalysis/gainrecord/listAll'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}
