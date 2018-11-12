<template>
    <el-upload
        drag
        :action="url"
        :before-upload="beforeUploadHandle"
        :on-success="successHandle"
        :multiple="isMultiple"
        :file-list="fileList"
        class="common-upload">
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <div class="el-upload__tip" slot="tip">只支持jpg、png、gif格式的图片！</div>
    </el-upload>
</template>

<script>
  import API from '@/api'
  export default {
    data () {
      return {
        isMultiple: false, //  是否支持多选文件
        url: '',
        num: 0,
        successNum: 0,
        fileList: []
      }
    },
    mounted () {
      this.$nextTick(() => {
        this.init()
      })
    },
    methods: {
      init (id) {
        this.url = API.oss.upload(this.$cookie.get('token'))
      },
      // 上传之前
      beforeUploadHandle (file) {
        if (file.type !== 'image/jpg' && file.type !== 'image/jpeg' && file.type !== 'image/png' && file.type !== 'image/gif') {
          this.$message.error('只支持jpg、png、gif格式的图片！')
          return false
        }
        this.num++
      },
      // 上传成功
      successHandle (response, file, fileList) {
        this.fileList = fileList
        this.successNum++
        if (response && response.code === 0) {
          if (this.num === this.successNum) {
            this.$message.success('上传成功')
            this.$emit('uploadFile', file.response)
            this.fileList = []
            // this.$confirm('操作成功, 是否继续操作?', '提示', {
            //   confirmButtonText: '确定',
            //   cancelButtonText: '取消',
            //   type: 'warning'
            // }).catch(() => {})
          }
        } else {
          this.$message.error(response.msg)
        }
      }
    }
  }
</script>
<style>
.common-upload {
  text-align: center;
}
</style>
