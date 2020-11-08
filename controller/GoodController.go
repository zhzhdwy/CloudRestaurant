package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GoodController struct {
}

func (gc *GoodController) Router(engine *gin.Engine) {
	engine.GET("/api/foods", gc.getGoods)
}

func (gc *GoodController) getGoods(context *gin.Context) {
	shopId, exist := context.GetQuery("shop_id")
	if !exist {
		tool.Failed(context, "请求参数错误，请重试")
		return
	}

	id, err := strconv.Atoi(shopId)
	if err != nil {
		tool.Failed(context, "id转换失败")
		return
	}
	goodService := service.NewGoodService()
	goods := goodService.GetFoods(int64(id))
	if len(goods) == 0 {
		tool.Failed(context, "未查询到相关数据")
		return
	}
	tool.Success(context, goods)
}
