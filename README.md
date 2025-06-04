# GoMobile iOS SDK Demo

这是一个使用 GoMobile 构建 iOS SDK 的示例项目。

## 项目结构

```
.
├── .github/
│   └── workflows/
│       └── build.yml    # GitHub Actions 构建脚本
├── go/
│   └── hello/
│       └── hello.go     # Go 源代码
├── build/              # 构建输出目录
├── release/           # 发布文件目录
├── go.mod            # Go 模块文件
└── README.md         # 项目说明文档
```

## 使用方法

1. Fork 或克隆此仓库
2. 修改 `go.mod` 中的模块路径为你的 GitHub 仓库路径
3. 根据需要修改 `go/hello/hello.go` 中的代码
4. 推送代码到 GitHub
5. GitHub Actions 将自动构建 iOS framework
6. 从 Actions 页面下载生成的 framework

## 本地开发

```bash
# 安装 gomobile
go install golang.org/x/mobile/cmd/gomobile@latest
go install golang.org/x/mobile/cmd/gobind@latest

# 初始化 gomobile
gomobile init

# 构建 iOS framework
cd go/hello
gomobile bind -target=ios -o ../../build/Hello.framework github.com/yourusername/gomobile_iosdemo/go/hello
```

## 注意事项

1. 确保 Go 代码中导出的函数使用 `//export` 注释
2. 导出的函数名必须以大写字母开头
3. 避免使用 Go 标准库中不支持的包 