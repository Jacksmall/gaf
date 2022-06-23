module github.com/Jacksmall/go-api-framework

go 1.16

require (
	github.com/gin-gonic/gin v1.8.0
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-redis/redis/v8 v8.11.5
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/net v0.0.0-20220531201128-c960675eff93 // indirect
	google.golang.org/genproto v0.0.0-20220602131408-e326c6e8e9c8 // indirect
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.3.4
	gorm.io/gorm v1.23.5
)

replace github.com/Jacksmall/go-api-framework/entry => ../entry
