import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'

// 获取验证码
export function captcha (uuid) {
  return requestUrl(`captcha.jpg?uuid=${uuid}`)
}

// 明文私钥导入地址 参数区别类型
export function login (params) {
  return request({
    url: requestUrl('/web/importAddress'),
    method: 'post',
    data: requestParam(params)
  })
}

// 登录时Google Authenticator的校验
export function loginVerify (params) {
  return request({
    url: requestUrl('web/loginVerify'),
    method: 'post',
    data: requestParam(params)
  })
}

// 找回密码
export function forget (params) {
  return request({
    url: requestUrl('web/forget'),
    method: 'post',
    data: requestParam(params)
  })
}

// 退出
export function logout () {
  return request({
    url: requestUrl('web/logout'),
    method: 'post',
    data: requestParam()
  })
}

// 接口msg：token失效，请重新登录
export function tokenMsg () {
  return 'token失效，请重新登录'
}
