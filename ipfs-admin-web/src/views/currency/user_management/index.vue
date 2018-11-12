<template>
<!-- 系统管理 / 管理员列表 -->
  <div class="mod-user_management">
    <!--                    搜索表单对象  -->
    <el-form :inline="true" :model="dataForm" @keyup.enter.native="getDataList()">
      <!-- 搜索栏目输入框 -->
      <el-form-item>
        <!--                搜索 关键词 -->
        <el-input v-model="dataForm.email" placeholder="邮箱" clearable style="width: 108%"></el-input>
      </el-form-item>
      <!-- 功能按钮 -->
      <el-form-item>
        <!-- 调用查询，新增，删除功能 -->
        <el-button @click="getEmailList()">查询</el-button>
      </el-form-item>
    </el-form>
    <!-- Loading 加载 v-loading="xxx"-->
    <el-table :data="dataList" border v-loading="dataListLoading" style="width: 100%;">          
      <!-- 邮箱 -->
      <el-table-column prop="cmcEmail" header-align="center" align="center" width="400%" label="邮箱 (双击加入搜索栏)" sortable>
        <template slot-scope="scope">
          <p class="tag-read" :data-clipboard-text="scope.row.cmcEmail" @dblclick="copy(scope.row.cmcEmail)">{{scope.row.cmcEmail}}</p>
        </template>
      </el-table-column>
      <!-- 注册时间 -->
      <el-table-column prop="cmcDjsj" header-align="center" align="center" width="" label="创建时间" sortable></el-table-column>
      <!-- 身份 -->
      <el-table-column prop="cmcVip" header-align="center" align="center" width="" label="身份" sortable></el-table-column>
      <!--  -->
      <el-table-column fixed="right" header-align="center" align="center" width="400%" label="操作">
        <template slot-scope="scope">
          <el-button v-if="isAuth('sys:user:update')" type="text" size="small" @click="addOrUpdateHandle(scope.row.cmcAutoid)">修改身份</el-button>
          <el-button v-if="isAuth('sys:user:delete')" type="text" size="small" @click="propertyHandle(scope.row.cmcAutoid)">资产</el-button>
          <el-button v-if="isAuth('sys:user:delete')" type="text" size="small" @click="walletAddressHandle(scope.row.cmcAutoid)">钱包地址</el-button>
          <el-button v-if="isAuth('sys:user:delete')" type="text" size="small" @click="addressBookHandle(scope.row.cmcAutoid)">通讯录</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--                              每页数                             当前页-->
    <el-pagination @size-change="sizeChangeHandle" @current-change="currentChangeHandle" :current-page="pageIndex" :page-sizes="[10, 20, 50, 100]" :page-size="pageSize" :total="totalPage" layout="total, sizes, prev, pager, next, jumper"></el-pagination>    
    <!-- 弹窗, 新增 / 修改 -->
    <!-- 弹窗组件（新增或者修改） -->
    <add-or-update v-if="addOrUpdateVisible" ref="addOrUpdate" @refreshDataList="getDataList"></add-or-update>
    <!-- 弹窗组件（资产） -->
    <property v-if="propertyVisible" ref="property" @refreshDataList="getDataList"></property>
    <!-- 弹窗组件（钱包地址） -->
    <wallet-address v-if="walletAddressVisible" ref="walletAddress" @refreshDataList="getDataList"></wallet-address>    
    <!-- 弹窗组件（通讯录） -->
    <address-book v-if="addressBookVisible" ref="addressBook" @refreshDataList="getDataList"></address-book>
  </div>
</template>
<script>
  import API from '@/api'
  import AddOrUpdate from './add-or-update'
  import Property from './property'
  import WalletAddress from './wallet-address'
  import AddressBook from './address-book'
  export default {
    data () {
      return {
        // form表单对象
        dataForm: {
          email: ''
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
        // 弹窗
        addOrUpdateVisible: false,
        // 资产弹窗
        propertyVisible: false,
        // 钱包地址弹窗
        walletAddressVisible: false,
        // 通讯录弹窗
        addressBookVisible: false
      }
    },
    // 新增or修改组件
    components: {
      AddOrUpdate,
      Property,
      WalletAddress,
      AddressBook
    },
    // keep-alive组件激活时调用
    activated () {
      this.getDataList()
    },
    // created () {
    //   this.getDataList()
    // },
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
        API.userManagement.list(params).then(({data}) => {
          // 请求成功
          if (data && data.code === 0) {
            // 赋值操作
            this.dataList = this.formatting(data.page.list)
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
      getEmailList () {
        if (this.dataForm.email === '') {
          return false
        }
        // 打开Loading 加载
        this.dataListLoading = true
        // 列表数据请求头
        var params = {
          // 用户名
          cmcEmail: this.dataForm.email
        }
        // 发起请求
        API.userManagement.getEmailInfo(params).then(({data}) => {
          // 请求成功
          if (data && data.code === 0) {
            // 赋值操作
            this.dataList = this.objectChangeArray(data.yhgl)
          } else {
            // 请求失败，数据为空
            this.dataList = []
          }
          // 隐藏等待
          this.dataListLoading = false
        })
      },
      copy (e) {
        // 双击将邮箱赋值到搜索框中
        this.dataForm.email = e
      },
      // 格式化
      objectChangeArray (obj) {
        let arr = []
        // 格式化类型
        if (obj.cmcVip === 0) {
          obj.cmcVip = '普通用户'
        } else if (obj.cmcVip === 1) {
          obj.cmcVip = 'vip1'
        } else if (obj.cmcVip === 2) {
          obj.cmcVip = 'vip2'
        } else if (obj.cmcVip === 3) {
          obj.cmcVip = 'vip3'
        }
        arr.push(obj)
        return arr
      },
      // 格式化
      formatting (array) {
        for (let i = 0; i < array.length; i++) {
          // 格式化类型
          if (array[i].cmcVip === 0) {
            array[i].cmcVip = '普通用户'
          } else if (array[i].cmcVip === 1) {
            array[i].cmcVip = 'vip1'
          } else if (array[i].cmcVip === 2) {
            array[i].cmcVip = 'vip2'
          } else if (array[i].cmcVip === 3) {
            array[i].cmcVip = 'vip3'
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
      // 修改身份
      addOrUpdateHandle (id) {
        this.addOrUpdateVisible = true
        this.$nextTick(() => {
          this.$refs.addOrUpdate.init(id)
        })
      },
      // 查看资产
      propertyHandle (id) {
        this.propertyVisible = true
        this.$nextTick(() => {
          this.$refs.property.init(id)
        })
      },
      // 查看钱包地址
      walletAddressHandle (id) {
        this.walletAddressVisible = true
        this.$nextTick(() => {
          this.$refs.walletAddress.init(id)
        })
      },
      // 查看通讯录
      addressBookHandle (id) {
        this.addressBookVisible = true
        this.$nextTick(() => {
          this.$refs.addressBook.init(id)
        })
      }
    }
  }
</script>
