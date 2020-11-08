package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"strings"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}

	//实例化mysql数据库
	err = tool.OrmEngine(cfg)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	// 实例化redis
	tool.InitRedisStore()

	app := gin.Default()
	registerRouter(app)

	// 设置全局跨域访问
	app.Use(Cors())

	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
	new(controller.FoodCategoryController).Router(router)
	new(controller.ShopController).Router(router)
	new(controller.GoodController).Router(router)
}

// cross origin resource share
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, "+
				"Token, session")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
				"Access-Control-Allow-Methods")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		context.Next()
	}
}
