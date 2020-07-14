package main

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"log"
)

func loger(key string) {
	fmt.Println(key)
}
func main()  {

	// create a client (safe to share across requests)
	client := graphql.NewClient("http://localhost:8080/query")
	client.Log=loger
	// make a request
	req := graphql.NewRequest(
		//{"query":"mutation  createTodo($input: NewTodo!){\n  createTodo(input:$input){\n    done\n    id\n    text\n  }\n}\n","variables":{"input":{"userId":"123424","text":"text"}}}
		//`{"query":"{people(id:\"3\") {\n  name\n  films{title}\n }\n}"}`)
		`{"query":"mutation  createTodo($input: NewTodo!){\n  createTodo(input:$input){\n    done\n    id\n    text\n  }\n}\n","variables":{"input":{"userId":"123424","text":"text"}}}`)

	//{"query":"query{\n  todos{\n    id\n    text\n    done\n  }\n    todo{\n    id\n    text\n    done\n  }\n}","variables":{"input":{"userId":"123424","text":"text"}}}
	/*`
	    query ($key: String!) {
	         people(id:$key) {
	             name
	            films{
	             title
	            }
	        }
	    }
	`)*/

	// set any variables
	//req.Var("key", "12445")
	//req = graphql.NewRequest("")
	fmt.Println(req)
	// set header fields
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	fmt.Println(respData)
}
