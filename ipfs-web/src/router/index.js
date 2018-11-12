import Vue from 'vue'
import Router from 'vue-router'

// 开发环境不使用懒加载, 因为懒加载页面太多的话会造成webpack热更新太慢, 所以只有开发环境使用懒加载
const _import = require('./import-' + process.env.NODE_ENV)

Vue.use(Router)

// 路由定义使用说明:
// 1. 代码中路由统一使用name属性跳转.
// 2. 开放path属性用做简短路由, 比如: '/a1', 访问地址: www.renren.io/#/a1
export default new Router({
  mode: 'hash',
  routes: [
    { path: '/404', component: _import('error/404'), name: '404', desc: '404未找到' },
    {
      path: '/',
      component: _import('content/index'),
      name: 'content',
      redirect: { name: 'home' },
      desc: '上下整体布局',
      children: [
        { path: '/login', component: _import('login/index'), name: 'login', desc: '登录' },
        { path: '/register', component: _import('register/index'), name: 'register', desc: '注册' },
        { path: '/user', component: _import('user/index'), name: 'user', desc: '个人中心' },
        { path: '/home', component: _import('home/index'), name: 'home', desc: '主页' },
        { path: '/upload', component: _import('upload/index'), name: 'upload', desc: '文件上链' },
        { path: '/file-list', component: _import('file-list/index'), name: 'file-list', desc: '文件列表' },
        { path: '/operate-logs', component: _import('operate-logs/index'), name: 'operate-logs', desc: '操作日志' }
      ],
      beforeEnter (to, from, next) {
        let token = Vue.cookie.get('token')
        if (!token || !/\S/.test(token)) {
          //  next({ name: 'login' })
        }
        next()
      }
    },
    { path: '*', redirect: { name: '404' } }
  ]
})
