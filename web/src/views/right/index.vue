<template>
  <div>
    <el-page-header >
      <template v-slot:title>
        首页
      </template>
      <template v-slot:content>
        <el-row>
          <span>页面描述</span>
        </el-row>
      </template>
    </el-page-header>
    <el-divider></el-divider>
    right
  </div>
</template>

<script>
export default {
  name: "rightContent"
}
</script>

<style >
.el-page-header i{
  display: none;
}
</style>
