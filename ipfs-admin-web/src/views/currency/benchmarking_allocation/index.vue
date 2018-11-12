<template>
  <!-- 系统管理 / 管理员列表 -->
  <div class="mod-benchmarking_allocation">
     <el-form :inline="true" @keyup.enter.native="getDataList()">      
      <el-form-item>
        <!-- 搜索 关键词 -->
        <el-select v-model="region" placeholder="对标代币" @change="editDriftStatus(region)">
          <el-option v-for="(item, index) in currencyList" :key="index" :label="item.cmcNumber" :value="item.cmcAutoid"></el-option>
        </el-select>
      </el-form-item>
     </el-form>
    <!-- Loading 加载 v-loading="xxx"-->    
    <el-transfer
      class="aaaaaaa"
      v-model="value1" 
      :props="{
        key: 'value',
        label: 'desc'
      }"
      :titles="['未分配币种列表', '已分配币种列表']"
      :data="data"
      @change="changeValue"
      ></el-transfer>    
  </div>
</template>

<script>
  import API from '@/api'
  export default {
    data () {
      return {
        // 显示对标代币名称
        region: '',
        // 对标代表 cmcAutoid
        cmcAutoid: '',
        // 等待Loading加载
        dataListLoading: false,
        // 未分配比重列表
        value1: [],
        // 对标代币列表
        currencyList: [],
        data: [],
        // 选中 数据列表
        dataListSelections: []
      }
    },
    created () {
      this.init()
    },
    methods: {
      // 初始化
      init () {
        this.getInfo()
      },
      // 获取对标代表列表
      getInfo () {
        // 列表数据请求头
        var params = {}
        // 发起请求
        API.benchmarkingAllocation.list(params).then(({data}) => {
          // 请求成功
          if (data && data.code === 0) {
            // 赋值
            this.currencyList = data.CurrencyREEntityList
          }
        })
      },
      // 选中代币名称请求表数据 cmcAutoid
      editDriftStatus (region) {
        this.cmcAutoid = region
        this.getDataList()
      },
      // 获取数据列表
      getDataList () {
        // 打开Loading 加载
        this.dataListLoading = true
        // 列表数据请求头
        var params = {
          // 代币id
          cmcAutoid: this.cmcAutoid
        }
        // 发起请求
        API.benchmarkingAllocation.info(params).then(({data}) => {
          // 请求成功
          if (data && data.code === 0) {
            this.data = this.formatting(data.distributionCurrencyList)
            this.value1 = this.formatting2(data.distributionCurrencyList)
            // 赋值操作
          } else {
            // 请求失败，数据为空
            console.log('请求失败')
            return false
          }
          // 隐藏等待
          this.dataListLoading = false
        })
      },
      changeValue () {
        let transCurrencyIdList = this.value1
         // 打开Loading 加载
        this.dataListLoading = true
        // 列表数据请求头
        var params = {
          // 代币id
          currencyAutoId: this.cmcAutoid,
          // transCurrencyIdList
          transCurrencyIdList: transCurrencyIdList
        }
        // 发起请求
        API.benchmarkingAllocation.changeInfo(params).then(({data}) => {
          // 请求成功
          if (data && data.code === 0) {
            // 赋值操作
          } else {
            // 请求失败，数据为空
            console.log('请求失败')
            return false
          }
          // 隐藏等待
          this.dataListLoading = false
        })
      },
      formatting (array) {
        let arr = []
        for (let i = 0; i < array.length; i++) {
          arr.push({
            value: array[i].cmcAutoid,
            desc: array[i].cmcName + ' ' + array[i].cmcNumber
          })
        }
        return arr
      },
      formatting2 (array) {
        let arr = []
        for (let i = 0; i < array.length; i++) {
          if (array[i].type === 1) {
            arr.push(array[i].cmcAutoid)
          }
        }
        return arr
      },
      formatting3 (array) {
        let arr = []
        for (let i = 0; i < array.length; i++) {
          if (array[i].type === 0) {
            arr.push(array[i].cmcAutoid)
          }
        }
        return arr
      }
    }
  }
</script>

<style>
  .aaaaaaa .el-transfer-panel__list {
   height: 450px;
}
  .aaaaaaa .el-transfer-panel__body {
   height: 450px;
 }
</style>
