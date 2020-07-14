$ go run gqlgen.go init
# github.com/99designs/gqlgen/internal/code
../../../99designs/gqlgen/internal/code/packages.go:11:12: undefined: packages.NeedName
../../../99designs/gqlgen/internal/code/packages.go:12:2: undefined: packages.NeedFiles
../../../99designs/gqlgen/internal/code/packages.go:13:2: undefined: packages.NeedImports
../../../99designs/gqlgen/internal/code/packages.go:14:2: undefined: packages.NeedTypes
../../../99designs/gqlgen/internal/code/packages.go:15:2: undefined: packages.NeedSyntax
../../../99designs/gqlgen/internal/code/packages.go:16:2: undefined: packages.NeedTypesInfo
../../../99designs/gqlgen/internal/code/packages.go:121:53: undefined: packages.NeedName
# github.com/99designs/gqlgen/internal/rewrite
../../../99designs/gqlgen/internal/rewrite/rewriter.go:24:9: undefined: packages.NeedSyntax
../../../99designs/gqlgen/internal/rewrite/rewriter.go:24:31: undefined: packages.NeedTypes


$go mod init
go: creating new go.mod: module github.com/xiazemin/graphql/exp10
export GOPROXY=https://goproxy.cn
$go run gqlgen.go init
Exec "go run ./server.go" to start GraphQL server

$go run server.go

mutation {
  updateArticle(articleId:2, content:"更新内容") {
    ok
    article {
      title
      content
    }
  }
}# Write your query or mutation here


curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br'
-H 'Content-Type: application/json' -H 'Accept: application/json'
 -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080'
 --data-binary '{"query":"mutation {\n  updateArticle(articleId:2, content:\"更新内容\") {\n    ok\n    article {\n      title\n      content\n    }\n  }\n}# Write your query or mutation here\n"}' --compressed

 $go run gqlgen.go
failed to load schema: graph/schema.graphqls:34: cannot use Article! as argument
input because OBJECT is not a valid input typeexit status 1

输入必须定义输入类型

https://github.com/graphql-go/graphql
