go run gqlgen.go init
Exec "go run ./server.go" to start GraphQL server

go run ./server.go


query{
  todos{
    id
    text
    done
  }
}

{
  "data": {
    "todos": [
      {
        "id": "123",
        "text": "test todo",
        "done": true
      }
    ]
  }
}

curl 'http://localhost:8080/query'
-H 'Accept-Encoding: gzip, deflate, br'
-H 'Content-Type: application/json'
-H 'Accept: application/json' -H 'Connection: keep-alive'
-H 'DNT: 1' -H 'Origin: http://localhost:8080'
--data-binary
'{"query":"query{\n  todos{\n    id\n    text\n    done\n  } \n}# Write your query or mutation here\n"}'
--compressed


mutation  createTodo($input: NewTodo!){
  createTodo(input:$input){
    done
    id
    text
  }
}

{
  "input": {
    "userId": "123424",
    "text": "text"
  }
}


{
  "data": {
    "createTodo": {
      "done": true,
      "id": "123424",
      "text": "text"
    }
  }
}


curl 'http://localhost:8080/query'
-H 'Accept-Encoding: gzip, deflate, br'
-H 'Content-Type: application/json'
-H 'Accept: application/json'
-H 'Connection: keep-alive'
-H 'DNT: 1'
-H 'Origin: http://localhost:8080'
--data-binary
'{"query":"mutation  createTodo($input: NewTodo!){\n  createTodo(input:$input){\n    done\n    id\n    text\n  }\n}\n","variables":{"input":{"userId":"123424","text":"text"}}}' --compressed


https://www.cnblogs.com/rongfengliang/p/9364783.html?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com
https://github.com/rongfengliang/gqlgen-demo

https://github.com/rongfengliang/gqlgen-demo



query{
  todos{
    id
    text
    done
  }
    todo{
    id
    text
    done
  }
}


{
  "data": {
    "todos": [
      {
        "id": "123",
        "text": "test todo",
        "done": true
      }
    ],
    "todo": {
      "id": "1234",
      "text": "to do func",
      "done": true
    }
  }
}


curl 'http://localhost:8080/query'
-H 'Accept-Encoding: gzip, deflate, br'
-H 'Content-Type: application/json'
-H 'Accept: application/json'
-H 'Connection: keep-alive'
-H 'DNT: 1'
-H 'Origin: http://localhost:8080'
--data-binary '{"query":"query{\n  todos{\n    id\n    text\n    done\n  }\n    todo{\n    id\n    text\n    done\n  }\n}","variables":{"input":{"userId":"123424","text":"text"}}}' --compressed