<template>
  <el-dialog
    :title="!dataForm.id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px">
      <el-form-item label="内容" prop="cmcNote">
        <el-input v-model="dataForm.cmcNote" placeholder="内容"></el-input>
      </el-form-item>
      <el-form-item label="路径" prop="cmcUrl">
        <el-input v-model="dataForm.cmcUrl" placeholder="路径"></el-input>
      </el-form-item>
      <upload ref="upload" @uploadFile="returnUploadFile"></upload>      
      <!-- <el-form-item label="密码" prop="password" :class="{ 'is-required': !dataForm.id }">
        <el-input v-model="dataForm.password" type="password" placeholder="密码"></el-input>
      </el-form-item>
      <el-form-item label="确认密码" prop="comfirmPassword" :class="{ 'is-required': !dataForm.id }">
        <el-input v-model="dataForm.comfirmPassword" type="password" placeholder="确认密码"></el-input>
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="dataForm.email" placeholder="邮箱"></el-input>
      </el-form-item>
      <el-form-item label="手机号" prop="mobile">
        <el-input v-model="dataForm.mobile" placeholder="手机号"></el-input>
      </el-form-item>
      <el-form-item label="角色" size="mini" prop="roleIdList">
        <el-checkbox-group v-model="dataForm.roleIdList">
          <el-checkbox v-for="role in roleList" :key="role.roleId" :label="role.roleId">{{ role.roleName }}</el-checkbox>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="状态" size="mini" prop="status">
        <el-radio-group v-model="dataForm.status">
          <el-radio :label="0">禁用</el-radio>
          <el-radio :label="1">正常</el-radio>
        </el-radio-group>
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
  import Upload from '@/components/picture-upload'
  export default {
    data () {
      return {
        visible: false,
        roleList: [],
        dataForm: {
          id: 0,
          cmcNote: '',
          cmcUrl: ''
        },
        dataRule: {
          cmcNote: [
            { required: true, message: '内容不能为空', trigger: 'blur' }
          ],
          cmcUrl: [
            { required: true, message: '路径不能为空', trigger: 'blur' }
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
            API.other.bannerInfo(this.dataForm.id).then(({data}) => {
              if (data && data.code === 0) {
                this.dataForm.cmcNote = data.banner.cmcNote
                this.dataForm.cmcUrl = data.banner.cmcUrl
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
              'cmcNote': this.dataForm.cmcNote,
              'cmcUrl': this.dataForm.cmcUrl
            }
            var tick = !this.dataForm.id ? API.other.bannerAdd(params) : API.other.bannerUpdate(params)
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
      },
      //  获取文件上传返回回来的数据
      returnUploadFile (data) {
        if (data && data.code === 0) {
          this.dataForm.cmcUrl = data.url
        }
      }
    },
    components: {
      Upload
    }
  }
</script>
