# Coca - toolbox for system refactoring and analysis

[![Build Status](https://travis-ci.org/phodal/coca.svg?branch=master)](https://travis-ci.org/phodal/coca)
[![Maintainability](https://api.codeclimate.com/v1/badges/d5a5e060522403b1f79b/maintainability)](https://codeclimate.com/github/phodal/coca/maintainability)
[![codecov](https://codecov.io/gh/phodal/coca/branch/master/graph/badge.svg)](https://codecov.io/gh/phodal/coca)

> Coca is a toolbox which is design for legacy system refactoring. Coca 是一个用于系统重构、系统迁移和系统分析的瑞士军刀。它可以分析代码中的 badsmell，行数统计，分析调用与依赖，进行 Git 分析，以及自动化重构等。

Refactoring Modeling:

![Refactoring Modeling](docs/images/model.svg)

Todo:

 - [x] Show all todos and history
 - [ ] Suggest API for Design Patterns
   - [x] factory pattern
   - [ ] strategy 
   - [ ] builder
 - [ ] Docs
    - [ ] Lifecycle for new projects: evaluate (cloc, bad smell, api, git, todo) -> design -> patterns (suggest) -> refactoring ()
 -  Evaluate API
    - [x] Average Method Length
    - [x] Average Class Method Count
 - Date Collections
    - [ ] monolithic
    - [ ] microservice 
    - [ ] big data

## Usage

install 

```bash
go get -u github.com/phodal/coca
```

help:

```bash
Usage:
  coca [command]


Available Commands:
  analysis    analysis package
  api         scan api
  bs          bad smell analysis
  call        call graph api
  cloc        cloc [FILE or DIRECTORY]
  concept     concept api
  count       count code
  evaluate    evaluate refactor effort
  git         git analysis
  help        Help about any command
  rcall       reverse call
  refactor    auto refactor code
  suggest     simple holmes
  todo        scan todo
  version     print the version number of Coca


```

### Analysis

```
coca analysis
```

### Find Bad Smells

```bash
coca bs -s type
```

Examples Result:

```json
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
coca cloc
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
───────────────────────────────────────────────────────────────────────────────
```

Results to json

```
coca cloc --by-file --format json
```

### Build Deps Tree

```
coca call -c com.phodal.pholedge.book.BookController.createBook -r com.phodal.pholedge.
```

Examples Results:

![Call Demo](docs/sample/call_demo.svg)

### Identify Spring API

```
coca api -f
```

![API Demo](docs/sample/api.svg)

With Count

```
coca api -r com.phodal.pholedge. -c 
``` 

or multi package:

`coca api  -r com.macro.mall.demo.controller.,com.zheng.cms.admin.,com.phodal.pholedge -c` 

```bash
+------+--------+------------------------------------------------+------------------------------------------------------------------------+
| SIZE | METHOD |                      URI                       |                                 CALLER                                 |
+------+--------+------------------------------------------------+------------------------------------------------------------------------+
|   36 | GET    | /aliyun/oss/policy                             | controller.OssController.policy                                        |
|   21 | POST   | /aliyun/osscallback                            | controller.OssController.callback                                      |
|   17 | GET    | /subject/list                                  | controller.CmsSubjectController.getList                                |
|   17 | GET    | /esProduct/search                              | search.controller.EsProductController.search                           |
|   17 | GET    | /order/list                                    | controller.OmsOrderController.list                                     |
|   17 | GET    | /productAttribute/list/{cid}                   | controller.PmsProductAttributeController.getList                       |
|   17 | GET    | /productCategory/list/{parentId}               | controller.PmsProductCategoryController.getList                        |
|   17 | GET    | /brand/list                                    | controller.PmsBrandController.getList                                  |
|   17 | GET    | /esProduct/search/simple                       | search.controller.EsProductController.search                           |
+------+--------+------------------------------------------------+------------------------------------------------------------------------+
```

### Git Analysis

```
coca git -t
```

Results: 

```bash
+---------------------------------------------------------------------------------------------------------------------+-----------+-------------+
|                                                     ENTITYNAME                                                      | REVSCOUNT | AUTHORCOUNT |
+---------------------------------------------------------------------------------------------------------------------+-----------+-------------+
| build.gradle                                                                                                        |      1326 |          36 |
| src/asciidoc/index.adoc                                                                                             |       239 |          20 |
| build-spring-framework/resources/changelog.txt                                                                      |       187 |          10 |
| spring-core/src/main/java/org/springframework/core/annotation/AnnotationUtils.java                                  |       170 |          10 |
| spring-beans/src/main/java/org/springframework/beans/factory/support/DefaultListableBeanFactory.java                |       159 |          15 |
| src/docs/asciidoc/web/webmvc.adoc                                                                                   |       121 |          24 |
| spring-context/src/main/java/org/springframework/context/annotation/ConfigurationClassParser.java                   |       118 |           9 |
| src/dist/changelog.txt                                                                                              |       118 |           9 |
| spring-webmvc/src/main/java/org/springframework/web/servlet/config/annotation/WebMvcConfigurationSupport.java       |       116 |          15 |
| spring-beans/src/main/java/org/springframework/beans/factory/support/AbstractAutowireCapableBeanFactory.java        |       113 |          15 |
| spring-web/src/main/java/org/springframework/http/HttpHeaders.java                                                  |       111 |          18 |
| src/docs/asciidoc/web/webflux.adoc                                                                                  |       108 |          21 |
| spring-core/src/main/java/org/springframework/core/annotation/AnnotatedElementUtils.java                            |       107 |           9 |
| spring-test/spring-test.gradle                                                                                      |       105 |           7 |
| spring-webmvc/src/main/java/org/springframework/web/servlet/mvc/method/annotation/RequestMappingHandlerAdapter.java |       105 |          13 |
| spring-messaging/src/main/java/org/springframework/messaging/simp/stomp/StompBrokerRelayMessageHandler.java         |       101 |          12 |
| spring-web/src/main/java/org/springframework/web/client/RestTemplate.java                                           |        98 |          17 |
| spring-webmvc/src/main/java/org/springframework/web/servlet/resource/ResourceHttpRequestHandler.java                |        96 |          14 |
| org.springframework.core/src/main/java/org/springframework/core/convert/TypeDescriptor.java                         |        93 |           4 |
| spring-core/src/main/java/org/springframework/core/ResolvableType.java                                              |        92 |          10 |
+---------------------------------------------------------------------------------------------------------------------+-----------+-------------+
```

### Concept Analyser

```
coca concept
```

Results Examples:

```
+------------------+--------+
|      WORDS       | COUNTS |
+------------------+--------+
| context          |    590 |
| resolve          |    531 |
| path             |    501 |
| content          |    423 |
| code             |    416 |
| resource         |    373 |
| property         |    372 |
| session          |    364 |
| attribute        |    349 |
| properties       |    343 |
| headers          |    330 |
+------------------+--------+
```

### Count Refs

```
coca count
```

Results:

```
+------------+--------------------------------------------------------------------------+
| REFS COUNT |                                  METHOD                                  |
+------------+--------------------------------------------------------------------------+
|          2 | com.phodal.pholedge.book.BookRepository.byId                             |
|          2 | com.phodal.pholedge.book.model.Book.toRepresentation                     |
|          2 | com.phodal.pholedge.book.BookRepository.save                             |
|          2 | com.phodal.coca.analysis.JavaCallApp.parse                               |
|          2 | com.phodal.pholedge.book.BookRepository.save                             |
|          2 | com.phodal.coca.analysis.JavaCallApp.parse                               |
|          1 | com.phodal.pholedge.book.model.Book.save                                 |
|          1 | evolution.analysis.jv.calls.JavaCallVisitor.parseNewType                 |
|          1 | evolution.analysis.jv.calls.JavaCallVisitor.isNotSpecialNewWord          |
|          1 | com.phodal.pholedge.book.BookMapper.byId                                 |
|          1 | com.phodal.pholedge.book.BookService.updateBook                          |
|          1 | com.phodal.pholedge.book.BookService.getBooksLists                       |
|          1 | com.phodal.pholedge.book.BookService.getBookById                         |
|          1 | com.phodal.pholedge.book.BookMapper.doSave                               |
|          1 | com.phodal.pholedge.book.BookMapper.list                                 |
|          1 | com.phodal.pholedge.book.BookService.createBook                          |
|          1 | com.phodal.pholedge.book.BookFactory.create                              |
|          1 | com.phodal.pholedge.book.BookRepository.list                             |
|          1 | com.phodal.pholedge.book.model.Book.create                               |
+------------+--------------------------------------------------------------------------+
```

### Reverse Call Graph

```
coca rcall -c org.bytedeco.javacpp.tools.TokenIndexer.get
```

Results:

```
digraph G { 
edge [dir="back"];

"org.bytedeco.javacpp.tools.Parser.extern" -> "org.bytedeco.javacpp.tools.Parser.declarations";
"org.bytedeco.javacpp.tools.Parser.declarations" -> "org.bytedeco.javacpp.tools.Parser.extern";
...
}
```

![RCall Demo](docs/sample/rcall.svg)

### Auto Refactor

support: 

 - rename
 - move
 - remove unused import
 - remove unused class

```
coca refactor -R rename.coca -p src/main
coca refactor -m move.config -p .
```

### Evaluate

```
coca evaluate
```

```
+--------------------------------+-------+-----------------------+-------+-----------+
|              TYPE              | COUNT |         LEVEL         | TOTAL |   RATE    |
+--------------------------------+-------+-----------------------+-------+-----------+
| Nullable / Return Null         |     0 | Method                |  5256 | 0.00%     |
| Utils                          |    18 | Class                 |   366 | 4.92%     |
| Static Method                  |     0 | Method                |  5256 | 0.34%     |
| Average Method Num.            |  5256 | Method/Class          |   366 | 14.360656 |
| Method Num. Std Dev / 标准差   |  5256 | Class                 | -     | 32.965336 |
| Average Method Length          | 19644 | Without Getter/Setter |  4328 |  4.538817 |
| Method Length Std Dev / 标准差 |  5256 | Method                | -     |  4.467935 |
+--------------------------------+-------+-----------------------+-------+-----------+
```

Evaluate.json examples

```json
{
	"Nullable": {
		"Items": [
			"nonnull.Name.testNull",
			"nonnull.Name.orElseNull",
			"org.ofbiz.base.util.UtilURL.fromResource",
			"org.ofbiz.base.util.UtilURL.getOfbizHomeRelativeLocationFromFilePath",
			"study.huhao.demo.adapters.outbound.persistence.blog.BlogPO.toDomainModel",
			"study.huhao.demo.adapters.outbound.persistence.blog.BlogPO.toDomainModel",
			"study.huhao.demo.adapters.outbound.persistence.blog.BlogPO.of",
			"study.huhao.demo.infrastructure.persistence.blog.BlogPO.convertDomain"
		]
	},
	"ServiceSummary": {
		"LifecycleMap": null,
		"ReturnTypeMap": {
			"BookRepresentaion": [
				"com.phodal.pholedge.book.BookService.getBookById",
				"com.phodal.pholedge.book.BookService.updateBook"
			]
		},
		"RelatedMethod": null
	},
	"UtilsSummary": {},
	"Summary": {
		"UtilsCount": 1,
		"ClassCount": 64,
		"MethodCount": 161,
		"StaticMethodCount": 19
	}
}
```

### Todo

```
coca todo
```

results:

```
+------------+--------------+--------------------------------+---------------------------------------------------+-------+
|    DATE    |    AUTHOR    |            MESSAGES            |                     FILENAME                      | LINE  |
+------------+--------------+--------------------------------+---------------------------------------------------+-------+
| 2019-12-19 | Phodal Huang | 支持 interface 在同一个包内    | core/adapter/api/JavaApiListener.go               |   200 |
| 2019-12-21 | Phodal Huang | 处理链试调用                   | core/adapter/bs/BadSmellListener.go               |   305 |
| 2019-12-18 | Phodal Huang | update this reflect            | core/adapter/bs/BadSmellListener.go               |   363 |
| 2019-12-15 | Phodal Huang | update for array               | core/adapter/bs/BadSmellListener.go               |   388 |
| 2019-12-24 | Phodal Huang | 支持依赖注入                   | core/adapter/call/JavaCallListener.go             |   108 |
| 2019-12-24 | Phodal Huang | add inner creator examples     | core/adapter/call/JavaCallListener.go             |   209 |
| 2019-12-24 | Phodal Huang | add inner creator examples     | core/adapter/call/JavaCallListener.go             |   215 |
| 2019-12-24 | Phodal Huang |                                | core/adapter/call/JavaCallListener.go             |   270 |
| 2019-12-20 | Phodal Huang | 处理链试调用                   | core/adapter/call/JavaCallListener.go             |   324 |
+------------+--------------+--------------------------------+---------------------------------------------------+-------+
```

### Suggest

```
coca suggest
```

results:

```
+--------+------------------+--------------------------------+
| CLASS  |     PATTERN      |             REASON             |
+--------+------------------+--------------------------------+
| Insect | factory          | too many constructor           |
| Bee    | factory, builder | complex constructor, too       |
|        |                  | many constructor, too many     |
|        |                  | parameters                     |
+--------+------------------+--------------------------------+
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
git clone https://github.com/phodal/coca
```

Test Frameworks

```
go get github.com/onsi/ginkgo
go get github.com/onsi/gomega
```

License
---

[![Phodal's Idea](http://brand.phodal.com/shields/idea-small.svg)](http://ideas.phodal.com/)

@ 2019 A [Phodal Huang](https://www.phodal.com)'s [Idea](http://github.com/phodal/ideas).  This code is distributed under the MPL license. See `LICENSE` in this directory.
