import * as common from './modules/common'
import * as user from './modules/user'
import * as role from './modules/role'
import * as menu from './modules/menu'
import * as log from './modules/log'
import * as config from './modules/config'
import * as oss from './modules/oss'
import * as schedule from './modules/schedule'
import * as app from './modules/appVersion'
import * as introduction from './modules/introduction'
import * as dataMaintain from './modules/data_maintain'
import * as userManagement from './modules/user_management'
import * as benchmarkingAllocation from './modules/benchmarking_allocation'
import * as revenue from './modules/revenue'
import * as audit from './modules/audit'
import * as transaction from './modules/transaction'  // 交易模块
import * as other from './modules/other'              // 其他模块

export default {
  common,                  // 公共
  user,                    // 管理员管理
  role,                    // 角色管理
  menu,                    // 菜单管理
  log,                     // 系统日志
  config,                  // 参数管理
  oss,                     // 文件服务
  schedule,                // 定时任务
  app,                     // app版本控制
  introduction,            // 币种介绍维护
  dataMaintain,            // 币种数据维护
  userManagement,          // 用户管理
  benchmarkingAllocation,  // 对标币种分配
  transaction,             // 交易管理
  other,                   // 其它管理
  audit,                   // 审核管理
  revenue                  // 收益记录管理
}
