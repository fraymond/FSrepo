<template>
  <!-- 子链管理 -->
  <div class="mod-user">
    <!-- 搜索表单对象  -->
    <el-form :inline="true" :model="dataForm" @keyup.enter.native="getDataList()">
      <!-- 功能按钮 -->
      <el-form-item>
        <!-- 调用查询，新增，删除功能 -->
        <!-- <el-button @click="getEList()">查询</el-button> -->
        <!-- <el-button v-if="isAuth('sys:user:save')" type="primary" @click="addOrUpdateHandle()">新增</el-button> -->
      </el-form-item>
    </el-form>
    <!-- Loading 加载 v-loading="xxx"-->
    <el-table :data="dataList" border v-loading="dataListLoading" @selection-change="selectionChangeHandle" style="width: 100%;">
      <!-- 勾选框 -->
      <el-table-column type="selection" header-align="center" align="center" width="50"></el-table-column>
      <!-- 操作系统用户名 -->
      <el-table-column prop="sysUserName" header-align="center" align="center" label="操作系统用户名"></el-table-column>
      <!-- 存储子链地址 -->
      <el-table-column prop="subchainAddress" header-align="center" align="center" label="存储子链地址"></el-table-column>      
      <!-- 操作类型 -->
      <el-table-column prop="operationType" header-align="center" align="center" label="操作类型"></el-table-column>
      <!-- 操作时间 -->
      <el-table-column prop="operationTime" header-align="center" align="center" label="操作时间"></el-table-column>
    </el-table>
    <!-- 分页 每页数 当前页-->
    <el-pagination v-show="totalPage>0" @size-change="sizeChangeHandle" @current-change="currentChangeHandle" :current-page="pageIndex" :page-sizes="[10, 20, 50, 100]" :page-size="pageSize" :total="totalPage" layout="total, sizes, prev, pager, next, jumper"></el-pagination>
  </div>
</template>

<script>
  import API from '@/api'
  // import AddOrUpdate from './add-or-update'
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
        // 选中 数据列表
        dataListSelections: [],
        // 是否显示新增，修改弹窗组件
        addOrUpdateVisible: false
      }
    },
    // 新增or修改组件
    // components: {
    //   AddOrUpdate
    // },
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
        API.user.logList(params).then(({data}) => {
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
      // 多选
      selectionChangeHandle (val) {
        this.dataListSelections = val
      },
      // 新增 / 修改
      addOrUpdateHandle (id) {
        this.addOrUpdateVisible = true
        this.$nextTick(() => {
          this.$refs.addOrUpdate.init(id)
        })
      },
      // 删除（单删，批量删除
      deleteHandle (obj) {
        // 改进，将id换为名称
        this.$confirm(`确定进行删除操作?`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          // 确定删除
          API.user.del1(obj).then(({data}) => {
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
        })
      },
      //  按用户名查询
      getEList () {
        this.getDataList()
      }
    }
  }
</script>
