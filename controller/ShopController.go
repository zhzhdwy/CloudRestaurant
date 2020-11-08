package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ShopController struct {
}

func (sc *ShopController) Router(engine *gin.Engine) {
	engine.GET("/api/shops", sc.GetShopList)
	engine.GET("/api/search_shops", sc.SearchShop)
}

func (sc *ShopController) SearchShop(context *gin.Context) {
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")
	keyword := context.Query("keyword")

	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34"
		latitude = "40.34"
	}

	if keyword == "" {
		tool.Failed(context, "请输入关键字")
	}

	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude, keyword)
	if len(shops) != 0 {
		tool.Success(context, shops)
		return
	}
	tool.Failed(context, "暂未获得店铺信息")
}

func (sc *ShopController) GetShopList(context *gin.Context) {
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")

	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34"
		latitude = "40.34"
	}
	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude, "")
	if len(shops) == 0 {
		tool.Failed(context, "暂未获得店铺信息")
		return
	}

	for index, shop := range shops {
		ss := shopService.GetService(shop.Id)
		if len(ss) == 0 {
			shops[index].Supports = nil
		} else {
			shops[index].Supports = ss
		}
	}

	fmt.Println(shops)

	tool.Success(context, shops)
}
