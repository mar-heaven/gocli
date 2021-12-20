### go编译

go build 在项目根目录下执行的时候（就是和go mod同级),如果发现哪个文件里有 **main** 函数，就会生成项目
名的可执行文件，本项目 用`go mod init github.com/mar-heaven/gocli` 初始化，项目名就是 **gocli**
上传到 **github** 后再 `go get github.com/mar-heaven/gocl`, 就可以使用 **gocli** 命令了。

## 补充
我的目录结构如下
```
├── .gitignore
|  ├── .gitignore
|  ├── gocli-learn.iml
|  ├── modules.xml
|  ├── vcs.xml
|  └── workspace.xml
├── README.md
├── cmd
|  ├── gocli
|  └── helper
├── go.mod
└── go.sum
```
如果想装 cmd下面的 **helper** mod，用下边的命令就好：
```
go get -u github.com/mar-heaven/gocli/cmd/helper
```