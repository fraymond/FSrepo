<template>
  <!-- 系统管理 / 管理员列表 -->
  <div class="mod-introduction">
    <!--                    搜索表单对象  -->
    <el-form :inline="true" :model="dataForm" @keyup.enter.native="getDataList()">
      <!-- 功能按钮 -->
      <el-form-item>
        <el-button v-if="isAuth('sys:user:save')" type="primary" @click="addOrUpdateHandle()">新增</el-button>
        <el-button v-if="isAuth('sys:user:delete')" type="danger" @click="deleteHandle()" :disabled="dataListSelections.length <= 1">批量删除</el-button>
      </el-form-item>
    </el-form>
    <!-- Loading 加载 v-loading="xxx"-->
    <el-table :data="dataList" border v-loading="dataListLoading" @selection-change="selectionChangeHandle" style="width: 100%;" tooltip-effect="dark">
      <!-- 勾选框 -->
      <el-table-column type="selection" header-align="center" align="center" width="40"></el-table-column>
      <!-- 图标 -->
      <el-table-column header-align="center" align="center" width="60" label="图标">
        <template slot-scope="scope">
          <img :src="scope.row.cmcWhUrl" />
        </template>
      </el-table-column>
      <!-- 币种编号 -->
      <el-table-column prop="cmcNumber" header-align="center" align="center" width="60" label="币种编号"></el-table-column>
      <!-- 中文名 -->
      <el-table-column prop="zhName" header-align="center" align="center" width="90" label="中文名" sortable></el-table-column>
      <!-- 英文名 -->
      <el-table-column prop="enName" header-align="center" align="center" width="" label="英文名"></el-table-column>
      <!-- 缩写 -->
      <el-table-column prop="sxName" header-align="center" align="center" width="60" label="缩写"></el-table-column>
      <!-- 币种介绍 -->
      <el-table-column prop="intro" header-align="center" align="center" width="200" label="币种介绍" show-overflow-tooltip></el-table-column>
      <!-- 推出时间 -->
      <el-table-column prop="addTime" header-align="center" align="center" width="156" label="推出时间" sortable></el-table-column>
      <!-- 发行总量 -->
      <el-table-column prop="fxCount" header-align="center" align="center" width="110" label="发行总量" sortable></el-table-column>
      <!-- 存量 -->
      <el-table-column prop="clCount" header-align="center" align="center" width="" label="存量" sortable></el-table-column>
      <!-- 市值 -->
      <el-table-column prop="sValue" header-align="center" align="center" width="" label="市值" sortable></el-table-column>
      <!-- 币种算法 -->
      <el-table-column prop="coinSf" header-align="center" align="center" width="" label="币种算法"></el-table-column>
      <!-- 区块速度 -->
      <el-table-column prop="speed" header-align="center" align="center" width="110" label="区块速度" sortable></el-table-column>
      <!-- 区块大小 -->
      <el-table-column prop="bigMin" header-align="center" align="center" width="60" label="区块大小" sortable></el-table-column>
      <!-- 调整难度 -->
      <el-table-column prop="difficult" header-align="center" align="center" width="50" label="调整难度" sortable></el-table-column>  
      <!-- 证明方式 -->
      <el-table-column prop="certify" header-align="center" align="center" width="60" label="证明方式"></el-table-column>
      <!-- 链接 -->
      <el-table-column prop="connects" header-align="center" align="center" width="" label="链接"></el-table-column>
      <el-table-column fixed="right" header-align="center" align="center" width="150" label="操作">
        <template slot-scope="scope">
          <el-button v-if="isAuth('sys:user:update')" type="text" size="small" @click="addOrUpdateHandle(scope.row.id)">修改</el-button>
          <el-button v-if="isAuth('sys:user:delete')" type="text" size="small" @click="deleteHandle(scope.row.id)">删除</el-button>
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
        // 选中 数据列表
        dataListSelections: [],
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
        API.introduction.list(params).then(({data}) => {
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
      deleteHandle (id) {
        var userIds = id ? [id] : this.dataListSelections.map(item => {
          return item.id
        })
        // 改进，将id换为名称 [id=${userIds.join(',')}]
        this.$confirm(`确定对选中项进行[${id ? '删除' : '批量删除'}]操作?`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          // 确定删除
          API.introduction.del(userIds).then(({data}) => {
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
      }
    }
  }
</script>
