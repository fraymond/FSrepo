<template>
  <el-dialog :title="!dataForm.id ? '新增' : '修改'" :close-on-click-modal="false" :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px">
      <el-form-item label="图标" prop="cmcWhUrl">
        <el-input v-model="dataForm.cmcWhUrl" placeholder="图标URL"></el-input>
      </el-form-item>
      <upload ref="upload" @uploadFile="returnUploadFile"></upload>
      <el-form-item label="币种编号" prop="cmcNumber">
        <el-input v-model="dataForm.cmcNumber" placeholder="币种编号"></el-input>
      </el-form-item>
      <el-form-item label="中文名" prop="zhName">
        <el-input v-model="dataForm.zhName" placeholder="中文名"></el-input>
      </el-form-item>
      <el-form-item label="英文名" prop="enName">
        <el-input v-model="dataForm.enName" placeholder="英文名"></el-input>
      </el-form-item>
      <el-form-item label="缩写" prop="sxName">
        <el-input v-model="dataForm.sxName" placeholder="缩写"></el-input>
      </el-form-item>
      <el-form-item label="币种介绍" prop="intro">
        <el-input v-model="dataForm.intro" placeholder="币种介绍"></el-input>
      </el-form-item>
      <el-form-item label="推出时间" prop="addTime">
        <el-input v-model="dataForm.addTime" placeholder="推出时间"></el-input>
      </el-form-item>
      <el-form-item label="发行总量" prop="fxCount">
        <el-input v-model="dataForm.fxCount" placeholder="发行总量"></el-input>
      </el-form-item>
      <el-form-item label="存量" prop="clCount">
        <el-input v-model="dataForm.clCount" placeholder="存量"></el-input>
      </el-form-item>
      <el-form-item label="市值" prop="sValue">
        <el-input v-model="dataForm.sValue" placeholder="市值"></el-input>
      </el-form-item>
      <el-form-item label="币种算法" prop="coinSf">
        <el-input v-model="dataForm.coinSf" placeholder="币种算法"></el-input>
      </el-form-item>
      <el-form-item label="区块速度" prop="speed">
        <el-input v-model="dataForm.speed" placeholder="区块速度"></el-input>
      </el-form-item>
      <el-form-item label="区块大小" prop="bigMin">
        <el-input v-model="dataForm.bigMin" placeholder="区块大小"></el-input>
      </el-form-item>
      <el-form-item label="调整难度" prop="cmcPriority">
        <el-input v-model="dataForm.cmcPriority" placeholder="调整难度"></el-input>
      </el-form-item>
      <el-form-item label="证明方式" prop="certify">
        <el-input v-model="dataForm.certify" placeholder="证明方式"></el-input>
      </el-form-item>
      <el-form-item label="链接" prop="connects">
        <el-input v-model="dataForm.connects" placeholder="链接"></el-input>
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
  import Upload from '@/components/picture-upload'
  export default {
    data () {
      return {
        visible: false,
        roleList: [],
        dataForm: {
          id: 0,
          cmcWhUrl: '',
          cmcNumber: '',
          zhName: '',
          enName: '',
          sxName: '',
          intro: '',
          addTime: '',
          fxCount: '',
          clCount: '',
          sValue: '',
          coinSf: '',
          speed: '',
          bigMin: '',
          cmcPriority: '',
          certify: '',
          connects: '',
          difficult: '',
          roleIdList: []
        },
        dataRule: {
          cmcWhUrl: [
            { required: true, message: '图标URL不能为空', trigger: 'blur' }
          ],
          cmcNumber: [
            { required: true, message: '币种编号不能为空', trigger: 'blur' }
          ],
          zhName: [
            { required: true, message: '中文名不能为空', trigger: 'blur' }
          ],
          enName: [
            { required: true, message: '英文名不能为空', trigger: 'blur' }
          ],
          sxName: [
            { required: true, message: '缩写不能为空', trigger: 'blur' }
          ],
          intro: [
            { required: true, message: '币种介绍不能为空', trigger: 'blur' }
          ],
          fxCount: [
            { required: true, message: '发行总量不能为空', trigger: 'blur' }
          ],
          clCount: [
            { required: true, message: '存量不能为空', trigger: 'blur' }
          ],
          sValue: [
            { required: true, message: '市值不能为空', trigger: 'blur' }
          ],
          coinSf: [
            { required: true, message: '币种算法不能为空', trigger: 'blur' }
          ],
          speed: [
            { required: true, message: '区块速度不能为空', trigger: 'blur' }
          ],
          addTime: [
            { required: true, message: '推出时间不能为空', trigger: 'blur' }
          ],
          bigMin: [
            { required: true, message: '区块大小不能为空', trigger: 'blur' }
          ],
          cmcPriority: [
            // 数字类型
            { required: true, message: '调整难度不能为空', trigger: 'blur' }
          ],
          certify: [
            { required: true, message: '证明方式不能为空', trigger: 'blur' }
          ],
          connects: [
            { required: true, message: '链接不能为空', trigger: 'blur' }
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
            API.introduction.info(this.dataForm.id).then(({data}) => {
              if (data && data.code === 0) {
                this.dataForm.cmcWhUrl = data.bzwh.cmcWhUrl
                this.dataForm.cmcNumber = data.bzwh.cmcNumber
                this.dataForm.zhName = data.bzwh.zhName
                this.dataForm.enName = data.bzwh.enName
                this.dataForm.sxName = data.bzwh.sxName
                this.dataForm.intro = data.bzwh.intro
                this.dataForm.addTime = data.bzwh.addTime
                this.dataForm.fxCount = data.bzwh.fxCount
                this.dataForm.clCount = data.bzwh.clCount
                this.dataForm.sValue = data.bzwh.sValue
                this.dataForm.coinSf = data.bzwh.coinSf
                this.dataForm.speed = data.bzwh.speed
                this.dataForm.bigMin = data.bzwh.bigMin
                this.dataForm.difficult = data.bzwh.difficult
                this.dataForm.certify = data.bzwh.certify
                this.dataForm.connects = data.bzwh.connects
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
              'cmcWhUrl': this.dataForm.cmcWhUrl,
              'cmcNumber': this.dataForm.cmcNumber,
              'zhName': this.dataForm.zhName,
              'enName': this.dataForm.enName,
              'sxName': this.dataForm.sxName,
              'intro': this.dataForm.intro,
              'fxCount': this.dataForm.fxCount,
              'clCount': this.dataForm.clCount,
              'sValue': this.dataForm.sValue,
              'coinSf': this.dataForm.coinSf,
              'speed': this.dataForm.speed,
              'bigMin': this.dataForm.bigMin,
              'difficult': this.dataForm.difficult,
              'cmcPriority': this.dataForm.cmcPriority,
              'certify': this.dataForm.certify,
              'connects': this.dataForm.connects,
              'roleIdList': this.dataForm.roleIdList
            }
            var tick = !this.dataForm.id ? API.introduction.add(params) : API.introduction.update(params)
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
          this.dataForm.cmcWhUrl = data.url
        }
      }
    },
    components: {
      Upload
    }
  }
</script>
