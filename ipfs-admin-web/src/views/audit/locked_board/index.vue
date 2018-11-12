<template>
  <!-- 系统管理 / 管理员列表 -->
  <div class="mod-user">
    <!-- 搜索表单对象  -->
    <el-form :inline="true" :model="dataForm" @keyup.enter.native="getDataList()">
      <!-- 搜索栏目输入框 -->
      <el-form-item>
        <!-- 搜索 关键词 -->
        <el-input v-model="dataForm.cmcEmail" placeholder="用户名" clearable></el-input>
      </el-form-item>
      <!-- 功能按钮 -->
      <el-form-item>
        <!-- 调用查询，新增，删除功能 -->
        <el-button @click="getEList()">查询</el-button>                
      </el-form-item>
    </el-form>    
    <!-- Loading 加载 v-loading="xxx"-->
    <el-table :data="dataList" border v-loading="dataListLoading" @selection-change="selectionChangeHandle" style="width: 100%;">            
      <!-- 币种 -->
      <el-table-column prop="cmcCurrencyNumber" header-align="center" align="center" width="80" label="币种"></el-table-column>
      <!-- 锁仓数量 -->
      <el-table-column prop="cmcScsl" header-align="center" align="center" width="90" label="锁仓数量"></el-table-column>
      <!-- 邮箱 -->
      <el-table-column prop="cmcEmail" header-align="center" align="center" width="220" label="邮箱 (双击加入搜索栏)" sortable>
        <template slot-scope="scope">
          <p class="tag-read" :data-clipboard-text="scope.row.cmcEmail" @dblclick="copy(scope.row.cmcEmail)">{{scope.row.cmcEmail}}</p>
        </template>
      </el-table-column>
      <!-- <el-table-column prop="cmcEmail" header-align="center" align="center" label="邮箱"></el-table-column> -->
      <!-- 锁仓时间 -->
      <el-table-column prop="cmcSj" header-align="center" align="center" width="160" label="锁仓时间"></el-table-column>
      <!-- 到期时间 -->
      <el-table-column prop="cmcDqsj" header-align="center" align="center" width="160" label="到期时间"></el-table-column>
      <!-- 期限 -->
      <el-table-column prop="cmcQxday" header-align="center" align="center" width="80" label="期限"></el-table-column>
      <!-- 利息 -->
      <el-table-column prop="cmcLx" header-align="center" align="center" width="80" label="利息"></el-table-column>
      <!-- 到期可得 -->
      <el-table-column prop="cmcDqkdNumber" header-align="center" align="center" width="80" label="到期可得"></el-table-column>
      <el-table-column fixed="right" header-align="center" align="center" width="150" label="操作">
        <template slot-scope="scope">          
          <el-button v-if="scope.row.type===0" type="text" size="small">暂未到期</el-button>      
          <el-button v-else type="success" size="small" @click="deleteHandle(scope.row.cmcId)">解锁</el-button>
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
        // form表单对象
        dataForm: {
          cmcEmail: ''
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
        dataListLoading: false
      }
    },
    created () {
      this.init()
    },
    methods: {
      init () {
        this.getDataList()
      },
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
          // 邮箱账号 搜索条件(选题项)
          cmcEmail: this.dataForm.cmcEmail
        }
        // 发起请求
        API.transaction.scList(params).then(({data}) => {
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
      // 删除（单删，批量删除
      deleteHandle (id) {
        // 改进，将id换为名称
        var userIds = id ? [id] : this.dataListSelections.map(item => {
          return item.cmcId
        })
        this.$confirm('确定对id为 ' + userIds + ' 的记录进行解锁操作?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          // 确定删除
          API.transaction.scRemove(userIds).then(({data}) => {
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
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '取消解锁操作'
          })
        })
      },
      copy (e) {
        // 双击将邮箱赋值到搜索框中
        this.dataForm.cmcEmail = e
      },
      //  按用户名查询
      getEList () {
        this.getDataList()
      }
    }
  }
</script>
