package tool

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSession(engine *gin.Engine) {
	config := GetConfig().RedisConfig
	store, err := redis.NewStore(10, "tcp", config.Addr+":"+config.Port, config.Password, []byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
	}

	engine.Use(sessions.Sessions("mysession", store))

}

func Setsess(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}

	session.Set(key, value)
	return session.Save()

}

func GetSess(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)

}
