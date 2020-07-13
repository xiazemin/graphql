go get github.com/machinebox/graphql

$curl -g http://localhost:8080/query
{"errors":[{"message":"operation  not found"}],"data":null}

$curl -g http://localhost:8080/query/
<!DOCTYPE html>
<html>
<head>
	<meta charset=utf-8/>
	<meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-s


$go run client.go
{people(id:"3") {
  name
  films{title}}

{"errors":[{"message":"Expected Name, found \u003cEOF\u003e","locations":
[{"line":3,"column":16}],"extensions":{"code":"GRAPHQL_PARSE_FAILED"}}],"data":null}


COPY CURL

$curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json'
 -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080'
  --data-binary '{"query":"{people(id:\"3\") {\n  name\n  films{title}\n }\n}"}' --compressed
{"data":{"people":{"name":"test","films":null}}}


$curl 'http://localhost:8080/query' --data-binary '{"query":"{people(id:\"3\") {\n  name\n  films{title}\n }\n}"}'
-H 'Content-Type: application/json'
{"data":{"people":{"name":"test","films":null}}}

$curl 'http://localhost:8080/query' --data '{"query":"{people(id:\"3\") {\n  name\n  films{title}\n }\n}"}'
-H 'Content-Type: application/json'
{"data":{"people":{"name":"test","films":null}}}

原因格式不对
>>reqquest body -> :{"query":"{\"query\":\"{people(id:\\\"3\\\") {\\n  name\\n  films{title}\\n }\\n}\"}","variables":null}

原因，请求的时候对string又格式化了一回，不合理，直接发原始string就可以了
r, err := http.NewRequest(http.MethodPost, c.endpoint, bytes.NewReader([]byte(req.q)))