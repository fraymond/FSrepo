<template>
  <!-- 系统管理 / 管理员列表 -->
  <div class="mod-user">
    <!--  搜索表单对象  -->
    <el-form :inline="true" :model="dataForm" @keyup.enter.native="getDataList()">
      <!-- 搜索栏目输入框 -->
      <el-form-item>
        <!-- 搜索 关键词 -->
        <el-input v-model="chooseCmcEmail" placeholder="用户邮箱" clearable></el-input>
      </el-form-item>
      <el-form-item>
        <el-select v-model="kind" placeholder="请选择类型" clearable>
          <el-option v-for="(item, index) in chooseCmcType" :key="index" :label="item.value" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-select v-model="region" placeholder="请选择交易代币" clearable>
          <el-option v-for="(item, index) in currencyList" :key="index" :label="item.cmcNumber" :value="item.cmcAutoid"></el-option>
        </el-select>
      </el-form-item>
      <!-- 搜索 开始时间 -->
      <el-form-item>
        <el-date-picker
          v-model="startDjsj"
          type="datetime"
          placeholder="选择开始日期时间"
          align="right"
          :picker-options="pickerOptions1" clearable>
        </el-date-picker>
      </el-form-item>
      <!-- 搜索结束时间 -->
      <el-form-item>
        <el-date-picker
          v-model="endDjsj"
          type="datetime"
          placeholder="选择结束日期时间"
          align="right"
          :picker-options="pickerOptions2" clearable>
        </el-date-picker>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="getDataList()" @keyup.enter.native="getDataList()">查询</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="danger" @click="exportExcel()">导出</el-button>
      </el-form-item>
    </el-form>
    <!-- Loading 加载 v-loading="xxx"-->
    <el-table :data="dataList" border v-loading="dataListLoading" tooltip-effect="dark" class="rw-table">
      <!-- 类型 -->
      <el-table-column prop="cmcType" header-align="center" align="center" label="类型" sortable></el-table-column>
      <!-- 交易单号 -->
      <el-table-column prop="cmcOrderNumber" header-align="center" align="center" label="交易单号" show-overflow-tooltip>
        <template slot-scope="scope">
          <p class="tag-read" :data-clipboard-text="scope.row.cmcOrderNumber" @dblclick="copy(),open9()" plain>{{scope.row.cmcOrderNumber}}</p>
        </template>
      </el-table-column>
      <!-- 手续费 -->
      <el-table-column prop="cmcSxf" header-align="center" align="center" label="手续费"></el-table-column>
      <!-- 登记时间 -->
      <el-table-column prop="cmcDjsj" header-align="center" align="center" label="登记时间"></el-table-column>
      <!-- 用户 -->
      <el-table-column prop="cmcEmail" header-align="center" align="center" label="用户" sortable show-overflow-tooltip>
        <template slot-scope="scope">
          <p class="tag-read" :data-clipboard-text="scope.row.cmcEmail" @dblclick="copy(),open9()">{{scope.row.cmcEmail}}</p>
        </template>
      </el-table-column>
      <!-- 币种编号 -->
      <el-table-column prop="cmcNumber" header-align="center" align="center" label="币种编号" sortable></el-table-column>
      <!-- 数量 -->
      <el-table-column prop="cmcSl" header-align="center" align="center" label="数量" sortable></el-table-column>
      <!-- 实际到账 -->
      <el-table-column prop="cmcSjdz" header-align="center" align="center" label="实际到账"></el-table-column>
      <!-- 接受地址 -->
      <el-table-column prop="cmcJsdz" header-align="center" align="center" label="接受地址" show-overflow-tooltip>
        <template slot-scope="scope">
          <p class="tag-read" :data-clipboard-text="scope.row.cmcJsdz" @dblclick="copy(),open9()">{{scope.row.cmcJsdz}}</p>
        </template>
      </el-table-column>
      <!-- 交易哈希 -->
      <el-table-column prop="cmcHashAddress" header-align="center" align="center" label="交易hash" show-overflow-tooltip>
        <template slot-scope="scope">
          <p class="tag-read" :data-clipboard-text="scope.row.cmcHashAddress" @dblclick="copy(),open9()">{{scope.row.cmcHashAddress}}</p>
        </template>
      </el-table-column>
      <el-table-column prop="cmcZt" header-align="center" align="center" label="状态">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 0" size="small" type="danger">禁用</el-tag>
          <el-tag v-else size="small">正常</el-tag>
        </template>
      </el-table-column>
      <!-- 备注 -->
      <el-table-column prop="cmcBz" header-align="center" align="center" label="备注" show-overflow-tooltip></el-table-column>
    </el-table>
    <!--                              每页数                             当前页-->
    <el-pagination @size-change="sizeChangeHandle" @current-change="currentChangeHandle" :current-page="pageIndex" :page-sizes="[10, 20, 50, 100]" :page-size="pageSize" :total="totalPage" layout="total, sizes, prev, pager, next, jumper"></el-pagination>
  </div>
</template>

