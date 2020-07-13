package mian

import (
	"fmt"
	"github.com/xiazemin/graphql/exp5/http/api"
	"net/http"
)

func main() {
h := api.Register()
handler := cors.Default().Handler(h)
http.Handle("/graphql", handler)
fmt.Println("The api server will run on port : ", apiPort)
http.ListenAndServe(apiPort, nil)
}