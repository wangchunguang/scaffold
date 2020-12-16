#!/usr/bin/env bash

protoDir="../proto"
outDir="../pb/gym"
protoc -I ${protoDir}/ ${protoDir}/*proto --go_out=plugins=grpc:${outDir}
# -i 指定import路径，可以指定多个参数 编译时按照顺序查找
# -go_out 指定og语言的访问类
# plugins 指定依赖的插件
# 使用gofast 将go_out改为gofast_out
protoc -I D:/go_project/src/business_master/proto D:/go_project/src/business_master/proto/*proto --gofast_out=plugins=grpc:D:/go_project/src/business_master/pb