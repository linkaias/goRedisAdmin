<script>
export default {
  name: "indexItem",
  data() {
    return {
      isLoading: false,
      InfoData: []
    }
  },
  mounted() {
    this.getInfo()
  },
  methods: {
    async getInfo() {
      this.isLoading = true
      let res = await this.$API.InfoApi.reqRedisInfo();
      console.log(res)
      this.InfoData = res.data
      this.isLoading = false
    }

  }
}
</script>

<template>
  <div style="display: flex;justify-content: center;margin-top: 20px">
    <el-card style="width: 50%;max-width: 1000px">
      <template #header>
        <div class="card-header">
          <span style="font-weight: bold">Redis Info</span>
        </div>
      </template>
      <div v-loading="isLoading" v-html="InfoData">
      </div>
    </el-card>
  </div>
</template>

<style scoped>

</style>