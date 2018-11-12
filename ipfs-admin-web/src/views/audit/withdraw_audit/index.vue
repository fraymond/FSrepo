<template>
  <!-- 系统管理 / 管理员列表 -->
  <div class="mod-user">    
    <!-- Loading 加载 v-loading="xxx"-->
    <el-table :data="dataList" border v-loading="dataListLoading" style="width: 100%;">
      <!-- 提币人 -->
      <el-table-column prop="cmcEmail" header-align="center" align="center" width="250" label="提币人"></el-table-column>
      <!-- 币种 -->
      <el-table-column prop="cmcNumber" header-align="center" align="center" width="150" label="币种"></el-table-column>
      <!-- 申请时间 -->
      <el-table-column prop="cmcDjsj" header-align="center" align="center" width="200" label="申请时间"></el-table-column>
      <!-- 手续费 -->
      <el-table-column prop="cmcSxf" header-align="center" align="center" width="120" label="手续费"></el-table-column>
      <!-- 提币数量 -->
      <el-table-column prop="cmcSl" header-align="center" align="center" width="150" label="提币数量"></el-table-column>
      <!-- 实际到帐 -->
      <el-table-column prop="cmcSjdz" header-align="center" align="center" width="150" label="实际到帐"></el-table-column>
      <!-- 发送地址 -->
      <el-table-column prop="cmcFsdz" header-align="center" align="center" width="330" label="发送地址"></el-table-column>
      <!-- 接受地址 -->
      <el-table-column prop="cmcJsdz" header-align="center" align="center" width="330" label="接受地址"></el-table-column>      
      <el-table-column fixed="right" header-align="center" align="center" width="150" label="操作">
        <template slot-scope="scope">
          <el-button v-if="isAuth('sys:user:update')" type="text" size="small" @click="passBtn(scope.row.cmcAutoid)">通过</el-button>
          <el-button v-if="isAuth('sys:user:delete')" type="text" size="small" @click="addOrUpdateHandle(scope.row.cmcAutoid)">驳回</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--                              每页数                             当前页-->
    <el-pagination @size-change="sizeChangeHandle" @current-change="currentChangeHandle" :current-page="pageIndex" :page-sizes="[10, 20, 50, 100]" :page-size="pageSize" :total="totalPage" layout="total, sizes, prev, pager, next, jumper"></el-pagination>
    <!-- 弹窗, 新增 / 修改 -->
    <!-- 弹窗组件（新增或者修改） -->
    <add-or-update v-if="addOrUpdateVisible" ref="addOrUpdate" @refreshDataList="getDataList"></add-or-update>
  </div>
</template>

<script>
  import API from '@/api'
  import AddOrUpdate from './add-or-update'
  export default {
    data () {
      return {
        // form表单对象
        dataForm: {
          userName: ''
        },
        // 列表数据
        dataList: [],
        // 第一页
        pageIndex: 1,
        // 每页显示数量
        pageSize: 10,
        // 有多少页
        totalPage: 0,
        // 等待Loading加载
        dataListLoading: false,
        // 是否显示新增，修改弹窗组件
        addOrUpdateVisible: false
      }
    },
    // 新增or修改组件
    components: {
      AddOrUpdate
    },
    // keep-alive组件激活时调用
    activated () {
      this.getDataList()
    },
    methods: {
      // 获取数据列表
      getDataList () {
        // 打开Loading 加载
        this.dataListLoading = true
        // 列表数据请求头
        var params = {
          // 第一页
          page: this.pageIndex,
          // 每页显示多少条数据
          limit: this.pageSize
        }
        // 发起请求
        API.transaction.tbAuditList(params).then(({data}) => {
          // 请求成功
          if (data && data.code === 0) {
            // 赋值操作
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
      // 通过
      passBtn (id) {
        // 列表数据请求头
        // var params = { cmcAutoid: id }
        var params = [id]
        API.transaction.checkPass(params).then(({data}) => {
          if (data && data.code === 0) {
            this.$message({
              message: '操作成功',
              type: 'success',
              duration: 1500,
              onClose: () => {
                this.getDataList()
              }
            })
          } else {
            this.$message.error(data.msg)
          }
        })
      },
      // 新增 / 修改
      addOrUpdateHandle (id) {
        this.addOrUpdateVisible = true
        this.$nextTick(() => {
          this.$refs.addOrUpdate.init(id)
        })
      }
    }
  }
</script>
