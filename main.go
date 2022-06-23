package main

import (
	"context"
	"fmt"
	"github.com/Jacksmall/go-api-framework/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Jacksmall/go-api-framework/database"
	"github.com/Jacksmall/go-api-framework/routes"
	"github.com/gin-gonic/gin"
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
	quit := make(chan os.Signal, 1)
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
	dsn := "root:chenkuanwo5@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	models.SetDB(database.DBConn)
	fmt.Println("success contect to mysql database")
}

func main() {
	initDatabase()

	database.DBConn.AutoMigrate(&models.Goods{})

	initRouter()

	// s := grpc.NewServer()
	// lis, _ := net.Listen("tcp", "localhost:50051")
	// s.Serve(lis)
}
