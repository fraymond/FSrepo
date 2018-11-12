<template>
  <el-dialog
    title=通讯录
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" ref="dataForm" label-width="150px">
      <el-table :data="dataList" border v-loading="dataListLoading" style="width: 100%;">
        <el-table-column prop="cmcNumber" header-align="center" align="center" width="" label="币种编号" sortable></el-table-column>
        <el-table-column prop="cmcAddress" header-align="center" align="center" width="500" label="钱包地址" sortable></el-table-column>
        <el-table-column prop="cmcNotes" header-align="center" align="center" width="" label="备注" sortable></el-table-column>        
      </el-table>
      <!--                              每页数                             当前页-->
      <el-pagination @size-change="sizeChangeHandle" @current-change="currentChangeHandle" :current-page="pageIndex" :page-sizes="[10, 20, 50, 100]" :page-size="pageSize" :total="totalPage" layout="total, sizes, prev, pager, next, jumper"></el-pagination>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button @click="visible = false">取消</el-button>
    </span>
  </el-dialog>
</template>

<script>
  import API from '@/api'
  export default {
    data () {
      return {
        visible: false,
        roleList: [],
        // 用户id
        cmcYhglId: '',
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
        dataForm: {
          id: 0
        }
      }
    },
    methods: {
      init (id) {
        this.dataForm.id = id || 0
        this.visible = true
        this.cmcYhglId = id
        if (this.cmcYhglId) {
          this.getDataList()
        }
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
          // 用户id
          cmcYhglId: this.cmcYhglId
        }
        // 发起请求
        API.userManagement.getTxlInfo(params).then(({data}) => {
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
      }
    }
  }
</script>
