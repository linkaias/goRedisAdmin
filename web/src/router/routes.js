export default [
    {
        path: "*",
        redirect: "/home",
    },
    {
        path: "/home",
        component: () => import("@/views/homeItem"),
        meta: {show: true},
    }

]
