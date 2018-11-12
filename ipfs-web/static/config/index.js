/**
 * 开发环境
 */
;(function () {
  window.SITE_CONFIG = {}
  // api接口请求地址 //47.74.224.104:8080/ //192.168.2.115:8080/
  window.SITE_CONFIG.baseUrl = '//192.168.1.23:8080/moacipfs'

  // 嵌套iframe地址
  window.SITE_CONFIG.nestIframeUrl = '//192.168.1.23:8080/moacipfs'
  // 嵌套iframe路由名称列表
  window.SITE_CONFIG.nestIframeRouteNameList = ['sql']
  // 静态资源文件夹名称
  window.SITE_CONFIG.staticFileName = ''
  // cdn地址
  window.SITE_CONFIG.cdnUrl = './' + window.SITE_CONFIG.staticFileName
})()
