# 如何使用这个脚手架

1.  clone 这个项目

```bash
git clone https://github.com/soolaugust/go-toolkit.git
```

2. 引入脚手架

    2.1 如果你的项目还没建立，可以直接修改根文件夹名为你的项目名称，然后使用下面的命令

    ```bash
    go mod init {你的项目名}
    go mod tidy
    ```

    2.2 如果你已经有现有的项目，直接将所有的文件夹复制到你项目的文件夹下，然后使用下面的命令

    ```bash
    go mod tidy
    ```

## 目前依赖的库

**log**

log采用 [uber-go/zap](https://github.com/uber-go/zap)

具体原因可参考： [在Go语言项目中使用Zap日志库](https://www.liwenzhou.com/posts/Go/zap/)

**配置读取**

配置采用[spf13/viper](https://github.com/spf13/viper)