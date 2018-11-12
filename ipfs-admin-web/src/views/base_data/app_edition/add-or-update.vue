<template>
  <el-dialog :title="!dataForm.id ? '新增' : '修改'" :close-on-click-modal="false" :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px">
      <el-form-item label="版本号" prop="version">
        <el-input v-model="dataForm.version" placeholder="版本号"></el-input>
      </el-form-item>
      <el-form-item label="版本名称" prop="versionname">
        <el-input v-model="dataForm.versionname" placeholder="版本名称"></el-input>
      </el-form-item>
      <el-form-item label="更新日志" prop="updatelog">
        <el-input v-model="dataForm.updatelog" placeholder="更新日志"></el-input>
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
          version: '',
          versionname: '',
          updatelog: '',
          roleIdList: []
        },
        dataRule: {
          version: [
            { required: true, message: '版本号不能为空', trigger: 'blur' }
          ],
          versionname: [
            { required: true, message: '版本名称不能为空', trigger: 'blur' }
          ],
          updatelog: [
            { required: true, message: '更新日志不能为空', trigger: 'blur' }
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
            API.app.info(this.dataForm.id).then(({data}) => {
              if (data && data.code === 0) {
                this.dataForm.version = data.appVersion.version
                this.dataForm.versionname = data.appVersion.versionname
                this.dataForm.updatelog = data.appVersion.updatelog
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
              'id': this.dataForm.id || undefined,
              'version': this.dataForm.version,
              'versionname': this.dataForm.versionname,
              'updatelog': this.dataForm.updatelog
            }
            var tick = !this.dataForm.id ? API.app.add(params) : API.app.update(params)
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
