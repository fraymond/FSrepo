import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'

// 注册
export function register (params) {
  return request({
    url: requestUrl('/web/createAddress'),
    method: 'post',
    data: requestParam(params)
  })
}
