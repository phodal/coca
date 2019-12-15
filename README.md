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
  ga          git analysis
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

Examples Result:

```
{
	"dataClass": [
		{
			"File": "examples/api/BookController.java",
			"BS": "dataClass"
		}
        ...
	],
	"lazyElement": [
		{
			"File": "examples/api/model/BookRepresentaion.java",
			"BS": "lazyElement"
		}
        ...
	]
}
```

### Code Line Count

```
coca scc
```

Results:

```
───────────────────────────────────────────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
───────────────────────────────────────────────────────────────────────────────
Go                          58     31763     7132       890    23741       2847
Java                        44       971      208        21      742         62
Markdown                     8       238       75         0      163          0
Gherkin Specificati…         2        32        2        16       14          0
Document Type Defin…         1       293       36         0      257          0
License                      1       201       32         0      169          0
SQL                          1         2        0         0        2          0
SVG                          1       199        0        34      165          0
Shell                        1         3        1         1        1          0
XML                          1        13        0         0       13          0
gitignore                    1        61        8         4       49          0
───────────────────────────────────────────────────────────────────────────────
Total                      119     33776     7494       966    25316       2909
───────────────────────────────────────────────────────────────────────────────
Estimated Cost to Develop $803,822
Estimated Schedule Effort 14.120551 months
Estimated People Required 6.743156
───────────────────────────────────────────────────────────────────────────────s
```

### Build Deps Tree

```
coca call -c com.phodal.pholedge.book.BookController.createBook -d deps.json -r com.phodal.pholedge.
```

Examples Results:

![Call Demo](docs/sample/call_demo.svg)

### Identify Spring API

```
coca api -p examples/api -d deps.json
```

### Git Analysis

```
coca ga -t -b 
```

```
+----------------------------------------------+-----------+-------------+
|                  ENTITYNAME                  | REVSCOUNT | AUTHORCOUNT |
+----------------------------------------------+-----------+-------------+
| adapter/call/JavaCallListener.go             |        35 |           2 |
| helloworld.go                                |        22 |           1 |
| refactor/base/JavaRefactorListener.go        |        16 |           2 |
| .gitignore                                   |        14 |           2 |
| refactor/rename/rename_method.go             |        12 |           2 |
| bs/BadSmellApp.go                            |        11 |           1 |
| cmd/analysis.go                              |        10 |           2 |
| README.md                                    |        10 |           2 |
| adapter/identifier/JavaIdentifierListener.go |         9 |           2 |
| bs/BadSmellListener.go                       |         8 |           1 |
| adapter/api/JavaApiListener.go               |         8 |           1 |
| cmd/refactor.go                              |         8 |           2 |
| adapter/identifier/JavaIdentifierApp.go      |         8 |           2 |
| refactor/main.go                             |         7 |           2 |
| src/domain/call_graph.go                     |         7 |           1 |
| refactor/base/JavaRefactorApp.go             |         7 |           1 |
| go.mod                                       |         7 |           1 |
| src/domain/concept_analyser.go               |         6 |           1 |
| cmd/root.go                                  |         6 |           2 |
| refactor/move_class_app.go                   |         6 |           1 |
+----------------------------------------------+-----------+-------------+
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
