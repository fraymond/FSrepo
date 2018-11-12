<template>
  <div class="home-wapper">
    <!-- 搜索表单对象  -->
    <el-form :inline="true" :model="dataForm" @keyup.enter.native="getDataList()">
      <!-- 搜索栏目输入框 -->
      <el-form-item>
        <!-- 搜索 关键词 -->
        <el-input v-model="dataForm.fileAddress" style="width:420px;float:left;margin-right:20px;" placeholder="请输入要查询的文件hash地址" clearable></el-input>
        <el-input v-model="dataForm.fileName" style="width:300px;float:left;" placeholder="请输入要查询的文件名" clearable></el-input>
      </el-form-item>
      <!-- 功能按钮 -->
      <el-form-item>
        <!-- 调用查询，新增，删除功能 -->
        <el-button @click="getEList()">查询</el-button>
        <el-button v-if="isAuth('sys:user:save')" type="primary" @click="addOrUpdateHandle()">新增</el-button>
      </el-form-item>
    </el-form>
    <el-table
      :data="dataList"
      border
      v-loading="dataListLoading"
      style="width: 100%;">
      <el-table-column
        prop="hash"
        header-align="center"
        align="center"
        min-width="32%"        
        label="文件hash地址">
      </el-table-column>
      <el-table-column
        prop="createTime"
        header-align="center"
        align="center"
        min-width="15%"
        label="创建时间">
        <template slot-scope="scope">
          <p>{{timestampToTime(scope.row.createTime)}}</p>
        </template>
      </el-table-column>
      <el-table-column
        prop="fileName"
        header-align="center"
        align="center"
        min-width="15%"
        label="文件名">
      </el-table-column>
      <el-table-column
        prop="fileSize"
        header-align="center"
        align="center"
        min-width="8%"        
        label="文件大小">
      </el-table-column>
      <el-table-column
        prop="remark"
        header-align="center"
        align="center"
        min-width="10%"
        label="备忘内容">
      </el-table-column>
      <el-table-column
        prop="encrypt"
        header-align="center"
        align="center"
        min-width="10%"        
        label="加密状态">
        <!-- 待确认 -->
        <template slot-scope="scope">
          <el-tag v-if="scope.row.encrypt === 1" size="small" type="danger">加密</el-tag>
          <el-tag v-else size="small">不加密</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        fixed="right"
        header-align="center"
        align="center"
        min-width="10%"        
        label="操作">
        <template slot-scope="scope">
          <el-button type="text" size="small" @click="lookfile(scope.row.hash)">查看</el-button>
          <el-button type="text" size="small" @click="deletefile(scope.row.hash)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      v-show="totalPage > 10"
      @size-change="sizeChangeHandle"
      @current-change="currentChangeHandle"
      :current-page="pageIndex"
      :page-sizes="[10, 20, 50, 100]"
      :page-size="pageSize"
      :total="totalPage"
      layout="total, sizes, prev, pager, next, jumper">
    </el-pagination>
    <!-- 弹窗, 新增 / 修改 -->
    <look-file v-if="addOrUpdateVisible" ref="addOrUpdate" @refreshDataList="getDataList"></look-file>
  </div>
</template>

<script>
  import API from '@/api'
  import LookFile from './lookfile'
  export default {
    data () {
      return {
        dataList: [],
        pageIndex: 1,
        pageSize: 10,
        totalPage: 0,
        address: '',
        dataListLoading: false,
        addOrUpdateVisible: false,
        dataListSelections: [],
        userInfo: [],
        dataForm: {
          fileAddress: '',
          fileName: ''
        }
      }
    },
    created () {
      this.init()
    },
    methods: {
      init () {
        this.userInfo = JSON.parse(localStorage.getItem('userInfo'))
        if (this.userInfo !== null) {
          this.getDataList()
        } else {
          this.$alert('查看文件列表需要导入地址', '温馨提示', {
            confirmButtonText: '立即导入'
          })
          this.$router.push({ path: '/login' })
        }
      },
      // 获取文件列表数据
      getDataList () {
        this.userInfo = JSON.parse(localStorage.getItem('userInfo'))
        let userId = this.userInfo.userId
        this.address = this.userInfo.address
        this.dataListLoading = true
        var params = {
          page: this.pageIndex,
          limit: this.pageSize,
          hash: this.dataForm.fileAddress,
          fileName: this.dataForm.fileName,
          userId: userId
        }
        API.main.fileList(params).then(({data}) => {
          if (data && data.code === 0) {
            this.dataList = data.userFileList.list
            this.totalPage = data.userFileList.totalCount
          } else {
            this.dataList = []
            this.totalPage = 0
          }
          this.dataListLoading = false
        })
      },
      getEList () {
        this.getDataList()
      },      // 每页数
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
      // 查看文件
      lookfile (id) {
        this.dataListLoading = true
        let params = {
          hash: id,
          address: this.address
        }
        API.main.queryFile(params).then(({data}) => {
          if (data && data.code === 0) {
            this.dataListLoading = false
            window.open(data.data, '_blank')
            this.addOrUpdateVisible = true
            this.$nextTick(() => {
              this.$refs.addOrUpdate.init(data.data)
            })
          } else {
            this.$message.error(data.msg)
          }
        })
      },
      // 删除
      deletefile (id) {
        this.$confirm(`确定要进行删除操作?`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          let params = {
            hash: id,
            address: this.address
          }
          this.dataListLoading = true
          API.main.deleteFile(params).then(({data}) => {
            if (data && data.code === 0) {
              this.dataListLoading = false
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
      // 转换时间戳
      timestampToTime (timestamp) {
        for (let i = 0; i < timestamp.length; i++) {
          console.log(timestamp[i].operationTime)
        }
        var date = new Date(timestamp) //  时间戳为10位需*1000，时间戳为13位的话不需乘1000
        var Y = date.getFullYear() + '-'
        var M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-'
        var D = date.getDate() + ' '
        var h = date.getHours() + ':'
        var m = date.getMinutes() + ':'
        var s = date.getSeconds()
        return Y + M + D + h + m + s
      }
    },
    components: {
      LookFile
    }
  }
</script>

<style>
  .home-wapper {
    min-width: 1330px;
    margin: 0 auto;
    padding: 0 18%;
  }
  .home-wapper .el-table {
    margin-bottom: 3%;
  }
</style>

