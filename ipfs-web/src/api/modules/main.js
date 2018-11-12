import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'

//  获取用户文件列表接口
export function fileList (params) {
  return request({
    url: requestUrl('/web/queryFileList'),
    method: 'post',
    data: requestParam(params)
  })
}

//  获取操作日志列表接口
export function logList (params) {
  return request({
    url: requestUrl('/web/queryFileLogList'),
    method: 'post',
    data: requestParam(params)
  })
}

// 查看文件接口 用户地址+hash
export function queryFile (params) {
  return request({
    url: requestUrl('/web/queryFile'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除文件接口 用户地址+hash
export function deleteFile (params) {
  return request({
    url: requestUrl('/web/deleteFile'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}
