//配置路由的地方
import Vue from "vue"

import VueRouter from "vue-router"
import routes from "@/router/routes";
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
/*
router.beforeEach(async function (to, from, next) {

    let token = store.state.user.token;
    if (token) {//已经登录
        if (to.path === "/login" || to.path === "/register") {
            next("/");
        } else {
            let name = store.state.user.userInfo.name;
            if (name) { //已经有登录数据
                next();
            } else {//重新获取登录数据
                try {
                    await store.dispatch("GetUserInfo")
                    next();
                } catch (e) {
                    //token失效，退出登录
                    await store.dispatch("Logout")
                    next("/login");
                }
            }
        }
    } else {
        let path = to.path;
        if (path.indexOf("/trade") !== -1 || path.indexOf("/pay") !== -1 || path.indexOf("/center") !== -1) {
            next("/login?redirect=" + path);
        }
        next();
    }
    next();
})*/

//配置路由
export default router;
