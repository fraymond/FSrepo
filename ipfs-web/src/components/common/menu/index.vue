<template>
<el-menu :default-active="$route.name" class="el-menu-demo" mode="horizontal">
  <el-menu-item index="home" @click="goTo(1)">主页</el-menu-item>
  <el-menu-item index="upload" @click="goTo(2)">文件上链</el-menu-item>
  <el-menu-item index="file-list" @click="goTo(3)">文件列表</el-menu-item>
  <el-menu-item index="operate-logs" @click="goTo(4)">操作日志</el-menu-item>
</el-menu>
</template>

<script>
  export default {
    data () {
      return {
        activeIndex: '1'
      }
    },
    methods: {
      goTo (e) {
        // 清除交易缓存
        sessionStorage.removeItem('BusinessTabIndex')
        sessionStorage.removeItem('changeCmc')
        //  判断用户是否登录
        // var user = sessionStorage.getItem('userInfo')
        if (e === 1) {
          this.$router.push({ path: '/home' })
        } else if (e === 2) {
          this.$router.push({ path: '/upload' })
          // if (user !== null && user !== 'null') {
          //   this.$router.push({ path: '/upload' })
          // } else {
          //   this.$message({
          //     message: '请先登录或注册',
          //     type: 'warning',
          //     duration: 1500,
          //     center: true
          //   })
          //   this.$router.push({ path: '/login', query: {rname: 'upload'} })
          // }
        } else if (e === 3) {
          this.$router.push({ path: '/file-list' })
          // if (user !== null && user !== 'null') {
          //   this.$router.push({ path: '/file-list' })
          // } else {
          //   this.$message({
          //     message: '请先登录或注册',
          //     type: 'warning',
          //     duration: 1500,
          //     center: true
          //   })
          //   this.$router.push({ path: '/login', query: {rname: 'file-list'} })
          // }
        } else if (e === 4) {
          this.$router.push({ path: '/operate-logs' })
          // if (user !== null && user !== 'null') {
          //   this.$router.push({ path: '/operate-logs' })
          // } else {
          //   this.$message({
          //     message: '请先登录或注册',
          //     type: 'warning',
          //     duration: 1500,
          //     center: true
          //   })
          //   this.$router.push({ path: '/login', query: {rname: 'operate-logs'} })
          // }
        }
      }
    }
  }
</script>

<style>
.el-menu-item {
  font-size:16px;
  padding: 0 35px;
}
.el-menu--horizontal>.el-menu-item {
  height: 78px;
  line-height: 78px;
  border: 0;
}
.el-menu--horizontal {
  border: 0px;
}
.el-menu--horizontal>.el-menu-item.is-active {
  border: 0;
  color: wheat;
}
.el-dropdown-menu__item--divided:before, 
.el-menu, 
.el-menu--horizontal>.el-menu-item:not(.is-disabled):focus, 
.el-menu--horizontal>.el-menu-item:not(.is-disabled):hover, 
.el-menu--horizontal>.el-submenu .el-submenu__title:hover {
  background-color: #2f4056;
  color: wheat;
  border: 0;
}
.hover-pointer {
  cursor: pointer;
}
</style>
