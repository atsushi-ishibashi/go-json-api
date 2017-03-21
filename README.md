# go-json-api
http://sgykfjsm.github.io/blog/2016/03/13/golang-json-api-tutorial/
を写経。

## 動かし方
```
glide install
go build
./go-json-api
```
他のプロジェクトで使用する場合は、
```
glide get github.com/atsushi-ishibashi/go-json-api
```
でできるはず

### 確認
Todo一覧 `curl http://localhost:8080/todos`  
Todo詳細 `curl http://localhost:8080/todos/:todo_id`  
Todo作成 `curl -H "Content-Type: application/json" -d '{"name":"New Todo", "due":"2017-03-27T00:00:00Z"}' http://localhost:8080/todos`  
Todo削除 DELETE /todos/:todo_id  `curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/todos/:todo_id`

＊POSTMANからうまくTodo作成のPOSTができない。。

### mysqlの設定
`repo.go` にdatabase設定諸々を(これだとuser:root, pass:, database: gotest)
tableは
```
CREATE TABLE `todos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `completed` tinyint(1) DEFAULT '0',
  `due` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)
```
