#!/bin/bash
###############################################################
# 说明：编译打包后端项目
###############################################################
# Harbor 或者其它镜像仓库地址
repository="harbor.ezops.cn"

# 生成版本信息文件
git rev-parse --short HEAD > config/version

# 获取版本信息变量
version=$(cat config/version)

# 镜像信息
image="${repository}/pandora:${version}"

# 2.镜像构建
docker build -t ${image} .

# 3.上传镜像
# docker push ${image}