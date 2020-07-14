$go run client.go
{"data":{"createTodo":{"done":true,"id":"123424","text":"text"}}}

query{
  todos{
    id,
    text,
    done
  },
    todo{
    id,
    text,
    done
  }
}