// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'       // api: https://router.vuejs.org/
import store from './store'         // api: https://vuex.vuejs.org/
import VueCookie from 'vue-cookie'  // api: https://github.com/alfhen/vue-cookie
// import '@/element-ui'               // api: https://github.com/ElemeFE/element
import '@/iconfont'                 // api: http://www.iconfont.cn/
import '@/assets/scss/index.scss'
import { isAuth } from '@/utils'
import ElementUI from 'element-ui'
import layer from 'vue-layer'
import 'element-ui/lib/theme-chalk/index.css'

Vue.config.productionTip = false
Vue.prototype.$layer = layer(Vue, {
  msgtime: 3
})
Vue.use(VueCookie)
Vue.use(ElementUI)
// Vue.use(Clipboard)
Vue.prototype.isAuth = isAuth // 挂载权限方法

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: { App }
})
