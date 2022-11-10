//获取数据库列表
import request from "@/api/request";

export const reqGetDbList = () => request({
    url: `/db/db_list`,
    method: "get",
})

export const reqGetKeys = (dbNum) => request({
    url: `/db/get_keys?db_num=${dbNum}`,
    method: "get",
})

export const reqDelKey = (dbNum, key) => request({
    url: `/db/key?db_num=${dbNum}&key=${key}`,
    method: "delete",
})

export const reqAddVal = (data) => request({
    url: `/db/key`,
    method: "post",
    data
})

export const reqFlush = (type, db) => request({
    url: `/db/flush?type=${type}&db_num=${db}`,
    method: "delete",
})
