<template>
  <div>
    <el-container>
      <el-header>
        <HeaderItem></HeaderItem>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>
<script>
import HeaderItem from "@/views/header"

export default {
  name: 'App',
  components: {
    HeaderItem
  }
}
</script>
<style>
body {
  margin: 0;
}
</style>
