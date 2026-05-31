import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false
import router from "./router"
import i18n, { setLanguage } from "./i18n"

//API
import API from "@/api"

Vue.prototype.$API = API
Vue.prototype.$setLanguage = setLanguage


import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)


new Vue({
    render: h => h(App),
    router,
    i18n
}).$mount('#app')
