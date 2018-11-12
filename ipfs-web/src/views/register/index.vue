<template>
  <div class="site-wrapper site-page--login">
    <div class="site-content__wrapper">
      <div class="site-content">
        <div class="login-main register-main">
          <h3 class="login-title">创建地址</h3>
          <el-form 
            :model="dataForm" 
            :rules="dataRule" 
            ref="dataForm" 
            @keyup.enter.native="dataFormSubmit()" 
            status-icon
            v-loading="dataListLoading"
            >
            <el-form-item prop="password">
              <el-input v-model="dataForm.password" type="password" placeholder="密码" clearable></el-input>
            </el-form-item>
            <el-form-item prop="confirmPassword">
              <el-input v-model="dataForm.confirmPassword" type="password" placeholder="确认密码" clearable></el-input>
            </el-form-item>
            <el-form-item prop="info">
              <el-input v-model="dataForm.info" type="text" placeholder="密码提示信息(选填)" clearable></el-input>
            </el-form-item>
            <el-form-item>
              <el-button :disabled="!isAgree" class="register-btn-submit" type="primary" @click="dataFormSubmit()">立即创建</el-button>
            </el-form-item>
          </el-form>
          <el-checkbox v-model="isAgree">我已阅读并同意</el-checkbox>
          <el-button type="text">IPFS服务条款</el-button>
          <!--  @click="openAgree()" -->
          <el-button type="text" @click="goLogin()" style="float:right;" class="hover-pointer">导入地址</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import QbHeader from '@/views/header'
  import API from '@/api'
  export default {
    data () {
      return {
        isAgree: false,
        dataListLoading: false,
        dataForm: {
          password: '',
          confirmPassword: '',
          info: ''
        },
        dataRule: {
          password: [
            { required: true, message: '密码不能为空', trigger: 'blur' },
            { min: 6, message: '密码长度大于6位', trigger: 'blur' },
            { pattern: /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]+$/, message: '密码必须由数字、英文组成', trigger: 'blur' }
          ],
          confirmPassword: [
            { required: true, message: '确认密码不能为空', trigger: 'blur' }
          ]
        },
        captchaPath: ''
      }
    },
    components: {
      QbHeader
    },
    created () {
    },
    methods: {
      // 去登录页面
      goLogin () {
        this.$router.push({ path: '/login' })
      },
      // 提交表单
      dataFormSubmit () {
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {
            var params = {
              'password': this.dataForm.password,
              'passwordHint': this.dataForm.info
            }
            this.dataListLoading = true
            API.register.register(params).then(({data}) => {
              if (data && data.code === 0) {
                this.dataListLoading = false
                this.$notify({
                  title: '成功',
                  message: '恭喜你，创建地址成功！',
                  type: 'success'
                })
                // 如果缓存中没有保存用户地址json文件
                if (localStorage.getItem('userJson') == null) {
                  // let userJson = []
                  // userJson.push(data.data.userId)
                  data.data = JSON.stringify(data.data)
                  localStorage.setItem('userInfo', data.data)
                  // localStorage.setItem('userJson', userJson)
                } else { // 如果已经存在则添加到json尾部
                  // let userJson = []
                  // userJson.push(data.data.userId)
                  // data.data = JSON.stringify(data.data)
                  localStorage.removeItem('userInfo')
                  data.data = JSON.stringify(data.data)
                  localStorage.setItem('userInfo', data.data)
                }
                window.location.reload()
                this.$router.push({ path: '/home' })
              } else {
                this.$message.error(data.msg)
              }
            })
          }
        })
      },
      //  打开用户协议
      openAgree () {
        this.$router.push({ path: '/agree' })
      }
    }
  }
</script>

<style lang="scss">
  .site-wrapper {
    background-color: transparent;
    min-height: 714px;
  }
  .site-wrapper.site-page--login {
    .site-content__wrapper {
      padding: 0;
      margin: 0;
      overflow-x: hidden;
      overflow-y: auto;
    }
    .site-content {
      width: 910px;
      margin: 0 auto;
      padding: 100px 0 0 405px;
      // background: url('~@/assets/img/bj.png') left center no-repeat;
      background-size: 40% auto;
    }
    .brand-info {
      width: 365px;
      margin-right: 40px;
    }
    .brand-info img {
      width: 100%;
    }
    .login-main {
      padding: 25px 50px;
      width: 470px;
      min-height: 100%;
      background-color: #fff;
    }
    .login-title {
      font-size: 16px;
      margin-bottom: 20px;
    }
    .login-captcha {
      overflow: hidden;
      > img {
        width: 100%;
        cursor: pointer;
      }
    }
    .register-btn-submit {
      width: 100%;
      margin-top: 10px;
    }
    .el-button--primary {
      background-color: #103990;
      border-color: #103990;
    }
    .el-checkbox__input.is-checked .el-checkbox__inner, .el-checkbox__input.is-indeterminate .el-checkbox__inner {
      background-color: #103990;
      border-color: #103990;
    }
    .el-checkbox__input.is-checked+.el-checkbox__label {
      color: #103990;
    }
    .el-button--text {
      color: #103990;
    }
    .el-button--primary.is-disabled {
      background-color: #ccc;
      border-color: #ccc;
    }
    .input-email {
      width:65%;
      margin-right:10px;
    }
  }
</style>
