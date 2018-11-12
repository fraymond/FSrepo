<template>
  <transition name="fade">
    <router-view></router-view>
  </transition>
</template>

<script>
  export default {
    beforeCreate () {
      //  页面初次加载 清除页面缓存 v-if="isRouterAlive"
      // localStorage.clear()
    },
    beforeDestory () {
      //  页面关闭时 清除页面缓存
      // localStorage.clear()
    }
  }
</script>
