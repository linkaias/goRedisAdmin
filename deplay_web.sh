#!/bin/bash

# 目标路径
target_path="/1panel/1panel/apps/openresty/uiucode-nginx/www/sites/gradmin.uiucode.com/index/web/"
# 要复制的文件夹路径
source_folder="./html/"

# 复制文件夹并覆盖目标文件夹
cp -r -f "$source_folder" "$target_path"

echo "Success!"