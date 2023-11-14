//配置路由的地方
import Vue from "vue"

import VueRouter from "vue-router"
import routes from "@/router/routes";
import {GetToken} from "@/utils/token";
//使用插件
Vue.use(VueRouter);

//备份push方法
let originPush = VueRouter.prototype.push;
let originReplace = VueRouter.prototype.replace;

//重写push方法
VueRouter.prototype.push = function (location, resolve, reject) {
    if (resolve && reject) {
        originPush.call(this, location, resolve, reject);
    } else {
        originPush.call(this, location, () => {
        });
    }
}

VueRouter.prototype.replace = function (location, resolve, reject) {
    if (resolve && reject) {
        originReplace.call(this, location, resolve, reject);
    } else {
        originReplace.call(this, location, () => {
        }, () => {
        })
    }
}

let router = new VueRouter({
    //配置路由
    routes,
    scrollBehavior(to, from, savedPosition) {
        // return 期望滚动到哪个的位置
        return {y: 0}
    }
})

//全局守卫，前置守卫
router.beforeEach(async function (to, from, next) {
    let token = GetToken();
    if (token) {//已经登录
        if (to.path === "/login") {
            next("/?haveLogin=true");
        }
        next();
    } else {
        if (to.path !== "/login") {
            next("/login");
            return
        }
        next()
    }
})

//配置路由
export default router;
