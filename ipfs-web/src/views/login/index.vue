<template>
  <div class="site-wrapper site-page--login">
    <div class="site-content__wrapper">
      <div class="site-content">
        <!-- 选择导入方式 -->
        <div class="login-main">
          <h3 class="login-title">请选择导入方式</h3>
          <el-radio-group v-model="importType">
            <el-radio-button :label="0">明文私钥</el-radio-button>
            <el-radio-button :label="1">keystore</el-radio-button>
            <el-radio-button :label="2">助记词</el-radio-button>
          </el-radio-group>
        </div>
        <div class="login-main" v-show="importType==0" v-loading="dataListLoading">
          <!-- 明文私钥导入 -->
          <el-form 
            :model="dataForm" 
            :rules="dataRule" 
            ref="dataForm0" 
            @keyup.enter.native="dataFormSubmit(0)" 
            status-icon
            >
            <!-- 明文私钥 -->
            <el-form-item prop="plaintextPrivateKey">
              <el-input v-model="dataForm.plaintextPrivateKey" type="textarea" placeholder="请粘贴明文私钥"></el-input>
            </el-form-item>
            <!-- 密码 -->
            <el-form-item prop="password">
              <el-input v-model="dataForm.password" type="password" placeholder="密码" clearable></el-input>
            </el-form-item>
            <!-- 重复密码 -->
            <el-form-item prop="confirmPassword">
              <el-input v-model="dataForm.confirmPassword" type="password" placeholder="确认密码" clearable></el-input>
            </el-form-item>
            <!-- 密码提示信息 -->
            <el-form-item prop="info">
              <el-input v-model="dataForm.info" type="text" placeholder="密码提示信息(选填)" clearable></el-input>
            </el-form-item>
            <!-- 导入和创建地址按钮 -->
            <el-form-item>
              <el-button class="login-btn-submit" type="primary" @click="dataFormSubmit(0)">导入地址</el-button>
              <el-button type="text" @click="goRegister()">创建地址</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="login-main" v-show="importType==1">
          <!-- keystore导入 -->
          <el-form :model="dataForm" :rules="dataRule" ref="dataForm1" @keyup.enter.native="dataFormSubmit(1)" status-icon>
            <!-- KEYSTORE -->
            <el-form-item prop="KEYSTORE_TYPE">
              <el-input v-model="dataForm.KEYSTORE_TYPE" type="textarea" placeholder="请粘贴KEYSTORE"></el-input>
            </el-form-item>
            <!-- 密码 -->
            <el-form-item prop="password">
              <el-input v-model="dataForm.password" type="password" placeholder="密码" clearable></el-input>
            </el-form-item>
            <!-- 重复密码 -->
            <el-form-item prop="confirmPassword">
              <el-input v-model="dataForm.confirmPassword" type="password" placeholder="确认密码" clearable></el-input>
            </el-form-item>
            <!-- 密码提示信息 -->
            <el-form-item prop="info">
              <el-input v-model="dataForm.info" type="text" placeholder="密码提示信息(选填)" clearable></el-input>
            </el-form-item>
            <!-- 导入和创建地址按钮 -->
            <el-form-item>
              <el-button class="login-btn-submit" type="primary" @click="dataFormSubmit(1)">导入地址</el-button>
              <el-button type="text" @click="goRegister()">创建地址</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="login-main" v-show="importType==2">
          <!-- 助记词 -->
          <el-form :model="dataForm" :rules="dataRule" ref="dataForm2" @keyup.enter.native="dataFormSubmit(2)" status-icon>
            <!-- 助记词 -->
            <el-form-item prop="MNEMONIC_TYPE">
              <el-input v-model="dataForm.MNEMONIC_TYPE" type="textarea" placeholder="请粘贴助记词"></el-input>
            </el-form-item>
            <!-- 密码 -->
            <el-form-item prop="password">
              <el-input v-model="dataForm.password" type="password" placeholder="密码" clearable></el-input>
            </el-form-item>
            <!-- 重复密码 -->
            <el-form-item prop="confirmPassword">
              <el-input v-model="dataForm.confirmPassword" type="password" placeholder="确认密码" clearable></el-input>
            </el-form-item>
            <!-- 密码提示信息 -->
            <el-form-item prop="info">
              <el-input v-model="dataForm.info" type="text" placeholder="密码提示信息(选填)" clearable></el-input>
            </el-form-item>
            <!-- 导入和创建地址按钮 -->
            <el-form-item>
              <el-button class="login-btn-submit" type="primary" @click="dataFormSubmit(2)">导入地址</el-button>
              <el-button type="text" @click="goRegister()">创建地址</el-button>
            </el-form-item>
          </el-form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import API from '@/api'
  import QbHeader from '@/views/header'
  export default {
    data () {
      return {
        importType: 0,              // 导入方式
        dataListLoading: false,
        dataForm: {
          password: '',             //  密码
          passwordHint: '',         //  密码提示信息
          plaintextPrivateKey: '',  //  明文私钥
          MNEMONIC_TYPE: '',        //  助记词
          KEYSTORE_TYPE: '',        //  keystore
          encryption: ''            //  加密方式
        },
        dataRule: {
          password: [               // 密码
            { required: true, message: '密码不能为空', trigger: 'blur' }
          ],
          confirmPassword: [        // 确认密码
            { required: true, message: '确认密码不能为空', trigger: 'blur' }
          ],
          plaintextPrivateKey: [    // 明文私钥
            { required: true, message: '明文私钥不能为空', trigger: 'blur' }
          ],
          MNEMONIC_TYPE: [          // 助记词
            { required: true, message: '助记词不能为空', trigger: 'blur' }
          ],
          KEYSTORE_TYPE: [          //  keystore
            { required: true, message: 'keystore不能为空', trigger: 'blur' }
          ],
          encryption: [             // 加密方式
            { required: true, message: '加密方式不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    components: {
      QbHeader
    },
    created () {
      //  登录提示语
      this.init()
    },
    methods: {
      init () {
        sessionStorage.clear('user-name')
      },
      // 去注册页面
      goRegister () {
        this.$router.push({ path: '/register' })
      },
      // 去找回密码页面
      goForget () {
        sessionStorage.removeItem('user-name')
        this.$router.push({ path: '/forget' })
      },
      // 提交表单
      dataFormSubmit (number) {
        this.dataListLoading = true
        if (number === 0) {
          this.$refs['dataForm0'].validate((valid) => {
            if (valid) {
              var params = {
                'importType': 'plaintextPrivateKey', // 导入方式
                'encryption': '',                    // 加密方式
                'plaintextPrivateKey': this.dataForm.plaintextPrivateKey, // 明文私钥
                'password': this.dataForm.password, // 密码
                'passwordHint': this.dataForm.info, // 密码提示语
                'keyStore': '',                     // keyStore
                'mnemonic': ''                      // 助记词
              }
              API.common.login(params).then(({data}) => {
                if (data && data.code === 0) {
                  this.dataListLoading = false
                  this.$message({
                    showClose: true,
                    message: '恭喜你，导入成功！',
                    type: 'success'
                  })
                  this.setUserInfo(data)
                } else {
                  this.$notify({
                    message: data.msg,
                    position: 'bottom-left',
                    type: 'error',
                    duration: 1500
                  })
                }
              })
            }
          })
        } else if (number === 1) {
          this.$refs['dataForm1'].validate((valid) => {
            if (valid) {
              var params = {
                'importType': 'KEYSTORE_TYPE',            // 导入方式
                'encryption': '',                         // 加密方式
                'keyStore': this.dataForm.KEYSTORE_TYPE,  // keyStore
                'password': this.dataForm.password,       // 密码
                'passwordHint': this.dataForm.info,       // 密码提示语
                'mnemonic': '',                           // 助记词
                'plaintextPrivateKey': ''                 // 明文私钥
              }
              API.common.login(params).then(({data}) => {
                if (data && data.code === 0) {
                  this.$message({
                    showClose: true,
                    message: '恭喜你，导入成功！',
                    type: 'success'
                  })
                  this.setUserInfo(data)
                } else {
                  this.$notify({
                    message: data.msg,
                    position: 'bottom-left',
                    type: 'error',
                    duration: 1500
                  })
                }
              })
            }
          })
        } else if (number === 2) {
          this.$message({
            showClose: true,
            message: '该方式暂未开放，敬请期待',
            type: 'warning'
          })
          this.importType = 0
          return false
          // this.$refs['dataForm2'].validate((valid) => {
          //   if (valid) {
          //     var params = {
          //       'importType': 'MNEMONIC_TYPE',       // 导入方式
          //       'encryption': '',                    // 加密方式
          //       'mnemonic': this.dataForm.mnemonic,  // mnemonic
          //       'password': this.dataForm.password,  // 密码
          //       'passwordHint': this.dataForm.info,  // 密码提示语
          //       'plaintextPrivateKey': '',           // 明文私钥
          //       'keyStore': ''
          //     }
          //     API.common.login(params).then(({data}) => {
          //       if (data && data.code === 0) {
          //         this.$notify({
          //           message: '登录成功',
          //           position: 'bottom-left',
          //           type: 'success',
          //           duration: 1500
          //         })
          //         this.setUserInfo(data)
          //       } else {
          //         this.$notify({
          //           message: data.msg,
          //           position: 'bottom-left',
          //           type: 'error',
          //           duration: 1500
          //         })
          //       }
          //     })
          //   }
          // })
        }
      },
      // 登录成功后缓存用户信息
      setUserInfo (data) {
        // 保存用户信息
        // 获取路由是从哪个页面进入登录界面，登录成功后回到对应的页面
        localStorage.removeItem('userInfo')
        let userInfo = data.data
        localStorage.setItem('userInfo', userInfo)
        window.location.reload()
        let rname = this.$route.query.rname
        this.$router.replace({ name: 'home' })
        if (rname !== undefined && rname !== 'undefined') {
          this.$router.push({ name: rname, query: {cmcName: this.dataForm.email} })
        } else {
          this.$router.push({ path: '/', query: {cmcName: this.dataForm.email} })
        }
      }
    }
  }
</script>

<style lang="scss">
  .txt-align-c {
    text-align: center;
  }
  .login-warning .el-button {
    padding-left: 60px;
    padding-right: 60px;
    font-size: 16px;
  }
  .login-warning-promit {
    font-size: 15px;
    color: #333;
    line-height: 22px;
    margin-bottom: 20px;
  }
  .login-warning-promit .col-red {
    color: red;
  }
  .el-notification__closeBtn {
    top: 10px;
    right: 5px;
  }
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
      padding: 5% 10%;
      width: 470px;
      min-height: 100%;
      margin-top: 10%;
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
    .login-btn-submit {
      width: 100%;
      margin-top: 38px;
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
  }
</style>
