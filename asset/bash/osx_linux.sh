#!/usr/bin/env sh

#https://gist.github.com/iamucil/7578dc7df7d72e1d78c8f5543db3fbcc
# 设置默认值
CUR="gin-api"
NEW="${NEW:-}"

# 检查 NEW 是否为空
if [ -z "$NEW" ]; then
    echo "Please enter the new module name (e.g., github.com/yourname/yourrepo):"
    read NEW
    if [ -z "$NEW" ]; then
        echo "No module name provided. Exiting script."
        exit 1
    fi
fi

# 修改 go mod 模块名
go mod edit -module "$NEW"

# 查找并替换所有 .go 文件中的模块导入
find . -type f -name '*.go' -exec perl -pi -e "s/$CUR/$NEW/g" {} \;