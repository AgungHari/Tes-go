```bash
mkdir todo-app
cd todo-app
```

```bash
go mod init todo-app
```

```bash
go run main.go
```

```bash
curl -X POST -d '{"task":"melajah go", "status":"konden usan"}' \
-H "Content-Type: application/json" http://localhost:8080/add-todo
```

