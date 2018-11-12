import * as common from './modules/common'
import * as user from './modules/user'
import * as role from './modules/role'
import * as menu from './modules/menu'
import * as main from './modules/main'
import * as config from './modules/config'
import * as oss from './modules/oss'
import * as schedule from './modules/schedule'
import * as register from './modules/register'
import * as home from './modules/home'
import * as setting from './modules/setting'
import * as trade from './modules/trade'
import * as upload from './modules/upload'

export default {
  common,     // 公共
  user,       // 管理员管理
  role,       // 角色管理
  menu,       // 菜单管理
  main,       // 系统日志
  config,     // 参数管理
  oss,        // 文件服务
  schedule,   // 定时任务
  register,   // 用户注册
  home,       // 首页
  setting,    // 设置
  trade,      // 交易
  upload      // 文件上传
}
