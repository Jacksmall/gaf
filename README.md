# go-api-framework
快速搭建go-api项目框架，集成了gin,gorm,rabbitmq,redis等插件
#### 1.目录结构
```
app
|----cache    // 缓存
      |--base.go
      |--redis.go
|----controllers // 控制器层
      |--base.go
      |--xxx_controller.go
|----database   // 数据库连接层
      |--base.go
|----entry    // 输入实体层
      |--base.go
      |--xxx.go
|----grpc	// gRPC
      |--client	// 客户端
      	|--client.go
      |--server // 服务器端
      	|--server.go
      |--testguide  // protocol buffer层
      	|--test_guide.proto
	|--test_guide.pb.go
	|--test_guide_grpc.pb.go
|----models   // 模型层
      |--base.go
      |--xxx.go
|----routes   // 路由层
      |--base.go
      |--xxx.go
      |--middleware.go
|----services // 业务服务层
      |--base.go
      |--xxx_service.go
```
#### 2.使用
请参照main.go
```
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Jacksmall/go-api-framework/database"
	"github.com/Jacksmall/go-api-framework/routes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initRouter() {
	// export GIN_MODE=release
	gin.SetMode(gin.ReleaseMode)
	server := &http.Server{
		Addr:           "127.0.0.1:3000",
		Handler:        routes.InitRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err == http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号优雅关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server existing")
}

func initDatabase() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("success contect to mysql database")
}

func main() {
	initDatabase()
	initRouter()

	s := grpc.NewServer()
	lis, _ := net.Listen("tcp", "localhost:50051")
	s.Serve(lis)
}


```
#### 3.命令行test
```
\>go run main.go
success contect to mysql database

\>curl http://127.0.0.1:3000/admin/api/v1/getAll?page=1&limit=5
{"code":0,"data":{"total":7,"list":[{"ID":2,"CreatedAt":"2022-05-28T23:14:36.754+08:00","UpdatedAt":"2022-05-28T23:14:36.754+08:00","DeletedAt":null,"code":"","Price":0},{"ID":3,"CreatedAt":"2022-05-28T23:14:38.171+08:00","UpdatedAt":"2022-05-28T23:14:38.171+08:00","DeletedAt":null,"code":"","Price":0},{"ID":4,"CreatedAt":"2022-05-28T23:14:39.627+08:00","UpdatedAt":"2022-05-28T23:14:39.627+08:00","DeletedAt":null,"code":"","Price":0},{"ID":5,"CreatedAt":"2022-05-28T23:18:31.595+08:00","UpdatedAt":"2022-05-28T23:18:31.595+08:00","DeletedAt":null,"code":"K33","Price":10},{"ID":6,"CreatedAt":"2022-05-28T23:18:31.595+08:00","UpdatedAt":"2022-05-28T23:18:31.595+08:00","DeletedAt":null,"code":"G66","Price":40},{"ID":7,"CreatedAt":"2022-05-30T22:40:44.99+08:00","UpdatedAt":"2022-05-30T22:40:44.99+08:00","DeletedAt":null,"code":"kkks","Price":28},{"ID":9,"CreatedAt":"2022-06-05T23:18:42.189+08:00","UpdatedAt":"2022-06-05T23:18:42.189+08:00","DeletedAt":null,"code":"CKKKK","Price":27}]},"msg":"success"}

#####gRPC
\$>cd grpc
\grpc$>protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative testguide/test_guide.proto

\$>go run grpc/server/server.go
2022/06/06 10:10:56 get feature: latitude:409146138 longitude:-746188906
2022/06/06 10:10:56 get feature:

另开终端
\$>go run grpc/client/client.go
2022/06/06 10:00:01 name:"Berkshire Valley Management Area Trail, Jefferson, NJ, USA"  location:{latitude:409146138  longitude:-746188906}
2022/06/06 10:00:01 location:{}
2022/06/06 10:00:01 Feature: name: "Patriots Path, Mendham, NJ 07945, USA", point:(407838351, -746143763)
2022/06/06 10:00:01 Feature: name: "101 New Jersey 10, Whippany, NJ 07981, USA", point:(408122808, -743999179)
2022/06/06 10:00:01 Feature: name: "U.S. 6, Shohola, PA 18458, USA", point:(413628156, -749015468)
2022/06/06 10:00:01 Feature: name: "5 Conners Road, Kingston, NY 12401, USA", point:(419999544, -740371136)
```
#### 4.持续更新中
      
       
