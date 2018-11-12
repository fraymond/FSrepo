<template>
  <el-dialog
    :title="!dataForm.id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px" v-loading.fullscreen.lock="sLoading" element-loading-text="提交中...">
      <el-form-item label="身份" prop="cmcVip" v-show="dataForm.id">
        <el-select v-model="dataForm.cmcVip" placeholder="请选择会员级别">
          <el-option v-for="(item, index) in chooseCmcVip" :key="index" :label="item.value" :value="item.id"></el-option>
        </el-select>
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
        sLoading: false,
        visible: false,
        roleList: [],
        chooseCmcVip: [{id: 0, value: '普通用户'}, {id: 1, value: 'vip1'}, {id: 2, value: 'vip2'}, {id: 3, value: 'vip3'}],
        dataForm: {
          id: 0,
          cmcVip: ''
        },
        dataRule: {
          cmcVip: [
            { required: true, message: '请选择会员级别', trigger: 'change' }
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
            API.userManagement.info(this.dataForm.id).then(({data}) => {
              if (data && data.code === 0) {
                this.dataForm.cmcVip = data.yhgl.cmcVip
              }
            })
          }
        })
      },
      // 表单提交
      dataFormSubmit () {
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {
            this.sLoading = true
            var params = {
              'cmcAutoid': this.dataForm.id || undefined,
              'cmcVip': this.dataForm.cmcVip
            }
            API.userManagement.update(params).then(({data}) => {
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
              this.sLoading = false
            })
          }
        })
      }
    }
  }
</script>
