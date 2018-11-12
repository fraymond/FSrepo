<template>
  <ul class="main">
    <li class="logo"><qb-logo ref="logo"></qb-logo></li>
    <li class="menu"><qb-menu></qb-menu></li>
    <li class="user"><qb-user res="user"></qb-user></li>
  </ul>
</template>
<script>
  import QbLogo from '@/components/common/logo'
  import QbUser from '@/components/common/user'
  import QbMenu from '@/components/common/menu'
  export default {
    data () {
      return {
        imgSrc: 'qiubi.oss-cn-shenzhen.aliyuncs.com/upload/20180413/1ef92fe230c7471394b502f9fee088d9.png',
        userName: ''
      }
    },
    created () {
      // 初始化
      this.init()
    },
    methods: {
      // 初始化
      init () {
        // 给组件logo传值
        this.$nextTick(() => {
          this.$refs.logo.init(this.imgSrc)
        })
      }
    },
    components: {
      QbLogo,
      QbMenu,
      QbUser
    }
  }
</script>
<style>
  html,body{
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
    background-color: #f1f2f6;
  }
  .main {
    width: 100%;
    min-width: 1330px;
    height: 80px;
    line-height: 80px;
    background-color: #ffffff;
    margin: 0 0 20px 0;
    padding: 0;
    -webkit-box-shadow:0 0 10px rgba(144, 145, 147, .8);
    -moz-box-shadow:0 0 10px rgba(144, 145, 147, .8);
         box-shadow:0 0 10px rgba(144, 145, 147, .8);
  }
  .main li {
    float: left;
    list-style-type: none;
  }
  .main .logo {
    width: 32%;
    text-align: right;
  }
  .main .menu {
    width: 35%;
    padding-left: 3%;
  }
  .main .user {
    width: 30%;
    padding-left: 5%;
  }
</style>
