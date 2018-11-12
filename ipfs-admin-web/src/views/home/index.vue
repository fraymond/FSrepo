<template>
  <!-- 首页 --> 
  <div class="site-content site-content--home">
    <el-card :style="contentViewStyles">
      <h3>交易所介绍</h3>
      <p>
        QBIT是一家包含多币种数字资产货币交易的平台，面向全球用户提供比特币，以太坊，莱特币等多种主流数字资产交易服务，
        是安全可信赖的数字资产交易平台！为用户提供安全，快捷，公正的数字资产交易平台！ 
      </p>
      <p>
        币种多样 平台上线多种主流数字资产货币，
        一键在线交易 极致体验 充值即时、提现迅速，快速响应 服务专业 专业的客服团队，7x24小时在线服务 系统可靠 不定期升级系统，多重交易验证，保障交易安全
      </p>
    </el-card>
  </div>
</template>

<script>
  export default {
    data () {
      return {
      }
    },
    computed: {
      contentViewStyles () {
        var height = this.$store.state.documentClientHeight
        height -= 50 // site-topbar
        height -= 15 // site-content padding-top
        height -= 15 // site-content padding-bottom
        height += 'px'
        return [
          { minHeight: height }
        ]
      }
    }
  }
</script>

<style>
  .site-content--home {
    line-height: 1.5;
  }
</style>

