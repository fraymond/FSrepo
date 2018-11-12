<template>
  <!-- 交易管理 / 锁仓记录 -->
  <div class="mod-user">      
    <!-- Loading 加载 v-loading="xxx"-->
    <el-table :data="dataList" border v-loading="dataListLoading" @selection-change="selectionChangeHandle" style="width: 100%;">
      <!-- 勾选框 -->
      <el-table-column type="selection" header-align="center" align="center" width="50"></el-table-column>      
      <!-- 币种编号 -->
      <el-table-column prop="cmcCurrencyNumber" header-align="center" align="center" width="100" label="币种编号"></el-table-column>      
      <!-- 锁仓数量 -->
      <el-table-column prop="cmcScsl" header-align="center" align="center" label="锁仓数量"></el-table-column>
      <!-- 用户 -->
      <el-table-column prop="cmcEmail" header-align="center" align="center" label="用户"></el-table-column>      
      <!-- 锁仓时间 -->
      <el-table-column prop="cmcSj" header-align="center" align="center" label="锁仓时间"></el-table-column>
      <!-- 到期时间 -->
      <el-table-column prop="cmcDqsj" header-align="center" align="center" label="到期时间"></el-table-column>
      <!-- 期限 -->
      <el-table-column prop="cmcQxday" header-align="center" align="center" label="期限"></el-table-column>
      <!-- 利息 -->
      <el-table-column prop="cmcLx" header-align="center" align="center" label="利息"></el-table-column>
      <!-- 到期可得 -->
      <el-table-column prop="cmcDqkdNumber" header-align="center" align="center" label="到期可得"></el-table-column>
      <!-- 状态 -->
      <el-table-column prop="cmcState" header-align="center" align="center" label="状态">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.cmcState === 0" size="small" type="danger">锁定</el-tag>
          <el-tag v-else size="small">解锁</el-tag>
        </template>
      </el-table-column>      
    </el-table>
    <!--                              每页数                             当前页-->
    <el-pagination @size-change="sizeChangeHandle" @current-change="currentChangeHandle" :current-page="pageIndex" :page-sizes="[10, 20, 50, 100]" :page-size="pageSize" :total="totalPage" layout="total, sizes, prev, pager, next, jumper"></el-pagination>    
  </div>
</template>

<script>
  import API from '@/api'
  export default {
    data () {
      return {
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
    created () {
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
        API.transaction.scList(params).then(({data}) => {
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
      }
    }
  }
</script>
