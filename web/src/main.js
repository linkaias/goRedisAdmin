import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false
//引入路由
import router from "./router"

//API
import API from "@/api"

Vue.prototype.$API = API


import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import locale from 'element-ui/lib/locale/lang/zh-CN' // lang i18n

// set ElementUI lang to EN
Vue.use(ElementUI, {locale})


new Vue({
    render: h => h(App),
    router
}).$mount('#app')
