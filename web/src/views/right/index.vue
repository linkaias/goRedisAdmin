<template>
  <div>
    <el-page-header>
      <template v-slot:title>
        {{ $t('right.home') }}
      </template>
      <template v-slot:content>
        <el-row>
          <span>{{ $t('right.description') }}</span>
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
<style>
.el-page-header i {
  display: none;
}
</style>
