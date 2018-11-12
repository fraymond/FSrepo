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
    { path: '/login', component: _import('login/index'), name: 'login', desc: '登录' },
    {
      path: '/',
      component: _import('layout/index'),
      name: 'layout',
      // 默认加载home
      redirect: { name: 'home' },
      desc: '上左右整体布局',
      children: [
        { path: '/home', component: _import('home/index'), name: 'home', desc: '首页' },
        {
          path: '/content-tabs',
          component: _import('layout/content-tabs'),
          name: 'content-tabs',
          redirect: { name: '404' },
          desc: '内容需通过tabs展示',
          children: [
            // 以'/n'开头统一拦截, 标记为左侧菜单导航tab展示内容方式路由
            { path: '/n/user', component: _import('user/index'), name: 'user', desc: '管理员管理' },
            { path: '/n/role', component: _import('role/index'), name: 'role', desc: '角色管理' },
            { path: '/n/menu', component: _import('menu/index'), name: 'menu', desc: '菜单管理' },
            { path: '/n/sql', component: _import('sql/index'), name: 'sql', desc: 'SQL监控' },
            { path: '/n/schedule', component: _import('schedule/index'), name: 'schedule', desc: '定时任务' },
            { path: '/n/config', component: _import('config/index'), name: 'config', desc: '参数管理' },
            { path: '/n/oss', component: _import('oss/index'), name: 'oss', desc: '文件上传' },
            { path: '/n/log', component: _import('log/index'), name: 'log', desc: '系统日志' },
            // 存储子链管理 modules/sys/user.html
            { path: '/n/zzgl', component: _import('zzgl/index'), name: 'zzgl', desc: '子链管理' },
            { path: '/n/zllog', component: _import('zllog/index'), name: 'zllog', desc: '子链日志' },
            // 基础数据管理
            { path: '/n/base_data/app_edition', component: _import('base_data/app_edition/index'), name: 'app_edition', desc: 'app版本控制' },
            // 交易基础数据维护
            { path: '/n/currency/data_maintain', component: _import('currency/data_maintain/index'), name: 'data_maintain', desc: '币种数据维护' },
            { path: '/n/currency/user_management', component: _import('currency/user_management/index'), name: 'user_management', desc: '用户管理' },
            { path: '/n/currency/benchmarking_allocation', component: _import('currency/benchmarking_allocation/index'), name: 'benchmarking_allocation', desc: '对标币种分配' },
            { path: '/n/currency/introduction', component: _import('currency/introduction/index'), name: 'introduction', desc: '币种介绍维护' },
            { path: '/n/currency/recharge_record_maintenance', component: _import('currency/recharge_record_maintenance/index'), name: 'recharge_record_maintenance', desc: '充值记录维护' },
            // 交易记录
            { path: '/n/transaction/recharge_withdraw_log', component: _import('transaction/recharge_withdraw_log/index'), name: 'recharge_withdraw_log', desc: '充币/提币记录' },
            { path: '/n/transaction/entrust_order', component: _import('transaction/entrust_order/index'), name: 'entrust_order', desc: '委托订单记录' },
            { path: '/n/transaction/record_of_transaction', component: _import('transaction/record_of_transaction/index'), name: 'record_of_transaction', desc: '成交记录' },
            { path: '/n/transaction/locked_record', component: _import('transaction/locked_record/index'), name: 'locked_record', desc: '锁仓记录' },
            // 审核中心
            { path: '/n/audit/locked_board', component: _import('audit/locked_board/index'), name: 'locked_board', desc: '锁仓看板' },
            { path: '/n/audit/withdraw_audit', component: _import('audit/withdraw_audit/index'), name: 'withdraw_audit', desc: '提币审核' },
             // 数据分析
             { path: '/n/data/revenue_record_statistics', component: _import('data/revenue_record_statistics/index'), name: 'revenue_record_statistics', desc: '收益记录统计' },
             { path: '/n/data/financial_statement', component: _import('data/financial_statement/index'), name: 'financial_statement', desc: '财产报表' },
            // 其他
            { path: '/n/other/banner', component: _import('other/banner/index'), name: 'banner', desc: '轮播图' },
            { path: '/n/other/notice', component: _import('other/notice/index'), name: 'notice', desc: '公告' },
            { path: '/n/other/message', component: _import('other/message/index'), name: 'message', desc: '留言板' },
            { path: '/n/other/help', component: _import('other/help/index'), name: 'help', desc: '帮助中心' }
          ]
        }
      ],
      beforeEnter (to, from, next) {
        let token = Vue.cookie.get('token')
        if (!token || !/\S/.test(token)) {
          next({ name: 'login' })
        }
        next()
      }
    },
    { path: '*', redirect: { name: '404' } }
  ]
})
