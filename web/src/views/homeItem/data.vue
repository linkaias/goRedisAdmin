<template>
  <div>
    <el-card v-if="nowInfo.type==='hash'|| nowInfo.type==='list'">
      <el-table
          :data="[]"
          style="width: 100%">
        <el-table-column
            prop="date"
            label="日期"
            width="180">
        </el-table-column>
        <el-table-column
            prop="name"
            label="姓名"
            width="180">
        </el-table-column>
        <el-table-column
            prop="address"
            label="地址">
        </el-table-column>
      </el-table>
    </el-card>
    <el-card v-if="nowInfo.type==='string'">
      <p>Key: <span>{{ nowInfo.key }}</span></p>
      <el-divider></el-divider>
      <p>Value:<span>{{ result }}</span></p>
    </el-card>

    <div style="text-align: right;margin-top: 20px">
      <el-button @click="close" size="small" icon="el-icon-close">关闭</el-button>
    </div>

  </div>
</template>

<script>
export default {
  name: "data",
  data() {
    return {
      dbNum: -1,
      nowInfo: {},
      result: {}
    }
  },
  methods: {
    initData(dbNum, info) {
      this.dbNum = dbNum
      this.nowInfo = info
      this.getData()
    },
    close() {
      this.$emit("closeData")
    },
    async getData() {
      let data = {
        key: this.nowInfo.key,
        type: this.nowInfo.type
      }
      let res = await this.$API.dbApi.reqGetValueByKey(this.dbNum, data)
      if (res.code === 0) {
        this.result = res.data
      }
    }
  }
}
</script>

<style scoped>

</style>
