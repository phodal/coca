# learn-go

## 安装 Go

```bash
brew install go
```

配置环境 

```bash
export GOROOT=/usr/local/opt/go/libexec
export GOPATH=$HOME/.go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

### 安装 Go 依赖管理工具：godep

```
brew install dep
```

TBD.

## Test

```
go get github.com/onsi/ginkgo
go get github.com/onsi/gomega
```

## Refs

[https://github.com/MontFerret/ferret](https://github.com/MontFerret/ferret)