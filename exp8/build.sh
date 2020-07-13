$go run gqlgen.go generate ./schema.graphql
$tree
.
|____build.sh
|____generated.go
|____gqlgen.go
|____schema.graphql

go run github.com/99designs/gqlgen generate
tree
.
|____build.sh
|____generated.go
|____gqlgen.go
|____schema.graphql

By default gqlgen will use any models in the model directory that match on name, this can be configured in gqlgen.yml.


https://gqlgen.com/getting-started

$go run gqlgen.go init
Exec "go run ./server.go" to start GraphQL server

$tree
.
|____build.sh
|____generated.go
|____gqlgen.go
|____gqlgen.yml
|____graph
| |____generated
| | |____generated.go
| |____model
| | |____models_gen.go
| |____resolver.go
| |____schema.graphqls
| |____schema.resolvers.go
|____schema.graphql
|____server.go


 go run github.com/99designs/gqlgen generate

 At the top of our resolver.go, between package and import, add the following line:

//go:generate go run github.com/99designs/gqlgen
This magic comment tells go generate what command to run when we want to regenerate our code. To run go generate recursively over your entire project, use this command:

go generate ./...

$go run github.com/99designs/gqlgen --verbose

默认会生成一个schema文件graph/schema.graphqls
所以必须替换这个文件才能生成我们需要的代码

$mv graph/schema.graphqls graph/schema.graphqls.bak
$mv schema.graphql graph/schema.graphqls

$go run gqlgen.go init
validation failed: packages.Load: /Users/didi/goLang/src/github.com/xiazemin/graphql/exp8/graph/schema.resolvers.go:33:72: NewTodo not declared by package model
/Users/didi/goLang/src/github.com/xiazemin/graphql/exp8/graph/schema.resolvers.go:33:89: Todo not declared by package model
/Users/didi/goLang/src/github.com/xiazemin/graphql/exp8/graph/schema.resolvers.go:36:62: Todo not declared by package model
/Users/didi/goLang/src/github.com/xiazemin/graphql/exp8/graph/schema.resolvers.go:39:41: MutationResolver not declared by package generated

Exec "go run ./server.go" to start GraphQL server


$ rm graph/schema.resolvers.go

$go run gqlgen.go init
Exec "go run ./server.go" to start GraphQL server

$ go run server.go
2020/07/13 17:54:27 connect to http://localhost:8080/ for GraphQL playground

https://blog.csdn.net/liuyh73/article/details/85028977

