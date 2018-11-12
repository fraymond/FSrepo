<template>
  <el-dialog
    :title="!dataForm.id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="120px" v-loading.fullscreen.lock="dataListLoading" element-loading-text="提交中...">
      <el-form-item label="图标" prop="cmcTburl">
        <el-input v-model="dataForm.cmcTburl" placeholder="图标"></el-input>
      </el-form-item>
      <upload ref="upload" @uploadFile="returnUploadFile"></upload>
      <el-form-item label="系列" prop="cmcSeries">
        <el-input v-model="dataForm.cmcSeries" placeholder="系列"></el-input>
      </el-form-item>
      <!-- <el-form-item label="系列" prop="cmcSeries">
        <el-select v-model="dataForm.cmcSeries" placeholder="请选择类型">
          <el-option v-for="(item, index) in chooseCmcType" :key="index" :label="item" :value="item"></el-option>
        </el-select>
      </el-form-item> -->
      <el-form-item label="货币编号" prop="cmcNumber">
        <el-input v-model="dataForm.cmcNumber" placeholder="货币编号"></el-input>
      </el-form-item>
      <el-form-item label="是否为交易货币" prop="cmcJy" v-if="dataForm.id===0 ? false : true">
        <el-select v-model="dataForm.cmcJy" placeholder="是否为交易货币">
          <el-option v-for="(item, index) in chooseCmcJy" :key="index" :label="item.value" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="是否上架交易所" prop="cmcJyzt" v-if="dataForm.id===0 ? false : true">
        <el-select v-model="dataForm.cmcJyzt" placeholder="是否上架交易所">
          <el-option v-for="(item, index) in chooseCmcJyzt" :key="index" :label="item.value" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="能否充币" prop="cmcCb">
        <el-select v-model="dataForm.cmcCb" placeholder="能否充币">
          <el-option v-for="(item, index) in chooseCmcCb" :key="index" :label="item.value" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="能否提币" prop="cmcTb">
        <el-select v-model="dataForm.cmcTb" placeholder="能否提币">
          <el-option v-for="(item, index) in chooseCmcTb" :key="index" :label="item.value" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="货币名称" prop="cmcName">
        <el-input v-model="dataForm.cmcName" placeholder="货币名称"></el-input>
      </el-form-item>
      <el-form-item label="最低可提" prop="cmcDbzd">
        <el-input v-model="dataForm.cmcDbzd" placeholder="最低可提"></el-input>
      </el-form-item>
      <el-form-item label="交易手续费" prop="cmcSxf">
        <el-input v-model="dataForm.cmcSxf" placeholder="交易手续费"></el-input>
      </el-form-item>
      <el-form-item label="提币手续费" prop="cmcTbsxf">
        <el-input v-model="dataForm.cmcTbsxf" placeholder="提币手续费"></el-input>
      </el-form-item>
      <el-form-item label="优先级" prop="cmcPriority">
        <el-input v-model="dataForm.cmcPriority" placeholder="优先级"></el-input>
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
        dataListLoading: false,
        visible: false,
        roleList: [],
        chooseCmcType: [],
        chooseCmcJy: [{id: 0, value: '否'}, {id: 1, value: '是'}],
        chooseCmcJyzt: [{id: 0, value: '否'}, {id: 1, value: '是'}],
        chooseCmcCb: [{id: 0, value: '否'}, {id: 1, value: '能'}],
        chooseCmcTb: [{id: 0, value: '否'}, {id: 1, value: '能'}],
        dataForm: {
          cmcNumber: '',
          cmcSeries: '',
          cmcTburl: '',
          cmcName: '',
          kind: '',
          cmcJy: '',
          cmcJyzt: '',
          cmcCb: '',
          cmcTb: '',
          cmcDbzd: '',
          cmcSxf: '',
          cmcTbsxf: '',
          cmcPriority: '',
          roleIdList: [],
          status: 1
        },
        dataRule: {
          cmcTburl: [
            { required: true, message: '图标不能为空', trigger: 'blur' }
          ],
          cmcSeries: [
            { required: true, message: '系列不能为空', trigger: 'blur' }
          ],
          cmcNumber: [
            { required: true, message: '用户名不能为空', trigger: 'blur' }
          ],
          cmcJy: [
            { required: true, message: '是否为交易货币未设置', trigger: 'change' }
          ],
          cmcJyzt: [
            { required: true, message: '是否上架交易所未设置', trigger: 'change' }
          ],
          cmcCb: [
            { required: true, message: '能否充币未设置', trigger: 'change' }
          ],
          cmcTb: [
            { required: true, message: '能否提币未设置', trigger: 'change' }
          ],
          cmcName: [
            { required: true, message: '货币名称不能为空', trigger: 'blur' }
          ],
          cmcDbzd: [
            { required: true, message: '最低可提未设置', trigger: 'blur' }
          ],
          cmcSxf: [
            { required: true, message: '交易手续费未设置', trigger: 'blur' }
          ],
          cmcTbsxf: [
            { required: true, message: '提币手续费未设置', trigger: 'blur' }
          ],
          cmcPriority: [
            { required: true, message: '优先级未设置', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      init (id) {
        // this.getInfo()
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
            API.dataMaintain.info(this.dataForm.id).then(({data}) => {
              this.formatting(data.transaction)
              if (data && data.code === 0) {
                this.dataForm.cmcNumber = data.transaction.cmcNumber
                this.dataForm.cmcTburl = data.transaction.cmcTburl
                this.dataForm.cmcName = data.transaction.cmcName
                this.dataForm.cmcSeries = data.transaction.cmcSeries
                this.dataForm.cmcJy = data.transaction.cmcJy
                this.dataForm.cmcJyzt = data.transaction.cmcJyzt
                this.dataForm.cmcCb = data.transaction.cmcCb
                this.dataForm.cmcTb = data.transaction.cmcTb
                this.dataForm.cmcDbzd = data.transaction.cmcDbzd
                this.dataForm.cmcSxf = data.transaction.cmcSxf
                this.dataForm.cmcTbsxf = data.transaction.cmcTbsxf
                this.dataForm.cmcPriority = data.transaction.cmcPriority
              }
            })
          }
        })
      },
      // getInfo () {
      //   // 获取系列数据
      //   var params = {}
      //   API.rechargeWithdrawLog.currList(params).then(({data}) => {
      //     // 请求成功
      //     if (data && data.code === 0) {
      //       // 赋值
      //       let arr = []
      //       let array = data.currencyList
      //       for (let i = 0; i < array.length; i++) {
      //         arr.push(array[i].cmcSeries)
      //       }
      //       arr = [...new Set(arr)]
      //       this.chooseCmcType = arr
      //     }
      //   })
      // },
      tool (array) {
        let arr = []
        for (let i = 0; i < array.length; i++) {
          arr.push(array[i].cmcSeries)
        }
        arr = [...new Set(arr)]
        return arr
      },
      // 格式化
      formatting (array) {
        // 格式化类型
        if (array.cmcJy === 0) {
          array.cmcJy = '否'
        } else if (array.cmcJy === 1) {
          array.cmcJy = '是'
        }
        if (array.cmcJyzt === 0) {
          array.cmcJyzt = '否'
        } else if (array.cmcJyzt === 1) {
          array.cmcJyzt = '是'
        }
        if (array.cmcCb === 0) {
          array.cmcCb = '否'
        } else if (array.cmcCb === 1) {
          array.cmcCb = '能'
        }
        if (array.cmcTb === 0) {
          array.cmcTb = '否'
        } else if (array.cmcTb === 1) {
          array.cmcTb = '能'
        }
        return array
      },
      unformatting (array) {
        // 格式化类型
        if (array.cmcJy === '否') {
          array.cmcJy = 0
        } else if (array.cmcJy === '是') {
          array.cmcJy = 1
        }
        if (array.cmcJyzt === '否') {
          array.cmcJyzt = 0
        } else if (array.cmcJyzt === '是') {
          array.cmcJyzt = 1
        }
        if (array.cmcCb === '否') {
          array.cmcCb = 0
        } else if (array.cmcCb === '能') {
          array.cmcCb = 1
        }
        if (array.cmcTb === '否') {
          array.cmcTb = 0
        } else if (array.cmcTb === '能') {
          array.cmcTb = 1
        }
        return array
      },
      // 表单提交
      dataFormSubmit () {
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {
            this.unformatting(this.dataForm)
            if (this.dataForm.id) {
              var paramss = {
                'cmcAutoid': this.dataForm.id,
                'cmcNumber': this.dataForm.cmcNumber,
                'cmcTburl': this.dataForm.cmcTburl,
                'cmcName': this.dataForm.cmcName,
                'cmcSeries': this.dataForm.cmcSeries,
                'cmcJy': this.dataForm.cmcJy,
                'cmcJyzt': this.dataForm.cmcJyzt,
                'cmcCb': this.dataForm.cmcCb,
                'cmcTb': this.dataForm.cmcTb,
                'cmcDbzd': this.dataForm.cmcDbzd,
                'cmcSxf': this.dataForm.cmcSxf,
                'cmcTbsxf': this.dataForm.cmcTbsxf,
                'cmcPriority': this.dataForm.cmcPriority
              }
            } else {
              var params = {
                'cmcNumber': this.dataForm.cmcNumber,
                'cmcTburl': this.dataForm.cmcTburl,
                'cmcName': this.dataForm.cmcName,
                'cmcSeries': this.dataForm.cmcSeries,
                'cmcCb': this.dataForm.cmcCb,
                'cmcTb': this.dataForm.cmcTb,
                'cmcDbzd': this.dataForm.cmcDbzd,
                'cmcSxf': this.dataForm.cmcSxf,
                'cmcTbsxf': this.dataForm.cmcTbsxf,
                'cmcPriority': this.dataForm.cmcPriority
              }
            }
            this.dataListLoading = true
            var tick = !this.dataForm.id ? API.dataMaintain.add(params) : API.dataMaintain.update(paramss)
            tick.then(({data}) => {
              if (data && data.code === 0) {
                this.dataListLoading = false
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
                this.dataListLoading = false
                this.$message.error(data.msg)
              }
            })
          }
        })
      },
      //  获取文件上传返回回来的数据
      returnUploadFile (data) {
        if (data && data.code === 0) {
          this.dataForm.cmcTburl = data.url
        }
      }
    },
    components: {
      Upload
    }
  }
</script>