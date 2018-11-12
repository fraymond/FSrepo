<template>
  <el-dialog
    :title="!dataForm.id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px">
      <el-form-item label="标题" prop="cmcTitle">
        <el-input v-model="dataForm.cmcTitle" placeholder="标题"></el-input>
      </el-form-item>
      <el-form-item label="内容" prop="cmcContent">
        <el-input v-model="dataForm.cmcContent" placeholder="内容"></el-input>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" @click="dataFormSubmit()">确定</el-button>
    </span>
  </el-dialog>
</template>

<script>
  import API from '@/api'
  export default {
    data () {
      return {
        visible: false,
        roleList: [],
        dataForm: {
          id: 0,
          cmcTitle: '',
          cmcContent: ''
        },
        dataRule: {
          cmcTitle: [
            { required: true, message: '标题不能为空', trigger: 'blur' }
          ],
          cmcContent: [
            { required: true, message: '内容不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      init (id) {
        this.dataForm.id = id || 0
        API.role.select().then(({data}) => {
          this.roleList = data && data.code === 0 ? data.list : []
        }).then(() => {
          this.visible = true
          this.$nextTick(() => {
            this.$refs['dataForm'].resetFields()
          })
        }).then(() => {
          if (this.dataForm.id) {
            API.other.noticeInfo(this.dataForm.id).then(({data}) => {
              if (data && data.code === 0) {
                this.dataForm.cmcTitle = data.notice.cmcTitle
                this.dataForm.cmcContent = data.notice.cmcContent
              }
            })
          }
        })
      },
      // 表单提交
      dataFormSubmit () {
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {
            var params = {
              'cmcAutoid': this.dataForm.id || undefined,
              'cmcTitle': this.dataForm.cmcTitle,
              'cmcContent': this.dataForm.cmcContent
            }
            var tick = !this.dataForm.id ? API.other.noticeAdd(params) : API.other.noticeUpdate(params)
            tick.then(({data}) => {
              if (data && data.code === 0) {
                this.$message({
                  message: '操作成功',
                  type: 'success',
                  duration: 1500,
                  onClose: () => {
                    this.visible = false
                    this.$emit('refreshDataList')
                  }
                })
              } else {
                this.$message.error(data.msg)
              }
            })
          }
        })
      }
    }
  }
</script>
