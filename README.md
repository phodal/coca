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
  bs          Bad Code Smell
  call        call graph api
  concept     concept api
  help        Help about any command
  refactor    auto refactor code
  scc         scc [FILE or DIRECTORY]


```

### Analysis

```
coca analysis -p [PATH]
```

### Find Bad Smells

```
coca bs -p examples/api -s type
```

### Code Line Count

```
coca cloc
```

### Build Deps Tree

```
coca call -c com.phodal.pholedge.book.BookController.createBook -d deps.json -r com.phodal.pholedge.
```

![Call Demo](docs/sample/call_demo.svg)

### Identify Spring API

```
coca api -p examples/api -d deps.json
```

### Refactor

support: 

 - rename
 - move
 - remove unused import
 - remove unused class

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
