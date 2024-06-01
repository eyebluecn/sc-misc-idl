# sc-misc-idl
smart-classroom-misc-idl 智慧课堂misc项目对应的idl接口定义 。

## 安装环境
如果首次运行，需要先安装thriftgo和kitex，参考[kitex官网](https://cloudwego.cn/zh/docs/kitex/getting-started/prerequisite/)
```shell
# 安装thriftgo 会讲thriftgo命令安装至GOPATH.
go install github.com/cloudwego/thriftgo@latest
# 验证是否安装好了thriftgo  eg:thriftgo 0.3.12
thriftgo --version

# 安装kitex 会将kitex命令安装至GOPATH.
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
# 验证是否安装好了kitex  eg:v0.9.1
kitex --version
```

## 运行脚本
```shell
kitex -module github.com/eyebluecn/sc-misc-idl idl/smart_classroom_misc.thrift
```