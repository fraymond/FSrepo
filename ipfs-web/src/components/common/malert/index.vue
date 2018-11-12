<template>
  <el-dialog :title="confirmTitle" :close-on-click-modal="false" :visible.sync="visible" class="alert-warning">
      <a href="javascript:;" @click="goLogin()">tocken实效，请前往登录 </a>
      <div class="txt-align-c">
        <el-button type="warning" @click="confirmOk()">确定</el-button>
      </div>
  </el-dialog>
</template>
<script>
export default {
  data () {
    return {
      confirmTitle: '提示',
      visible: true
    }
  },
  methods: {
    goLogin () {
      this.$router.push({ path: '/login' })
    },
    confirmOk () {
      this.visible = false
    }
  }
}
</script>
<style>
 .alert-warning {
   text-align: center
 }
</style>
