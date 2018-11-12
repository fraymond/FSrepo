<template>
  <el-dialog
    title=资产
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" ref="dataForm" label-width="150px">
      <el-table :data="dataList" border v-loading="dataListLoading" style="width: 100%;">
        <el-table-column prop="cmcNumber" header-align="center" align="center" width="" label="货币编号" sortable></el-table-column>
        <el-table-column prop="cmcYe" header-align="center" align="center" width="" label="可用" sortable></el-table-column>
        <el-table-column prop="cmcSd" header-align="center" align="center" width="" label="锁定" sortable></el-table-column>
        <el-table-column fixed="right" header-align="center" align="center" width="" label="操作">
          <template slot-scope="scope">
            <el-button v-show="scope.row.cmcCurrencyId===12" type="text" size="small" @click="pay(scope.row.cmcCurrencyId)">充值</el-button>
          </template>
        </el-table-column>        
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
        // cmcAutoid
        cmcAutoid: '',
        // 充值数量
        payValue: '',
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
        API.userManagement.getZcInfo(params).then(({data}) => {
          console.log(data)
          // 请求成功
          if (data && data.code === 0) {
            // 赋值操作
            this.dataList = data.page.list
            this.cmcAutoid = this.getCmcAutoid(this.dataList)
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
      getCmcAutoid (array) {
        let val
        for (let i = 0; i < array.length; i++) {
          if (array[i].cmcCurrencyId === 12) {
            val = array[i].cmcAutoid
          }
        }
        return val
      },
      pay () {
        // cmcYhglId+++ 请输入充值数量
        this.$prompt('请输入充值数量', '充值', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          required: true,
          inputPattern: /^([1-9]\d*(\.\d*[1-9])?)|(0\.\d*[1-9])$/,
          inputErrorMessage: '充值数量必须大于0'
        }).then(({ value }) => {
          this.payValue = parseInt(value)
          this.paySome()
          this.$message({
            type: 'success',
            message: '充值成功，充值数量为: ' + value
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '取消充值'
          })
        })
      },
      paySome () {
        // 列表数据请求头
        var params = {
          // xx
          cmcYe: this.payValue,
          // 用户id
          cmcAutoid: this.cmcAutoid
        }
        // 发起请求
        API.userManagement.updateZc(params).then(({data}) => {
          console.log(data)
          // 请求成功
          if (data && data.code === 0) {
            // 赋值操作
            // dataList
            this.getDataList()
          } else {
            // 请求失败，数据为空
          }
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
      }
    }
  }
</script>
