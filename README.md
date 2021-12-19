### go编译

go build 在项目根目录下执行的时候（就是和go mod同级),如果发现哪个文件里有 **main** 函数，就会生成项目
名的可执行文件，本项目 用`go mod init github.com/mar-heaven/gocli` 初始化，项目名就是 **gocli**
上传到 **github** 后再 `go get github.com/mar-heaven/gocl`, 就可以使用 **gocli** 命令了。

[Compile_packages_and_dependencies](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)