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

export const reqExportKeyOld = (dbNum, key) => request({
    url: `/db/export_keys?db_num=${dbNum}&key=${key}`,
    method: "post",
})

export const reqExportKey = async function (dbNum, key) {
    try {
        let response = await request.post(
            `/db/export_keys?db_num=${dbNum}&key=${key}`,
            {responseType: 'blob'} // 设置响应类型为 blob
        );

        // 确保返回的是数组
        if (!Array.isArray(response.data)) {
            console.error("返回的数据不是一个数组！");
            return;
        }
// 将数组转换为 JSON 字符串
        const jsonString = JSON.stringify(response.data, null, 2); // 格式化 JSON 字符串

        // 创建一个 Blob 对象
        const blob = new Blob([jsonString], {type: 'application/json'});

        // 创建下载链接
        const link = document.createElement('a');
        link.href = URL.createObjectURL(blob);

        // 设置下载文件名
        link.download = 'keys.json'; // 设置下载文件名

        // 触发下载
        link.click();

        // 释放 URL 对象
        URL.revokeObjectURL(link.href);
    } catch (error) {
        console.error('下载文件失败：', error);
    }

}

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

export const reqExpireKey = (dbNum, expire, key) => request({
    url: `/db/key/expire?db_num=${dbNum}&expire=${expire}&key=${key}`,
    method: "post",
})
