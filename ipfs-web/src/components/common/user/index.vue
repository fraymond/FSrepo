<template>
  <div>
    <div class="user-item" v-if="isShow">
      <el-button type="text" @click="goLogin()">导入地址</el-button>
    </div>
    <div class="user-item" v-else>
      <!-- 模拟数据 -->
      <el-button type="text" @click="goUser()">当前地址：{{userName}} </el-button>
      <el-button type="text" @click="goRegister()">点我切换地址</el-button>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      isShow: true,
      userInfo: [],
      userName: '请登录'
    }
  },
  created () {
    this.init()
  },
  methods: {
    //  判断用户登陆
    init () {
      this.userInfo = JSON.parse(localStorage.getItem('userInfo'))
      if (this.userInfo !== null) {
        this.isShow = false
        this.userName = this.userInfo.address
      } else {
        this.isShow = true
      }
    },
    bufen (value) {
      return value.slice(0, 10)
    },
    // 去个人中心
    goUser () {
      this.$router.push({ path: '/' })
    },
    // 去登录页面
    goLogin () {
      this.$router.push({ path: '/login' })
    },
    // 去注册页面
    goRegister () {
      this.$router.push({ path: '/register' })
    }
  },
  watch: {
    'ifLogin': function () {
      this.userName = this.ifLogin
      this.checkUserName()
    }
  }
}
</script>
<style>
  .user .el-button--text {
    color: #909399;
  }
  .user .user-item {
    height:80px;
    line-height:80px;
  }
</style>
