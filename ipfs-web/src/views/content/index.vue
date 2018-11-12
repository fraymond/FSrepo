<template>
    <div class="qibui-body">
      <ul class="main adaper-header">
        <li class="logo"><qb-logo ref="ssss"></qb-logo></li>
        <li class="menu"><qb-menu></qb-menu></li>
        <li class="user"><qb-user :ifLogin="ifLogin"></qb-user></li>
      </ul>
      <router-view></router-view>
    </div>
</template>

<script>
  import '@/assets/css/pcAdaper.css'
  import QbLogo from '@/components/common/logo'
  import QbUser from '@/components/common/user'
  import QbMenu from '@/components/common/menu'

  export default {
    data () {
      return {
        imgSrc: 'qiub-wallet.oss-cn-shenzhen.aliyuncs.com/img/chainImg/w-logo.png',
        isShow: true,
        ifLogin: '',
        isShowNoty: true
      }
    },
    created () {
      // 初始化
      this.init()
    },
    methods: {
      // 初始化
      init () {
        // 给组件ssss传值
        this.$nextTick(() => {
          this.$refs.ssss.init(this.imgSrc)
        })
      },
      goForget () {
        this.$router.push({ path: '/forget' })
      }
    },
    components: {
      QbLogo,
      QbMenu,
      QbUser
    },
    watch: {
      '$route': function () {
        let cmcName = this.$route.query.cmcName
        if (cmcName !== undefined && cmcName !== 'undefined') {
          this.ifLogin = this.$route.query.cmcName
        } else {
          this.ifLogin = ''
        }
      }
    }
  }
</script>
<style scoped src = "../../assets/css/pcAdaper.css"></style>
<style>
  html,body{
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
  }
  .main {
    width: 100%;
    min-width: 1330px;
    height: 80px;
    line-height: 80px;
    background-color: #2f4056;
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
    width: 25%;
    text-align: right;
  }
  .main .menu {
    width: 45%;
    height: 100%;
    padding-left: 2%;
  }
  .main .user {
    width: 30%;
  }
</style>
