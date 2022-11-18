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
        path: "/login",
        component: () => import("@/views/login"),
    }
]
