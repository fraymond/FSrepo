<template>
  <!-- 数据分析 / 收益记录列表 -->
  <div class="mod-revenue_record_statistics">
    <!--                    搜索表单对象  -->
    <el-form :inline="true" :model="dataForm" @keyup.enter.native="getDataList()">
      <!-- 搜索栏目输入框 -->
      <el-form-item>
        <!--                搜索 关键词 -->
        <!-- <el-input v-model="dataForm.userName" placeholder="用户名" clearable></el-input> -->
      </el-form-item>
      <!-- 功能按钮 -->
      <el-form-item>
        <!-- 调用查询，新增，删除功能 -->
        <!-- <el-button @click="getEList()">查询</el-button> -->
        <!-- <el-button v-if="isAuth('sys:user:save')" type="primary" @click="addOrUpdateHandle()">新增</el-button> -->
        <!-- <el-button v-if="isAuth('sys:user:delete')" type="danger" @click="deleteHandle()" :disabled="dataListSelections.length <= 0">批量删除</el-button> -->
      </el-form-item>
    </el-form>
    <!-- Loading 加载 v-loading="xxx"-->
    <el-table :data="dataList" border v-loading="dataListLoading" @selection-change="selectionChangeHandle" style="width: 100%;">
      <!-- 勾选框 -->
      <!-- <el-table-column type="selection" header-align="center" align="center" width="50"></el-table-column> -->
      <!-- 币种 -->
      <el-table-column prop="cmcNumber" header-align="center" align="center" width="" label="币种"></el-table-column>      
      <!-- 用户 -->
      <el-table-column prop="cmcEmail" header-align="center" align="center" label="用户"></el-table-column>
      <!-- 时间 -->
      <el-table-column prop="cmcSj" header-align="center" align="center" width="" label="时间"></el-table-column>
      <!-- 类型 -->
      <!-- <el-table-column prop="cmcType" header-align="center" align="center" width="" label="类型"></el-table-column> -->
      <el-table-column prop="cmcType" header-align="center" align="center" width="" label="类型">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.cmcCb === 0" size="small" type="danger">交易</el-tag>
          <el-tag v-else size="small">提币</el-tag>
        </template>
      </el-table-column>
      <!-- 提币id -->
      <el-table-column prop="cmcCzjlId" header-align="center" align="center" label="提币id"></el-table-column>
      <!-- 交易id -->
      <el-table-column prop="cmcTransBillId" header-align="center" align="center" label="交易id"></el-table-column>
      <!-- 数量 -->
      <el-table-column prop="cmcQuantity" header-align="center" align="center" label="数量"></el-table-column>
      <!-- <el-table-column fixed="right" header-align="center" align="center" width="150" label="操作">
        <template slot-scope="scope">
          <el-button v-if="isAuth('sys:user:update')" type="text" size="small" @click="addOrUpdateHandle(scope.row.cmcAutoid)">修改</el-button>
          <el-button v-if="isAuth('sys:user:delete')" type="text" size="small" @click="deleteHandle(scope.row.cmcAutoid)">删除</el-button>
        </template>
      </el-table-column> -->
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
          limit: this.pageSize,
          // 用户名
          username: this.dataForm.userName
        }
        // 发起请求
        API.revenue.list(params).then(({data}) => {
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
      // 格式化
      formatting (array) {
        for (let i = 0; i < array.length; i++) {
          // 格式化类型
          if (array[i].cmcType === 0) {
            array[i].cmcType = '交易'
          } else if (array[i].cmcType === 1) {
            array[i].cmcType = '提币'
          }
        }
        return array
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
          return item.cmcAutoid
        })
        // 改进，将id换为名称
        this.$confirm(`确定对[id=${userIds.join(',')}]进行[${id ? '删除' : '批量删除'}]操作?`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          // 确定删除
          API.user.del(userIds).then(({data}) => {
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
