import request from "@/api/request";

export const reqRedisInfo = () => request({
    url: `/info/get_info`,
    method: "get",
})