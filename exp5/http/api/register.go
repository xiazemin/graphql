package api
import "github.com/graphql-go/handler"

func Register() *handler.Handler {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	return h
}
