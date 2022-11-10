const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
    transpileDependencies: true,
    productionSourceMap: false,
    lintOnSave: false,
    devServer: {
      proxy: {
        "/api/v1": {
          target: "http://127.0.0.1:9527",
        }
      }
    }
})
