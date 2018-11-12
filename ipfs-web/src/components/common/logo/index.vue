<template>
  <div @click="goHome()">
    <img :src="img" alt="logo">
  </div>
</template>

<script>
export default {
  data () {
    return {
      img: 'https://qiub-wallet.oss-cn-shenzhen.aliyuncs.com/img/chainImg/w-logo.png'
    }
  },
  created () {
  },
  methods: {
    init (id) {
      this.img = this.formatting(id)
    },
    // 回主页
    goHome () {
      this.$router.push({ path: '/home' })
    },
    formatting (id) {
      id = 'http://' + id
      return id
    }
  }
}
</script>
<style>
 .main .logo img {
   width: 100px;
   height: 60px;
   max-width: 100%;
 }
</style>
