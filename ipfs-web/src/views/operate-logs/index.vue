<template>
  <div class="home-wapper">
    <el-table
      :data="dataList"
      border
      v-loading="dataListLoading"
      style="width: 100%;">
      <el-table-column
        prop="fileHash"
        header-align="center"
        align="center"
        min-width="45%"
        label="文件hash地址">
      </el-table-column>
      <el-table-column
        prop="operationTime"
        header-align="center"
        align="center"
        min-width="25%"        
        label="删除时间">
        <template slot-scope="scope">
          <p>{{timestampToTime(scope.row.operationTime)}}</p>
        </template>
      </el-table-column>
      <el-table-column
        prop="fileRemark"
        header-align="center"
        align="center"
        min-width="15%"
        label="备忘内容">
      </el-table-column>
      <el-table-column
        prop="operationType"
        header-align="center"
        align="center"
        min-width="15%"        
        label="操作类型">
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
  </div>
</template>

<script>
  import API from '@/api'
  export default {
    data () {
      return {
        dataList: [],
        pageIndex: 1,
        pageSize: 10,
        totalPage: 0,
        userInfo: [],
        address: '',
        dataListLoading: false,
        addOrUpdateVisible: false,
        dataListSelections: []
      }
    },
    created () {
      this.init()
    },
    methods: {
      // 判断用户是否登陆
      init () {
        this.userInfo = JSON.parse(localStorage.getItem('userInfo'))
        if (this.userInfo !== null) {
          this.getDataList()
        } else {
          this.$alert('查看操作日志需要导入地址', '温馨提示', {
            confirmButtonText: '立即导入'
          })
          this.$router.push({ path: '/login' })
        }
      },
      // 获取文件列表数据
      getDataList () {
        this.userInfo = JSON.parse(localStorage.getItem('userInfo'))
        let userId = this.userInfo.userId
        this.dataListLoading = true
        var params = {
          page: this.pageIndex,
          limit: this.pageSize,
          userId: userId
        }
        API.main.logList(params).then(({data}) => {
          if (data && data.code === 0) {
            // 时间格式化
            this.dataList = data.page.list
            this.totalPage = data.page.totalCount
          } else {
            this.dataList = []
            this.totalPage = 0
          }
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