import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'

// 上传文件接口
export function upload (params) {
  return request({
    url: requestUrl('/web/addFile'),
    method: 'post',
    data: requestParam(params)
  })
}
