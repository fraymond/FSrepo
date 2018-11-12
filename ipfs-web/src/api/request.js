import Vue from 'vue'
import axios from 'axios'
import router from '@/router'
// 创建axios实例
const service = axios.create({
  timeout: 1000 * 30,
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json; charset=utf-8'
  }
})
// request拦截器
service.interceptors.request.use(config => {
  config.headers['token'] = Vue.cookie.get('token')
  return config
}, error => {
  return Promise.reject(error)
})
// response拦截器
service.interceptors.response.use(response => {
  // let ifGetCurrRelation = (response.request.responseURL).includes('home/getCurrRelation')
  // let ifGetAllJYCoin = (response.request.responseURL).includes('home/getAllJYCoin')
  // let href = window.location.href.includes('business')
  if (response.data && response.data.code === 401) { // 401, token失效
    Vue.cookie.delete('token')
    router.push({ name: 'login' })
  }
  return response
}, error => {
  //  判断用户是否在其他地点登录，登录后调回登录页面
  let href = window.location.href.includes('business')
  if (href === true && (error.response === undefined || error.response === 'undefined')) {
    Vue.cookie.delete('token')
    sessionStorage.clear()
    router.push({ path: '/login', query: {rname: 'home'} })
    return false
  }
  return Promise.reject(error)
})

export default service
