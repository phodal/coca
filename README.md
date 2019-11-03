# Coca

## Usage

install 

```
go get https://github.com/phodal/coca
```

Analysis

```
coca analysis -p [PATH]
```

Refactor

```
coca refactor -R rename.coca -D deps.json -p src/main
coca refactor -m move.config -p .
```

## Dev

Install Go

```bash
brew install go
```

Env

```bash
export GOROOT=/usr/local/opt/go/libexec
export GOPATH=$HOME/.go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

clone

```
go get https://github.com/phodal/coca
```

Test Frameworks

```
go get github.com/onsi/ginkgo
go get github.com/onsi/gomega
```

### Refs

[https://github.com/MontFerret/ferret](https://github.com/MontFerret/ferret)
