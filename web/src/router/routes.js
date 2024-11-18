export default [
    {
        path: "*",
        redirect: "/home",
    },
    {
        path: "/home",
        component: () => import("@/views/homeItem"),
    },
    {
        path: "/info",
        component: () => import("@/views/infoItem/indexItem.vue"),
    },
    {
        path: "/login",
        component: () => import("@/views/login"),
    }
]
