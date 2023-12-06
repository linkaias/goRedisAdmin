//获取数据库列表
import request from "@/api/request";

export const reqGetDbList = () => request({
    url: `/db/db_list`,
    method: "get",
})

export const reqGetKeys = (dbNum, filter) => request({
    url: `/db/get_keys?db_num=${dbNum}&filter=${filter}`,
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


export const reqLogin = (data) => request({
    url: `/user/login`,
    method: "post",
    data
})


export const reqGetValueByKey = (dbNum, data) => request({
    url: `/db/get_val_by_key?db_num=${dbNum}`,
    method: "post",
    data
})

export const reqExpireKey = (dbNum,expire,key) => request({
    url: `/db/key/expire?db_num=${dbNum}&expire=${expire}&key=${key}`,
    method: "post",
})
