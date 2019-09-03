# 微信文章资源下载工具

## 命令行下载工具
提供了命令行形式的下载工具。

### 构建命令行下载工具
`go build -o downloader cmd/cli/downloader.go`

### 使用命令行下载工具
```
./downloader --folder=/Users/mj/Downloads --url=https://mp.weixin.qq.com/s/5B1zecPcJ-LX_YKZJpNJyg
2019/09/03 16:47:50 创建资源目录: /Users/mj/Downloads/人类史上首个太空 AI 机器人，IBM 和空客如何两年开发了它？
2019/09/03 16:47:50 保存视频地址...
2019/09/03 16:47:50 下载图片...
2019/09/03 16:47:52 资源下载成功
```

## GUI 下载工具
使用 Fyne 制作的图形化界面。

### 构建 GUI 下载工具
`go build -o downloader_gui cmd/gui/downloader.go`

### 打包（可选）
如果你是自己使用，打不打包其实不重要，如果需要打包，参考 Fyne [官方文档](https://fyne.io/develop/distribution.html)。

### 使用 GUI 下载工具
![使用 GUI 下载工具](gui_screenshot.png?raw=true "使用 GUI 下载工具")
