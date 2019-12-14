# Coca

## Usage

install 

```
go get https://github.com/phodal/coca
```


help:

```
Usage:
  coca [command]

Available Commands:
  analysis    analysis package
  api         scan api
  call        call graph api
  concept     concept api
  help        Help about any command
  refactor    auto refactor code

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

License
---

[![Phodal's Idea](http://brand.phodal.com/shields/idea-small.svg)](http://ideas.phodal.com/)

@ 2019 A [Phodal Huang](https://www.phodal.com)'s [Idea](http://github.com/phodal/ideas).  This code is distributed under the MIT license. See `LICENSE` in this directory.
