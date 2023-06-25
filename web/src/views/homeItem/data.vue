<template>
  <div>
    <el-card v-if="nowInfo.type==='hash'">
      <el-table
          :data="data"
          style="width: 100%;max-height: 600px;overflow-y: auto">
        <el-table-column
            type="index"
            width="50">
        </el-table-column>
        <el-table-column
            prop="key"
            label="Key"
        >
        </el-table-column>
        <el-table-column
            prop="value"
            label="值"
        >
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
      data: [],
      cursor: 0,
      count: 0,
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
        this.data = res.data.data
        this.cursor = res.data.cursor
        this.count = res.data.count
      }
    }
  }
}
</script>

<style scoped>

</style>
