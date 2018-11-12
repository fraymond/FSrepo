import request from '../request'
import requestUrl from '../requestUrl'
import requestParam from '../requestParam'

// 获取某用户地址列表
export function addAddresslist (params) {
  return request({
    url: requestUrl('web/setting/queryMailAddress'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取用户信息
export function getYhgMail (params) {
  return request({
    url: requestUrl('web/setting/queryYhglMail'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 获取可交易的币种列表
export function getTradeCurrency (params) {
  return request({
    url: requestUrl('web/trade/getTradeCurrency'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 新增地址
export function addAddress (params) {
  return request({
    url: requestUrl('web/setting/addMailAddress'),
    method: 'post',
    data: requestParam(params)
  })
}

// 删除地址
export function removeAddress (params) {
  return request({
    url: requestUrl('web/setting/deleteMailAddress'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 获取某用户资产列表
export function propertyList (params) {
  return request({
    url: requestUrl('web/setting/queryYhzc'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 修改交易密码
export function updateTradePassword (params) {
  return request({
    url: requestUrl('web/setting/tradePassword'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 修改登陆密码by 邮箱
export function updateEmailChangePwd (params) {
  return request({
    url: requestUrl('web/emailChangePwd'),
    method: 'post',
    data: requestParam(params, 'post', false)
  })
}

// 获取手机验证码 POST /web/setting/bindMobile
export function getMobileCode (params) {
  return request({
    url: requestUrl('web/setting/getMobileCode'),
    method: 'post',
    data: requestParam(params)
  })
}

// 绑定手机号
export function bindMobile (params) {
  return request({
    url: requestUrl('web/setting/bindMobile'),
    method: 'post',
    data: requestParam(params)
  })
}

// 绑定谷歌验证码
export function bindVerify (params) {
  return request({
    url: requestUrl('web/bindVerify'),
    method: 'post',
    data: requestParam(params)
  })
}

// 解绑谷歌验证码
export function unbindVerify (params) {
  return request({
    url: requestUrl('web/unbindVerify'),
    method: 'post',
    data: requestParam(params)
  })
}

// 生成包含用户信息,密钥的二维码图片
export function googleAuthQrCode (token) {
  return requestUrl(`web/googleAuthQrCode?token=${token}`)
}

// 获取secretKey
export function googleAuthStr () {
  return request({
    url: requestUrl('web/googleAuthStr'),
    method: 'get',
    params: requestParam({}, 'get')
  })
}

// 充币地址串
export function chargeCoin (params) {
  return request({
    url: requestUrl('web/setting/chargeCoin'),
    method: 'get',
    params: requestParam(params, 'get')
  })
}

// 充币二维码
export function chargeCoinQrCord (cmcCurrencyId, token) {
  return requestUrl(`web/setting/chargeCoinQrCord?cmcCurrencyId=${cmcCurrencyId}&token=${token}`)
}

// 用户提币
export function mentionCoin (params) {
  return request({
    url: requestUrl('web/setting/mentionCoin'),
    method: 'post',
    data: requestParam(params)
  })
}

// WEB用户充币检查
export function checkChargeCoin (params) {
  return request({
    url: requestUrl('web/setting/checkChargeCoin'),
    method: 'post',
    data: requestParam(params)
  })
}

// WEB用户提币检查
export function checkMentionCoin (params) {
  return request({
    url: requestUrl('web/setting/checkMentionCoin'),
    method: 'post',
    data: requestParam(params)
  })
}

// 设置是否开启谷歌验证
export function googleAuth (params) {
  return request({
    url: requestUrl('web/setting/googleAuth'),
    method: 'post',
    data: requestParam(params)
  })
}

// WEB忘记资金密码
export function forgetPassword (params) {
  return request({
    url: requestUrl('web/setting/forgetPassword'),
    method: 'post',
    data: requestParam(params)
  })
}

// WEB修改资金密码状态(是否交易时输入资金密码)
export function tradePasswordType (params) {
  return request({
    url: requestUrl('web/setting/tradePasswordType'),
    method: 'post',
    data: requestParam(params)
  })
}
