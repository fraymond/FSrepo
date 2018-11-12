export default {
  // 判断登录密码安全等级
  passwordSecurityLevel (password) {
    let securityLevelFlag = 0
    if (password.length < 6) {
      return 0
    } else {
      // lowercase
      if (/[a-z]/.test(password)) {
        securityLevelFlag++
      }
      // uppercase
      if (/[A-Z]/.test(password)) {
        securityLevelFlag++
      }
      // digital
      if (/[0-9]/.test(password)) {
        securityLevelFlag++
      }
      return securityLevelFlag
    }
  },
  setCookie (name, val, time) {
    let str = name + '=' + escape(val)
    if (time > 0) {
      let date = new Date()
      let ms = time * 3600 * 1000
      date.setTime(date.getTime() + ms)
      str += '; expires=' + date.toGMTString()
    }
    document.cookie = str
  },
  getCookie (name) {
    let arrStr = document.cookie.split('; ')
    for (let i = 0; i < arrStr.length; i++) {
      let temp = arrStr[i].split('=')
      if (temp[0] === name) {
        return unescape(temp[1])
      }
    }
  },
  delCookie (name) {
    let date = new Date()
    date.setTime(date.getTime() - 10000)
    document.cookie = name + '=a; expires=' + date.toGMTString()
  }
}
