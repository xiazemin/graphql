go get -u github.com/graphql-go/graphql
go get github.com/99designs/gqlgen
 # github.com/99designs/gqlgen/internal/code
goLang/src/github.com/99designs/gqlgen/internal/code/packages.go:11:12: undefined: packages.NeedName
goLang/src/github.com/99designs/gqlgen/internal/code/packages.go:12:2: undefined: packages.NeedFiles
goLang/src/github.com/99designs/gqlgen/internal/code/packages.go:13:2: undefined: packages.NeedImports
goLang/src/github.com/99designs/gqlgen/internal/code/packages.go:14:2: undefined: packages.NeedTypes
goLang/src/github.com/99designs/gqlgen/internal/code/packages.go:15:2: undefined: packages.NeedSyntax
goLang/src/github.com/99designs/gqlgen/internal/code/packages.go:16:2: undefined: packages.NeedTypesInfo
goLang/src/github.com/99designs/gqlgen/internal/code/packages.go:121:53: undefined: packages.NeedName
# github.com/99designs/gqlgen/internal/rewrite
goLang/src/github.com/99designs/gqlgen/internal/rewrite/rewriter.go:24:9: undefined: packages.NeedSyntax
goLang/src/github.com/99designs/gqlgen/internal/rewrite/rewriter.go:24:31: undefined: packages.NeedTypes


validation failed: packages.Load: generated.go:14:2: cannot find package "github.com/vektah/gqlparser/v2/ast" in any of:
        /usr/local/go/src/github.com/vektah/gqlparser/v2/ast (from $GOROOT)
        /Users/didi/goLang/src/github.com/vektah/gqlparser/v2/ast (from $GOPATH)
        /Users/didi/PhpstormProjects/go/src/github.com/vektah/gqlparser/v2/ast
/Users/didi/goLang/src/generated.go:14:2: could not import github.com/vektah/gqlparser/v2/ast (invalid package name: "")
/Users/didi/goLang/src/generated.go:113:20: undeclared name: gqlparser
/Users/didi/goLang/src/generated.go:21:9: cannot use &(executableSchema literal) (value of type *executableSchema) as graphql.ExecutableSchema value in return statement: wrong type for method Schema

go get github.com/vektah/gqlparser
go: finding github.com/vektah/gqlparser v1.3.1
go: downloading github.com/vektah/gqlparser v1.3.1
go: extracting github.com/vektah/gqlparser v1.3.1


	"github.com/vektah/gqlparser/v2/ast"改成
	"github.com/vektah/gqlparser/ast"

$go mod init
go: creating new go.mod: module github.com/xiazemin/graphql/exp7
$go mod edit -replace github.com/vektah/gqlparser/v2/ast=/Users/didi/goLang/src/github.com/vektah/gqlparser/ast

cd  ~/goLang/src/github.com/vektah

mkdir v2
cp -r gqlparser/ v2/
$mv v2 gqlparser/
编译成功

$go run gqlgen.go init
Exec "go run ./server.go" to start GraphQL server

$tree
.
|____build.sh
|____go.mod
|____go.sum
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


$go run ./server.go
2020/07/13 12:19:23 connect to http://localhost:8080/ for GraphQL playground

