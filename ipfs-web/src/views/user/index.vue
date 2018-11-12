<template>
  <div class="home-wapper">
    <h2>我的地址列表</h2>
    <el-table
      :data="dataList"
      border
      style="width: 100%;">
      <el-table-column
        prop="fileHash"
        header-align="center"
        align="center"
        width="600"
        label="我的地址">
      </el-table-column>
      <el-table-column
        prop="operationTime"
        header-align="center"
        align="center"
        width="450"
        label="创建/导入时间">
        <template slot-scope="scope">
          <p>{{(scope.row.operationTime)}}</p>
        </template>
      </el-table-column>
      <el-table-column
        prop="operationType"
        header-align="center"
        align="center"
        width="200"
        label="操作">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        dataList: []
      }
    },
    created () {
      this.init()
    },
    methods: {
      init () {
        let userInfo = localStorage.getItem('userInfo')
        this.data = userInfo
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