<template>
  <el-dialog
    :title="!dataForm.id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px"> 
      <!-- 子链合约地址 -->
      <el-form-item label="子链合约地址" prop="subchainAddress">
        <el-input v-model="dataForm.subchainAddress" placeholder="请输入子链合约地址"></el-input>
      </el-form-item>
      <!-- 存储大小 -->
      <el-form-item label="存储大小" prop="subchainSize">
        <el-input v-model="dataForm.subchainSize" placeholder="存储大小"></el-input>
        <!-- 500G 1T 2T 4T 8T 16T 32T -->
      </el-form-item>
      <!-- 已使用存储大小
      <el-form-item label="已使用存储大小" prop="useSize">
        <el-input v-model="dataForm.useSize" placeholder="已使用存储大小"></el-input>
      </el-form-item>
      已使用百分比
      <el-form-item label="已使用百分比" prop="percentageUse">
        <el-input v-model="dataForm.percentageUse" placeholder="已使用百分比"></el-input>
      </el-form-item>
      剩余存储量
      <el-form-item label="剩余存储量" prop="remainSize">
        <el-input v-model="dataForm.remainSize" placeholder="剩余存储量"></el-input>
      </el-form-item> -->
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
          subchainAddress: '',
          subchainSize: '',
          useSize: 0,
          percentageUse: 0,
          remainSize: 0
        },
        dataRule: {
          userName: [
            { required: true, message: '用户名不能为空', trigger: 'blur' }
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
            API.user.info1(this.dataForm.id).then(({data}) => {
              if (data && data.code === 0) {
                this.dataForm.subchainAddress = data.user.subchainAddress
                this.dataForm.subchainSize = data.user.subchainSize
                this.dataForm.useSize = data.user.useSize
                this.dataForm.percentageUse = data.user.percentageUse
                this.dataForm.remainSize = data.user.remainSize
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
              'userId': this.dataForm.id || undefined,
              'subchainAddress': this.dataForm.subchainAddress,
              'subchainSize': this.dataForm.subchainSize,
              'useSize': this.dataForm.useSize,
              'percentageUse': this.dataForm.percentageUse,
              'remainSize': this.dataForm.remainSize
            }
            var tick = !this.dataForm.id ? API.user.add1(params) : API.user.update1(params)
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
