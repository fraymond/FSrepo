<template>
  <!-- 上传的提示语待优化 -->
  <section
    class="fileBox"      
    title="上传文件"
    v-loading="dataListLoading"
    >        
      <el-input placeholder="请输入文件备注内容(选填)" v-model="info" clearable class="elinput"></el-input>
      <input type="file" id="myFile" class="inputfile" @change="handlerUpload($event)">
      <label for="myFile" v-show="uploadNumber==0">点击上传</label>
      <label for="myFile" v-show="uploadNumber==1">继续上传</label>
      <label for="myFile" v-show="uploadNumber==2">正在上传</label>
      <el-radio-group v-model="encrypt" class="elencrypt">
        <el-radio :label="0">不加密</el-radio>
        <el-radio :label="1">加密</el-radio>
      </el-radio-group>
      <span class="span" slot="tip">温馨提醒：IPFS1.0体验版仅支持50M以内的文件上链！</span>
  </section>
</template>

<script>
  import API from '@/api'
  import axios from 'axios'
  export default {
    data () {
      return {
        // 给go文件对象
        files: [],
        // 给java文件对象
        theFiles: [],
        // 文件加密
        encrypt: 0,
        // 上传次数，默认 0
        uploadNumber: 0,
        // 用户对象
        userInfo: [],
        dataListLoading: false,
        //  文件备注
        info: ''
      }
    },
    created () {
      this.init()
    },
    methods: {
      // 判断用户是否登陆
      init () {
        this.userInfo = localStorage.getItem('userInfo')
        if (this.userInfo !== null) {
        } else {
          this.$alert('上传文件需要导入地址', '温馨提示', {
            confirmButtonText: '立即导入'
          })
          this.$router.push({ path: '/login' })
        }
      },
      // 上传文件
      handlerUpload: function (e) {
        //  获取选定的文件
        let tFiles = e.target.files
        this.theFiles = tFiles
        let len = tFiles.length
        //  文件hash
        //  是否加密
        for (var i = 0; i < len; i++) {
          //  开始上传每一个文件
          var item = {
            name: tFiles[i].name,
            uploadPercentage: 1,
            size: this.formatFileSize(tFiles[i].size, 0)
          }
          this.files.push(item)
          //  开始上传文件 新建一个formData
          let param = new FormData()
          // 通过append向form对象添加数据
          param.append('uploadfile', tFiles[i])
          if (!this.checkFileSize(tFiles[i].size)) {
            // 文件超过50m
            return false
          }
          // 通过axios上传文件
          // 配置
          let config = {
            // 添加请求头
            headers: {
              'Content-Type': 'multipart/form-data'
            }
          }
          let _this = this
          axios.post('http://47.107.99.26:9090/upload', param, config).then(function (response) {
            _this.send(response.data.resultData.path)
          }).catch(function (error) {
            console.log(error)
          })
        }
      },
      send (filePath) {
        //  用户地址
        let _this = this
        let userInfo = JSON.parse(localStorage.getItem('userInfo'))
        let address = userInfo.address
        let obj = {}
        obj.address = address                               // 用户地址
        obj.fileRealSize = _this.theFiles[0].size           // 文件大小
        obj.fileName = _this.theFiles[0].name               // 文件名称
        obj.filePath = filePath                             // go文件路径
        obj.encrypt = _this.encrypt                         // 文件是否加密
        obj.remark = _this.info                             // 文件备注
        _this.dataListLoading = true
        API.upload.upload(obj).then(function (data) {
          if (data.data.code === 0) {
            _this.dataListLoading = false
            _this.$message({
              message: '恭喜你，上链成功！',
              type: 'success'
            })
          } else {
            _this.$message({
              message: data.data.msg,
              type: 'warning'
            })
          }
        }).catch(function (error) {
          console.log(error)
        })
      },
      formatFileSize: function (fileSize, idx) {
        var units = ['B', 'KB', 'MB', 'GB']
        idx = idx || 0
        if (fileSize < 1024 || idx === units.length - 1) {
          return fileSize.toFixed(1) + units[idx]
        }
        return this.formatFileSize(fileSize / 1024, ++idx)
      },
      checkFileSize: function (fileSize) {
        // 2M
        const MAX_SIZE = 2 * 1024 * 1024
        if (fileSize > MAX_SIZE) {
          return false
        }
        return true
      }
    }
  }
</script>

<style>
  .fileBox {
      margin: auto;
      text-align: center;
      padding-top: 10%;
      width: 400px;
      font-size: 16px;
  }
  .inputfile {
      opacity: 0;
      overflow: hidden;
      position: absolute;
      z-index: -1;
  }
  .el-input {
    width: 300px;
  }
  .el-radio-group {
    height: 40px;
    line-height: 40px;
    float: left;
    margin-top: 5%;
  }
  .span {
    float: left;
    margin-top: 5%;
    font-size: 12px;
  }
  /*E + F 毗邻元素选择器，匹配所有紧随E元素之后的同级元素F*/
  .inputfile+label {
      float: right;
      width: 80px;
      height: 40px;
      line-height: 40px;
      background-color: indianred;
      color: white;
      border-radius: 5%;
      margin-left: 5%;
  }
  .inputfile:focus+label,
  .inputfile+label:hover {
      color: white;
  }
</style>