<script>
  import API from '@/api'
  import Vue from 'vue'
  import Clipboard from 'clipboard'
  export default {
    data () {
      return {
        // form表单对象
        pickerOptions1: {
          shortcuts: [{
            text: '今天',
            onClick (picker) {
              picker.$emit('pick', new Date())
            }
          }, {
            text: '昨天',
            onClick (picker) {
              const date = new Date()
              date.setTime(date.getTime() - 3600 * 1000 * 24)
              picker.$emit('pick', date)
            }
          }, {
            text: '一周前',
            onClick (picker) {
              const date = new Date()
              date.setTime(date.getTime() - 3600 * 1000 * 24 * 7)
              picker.$emit('pick', date)
            }
          }]
        },
        pickerOptions2: {
          shortcuts: [{
            text: '今天',
            onClick (picker) {
              picker.$emit('pick', new Date())
            }
          }, {
            text: '昨天',
            onClick (picker) {
              const date = new Date()
              date.setTime(date.getTime() - 3600 * 1000 * 24)
              picker.$emit('pick', date)
            }
          }, {
            text: '一周前',
            onClick (picker) {
              const date = new Date()
              date.setTime(date.getTime() - 3600 * 1000 * 24 * 7)
              picker.$emit('pick', date)
            }
          }]
        },
        startDjsj: '',
        endDjsj: '',
        chooseCmcEmail: '',
        kind: '',
        region: '',
        chooseCmcType: [{id: 0, value: '充币'}, {id: 1, value: '提币'}],
        dataForm: {
        },
        // 交易代币列表
        currencyList: [],
        // 列表数据
        dataList: [],
        // 第一页
        pageIndex: 1,
        // 每页显示数量
        pageSize: 10,
        // 有多少页
        totalPage: 0,
        // 等待Loading加载
        dataListLoading: false
      }
    },
    created () {
      this.init()
    },
    methods: {
      init () {
        this.getInfo()
        this.getDataList()
      },
      getInfo () {
        // 列表数据请求头
        var params = {}
        // 发起请求
        API.transaction.currList(params).then(({data}) => {
          // 请求成功
          if (data && data.code === 0) {
            // 赋值
            this.currencyList = data.currencyList
          }
        })
      },
      // 复制成功提示语
      copy () {
        let clipboard = new Clipboard('.tag-read')
        clipboard.on('success', e => {
          console.log('成功')
          // this.$layer.alert('复制成功!')
          clipboard.destroy()
        })
        clipboard.on('error', e => {
          console.log('失败')
          clipboard.destroy()
        })
      },
      open9 () {
        this.$notify({
          message: '复制成功',
          position: 'bottom-left',
          type: 'success'
        })
      },
      // 获取数据列表
      getDataList () {
        // 打开Loading 加载
        this.dataListLoading = true
        // dffdsdfg
        if (this.startDjsj === '') {
          this.startDjsj = null
        }
        if (this.endDjsj === '') {
          this.endDjsj = null
        }
        if (this.chooseCmcEmail === '') {
          this.chooseCmcEmail = null
        }
        if (this.kind === '') {
          this.kind = null
        }
        if (this.region === '') {
          this.region = null
        }
        // 列表数据请求头
        var params = {
          // 第一页
          page: this.pageIndex,
          // 每页显示多少条数据
          limit: this.pageSize,
           // 开始日期时间
          startDjsj: this.startDjsj,
          // 结束日期时间
          endDjsj: this.endDjsj,
          // 查询邮箱
          cmcEmail: this.chooseCmcEmail,
          // 查询类型
          cmcType: this.kind,
          // 查询代币
          cmcCurrencyId: this.region
        }
        // 发起请求
        API.transaction.list(params).then(({data}) => {
          // 请求成功
          if (data && data.code === 0) {
            // 赋值操作
            // 格式化处理
            this.formatting(data.page.list)
            this.dataList = data.page.list
            // 反馈有多少页
            this.totalPage = data.page.totalCount
          } else {
            // 请求失败，数据为空
            this.dataList = []
            this.totalPage = 0
          }
          // 隐藏等待
          this.dataListLoading = false
        })
      },
      // 每页数
      sizeChangeHandle (val) {
        this.pageSize = val
        this.pageIndex = 1
        this.getDataList()
      },
      // 当前页
      currentChangeHandle (val) {
        this.pageIndex = val
        this.getDataList()
      },
      // 格式化
      formatting (array) {
        let arr = []
        for (let i = 0; i < array.length; i++) {
          // 格式化类型
          if (array[i].cmcType === 0) {
            array[i].cmcType = '充币'
          } else if (array[i].cmcType === 1) {
            array[i].cmcType = '提币'
          }
          // 格式化状态
          if (array[i].cmcZt === 0) {
            array[i].cmcZt = '通过'
          } else if (array[i].cmcZt === 1) {
            array[i].cmcZt = '审核中'
          } else if (array[i].cmcZt === 2) {
            array[i].cmcZt = '失败'
          }
        }
        return arr
      },
      // excel导出数据
      exportExcel () {
        this.dataListLoading = true
        let params = {
          'page': this.pageIndex,
          'limit': this.pageSize,
          'startDjsj': this.startDjsj,
          'endDjsj': this.endDjsj,
          'cmcEmail': this.chooseCmcEmail,
          'cmcType': this.kind,
          'cmcCurrencyId': this.region
        }
        API.transaction.exportExcel(params).then(({data}) => {
          if (data && data.code === 0) {
            window.location.href = data.url + '?token=' + Vue.cookie.get('token')
          } else {
            this.$message.error(data.msg)
          }
          this.dataListLoading = false
        })
        // let url = 'http://192.168.2.115:8080/qiubi/transactionRecords/czjl/exportExcel'
        // axios.get(url, params, {responseType: 'arraybuffer'}).then((res) => {
        //   let blob = new Blob([res.data], {type: 'application/vnd.ms-excel'})
        //   let objectUrl = URL.createObjectURL(blob)
        //   window.location.href = objectUrl
        // }).catch(function (res) {

        // })
      }
    }
  }
</script>
