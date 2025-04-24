# scanEVMBlockData

1. 启动mongodb数据库 dock容器

> docker run -d --name my-mongodb   -e MONGO_INITDB_ROOT_USERNAME=root   -e MONGO_INITDB_ROOT_PASSWORD=password   -p 27017:27017   mongo:latest

通过 mongoDB Compass 软件连接本地 数据库, 创建数据库： drugtra

2. 启动后端
> go run ./scanblockdata.go

>GOOS=windows GOARCH=amd64 go build -o bin/app-amd64.exe org.go

> http://localhost:3005/




