<template>
  <!-- 系统管理 / 管理员列表 -->
  <div class="mod-data_maintain">
    <!--                    搜索表单对象  -->
    <el-form :inline="true" :model="dataForm" @keyup.enter.native="getDataList()">
      <!-- 功能按钮 -->
      <el-form-item>
        <!-- 调用查询，新增，删除功能 -->
        <el-button v-if="isAuth('sys:user:save')" type="primary" @click="addOrUpdateHandle()">新增</el-button>
        <el-button v-if="isAuth('sys:user:delete')" type="danger" @click="deleteHandle()" :disabled="dataListSelections.length <= 1">批量删除</el-button>
      </el-form-item>
    </el-form>
    <!-- Loading 加载 v-loading="xxx"-->
    <el-table :data="dataList" border v-loading="dataListLoading" @selection-change="selectionChangeHandle" style="width: 100%;" tooltip-effect="dark">
      <!-- 勾选框 -->
      <el-table-column type="selection" header-align="center" align="center" width="50"></el-table-column>
      <!-- 系列 -->
      <el-table-column prop="cmcSeries" header-align="center" align="center" width="150" label="系列"></el-table-column>
      <!-- 图标 -->
      <el-table-column header-align="center" align="center" width="100" label="图标">
        <template slot-scope="scope">
          <img :src="scope.row.cmcTburl" />
        </template>
      </el-table-column>
      <!-- 编号 -->
      <el-table-column prop="cmcNumber" header-align="center" align="center" width="150" label="编号"></el-table-column>
      <!-- 是否为交易货币 -->
      <el-table-column prop="cmcJy" header-align="center" align="center" width="150" label="是否为交易货币">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.cmcJy === 0" size="small" type="danger">否</el-tag>
          <el-tag v-else size="small">是</el-tag>
        </template>
      </el-table-column>
      <!-- 是否上架交易所 -->
      <el-table-column prop="cmcJyzt" header-align="center" align="center" width="150" label="是否上架交易所">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.cmcJyzt === 0" size="small" type="danger">否</el-tag>
          <el-tag v-else size="small">是</el-tag>
        </template>
      </el-table-column>
      <!-- 能否充币 -->
      <el-table-column prop="cmcCb" header-align="center" align="center" width="150" label="能否充币">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.cmcCb === 0" size="small" type="danger">否</el-tag>
          <el-tag v-else size="small">是</el-tag>
        </template>
      </el-table-column>
      <!-- 能否提币 -->
      <el-table-column prop="cmcTb" header-align="center" align="center" width="150" label="能否提币">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.cmcTb === 0" size="small" type="danger">否</el-tag>
          <el-tag v-else size="small">是</el-tag>
        </template>
      </el-table-column>
      <!-- 交易手续费(%) -->
      <el-table-column prop="cmcSxf" header-align="center" align="center" width="150" label="交易手续费(%)"></el-table-column>
      <!-- 提币手续费(%) -->
      <el-table-column prop="cmcTbsxf" header-align="center" align="center" width="150" label="提币手续费(%)"></el-table-column>
      <!-- 最低可提 -->
      <el-table-column prop="cmcDbzd" header-align="center" align="center" width="150" label="最低可提"></el-table-column>
      <!-- 优先级 -->
      <el-table-column prop="cmcPriority" header-align="center" align="center" width="150" label="优先级"></el-table-column>
      <el-table-column fixed="right" header-align="center" align="center" width="150" label="操作">
        <template slot-scope="scope">
          <el-button v-if="isAuth('sys:user:update')" type="text" size="small" @click="addOrUpdateHandle(scope.row.cmcAutoid)">修改</el-button>
          <el-button v-if="isAuth('sys:user:delete')" type="text" size="small" @click="deleteHandle(scope.row.cmcAutoid)">删除</el-button>
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
          limit: this.pageSize,
          // 用户名
          username: this.dataForm.userName
        }
        // 发起请求
        API.dataMaintain.list(params).then(({data}) => {
          console.log(data)
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
          return item.cmcAutoid
        })
        // 改进，将id换为名称
        this.$confirm(`确定对选中项进行[${id ? '删除' : '批量删除'}]操作?`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          // 确定删除
          API.dataMaintain.del(userIds).then(({data}) => {
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
