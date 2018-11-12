// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'       // api: https://router.vuejs.org/
import VueCookie from 'vue-cookie'  // api: https://github.com/alfhen/vue-cookie
import '@/iconfont'                 // api: http://www.iconfont.cn/
import '@/assets/scss/index.scss'
import { isAuth } from '@/utils'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.config.productionTip = false
Vue.use(VueCookie)
Vue.use(ElementUI)
// vue-router切换页面时滚动条自动滚动到顶部
router.afterEach((to, from, next) => {
  window.scrollTo(0, 0)
})
Vue.prototype.isAuth = isAuth // 挂载权限方法
// 设置全局token实效变量
Vue.prototype.notToken = false
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
